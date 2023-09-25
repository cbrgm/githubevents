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

func TestOnIssuesEventAny(t *testing.T) {
	type args struct {
		callbacks []IssuesEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single IssuesEventHandleFunc",
			args: args{
				[]IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple IssuesEventHandleFuncs",
			args: args{
				[]IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnIssuesEventAny(tt.args.callbacks...)
			if len(g.onIssuesEvent[IssuesEventAnyAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onIssuesEvent[IssuesEventAnyAction]))
			}
		})
	}
}

func TestSetOnIssuesEventAny(t *testing.T) {
	type args struct {
		callbacks []IssuesEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single IssuesEventHandleFunc",
			args: args{
				[]IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple IssuesEventHandleFuncs",
			args: args{
				[]IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
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
			g.SetOnIssuesEventAny(func(deliveryID string, eventName string, event *github.IssuesEvent) error {
				return nil
			})
			g.SetOnIssuesEventAny(tt.args.callbacks...)
			if len(g.onIssuesEvent[IssuesEventAnyAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onIssuesEvent[IssuesEventAnyAction]), tt.want)
			}
		})
	}
}

func TestHandleIssuesEventAny(t *testing.T) {

	action := "*"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.IssuesEvent
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
				eventName:  "issues",

				event: &github.IssuesEvent{Action: &action},

				fail: false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "issues",

				event: &github.IssuesEvent{Action: &action},

				fail: true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnIssuesEventAny(func(deliveryID string, eventName string, event *github.IssuesEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleIssuesEventAny(tt.args.deliveryID, tt.args.deliveryID, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("TestHandleIssuesEventAny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnIssuesEventOpened(t *testing.T) {
	type args struct {
		callbacks []IssuesEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single IssuesEventHandleFunc",
			args: args{
				callbacks: []IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple IssuesEventHandleFunc",
			args: args{
				callbacks: []IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnIssuesEventOpened(tt.args.callbacks...)
			if len(g.onIssuesEvent[IssuesEventOpenedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onIssuesEvent[IssuesEventOpenedAction]))
			}
		})
	}
}

func TestSetOnIssuesEventOpened(t *testing.T) {
	type args struct {
		callbacks []IssuesEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single IssuesEventHandleFunc",
			args: args{
				[]IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple IssuesEventHandleFuncs",
			args: args{
				[]IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
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
			g.SetOnIssuesEventOpened(func(deliveryID string, eventName string, event *github.IssuesEvent) error {
				return nil
			})
			g.SetOnIssuesEventOpened(tt.args.callbacks...)
			if len(g.onIssuesEvent[IssuesEventOpenedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onIssuesEvent[IssuesEventOpenedAction]), tt.want)
			}
		})
	}
}

func TestHandleIssuesEventOpened(t *testing.T) {
	action := IssuesEventOpenedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.IssuesEvent
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
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnIssuesEventOpened(func(deliveryID string, eventName string, event *github.IssuesEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleIssuesEventOpened(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleIssuesEventOpened() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnIssuesEventEdited(t *testing.T) {
	type args struct {
		callbacks []IssuesEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single IssuesEventHandleFunc",
			args: args{
				callbacks: []IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple IssuesEventHandleFunc",
			args: args{
				callbacks: []IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnIssuesEventEdited(tt.args.callbacks...)
			if len(g.onIssuesEvent[IssuesEventEditedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onIssuesEvent[IssuesEventEditedAction]))
			}
		})
	}
}

func TestSetOnIssuesEventEdited(t *testing.T) {
	type args struct {
		callbacks []IssuesEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single IssuesEventHandleFunc",
			args: args{
				[]IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple IssuesEventHandleFuncs",
			args: args{
				[]IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
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
			g.SetOnIssuesEventEdited(func(deliveryID string, eventName string, event *github.IssuesEvent) error {
				return nil
			})
			g.SetOnIssuesEventEdited(tt.args.callbacks...)
			if len(g.onIssuesEvent[IssuesEventEditedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onIssuesEvent[IssuesEventEditedAction]), tt.want)
			}
		})
	}
}

func TestHandleIssuesEventEdited(t *testing.T) {
	action := IssuesEventEditedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.IssuesEvent
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
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnIssuesEventEdited(func(deliveryID string, eventName string, event *github.IssuesEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleIssuesEventEdited(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleIssuesEventEdited() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnIssuesEventDeleted(t *testing.T) {
	type args struct {
		callbacks []IssuesEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single IssuesEventHandleFunc",
			args: args{
				callbacks: []IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple IssuesEventHandleFunc",
			args: args{
				callbacks: []IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnIssuesEventDeleted(tt.args.callbacks...)
			if len(g.onIssuesEvent[IssuesEventDeletedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onIssuesEvent[IssuesEventDeletedAction]))
			}
		})
	}
}

func TestSetOnIssuesEventDeleted(t *testing.T) {
	type args struct {
		callbacks []IssuesEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single IssuesEventHandleFunc",
			args: args{
				[]IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple IssuesEventHandleFuncs",
			args: args{
				[]IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
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
			g.SetOnIssuesEventDeleted(func(deliveryID string, eventName string, event *github.IssuesEvent) error {
				return nil
			})
			g.SetOnIssuesEventDeleted(tt.args.callbacks...)
			if len(g.onIssuesEvent[IssuesEventDeletedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onIssuesEvent[IssuesEventDeletedAction]), tt.want)
			}
		})
	}
}

func TestHandleIssuesEventDeleted(t *testing.T) {
	action := IssuesEventDeletedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.IssuesEvent
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
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnIssuesEventDeleted(func(deliveryID string, eventName string, event *github.IssuesEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleIssuesEventDeleted(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleIssuesEventDeleted() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnIssuesEventPinned(t *testing.T) {
	type args struct {
		callbacks []IssuesEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single IssuesEventHandleFunc",
			args: args{
				callbacks: []IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple IssuesEventHandleFunc",
			args: args{
				callbacks: []IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnIssuesEventPinned(tt.args.callbacks...)
			if len(g.onIssuesEvent[IssuesEventPinnedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onIssuesEvent[IssuesEventPinnedAction]))
			}
		})
	}
}

func TestSetOnIssuesEventPinned(t *testing.T) {
	type args struct {
		callbacks []IssuesEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single IssuesEventHandleFunc",
			args: args{
				[]IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple IssuesEventHandleFuncs",
			args: args{
				[]IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
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
			g.SetOnIssuesEventPinned(func(deliveryID string, eventName string, event *github.IssuesEvent) error {
				return nil
			})
			g.SetOnIssuesEventPinned(tt.args.callbacks...)
			if len(g.onIssuesEvent[IssuesEventPinnedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onIssuesEvent[IssuesEventPinnedAction]), tt.want)
			}
		})
	}
}

func TestHandleIssuesEventPinned(t *testing.T) {
	action := IssuesEventPinnedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.IssuesEvent
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
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnIssuesEventPinned(func(deliveryID string, eventName string, event *github.IssuesEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleIssuesEventPinned(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleIssuesEventPinned() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnIssuesEventUnpinned(t *testing.T) {
	type args struct {
		callbacks []IssuesEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single IssuesEventHandleFunc",
			args: args{
				callbacks: []IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple IssuesEventHandleFunc",
			args: args{
				callbacks: []IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnIssuesEventUnpinned(tt.args.callbacks...)
			if len(g.onIssuesEvent[IssuesEventUnpinnedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onIssuesEvent[IssuesEventUnpinnedAction]))
			}
		})
	}
}

func TestSetOnIssuesEventUnpinned(t *testing.T) {
	type args struct {
		callbacks []IssuesEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single IssuesEventHandleFunc",
			args: args{
				[]IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple IssuesEventHandleFuncs",
			args: args{
				[]IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
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
			g.SetOnIssuesEventUnpinned(func(deliveryID string, eventName string, event *github.IssuesEvent) error {
				return nil
			})
			g.SetOnIssuesEventUnpinned(tt.args.callbacks...)
			if len(g.onIssuesEvent[IssuesEventUnpinnedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onIssuesEvent[IssuesEventUnpinnedAction]), tt.want)
			}
		})
	}
}

func TestHandleIssuesEventUnpinned(t *testing.T) {
	action := IssuesEventUnpinnedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.IssuesEvent
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
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnIssuesEventUnpinned(func(deliveryID string, eventName string, event *github.IssuesEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleIssuesEventUnpinned(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleIssuesEventUnpinned() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnIssuesEventClosed(t *testing.T) {
	type args struct {
		callbacks []IssuesEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single IssuesEventHandleFunc",
			args: args{
				callbacks: []IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple IssuesEventHandleFunc",
			args: args{
				callbacks: []IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnIssuesEventClosed(tt.args.callbacks...)
			if len(g.onIssuesEvent[IssuesEventClosedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onIssuesEvent[IssuesEventClosedAction]))
			}
		})
	}
}

func TestSetOnIssuesEventClosed(t *testing.T) {
	type args struct {
		callbacks []IssuesEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single IssuesEventHandleFunc",
			args: args{
				[]IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple IssuesEventHandleFuncs",
			args: args{
				[]IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
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
			g.SetOnIssuesEventClosed(func(deliveryID string, eventName string, event *github.IssuesEvent) error {
				return nil
			})
			g.SetOnIssuesEventClosed(tt.args.callbacks...)
			if len(g.onIssuesEvent[IssuesEventClosedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onIssuesEvent[IssuesEventClosedAction]), tt.want)
			}
		})
	}
}

func TestHandleIssuesEventClosed(t *testing.T) {
	action := IssuesEventClosedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.IssuesEvent
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
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnIssuesEventClosed(func(deliveryID string, eventName string, event *github.IssuesEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleIssuesEventClosed(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleIssuesEventClosed() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnIssuesEventReopened(t *testing.T) {
	type args struct {
		callbacks []IssuesEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single IssuesEventHandleFunc",
			args: args{
				callbacks: []IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple IssuesEventHandleFunc",
			args: args{
				callbacks: []IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnIssuesEventReopened(tt.args.callbacks...)
			if len(g.onIssuesEvent[IssuesEventReopenedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onIssuesEvent[IssuesEventReopenedAction]))
			}
		})
	}
}

func TestSetOnIssuesEventReopened(t *testing.T) {
	type args struct {
		callbacks []IssuesEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single IssuesEventHandleFunc",
			args: args{
				[]IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple IssuesEventHandleFuncs",
			args: args{
				[]IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
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
			g.SetOnIssuesEventReopened(func(deliveryID string, eventName string, event *github.IssuesEvent) error {
				return nil
			})
			g.SetOnIssuesEventReopened(tt.args.callbacks...)
			if len(g.onIssuesEvent[IssuesEventReopenedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onIssuesEvent[IssuesEventReopenedAction]), tt.want)
			}
		})
	}
}

func TestHandleIssuesEventReopened(t *testing.T) {
	action := IssuesEventReopenedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.IssuesEvent
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
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnIssuesEventReopened(func(deliveryID string, eventName string, event *github.IssuesEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleIssuesEventReopened(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleIssuesEventReopened() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnIssuesEventAssigned(t *testing.T) {
	type args struct {
		callbacks []IssuesEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single IssuesEventHandleFunc",
			args: args{
				callbacks: []IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple IssuesEventHandleFunc",
			args: args{
				callbacks: []IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnIssuesEventAssigned(tt.args.callbacks...)
			if len(g.onIssuesEvent[IssuesEventAssignedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onIssuesEvent[IssuesEventAssignedAction]))
			}
		})
	}
}

func TestSetOnIssuesEventAssigned(t *testing.T) {
	type args struct {
		callbacks []IssuesEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single IssuesEventHandleFunc",
			args: args{
				[]IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple IssuesEventHandleFuncs",
			args: args{
				[]IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
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
			g.SetOnIssuesEventAssigned(func(deliveryID string, eventName string, event *github.IssuesEvent) error {
				return nil
			})
			g.SetOnIssuesEventAssigned(tt.args.callbacks...)
			if len(g.onIssuesEvent[IssuesEventAssignedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onIssuesEvent[IssuesEventAssignedAction]), tt.want)
			}
		})
	}
}

func TestHandleIssuesEventAssigned(t *testing.T) {
	action := IssuesEventAssignedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.IssuesEvent
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
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnIssuesEventAssigned(func(deliveryID string, eventName string, event *github.IssuesEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleIssuesEventAssigned(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleIssuesEventAssigned() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnIssuesEventUnassigned(t *testing.T) {
	type args struct {
		callbacks []IssuesEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single IssuesEventHandleFunc",
			args: args{
				callbacks: []IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple IssuesEventHandleFunc",
			args: args{
				callbacks: []IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnIssuesEventUnassigned(tt.args.callbacks...)
			if len(g.onIssuesEvent[IssuesEventUnassignedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onIssuesEvent[IssuesEventUnassignedAction]))
			}
		})
	}
}

func TestSetOnIssuesEventUnassigned(t *testing.T) {
	type args struct {
		callbacks []IssuesEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single IssuesEventHandleFunc",
			args: args{
				[]IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple IssuesEventHandleFuncs",
			args: args{
				[]IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
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
			g.SetOnIssuesEventUnassigned(func(deliveryID string, eventName string, event *github.IssuesEvent) error {
				return nil
			})
			g.SetOnIssuesEventUnassigned(tt.args.callbacks...)
			if len(g.onIssuesEvent[IssuesEventUnassignedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onIssuesEvent[IssuesEventUnassignedAction]), tt.want)
			}
		})
	}
}

func TestHandleIssuesEventUnassigned(t *testing.T) {
	action := IssuesEventUnassignedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.IssuesEvent
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
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnIssuesEventUnassigned(func(deliveryID string, eventName string, event *github.IssuesEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleIssuesEventUnassigned(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleIssuesEventUnassigned() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnIssuesEventLabeled(t *testing.T) {
	type args struct {
		callbacks []IssuesEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single IssuesEventHandleFunc",
			args: args{
				callbacks: []IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple IssuesEventHandleFunc",
			args: args{
				callbacks: []IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnIssuesEventLabeled(tt.args.callbacks...)
			if len(g.onIssuesEvent[IssuesEventLabeledAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onIssuesEvent[IssuesEventLabeledAction]))
			}
		})
	}
}

func TestSetOnIssuesEventLabeled(t *testing.T) {
	type args struct {
		callbacks []IssuesEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single IssuesEventHandleFunc",
			args: args{
				[]IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple IssuesEventHandleFuncs",
			args: args{
				[]IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
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
			g.SetOnIssuesEventLabeled(func(deliveryID string, eventName string, event *github.IssuesEvent) error {
				return nil
			})
			g.SetOnIssuesEventLabeled(tt.args.callbacks...)
			if len(g.onIssuesEvent[IssuesEventLabeledAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onIssuesEvent[IssuesEventLabeledAction]), tt.want)
			}
		})
	}
}

func TestHandleIssuesEventLabeled(t *testing.T) {
	action := IssuesEventLabeledAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.IssuesEvent
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
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnIssuesEventLabeled(func(deliveryID string, eventName string, event *github.IssuesEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleIssuesEventLabeled(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleIssuesEventLabeled() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnIssuesEventUnlabeled(t *testing.T) {
	type args struct {
		callbacks []IssuesEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single IssuesEventHandleFunc",
			args: args{
				callbacks: []IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple IssuesEventHandleFunc",
			args: args{
				callbacks: []IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnIssuesEventUnlabeled(tt.args.callbacks...)
			if len(g.onIssuesEvent[IssuesEventUnlabeledAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onIssuesEvent[IssuesEventUnlabeledAction]))
			}
		})
	}
}

func TestSetOnIssuesEventUnlabeled(t *testing.T) {
	type args struct {
		callbacks []IssuesEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single IssuesEventHandleFunc",
			args: args{
				[]IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple IssuesEventHandleFuncs",
			args: args{
				[]IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
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
			g.SetOnIssuesEventUnlabeled(func(deliveryID string, eventName string, event *github.IssuesEvent) error {
				return nil
			})
			g.SetOnIssuesEventUnlabeled(tt.args.callbacks...)
			if len(g.onIssuesEvent[IssuesEventUnlabeledAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onIssuesEvent[IssuesEventUnlabeledAction]), tt.want)
			}
		})
	}
}

func TestHandleIssuesEventUnlabeled(t *testing.T) {
	action := IssuesEventUnlabeledAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.IssuesEvent
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
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnIssuesEventUnlabeled(func(deliveryID string, eventName string, event *github.IssuesEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleIssuesEventUnlabeled(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleIssuesEventUnlabeled() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnIssuesEventLocked(t *testing.T) {
	type args struct {
		callbacks []IssuesEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single IssuesEventHandleFunc",
			args: args{
				callbacks: []IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple IssuesEventHandleFunc",
			args: args{
				callbacks: []IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnIssuesEventLocked(tt.args.callbacks...)
			if len(g.onIssuesEvent[IssuesEventLockedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onIssuesEvent[IssuesEventLockedAction]))
			}
		})
	}
}

func TestSetOnIssuesEventLocked(t *testing.T) {
	type args struct {
		callbacks []IssuesEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single IssuesEventHandleFunc",
			args: args{
				[]IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple IssuesEventHandleFuncs",
			args: args{
				[]IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
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
			g.SetOnIssuesEventLocked(func(deliveryID string, eventName string, event *github.IssuesEvent) error {
				return nil
			})
			g.SetOnIssuesEventLocked(tt.args.callbacks...)
			if len(g.onIssuesEvent[IssuesEventLockedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onIssuesEvent[IssuesEventLockedAction]), tt.want)
			}
		})
	}
}

func TestHandleIssuesEventLocked(t *testing.T) {
	action := IssuesEventLockedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.IssuesEvent
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
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnIssuesEventLocked(func(deliveryID string, eventName string, event *github.IssuesEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleIssuesEventLocked(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleIssuesEventLocked() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnIssuesEventUnlocked(t *testing.T) {
	type args struct {
		callbacks []IssuesEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single IssuesEventHandleFunc",
			args: args{
				callbacks: []IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple IssuesEventHandleFunc",
			args: args{
				callbacks: []IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnIssuesEventUnlocked(tt.args.callbacks...)
			if len(g.onIssuesEvent[IssuesEventUnlockedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onIssuesEvent[IssuesEventUnlockedAction]))
			}
		})
	}
}

func TestSetOnIssuesEventUnlocked(t *testing.T) {
	type args struct {
		callbacks []IssuesEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single IssuesEventHandleFunc",
			args: args{
				[]IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple IssuesEventHandleFuncs",
			args: args{
				[]IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
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
			g.SetOnIssuesEventUnlocked(func(deliveryID string, eventName string, event *github.IssuesEvent) error {
				return nil
			})
			g.SetOnIssuesEventUnlocked(tt.args.callbacks...)
			if len(g.onIssuesEvent[IssuesEventUnlockedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onIssuesEvent[IssuesEventUnlockedAction]), tt.want)
			}
		})
	}
}

func TestHandleIssuesEventUnlocked(t *testing.T) {
	action := IssuesEventUnlockedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.IssuesEvent
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
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnIssuesEventUnlocked(func(deliveryID string, eventName string, event *github.IssuesEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleIssuesEventUnlocked(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleIssuesEventUnlocked() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnIssuesEventTransferred(t *testing.T) {
	type args struct {
		callbacks []IssuesEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single IssuesEventHandleFunc",
			args: args{
				callbacks: []IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple IssuesEventHandleFunc",
			args: args{
				callbacks: []IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnIssuesEventTransferred(tt.args.callbacks...)
			if len(g.onIssuesEvent[IssuesEventTransferredAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onIssuesEvent[IssuesEventTransferredAction]))
			}
		})
	}
}

func TestSetOnIssuesEventTransferred(t *testing.T) {
	type args struct {
		callbacks []IssuesEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single IssuesEventHandleFunc",
			args: args{
				[]IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple IssuesEventHandleFuncs",
			args: args{
				[]IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
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
			g.SetOnIssuesEventTransferred(func(deliveryID string, eventName string, event *github.IssuesEvent) error {
				return nil
			})
			g.SetOnIssuesEventTransferred(tt.args.callbacks...)
			if len(g.onIssuesEvent[IssuesEventTransferredAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onIssuesEvent[IssuesEventTransferredAction]), tt.want)
			}
		})
	}
}

func TestHandleIssuesEventTransferred(t *testing.T) {
	action := IssuesEventTransferredAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.IssuesEvent
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
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnIssuesEventTransferred(func(deliveryID string, eventName string, event *github.IssuesEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleIssuesEventTransferred(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleIssuesEventTransferred() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnIssuesEventMilestoned(t *testing.T) {
	type args struct {
		callbacks []IssuesEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single IssuesEventHandleFunc",
			args: args{
				callbacks: []IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple IssuesEventHandleFunc",
			args: args{
				callbacks: []IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnIssuesEventMilestoned(tt.args.callbacks...)
			if len(g.onIssuesEvent[IssuesEventMilestonedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onIssuesEvent[IssuesEventMilestonedAction]))
			}
		})
	}
}

func TestSetOnIssuesEventMilestoned(t *testing.T) {
	type args struct {
		callbacks []IssuesEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single IssuesEventHandleFunc",
			args: args{
				[]IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple IssuesEventHandleFuncs",
			args: args{
				[]IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
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
			g.SetOnIssuesEventMilestoned(func(deliveryID string, eventName string, event *github.IssuesEvent) error {
				return nil
			})
			g.SetOnIssuesEventMilestoned(tt.args.callbacks...)
			if len(g.onIssuesEvent[IssuesEventMilestonedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onIssuesEvent[IssuesEventMilestonedAction]), tt.want)
			}
		})
	}
}

func TestHandleIssuesEventMilestoned(t *testing.T) {
	action := IssuesEventMilestonedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.IssuesEvent
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
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnIssuesEventMilestoned(func(deliveryID string, eventName string, event *github.IssuesEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleIssuesEventMilestoned(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleIssuesEventMilestoned() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnIssuesEventDeMilestoned(t *testing.T) {
	type args struct {
		callbacks []IssuesEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single IssuesEventHandleFunc",
			args: args{
				callbacks: []IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple IssuesEventHandleFunc",
			args: args{
				callbacks: []IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnIssuesEventDeMilestoned(tt.args.callbacks...)
			if len(g.onIssuesEvent[IssuesEventDeMilestonedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onIssuesEvent[IssuesEventDeMilestonedAction]))
			}
		})
	}
}

func TestSetOnIssuesEventDeMilestoned(t *testing.T) {
	type args struct {
		callbacks []IssuesEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single IssuesEventHandleFunc",
			args: args{
				[]IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple IssuesEventHandleFuncs",
			args: args{
				[]IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
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
			g.SetOnIssuesEventDeMilestoned(func(deliveryID string, eventName string, event *github.IssuesEvent) error {
				return nil
			})
			g.SetOnIssuesEventDeMilestoned(tt.args.callbacks...)
			if len(g.onIssuesEvent[IssuesEventDeMilestonedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onIssuesEvent[IssuesEventDeMilestonedAction]), tt.want)
			}
		})
	}
}

func TestHandleIssuesEventDeMilestoned(t *testing.T) {
	action := IssuesEventDeMilestonedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.IssuesEvent
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
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnIssuesEventDeMilestoned(func(deliveryID string, eventName string, event *github.IssuesEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleIssuesEventDeMilestoned(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleIssuesEventDeMilestoned() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIssuesEvent(t *testing.T) {
	type fields struct {
		handler *EventHandler
	}
	type args struct {
		deliveryID string
		eventName  string
		event      *github.IssuesEvent
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "must trigger IssuesEventAny with unknown event action",
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
					onIssuesEvent: map[string][]IssuesEventHandleFunc{
						IssuesEventAnyAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  IssuesEvent,

				event: &github.IssuesEvent{Action: ptrString("unknown")},
			},
			wantErr: false,
		},

		{
			name: "must trigger IssuesEventOpened",
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
					onIssuesEvent: map[string][]IssuesEventHandleFunc{
						IssuesEventAnyAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						IssuesEventOpenedAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Logf("%s action called", IssuesEventOpenedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: ptrString(IssuesEventOpenedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail IssuesEventOpened with empty action",
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
					onIssuesEvent: map[string][]IssuesEventHandleFunc{
						IssuesEventAnyAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						IssuesEventOpenedAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Logf("%s action called", IssuesEventOpenedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail IssuesEventOpened with nil action",
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
					onIssuesEvent: map[string][]IssuesEventHandleFunc{
						IssuesEventAnyAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						IssuesEventOpenedAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Logf("%s action called", IssuesEventOpenedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger IssuesEventEdited",
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
					onIssuesEvent: map[string][]IssuesEventHandleFunc{
						IssuesEventAnyAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						IssuesEventEditedAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Logf("%s action called", IssuesEventEditedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: ptrString(IssuesEventEditedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail IssuesEventEdited with empty action",
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
					onIssuesEvent: map[string][]IssuesEventHandleFunc{
						IssuesEventAnyAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						IssuesEventEditedAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Logf("%s action called", IssuesEventEditedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail IssuesEventEdited with nil action",
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
					onIssuesEvent: map[string][]IssuesEventHandleFunc{
						IssuesEventAnyAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						IssuesEventEditedAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Logf("%s action called", IssuesEventEditedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger IssuesEventDeleted",
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
					onIssuesEvent: map[string][]IssuesEventHandleFunc{
						IssuesEventAnyAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						IssuesEventDeletedAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Logf("%s action called", IssuesEventDeletedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: ptrString(IssuesEventDeletedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail IssuesEventDeleted with empty action",
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
					onIssuesEvent: map[string][]IssuesEventHandleFunc{
						IssuesEventAnyAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						IssuesEventDeletedAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Logf("%s action called", IssuesEventDeletedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail IssuesEventDeleted with nil action",
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
					onIssuesEvent: map[string][]IssuesEventHandleFunc{
						IssuesEventAnyAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						IssuesEventDeletedAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Logf("%s action called", IssuesEventDeletedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger IssuesEventPinned",
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
					onIssuesEvent: map[string][]IssuesEventHandleFunc{
						IssuesEventAnyAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						IssuesEventPinnedAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Logf("%s action called", IssuesEventPinnedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: ptrString(IssuesEventPinnedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail IssuesEventPinned with empty action",
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
					onIssuesEvent: map[string][]IssuesEventHandleFunc{
						IssuesEventAnyAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						IssuesEventPinnedAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Logf("%s action called", IssuesEventPinnedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail IssuesEventPinned with nil action",
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
					onIssuesEvent: map[string][]IssuesEventHandleFunc{
						IssuesEventAnyAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						IssuesEventPinnedAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Logf("%s action called", IssuesEventPinnedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger IssuesEventUnpinned",
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
					onIssuesEvent: map[string][]IssuesEventHandleFunc{
						IssuesEventAnyAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						IssuesEventUnpinnedAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Logf("%s action called", IssuesEventUnpinnedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: ptrString(IssuesEventUnpinnedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail IssuesEventUnpinned with empty action",
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
					onIssuesEvent: map[string][]IssuesEventHandleFunc{
						IssuesEventAnyAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						IssuesEventUnpinnedAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Logf("%s action called", IssuesEventUnpinnedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail IssuesEventUnpinned with nil action",
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
					onIssuesEvent: map[string][]IssuesEventHandleFunc{
						IssuesEventAnyAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						IssuesEventUnpinnedAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Logf("%s action called", IssuesEventUnpinnedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger IssuesEventClosed",
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
					onIssuesEvent: map[string][]IssuesEventHandleFunc{
						IssuesEventAnyAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						IssuesEventClosedAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Logf("%s action called", IssuesEventClosedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: ptrString(IssuesEventClosedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail IssuesEventClosed with empty action",
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
					onIssuesEvent: map[string][]IssuesEventHandleFunc{
						IssuesEventAnyAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						IssuesEventClosedAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Logf("%s action called", IssuesEventClosedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail IssuesEventClosed with nil action",
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
					onIssuesEvent: map[string][]IssuesEventHandleFunc{
						IssuesEventAnyAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						IssuesEventClosedAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Logf("%s action called", IssuesEventClosedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger IssuesEventReopened",
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
					onIssuesEvent: map[string][]IssuesEventHandleFunc{
						IssuesEventAnyAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						IssuesEventReopenedAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Logf("%s action called", IssuesEventReopenedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: ptrString(IssuesEventReopenedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail IssuesEventReopened with empty action",
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
					onIssuesEvent: map[string][]IssuesEventHandleFunc{
						IssuesEventAnyAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						IssuesEventReopenedAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Logf("%s action called", IssuesEventReopenedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail IssuesEventReopened with nil action",
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
					onIssuesEvent: map[string][]IssuesEventHandleFunc{
						IssuesEventAnyAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						IssuesEventReopenedAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Logf("%s action called", IssuesEventReopenedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger IssuesEventAssigned",
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
					onIssuesEvent: map[string][]IssuesEventHandleFunc{
						IssuesEventAnyAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						IssuesEventAssignedAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Logf("%s action called", IssuesEventAssignedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: ptrString(IssuesEventAssignedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail IssuesEventAssigned with empty action",
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
					onIssuesEvent: map[string][]IssuesEventHandleFunc{
						IssuesEventAnyAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						IssuesEventAssignedAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Logf("%s action called", IssuesEventAssignedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail IssuesEventAssigned with nil action",
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
					onIssuesEvent: map[string][]IssuesEventHandleFunc{
						IssuesEventAnyAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						IssuesEventAssignedAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Logf("%s action called", IssuesEventAssignedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger IssuesEventUnassigned",
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
					onIssuesEvent: map[string][]IssuesEventHandleFunc{
						IssuesEventAnyAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						IssuesEventUnassignedAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Logf("%s action called", IssuesEventUnassignedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: ptrString(IssuesEventUnassignedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail IssuesEventUnassigned with empty action",
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
					onIssuesEvent: map[string][]IssuesEventHandleFunc{
						IssuesEventAnyAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						IssuesEventUnassignedAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Logf("%s action called", IssuesEventUnassignedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail IssuesEventUnassigned with nil action",
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
					onIssuesEvent: map[string][]IssuesEventHandleFunc{
						IssuesEventAnyAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						IssuesEventUnassignedAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Logf("%s action called", IssuesEventUnassignedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger IssuesEventLabeled",
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
					onIssuesEvent: map[string][]IssuesEventHandleFunc{
						IssuesEventAnyAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						IssuesEventLabeledAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Logf("%s action called", IssuesEventLabeledAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: ptrString(IssuesEventLabeledAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail IssuesEventLabeled with empty action",
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
					onIssuesEvent: map[string][]IssuesEventHandleFunc{
						IssuesEventAnyAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						IssuesEventLabeledAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Logf("%s action called", IssuesEventLabeledAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail IssuesEventLabeled with nil action",
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
					onIssuesEvent: map[string][]IssuesEventHandleFunc{
						IssuesEventAnyAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						IssuesEventLabeledAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Logf("%s action called", IssuesEventLabeledAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger IssuesEventUnlabeled",
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
					onIssuesEvent: map[string][]IssuesEventHandleFunc{
						IssuesEventAnyAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						IssuesEventUnlabeledAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Logf("%s action called", IssuesEventUnlabeledAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: ptrString(IssuesEventUnlabeledAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail IssuesEventUnlabeled with empty action",
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
					onIssuesEvent: map[string][]IssuesEventHandleFunc{
						IssuesEventAnyAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						IssuesEventUnlabeledAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Logf("%s action called", IssuesEventUnlabeledAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail IssuesEventUnlabeled with nil action",
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
					onIssuesEvent: map[string][]IssuesEventHandleFunc{
						IssuesEventAnyAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						IssuesEventUnlabeledAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Logf("%s action called", IssuesEventUnlabeledAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger IssuesEventLocked",
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
					onIssuesEvent: map[string][]IssuesEventHandleFunc{
						IssuesEventAnyAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						IssuesEventLockedAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Logf("%s action called", IssuesEventLockedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: ptrString(IssuesEventLockedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail IssuesEventLocked with empty action",
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
					onIssuesEvent: map[string][]IssuesEventHandleFunc{
						IssuesEventAnyAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						IssuesEventLockedAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Logf("%s action called", IssuesEventLockedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail IssuesEventLocked with nil action",
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
					onIssuesEvent: map[string][]IssuesEventHandleFunc{
						IssuesEventAnyAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						IssuesEventLockedAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Logf("%s action called", IssuesEventLockedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger IssuesEventUnlocked",
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
					onIssuesEvent: map[string][]IssuesEventHandleFunc{
						IssuesEventAnyAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						IssuesEventUnlockedAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Logf("%s action called", IssuesEventUnlockedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: ptrString(IssuesEventUnlockedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail IssuesEventUnlocked with empty action",
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
					onIssuesEvent: map[string][]IssuesEventHandleFunc{
						IssuesEventAnyAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						IssuesEventUnlockedAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Logf("%s action called", IssuesEventUnlockedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail IssuesEventUnlocked with nil action",
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
					onIssuesEvent: map[string][]IssuesEventHandleFunc{
						IssuesEventAnyAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						IssuesEventUnlockedAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Logf("%s action called", IssuesEventUnlockedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger IssuesEventTransferred",
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
					onIssuesEvent: map[string][]IssuesEventHandleFunc{
						IssuesEventAnyAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						IssuesEventTransferredAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Logf("%s action called", IssuesEventTransferredAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: ptrString(IssuesEventTransferredAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail IssuesEventTransferred with empty action",
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
					onIssuesEvent: map[string][]IssuesEventHandleFunc{
						IssuesEventAnyAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						IssuesEventTransferredAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Logf("%s action called", IssuesEventTransferredAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail IssuesEventTransferred with nil action",
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
					onIssuesEvent: map[string][]IssuesEventHandleFunc{
						IssuesEventAnyAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						IssuesEventTransferredAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Logf("%s action called", IssuesEventTransferredAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger IssuesEventMilestoned",
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
					onIssuesEvent: map[string][]IssuesEventHandleFunc{
						IssuesEventAnyAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						IssuesEventMilestonedAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Logf("%s action called", IssuesEventMilestonedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: ptrString(IssuesEventMilestonedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail IssuesEventMilestoned with empty action",
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
					onIssuesEvent: map[string][]IssuesEventHandleFunc{
						IssuesEventAnyAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						IssuesEventMilestonedAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Logf("%s action called", IssuesEventMilestonedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail IssuesEventMilestoned with nil action",
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
					onIssuesEvent: map[string][]IssuesEventHandleFunc{
						IssuesEventAnyAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						IssuesEventMilestonedAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Logf("%s action called", IssuesEventMilestonedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger IssuesEventDeMilestoned",
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
					onIssuesEvent: map[string][]IssuesEventHandleFunc{
						IssuesEventAnyAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						IssuesEventDeMilestonedAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Logf("%s action called", IssuesEventDeMilestonedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: ptrString(IssuesEventDeMilestonedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail IssuesEventDeMilestoned with empty action",
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
					onIssuesEvent: map[string][]IssuesEventHandleFunc{
						IssuesEventAnyAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						IssuesEventDeMilestonedAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Logf("%s action called", IssuesEventDeMilestonedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail IssuesEventDeMilestoned with nil action",
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
					onIssuesEvent: map[string][]IssuesEventHandleFunc{
						IssuesEventAnyAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						IssuesEventDeMilestonedAction: {
							func(deliveryID string, eventName string, event *github.IssuesEvent) error {
								t.Logf("%s action called", IssuesEventDeMilestonedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: nil},
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
			if err := g.IssuesEvent(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("IssuesEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
