package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"errors"
	"github.com/google/go-github/v43/github"
	"testing"
)

func TestOnTeamAddEventAny(t *testing.T) {
	type args struct {
		callbacks []TeamAddEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single TeamAddEventHandleFunc",
			args: args{
				[]TeamAddEventHandleFunc{
					func(deliveryID string, eventName string, event *github.TeamAddEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple TeamAddEventHandleFuncs",
			args: args{
				[]TeamAddEventHandleFunc{
					func(deliveryID string, eventName string, event *github.TeamAddEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.TeamAddEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnTeamAddEventAny(tt.args.callbacks...)
			if len(g.onTeamAddEvent["*"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onTeamAddEvent["*"]))
			}
		})
	}
}

func TestSetOnTeamAddEventAny(t *testing.T) {
	type args struct {
		callbacks []TeamAddEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single TeamAddEventHandleFunc",
			args: args{
				[]TeamAddEventHandleFunc{
					func(deliveryID string, eventName string, event *github.TeamAddEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple TeamAddEventHandleFuncs",
			args: args{
				[]TeamAddEventHandleFunc{
					func(deliveryID string, eventName string, event *github.TeamAddEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.TeamAddEvent) error {
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
			g.SetOnTeamAddEventAny(func(deliveryID string, eventName string, event *github.TeamAddEvent) error {
				return nil
			})
			g.SetOnTeamAddEventAny(tt.args.callbacks...)
			if len(g.onTeamAddEvent["*"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onTeamAddEvent["*"]), tt.want)
			}
		})
	}
}

func TestHandleTeamAddEventAny(t *testing.T) {

	type args struct {
		deliveryID string
		eventName  string
		event      *github.TeamAddEvent
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
				eventName:  "team_add",

				event: &github.TeamAddEvent{},

				fail: false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "team_add",

				event: &github.TeamAddEvent{},

				fail: true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "team_add",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnTeamAddEventAny(func(deliveryID string, eventName string, event *github.TeamAddEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleTeamAddEventAny(tt.args.deliveryID, tt.args.deliveryID, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("TestHandleTeamAddEventAny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
