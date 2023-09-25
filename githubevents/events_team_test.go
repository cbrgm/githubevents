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

func TestOnTeamEventAny(t *testing.T) {
	type args struct {
		callbacks []TeamEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single TeamEventHandleFunc",
			args: args{
				[]TeamEventHandleFunc{
					func(deliveryID string, eventName string, event *github.TeamEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple TeamEventHandleFuncs",
			args: args{
				[]TeamEventHandleFunc{
					func(deliveryID string, eventName string, event *github.TeamEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.TeamEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnTeamEventAny(tt.args.callbacks...)
			if len(g.onTeamEvent[TeamEventAnyAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onTeamEvent[TeamEventAnyAction]))
			}
		})
	}
}

func TestSetOnTeamEventAny(t *testing.T) {
	type args struct {
		callbacks []TeamEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single TeamEventHandleFunc",
			args: args{
				[]TeamEventHandleFunc{
					func(deliveryID string, eventName string, event *github.TeamEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple TeamEventHandleFuncs",
			args: args{
				[]TeamEventHandleFunc{
					func(deliveryID string, eventName string, event *github.TeamEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.TeamEvent) error {
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
			g.SetOnTeamEventAny(func(deliveryID string, eventName string, event *github.TeamEvent) error {
				return nil
			})
			g.SetOnTeamEventAny(tt.args.callbacks...)
			if len(g.onTeamEvent[TeamEventAnyAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onTeamEvent[TeamEventAnyAction]), tt.want)
			}
		})
	}
}

func TestHandleTeamEventAny(t *testing.T) {

	action := "*"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.TeamEvent
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
				eventName:  "team",

				event: &github.TeamEvent{Action: &action},

				fail: false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "team",

				event: &github.TeamEvent{Action: &action},

				fail: true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "team",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnTeamEventAny(func(deliveryID string, eventName string, event *github.TeamEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleTeamEventAny(tt.args.deliveryID, tt.args.deliveryID, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("TestHandleTeamEventAny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnTeamEventCreated(t *testing.T) {
	type args struct {
		callbacks []TeamEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single TeamEventHandleFunc",
			args: args{
				callbacks: []TeamEventHandleFunc{
					func(deliveryID string, eventName string, event *github.TeamEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple TeamEventHandleFunc",
			args: args{
				callbacks: []TeamEventHandleFunc{
					func(deliveryID string, eventName string, event *github.TeamEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.TeamEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnTeamEventCreated(tt.args.callbacks...)
			if len(g.onTeamEvent[TeamEventCreatedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onTeamEvent[TeamEventCreatedAction]))
			}
		})
	}
}

func TestSetOnTeamEventCreated(t *testing.T) {
	type args struct {
		callbacks []TeamEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single TeamEventHandleFunc",
			args: args{
				[]TeamEventHandleFunc{
					func(deliveryID string, eventName string, event *github.TeamEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple TeamEventHandleFuncs",
			args: args{
				[]TeamEventHandleFunc{
					func(deliveryID string, eventName string, event *github.TeamEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.TeamEvent) error {
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
			g.SetOnTeamEventCreated(func(deliveryID string, eventName string, event *github.TeamEvent) error {
				return nil
			})
			g.SetOnTeamEventCreated(tt.args.callbacks...)
			if len(g.onTeamEvent[TeamEventCreatedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onTeamEvent[TeamEventCreatedAction]), tt.want)
			}
		})
	}
}

func TestHandleTeamEventCreated(t *testing.T) {
	action := TeamEventCreatedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.TeamEvent
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
				eventName:  "team",
				event:      &github.TeamEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "team",
				event:      &github.TeamEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "team",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "team",
				event:      &github.TeamEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "team",
				event:      &github.TeamEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "team",
				event:      &github.TeamEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnTeamEventCreated(func(deliveryID string, eventName string, event *github.TeamEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleTeamEventCreated(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleTeamEventCreated() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnTeamEventDeleted(t *testing.T) {
	type args struct {
		callbacks []TeamEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single TeamEventHandleFunc",
			args: args{
				callbacks: []TeamEventHandleFunc{
					func(deliveryID string, eventName string, event *github.TeamEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple TeamEventHandleFunc",
			args: args{
				callbacks: []TeamEventHandleFunc{
					func(deliveryID string, eventName string, event *github.TeamEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.TeamEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnTeamEventDeleted(tt.args.callbacks...)
			if len(g.onTeamEvent[TeamEventDeletedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onTeamEvent[TeamEventDeletedAction]))
			}
		})
	}
}

func TestSetOnTeamEventDeleted(t *testing.T) {
	type args struct {
		callbacks []TeamEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single TeamEventHandleFunc",
			args: args{
				[]TeamEventHandleFunc{
					func(deliveryID string, eventName string, event *github.TeamEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple TeamEventHandleFuncs",
			args: args{
				[]TeamEventHandleFunc{
					func(deliveryID string, eventName string, event *github.TeamEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.TeamEvent) error {
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
			g.SetOnTeamEventDeleted(func(deliveryID string, eventName string, event *github.TeamEvent) error {
				return nil
			})
			g.SetOnTeamEventDeleted(tt.args.callbacks...)
			if len(g.onTeamEvent[TeamEventDeletedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onTeamEvent[TeamEventDeletedAction]), tt.want)
			}
		})
	}
}

func TestHandleTeamEventDeleted(t *testing.T) {
	action := TeamEventDeletedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.TeamEvent
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
				eventName:  "team",
				event:      &github.TeamEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "team",
				event:      &github.TeamEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "team",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "team",
				event:      &github.TeamEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "team",
				event:      &github.TeamEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "team",
				event:      &github.TeamEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnTeamEventDeleted(func(deliveryID string, eventName string, event *github.TeamEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleTeamEventDeleted(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleTeamEventDeleted() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnTeamEventEdited(t *testing.T) {
	type args struct {
		callbacks []TeamEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single TeamEventHandleFunc",
			args: args{
				callbacks: []TeamEventHandleFunc{
					func(deliveryID string, eventName string, event *github.TeamEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple TeamEventHandleFunc",
			args: args{
				callbacks: []TeamEventHandleFunc{
					func(deliveryID string, eventName string, event *github.TeamEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.TeamEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnTeamEventEdited(tt.args.callbacks...)
			if len(g.onTeamEvent[TeamEventEditedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onTeamEvent[TeamEventEditedAction]))
			}
		})
	}
}

func TestSetOnTeamEventEdited(t *testing.T) {
	type args struct {
		callbacks []TeamEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single TeamEventHandleFunc",
			args: args{
				[]TeamEventHandleFunc{
					func(deliveryID string, eventName string, event *github.TeamEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple TeamEventHandleFuncs",
			args: args{
				[]TeamEventHandleFunc{
					func(deliveryID string, eventName string, event *github.TeamEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.TeamEvent) error {
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
			g.SetOnTeamEventEdited(func(deliveryID string, eventName string, event *github.TeamEvent) error {
				return nil
			})
			g.SetOnTeamEventEdited(tt.args.callbacks...)
			if len(g.onTeamEvent[TeamEventEditedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onTeamEvent[TeamEventEditedAction]), tt.want)
			}
		})
	}
}

func TestHandleTeamEventEdited(t *testing.T) {
	action := TeamEventEditedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.TeamEvent
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
				eventName:  "team",
				event:      &github.TeamEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "team",
				event:      &github.TeamEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "team",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "team",
				event:      &github.TeamEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "team",
				event:      &github.TeamEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "team",
				event:      &github.TeamEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnTeamEventEdited(func(deliveryID string, eventName string, event *github.TeamEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleTeamEventEdited(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleTeamEventEdited() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnTeamEventAddedToRepository(t *testing.T) {
	type args struct {
		callbacks []TeamEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single TeamEventHandleFunc",
			args: args{
				callbacks: []TeamEventHandleFunc{
					func(deliveryID string, eventName string, event *github.TeamEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple TeamEventHandleFunc",
			args: args{
				callbacks: []TeamEventHandleFunc{
					func(deliveryID string, eventName string, event *github.TeamEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.TeamEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnTeamEventAddedToRepository(tt.args.callbacks...)
			if len(g.onTeamEvent[TeamEventAddedToRepositoryAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onTeamEvent[TeamEventAddedToRepositoryAction]))
			}
		})
	}
}

func TestSetOnTeamEventAddedToRepository(t *testing.T) {
	type args struct {
		callbacks []TeamEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single TeamEventHandleFunc",
			args: args{
				[]TeamEventHandleFunc{
					func(deliveryID string, eventName string, event *github.TeamEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple TeamEventHandleFuncs",
			args: args{
				[]TeamEventHandleFunc{
					func(deliveryID string, eventName string, event *github.TeamEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.TeamEvent) error {
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
			g.SetOnTeamEventAddedToRepository(func(deliveryID string, eventName string, event *github.TeamEvent) error {
				return nil
			})
			g.SetOnTeamEventAddedToRepository(tt.args.callbacks...)
			if len(g.onTeamEvent[TeamEventAddedToRepositoryAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onTeamEvent[TeamEventAddedToRepositoryAction]), tt.want)
			}
		})
	}
}

func TestHandleTeamEventAddedToRepository(t *testing.T) {
	action := TeamEventAddedToRepositoryAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.TeamEvent
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
				eventName:  "team",
				event:      &github.TeamEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "team",
				event:      &github.TeamEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "team",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "team",
				event:      &github.TeamEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "team",
				event:      &github.TeamEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "team",
				event:      &github.TeamEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnTeamEventAddedToRepository(func(deliveryID string, eventName string, event *github.TeamEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleTeamEventAddedToRepository(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleTeamEventAddedToRepository() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnTeamEventRemovedFromRepository(t *testing.T) {
	type args struct {
		callbacks []TeamEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single TeamEventHandleFunc",
			args: args{
				callbacks: []TeamEventHandleFunc{
					func(deliveryID string, eventName string, event *github.TeamEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple TeamEventHandleFunc",
			args: args{
				callbacks: []TeamEventHandleFunc{
					func(deliveryID string, eventName string, event *github.TeamEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.TeamEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnTeamEventRemovedFromRepository(tt.args.callbacks...)
			if len(g.onTeamEvent[TeamEventRemovedFromRepositoryAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onTeamEvent[TeamEventRemovedFromRepositoryAction]))
			}
		})
	}
}

func TestSetOnTeamEventRemovedFromRepository(t *testing.T) {
	type args struct {
		callbacks []TeamEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single TeamEventHandleFunc",
			args: args{
				[]TeamEventHandleFunc{
					func(deliveryID string, eventName string, event *github.TeamEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple TeamEventHandleFuncs",
			args: args{
				[]TeamEventHandleFunc{
					func(deliveryID string, eventName string, event *github.TeamEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.TeamEvent) error {
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
			g.SetOnTeamEventRemovedFromRepository(func(deliveryID string, eventName string, event *github.TeamEvent) error {
				return nil
			})
			g.SetOnTeamEventRemovedFromRepository(tt.args.callbacks...)
			if len(g.onTeamEvent[TeamEventRemovedFromRepositoryAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onTeamEvent[TeamEventRemovedFromRepositoryAction]), tt.want)
			}
		})
	}
}

func TestHandleTeamEventRemovedFromRepository(t *testing.T) {
	action := TeamEventRemovedFromRepositoryAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.TeamEvent
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
				eventName:  "team",
				event:      &github.TeamEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "team",
				event:      &github.TeamEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "team",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "team",
				event:      &github.TeamEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "team",
				event:      &github.TeamEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "team",
				event:      &github.TeamEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnTeamEventRemovedFromRepository(func(deliveryID string, eventName string, event *github.TeamEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleTeamEventRemovedFromRepository(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleTeamEventRemovedFromRepository() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTeamEvent(t *testing.T) {
	type fields struct {
		handler *EventHandler
	}
	type args struct {
		deliveryID string
		eventName  string
		event      *github.TeamEvent
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "must trigger TeamEventAny with unknown event action",
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
					onTeamEvent: map[string][]TeamEventHandleFunc{
						TeamEventAnyAction: {
							func(deliveryID string, eventName string, event *github.TeamEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  TeamEvent,

				event: &github.TeamEvent{Action: ptrString("unknown")},
			},
			wantErr: false,
		},

		{
			name: "must trigger TeamEventCreated",
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
					onTeamEvent: map[string][]TeamEventHandleFunc{
						TeamEventAnyAction: {
							func(deliveryID string, eventName string, event *github.TeamEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						TeamEventCreatedAction: {
							func(deliveryID string, eventName string, event *github.TeamEvent) error {
								t.Logf("%s action called", TeamEventCreatedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "team",
				event:      &github.TeamEvent{Action: ptrString(TeamEventCreatedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail TeamEventCreated with empty action",
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
					onTeamEvent: map[string][]TeamEventHandleFunc{
						TeamEventAnyAction: {
							func(deliveryID string, eventName string, event *github.TeamEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						TeamEventCreatedAction: {
							func(deliveryID string, eventName string, event *github.TeamEvent) error {
								t.Logf("%s action called", TeamEventCreatedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "team",
				event:      &github.TeamEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail TeamEventCreated with nil action",
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
					onTeamEvent: map[string][]TeamEventHandleFunc{
						TeamEventAnyAction: {
							func(deliveryID string, eventName string, event *github.TeamEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						TeamEventCreatedAction: {
							func(deliveryID string, eventName string, event *github.TeamEvent) error {
								t.Logf("%s action called", TeamEventCreatedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "team",
				event:      &github.TeamEvent{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger TeamEventDeleted",
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
					onTeamEvent: map[string][]TeamEventHandleFunc{
						TeamEventAnyAction: {
							func(deliveryID string, eventName string, event *github.TeamEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						TeamEventDeletedAction: {
							func(deliveryID string, eventName string, event *github.TeamEvent) error {
								t.Logf("%s action called", TeamEventDeletedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "team",
				event:      &github.TeamEvent{Action: ptrString(TeamEventDeletedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail TeamEventDeleted with empty action",
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
					onTeamEvent: map[string][]TeamEventHandleFunc{
						TeamEventAnyAction: {
							func(deliveryID string, eventName string, event *github.TeamEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						TeamEventDeletedAction: {
							func(deliveryID string, eventName string, event *github.TeamEvent) error {
								t.Logf("%s action called", TeamEventDeletedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "team",
				event:      &github.TeamEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail TeamEventDeleted with nil action",
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
					onTeamEvent: map[string][]TeamEventHandleFunc{
						TeamEventAnyAction: {
							func(deliveryID string, eventName string, event *github.TeamEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						TeamEventDeletedAction: {
							func(deliveryID string, eventName string, event *github.TeamEvent) error {
								t.Logf("%s action called", TeamEventDeletedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "team",
				event:      &github.TeamEvent{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger TeamEventEdited",
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
					onTeamEvent: map[string][]TeamEventHandleFunc{
						TeamEventAnyAction: {
							func(deliveryID string, eventName string, event *github.TeamEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						TeamEventEditedAction: {
							func(deliveryID string, eventName string, event *github.TeamEvent) error {
								t.Logf("%s action called", TeamEventEditedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "team",
				event:      &github.TeamEvent{Action: ptrString(TeamEventEditedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail TeamEventEdited with empty action",
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
					onTeamEvent: map[string][]TeamEventHandleFunc{
						TeamEventAnyAction: {
							func(deliveryID string, eventName string, event *github.TeamEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						TeamEventEditedAction: {
							func(deliveryID string, eventName string, event *github.TeamEvent) error {
								t.Logf("%s action called", TeamEventEditedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "team",
				event:      &github.TeamEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail TeamEventEdited with nil action",
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
					onTeamEvent: map[string][]TeamEventHandleFunc{
						TeamEventAnyAction: {
							func(deliveryID string, eventName string, event *github.TeamEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						TeamEventEditedAction: {
							func(deliveryID string, eventName string, event *github.TeamEvent) error {
								t.Logf("%s action called", TeamEventEditedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "team",
				event:      &github.TeamEvent{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger TeamEventAddedToRepository",
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
					onTeamEvent: map[string][]TeamEventHandleFunc{
						TeamEventAnyAction: {
							func(deliveryID string, eventName string, event *github.TeamEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						TeamEventAddedToRepositoryAction: {
							func(deliveryID string, eventName string, event *github.TeamEvent) error {
								t.Logf("%s action called", TeamEventAddedToRepositoryAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "team",
				event:      &github.TeamEvent{Action: ptrString(TeamEventAddedToRepositoryAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail TeamEventAddedToRepository with empty action",
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
					onTeamEvent: map[string][]TeamEventHandleFunc{
						TeamEventAnyAction: {
							func(deliveryID string, eventName string, event *github.TeamEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						TeamEventAddedToRepositoryAction: {
							func(deliveryID string, eventName string, event *github.TeamEvent) error {
								t.Logf("%s action called", TeamEventAddedToRepositoryAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "team",
				event:      &github.TeamEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail TeamEventAddedToRepository with nil action",
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
					onTeamEvent: map[string][]TeamEventHandleFunc{
						TeamEventAnyAction: {
							func(deliveryID string, eventName string, event *github.TeamEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						TeamEventAddedToRepositoryAction: {
							func(deliveryID string, eventName string, event *github.TeamEvent) error {
								t.Logf("%s action called", TeamEventAddedToRepositoryAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "team",
				event:      &github.TeamEvent{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger TeamEventRemovedFromRepository",
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
					onTeamEvent: map[string][]TeamEventHandleFunc{
						TeamEventAnyAction: {
							func(deliveryID string, eventName string, event *github.TeamEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						TeamEventRemovedFromRepositoryAction: {
							func(deliveryID string, eventName string, event *github.TeamEvent) error {
								t.Logf("%s action called", TeamEventRemovedFromRepositoryAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "team",
				event:      &github.TeamEvent{Action: ptrString(TeamEventRemovedFromRepositoryAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail TeamEventRemovedFromRepository with empty action",
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
					onTeamEvent: map[string][]TeamEventHandleFunc{
						TeamEventAnyAction: {
							func(deliveryID string, eventName string, event *github.TeamEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						TeamEventRemovedFromRepositoryAction: {
							func(deliveryID string, eventName string, event *github.TeamEvent) error {
								t.Logf("%s action called", TeamEventRemovedFromRepositoryAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "team",
				event:      &github.TeamEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail TeamEventRemovedFromRepository with nil action",
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
					onTeamEvent: map[string][]TeamEventHandleFunc{
						TeamEventAnyAction: {
							func(deliveryID string, eventName string, event *github.TeamEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						TeamEventRemovedFromRepositoryAction: {
							func(deliveryID string, eventName string, event *github.TeamEvent) error {
								t.Logf("%s action called", TeamEventRemovedFromRepositoryAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "team",
				event:      &github.TeamEvent{Action: nil},
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
			if err := g.TeamEvent(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("TeamEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
