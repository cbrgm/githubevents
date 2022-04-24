package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"errors"
	"github.com/google/go-github/v43/github"
	"testing"
)

func TestOnWatchEventAny(t *testing.T) {
	type args struct {
		callbacks []WatchEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single WatchEventHandleFunc",
			args: args{
				[]WatchEventHandleFunc{
					func(deliveryID string, eventName string, event *github.WatchEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple WatchEventHandleFuncs",
			args: args{
				[]WatchEventHandleFunc{
					func(deliveryID string, eventName string, event *github.WatchEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.WatchEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnWatchEventAny(tt.args.callbacks...)
			if len(g.onWatchEvent["*"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onWatchEvent["*"]))
			}
		})
	}
}

func TestSetOnWatchEventAny(t *testing.T) {
	type args struct {
		callbacks []WatchEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single WatchEventHandleFunc",
			args: args{
				[]WatchEventHandleFunc{
					func(deliveryID string, eventName string, event *github.WatchEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple WatchEventHandleFuncs",
			args: args{
				[]WatchEventHandleFunc{
					func(deliveryID string, eventName string, event *github.WatchEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.WatchEvent) error {
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
			g.SetOnWatchEventAny(func(deliveryID string, eventName string, event *github.WatchEvent) error {
				return nil
			})
			g.SetOnWatchEventAny(tt.args.callbacks...)
			if len(g.onWatchEvent["*"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onWatchEvent["*"]), tt.want)
			}
		})
	}
}

func TestHandleWatchEventAny(t *testing.T) {

	type args struct {
		deliveryID string
		eventName  string
		event      *github.WatchEvent
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
				eventName:  "watch",

				event: &github.WatchEvent{},

				fail: false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "watch",

				event: &github.WatchEvent{},

				fail: true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "watch",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnWatchEventAny(func(deliveryID string, eventName string, event *github.WatchEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleWatchEventAny(tt.args.deliveryID, tt.args.deliveryID, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("TestHandleWatchEventAny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
