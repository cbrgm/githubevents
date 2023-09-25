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

func TestOnRepositoryEventAny(t *testing.T) {
	type args struct {
		callbacks []RepositoryEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single RepositoryEventHandleFunc",
			args: args{
				[]RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple RepositoryEventHandleFuncs",
			args: args{
				[]RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnRepositoryEventAny(tt.args.callbacks...)
			if len(g.onRepositoryEvent[RepositoryEventAnyAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onRepositoryEvent[RepositoryEventAnyAction]))
			}
		})
	}
}

func TestSetOnRepositoryEventAny(t *testing.T) {
	type args struct {
		callbacks []RepositoryEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single RepositoryEventHandleFunc",
			args: args{
				[]RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple RepositoryEventHandleFuncs",
			args: args{
				[]RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
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
			g.SetOnRepositoryEventAny(func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
				return nil
			})
			g.SetOnRepositoryEventAny(tt.args.callbacks...)
			if len(g.onRepositoryEvent[RepositoryEventAnyAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onRepositoryEvent[RepositoryEventAnyAction]), tt.want)
			}
		})
	}
}

func TestHandleRepositoryEventAny(t *testing.T) {

	action := "*"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.RepositoryEvent
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
				eventName:  "repository",

				event: &github.RepositoryEvent{Action: &action},

				fail: false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "repository",

				event: &github.RepositoryEvent{Action: &action},

				fail: true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnRepositoryEventAny(func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleRepositoryEventAny(tt.args.deliveryID, tt.args.deliveryID, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("TestHandleRepositoryEventAny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnRepositoryEventCreated(t *testing.T) {
	type args struct {
		callbacks []RepositoryEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single RepositoryEventHandleFunc",
			args: args{
				callbacks: []RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple RepositoryEventHandleFunc",
			args: args{
				callbacks: []RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnRepositoryEventCreated(tt.args.callbacks...)
			if len(g.onRepositoryEvent[RepositoryEventCreatedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onRepositoryEvent[RepositoryEventCreatedAction]))
			}
		})
	}
}

func TestSetOnRepositoryEventCreated(t *testing.T) {
	type args struct {
		callbacks []RepositoryEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single RepositoryEventHandleFunc",
			args: args{
				[]RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple RepositoryEventHandleFuncs",
			args: args{
				[]RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
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
			g.SetOnRepositoryEventCreated(func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
				return nil
			})
			g.SetOnRepositoryEventCreated(tt.args.callbacks...)
			if len(g.onRepositoryEvent[RepositoryEventCreatedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onRepositoryEvent[RepositoryEventCreatedAction]), tt.want)
			}
		})
	}
}

func TestHandleRepositoryEventCreated(t *testing.T) {
	action := RepositoryEventCreatedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.RepositoryEvent
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
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnRepositoryEventCreated(func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleRepositoryEventCreated(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleRepositoryEventCreated() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnRepositoryEventDeleted(t *testing.T) {
	type args struct {
		callbacks []RepositoryEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single RepositoryEventHandleFunc",
			args: args{
				callbacks: []RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple RepositoryEventHandleFunc",
			args: args{
				callbacks: []RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnRepositoryEventDeleted(tt.args.callbacks...)
			if len(g.onRepositoryEvent[RepositoryEventDeletedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onRepositoryEvent[RepositoryEventDeletedAction]))
			}
		})
	}
}

func TestSetOnRepositoryEventDeleted(t *testing.T) {
	type args struct {
		callbacks []RepositoryEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single RepositoryEventHandleFunc",
			args: args{
				[]RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple RepositoryEventHandleFuncs",
			args: args{
				[]RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
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
			g.SetOnRepositoryEventDeleted(func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
				return nil
			})
			g.SetOnRepositoryEventDeleted(tt.args.callbacks...)
			if len(g.onRepositoryEvent[RepositoryEventDeletedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onRepositoryEvent[RepositoryEventDeletedAction]), tt.want)
			}
		})
	}
}

