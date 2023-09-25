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

func TestOnMetaEventAny(t *testing.T) {
	type args struct {
		callbacks []MetaEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single MetaEventHandleFunc",
			args: args{
				[]MetaEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MetaEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple MetaEventHandleFuncs",
			args: args{
				[]MetaEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MetaEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.MetaEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnMetaEventAny(tt.args.callbacks...)
			if len(g.onMetaEvent[MetaEventAnyAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onMetaEvent[MetaEventAnyAction]))
			}
		})
	}
}

func TestSetOnMetaEventAny(t *testing.T) {
	type args struct {
		callbacks []MetaEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single MetaEventHandleFunc",
			args: args{
				[]MetaEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MetaEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple MetaEventHandleFuncs",
			args: args{
				[]MetaEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MetaEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.MetaEvent) error {
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
			g.SetOnMetaEventAny(func(deliveryID string, eventName string, event *github.MetaEvent) error {
				return nil
			})
			g.SetOnMetaEventAny(tt.args.callbacks...)
			if len(g.onMetaEvent[MetaEventAnyAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onMetaEvent[MetaEventAnyAction]), tt.want)
			}
		})
	}
}

func TestHandleMetaEventAny(t *testing.T) {

	type args struct {
		deliveryID string
		eventName  string
		event      *github.MetaEvent
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
				eventName:  "meta",

				event: &github.MetaEvent{},

				fail: false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "meta",

				event: &github.MetaEvent{},

				fail: true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "meta",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnMetaEventAny(func(deliveryID string, eventName string, event *github.MetaEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleMetaEventAny(tt.args.deliveryID, tt.args.deliveryID, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("TestHandleMetaEventAny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMetaEvent(t *testing.T) {
	type fields struct {
		handler *EventHandler
	}
	type args struct {
		deliveryID string
		eventName  string
		event      *github.MetaEvent
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "must trigger MetaEventAny with unknown event action",
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
					onMetaEvent: map[string][]MetaEventHandleFunc{
						MetaEventAnyAction: {
							func(deliveryID string, eventName string, event *github.MetaEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  MetaEvent,

				event: &github.MetaEvent{},
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
			if err := g.MetaEvent(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("MetaEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
