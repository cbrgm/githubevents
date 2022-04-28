// Copyright 2022 The GithubEvents Authors. All rights reserved.
// Use of this source code is governed by the MIT License
// that can be found in the LICENSE file.

package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"errors"
	"github.com/google/go-github/v43/github"
	"testing"
)

func TestOnPullRequestEventAny(t *testing.T) {
	type args struct {
		callbacks []PullRequestEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single PullRequestEventHandleFunc",
			args: args{
				[]PullRequestEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple PullRequestEventHandleFuncs",
			args: args{
				[]PullRequestEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnPullRequestEventAny(tt.args.callbacks...)
			if len(g.onPullRequestEvent[PullRequestEventAnyAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onPullRequestEvent[PullRequestEventAnyAction]))
			}
		})
	}
}

func TestSetOnPullRequestEventAny(t *testing.T) {
	type args struct {
		callbacks []PullRequestEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single PullRequestEventHandleFunc",
			args: args{
				[]PullRequestEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple PullRequestEventHandleFuncs",
			args: args{
				[]PullRequestEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
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
			g.SetOnPullRequestEventAny(func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
				return nil
			})
			g.SetOnPullRequestEventAny(tt.args.callbacks...)
			if len(g.onPullRequestEvent[PullRequestEventAnyAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onPullRequestEvent[PullRequestEventAnyAction]), tt.want)
			}
		})
	}
}

func TestHandlePullRequestEventAny(t *testing.T) {

	action := "*"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.PullRequestEvent
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
				eventName:  "pull_request",

				event: &github.PullRequestEvent{Action: &action},

				fail: false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",

				event: &github.PullRequestEvent{Action: &action},

				fail: true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnPullRequestEventAny(func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handlePullRequestEventAny(tt.args.deliveryID, tt.args.deliveryID, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("TestHandlePullRequestEventAny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnPullRequestEventAssigned(t *testing.T) {
	type args struct {
		callbacks []PullRequestEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single PullRequestEventHandleFunc",
			args: args{
				callbacks: []PullRequestEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple PullRequestEventHandleFunc",
			args: args{
				callbacks: []PullRequestEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnPullRequestEventAssigned(tt.args.callbacks...)
			if len(g.onPullRequestEvent[PullRequestEventAssignedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onPullRequestEvent[PullRequestEventAssignedAction]))
			}
		})
	}
}

func TestSetOnPullRequestEventAssigned(t *testing.T) {
	type args struct {
		callbacks []PullRequestEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single PullRequestEventHandleFunc",
			args: args{
				[]PullRequestEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple PullRequestEventHandleFuncs",
			args: args{
				[]PullRequestEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
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
			g.SetOnPullRequestEventAssigned(func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
				return nil
			})
			g.SetOnPullRequestEventAssigned(tt.args.callbacks...)
			if len(g.onPullRequestEvent[PullRequestEventAssignedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onPullRequestEvent[PullRequestEventAssignedAction]), tt.want)
			}
		})
	}
}

func TestHandlePullRequestEventAssigned(t *testing.T) {
	action := PullRequestEventAssignedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.PullRequestEvent
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
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnPullRequestEventAssigned(func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handlePullRequestEventAssigned(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handlePullRequestEventAssigned() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnPullRequestEventAutoMergeDisabled(t *testing.T) {
	type args struct {
		callbacks []PullRequestEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single PullRequestEventHandleFunc",
			args: args{
				callbacks: []PullRequestEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple PullRequestEventHandleFunc",
			args: args{
				callbacks: []PullRequestEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnPullRequestEventAutoMergeDisabled(tt.args.callbacks...)
			if len(g.onPullRequestEvent[PullRequestEventAutoMergeDisabledAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onPullRequestEvent[PullRequestEventAutoMergeDisabledAction]))
			}
		})
	}
}

func TestSetOnPullRequestEventAutoMergeDisabled(t *testing.T) {
	type args struct {
		callbacks []PullRequestEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single PullRequestEventHandleFunc",
			args: args{
				[]PullRequestEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple PullRequestEventHandleFuncs",
			args: args{
				[]PullRequestEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
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
			g.SetOnPullRequestEventAutoMergeDisabled(func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
				return nil
			})
			g.SetOnPullRequestEventAutoMergeDisabled(tt.args.callbacks...)
			if len(g.onPullRequestEvent[PullRequestEventAutoMergeDisabledAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onPullRequestEvent[PullRequestEventAutoMergeDisabledAction]), tt.want)
			}
		})
	}
}

