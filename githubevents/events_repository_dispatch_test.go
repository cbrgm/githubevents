package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"errors"
	"github.com/google/go-github/v43/github"
	"testing"
)

func TestOnRepositoryDispatchEventAny(t *testing.T) {
	type args struct {
		callbacks []RepositoryDispatchEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single RepositoryDispatchEventHandleFunc",
			args: args{
				[]RepositoryDispatchEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryDispatchEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple RepositoryDispatchEventHandleFuncs",
			args: args{
				[]RepositoryDispatchEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryDispatchEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.RepositoryDispatchEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnRepositoryDispatchEventAny(tt.args.callbacks...)
			if len(g.onRepositoryDispatchEvent["*"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onRepositoryDispatchEvent["*"]))
			}
		})
	}
}

func TestSetOnRepositoryDispatchEventAny(t *testing.T) {
	type args struct {
		callbacks []RepositoryDispatchEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single RepositoryDispatchEventHandleFunc",
			args: args{
				[]RepositoryDispatchEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryDispatchEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple RepositoryDispatchEventHandleFuncs",
			args: args{
				[]RepositoryDispatchEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryDispatchEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.RepositoryDispatchEvent) error {
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
			g.SetOnRepositoryDispatchEventAny(func(deliveryID string, eventName string, event *github.RepositoryDispatchEvent) error {
				return nil
			})
			g.SetOnRepositoryDispatchEventAny(tt.args.callbacks...)
			if len(g.onRepositoryDispatchEvent["*"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onRepositoryDispatchEvent["*"]), tt.want)
			}
		})
	}
}

func TestHandleRepositoryDispatchEventAny(t *testing.T) {

	type args struct {
		deliveryID string
		eventName  string
		event      *github.RepositoryDispatchEvent
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
				eventName:  "repository_dispatch",

				event: &github.RepositoryDispatchEvent{},

				fail: false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "repository_dispatch",

				event: &github.RepositoryDispatchEvent{},

				fail: true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "repository_dispatch",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnRepositoryDispatchEventAny(func(deliveryID string, eventName string, event *github.RepositoryDispatchEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleRepositoryDispatchEventAny(tt.args.deliveryID, tt.args.deliveryID, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("TestHandleRepositoryDispatchEventAny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
