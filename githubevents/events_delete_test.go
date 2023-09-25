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
			if len(g.onDeleteEvent[DeleteEventAnyAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onDeleteEvent[DeleteEventAnyAction]))
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
			if len(g.onDeleteEvent[DeleteEventAnyAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onDeleteEvent[DeleteEventAnyAction]), tt.want)
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

func TestDeleteEvent(t *testing.T) {
	type fields struct {
		handler *EventHandler
	}
	type args struct {
		deliveryID string
		eventName  string
		event      *github.DeleteEvent
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "must trigger DeleteEventAny with unknown event action",
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
					onDeleteEvent: map[string][]DeleteEventHandleFunc{
						DeleteEventAnyAction: {
							func(deliveryID string, eventName string, event *github.DeleteEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  DeleteEvent,

				event: &github.DeleteEvent{},
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
			if err := g.DeleteEvent(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("DeleteEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