func TestHandlePullRequestEventAutoMergeDisabled(t *testing.T) {
	action := PullRequestEventAutoMergeDisabledAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.PullRequestEvent
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
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnPullRequestEventAutoMergeDisabled(func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handlePullRequestEventAutoMergeDisabled(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handlePullRequestEventAutoMergeDisabled() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnPullRequestEventAutoMergeEnabled(t *testing.T) {
	type args struct {
		callbacks []PullRequestEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single PullRequestEventHandleFunc",
			args: args{
				callbacks: []PullRequestEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple PullRequestEventHandleFunc",
			args: args{
				callbacks: []PullRequestEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnPullRequestEventAutoMergeEnabled(tt.args.callbacks...)
			if len(g.onPullRequestEvent[PullRequestEventAutoMergeEnabledAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onPullRequestEvent[PullRequestEventAutoMergeEnabledAction]))
			}
		})
	}
}

func TestSetOnPullRequestEventAutoMergeEnabled(t *testing.T) {
	type args struct {
		callbacks []PullRequestEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single PullRequestEventHandleFunc",
			args: args{
				[]PullRequestEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple PullRequestEventHandleFuncs",
			args: args{
				[]PullRequestEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
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
			g.SetOnPullRequestEventAutoMergeEnabled(func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
				return nil
			})
			g.SetOnPullRequestEventAutoMergeEnabled(tt.args.callbacks...)
			if len(g.onPullRequestEvent[PullRequestEventAutoMergeEnabledAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onPullRequestEvent[PullRequestEventAutoMergeEnabledAction]), tt.want)
			}
		})
	}
}

func TestHandlePullRequestEventAutoMergeEnabled(t *testing.T) {
	action := PullRequestEventAutoMergeEnabledAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.PullRequestEvent
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
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnPullRequestEventAutoMergeEnabled(func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handlePullRequestEventAutoMergeEnabled(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handlePullRequestEventAutoMergeEnabled() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnPullRequestEventClosed(t *testing.T) {
	type args struct {
		callbacks []PullRequestEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single PullRequestEventHandleFunc",
			args: args{
				callbacks: []PullRequestEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple PullRequestEventHandleFunc",
			args: args{
				callbacks: []PullRequestEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnPullRequestEventClosed(tt.args.callbacks...)
			if len(g.onPullRequestEvent[PullRequestEventClosedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onPullRequestEvent[PullRequestEventClosedAction]))
			}
		})
	}
}

func TestSetOnPullRequestEventClosed(t *testing.T) {
	type args struct {
		callbacks []PullRequestEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single PullRequestEventHandleFunc",
			args: args{
				[]PullRequestEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple PullRequestEventHandleFuncs",
			args: args{
				[]PullRequestEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
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
			g.SetOnPullRequestEventClosed(func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
				return nil
			})
			g.SetOnPullRequestEventClosed(tt.args.callbacks...)
			if len(g.onPullRequestEvent[PullRequestEventClosedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onPullRequestEvent[PullRequestEventClosedAction]), tt.want)
			}
		})
	}
}

func TestHandlePullRequestEventClosed(t *testing.T) {
	action := PullRequestEventClosedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.PullRequestEvent
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
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnPullRequestEventClosed(func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handlePullRequestEventClosed(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handlePullRequestEventClosed() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnPullRequestEventConvertedToDraft(t *testing.T) {
	type args struct {
		callbacks []PullRequestEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single PullRequestEventHandleFunc",
			args: args{
				callbacks: []PullRequestEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple PullRequestEventHandleFunc",
			args: args{
				callbacks: []PullRequestEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnPullRequestEventConvertedToDraft(tt.args.callbacks...)
			if len(g.onPullRequestEvent[PullRequestEventConvertedToDraftAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onPullRequestEvent[PullRequestEventConvertedToDraftAction]))
			}
		})
	}
}

func TestSetOnPullRequestEventConvertedToDraft(t *testing.T) {
	type args struct {
		callbacks []PullRequestEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single PullRequestEventHandleFunc",
			args: args{
				[]PullRequestEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple PullRequestEventHandleFuncs",
			args: args{
				[]PullRequestEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
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
			g.SetOnPullRequestEventConvertedToDraft(func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
				return nil
			})
			g.SetOnPullRequestEventConvertedToDraft(tt.args.callbacks...)
			if len(g.onPullRequestEvent[PullRequestEventConvertedToDraftAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onPullRequestEvent[PullRequestEventConvertedToDraftAction]), tt.want)
			}
		})
	}
}

