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

func TestOnProjectCardEventAny(t *testing.T) {
	type args struct {
		callbacks []ProjectCardEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single ProjectCardEventHandleFunc",
			args: args{
				[]ProjectCardEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple ProjectCardEventHandleFuncs",
			args: args{
				[]ProjectCardEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectCardEventAny(tt.args.callbacks...)
			if len(g.onProjectCardEvent[ProjectCardEventAnyAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onProjectCardEvent[ProjectCardEventAnyAction]))
			}
		})
	}
}

func TestSetOnProjectCardEventAny(t *testing.T) {
	type args struct {
		callbacks []ProjectCardEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single ProjectCardEventHandleFunc",
			args: args{
				[]ProjectCardEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple ProjectCardEventHandleFuncs",
			args: args{
				[]ProjectCardEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
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
			g.SetOnProjectCardEventAny(func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
				return nil
			})
			g.SetOnProjectCardEventAny(tt.args.callbacks...)
			if len(g.onProjectCardEvent[ProjectCardEventAnyAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onProjectCardEvent[ProjectCardEventAnyAction]), tt.want)
			}
		})
	}
}

func TestHandleProjectCardEventAny(t *testing.T) {

	action := "*"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.ProjectCardEvent
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
				eventName:  "project_card",

				event: &github.ProjectCardEvent{Action: &action},

				fail: false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "project_card",

				event: &github.ProjectCardEvent{Action: &action},

				fail: true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "project_card",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectCardEventAny(func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleProjectCardEventAny(tt.args.deliveryID, tt.args.deliveryID, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("TestHandleProjectCardEventAny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnProjectCardEventCreated(t *testing.T) {
	type args struct {
		callbacks []ProjectCardEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single ProjectCardEventHandleFunc",
			args: args{
				callbacks: []ProjectCardEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple ProjectCardEventHandleFunc",
			args: args{
				callbacks: []ProjectCardEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectCardEventCreated(tt.args.callbacks...)
			if len(g.onProjectCardEvent[ProjectCardEventCreatedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onProjectCardEvent[ProjectCardEventCreatedAction]))
			}
		})
	}
}

func TestSetOnProjectCardEventCreated(t *testing.T) {
	type args struct {
		callbacks []ProjectCardEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single ProjectCardEventHandleFunc",
			args: args{
				[]ProjectCardEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple ProjectCardEventHandleFuncs",
			args: args{
				[]ProjectCardEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
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
			g.SetOnProjectCardEventCreated(func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
				return nil
			})
			g.SetOnProjectCardEventCreated(tt.args.callbacks...)
			if len(g.onProjectCardEvent[ProjectCardEventCreatedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onProjectCardEvent[ProjectCardEventCreatedAction]), tt.want)
			}
		})
	}
}

func TestHandleProjectCardEventCreated(t *testing.T) {
	action := ProjectCardEventCreatedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.ProjectCardEvent
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
				eventName:  "project_card",
				event:      &github.ProjectCardEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "project_card",
				event:      &github.ProjectCardEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "project_card",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "project_card",
				event:      &github.ProjectCardEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "project_card",
				event:      &github.ProjectCardEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "project_card",
				event:      &github.ProjectCardEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectCardEventCreated(func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleProjectCardEventCreated(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleProjectCardEventCreated() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnProjectCardEventEdited(t *testing.T) {
	type args struct {
		callbacks []ProjectCardEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single ProjectCardEventHandleFunc",
			args: args{
				callbacks: []ProjectCardEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple ProjectCardEventHandleFunc",
			args: args{
				callbacks: []ProjectCardEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectCardEventEdited(tt.args.callbacks...)
			if len(g.onProjectCardEvent[ProjectCardEventEditedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onProjectCardEvent[ProjectCardEventEditedAction]))
			}
		})
	}
}

func TestSetOnProjectCardEventEdited(t *testing.T) {
	type args struct {
		callbacks []ProjectCardEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single ProjectCardEventHandleFunc",
			args: args{
				[]ProjectCardEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple ProjectCardEventHandleFuncs",
			args: args{
				[]ProjectCardEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
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
			g.SetOnProjectCardEventEdited(func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
				return nil
			})
			g.SetOnProjectCardEventEdited(tt.args.callbacks...)
			if len(g.onProjectCardEvent[ProjectCardEventEditedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onProjectCardEvent[ProjectCardEventEditedAction]), tt.want)
			}
		})
	}
}

