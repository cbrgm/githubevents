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

func TestOnForkEventAny(t *testing.T) {
	type args struct {
		callbacks []ForkEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single ForkEventHandleFunc",
			args: args{
				[]ForkEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ForkEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple ForkEventHandleFuncs",
			args: args{
				[]ForkEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ForkEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ForkEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnForkEventAny(tt.args.callbacks...)
			if len(g.onForkEvent[ForkEventAnyAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onForkEvent[ForkEventAnyAction]))
			}
		})
	}
}

func TestSetOnForkEventAny(t *testing.T) {
	type args struct {
		callbacks []ForkEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single ForkEventHandleFunc",
			args: args{
				[]ForkEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ForkEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple ForkEventHandleFuncs",
			args: args{
				[]ForkEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ForkEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ForkEvent) error {
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
			g.SetOnForkEventAny(func(deliveryID string, eventName string, event *github.ForkEvent) error {
				return nil
			})
			g.SetOnForkEventAny(tt.args.callbacks...)
			if len(g.onForkEvent[ForkEventAnyAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onForkEvent[ForkEventAnyAction]), tt.want)
			}
		})
	}
}

func TestHandleForkEventAny(t *testing.T) {

	type args struct {
		deliveryID string
		eventName  string
		event      *github.ForkEvent
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
				eventName:  "fork",

				event: &github.ForkEvent{},

				fail: false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "fork",

				event: &github.ForkEvent{},

				fail: true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "fork",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnForkEventAny(func(deliveryID string, eventName string, event *github.ForkEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleForkEventAny(tt.args.deliveryID, tt.args.deliveryID, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("TestHandleForkEventAny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestForkEvent(t *testing.T) {
	type fields struct {
		handler *EventHandler
	}
	type args struct {
		deliveryID string
		eventName  string
		event      *github.ForkEvent
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "must trigger ForkEventAny with unknown event action",
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
					onForkEvent: map[string][]ForkEventHandleFunc{
						ForkEventAnyAction: {
							func(deliveryID string, eventName string, event *github.ForkEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  ForkEvent,

				event: &github.ForkEvent{},
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
			if err := g.ForkEvent(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("ForkEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