func TestHandlePullRequestEventConvertedToDraft(t *testing.T) {
	action := PullRequestEventConvertedToDraftAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.PullRequestEvent
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
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnPullRequestEventConvertedToDraft(func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handlePullRequestEventConvertedToDraft(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handlePullRequestEventConvertedToDraft() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnPullRequestEventEdited(t *testing.T) {
	type args struct {
		callbacks []PullRequestEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single PullRequestEventHandleFunc",
			args: args{
				callbacks: []PullRequestEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple PullRequestEventHandleFunc",
			args: args{
				callbacks: []PullRequestEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnPullRequestEventEdited(tt.args.callbacks...)
			if len(g.onPullRequestEvent[PullRequestEventEditedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onPullRequestEvent[PullRequestEventEditedAction]))
			}
		})
	}
}

func TestSetOnPullRequestEventEdited(t *testing.T) {
	type args struct {
		callbacks []PullRequestEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single PullRequestEventHandleFunc",
			args: args{
				[]PullRequestEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple PullRequestEventHandleFuncs",
			args: args{
				[]PullRequestEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
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
			g.SetOnPullRequestEventEdited(func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
				return nil
			})
			g.SetOnPullRequestEventEdited(tt.args.callbacks...)
			if len(g.onPullRequestEvent[PullRequestEventEditedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onPullRequestEvent[PullRequestEventEditedAction]), tt.want)
			}
		})
	}
}

func TestHandlePullRequestEventEdited(t *testing.T) {
	action := PullRequestEventEditedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.PullRequestEvent
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
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnPullRequestEventEdited(func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handlePullRequestEventEdited(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handlePullRequestEventEdited() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnPullRequestEventLabeled(t *testing.T) {
	type args struct {
		callbacks []PullRequestEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single PullRequestEventHandleFunc",
			args: args{
				callbacks: []PullRequestEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple PullRequestEventHandleFunc",
			args: args{
				callbacks: []PullRequestEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnPullRequestEventLabeled(tt.args.callbacks...)
			if len(g.onPullRequestEvent[PullRequestEventLabeledAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onPullRequestEvent[PullRequestEventLabeledAction]))
			}
		})
	}
}

func TestSetOnPullRequestEventLabeled(t *testing.T) {
	type args struct {
		callbacks []PullRequestEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single PullRequestEventHandleFunc",
			args: args{
				[]PullRequestEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple PullRequestEventHandleFuncs",
			args: args{
				[]PullRequestEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
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
			g.SetOnPullRequestEventLabeled(func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
				return nil
			})
			g.SetOnPullRequestEventLabeled(tt.args.callbacks...)
			if len(g.onPullRequestEvent[PullRequestEventLabeledAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onPullRequestEvent[PullRequestEventLabeledAction]), tt.want)
			}
		})
	}
}

func TestHandlePullRequestEventLabeled(t *testing.T) {
	action := PullRequestEventLabeledAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.PullRequestEvent
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
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnPullRequestEventLabeled(func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handlePullRequestEventLabeled(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handlePullRequestEventLabeled() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnPullRequestEventLocked(t *testing.T) {
	type args struct {
		callbacks []PullRequestEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single PullRequestEventHandleFunc",
			args: args{
				callbacks: []PullRequestEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple PullRequestEventHandleFunc",
			args: args{
				callbacks: []PullRequestEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnPullRequestEventLocked(tt.args.callbacks...)
			if len(g.onPullRequestEvent[PullRequestEventLockedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onPullRequestEvent[PullRequestEventLockedAction]))
			}
		})
	}
}

func TestSetOnPullRequestEventLocked(t *testing.T) {
	type args struct {
		callbacks []PullRequestEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single PullRequestEventHandleFunc",
			args: args{
				[]PullRequestEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple PullRequestEventHandleFuncs",
			args: args{
				[]PullRequestEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
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
			g.SetOnPullRequestEventLocked(func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
				return nil
			})
			g.SetOnPullRequestEventLocked(tt.args.callbacks...)
			if len(g.onPullRequestEvent[PullRequestEventLockedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onPullRequestEvent[PullRequestEventLockedAction]), tt.want)
			}
		})
	}
}

