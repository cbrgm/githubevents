package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"errors"
	"github.com/google/go-github/v43/github"
	"testing"
)

func TestOnStatusEventAny(t *testing.T) {
	type args struct {
		callbacks []StatusEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single StatusEventHandleFunc",
			args: args{
				[]StatusEventHandleFunc{
					func(deliveryID string, eventName string, event *github.StatusEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple StatusEventHandleFuncs",
			args: args{
				[]StatusEventHandleFunc{
					func(deliveryID string, eventName string, event *github.StatusEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.StatusEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnStatusEventAny(tt.args.callbacks...)
			if len(g.onStatusEvent["*"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onStatusEvent["*"]))
			}
		})
	}
}

func TestSetOnStatusEventAny(t *testing.T) {
	type args struct {
		callbacks []StatusEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single StatusEventHandleFunc",
			args: args{
				[]StatusEventHandleFunc{
					func(deliveryID string, eventName string, event *github.StatusEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple StatusEventHandleFuncs",
			args: args{
				[]StatusEventHandleFunc{
					func(deliveryID string, eventName string, event *github.StatusEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.StatusEvent) error {
						return nil
					},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			// add callbacks to be overwritten
			g.SetOnStatusEventAny(func(deliveryID string, eventName string, event *github.StatusEvent) error {
				return nil
			})
			g.SetOnStatusEventAny(tt.args.callbacks...)
			if len(g.onStatusEvent["*"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onStatusEvent["*"]), tt.want)
			}
		})
	}
}

func TestHandleStatusEventAny(t *testing.T) {

	type args struct {
		deliveryID string
		eventName  string
		event      *github.StatusEvent
		fail       bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "must pass",
			args: args{
				deliveryID: "42",
				eventName:  "status",

				event: &github.StatusEvent{},

				fail: false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "status",

				event: &github.StatusEvent{},

				fail: true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "status",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnStatusEventAny(func(deliveryID string, eventName string, event *github.StatusEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleStatusEventAny(tt.args.deliveryID, tt.args.deliveryID, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("TestHandleStatusEventAny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
