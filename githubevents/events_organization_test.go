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

func TestOnOrganizationEventAny(t *testing.T) {
	type args struct {
		callbacks []OrganizationEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single OrganizationEventHandleFunc",
			args: args{
				[]OrganizationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple OrganizationEventHandleFuncs",
			args: args{
				[]OrganizationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnOrganizationEventAny(tt.args.callbacks...)
			if len(g.onOrganizationEvent[OrganizationEventAnyAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onOrganizationEvent[OrganizationEventAnyAction]))
			}
		})
	}
}

func TestSetOnOrganizationEventAny(t *testing.T) {
	type args struct {
		callbacks []OrganizationEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single OrganizationEventHandleFunc",
			args: args{
				[]OrganizationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple OrganizationEventHandleFuncs",
			args: args{
				[]OrganizationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
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
			g.SetOnOrganizationEventAny(func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
				return nil
			})
			g.SetOnOrganizationEventAny(tt.args.callbacks...)
			if len(g.onOrganizationEvent[OrganizationEventAnyAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onOrganizationEvent[OrganizationEventAnyAction]), tt.want)
			}
		})
	}
}

func TestHandleOrganizationEventAny(t *testing.T) {

	action := "*"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.OrganizationEvent
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
				eventName:  "organization",

				event: &github.OrganizationEvent{Action: &action},

				fail: false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "organization",

				event: &github.OrganizationEvent{Action: &action},

				fail: true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "organization",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnOrganizationEventAny(func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleOrganizationEventAny(tt.args.deliveryID, tt.args.deliveryID, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("TestHandleOrganizationEventAny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnOrganizationEventDeleted(t *testing.T) {
	type args struct {
		callbacks []OrganizationEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single OrganizationEventHandleFunc",
			args: args{
				callbacks: []OrganizationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple OrganizationEventHandleFunc",
			args: args{
				callbacks: []OrganizationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnOrganizationEventDeleted(tt.args.callbacks...)
			if len(g.onOrganizationEvent[OrganizationEventDeletedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onOrganizationEvent[OrganizationEventDeletedAction]))
			}
		})
	}
}

func TestSetOnOrganizationEventDeleted(t *testing.T) {
	type args struct {
		callbacks []OrganizationEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single OrganizationEventHandleFunc",
			args: args{
				[]OrganizationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple OrganizationEventHandleFuncs",
			args: args{
				[]OrganizationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
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
			g.SetOnOrganizationEventDeleted(func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
				return nil
			})
			g.SetOnOrganizationEventDeleted(tt.args.callbacks...)
			if len(g.onOrganizationEvent[OrganizationEventDeletedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onOrganizationEvent[OrganizationEventDeletedAction]), tt.want)
			}
		})
	}
}

func TestHandleOrganizationEventDeleted(t *testing.T) {
	action := OrganizationEventDeletedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.OrganizationEvent
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
				eventName:  "organization",
				event:      &github.OrganizationEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "organization",
				event:      &github.OrganizationEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "organization",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "organization",
				event:      &github.OrganizationEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "organization",
				event:      &github.OrganizationEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "organization",
				event:      &github.OrganizationEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnOrganizationEventDeleted(func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleOrganizationEventDeleted(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleOrganizationEventDeleted() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnOrganizationEventRenamed(t *testing.T) {
	type args struct {
		callbacks []OrganizationEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single OrganizationEventHandleFunc",
			args: args{
				callbacks: []OrganizationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple OrganizationEventHandleFunc",
			args: args{
				callbacks: []OrganizationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnOrganizationEventRenamed(tt.args.callbacks...)
			if len(g.onOrganizationEvent[OrganizationEventRenamedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onOrganizationEvent[OrganizationEventRenamedAction]))
			}
		})
	}
}

func TestSetOnOrganizationEventRenamed(t *testing.T) {
	type args struct {
		callbacks []OrganizationEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single OrganizationEventHandleFunc",
			args: args{
				[]OrganizationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple OrganizationEventHandleFuncs",
			args: args{
				[]OrganizationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
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
			g.SetOnOrganizationEventRenamed(func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
				return nil
			})
			g.SetOnOrganizationEventRenamed(tt.args.callbacks...)
			if len(g.onOrganizationEvent[OrganizationEventRenamedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onOrganizationEvent[OrganizationEventRenamedAction]), tt.want)
			}
		})
	}
}

