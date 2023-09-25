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

func TestOnPublicEventAny(t *testing.T) {
	type args struct {
		callbacks []PublicEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single PublicEventHandleFunc",
			args: args{
				[]PublicEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PublicEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple PublicEventHandleFuncs",
			args: args{
				[]PublicEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PublicEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.PublicEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnPublicEventAny(tt.args.callbacks...)
			if len(g.onPublicEvent[PublicEventAnyAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onPublicEvent[PublicEventAnyAction]))
			}
		})
	}
}

func TestSetOnPublicEventAny(t *testing.T) {
	type args struct {
		callbacks []PublicEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single PublicEventHandleFunc",
			args: args{
				[]PublicEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PublicEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple PublicEventHandleFuncs",
			args: args{
				[]PublicEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PublicEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.PublicEvent) error {
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
			g.SetOnPublicEventAny(func(deliveryID string, eventName string, event *github.PublicEvent) error {
				return nil
			})
			g.SetOnPublicEventAny(tt.args.callbacks...)
			if len(g.onPublicEvent[PublicEventAnyAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onPublicEvent[PublicEventAnyAction]), tt.want)
			}
		})
	}
}

func TestHandlePublicEventAny(t *testing.T) {

	type args struct {
		deliveryID string
		eventName  string
		event      *github.PublicEvent
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
				eventName:  "public",

				event: &github.PublicEvent{},

				fail: false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "public",

				event: &github.PublicEvent{},

				fail: true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "public",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnPublicEventAny(func(deliveryID string, eventName string, event *github.PublicEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handlePublicEventAny(tt.args.deliveryID, tt.args.deliveryID, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("TestHandlePublicEventAny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPublicEvent(t *testing.T) {
	type fields struct {
		handler *EventHandler
	}
	type args struct {
		deliveryID string
		eventName  string
		event      *github.PublicEvent
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "must trigger PublicEventAny with unknown event action",
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
					onPublicEvent: map[string][]PublicEventHandleFunc{
						PublicEventAnyAction: {
							func(deliveryID string, eventName string, event *github.PublicEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  PublicEvent,

				event: &github.PublicEvent{},
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
			if err := g.PublicEvent(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("PublicEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
