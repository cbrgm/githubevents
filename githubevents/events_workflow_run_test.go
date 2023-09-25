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

func TestOnWorkflowRunEventAny(t *testing.T) {
	type args struct {
		callbacks []WorkflowRunEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single WorkflowRunEventHandleFunc",
			args: args{
				[]WorkflowRunEventHandleFunc{
					func(deliveryID string, eventName string, event *github.WorkflowRunEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple WorkflowRunEventHandleFuncs",
			args: args{
				[]WorkflowRunEventHandleFunc{
					func(deliveryID string, eventName string, event *github.WorkflowRunEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.WorkflowRunEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnWorkflowRunEventAny(tt.args.callbacks...)
			if len(g.onWorkflowRunEvent[WorkflowRunEventAnyAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onWorkflowRunEvent[WorkflowRunEventAnyAction]))
			}
		})
	}
}

func TestSetOnWorkflowRunEventAny(t *testing.T) {
	type args struct {
		callbacks []WorkflowRunEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single WorkflowRunEventHandleFunc",
			args: args{
				[]WorkflowRunEventHandleFunc{
					func(deliveryID string, eventName string, event *github.WorkflowRunEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple WorkflowRunEventHandleFuncs",
			args: args{
				[]WorkflowRunEventHandleFunc{
					func(deliveryID string, eventName string, event *github.WorkflowRunEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.WorkflowRunEvent) error {
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
			g.SetOnWorkflowRunEventAny(func(deliveryID string, eventName string, event *github.WorkflowRunEvent) error {
				return nil
			})
			g.SetOnWorkflowRunEventAny(tt.args.callbacks...)
			if len(g.onWorkflowRunEvent[WorkflowRunEventAnyAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onWorkflowRunEvent[WorkflowRunEventAnyAction]), tt.want)
			}
		})
	}
}

func TestHandleWorkflowRunEventAny(t *testing.T) {

	action := "*"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.WorkflowRunEvent
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
				eventName:  "workflow_run",

				event: &github.WorkflowRunEvent{Action: &action},

				fail: false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "workflow_run",

				event: &github.WorkflowRunEvent{Action: &action},

				fail: true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "workflow_run",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnWorkflowRunEventAny(func(deliveryID string, eventName string, event *github.WorkflowRunEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleWorkflowRunEventAny(tt.args.deliveryID, tt.args.deliveryID, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("TestHandleWorkflowRunEventAny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnWorkflowRunEventRequested(t *testing.T) {
	type args struct {
		callbacks []WorkflowRunEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single WorkflowRunEventHandleFunc",
			args: args{
				callbacks: []WorkflowRunEventHandleFunc{
					func(deliveryID string, eventName string, event *github.WorkflowRunEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple WorkflowRunEventHandleFunc",
			args: args{
				callbacks: []WorkflowRunEventHandleFunc{
					func(deliveryID string, eventName string, event *github.WorkflowRunEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.WorkflowRunEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnWorkflowRunEventRequested(tt.args.callbacks...)
			if len(g.onWorkflowRunEvent[WorkflowRunEventRequestedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onWorkflowRunEvent[WorkflowRunEventRequestedAction]))
			}
		})
	}
}

func TestSetOnWorkflowRunEventRequested(t *testing.T) {
	type args struct {
		callbacks []WorkflowRunEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single WorkflowRunEventHandleFunc",
			args: args{
				[]WorkflowRunEventHandleFunc{
					func(deliveryID string, eventName string, event *github.WorkflowRunEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple WorkflowRunEventHandleFuncs",
			args: args{
				[]WorkflowRunEventHandleFunc{
					func(deliveryID string, eventName string, event *github.WorkflowRunEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.WorkflowRunEvent) error {
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
			g.SetOnWorkflowRunEventRequested(func(deliveryID string, eventName string, event *github.WorkflowRunEvent) error {
				return nil
			})
			g.SetOnWorkflowRunEventRequested(tt.args.callbacks...)
			if len(g.onWorkflowRunEvent[WorkflowRunEventRequestedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onWorkflowRunEvent[WorkflowRunEventRequestedAction]), tt.want)
			}
		})
	}
}

func TestHandleWorkflowRunEventRequested(t *testing.T) {
	action := WorkflowRunEventRequestedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.WorkflowRunEvent
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
				eventName:  "workflow_run",
				event:      &github.WorkflowRunEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "workflow_run",
				event:      &github.WorkflowRunEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "workflow_run",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "workflow_run",
				event:      &github.WorkflowRunEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "workflow_run",
				event:      &github.WorkflowRunEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "workflow_run",
				event:      &github.WorkflowRunEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnWorkflowRunEventRequested(func(deliveryID string, eventName string, event *github.WorkflowRunEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleWorkflowRunEventRequested(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleWorkflowRunEventRequested() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnWorkflowRunEventCompleted(t *testing.T) {
	type args struct {
		callbacks []WorkflowRunEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single WorkflowRunEventHandleFunc",
			args: args{
				callbacks: []WorkflowRunEventHandleFunc{
					func(deliveryID string, eventName string, event *github.WorkflowRunEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple WorkflowRunEventHandleFunc",
			args: args{
				callbacks: []WorkflowRunEventHandleFunc{
					func(deliveryID string, eventName string, event *github.WorkflowRunEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.WorkflowRunEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnWorkflowRunEventCompleted(tt.args.callbacks...)
			if len(g.onWorkflowRunEvent[WorkflowRunEventCompletedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onWorkflowRunEvent[WorkflowRunEventCompletedAction]))
			}
		})
	}
}

func TestSetOnWorkflowRunEventCompleted(t *testing.T) {
	type args struct {
		callbacks []WorkflowRunEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single WorkflowRunEventHandleFunc",
			args: args{
				[]WorkflowRunEventHandleFunc{
					func(deliveryID string, eventName string, event *github.WorkflowRunEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple WorkflowRunEventHandleFuncs",
			args: args{
				[]WorkflowRunEventHandleFunc{
					func(deliveryID string, eventName string, event *github.WorkflowRunEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.WorkflowRunEvent) error {
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
			g.SetOnWorkflowRunEventCompleted(func(deliveryID string, eventName string, event *github.WorkflowRunEvent) error {
				return nil
			})
			g.SetOnWorkflowRunEventCompleted(tt.args.callbacks...)
			if len(g.onWorkflowRunEvent[WorkflowRunEventCompletedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onWorkflowRunEvent[WorkflowRunEventCompletedAction]), tt.want)
			}
		})
	}
}

func TestHandleWorkflowRunEventCompleted(t *testing.T) {
	action := WorkflowRunEventCompletedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.WorkflowRunEvent
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
				eventName:  "workflow_run",
				event:      &github.WorkflowRunEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "workflow_run",
				event:      &github.WorkflowRunEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "workflow_run",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "workflow_run",
				event:      &github.WorkflowRunEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "workflow_run",
				event:      &github.WorkflowRunEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "workflow_run",
				event:      &github.WorkflowRunEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnWorkflowRunEventCompleted(func(deliveryID string, eventName string, event *github.WorkflowRunEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleWorkflowRunEventCompleted(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleWorkflowRunEventCompleted() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWorkflowRunEvent(t *testing.T) {
	type fields struct {
		handler *EventHandler
	}
	type args struct {
		deliveryID string
		eventName  string
		event      *github.WorkflowRunEvent
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "must trigger WorkflowRunEventAny with unknown event action",
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
					onWorkflowRunEvent: map[string][]WorkflowRunEventHandleFunc{
						WorkflowRunEventAnyAction: {
							func(deliveryID string, eventName string, event *github.WorkflowRunEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  WorkflowRunEvent,

				event: &github.WorkflowRunEvent{Action: ptrString("unknown")},
			},
			wantErr: false,
		},

		{
			name: "must trigger WorkflowRunEventRequested",
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
					onWorkflowRunEvent: map[string][]WorkflowRunEventHandleFunc{
						WorkflowRunEventAnyAction: {
							func(deliveryID string, eventName string, event *github.WorkflowRunEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						WorkflowRunEventRequestedAction: {
							func(deliveryID string, eventName string, event *github.WorkflowRunEvent) error {
								t.Logf("%s action called", WorkflowRunEventRequestedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "workflow_run",
				event:      &github.WorkflowRunEvent{Action: ptrString(WorkflowRunEventRequestedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail WorkflowRunEventRequested with empty action",
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
					onWorkflowRunEvent: map[string][]WorkflowRunEventHandleFunc{
						WorkflowRunEventAnyAction: {
							func(deliveryID string, eventName string, event *github.WorkflowRunEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						WorkflowRunEventRequestedAction: {
							func(deliveryID string, eventName string, event *github.WorkflowRunEvent) error {
								t.Logf("%s action called", WorkflowRunEventRequestedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "workflow_run",
				event:      &github.WorkflowRunEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail WorkflowRunEventRequested with nil action",
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
					onWorkflowRunEvent: map[string][]WorkflowRunEventHandleFunc{
						WorkflowRunEventAnyAction: {
							func(deliveryID string, eventName string, event *github.WorkflowRunEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						WorkflowRunEventRequestedAction: {
							func(deliveryID string, eventName string, event *github.WorkflowRunEvent) error {
								t.Logf("%s action called", WorkflowRunEventRequestedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "workflow_run",
				event:      &github.WorkflowRunEvent{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger WorkflowRunEventCompleted",
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
					onWorkflowRunEvent: map[string][]WorkflowRunEventHandleFunc{
						WorkflowRunEventAnyAction: {
							func(deliveryID string, eventName string, event *github.WorkflowRunEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						WorkflowRunEventCompletedAction: {
							func(deliveryID string, eventName string, event *github.WorkflowRunEvent) error {
								t.Logf("%s action called", WorkflowRunEventCompletedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "workflow_run",
				event:      &github.WorkflowRunEvent{Action: ptrString(WorkflowRunEventCompletedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail WorkflowRunEventCompleted with empty action",
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
					onWorkflowRunEvent: map[string][]WorkflowRunEventHandleFunc{
						WorkflowRunEventAnyAction: {
							func(deliveryID string, eventName string, event *github.WorkflowRunEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						WorkflowRunEventCompletedAction: {
							func(deliveryID string, eventName string, event *github.WorkflowRunEvent) error {
								t.Logf("%s action called", WorkflowRunEventCompletedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "workflow_run",
				event:      &github.WorkflowRunEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail WorkflowRunEventCompleted with nil action",
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
					onWorkflowRunEvent: map[string][]WorkflowRunEventHandleFunc{
						WorkflowRunEventAnyAction: {
							func(deliveryID string, eventName string, event *github.WorkflowRunEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						WorkflowRunEventCompletedAction: {
							func(deliveryID string, eventName string, event *github.WorkflowRunEvent) error {
								t.Logf("%s action called", WorkflowRunEventCompletedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "workflow_run",
				event:      &github.WorkflowRunEvent{Action: nil},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &EventHandler{
				WebhookSecret: "fake",
				mu:            sync.RWMutex{},
			}
			if err := g.WorkflowRunEvent(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("WorkflowRunEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
