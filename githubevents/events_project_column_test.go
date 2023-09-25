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

func TestOnProjectColumnEventAny(t *testing.T) {
	type args struct {
		callbacks []ProjectColumnEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single ProjectColumnEventHandleFunc",
			args: args{
				[]ProjectColumnEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple ProjectColumnEventHandleFuncs",
			args: args{
				[]ProjectColumnEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectColumnEventAny(tt.args.callbacks...)
			if len(g.onProjectColumnEvent[ProjectColumnEventAnyAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onProjectColumnEvent[ProjectColumnEventAnyAction]))
			}
		})
	}
}

func TestSetOnProjectColumnEventAny(t *testing.T) {
	type args struct {
		callbacks []ProjectColumnEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single ProjectColumnEventHandleFunc",
			args: args{
				[]ProjectColumnEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple ProjectColumnEventHandleFuncs",
			args: args{
				[]ProjectColumnEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
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
			g.SetOnProjectColumnEventAny(func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
				return nil
			})
			g.SetOnProjectColumnEventAny(tt.args.callbacks...)
			if len(g.onProjectColumnEvent[ProjectColumnEventAnyAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onProjectColumnEvent[ProjectColumnEventAnyAction]), tt.want)
			}
		})
	}
}

func TestHandleProjectColumnEventAny(t *testing.T) {

	action := "*"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.ProjectColumnEvent
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
				eventName:  "project_column",

				event: &github.ProjectColumnEvent{Action: &action},

				fail: false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "project_column",

				event: &github.ProjectColumnEvent{Action: &action},

				fail: true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "project_column",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectColumnEventAny(func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleProjectColumnEventAny(tt.args.deliveryID, tt.args.deliveryID, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("TestHandleProjectColumnEventAny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnProjectColumnEventCreated(t *testing.T) {
	type args struct {
		callbacks []ProjectColumnEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single ProjectColumnEventHandleFunc",
			args: args{
				callbacks: []ProjectColumnEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple ProjectColumnEventHandleFunc",
			args: args{
				callbacks: []ProjectColumnEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectColumnEventCreated(tt.args.callbacks...)
			if len(g.onProjectColumnEvent[ProjectColumnEventCreatedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onProjectColumnEvent[ProjectColumnEventCreatedAction]))
			}
		})
	}
}

func TestSetOnProjectColumnEventCreated(t *testing.T) {
	type args struct {
		callbacks []ProjectColumnEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single ProjectColumnEventHandleFunc",
			args: args{
				[]ProjectColumnEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple ProjectColumnEventHandleFuncs",
			args: args{
				[]ProjectColumnEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
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
			g.SetOnProjectColumnEventCreated(func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
				return nil
			})
			g.SetOnProjectColumnEventCreated(tt.args.callbacks...)
			if len(g.onProjectColumnEvent[ProjectColumnEventCreatedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onProjectColumnEvent[ProjectColumnEventCreatedAction]), tt.want)
			}
		})
	}
}

func TestHandleProjectColumnEventCreated(t *testing.T) {
	action := ProjectColumnEventCreatedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.ProjectColumnEvent
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
				eventName:  "project_column",
				event:      &github.ProjectColumnEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "project_column",
				event:      &github.ProjectColumnEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "project_column",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "project_column",
				event:      &github.ProjectColumnEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "project_column",
				event:      &github.ProjectColumnEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "project_column",
				event:      &github.ProjectColumnEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectColumnEventCreated(func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleProjectColumnEventCreated(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleProjectColumnEventCreated() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnProjectColumnEventEdited(t *testing.T) {
	type args struct {
		callbacks []ProjectColumnEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single ProjectColumnEventHandleFunc",
			args: args{
				callbacks: []ProjectColumnEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple ProjectColumnEventHandleFunc",
			args: args{
				callbacks: []ProjectColumnEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectColumnEventEdited(tt.args.callbacks...)
			if len(g.onProjectColumnEvent[ProjectColumnEventEditedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onProjectColumnEvent[ProjectColumnEventEditedAction]))
			}
		})
	}
}

func TestSetOnProjectColumnEventEdited(t *testing.T) {
	type args struct {
		callbacks []ProjectColumnEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single ProjectColumnEventHandleFunc",
			args: args{
				[]ProjectColumnEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple ProjectColumnEventHandleFuncs",
			args: args{
				[]ProjectColumnEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
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
			g.SetOnProjectColumnEventEdited(func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
				return nil
			})
			g.SetOnProjectColumnEventEdited(tt.args.callbacks...)
			if len(g.onProjectColumnEvent[ProjectColumnEventEditedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onProjectColumnEvent[ProjectColumnEventEditedAction]), tt.want)
			}
		})
	}
}

