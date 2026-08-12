package githubevents

import (
	"context"
	"errors"
	"strings"
	"sync/atomic"
	"testing"

	"github.com/google/go-github/v89/github"
)

func labelEvent(action string) *github.LabelEvent {
	return &github.LabelEvent{Action: github.Ptr(action)}
}

// panic in one callback surfaces as a recovered-from-panic error.
func TestDispatchRecoversPanic(t *testing.T) {
	g := New("")
	g.OnLabelEventCreated(func(ctx context.Context, id, name string, e *github.LabelEvent) error {
		panic("boom")
	})
	err := g.LabelEvent(context.Background(), "id", "label", labelEvent("created"))
	if err == nil || !strings.Contains(err.Error(), "recovered from panic:") {
		t.Fatalf("want recovered-from-panic error, got %v", err)
	}
}

// every registered callback runs (parallel fan-out).
func TestDispatchRunsAll(t *testing.T) {
	g := New("")
	var n atomic.Int32
	inc := func(ctx context.Context, id, name string, e *github.LabelEvent) error {
		n.Add(1)
		return nil
	}
	g.OnLabelEventCreated(inc, inc, inc)
	if err := g.LabelEvent(context.Background(), "id", "label", labelEvent("created")); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if n.Load() != 3 {
		t.Fatalf("want 3 callbacks run, got %d", n.Load())
	}
}

// a returned error propagates.
func TestDispatchPropagatesError(t *testing.T) {
	g := New("")
	want := errors.New("nope")
	g.OnLabelEventCreated(func(ctx context.Context, id, name string, e *github.LabelEvent) error {
		return want
	})
	if err := g.LabelEvent(context.Background(), "id", "label", labelEvent("created")); err == nil {
		t.Fatal("want error, got nil")
	}
}

// wrong / empty action is rejected by the action handler contract.
func TestDispatchRejectsEmptyAction(t *testing.T) {
	g := New("")
	err := g.LabelEvent(context.Background(), "id", "label", &github.LabelEvent{})
	if err == nil || !strings.Contains(err.Error(), "empty or nil") {
		t.Fatalf("want empty-action error, got %v", err)
	}
}

// the Any path (event without a matching action handler) still fans out.
func TestDispatchAnyPath(t *testing.T) {
	g := New("")
	var n atomic.Int32
	g.OnLabelEventAny(func(ctx context.Context, id, name string, e *github.LabelEvent) error {
		n.Add(1)
		return nil
	})
	if err := g.LabelEvent(context.Background(), "id", "label", labelEvent("created")); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if n.Load() != 1 {
		t.Fatalf("want Any callback to run once, got %d", n.Load())
	}
}