func TestHandleOrganizationEventRenamed(t *testing.T) {
	action := OrganizationEventRenamedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.OrganizationEvent
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
				eventName:  "organization",
				event:      &github.OrganizationEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "organization",
				event:      &github.OrganizationEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "organization",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "organization",
				event:      &github.OrganizationEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "organization",
				event:      &github.OrganizationEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "organization",
				event:      &github.OrganizationEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnOrganizationEventRenamed(func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleOrganizationEventRenamed(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleOrganizationEventRenamed() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnOrganizationEventMemberAdded(t *testing.T) {
	type args struct {
		callbacks []OrganizationEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single OrganizationEventHandleFunc",
			args: args{
				callbacks: []OrganizationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple OrganizationEventHandleFunc",
			args: args{
				callbacks: []OrganizationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnOrganizationEventMemberAdded(tt.args.callbacks...)
			if len(g.onOrganizationEvent[OrganizationEventMemberAddedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onOrganizationEvent[OrganizationEventMemberAddedAction]))
			}
		})
	}
}

func TestSetOnOrganizationEventMemberAdded(t *testing.T) {
	type args struct {
		callbacks []OrganizationEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single OrganizationEventHandleFunc",
			args: args{
				[]OrganizationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple OrganizationEventHandleFuncs",
			args: args{
				[]OrganizationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
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
			g.SetOnOrganizationEventMemberAdded(func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
				return nil
			})
			g.SetOnOrganizationEventMemberAdded(tt.args.callbacks...)
			if len(g.onOrganizationEvent[OrganizationEventMemberAddedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onOrganizationEvent[OrganizationEventMemberAddedAction]), tt.want)
			}
		})
	}
}

func TestHandleOrganizationEventMemberAdded(t *testing.T) {
	action := OrganizationEventMemberAddedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.OrganizationEvent
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
				eventName:  "organization",
				event:      &github.OrganizationEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "organization",
				event:      &github.OrganizationEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "organization",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "organization",
				event:      &github.OrganizationEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "organization",
				event:      &github.OrganizationEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "organization",
				event:      &github.OrganizationEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnOrganizationEventMemberAdded(func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleOrganizationEventMemberAdded(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleOrganizationEventMemberAdded() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnOrganizationEventMemberRemoved(t *testing.T) {
	type args struct {
		callbacks []OrganizationEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single OrganizationEventHandleFunc",
			args: args{
				callbacks: []OrganizationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple OrganizationEventHandleFunc",
			args: args{
				callbacks: []OrganizationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnOrganizationEventMemberRemoved(tt.args.callbacks...)
			if len(g.onOrganizationEvent[OrganizationEventMemberRemovedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onOrganizationEvent[OrganizationEventMemberRemovedAction]))
			}
		})
	}
}

func TestSetOnOrganizationEventMemberRemoved(t *testing.T) {
	type args struct {
		callbacks []OrganizationEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single OrganizationEventHandleFunc",
			args: args{
				[]OrganizationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple OrganizationEventHandleFuncs",
			args: args{
				[]OrganizationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
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
			g.SetOnOrganizationEventMemberRemoved(func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
				return nil
			})
			g.SetOnOrganizationEventMemberRemoved(tt.args.callbacks...)
			if len(g.onOrganizationEvent[OrganizationEventMemberRemovedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onOrganizationEvent[OrganizationEventMemberRemovedAction]), tt.want)
			}
		})
	}
}