func TestHandleProjectCardEventEdited(t *testing.T) {
	action := ProjectCardEventEditedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.ProjectCardEvent
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
				eventName:  "project_card",
				event:      &github.ProjectCardEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "project_card",
				event:      &github.ProjectCardEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "project_card",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "project_card",
				event:      &github.ProjectCardEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "project_card",
				event:      &github.ProjectCardEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "project_card",
				event:      &github.ProjectCardEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectCardEventEdited(func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleProjectCardEventEdited(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleProjectCardEventEdited() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnProjectCardEventConverted(t *testing.T) {
	type args struct {
		callbacks []ProjectCardEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single ProjectCardEventHandleFunc",
			args: args{
				callbacks: []ProjectCardEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple ProjectCardEventHandleFunc",
			args: args{
				callbacks: []ProjectCardEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectCardEventConverted(tt.args.callbacks...)
			if len(g.onProjectCardEvent[ProjectCardEventConvertedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onProjectCardEvent[ProjectCardEventConvertedAction]))
			}
		})
	}
}

func TestSetOnProjectCardEventConverted(t *testing.T) {
	type args struct {
		callbacks []ProjectCardEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single ProjectCardEventHandleFunc",
			args: args{
				[]ProjectCardEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple ProjectCardEventHandleFuncs",
			args: args{
				[]ProjectCardEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
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
			g.SetOnProjectCardEventConverted(func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
				return nil
			})
			g.SetOnProjectCardEventConverted(tt.args.callbacks...)
			if len(g.onProjectCardEvent[ProjectCardEventConvertedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onProjectCardEvent[ProjectCardEventConvertedAction]), tt.want)
			}
		})
	}
}

func TestHandleProjectCardEventConverted(t *testing.T) {
	action := ProjectCardEventConvertedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.ProjectCardEvent
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
				eventName:  "project_card",
				event:      &github.ProjectCardEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "project_card",
				event:      &github.ProjectCardEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "project_card",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "project_card",
				event:      &github.ProjectCardEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "project_card",
				event:      &github.ProjectCardEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "project_card",
				event:      &github.ProjectCardEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectCardEventConverted(func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleProjectCardEventConverted(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleProjectCardEventConverted() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnProjectCardEventMoved(t *testing.T) {
	type args struct {
		callbacks []ProjectCardEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single ProjectCardEventHandleFunc",
			args: args{
				callbacks: []ProjectCardEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple ProjectCardEventHandleFunc",
			args: args{
				callbacks: []ProjectCardEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectCardEventMoved(tt.args.callbacks...)
			if len(g.onProjectCardEvent[ProjectCardEventMovedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onProjectCardEvent[ProjectCardEventMovedAction]))
			}
		})
	}
}

func TestSetOnProjectCardEventMoved(t *testing.T) {
	type args struct {
		callbacks []ProjectCardEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single ProjectCardEventHandleFunc",
			args: args{
				[]ProjectCardEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple ProjectCardEventHandleFuncs",
			args: args{
				[]ProjectCardEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
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
			g.SetOnProjectCardEventMoved(func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
				return nil
			})
			g.SetOnProjectCardEventMoved(tt.args.callbacks...)
			if len(g.onProjectCardEvent[ProjectCardEventMovedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onProjectCardEvent[ProjectCardEventMovedAction]), tt.want)
			}
		})
	}
}

