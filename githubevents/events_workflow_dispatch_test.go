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

func TestOnWorkflowDispatchEventAny(t *testing.T) {
	type args struct {
		callbacks []WorkflowDispatchEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single WorkflowDispatchEventHandleFunc",
			args: args{
				[]WorkflowDispatchEventHandleFunc{
					func(deliveryID string, eventName string, event *github.WorkflowDispatchEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple WorkflowDispatchEventHandleFuncs",
			args: args{
				[]WorkflowDispatchEventHandleFunc{
					func(deliveryID string, eventName string, event *github.WorkflowDispatchEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.WorkflowDispatchEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnWorkflowDispatchEventAny(tt.args.callbacks...)
			if len(g.onWorkflowDispatchEvent[WorkflowDispatchEventAnyAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onWorkflowDispatchEvent[WorkflowDispatchEventAnyAction]))
			}
		})
	}
}

func TestSetOnWorkflowDispatchEventAny(t *testing.T) {
	type args struct {
		callbacks []WorkflowDispatchEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single WorkflowDispatchEventHandleFunc",
			args: args{
				[]WorkflowDispatchEventHandleFunc{
					func(deliveryID string, eventName string, event *github.WorkflowDispatchEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple WorkflowDispatchEventHandleFuncs",
			args: args{
				[]WorkflowDispatchEventHandleFunc{
					func(deliveryID string, eventName string, event *github.WorkflowDispatchEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.WorkflowDispatchEvent) error {
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
			g.SetOnWorkflowDispatchEventAny(func(deliveryID string, eventName string, event *github.WorkflowDispatchEvent) error {
				return nil
			})
			g.SetOnWorkflowDispatchEventAny(tt.args.callbacks...)
			if len(g.onWorkflowDispatchEvent[WorkflowDispatchEventAnyAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onWorkflowDispatchEvent[WorkflowDispatchEventAnyAction]), tt.want)
			}
		})
	}
}

func TestHandleWorkflowDispatchEventAny(t *testing.T) {

	type args struct {
		deliveryID string
		eventName  string
		event      *github.WorkflowDispatchEvent
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
				eventName:  "workflow_dispatch",

				event: &github.WorkflowDispatchEvent{},

				fail: false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "workflow_dispatch",

				event: &github.WorkflowDispatchEvent{},

				fail: true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "workflow_dispatch",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnWorkflowDispatchEventAny(func(deliveryID string, eventName string, event *github.WorkflowDispatchEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleWorkflowDispatchEventAny(tt.args.deliveryID, tt.args.deliveryID, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("TestHandleWorkflowDispatchEventAny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWorkflowDispatchEvent(t *testing.T) {
	type fields struct {
		handler *EventHandler
	}
	type args struct {
		deliveryID string
		eventName  string
		event      *github.WorkflowDispatchEvent
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "must trigger WorkflowDispatchEventAny with unknown event action",
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
					onWorkflowDispatchEvent: map[string][]WorkflowDispatchEventHandleFunc{
						WorkflowDispatchEventAnyAction: {
							func(deliveryID string, eventName string, event *github.WorkflowDispatchEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  WorkflowDispatchEvent,

				event: &github.WorkflowDispatchEvent{},
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
			if err := g.WorkflowDispatchEvent(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("WorkflowDispatchEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