func TestHandleOrganizationEventMemberRemoved(t *testing.T) {
	action := OrganizationEventMemberRemovedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.OrganizationEvent
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
				eventName:  "organization",
				event:      &github.OrganizationEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "organization",
				event:      &github.OrganizationEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "organization",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "organization",
				event:      &github.OrganizationEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "organization",
				event:      &github.OrganizationEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "organization",
				event:      &github.OrganizationEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnOrganizationEventMemberRemoved(func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleOrganizationEventMemberRemoved(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleOrganizationEventMemberRemoved() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnOrganizationEventMemberInvited(t *testing.T) {
	type args struct {
		callbacks []OrganizationEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single OrganizationEventHandleFunc",
			args: args{
				callbacks: []OrganizationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple OrganizationEventHandleFunc",
			args: args{
				callbacks: []OrganizationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnOrganizationEventMemberInvited(tt.args.callbacks...)
			if len(g.onOrganizationEvent[OrganizationEventMemberInvitedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onOrganizationEvent[OrganizationEventMemberInvitedAction]))
			}
		})
	}
}

func TestSetOnOrganizationEventMemberInvited(t *testing.T) {
	type args struct {
		callbacks []OrganizationEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single OrganizationEventHandleFunc",
			args: args{
				[]OrganizationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple OrganizationEventHandleFuncs",
			args: args{
				[]OrganizationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
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
			g.SetOnOrganizationEventMemberInvited(func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
				return nil
			})
			g.SetOnOrganizationEventMemberInvited(tt.args.callbacks...)
			if len(g.onOrganizationEvent[OrganizationEventMemberInvitedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onOrganizationEvent[OrganizationEventMemberInvitedAction]), tt.want)
			}
		})
	}
}