func TestHandleProjectCardEventMoved(t *testing.T) {
	action := ProjectCardEventMovedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.ProjectCardEvent
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
				eventName:  "project_card",
				event:      &github.ProjectCardEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "project_card",
				event:      &github.ProjectCardEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "project_card",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "project_card",
				event:      &github.ProjectCardEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "project_card",
				event:      &github.ProjectCardEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "project_card",
				event:      &github.ProjectCardEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectCardEventMoved(func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleProjectCardEventMoved(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleProjectCardEventMoved() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnProjectCardEventDeleted(t *testing.T) {
	type args struct {
		callbacks []ProjectCardEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single ProjectCardEventHandleFunc",
			args: args{
				callbacks: []ProjectCardEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple ProjectCardEventHandleFunc",
			args: args{
				callbacks: []ProjectCardEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectCardEventDeleted(tt.args.callbacks...)
			if len(g.onProjectCardEvent[ProjectCardEventDeletedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onProjectCardEvent[ProjectCardEventDeletedAction]))
			}
		})
	}
}

func TestSetOnProjectCardEventDeleted(t *testing.T) {
	type args struct {
		callbacks []ProjectCardEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single ProjectCardEventHandleFunc",
			args: args{
				[]ProjectCardEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple ProjectCardEventHandleFuncs",
			args: args{
				[]ProjectCardEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
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
			g.SetOnProjectCardEventDeleted(func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
				return nil
			})
			g.SetOnProjectCardEventDeleted(tt.args.callbacks...)
			if len(g.onProjectCardEvent[ProjectCardEventDeletedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onProjectCardEvent[ProjectCardEventDeletedAction]), tt.want)
			}
		})
	}
}

