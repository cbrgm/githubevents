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

func TestOnGollumEventAny(t *testing.T) {
	type args struct {
		callbacks []GollumEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single GollumEventHandleFunc",
			args: args{
				[]GollumEventHandleFunc{
					func(deliveryID string, eventName string, event *github.GollumEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple GollumEventHandleFuncs",
			args: args{
				[]GollumEventHandleFunc{
					func(deliveryID string, eventName string, event *github.GollumEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.GollumEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnGollumEventAny(tt.args.callbacks...)
			if len(g.onGollumEvent[GollumEventAnyAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onGollumEvent[GollumEventAnyAction]))
			}
		})
	}
}

func TestSetOnGollumEventAny(t *testing.T) {
	type args struct {
		callbacks []GollumEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single GollumEventHandleFunc",
			args: args{
				[]GollumEventHandleFunc{
					func(deliveryID string, eventName string, event *github.GollumEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple GollumEventHandleFuncs",
			args: args{
				[]GollumEventHandleFunc{
					func(deliveryID string, eventName string, event *github.GollumEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.GollumEvent) error {
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
			g.SetOnGollumEventAny(func(deliveryID string, eventName string, event *github.GollumEvent) error {
				return nil
			})
			g.SetOnGollumEventAny(tt.args.callbacks...)
			if len(g.onGollumEvent[GollumEventAnyAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onGollumEvent[GollumEventAnyAction]), tt.want)
			}
		})
	}
}

func TestHandleGollumEventAny(t *testing.T) {

	type args struct {
		deliveryID string
		eventName  string
		event      *github.GollumEvent
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
				eventName:  "gollum",

				event: &github.GollumEvent{},

				fail: false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "gollum",

				event: &github.GollumEvent{},

				fail: true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "gollum",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnGollumEventAny(func(deliveryID string, eventName string, event *github.GollumEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleGollumEventAny(tt.args.deliveryID, tt.args.deliveryID, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("TestHandleGollumEventAny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGollumEvent(t *testing.T) {
	type fields struct {
		handler *EventHandler
	}
	type args struct {
		deliveryID string
		eventName  string
		event      *github.GollumEvent
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "must trigger GollumEventAny with unknown event action",
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
					onGollumEvent: map[string][]GollumEventHandleFunc{
						GollumEventAnyAction: {
							func(deliveryID string, eventName string, event *github.GollumEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  GollumEvent,

				event: &github.GollumEvent{},
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
			if err := g.GollumEvent(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("GollumEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