func TestHandleRepositoryEventDeleted(t *testing.T) {
	action := RepositoryEventDeletedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.RepositoryEvent
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
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnRepositoryEventDeleted(func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleRepositoryEventDeleted(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleRepositoryEventDeleted() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnRepositoryEventArchived(t *testing.T) {
	type args struct {
		callbacks []RepositoryEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single RepositoryEventHandleFunc",
			args: args{
				callbacks: []RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple RepositoryEventHandleFunc",
			args: args{
				callbacks: []RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnRepositoryEventArchived(tt.args.callbacks...)
			if len(g.onRepositoryEvent[RepositoryEventArchivedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onRepositoryEvent[RepositoryEventArchivedAction]))
			}
		})
	}
}

func TestSetOnRepositoryEventArchived(t *testing.T) {
	type args struct {
		callbacks []RepositoryEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single RepositoryEventHandleFunc",
			args: args{
				[]RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple RepositoryEventHandleFuncs",
			args: args{
				[]RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
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
			g.SetOnRepositoryEventArchived(func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
				return nil
			})
			g.SetOnRepositoryEventArchived(tt.args.callbacks...)
			if len(g.onRepositoryEvent[RepositoryEventArchivedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onRepositoryEvent[RepositoryEventArchivedAction]), tt.want)
			}
		})
	}
}

func TestHandleRepositoryEventArchived(t *testing.T) {
	action := RepositoryEventArchivedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.RepositoryEvent
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
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnRepositoryEventArchived(func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleRepositoryEventArchived(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleRepositoryEventArchived() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnRepositoryEventUnarchived(t *testing.T) {
	type args struct {
		callbacks []RepositoryEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single RepositoryEventHandleFunc",
			args: args{
				callbacks: []RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple RepositoryEventHandleFunc",
			args: args{
				callbacks: []RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnRepositoryEventUnarchived(tt.args.callbacks...)
			if len(g.onRepositoryEvent[RepositoryEventUnarchivedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onRepositoryEvent[RepositoryEventUnarchivedAction]))
			}
		})
	}
}

func TestSetOnRepositoryEventUnarchived(t *testing.T) {
	type args struct {
		callbacks []RepositoryEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single RepositoryEventHandleFunc",
			args: args{
				[]RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple RepositoryEventHandleFuncs",
			args: args{
				[]RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
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
			g.SetOnRepositoryEventUnarchived(func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
				return nil
			})
			g.SetOnRepositoryEventUnarchived(tt.args.callbacks...)
			if len(g.onRepositoryEvent[RepositoryEventUnarchivedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onRepositoryEvent[RepositoryEventUnarchivedAction]), tt.want)
			}
		})
	}
}

func TestHandleRepositoryEventUnarchived(t *testing.T) {
	action := RepositoryEventUnarchivedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.RepositoryEvent
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
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnRepositoryEventUnarchived(func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleRepositoryEventUnarchived(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleRepositoryEventUnarchived() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnRepositoryEventEdited(t *testing.T) {
	type args struct {
		callbacks []RepositoryEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single RepositoryEventHandleFunc",
			args: args{
				callbacks: []RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple RepositoryEventHandleFunc",
			args: args{
				callbacks: []RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnRepositoryEventEdited(tt.args.callbacks...)
			if len(g.onRepositoryEvent[RepositoryEventEditedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onRepositoryEvent[RepositoryEventEditedAction]))
			}
		})
	}
}

func TestSetOnRepositoryEventEdited(t *testing.T) {
	type args struct {
		callbacks []RepositoryEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single RepositoryEventHandleFunc",
			args: args{
				[]RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple RepositoryEventHandleFuncs",
			args: args{
				[]RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
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
			g.SetOnRepositoryEventEdited(func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
				return nil
			})
			g.SetOnRepositoryEventEdited(tt.args.callbacks...)
			if len(g.onRepositoryEvent[RepositoryEventEditedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onRepositoryEvent[RepositoryEventEditedAction]), tt.want)
			}
		})
	}
}

func TestHandleRepositoryEventEdited(t *testing.T) {
	action := RepositoryEventEditedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.RepositoryEvent
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
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnRepositoryEventEdited(func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleRepositoryEventEdited(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleRepositoryEventEdited() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnRepositoryEventRenamed(t *testing.T) {
	type args struct {
		callbacks []RepositoryEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single RepositoryEventHandleFunc",
			args: args{
				callbacks: []RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple RepositoryEventHandleFunc",
			args: args{
				callbacks: []RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnRepositoryEventRenamed(tt.args.callbacks...)
			if len(g.onRepositoryEvent[RepositoryEventRenamedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onRepositoryEvent[RepositoryEventRenamedAction]))
			}
		})
	}
}

func TestSetOnRepositoryEventRenamed(t *testing.T) {
	type args struct {
		callbacks []RepositoryEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single RepositoryEventHandleFunc",
			args: args{
				[]RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple RepositoryEventHandleFuncs",
			args: args{
				[]RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
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
			g.SetOnRepositoryEventRenamed(func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
				return nil
			})
			g.SetOnRepositoryEventRenamed(tt.args.callbacks...)
			if len(g.onRepositoryEvent[RepositoryEventRenamedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onRepositoryEvent[RepositoryEventRenamedAction]), tt.want)
			}
		})
	}
}