func TestHandlePullRequestEventLocked(t *testing.T) {
	action := PullRequestEventLockedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.PullRequestEvent
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
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnPullRequestEventLocked(func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handlePullRequestEventLocked(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handlePullRequestEventLocked() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnPullRequestEventOpened(t *testing.T) {
	type args struct {
		callbacks []PullRequestEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single PullRequestEventHandleFunc",
			args: args{
				callbacks: []PullRequestEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple PullRequestEventHandleFunc",
			args: args{
				callbacks: []PullRequestEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnPullRequestEventOpened(tt.args.callbacks...)
			if len(g.onPullRequestEvent[PullRequestEventOpenedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onPullRequestEvent[PullRequestEventOpenedAction]))
			}
		})
	}
}

func TestSetOnPullRequestEventOpened(t *testing.T) {
	type args struct {
		callbacks []PullRequestEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single PullRequestEventHandleFunc",
			args: args{
				[]PullRequestEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple PullRequestEventHandleFuncs",
			args: args{
				[]PullRequestEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
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
			g.SetOnPullRequestEventOpened(func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
				return nil
			})
			g.SetOnPullRequestEventOpened(tt.args.callbacks...)
			if len(g.onPullRequestEvent[PullRequestEventOpenedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onPullRequestEvent[PullRequestEventOpenedAction]), tt.want)
			}
		})
	}
}

func TestHandlePullRequestEventOpened(t *testing.T) {
	action := PullRequestEventOpenedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.PullRequestEvent
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
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnPullRequestEventOpened(func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handlePullRequestEventOpened(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handlePullRequestEventOpened() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnPullRequestEventReadyForReview(t *testing.T) {
	type args struct {
		callbacks []PullRequestEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single PullRequestEventHandleFunc",
			args: args{
				callbacks: []PullRequestEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple PullRequestEventHandleFunc",
			args: args{
				callbacks: []PullRequestEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnPullRequestEventReadyForReview(tt.args.callbacks...)
			if len(g.onPullRequestEvent[PullRequestEventReadyForReviewAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onPullRequestEvent[PullRequestEventReadyForReviewAction]))
			}
		})
	}
}

func TestSetOnPullRequestEventReadyForReview(t *testing.T) {
	type args struct {
		callbacks []PullRequestEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single PullRequestEventHandleFunc",
			args: args{
				[]PullRequestEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple PullRequestEventHandleFuncs",
			args: args{
				[]PullRequestEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
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
			g.SetOnPullRequestEventReadyForReview(func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
				return nil
			})
			g.SetOnPullRequestEventReadyForReview(tt.args.callbacks...)
			if len(g.onPullRequestEvent[PullRequestEventReadyForReviewAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onPullRequestEvent[PullRequestEventReadyForReviewAction]), tt.want)
			}
		})
	}
}

func TestHandlePullRequestEventReadyForReview(t *testing.T) {
	action := PullRequestEventReadyForReviewAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.PullRequestEvent
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
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnPullRequestEventReadyForReview(func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handlePullRequestEventReadyForReview(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handlePullRequestEventReadyForReview() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnPullRequestEventReopened(t *testing.T) {
	type args struct {
		callbacks []PullRequestEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single PullRequestEventHandleFunc",
			args: args{
				callbacks: []PullRequestEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple PullRequestEventHandleFunc",
			args: args{
				callbacks: []PullRequestEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnPullRequestEventReopened(tt.args.callbacks...)
			if len(g.onPullRequestEvent[PullRequestEventReopenedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onPullRequestEvent[PullRequestEventReopenedAction]))
			}
		})
	}
}

func TestSetOnPullRequestEventReopened(t *testing.T) {
	type args struct {
		callbacks []PullRequestEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single PullRequestEventHandleFunc",
			args: args{
				[]PullRequestEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple PullRequestEventHandleFuncs",
			args: args{
				[]PullRequestEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
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
			g.SetOnPullRequestEventReopened(func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
				return nil
			})
			g.SetOnPullRequestEventReopened(tt.args.callbacks...)
			if len(g.onPullRequestEvent[PullRequestEventReopenedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onPullRequestEvent[PullRequestEventReopenedAction]), tt.want)
			}
		})
	}
}

