// Copyright 2022 The GithubEvents Authors. All rights reserved.
// Use of this source code is governed by the MIT License
// that can be found in the LICENSE file.

package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"errors"
	"github.com/google/go-github/v55/github"
	"sync"
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
			if len(g.onTeamAddEvent[TeamAddEventAnyAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onTeamAddEvent[TeamAddEventAnyAction]))
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
			if len(g.onTeamAddEvent[TeamAddEventAnyAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onTeamAddEvent[TeamAddEventAnyAction]), tt.want)
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

func TestTeamAddEvent(t *testing.T) {
	type fields struct {
		handler *EventHandler
	}
	type args struct {
		deliveryID string
		eventName  string
		event      *github.TeamAddEvent
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "must trigger TeamAddEventAny with unknown event action",
			fields: fields{
				handler: &EventHandler{
					WebhookSecret: "fake",
					onBeforeAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(deliveryID string, eventName string, event any) error {
								t.Log("onBeforeAny called")
								return nil
							},
						},
					},
					onAfterAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(deliveryID string, eventName string, event any) error {
								t.Log("onAfterAny called")
								return nil
							},
						},
					},
					onTeamAddEvent: map[string][]TeamAddEventHandleFunc{
						TeamAddEventAnyAction: {
							func(deliveryID string, eventName string, event *github.TeamAddEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  TeamAddEvent,

				event: &github.TeamAddEvent{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &EventHandler{
				WebhookSecret: "fake",
				mu:            sync.RWMutex{},
			}
			if err := g.TeamAddEvent(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("TeamAddEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
