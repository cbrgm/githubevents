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

func TestOnStarEventAny(t *testing.T) {
	type args struct {
		callbacks []StarEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single StarEventHandleFunc",
			args: args{
				[]StarEventHandleFunc{
					func(deliveryID string, eventName string, event *github.StarEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple StarEventHandleFuncs",
			args: args{
				[]StarEventHandleFunc{
					func(deliveryID string, eventName string, event *github.StarEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.StarEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnStarEventAny(tt.args.callbacks...)
			if len(g.onStarEvent[StarEventAnyAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onStarEvent[StarEventAnyAction]))
			}
		})
	}
}

func TestSetOnStarEventAny(t *testing.T) {
	type args struct {
		callbacks []StarEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single StarEventHandleFunc",
			args: args{
				[]StarEventHandleFunc{
					func(deliveryID string, eventName string, event *github.StarEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple StarEventHandleFuncs",
			args: args{
				[]StarEventHandleFunc{
					func(deliveryID string, eventName string, event *github.StarEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.StarEvent) error {
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
			g.SetOnStarEventAny(func(deliveryID string, eventName string, event *github.StarEvent) error {
				return nil
			})
			g.SetOnStarEventAny(tt.args.callbacks...)
			if len(g.onStarEvent[StarEventAnyAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onStarEvent[StarEventAnyAction]), tt.want)
			}
		})
	}
}

func TestHandleStarEventAny(t *testing.T) {

	action := "*"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.StarEvent
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
				eventName:  "star",

				event: &github.StarEvent{Action: &action},

				fail: false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "star",

				event: &github.StarEvent{Action: &action},

				fail: true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "star",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnStarEventAny(func(deliveryID string, eventName string, event *github.StarEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleStarEventAny(tt.args.deliveryID, tt.args.deliveryID, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("TestHandleStarEventAny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnStarEventCreated(t *testing.T) {
	type args struct {
		callbacks []StarEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single StarEventHandleFunc",
			args: args{
				callbacks: []StarEventHandleFunc{
					func(deliveryID string, eventName string, event *github.StarEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple StarEventHandleFunc",
			args: args{
				callbacks: []StarEventHandleFunc{
					func(deliveryID string, eventName string, event *github.StarEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.StarEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnStarEventCreated(tt.args.callbacks...)
			if len(g.onStarEvent[StarEventCreatedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onStarEvent[StarEventCreatedAction]))
			}
		})
	}
}

func TestSetOnStarEventCreated(t *testing.T) {
	type args struct {
		callbacks []StarEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single StarEventHandleFunc",
			args: args{
				[]StarEventHandleFunc{
					func(deliveryID string, eventName string, event *github.StarEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple StarEventHandleFuncs",
			args: args{
				[]StarEventHandleFunc{
					func(deliveryID string, eventName string, event *github.StarEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.StarEvent) error {
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
			g.SetOnStarEventCreated(func(deliveryID string, eventName string, event *github.StarEvent) error {
				return nil
			})
			g.SetOnStarEventCreated(tt.args.callbacks...)
			if len(g.onStarEvent[StarEventCreatedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onStarEvent[StarEventCreatedAction]), tt.want)
			}
		})
	}
}

func TestHandleStarEventCreated(t *testing.T) {
	action := StarEventCreatedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.StarEvent
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
				eventName:  "star",
				event:      &github.StarEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "star",
				event:      &github.StarEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "star",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "star",
				event:      &github.StarEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "star",
				event:      &github.StarEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "star",
				event:      &github.StarEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnStarEventCreated(func(deliveryID string, eventName string, event *github.StarEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleStarEventCreated(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleStarEventCreated() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnStarEventDeleted(t *testing.T) {
	type args struct {
		callbacks []StarEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single StarEventHandleFunc",
			args: args{
				callbacks: []StarEventHandleFunc{
					func(deliveryID string, eventName string, event *github.StarEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple StarEventHandleFunc",
			args: args{
				callbacks: []StarEventHandleFunc{
					func(deliveryID string, eventName string, event *github.StarEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.StarEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnStarEventDeleted(tt.args.callbacks...)
			if len(g.onStarEvent[StarEventDeletedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onStarEvent[StarEventDeletedAction]))
			}
		})
	}
}

func TestSetOnStarEventDeleted(t *testing.T) {
	type args struct {
		callbacks []StarEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single StarEventHandleFunc",
			args: args{
				[]StarEventHandleFunc{
					func(deliveryID string, eventName string, event *github.StarEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple StarEventHandleFuncs",
			args: args{
				[]StarEventHandleFunc{
					func(deliveryID string, eventName string, event *github.StarEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.StarEvent) error {
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
			g.SetOnStarEventDeleted(func(deliveryID string, eventName string, event *github.StarEvent) error {
				return nil
			})
			g.SetOnStarEventDeleted(tt.args.callbacks...)
			if len(g.onStarEvent[StarEventDeletedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onStarEvent[StarEventDeletedAction]), tt.want)
			}
		})
	}
}

func TestHandleStarEventDeleted(t *testing.T) {
	action := StarEventDeletedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.StarEvent
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
				eventName:  "star",
				event:      &github.StarEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "star",
				event:      &github.StarEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "star",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "star",
				event:      &github.StarEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "star",
				event:      &github.StarEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "star",
				event:      &github.StarEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnStarEventDeleted(func(deliveryID string, eventName string, event *github.StarEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleStarEventDeleted(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleStarEventDeleted() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