func TestHandlePullRequestEventReopened(t *testing.T) {
	action := PullRequestEventReopenedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.PullRequestEvent
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
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnPullRequestEventReopened(func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handlePullRequestEventReopened(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handlePullRequestEventReopened() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnPullRequestEventReviewRequestRemoved(t *testing.T) {
	type args struct {
		callbacks []PullRequestEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single PullRequestEventHandleFunc",
			args: args{
				callbacks: []PullRequestEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple PullRequestEventHandleFunc",
			args: args{
				callbacks: []PullRequestEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnPullRequestEventReviewRequestRemoved(tt.args.callbacks...)
			if len(g.onPullRequestEvent[PullRequestEventReviewRequestRemovedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onPullRequestEvent[PullRequestEventReviewRequestRemovedAction]))
			}
		})
	}
}

func TestSetOnPullRequestEventReviewRequestRemoved(t *testing.T) {
	type args struct {
		callbacks []PullRequestEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single PullRequestEventHandleFunc",
			args: args{
				[]PullRequestEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple PullRequestEventHandleFuncs",
			args: args{
				[]PullRequestEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
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
			g.SetOnPullRequestEventReviewRequestRemoved(func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
				return nil
			})
			g.SetOnPullRequestEventReviewRequestRemoved(tt.args.callbacks...)
			if len(g.onPullRequestEvent[PullRequestEventReviewRequestRemovedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onPullRequestEvent[PullRequestEventReviewRequestRemovedAction]), tt.want)
			}
		})
	}
}

func TestHandlePullRequestEventReviewRequestRemoved(t *testing.T) {
	action := PullRequestEventReviewRequestRemovedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.PullRequestEvent
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
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnPullRequestEventReviewRequestRemoved(func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handlePullRequestEventReviewRequestRemoved(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handlePullRequestEventReviewRequestRemoved() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnPullRequestEventReviewRequested(t *testing.T) {
	type args struct {
		callbacks []PullRequestEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single PullRequestEventHandleFunc",
			args: args{
				callbacks: []PullRequestEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple PullRequestEventHandleFunc",
			args: args{
				callbacks: []PullRequestEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnPullRequestEventReviewRequested(tt.args.callbacks...)
			if len(g.onPullRequestEvent[PullRequestEventReviewRequestedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onPullRequestEvent[PullRequestEventReviewRequestedAction]))
			}
		})
	}
}

func TestSetOnPullRequestEventReviewRequested(t *testing.T) {
	type args struct {
		callbacks []PullRequestEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single PullRequestEventHandleFunc",
			args: args{
				[]PullRequestEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple PullRequestEventHandleFuncs",
			args: args{
				[]PullRequestEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
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
			g.SetOnPullRequestEventReviewRequested(func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
				return nil
			})
			g.SetOnPullRequestEventReviewRequested(tt.args.callbacks...)
			if len(g.onPullRequestEvent[PullRequestEventReviewRequestedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onPullRequestEvent[PullRequestEventReviewRequestedAction]), tt.want)
			}
		})
	}
}

func TestHandlePullRequestEventReviewRequested(t *testing.T) {
	action := PullRequestEventReviewRequestedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.PullRequestEvent
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
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnPullRequestEventReviewRequested(func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handlePullRequestEventReviewRequested(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handlePullRequestEventReviewRequested() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnPullRequestEventSynchronize(t *testing.T) {
	type args struct {
		callbacks []PullRequestEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single PullRequestEventHandleFunc",
			args: args{
				callbacks: []PullRequestEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple PullRequestEventHandleFunc",
			args: args{
				callbacks: []PullRequestEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnPullRequestEventSynchronize(tt.args.callbacks...)
			if len(g.onPullRequestEvent[PullRequestEventSynchronizeAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onPullRequestEvent[PullRequestEventSynchronizeAction]))
			}
		})
	}
}

func TestSetOnPullRequestEventSynchronize(t *testing.T) {
	type args struct {
		callbacks []PullRequestEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single PullRequestEventHandleFunc",
			args: args{
				[]PullRequestEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple PullRequestEventHandleFuncs",
			args: args{
				[]PullRequestEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
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
			g.SetOnPullRequestEventSynchronize(func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
				return nil
			})
			g.SetOnPullRequestEventSynchronize(tt.args.callbacks...)
			if len(g.onPullRequestEvent[PullRequestEventSynchronizeAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onPullRequestEvent[PullRequestEventSynchronizeAction]), tt.want)
			}
		})
	}
}

