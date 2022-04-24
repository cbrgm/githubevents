package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"errors"
	"github.com/google/go-github/v43/github"
	"testing"
)

func TestOnCreateEventAny(t *testing.T) {
	type args struct {
		callbacks []CreateEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single CreateEventHandleFunc",
			args: args{
				[]CreateEventHandleFunc{
					func(deliveryID string, eventName string, event *github.CreateEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple CreateEventHandleFuncs",
			args: args{
				[]CreateEventHandleFunc{
					func(deliveryID string, eventName string, event *github.CreateEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.CreateEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnCreateEventAny(tt.args.callbacks...)
			if len(g.onCreateEvent["*"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onCreateEvent["*"]))
			}
		})
	}
}

func TestSetOnCreateEventAny(t *testing.T) {
	type args struct {
		callbacks []CreateEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single CreateEventHandleFunc",
			args: args{
				[]CreateEventHandleFunc{
					func(deliveryID string, eventName string, event *github.CreateEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple CreateEventHandleFuncs",
			args: args{
				[]CreateEventHandleFunc{
					func(deliveryID string, eventName string, event *github.CreateEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.CreateEvent) error {
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
			g.SetOnCreateEventAny(func(deliveryID string, eventName string, event *github.CreateEvent) error {
				return nil
			})
			g.SetOnCreateEventAny(tt.args.callbacks...)
			if len(g.onCreateEvent["*"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onCreateEvent["*"]), tt.want)
			}
		})
	}
}

func TestHandleCreateEventAny(t *testing.T) {

	type args struct {
		deliveryID string
		eventName  string
		event      *github.CreateEvent
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
				eventName:  "create",

				event: &github.CreateEvent{},

				fail: false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "create",

				event: &github.CreateEvent{},

				fail: true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "create",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnCreateEventAny(func(deliveryID string, eventName string, event *github.CreateEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleCreateEventAny(tt.args.deliveryID, tt.args.deliveryID, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("TestHandleCreateEventAny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
