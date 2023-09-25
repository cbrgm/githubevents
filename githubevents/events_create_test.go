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
			if len(g.onCreateEvent[CreateEventAnyAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onCreateEvent[CreateEventAnyAction]))
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
			if len(g.onCreateEvent[CreateEventAnyAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onCreateEvent[CreateEventAnyAction]), tt.want)
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

func TestCreateEvent(t *testing.T) {
	type fields struct {
		handler *EventHandler
	}
	type args struct {
		deliveryID string
		eventName  string
		event      *github.CreateEvent
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "must trigger CreateEventAny with unknown event action",
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
					onCreateEvent: map[string][]CreateEventHandleFunc{
						CreateEventAnyAction: {
							func(deliveryID string, eventName string, event *github.CreateEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  CreateEvent,

				event: &github.CreateEvent{},
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
			if err := g.CreateEvent(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("CreateEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
