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

func TestOnRepositoryDispatchEventAny(t *testing.T) {
	type args struct {
		callbacks []RepositoryDispatchEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single RepositoryDispatchEventHandleFunc",
			args: args{
				[]RepositoryDispatchEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryDispatchEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple RepositoryDispatchEventHandleFuncs",
			args: args{
				[]RepositoryDispatchEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryDispatchEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.RepositoryDispatchEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnRepositoryDispatchEventAny(tt.args.callbacks...)
			if len(g.onRepositoryDispatchEvent[RepositoryDispatchEventAnyAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onRepositoryDispatchEvent[RepositoryDispatchEventAnyAction]))
			}
		})
	}
}

func TestSetOnRepositoryDispatchEventAny(t *testing.T) {
	type args struct {
		callbacks []RepositoryDispatchEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single RepositoryDispatchEventHandleFunc",
			args: args{
				[]RepositoryDispatchEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryDispatchEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple RepositoryDispatchEventHandleFuncs",
			args: args{
				[]RepositoryDispatchEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryDispatchEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.RepositoryDispatchEvent) error {
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
			g.SetOnRepositoryDispatchEventAny(func(deliveryID string, eventName string, event *github.RepositoryDispatchEvent) error {
				return nil
			})
			g.SetOnRepositoryDispatchEventAny(tt.args.callbacks...)
			if len(g.onRepositoryDispatchEvent[RepositoryDispatchEventAnyAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onRepositoryDispatchEvent[RepositoryDispatchEventAnyAction]), tt.want)
			}
		})
	}
}

func TestHandleRepositoryDispatchEventAny(t *testing.T) {

	type args struct {
		deliveryID string
		eventName  string
		event      *github.RepositoryDispatchEvent
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
				eventName:  "repository_dispatch",

				event: &github.RepositoryDispatchEvent{},

				fail: false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "repository_dispatch",

				event: &github.RepositoryDispatchEvent{},

				fail: true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "repository_dispatch",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnRepositoryDispatchEventAny(func(deliveryID string, eventName string, event *github.RepositoryDispatchEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleRepositoryDispatchEventAny(tt.args.deliveryID, tt.args.deliveryID, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("TestHandleRepositoryDispatchEventAny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRepositoryDispatchEvent(t *testing.T) {
	type fields struct {
		handler *EventHandler
	}
	type args struct {
		deliveryID string
		eventName  string
		event      *github.RepositoryDispatchEvent
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "must trigger RepositoryDispatchEventAny with unknown event action",
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
					onRepositoryDispatchEvent: map[string][]RepositoryDispatchEventHandleFunc{
						RepositoryDispatchEventAnyAction: {
							func(deliveryID string, eventName string, event *github.RepositoryDispatchEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  RepositoryDispatchEvent,

				event: &github.RepositoryDispatchEvent{},
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
			if err := g.RepositoryDispatchEvent(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("RepositoryDispatchEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
