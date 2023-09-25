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

func TestOnWorkflowJobEventAny(t *testing.T) {
	type args struct {
		callbacks []WorkflowJobEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single WorkflowJobEventHandleFunc",
			args: args{
				[]WorkflowJobEventHandleFunc{
					func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple WorkflowJobEventHandleFuncs",
			args: args{
				[]WorkflowJobEventHandleFunc{
					func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnWorkflowJobEventAny(tt.args.callbacks...)
			if len(g.onWorkflowJobEvent[WorkflowJobEventAnyAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onWorkflowJobEvent[WorkflowJobEventAnyAction]))
			}
		})
	}
}

func TestSetOnWorkflowJobEventAny(t *testing.T) {
	type args struct {
		callbacks []WorkflowJobEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single WorkflowJobEventHandleFunc",
			args: args{
				[]WorkflowJobEventHandleFunc{
					func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple WorkflowJobEventHandleFuncs",
			args: args{
				[]WorkflowJobEventHandleFunc{
					func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
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
			g.SetOnWorkflowJobEventAny(func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
				return nil
			})
			g.SetOnWorkflowJobEventAny(tt.args.callbacks...)
			if len(g.onWorkflowJobEvent[WorkflowJobEventAnyAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onWorkflowJobEvent[WorkflowJobEventAnyAction]), tt.want)
			}
		})
	}
}

func TestHandleWorkflowJobEventAny(t *testing.T) {

	action := "*"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.WorkflowJobEvent
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
				eventName:  "workflow_job",

				event: &github.WorkflowJobEvent{Action: &action},

				fail: false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "workflow_job",

				event: &github.WorkflowJobEvent{Action: &action},

				fail: true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "workflow_job",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnWorkflowJobEventAny(func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleWorkflowJobEventAny(tt.args.deliveryID, tt.args.deliveryID, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("TestHandleWorkflowJobEventAny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnWorkflowJobEventQueued(t *testing.T) {
	type args struct {
		callbacks []WorkflowJobEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single WorkflowJobEventHandleFunc",
			args: args{
				callbacks: []WorkflowJobEventHandleFunc{
					func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple WorkflowJobEventHandleFunc",
			args: args{
				callbacks: []WorkflowJobEventHandleFunc{
					func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnWorkflowJobEventQueued(tt.args.callbacks...)
			if len(g.onWorkflowJobEvent[WorkflowJobEventQueuedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onWorkflowJobEvent[WorkflowJobEventQueuedAction]))
			}
		})
	}
}

func TestSetOnWorkflowJobEventQueued(t *testing.T) {
	type args struct {
		callbacks []WorkflowJobEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single WorkflowJobEventHandleFunc",
			args: args{
				[]WorkflowJobEventHandleFunc{
					func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple WorkflowJobEventHandleFuncs",
			args: args{
				[]WorkflowJobEventHandleFunc{
					func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
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
			g.SetOnWorkflowJobEventQueued(func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
				return nil
			})
			g.SetOnWorkflowJobEventQueued(tt.args.callbacks...)
			if len(g.onWorkflowJobEvent[WorkflowJobEventQueuedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onWorkflowJobEvent[WorkflowJobEventQueuedAction]), tt.want)
			}
		})
	}
}

func TestHandleWorkflowJobEventQueued(t *testing.T) {
	action := WorkflowJobEventQueuedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.WorkflowJobEvent
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
				eventName:  "workflow_job",
				event:      &github.WorkflowJobEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "workflow_job",
				event:      &github.WorkflowJobEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "workflow_job",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "workflow_job",
				event:      &github.WorkflowJobEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "workflow_job",
				event:      &github.WorkflowJobEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "workflow_job",
				event:      &github.WorkflowJobEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnWorkflowJobEventQueued(func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleWorkflowJobEventQueued(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleWorkflowJobEventQueued() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnWorkflowJobEventInProgress(t *testing.T) {
	type args struct {
		callbacks []WorkflowJobEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single WorkflowJobEventHandleFunc",
			args: args{
				callbacks: []WorkflowJobEventHandleFunc{
					func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple WorkflowJobEventHandleFunc",
			args: args{
				callbacks: []WorkflowJobEventHandleFunc{
					func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnWorkflowJobEventInProgress(tt.args.callbacks...)
			if len(g.onWorkflowJobEvent[WorkflowJobEventInProgressAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onWorkflowJobEvent[WorkflowJobEventInProgressAction]))
			}
		})
	}
}

func TestSetOnWorkflowJobEventInProgress(t *testing.T) {
	type args struct {
		callbacks []WorkflowJobEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single WorkflowJobEventHandleFunc",
			args: args{
				[]WorkflowJobEventHandleFunc{
					func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple WorkflowJobEventHandleFuncs",
			args: args{
				[]WorkflowJobEventHandleFunc{
					func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
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
			g.SetOnWorkflowJobEventInProgress(func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
				return nil
			})
			g.SetOnWorkflowJobEventInProgress(tt.args.callbacks...)
			if len(g.onWorkflowJobEvent[WorkflowJobEventInProgressAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onWorkflowJobEvent[WorkflowJobEventInProgressAction]), tt.want)
			}
		})
	}
}

func TestHandleWorkflowJobEventInProgress(t *testing.T) {
	action := WorkflowJobEventInProgressAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.WorkflowJobEvent
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
				eventName:  "workflow_job",
				event:      &github.WorkflowJobEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "workflow_job",
				event:      &github.WorkflowJobEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "workflow_job",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "workflow_job",
				event:      &github.WorkflowJobEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "workflow_job",
				event:      &github.WorkflowJobEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "workflow_job",
				event:      &github.WorkflowJobEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnWorkflowJobEventInProgress(func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleWorkflowJobEventInProgress(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleWorkflowJobEventInProgress() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnWorkflowJobEventCompleted(t *testing.T) {
	type args struct {
		callbacks []WorkflowJobEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single WorkflowJobEventHandleFunc",
			args: args{
				callbacks: []WorkflowJobEventHandleFunc{
					func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple WorkflowJobEventHandleFunc",
			args: args{
				callbacks: []WorkflowJobEventHandleFunc{
					func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnWorkflowJobEventCompleted(tt.args.callbacks...)
			if len(g.onWorkflowJobEvent[WorkflowJobEventCompletedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onWorkflowJobEvent[WorkflowJobEventCompletedAction]))
			}
		})
	}
}

func TestSetOnWorkflowJobEventCompleted(t *testing.T) {
	type args struct {
		callbacks []WorkflowJobEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single WorkflowJobEventHandleFunc",
			args: args{
				[]WorkflowJobEventHandleFunc{
					func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple WorkflowJobEventHandleFuncs",
			args: args{
				[]WorkflowJobEventHandleFunc{
					func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
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
			g.SetOnWorkflowJobEventCompleted(func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
				return nil
			})
			g.SetOnWorkflowJobEventCompleted(tt.args.callbacks...)
			if len(g.onWorkflowJobEvent[WorkflowJobEventCompletedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onWorkflowJobEvent[WorkflowJobEventCompletedAction]), tt.want)
			}
		})
	}
}

func TestHandleWorkflowJobEventCompleted(t *testing.T) {
	action := WorkflowJobEventCompletedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.WorkflowJobEvent
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
				eventName:  "workflow_job",
				event:      &github.WorkflowJobEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "workflow_job",
				event:      &github.WorkflowJobEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "workflow_job",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "workflow_job",
				event:      &github.WorkflowJobEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "workflow_job",
				event:      &github.WorkflowJobEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "workflow_job",
				event:      &github.WorkflowJobEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnWorkflowJobEventCompleted(func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleWorkflowJobEventCompleted(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleWorkflowJobEventCompleted() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWorkflowJobEvent(t *testing.T) {
	type fields struct {
		handler *EventHandler
	}
	type args struct {
		deliveryID string
		eventName  string
		event      *github.WorkflowJobEvent
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "must trigger WorkflowJobEventAny with unknown event action",
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
					onWorkflowJobEvent: map[string][]WorkflowJobEventHandleFunc{
						WorkflowJobEventAnyAction: {
							func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  WorkflowJobEvent,

				event: &github.WorkflowJobEvent{Action: ptrString("unknown")},
			},
			wantErr: false,
		},

		{
			name: "must trigger WorkflowJobEventQueued",
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
					onWorkflowJobEvent: map[string][]WorkflowJobEventHandleFunc{
						WorkflowJobEventAnyAction: {
							func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						WorkflowJobEventQueuedAction: {
							func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
								t.Logf("%s action called", WorkflowJobEventQueuedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "workflow_job",
				event:      &github.WorkflowJobEvent{Action: ptrString(WorkflowJobEventQueuedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail WorkflowJobEventQueued with empty action",
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
					onWorkflowJobEvent: map[string][]WorkflowJobEventHandleFunc{
						WorkflowJobEventAnyAction: {
							func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						WorkflowJobEventQueuedAction: {
							func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
								t.Logf("%s action called", WorkflowJobEventQueuedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "workflow_job",
				event:      &github.WorkflowJobEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail WorkflowJobEventQueued with nil action",
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
					onWorkflowJobEvent: map[string][]WorkflowJobEventHandleFunc{
						WorkflowJobEventAnyAction: {
							func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						WorkflowJobEventQueuedAction: {
							func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
								t.Logf("%s action called", WorkflowJobEventQueuedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "workflow_job",
				event:      &github.WorkflowJobEvent{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger WorkflowJobEventInProgress",
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
					onWorkflowJobEvent: map[string][]WorkflowJobEventHandleFunc{
						WorkflowJobEventAnyAction: {
							func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						WorkflowJobEventInProgressAction: {
							func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
								t.Logf("%s action called", WorkflowJobEventInProgressAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "workflow_job",
				event:      &github.WorkflowJobEvent{Action: ptrString(WorkflowJobEventInProgressAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail WorkflowJobEventInProgress with empty action",
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
					onWorkflowJobEvent: map[string][]WorkflowJobEventHandleFunc{
						WorkflowJobEventAnyAction: {
							func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						WorkflowJobEventInProgressAction: {
							func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
								t.Logf("%s action called", WorkflowJobEventInProgressAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "workflow_job",
				event:      &github.WorkflowJobEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail WorkflowJobEventInProgress with nil action",
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
					onWorkflowJobEvent: map[string][]WorkflowJobEventHandleFunc{
						WorkflowJobEventAnyAction: {
							func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						WorkflowJobEventInProgressAction: {
							func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
								t.Logf("%s action called", WorkflowJobEventInProgressAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "workflow_job",
				event:      &github.WorkflowJobEvent{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger WorkflowJobEventCompleted",
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
					onWorkflowJobEvent: map[string][]WorkflowJobEventHandleFunc{
						WorkflowJobEventAnyAction: {
							func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						WorkflowJobEventCompletedAction: {
							func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
								t.Logf("%s action called", WorkflowJobEventCompletedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "workflow_job",
				event:      &github.WorkflowJobEvent{Action: ptrString(WorkflowJobEventCompletedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail WorkflowJobEventCompleted with empty action",
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
					onWorkflowJobEvent: map[string][]WorkflowJobEventHandleFunc{
						WorkflowJobEventAnyAction: {
							func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						WorkflowJobEventCompletedAction: {
							func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
								t.Logf("%s action called", WorkflowJobEventCompletedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "workflow_job",
				event:      &github.WorkflowJobEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail WorkflowJobEventCompleted with nil action",
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
					onWorkflowJobEvent: map[string][]WorkflowJobEventHandleFunc{
						WorkflowJobEventAnyAction: {
							func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						WorkflowJobEventCompletedAction: {
							func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
								t.Logf("%s action called", WorkflowJobEventCompletedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "workflow_job",
				event:      &github.WorkflowJobEvent{Action: nil},
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
			if err := g.WorkflowJobEvent(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("WorkflowJobEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