func TestHandlePullRequestEventSynchronize(t *testing.T) {
	action := PullRequestEventSynchronizeAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.PullRequestEvent
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
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnPullRequestEventSynchronize(func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handlePullRequestEventSynchronize(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handlePullRequestEventSynchronize() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnPullRequestEventUnassigned(t *testing.T) {
	type args struct {
		callbacks []PullRequestEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single PullRequestEventHandleFunc",
			args: args{
				callbacks: []PullRequestEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple PullRequestEventHandleFunc",
			args: args{
				callbacks: []PullRequestEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnPullRequestEventUnassigned(tt.args.callbacks...)
			if len(g.onPullRequestEvent[PullRequestEventUnassignedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onPullRequestEvent[PullRequestEventUnassignedAction]))
			}
		})
	}
}

func TestSetOnPullRequestEventUnassigned(t *testing.T) {
	type args struct {
		callbacks []PullRequestEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single PullRequestEventHandleFunc",
			args: args{
				[]PullRequestEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple PullRequestEventHandleFuncs",
			args: args{
				[]PullRequestEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
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
			g.SetOnPullRequestEventUnassigned(func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
				return nil
			})
			g.SetOnPullRequestEventUnassigned(tt.args.callbacks...)
			if len(g.onPullRequestEvent[PullRequestEventUnassignedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onPullRequestEvent[PullRequestEventUnassignedAction]), tt.want)
			}
		})
	}
}

func TestHandlePullRequestEventUnassigned(t *testing.T) {
	action := PullRequestEventUnassignedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.PullRequestEvent
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
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnPullRequestEventUnassigned(func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handlePullRequestEventUnassigned(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handlePullRequestEventUnassigned() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnPullRequestEventUnlabeled(t *testing.T) {
	type args struct {
		callbacks []PullRequestEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single PullRequestEventHandleFunc",
			args: args{
				callbacks: []PullRequestEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple PullRequestEventHandleFunc",
			args: args{
				callbacks: []PullRequestEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnPullRequestEventUnlabeled(tt.args.callbacks...)
			if len(g.onPullRequestEvent[PullRequestEventUnlabeledAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onPullRequestEvent[PullRequestEventUnlabeledAction]))
			}
		})
	}
}

func TestSetOnPullRequestEventUnlabeled(t *testing.T) {
	type args struct {
		callbacks []PullRequestEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single PullRequestEventHandleFunc",
			args: args{
				[]PullRequestEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple PullRequestEventHandleFuncs",
			args: args{
				[]PullRequestEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
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
			g.SetOnPullRequestEventUnlabeled(func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
				return nil
			})
			g.SetOnPullRequestEventUnlabeled(tt.args.callbacks...)
			if len(g.onPullRequestEvent[PullRequestEventUnlabeledAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onPullRequestEvent[PullRequestEventUnlabeledAction]), tt.want)
			}
		})
	}
}

func TestHandlePullRequestEventUnlabeled(t *testing.T) {
	action := PullRequestEventUnlabeledAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.PullRequestEvent
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
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnPullRequestEventUnlabeled(func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handlePullRequestEventUnlabeled(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handlePullRequestEventUnlabeled() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnPullRequestEventUnlocked(t *testing.T) {
	type args struct {
		callbacks []PullRequestEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single PullRequestEventHandleFunc",
			args: args{
				callbacks: []PullRequestEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple PullRequestEventHandleFunc",
			args: args{
				callbacks: []PullRequestEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnPullRequestEventUnlocked(tt.args.callbacks...)
			if len(g.onPullRequestEvent[PullRequestEventUnlockedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onPullRequestEvent[PullRequestEventUnlockedAction]))
			}
		})
	}
}

func TestSetOnPullRequestEventUnlocked(t *testing.T) {
	type args struct {
		callbacks []PullRequestEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single PullRequestEventHandleFunc",
			args: args{
				[]PullRequestEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple PullRequestEventHandleFuncs",
			args: args{
				[]PullRequestEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
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
			g.SetOnPullRequestEventUnlocked(func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
				return nil
			})
			g.SetOnPullRequestEventUnlocked(tt.args.callbacks...)
			if len(g.onPullRequestEvent[PullRequestEventUnlockedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onPullRequestEvent[PullRequestEventUnlockedAction]), tt.want)
			}
		})
	}
}

func TestHandlePullRequestEventUnlocked(t *testing.T) {
	action := PullRequestEventUnlockedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.PullRequestEvent
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
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request",
				event:      &github.PullRequestEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnPullRequestEventUnlocked(func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handlePullRequestEventUnlocked(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handlePullRequestEventUnlocked() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
