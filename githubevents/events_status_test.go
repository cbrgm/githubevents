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

func TestOnStatusEventAny(t *testing.T) {
	type args struct {
		callbacks []StatusEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single StatusEventHandleFunc",
			args: args{
				[]StatusEventHandleFunc{
					func(deliveryID string, eventName string, event *github.StatusEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple StatusEventHandleFuncs",
			args: args{
				[]StatusEventHandleFunc{
					func(deliveryID string, eventName string, event *github.StatusEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.StatusEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnStatusEventAny(tt.args.callbacks...)
			if len(g.onStatusEvent[StatusEventAnyAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onStatusEvent[StatusEventAnyAction]))
			}
		})
	}
}

func TestSetOnStatusEventAny(t *testing.T) {
	type args struct {
		callbacks []StatusEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single StatusEventHandleFunc",
			args: args{
				[]StatusEventHandleFunc{
					func(deliveryID string, eventName string, event *github.StatusEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple StatusEventHandleFuncs",
			args: args{
				[]StatusEventHandleFunc{
					func(deliveryID string, eventName string, event *github.StatusEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.StatusEvent) error {
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
			g.SetOnStatusEventAny(func(deliveryID string, eventName string, event *github.StatusEvent) error {
				return nil
			})
			g.SetOnStatusEventAny(tt.args.callbacks...)
			if len(g.onStatusEvent[StatusEventAnyAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onStatusEvent[StatusEventAnyAction]), tt.want)
			}
		})
	}
}

func TestHandleStatusEventAny(t *testing.T) {

	type args struct {
		deliveryID string
		eventName  string
		event      *github.StatusEvent
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
				eventName:  "status",

				event: &github.StatusEvent{},

				fail: false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "status",

				event: &github.StatusEvent{},

				fail: true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "status",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnStatusEventAny(func(deliveryID string, eventName string, event *github.StatusEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleStatusEventAny(tt.args.deliveryID, tt.args.deliveryID, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("TestHandleStatusEventAny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStatusEvent(t *testing.T) {
	type fields struct {
		handler *EventHandler
	}
	type args struct {
		deliveryID string
		eventName  string
		event      *github.StatusEvent
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "must trigger StatusEventAny with unknown event action",
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
					onStatusEvent: map[string][]StatusEventHandleFunc{
						StatusEventAnyAction: {
							func(deliveryID string, eventName string, event *github.StatusEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  StatusEvent,

				event: &github.StatusEvent{},
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
			if err := g.StatusEvent(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("StatusEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