func TestHandleProjectColumnEventEdited(t *testing.T) {
	action := ProjectColumnEventEditedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.ProjectColumnEvent
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
				eventName:  "project_column",
				event:      &github.ProjectColumnEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "project_column",
				event:      &github.ProjectColumnEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "project_column",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "project_column",
				event:      &github.ProjectColumnEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "project_column",
				event:      &github.ProjectColumnEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "project_column",
				event:      &github.ProjectColumnEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectColumnEventEdited(func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleProjectColumnEventEdited(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleProjectColumnEventEdited() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnProjectColumnEventMoved(t *testing.T) {
	type args struct {
		callbacks []ProjectColumnEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single ProjectColumnEventHandleFunc",
			args: args{
				callbacks: []ProjectColumnEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple ProjectColumnEventHandleFunc",
			args: args{
				callbacks: []ProjectColumnEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectColumnEventMoved(tt.args.callbacks...)
			if len(g.onProjectColumnEvent[ProjectColumnEventMovedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onProjectColumnEvent[ProjectColumnEventMovedAction]))
			}
		})
	}
}

func TestSetOnProjectColumnEventMoved(t *testing.T) {
	type args struct {
		callbacks []ProjectColumnEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single ProjectColumnEventHandleFunc",
			args: args{
				[]ProjectColumnEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple ProjectColumnEventHandleFuncs",
			args: args{
				[]ProjectColumnEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
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
			g.SetOnProjectColumnEventMoved(func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
				return nil
			})
			g.SetOnProjectColumnEventMoved(tt.args.callbacks...)
			if len(g.onProjectColumnEvent[ProjectColumnEventMovedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onProjectColumnEvent[ProjectColumnEventMovedAction]), tt.want)
			}
		})
	}
}

func TestHandleProjectColumnEventMoved(t *testing.T) {
	action := ProjectColumnEventMovedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.ProjectColumnEvent
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
				eventName:  "project_column",
				event:      &github.ProjectColumnEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "project_column",
				event:      &github.ProjectColumnEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "project_column",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "project_column",
				event:      &github.ProjectColumnEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "project_column",
				event:      &github.ProjectColumnEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "project_column",
				event:      &github.ProjectColumnEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectColumnEventMoved(func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleProjectColumnEventMoved(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleProjectColumnEventMoved() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnProjectColumnEventDeleted(t *testing.T) {
	type args struct {
		callbacks []ProjectColumnEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single ProjectColumnEventHandleFunc",
			args: args{
				callbacks: []ProjectColumnEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple ProjectColumnEventHandleFunc",
			args: args{
				callbacks: []ProjectColumnEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectColumnEventDeleted(tt.args.callbacks...)
			if len(g.onProjectColumnEvent[ProjectColumnEventDeletedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onProjectColumnEvent[ProjectColumnEventDeletedAction]))
			}
		})
	}
}

func TestSetOnProjectColumnEventDeleted(t *testing.T) {
	type args struct {
		callbacks []ProjectColumnEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single ProjectColumnEventHandleFunc",
			args: args{
				[]ProjectColumnEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple ProjectColumnEventHandleFuncs",
			args: args{
				[]ProjectColumnEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
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
			g.SetOnProjectColumnEventDeleted(func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
				return nil
			})
			g.SetOnProjectColumnEventDeleted(tt.args.callbacks...)
			if len(g.onProjectColumnEvent[ProjectColumnEventDeletedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onProjectColumnEvent[ProjectColumnEventDeletedAction]), tt.want)
			}
		})
	}
}

