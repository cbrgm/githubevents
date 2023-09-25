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

func TestOnProjectEventAny(t *testing.T) {
	type args struct {
		callbacks []ProjectEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single ProjectEventHandleFunc",
			args: args{
				[]ProjectEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple ProjectEventHandleFuncs",
			args: args{
				[]ProjectEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectEventAny(tt.args.callbacks...)
			if len(g.onProjectEvent[ProjectEventAnyAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onProjectEvent[ProjectEventAnyAction]))
			}
		})
	}
}

func TestSetOnProjectEventAny(t *testing.T) {
	type args struct {
		callbacks []ProjectEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single ProjectEventHandleFunc",
			args: args{
				[]ProjectEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple ProjectEventHandleFuncs",
			args: args{
				[]ProjectEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectEvent) error {
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
			g.SetOnProjectEventAny(func(deliveryID string, eventName string, event *github.ProjectEvent) error {
				return nil
			})
			g.SetOnProjectEventAny(tt.args.callbacks...)
			if len(g.onProjectEvent[ProjectEventAnyAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onProjectEvent[ProjectEventAnyAction]), tt.want)
			}
		})
	}
}

func TestHandleProjectEventAny(t *testing.T) {

	action := "*"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.ProjectEvent
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
				eventName:  "project",

				event: &github.ProjectEvent{Action: &action},

				fail: false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "project",

				event: &github.ProjectEvent{Action: &action},

				fail: true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "project",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectEventAny(func(deliveryID string, eventName string, event *github.ProjectEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleProjectEventAny(tt.args.deliveryID, tt.args.deliveryID, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("TestHandleProjectEventAny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnProjectEventCreated(t *testing.T) {
	type args struct {
		callbacks []ProjectEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single ProjectEventHandleFunc",
			args: args{
				callbacks: []ProjectEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple ProjectEventHandleFunc",
			args: args{
				callbacks: []ProjectEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectEventCreated(tt.args.callbacks...)
			if len(g.onProjectEvent[ProjectEventCreatedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onProjectEvent[ProjectEventCreatedAction]))
			}
		})
	}
}

func TestSetOnProjectEventCreated(t *testing.T) {
	type args struct {
		callbacks []ProjectEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single ProjectEventHandleFunc",
			args: args{
				[]ProjectEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple ProjectEventHandleFuncs",
			args: args{
				[]ProjectEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectEvent) error {
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
			g.SetOnProjectEventCreated(func(deliveryID string, eventName string, event *github.ProjectEvent) error {
				return nil
			})
			g.SetOnProjectEventCreated(tt.args.callbacks...)
			if len(g.onProjectEvent[ProjectEventCreatedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onProjectEvent[ProjectEventCreatedAction]), tt.want)
			}
		})
	}
}

func TestHandleProjectEventCreated(t *testing.T) {
	action := ProjectEventCreatedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.ProjectEvent
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
				eventName:  "project",
				event:      &github.ProjectEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "project",
				event:      &github.ProjectEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "project",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "project",
				event:      &github.ProjectEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "project",
				event:      &github.ProjectEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "project",
				event:      &github.ProjectEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectEventCreated(func(deliveryID string, eventName string, event *github.ProjectEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleProjectEventCreated(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleProjectEventCreated() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnProjectEventEdited(t *testing.T) {
	type args struct {
		callbacks []ProjectEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single ProjectEventHandleFunc",
			args: args{
				callbacks: []ProjectEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple ProjectEventHandleFunc",
			args: args{
				callbacks: []ProjectEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectEventEdited(tt.args.callbacks...)
			if len(g.onProjectEvent[ProjectEventEditedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onProjectEvent[ProjectEventEditedAction]))
			}
		})
	}
}

func TestSetOnProjectEventEdited(t *testing.T) {
	type args struct {
		callbacks []ProjectEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single ProjectEventHandleFunc",
			args: args{
				[]ProjectEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple ProjectEventHandleFuncs",
			args: args{
				[]ProjectEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectEvent) error {
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
			g.SetOnProjectEventEdited(func(deliveryID string, eventName string, event *github.ProjectEvent) error {
				return nil
			})
			g.SetOnProjectEventEdited(tt.args.callbacks...)
			if len(g.onProjectEvent[ProjectEventEditedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onProjectEvent[ProjectEventEditedAction]), tt.want)
			}
		})
	}
}

func TestHandleProjectEventEdited(t *testing.T) {
	action := ProjectEventEditedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.ProjectEvent
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
				eventName:  "project",
				event:      &github.ProjectEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "project",
				event:      &github.ProjectEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "project",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "project",
				event:      &github.ProjectEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "project",
				event:      &github.ProjectEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "project",
				event:      &github.ProjectEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectEventEdited(func(deliveryID string, eventName string, event *github.ProjectEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleProjectEventEdited(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleProjectEventEdited() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnProjectEventClosed(t *testing.T) {
	type args struct {
		callbacks []ProjectEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single ProjectEventHandleFunc",
			args: args{
				callbacks: []ProjectEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple ProjectEventHandleFunc",
			args: args{
				callbacks: []ProjectEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectEventClosed(tt.args.callbacks...)
			if len(g.onProjectEvent[ProjectEventClosedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onProjectEvent[ProjectEventClosedAction]))
			}
		})
	}
}

func TestSetOnProjectEventClosed(t *testing.T) {
	type args struct {
		callbacks []ProjectEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single ProjectEventHandleFunc",
			args: args{
				[]ProjectEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple ProjectEventHandleFuncs",
			args: args{
				[]ProjectEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectEvent) error {
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
			g.SetOnProjectEventClosed(func(deliveryID string, eventName string, event *github.ProjectEvent) error {
				return nil
			})
			g.SetOnProjectEventClosed(tt.args.callbacks...)
			if len(g.onProjectEvent[ProjectEventClosedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onProjectEvent[ProjectEventClosedAction]), tt.want)
			}
		})
	}
}

func TestHandleProjectEventClosed(t *testing.T) {
	action := ProjectEventClosedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.ProjectEvent
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
				eventName:  "project",
				event:      &github.ProjectEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "project",
				event:      &github.ProjectEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "project",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "project",
				event:      &github.ProjectEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "project",
				event:      &github.ProjectEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "project",
				event:      &github.ProjectEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectEventClosed(func(deliveryID string, eventName string, event *github.ProjectEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleProjectEventClosed(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleProjectEventClosed() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnProjectEventReopened(t *testing.T) {
	type args struct {
		callbacks []ProjectEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single ProjectEventHandleFunc",
			args: args{
				callbacks: []ProjectEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple ProjectEventHandleFunc",
			args: args{
				callbacks: []ProjectEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectEventReopened(tt.args.callbacks...)
			if len(g.onProjectEvent[ProjectEventReopenedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onProjectEvent[ProjectEventReopenedAction]))
			}
		})
	}
}

func TestSetOnProjectEventReopened(t *testing.T) {
	type args struct {
		callbacks []ProjectEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single ProjectEventHandleFunc",
			args: args{
				[]ProjectEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple ProjectEventHandleFuncs",
			args: args{
				[]ProjectEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectEvent) error {
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
			g.SetOnProjectEventReopened(func(deliveryID string, eventName string, event *github.ProjectEvent) error {
				return nil
			})
			g.SetOnProjectEventReopened(tt.args.callbacks...)
			if len(g.onProjectEvent[ProjectEventReopenedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onProjectEvent[ProjectEventReopenedAction]), tt.want)
			}
		})
	}
}

func TestHandleProjectEventReopened(t *testing.T) {
	action := ProjectEventReopenedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.ProjectEvent
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
				eventName:  "project",
				event:      &github.ProjectEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "project",
				event:      &github.ProjectEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "project",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "project",
				event:      &github.ProjectEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "project",
				event:      &github.ProjectEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "project",
				event:      &github.ProjectEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectEventReopened(func(deliveryID string, eventName string, event *github.ProjectEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleProjectEventReopened(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleProjectEventReopened() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnProjectEventDeleted(t *testing.T) {
	type args struct {
		callbacks []ProjectEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single ProjectEventHandleFunc",
			args: args{
				callbacks: []ProjectEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple ProjectEventHandleFunc",
			args: args{
				callbacks: []ProjectEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectEventDeleted(tt.args.callbacks...)
			if len(g.onProjectEvent[ProjectEventDeletedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onProjectEvent[ProjectEventDeletedAction]))
			}
		})
	}
}

func TestSetOnProjectEventDeleted(t *testing.T) {
	type args struct {
		callbacks []ProjectEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single ProjectEventHandleFunc",
			args: args{
				[]ProjectEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple ProjectEventHandleFuncs",
			args: args{
				[]ProjectEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectEvent) error {
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
			g.SetOnProjectEventDeleted(func(deliveryID string, eventName string, event *github.ProjectEvent) error {
				return nil
			})
			g.SetOnProjectEventDeleted(tt.args.callbacks...)
			if len(g.onProjectEvent[ProjectEventDeletedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onProjectEvent[ProjectEventDeletedAction]), tt.want)
			}
		})
	}
}

func TestHandleProjectEventDeleted(t *testing.T) {
	action := ProjectEventDeletedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.ProjectEvent
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
				eventName:  "project",
				event:      &github.ProjectEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "project",
				event:      &github.ProjectEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "project",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "project",
				event:      &github.ProjectEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "project",
				event:      &github.ProjectEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "project",
				event:      &github.ProjectEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectEventDeleted(func(deliveryID string, eventName string, event *github.ProjectEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleProjectEventDeleted(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleProjectEventDeleted() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProjectEvent(t *testing.T) {
	type fields struct {
		handler *EventHandler
	}
	type args struct {
		deliveryID string
		eventName  string
		event      *github.ProjectEvent
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "must trigger ProjectEventAny with unknown event action",
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
					onProjectEvent: map[string][]ProjectEventHandleFunc{
						ProjectEventAnyAction: {
							func(deliveryID string, eventName string, event *github.ProjectEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  ProjectEvent,

				event: &github.ProjectEvent{Action: ptrString("unknown")},
			},
			wantErr: false,
		},

		{
			name: "must trigger ProjectEventCreated",
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
					onProjectEvent: map[string][]ProjectEventHandleFunc{
						ProjectEventAnyAction: {
							func(deliveryID string, eventName string, event *github.ProjectEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						ProjectEventCreatedAction: {
							func(deliveryID string, eventName string, event *github.ProjectEvent) error {
								t.Logf("%s action called", ProjectEventCreatedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "project",
				event:      &github.ProjectEvent{Action: ptrString(ProjectEventCreatedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail ProjectEventCreated with empty action",
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
					onProjectEvent: map[string][]ProjectEventHandleFunc{
						ProjectEventAnyAction: {
							func(deliveryID string, eventName string, event *github.ProjectEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						ProjectEventCreatedAction: {
							func(deliveryID string, eventName string, event *github.ProjectEvent) error {
								t.Logf("%s action called", ProjectEventCreatedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "project",
				event:      &github.ProjectEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail ProjectEventCreated with nil action",
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
					onProjectEvent: map[string][]ProjectEventHandleFunc{
						ProjectEventAnyAction: {
							func(deliveryID string, eventName string, event *github.ProjectEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						ProjectEventCreatedAction: {
							func(deliveryID string, eventName string, event *github.ProjectEvent) error {
								t.Logf("%s action called", ProjectEventCreatedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "project",
				event:      &github.ProjectEvent{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger ProjectEventEdited",
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
					onProjectEvent: map[string][]ProjectEventHandleFunc{
						ProjectEventAnyAction: {
							func(deliveryID string, eventName string, event *github.ProjectEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						ProjectEventEditedAction: {
							func(deliveryID string, eventName string, event *github.ProjectEvent) error {
								t.Logf("%s action called", ProjectEventEditedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "project",
				event:      &github.ProjectEvent{Action: ptrString(ProjectEventEditedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail ProjectEventEdited with empty action",
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
					onProjectEvent: map[string][]ProjectEventHandleFunc{
						ProjectEventAnyAction: {
							func(deliveryID string, eventName string, event *github.ProjectEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						ProjectEventEditedAction: {
							func(deliveryID string, eventName string, event *github.ProjectEvent) error {
								t.Logf("%s action called", ProjectEventEditedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "project",
				event:      &github.ProjectEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail ProjectEventEdited with nil action",
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
					onProjectEvent: map[string][]ProjectEventHandleFunc{
						ProjectEventAnyAction: {
							func(deliveryID string, eventName string, event *github.ProjectEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						ProjectEventEditedAction: {
							func(deliveryID string, eventName string, event *github.ProjectEvent) error {
								t.Logf("%s action called", ProjectEventEditedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "project",
				event:      &github.ProjectEvent{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger ProjectEventClosed",
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
					onProjectEvent: map[string][]ProjectEventHandleFunc{
						ProjectEventAnyAction: {
							func(deliveryID string, eventName string, event *github.ProjectEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						ProjectEventClosedAction: {
							func(deliveryID string, eventName string, event *github.ProjectEvent) error {
								t.Logf("%s action called", ProjectEventClosedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "project",
				event:      &github.ProjectEvent{Action: ptrString(ProjectEventClosedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail ProjectEventClosed with empty action",
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
					onProjectEvent: map[string][]ProjectEventHandleFunc{
						ProjectEventAnyAction: {
							func(deliveryID string, eventName string, event *github.ProjectEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						ProjectEventClosedAction: {
							func(deliveryID string, eventName string, event *github.ProjectEvent) error {
								t.Logf("%s action called", ProjectEventClosedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "project",
				event:      &github.ProjectEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail ProjectEventClosed with nil action",
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
					onProjectEvent: map[string][]ProjectEventHandleFunc{
						ProjectEventAnyAction: {
							func(deliveryID string, eventName string, event *github.ProjectEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						ProjectEventClosedAction: {
							func(deliveryID string, eventName string, event *github.ProjectEvent) error {
								t.Logf("%s action called", ProjectEventClosedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "project",
				event:      &github.ProjectEvent{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger ProjectEventReopened",
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
					onProjectEvent: map[string][]ProjectEventHandleFunc{
						ProjectEventAnyAction: {
							func(deliveryID string, eventName string, event *github.ProjectEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						ProjectEventReopenedAction: {
							func(deliveryID string, eventName string, event *github.ProjectEvent) error {
								t.Logf("%s action called", ProjectEventReopenedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "project",
				event:      &github.ProjectEvent{Action: ptrString(ProjectEventReopenedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail ProjectEventReopened with empty action",
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
					onProjectEvent: map[string][]ProjectEventHandleFunc{
						ProjectEventAnyAction: {
							func(deliveryID string, eventName string, event *github.ProjectEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						ProjectEventReopenedAction: {
							func(deliveryID string, eventName string, event *github.ProjectEvent) error {
								t.Logf("%s action called", ProjectEventReopenedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "project",
				event:      &github.ProjectEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail ProjectEventReopened with nil action",
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
					onProjectEvent: map[string][]ProjectEventHandleFunc{
						ProjectEventAnyAction: {
							func(deliveryID string, eventName string, event *github.ProjectEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						ProjectEventReopenedAction: {
							func(deliveryID string, eventName string, event *github.ProjectEvent) error {
								t.Logf("%s action called", ProjectEventReopenedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "project",
				event:      &github.ProjectEvent{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger ProjectEventDeleted",
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
					onProjectEvent: map[string][]ProjectEventHandleFunc{
						ProjectEventAnyAction: {
							func(deliveryID string, eventName string, event *github.ProjectEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						ProjectEventDeletedAction: {
							func(deliveryID string, eventName string, event *github.ProjectEvent) error {
								t.Logf("%s action called", ProjectEventDeletedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "project",
				event:      &github.ProjectEvent{Action: ptrString(ProjectEventDeletedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail ProjectEventDeleted with empty action",
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
					onProjectEvent: map[string][]ProjectEventHandleFunc{
						ProjectEventAnyAction: {
							func(deliveryID string, eventName string, event *github.ProjectEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						ProjectEventDeletedAction: {
							func(deliveryID string, eventName string, event *github.ProjectEvent) error {
								t.Logf("%s action called", ProjectEventDeletedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "project",
				event:      &github.ProjectEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail ProjectEventDeleted with nil action",
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
					onProjectEvent: map[string][]ProjectEventHandleFunc{
						ProjectEventAnyAction: {
							func(deliveryID string, eventName string, event *github.ProjectEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						ProjectEventDeletedAction: {
							func(deliveryID string, eventName string, event *github.ProjectEvent) error {
								t.Logf("%s action called", ProjectEventDeletedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "project",
				event:      &github.ProjectEvent{Action: nil},
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
			if err := g.ProjectEvent(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("ProjectEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