func TestHandleProjectCardEventDeleted(t *testing.T) {
	action := ProjectCardEventDeletedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.ProjectCardEvent
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
				eventName:  "project_card",
				event:      &github.ProjectCardEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "project_card",
				event:      &github.ProjectCardEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "project_card",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "project_card",
				event:      &github.ProjectCardEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "project_card",
				event:      &github.ProjectCardEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "project_card",
				event:      &github.ProjectCardEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectCardEventDeleted(func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleProjectCardEventDeleted(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleProjectCardEventDeleted() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProjectCardEvent(t *testing.T) {
	type fields struct {
		handler *EventHandler
	}
	type args struct {
		deliveryID string
		eventName  string
		event      *github.ProjectCardEvent
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "must trigger ProjectCardEventAny with unknown event action",
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
					onProjectCardEvent: map[string][]ProjectCardEventHandleFunc{
						ProjectCardEventAnyAction: {
							func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  ProjectCardEvent,

				event: &github.ProjectCardEvent{Action: ptrString("unknown")},
			},
			wantErr: false,
		},

		{
			name: "must trigger ProjectCardEventCreated",
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
					onProjectCardEvent: map[string][]ProjectCardEventHandleFunc{
						ProjectCardEventAnyAction: {
							func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						ProjectCardEventCreatedAction: {
							func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
								t.Logf("%s action called", ProjectCardEventCreatedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "project_card",
				event:      &github.ProjectCardEvent{Action: ptrString(ProjectCardEventCreatedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail ProjectCardEventCreated with empty action",
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
					onProjectCardEvent: map[string][]ProjectCardEventHandleFunc{
						ProjectCardEventAnyAction: {
							func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						ProjectCardEventCreatedAction: {
							func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
								t.Logf("%s action called", ProjectCardEventCreatedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "project_card",
				event:      &github.ProjectCardEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail ProjectCardEventCreated with nil action",
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
					onProjectCardEvent: map[string][]ProjectCardEventHandleFunc{
						ProjectCardEventAnyAction: {
							func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						ProjectCardEventCreatedAction: {
							func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
								t.Logf("%s action called", ProjectCardEventCreatedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "project_card",
				event:      &github.ProjectCardEvent{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger ProjectCardEventEdited",
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
					onProjectCardEvent: map[string][]ProjectCardEventHandleFunc{
						ProjectCardEventAnyAction: {
							func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						ProjectCardEventEditedAction: {
							func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
								t.Logf("%s action called", ProjectCardEventEditedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "project_card",
				event:      &github.ProjectCardEvent{Action: ptrString(ProjectCardEventEditedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail ProjectCardEventEdited with empty action",
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
					onProjectCardEvent: map[string][]ProjectCardEventHandleFunc{
						ProjectCardEventAnyAction: {
							func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						ProjectCardEventEditedAction: {
							func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
								t.Logf("%s action called", ProjectCardEventEditedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "project_card",
				event:      &github.ProjectCardEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail ProjectCardEventEdited with nil action",
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
					onProjectCardEvent: map[string][]ProjectCardEventHandleFunc{
						ProjectCardEventAnyAction: {
							func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						ProjectCardEventEditedAction: {
							func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
								t.Logf("%s action called", ProjectCardEventEditedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "project_card",
				event:      &github.ProjectCardEvent{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger ProjectCardEventConverted",
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
					onProjectCardEvent: map[string][]ProjectCardEventHandleFunc{
						ProjectCardEventAnyAction: {
							func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						ProjectCardEventConvertedAction: {
							func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
								t.Logf("%s action called", ProjectCardEventConvertedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "project_card",
				event:      &github.ProjectCardEvent{Action: ptrString(ProjectCardEventConvertedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail ProjectCardEventConverted with empty action",
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
					onProjectCardEvent: map[string][]ProjectCardEventHandleFunc{
						ProjectCardEventAnyAction: {
							func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						ProjectCardEventConvertedAction: {
							func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
								t.Logf("%s action called", ProjectCardEventConvertedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "project_card",
				event:      &github.ProjectCardEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail ProjectCardEventConverted with nil action",
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
					onProjectCardEvent: map[string][]ProjectCardEventHandleFunc{
						ProjectCardEventAnyAction: {
							func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						ProjectCardEventConvertedAction: {
							func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
								t.Logf("%s action called", ProjectCardEventConvertedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "project_card",
				event:      &github.ProjectCardEvent{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger ProjectCardEventMoved",
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
					onProjectCardEvent: map[string][]ProjectCardEventHandleFunc{
						ProjectCardEventAnyAction: {
							func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						ProjectCardEventMovedAction: {
							func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
								t.Logf("%s action called", ProjectCardEventMovedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "project_card",
				event:      &github.ProjectCardEvent{Action: ptrString(ProjectCardEventMovedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail ProjectCardEventMoved with empty action",
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
					onProjectCardEvent: map[string][]ProjectCardEventHandleFunc{
						ProjectCardEventAnyAction: {
							func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						ProjectCardEventMovedAction: {
							func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
								t.Logf("%s action called", ProjectCardEventMovedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "project_card",
				event:      &github.ProjectCardEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail ProjectCardEventMoved with nil action",
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
					onProjectCardEvent: map[string][]ProjectCardEventHandleFunc{
						ProjectCardEventAnyAction: {
							func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						ProjectCardEventMovedAction: {
							func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
								t.Logf("%s action called", ProjectCardEventMovedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "project_card",
				event:      &github.ProjectCardEvent{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger ProjectCardEventDeleted",
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
					onProjectCardEvent: map[string][]ProjectCardEventHandleFunc{
						ProjectCardEventAnyAction: {
							func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						ProjectCardEventDeletedAction: {
							func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
								t.Logf("%s action called", ProjectCardEventDeletedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "project_card",
				event:      &github.ProjectCardEvent{Action: ptrString(ProjectCardEventDeletedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail ProjectCardEventDeleted with empty action",
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
					onProjectCardEvent: map[string][]ProjectCardEventHandleFunc{
						ProjectCardEventAnyAction: {
							func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						ProjectCardEventDeletedAction: {
							func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
								t.Logf("%s action called", ProjectCardEventDeletedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "project_card",
				event:      &github.ProjectCardEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail ProjectCardEventDeleted with nil action",
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
					onProjectCardEvent: map[string][]ProjectCardEventHandleFunc{
						ProjectCardEventAnyAction: {
							func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						ProjectCardEventDeletedAction: {
							func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
								t.Logf("%s action called", ProjectCardEventDeletedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "project_card",
				event:      &github.ProjectCardEvent{Action: nil},
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
			if err := g.ProjectCardEvent(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("ProjectCardEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
