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

func TestOnDeploymentEventAny(t *testing.T) {
	type args struct {
		callbacks []DeploymentEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single DeploymentEventHandleFunc",
			args: args{
				[]DeploymentEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DeploymentEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple DeploymentEventHandleFuncs",
			args: args{
				[]DeploymentEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DeploymentEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.DeploymentEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnDeploymentEventAny(tt.args.callbacks...)
			if len(g.onDeploymentEvent[DeploymentEventAnyAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onDeploymentEvent[DeploymentEventAnyAction]))
			}
		})
	}
}

func TestSetOnDeploymentEventAny(t *testing.T) {
	type args struct {
		callbacks []DeploymentEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single DeploymentEventHandleFunc",
			args: args{
				[]DeploymentEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DeploymentEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple DeploymentEventHandleFuncs",
			args: args{
				[]DeploymentEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DeploymentEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.DeploymentEvent) error {
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
			g.SetOnDeploymentEventAny(func(deliveryID string, eventName string, event *github.DeploymentEvent) error {
				return nil
			})
			g.SetOnDeploymentEventAny(tt.args.callbacks...)
			if len(g.onDeploymentEvent[DeploymentEventAnyAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onDeploymentEvent[DeploymentEventAnyAction]), tt.want)
			}
		})
	}
}

func TestHandleDeploymentEventAny(t *testing.T) {

	type args struct {
		deliveryID string
		eventName  string
		event      *github.DeploymentEvent
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
				eventName:  "deployment",

				event: &github.DeploymentEvent{},

				fail: false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "deployment",

				event: &github.DeploymentEvent{},

				fail: true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "deployment",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnDeploymentEventAny(func(deliveryID string, eventName string, event *github.DeploymentEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleDeploymentEventAny(tt.args.deliveryID, tt.args.deliveryID, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("TestHandleDeploymentEventAny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeploymentEvent(t *testing.T) {
	type fields struct {
		handler *EventHandler
	}
	type args struct {
		deliveryID string
		eventName  string
		event      *github.DeploymentEvent
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "must trigger DeploymentEventAny with unknown event action",
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
					onDeploymentEvent: map[string][]DeploymentEventHandleFunc{
						DeploymentEventAnyAction: {
							func(deliveryID string, eventName string, event *github.DeploymentEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  DeploymentEvent,

				event: &github.DeploymentEvent{},
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
			if err := g.DeploymentEvent(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("DeploymentEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
