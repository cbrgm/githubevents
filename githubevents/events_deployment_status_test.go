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

func TestOnDeploymentStatusEventAny(t *testing.T) {
	type args struct {
		callbacks []DeploymentStatusEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single DeploymentStatusEventHandleFunc",
			args: args{
				[]DeploymentStatusEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DeploymentStatusEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple DeploymentStatusEventHandleFuncs",
			args: args{
				[]DeploymentStatusEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DeploymentStatusEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.DeploymentStatusEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnDeploymentStatusEventAny(tt.args.callbacks...)
			if len(g.onDeploymentStatusEvent[DeploymentStatusEventAnyAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onDeploymentStatusEvent[DeploymentStatusEventAnyAction]))
			}
		})
	}
}

func TestSetOnDeploymentStatusEventAny(t *testing.T) {
	type args struct {
		callbacks []DeploymentStatusEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single DeploymentStatusEventHandleFunc",
			args: args{
				[]DeploymentStatusEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DeploymentStatusEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple DeploymentStatusEventHandleFuncs",
			args: args{
				[]DeploymentStatusEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DeploymentStatusEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.DeploymentStatusEvent) error {
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
			g.SetOnDeploymentStatusEventAny(func(deliveryID string, eventName string, event *github.DeploymentStatusEvent) error {
				return nil
			})
			g.SetOnDeploymentStatusEventAny(tt.args.callbacks...)
			if len(g.onDeploymentStatusEvent[DeploymentStatusEventAnyAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onDeploymentStatusEvent[DeploymentStatusEventAnyAction]), tt.want)
			}
		})
	}
}

func TestHandleDeploymentStatusEventAny(t *testing.T) {

	type args struct {
		deliveryID string
		eventName  string
		event      *github.DeploymentStatusEvent
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
				eventName:  "deployment_status",

				event: &github.DeploymentStatusEvent{},

				fail: false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "deployment_status",

				event: &github.DeploymentStatusEvent{},

				fail: true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "deployment_status",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnDeploymentStatusEventAny(func(deliveryID string, eventName string, event *github.DeploymentStatusEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleDeploymentStatusEventAny(tt.args.deliveryID, tt.args.deliveryID, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("TestHandleDeploymentStatusEventAny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeploymentStatusEvent(t *testing.T) {
	type fields struct {
		handler *EventHandler
	}
	type args struct {
		deliveryID string
		eventName  string
		event      *github.DeploymentStatusEvent
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "must trigger DeploymentStatusEventAny with unknown event action",
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
					onDeploymentStatusEvent: map[string][]DeploymentStatusEventHandleFunc{
						DeploymentStatusEventAnyAction: {
							func(deliveryID string, eventName string, event *github.DeploymentStatusEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  DeploymentStatusEvent,

				event: &github.DeploymentStatusEvent{},
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
			if err := g.DeploymentStatusEvent(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("DeploymentStatusEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
