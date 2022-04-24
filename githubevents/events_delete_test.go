package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"errors"
	"github.com/google/go-github/v43/github"
	"testing"
)

func TestOnDeleteEventAny(t *testing.T) {
	type args struct {
		callbacks []DeleteEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single DeleteEventHandleFunc",
			args: args{
				[]DeleteEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DeleteEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple DeleteEventHandleFuncs",
			args: args{
				[]DeleteEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DeleteEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.DeleteEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnDeleteEventAny(tt.args.callbacks...)
			if len(g.onDeleteEvent["*"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onDeleteEvent["*"]))
			}
		})
	}
}

func TestSetOnDeleteEventAny(t *testing.T) {
	type args struct {
		callbacks []DeleteEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single DeleteEventHandleFunc",
			args: args{
				[]DeleteEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DeleteEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple DeleteEventHandleFuncs",
			args: args{
				[]DeleteEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DeleteEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.DeleteEvent) error {
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
			g.SetOnDeleteEventAny(func(deliveryID string, eventName string, event *github.DeleteEvent) error {
				return nil
			})
			g.SetOnDeleteEventAny(tt.args.callbacks...)
			if len(g.onDeleteEvent["*"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onDeleteEvent["*"]), tt.want)
			}
		})
	}
}

func TestHandleDeleteEventAny(t *testing.T) {

	type args struct {
		deliveryID string
		eventName  string
		event      *github.DeleteEvent
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
				eventName:  "delete",

				event: &github.DeleteEvent{},

				fail: false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "delete",

				event: &github.DeleteEvent{},

				fail: true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "delete",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnDeleteEventAny(func(deliveryID string, eventName string, event *github.DeleteEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleDeleteEventAny(tt.args.deliveryID, tt.args.deliveryID, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("TestHandleDeleteEventAny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