func TestHandleRepositoryEventRenamed(t *testing.T) {
	action := RepositoryEventRenamedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.RepositoryEvent
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
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnRepositoryEventRenamed(func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleRepositoryEventRenamed(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleRepositoryEventRenamed() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnRepositoryEventTransferred(t *testing.T) {
	type args struct {
		callbacks []RepositoryEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single RepositoryEventHandleFunc",
			args: args{
				callbacks: []RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple RepositoryEventHandleFunc",
			args: args{
				callbacks: []RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnRepositoryEventTransferred(tt.args.callbacks...)
			if len(g.onRepositoryEvent[RepositoryEventTransferredAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onRepositoryEvent[RepositoryEventTransferredAction]))
			}
		})
	}
}

func TestSetOnRepositoryEventTransferred(t *testing.T) {
	type args struct {
		callbacks []RepositoryEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single RepositoryEventHandleFunc",
			args: args{
				[]RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple RepositoryEventHandleFuncs",
			args: args{
				[]RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
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
			g.SetOnRepositoryEventTransferred(func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
				return nil
			})
			g.SetOnRepositoryEventTransferred(tt.args.callbacks...)
			if len(g.onRepositoryEvent[RepositoryEventTransferredAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onRepositoryEvent[RepositoryEventTransferredAction]), tt.want)
			}
		})
	}
}

func TestHandleRepositoryEventTransferred(t *testing.T) {
	action := RepositoryEventTransferredAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.RepositoryEvent
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
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnRepositoryEventTransferred(func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleRepositoryEventTransferred(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleRepositoryEventTransferred() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnRepositoryEventPublicized(t *testing.T) {
	type args struct {
		callbacks []RepositoryEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single RepositoryEventHandleFunc",
			args: args{
				callbacks: []RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple RepositoryEventHandleFunc",
			args: args{
				callbacks: []RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnRepositoryEventPublicized(tt.args.callbacks...)
			if len(g.onRepositoryEvent[RepositoryEventPublicizedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onRepositoryEvent[RepositoryEventPublicizedAction]))
			}
		})
	}
}

func TestSetOnRepositoryEventPublicized(t *testing.T) {
	type args struct {
		callbacks []RepositoryEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single RepositoryEventHandleFunc",
			args: args{
				[]RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple RepositoryEventHandleFuncs",
			args: args{
				[]RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
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
			g.SetOnRepositoryEventPublicized(func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
				return nil
			})
			g.SetOnRepositoryEventPublicized(tt.args.callbacks...)
			if len(g.onRepositoryEvent[RepositoryEventPublicizedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onRepositoryEvent[RepositoryEventPublicizedAction]), tt.want)
			}
		})
	}
}

func TestHandleRepositoryEventPublicized(t *testing.T) {
	action := RepositoryEventPublicizedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.RepositoryEvent
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
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnRepositoryEventPublicized(func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleRepositoryEventPublicized(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleRepositoryEventPublicized() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnRepositoryEventPrivatized(t *testing.T) {
	type args struct {
		callbacks []RepositoryEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single RepositoryEventHandleFunc",
			args: args{
				callbacks: []RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple RepositoryEventHandleFunc",
			args: args{
				callbacks: []RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnRepositoryEventPrivatized(tt.args.callbacks...)
			if len(g.onRepositoryEvent[RepositoryEventPrivatizedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onRepositoryEvent[RepositoryEventPrivatizedAction]))
			}
		})
	}
}

func TestSetOnRepositoryEventPrivatized(t *testing.T) {
	type args struct {
		callbacks []RepositoryEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single RepositoryEventHandleFunc",
			args: args{
				[]RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple RepositoryEventHandleFuncs",
			args: args{
				[]RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
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
			g.SetOnRepositoryEventPrivatized(func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
				return nil
			})
			g.SetOnRepositoryEventPrivatized(tt.args.callbacks...)
			if len(g.onRepositoryEvent[RepositoryEventPrivatizedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onRepositoryEvent[RepositoryEventPrivatizedAction]), tt.want)
			}
		})
	}
}