func TestHandleOrganizationEventMemberInvited(t *testing.T) {
	action := OrganizationEventMemberInvitedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.OrganizationEvent
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
				eventName:  "organization",
				event:      &github.OrganizationEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "organization",
				event:      &github.OrganizationEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "organization",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "organization",
				event:      &github.OrganizationEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "organization",
				event:      &github.OrganizationEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "organization",
				event:      &github.OrganizationEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnOrganizationEventMemberInvited(func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleOrganizationEventMemberInvited(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleOrganizationEventMemberInvited() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOrganizationEvent(t *testing.T) {
	type fields struct {
		handler *EventHandler
	}
	type args struct {
		deliveryID string
		eventName  string
		event      *github.OrganizationEvent
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "must trigger OrganizationEventAny with unknown event action",
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
					onOrganizationEvent: map[string][]OrganizationEventHandleFunc{
						OrganizationEventAnyAction: {
							func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  OrganizationEvent,

				event: &github.OrganizationEvent{Action: ptrString("unknown")},
			},
			wantErr: false,
		},

		{
			name: "must trigger OrganizationEventDeleted",
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
					onOrganizationEvent: map[string][]OrganizationEventHandleFunc{
						OrganizationEventAnyAction: {
							func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						OrganizationEventDeletedAction: {
							func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
								t.Logf("%s action called", OrganizationEventDeletedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "organization",
				event:      &github.OrganizationEvent{Action: ptrString(OrganizationEventDeletedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail OrganizationEventDeleted with empty action",
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
					onOrganizationEvent: map[string][]OrganizationEventHandleFunc{
						OrganizationEventAnyAction: {
							func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						OrganizationEventDeletedAction: {
							func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
								t.Logf("%s action called", OrganizationEventDeletedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "organization",
				event:      &github.OrganizationEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail OrganizationEventDeleted with nil action",
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
					onOrganizationEvent: map[string][]OrganizationEventHandleFunc{
						OrganizationEventAnyAction: {
							func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						OrganizationEventDeletedAction: {
							func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
								t.Logf("%s action called", OrganizationEventDeletedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "organization",
				event:      &github.OrganizationEvent{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger OrganizationEventRenamed",
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
					onOrganizationEvent: map[string][]OrganizationEventHandleFunc{
						OrganizationEventAnyAction: {
							func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						OrganizationEventRenamedAction: {
							func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
								t.Logf("%s action called", OrganizationEventRenamedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "organization",
				event:      &github.OrganizationEvent{Action: ptrString(OrganizationEventRenamedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail OrganizationEventRenamed with empty action",
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
					onOrganizationEvent: map[string][]OrganizationEventHandleFunc{
						OrganizationEventAnyAction: {
							func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						OrganizationEventRenamedAction: {
							func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
								t.Logf("%s action called", OrganizationEventRenamedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "organization",
				event:      &github.OrganizationEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail OrganizationEventRenamed with nil action",
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
					onOrganizationEvent: map[string][]OrganizationEventHandleFunc{
						OrganizationEventAnyAction: {
							func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						OrganizationEventRenamedAction: {
							func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
								t.Logf("%s action called", OrganizationEventRenamedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "organization",
				event:      &github.OrganizationEvent{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger OrganizationEventMemberAdded",
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
					onOrganizationEvent: map[string][]OrganizationEventHandleFunc{
						OrganizationEventAnyAction: {
							func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						OrganizationEventMemberAddedAction: {
							func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
								t.Logf("%s action called", OrganizationEventMemberAddedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "organization",
				event:      &github.OrganizationEvent{Action: ptrString(OrganizationEventMemberAddedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail OrganizationEventMemberAdded with empty action",
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
					onOrganizationEvent: map[string][]OrganizationEventHandleFunc{
						OrganizationEventAnyAction: {
							func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						OrganizationEventMemberAddedAction: {
							func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
								t.Logf("%s action called", OrganizationEventMemberAddedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "organization",
				event:      &github.OrganizationEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail OrganizationEventMemberAdded with nil action",
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
					onOrganizationEvent: map[string][]OrganizationEventHandleFunc{
						OrganizationEventAnyAction: {
							func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						OrganizationEventMemberAddedAction: {
							func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
								t.Logf("%s action called", OrganizationEventMemberAddedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "organization",
				event:      &github.OrganizationEvent{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger OrganizationEventMemberRemoved",
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
					onOrganizationEvent: map[string][]OrganizationEventHandleFunc{
						OrganizationEventAnyAction: {
							func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						OrganizationEventMemberRemovedAction: {
							func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
								t.Logf("%s action called", OrganizationEventMemberRemovedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "organization",
				event:      &github.OrganizationEvent{Action: ptrString(OrganizationEventMemberRemovedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail OrganizationEventMemberRemoved with empty action",
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
					onOrganizationEvent: map[string][]OrganizationEventHandleFunc{
						OrganizationEventAnyAction: {
							func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						OrganizationEventMemberRemovedAction: {
							func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
								t.Logf("%s action called", OrganizationEventMemberRemovedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "organization",
				event:      &github.OrganizationEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail OrganizationEventMemberRemoved with nil action",
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
					onOrganizationEvent: map[string][]OrganizationEventHandleFunc{
						OrganizationEventAnyAction: {
							func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						OrganizationEventMemberRemovedAction: {
							func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
								t.Logf("%s action called", OrganizationEventMemberRemovedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "organization",
				event:      &github.OrganizationEvent{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger OrganizationEventMemberInvited",
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
					onOrganizationEvent: map[string][]OrganizationEventHandleFunc{
						OrganizationEventAnyAction: {
							func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						OrganizationEventMemberInvitedAction: {
							func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
								t.Logf("%s action called", OrganizationEventMemberInvitedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "organization",
				event:      &github.OrganizationEvent{Action: ptrString(OrganizationEventMemberInvitedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail OrganizationEventMemberInvited with empty action",
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
					onOrganizationEvent: map[string][]OrganizationEventHandleFunc{
						OrganizationEventAnyAction: {
							func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						OrganizationEventMemberInvitedAction: {
							func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
								t.Logf("%s action called", OrganizationEventMemberInvitedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "organization",
				event:      &github.OrganizationEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail OrganizationEventMemberInvited with nil action",
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
					onOrganizationEvent: map[string][]OrganizationEventHandleFunc{
						OrganizationEventAnyAction: {
							func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						OrganizationEventMemberInvitedAction: {
							func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
								t.Logf("%s action called", OrganizationEventMemberInvitedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "organization",
				event:      &github.OrganizationEvent{Action: nil},
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
			if err := g.OrganizationEvent(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("OrganizationEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
