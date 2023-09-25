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
			if len(g.onWatchEvent[WatchEventAnyAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onWatchEvent[WatchEventAnyAction]))
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
			if len(g.onWatchEvent[WatchEventAnyAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onWatchEvent[WatchEventAnyAction]), tt.want)
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

func TestWatchEvent(t *testing.T) {
	type fields struct {
		handler *EventHandler
	}
	type args struct {
		deliveryID string
		eventName  string
		event      *github.WatchEvent
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "must trigger WatchEventAny with unknown event action",
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
					onWatchEvent: map[string][]WatchEventHandleFunc{
						WatchEventAnyAction: {
							func(deliveryID string, eventName string, event *github.WatchEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  WatchEvent,

				event: &github.WatchEvent{},
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
			if err := g.WatchEvent(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("WatchEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