func TestHandleProjectColumnEventDeleted(t *testing.T) {
	action := ProjectColumnEventDeletedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.ProjectColumnEvent
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
				eventName:  "project_column",
				event:      &github.ProjectColumnEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "project_column",
				event:      &github.ProjectColumnEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "project_column",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "project_column",
				event:      &github.ProjectColumnEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "project_column",
				event:      &github.ProjectColumnEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "project_column",
				event:      &github.ProjectColumnEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectColumnEventDeleted(func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleProjectColumnEventDeleted(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleProjectColumnEventDeleted() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProjectColumnEvent(t *testing.T) {
	type fields struct {
		handler *EventHandler
	}
	type args struct {
		deliveryID string
		eventName  string
		event      *github.ProjectColumnEvent
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "must trigger ProjectColumnEventAny with unknown event action",
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
					onProjectColumnEvent: map[string][]ProjectColumnEventHandleFunc{
						ProjectColumnEventAnyAction: {
							func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  ProjectColumnEvent,

				event: &github.ProjectColumnEvent{Action: ptrString("unknown")},
			},
			wantErr: false,
		},

		{
			name: "must trigger ProjectColumnEventCreated",
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
					onProjectColumnEvent: map[string][]ProjectColumnEventHandleFunc{
						ProjectColumnEventAnyAction: {
							func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						ProjectColumnEventCreatedAction: {
							func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
								t.Logf("%s action called", ProjectColumnEventCreatedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "project_column",
				event:      &github.ProjectColumnEvent{Action: ptrString(ProjectColumnEventCreatedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail ProjectColumnEventCreated with empty action",
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
					onProjectColumnEvent: map[string][]ProjectColumnEventHandleFunc{
						ProjectColumnEventAnyAction: {
							func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						ProjectColumnEventCreatedAction: {
							func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
								t.Logf("%s action called", ProjectColumnEventCreatedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "project_column",
				event:      &github.ProjectColumnEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail ProjectColumnEventCreated with nil action",
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
					onProjectColumnEvent: map[string][]ProjectColumnEventHandleFunc{
						ProjectColumnEventAnyAction: {
							func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						ProjectColumnEventCreatedAction: {
							func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
								t.Logf("%s action called", ProjectColumnEventCreatedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "project_column",
				event:      &github.ProjectColumnEvent{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger ProjectColumnEventEdited",
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
					onProjectColumnEvent: map[string][]ProjectColumnEventHandleFunc{
						ProjectColumnEventAnyAction: {
							func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						ProjectColumnEventEditedAction: {
							func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
								t.Logf("%s action called", ProjectColumnEventEditedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "project_column",
				event:      &github.ProjectColumnEvent{Action: ptrString(ProjectColumnEventEditedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail ProjectColumnEventEdited with empty action",
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
					onProjectColumnEvent: map[string][]ProjectColumnEventHandleFunc{
						ProjectColumnEventAnyAction: {
							func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						ProjectColumnEventEditedAction: {
							func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
								t.Logf("%s action called", ProjectColumnEventEditedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "project_column",
				event:      &github.ProjectColumnEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail ProjectColumnEventEdited with nil action",
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
					onProjectColumnEvent: map[string][]ProjectColumnEventHandleFunc{
						ProjectColumnEventAnyAction: {
							func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						ProjectColumnEventEditedAction: {
							func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
								t.Logf("%s action called", ProjectColumnEventEditedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "project_column",
				event:      &github.ProjectColumnEvent{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger ProjectColumnEventMoved",
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
					onProjectColumnEvent: map[string][]ProjectColumnEventHandleFunc{
						ProjectColumnEventAnyAction: {
							func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						ProjectColumnEventMovedAction: {
							func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
								t.Logf("%s action called", ProjectColumnEventMovedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "project_column",
				event:      &github.ProjectColumnEvent{Action: ptrString(ProjectColumnEventMovedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail ProjectColumnEventMoved with empty action",
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
					onProjectColumnEvent: map[string][]ProjectColumnEventHandleFunc{
						ProjectColumnEventAnyAction: {
							func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						ProjectColumnEventMovedAction: {
							func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
								t.Logf("%s action called", ProjectColumnEventMovedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "project_column",
				event:      &github.ProjectColumnEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail ProjectColumnEventMoved with nil action",
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
					onProjectColumnEvent: map[string][]ProjectColumnEventHandleFunc{
						ProjectColumnEventAnyAction: {
							func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						ProjectColumnEventMovedAction: {
							func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
								t.Logf("%s action called", ProjectColumnEventMovedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "project_column",
				event:      &github.ProjectColumnEvent{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger ProjectColumnEventDeleted",
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
					onProjectColumnEvent: map[string][]ProjectColumnEventHandleFunc{
						ProjectColumnEventAnyAction: {
							func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						ProjectColumnEventDeletedAction: {
							func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
								t.Logf("%s action called", ProjectColumnEventDeletedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "project_column",
				event:      &github.ProjectColumnEvent{Action: ptrString(ProjectColumnEventDeletedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail ProjectColumnEventDeleted with empty action",
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
					onProjectColumnEvent: map[string][]ProjectColumnEventHandleFunc{
						ProjectColumnEventAnyAction: {
							func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						ProjectColumnEventDeletedAction: {
							func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
								t.Logf("%s action called", ProjectColumnEventDeletedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "project_column",
				event:      &github.ProjectColumnEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail ProjectColumnEventDeleted with nil action",
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
					onProjectColumnEvent: map[string][]ProjectColumnEventHandleFunc{
						ProjectColumnEventAnyAction: {
							func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						ProjectColumnEventDeletedAction: {
							func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
								t.Logf("%s action called", ProjectColumnEventDeletedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "project_column",
				event:      &github.ProjectColumnEvent{Action: nil},
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
			if err := g.ProjectColumnEvent(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("ProjectColumnEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