func TestHandleRepositoryEventPrivatized(t *testing.T) {
	action := RepositoryEventPrivatizedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.RepositoryEvent
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
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnRepositoryEventPrivatized(func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleRepositoryEventPrivatized(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleRepositoryEventPrivatized() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRepositoryEvent(t *testing.T) {
	type fields struct {
		handler *EventHandler
	}
	type args struct {
		deliveryID string
		eventName  string
		event      *github.RepositoryEvent
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "must trigger RepositoryEventAny with unknown event action",
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
					onRepositoryEvent: map[string][]RepositoryEventHandleFunc{
						RepositoryEventAnyAction: {
							func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  RepositoryEvent,

				event: &github.RepositoryEvent{Action: ptrString("unknown")},
			},
			wantErr: false,
		},

		{
			name: "must trigger RepositoryEventCreated",
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
					onRepositoryEvent: map[string][]RepositoryEventHandleFunc{
						RepositoryEventAnyAction: {
							func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						RepositoryEventCreatedAction: {
							func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
								t.Logf("%s action called", RepositoryEventCreatedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: ptrString(RepositoryEventCreatedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail RepositoryEventCreated with empty action",
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
					onRepositoryEvent: map[string][]RepositoryEventHandleFunc{
						RepositoryEventAnyAction: {
							func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						RepositoryEventCreatedAction: {
							func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
								t.Logf("%s action called", RepositoryEventCreatedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail RepositoryEventCreated with nil action",
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
					onRepositoryEvent: map[string][]RepositoryEventHandleFunc{
						RepositoryEventAnyAction: {
							func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						RepositoryEventCreatedAction: {
							func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
								t.Logf("%s action called", RepositoryEventCreatedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger RepositoryEventDeleted",
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
					onRepositoryEvent: map[string][]RepositoryEventHandleFunc{
						RepositoryEventAnyAction: {
							func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						RepositoryEventDeletedAction: {
							func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
								t.Logf("%s action called", RepositoryEventDeletedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: ptrString(RepositoryEventDeletedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail RepositoryEventDeleted with empty action",
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
					onRepositoryEvent: map[string][]RepositoryEventHandleFunc{
						RepositoryEventAnyAction: {
							func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						RepositoryEventDeletedAction: {
							func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
								t.Logf("%s action called", RepositoryEventDeletedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail RepositoryEventDeleted with nil action",
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
					onRepositoryEvent: map[string][]RepositoryEventHandleFunc{
						RepositoryEventAnyAction: {
							func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						RepositoryEventDeletedAction: {
							func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
								t.Logf("%s action called", RepositoryEventDeletedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger RepositoryEventArchived",
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
					onRepositoryEvent: map[string][]RepositoryEventHandleFunc{
						RepositoryEventAnyAction: {
							func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						RepositoryEventArchivedAction: {
							func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
								t.Logf("%s action called", RepositoryEventArchivedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: ptrString(RepositoryEventArchivedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail RepositoryEventArchived with empty action",
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
					onRepositoryEvent: map[string][]RepositoryEventHandleFunc{
						RepositoryEventAnyAction: {
							func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						RepositoryEventArchivedAction: {
							func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
								t.Logf("%s action called", RepositoryEventArchivedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail RepositoryEventArchived with nil action",
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
					onRepositoryEvent: map[string][]RepositoryEventHandleFunc{
						RepositoryEventAnyAction: {
							func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						RepositoryEventArchivedAction: {
							func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
								t.Logf("%s action called", RepositoryEventArchivedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger RepositoryEventUnarchived",
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
					onRepositoryEvent: map[string][]RepositoryEventHandleFunc{
						RepositoryEventAnyAction: {
							func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						RepositoryEventUnarchivedAction: {
							func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
								t.Logf("%s action called", RepositoryEventUnarchivedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: ptrString(RepositoryEventUnarchivedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail RepositoryEventUnarchived with empty action",
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
					onRepositoryEvent: map[string][]RepositoryEventHandleFunc{
						RepositoryEventAnyAction: {
							func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						RepositoryEventUnarchivedAction: {
							func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
								t.Logf("%s action called", RepositoryEventUnarchivedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail RepositoryEventUnarchived with nil action",
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
					onRepositoryEvent: map[string][]RepositoryEventHandleFunc{
						RepositoryEventAnyAction: {
							func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						RepositoryEventUnarchivedAction: {
							func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
								t.Logf("%s action called", RepositoryEventUnarchivedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger RepositoryEventEdited",
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
					onRepositoryEvent: map[string][]RepositoryEventHandleFunc{
						RepositoryEventAnyAction: {
							func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						RepositoryEventEditedAction: {
							func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
								t.Logf("%s action called", RepositoryEventEditedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: ptrString(RepositoryEventEditedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail RepositoryEventEdited with empty action",
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
					onRepositoryEvent: map[string][]RepositoryEventHandleFunc{
						RepositoryEventAnyAction: {
							func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						RepositoryEventEditedAction: {
							func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
								t.Logf("%s action called", RepositoryEventEditedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail RepositoryEventEdited with nil action",
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
					onRepositoryEvent: map[string][]RepositoryEventHandleFunc{
						RepositoryEventAnyAction: {
							func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						RepositoryEventEditedAction: {
							func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
								t.Logf("%s action called", RepositoryEventEditedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger RepositoryEventRenamed",
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
					onRepositoryEvent: map[string][]RepositoryEventHandleFunc{
						RepositoryEventAnyAction: {
							func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						RepositoryEventRenamedAction: {
							func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
								t.Logf("%s action called", RepositoryEventRenamedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: ptrString(RepositoryEventRenamedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail RepositoryEventRenamed with empty action",
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
					onRepositoryEvent: map[string][]RepositoryEventHandleFunc{
						RepositoryEventAnyAction: {
							func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						RepositoryEventRenamedAction: {
							func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
								t.Logf("%s action called", RepositoryEventRenamedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail RepositoryEventRenamed with nil action",
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
					onRepositoryEvent: map[string][]RepositoryEventHandleFunc{
						RepositoryEventAnyAction: {
							func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						RepositoryEventRenamedAction: {
							func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
								t.Logf("%s action called", RepositoryEventRenamedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger RepositoryEventTransferred",
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
					onRepositoryEvent: map[string][]RepositoryEventHandleFunc{
						RepositoryEventAnyAction: {
							func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						RepositoryEventTransferredAction: {
							func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
								t.Logf("%s action called", RepositoryEventTransferredAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: ptrString(RepositoryEventTransferredAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail RepositoryEventTransferred with empty action",
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
					onRepositoryEvent: map[string][]RepositoryEventHandleFunc{
						RepositoryEventAnyAction: {
							func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						RepositoryEventTransferredAction: {
							func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
								t.Logf("%s action called", RepositoryEventTransferredAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail RepositoryEventTransferred with nil action",
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
					onRepositoryEvent: map[string][]RepositoryEventHandleFunc{
						RepositoryEventAnyAction: {
							func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						RepositoryEventTransferredAction: {
							func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
								t.Logf("%s action called", RepositoryEventTransferredAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger RepositoryEventPublicized",
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
					onRepositoryEvent: map[string][]RepositoryEventHandleFunc{
						RepositoryEventAnyAction: {
							func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						RepositoryEventPublicizedAction: {
							func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
								t.Logf("%s action called", RepositoryEventPublicizedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: ptrString(RepositoryEventPublicizedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail RepositoryEventPublicized with empty action",
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
					onRepositoryEvent: map[string][]RepositoryEventHandleFunc{
						RepositoryEventAnyAction: {
							func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						RepositoryEventPublicizedAction: {
							func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
								t.Logf("%s action called", RepositoryEventPublicizedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail RepositoryEventPublicized with nil action",
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
					onRepositoryEvent: map[string][]RepositoryEventHandleFunc{
						RepositoryEventAnyAction: {
							func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						RepositoryEventPublicizedAction: {
							func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
								t.Logf("%s action called", RepositoryEventPublicizedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger RepositoryEventPrivatized",
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
					onRepositoryEvent: map[string][]RepositoryEventHandleFunc{
						RepositoryEventAnyAction: {
							func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						RepositoryEventPrivatizedAction: {
							func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
								t.Logf("%s action called", RepositoryEventPrivatizedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: ptrString(RepositoryEventPrivatizedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail RepositoryEventPrivatized with empty action",
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
					onRepositoryEvent: map[string][]RepositoryEventHandleFunc{
						RepositoryEventAnyAction: {
							func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						RepositoryEventPrivatizedAction: {
							func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
								t.Logf("%s action called", RepositoryEventPrivatizedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail RepositoryEventPrivatized with nil action",
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
					onRepositoryEvent: map[string][]RepositoryEventHandleFunc{
						RepositoryEventAnyAction: {
							func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						RepositoryEventPrivatizedAction: {
							func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
								t.Logf("%s action called", RepositoryEventPrivatizedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: nil},
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
			if err := g.RepositoryEvent(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("RepositoryEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
