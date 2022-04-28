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

func TestOnReleaseEventAny(t *testing.T) {
	type args struct {
		callbacks []ReleaseEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single ReleaseEventHandleFunc",
			args: args{
				[]ReleaseEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ReleaseEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple ReleaseEventHandleFuncs",
			args: args{
				[]ReleaseEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ReleaseEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ReleaseEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnReleaseEventAny(tt.args.callbacks...)
			if len(g.onReleaseEvent[ReleaseEventAnyAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onReleaseEvent[ReleaseEventAnyAction]))
			}
		})
	}
}

func TestSetOnReleaseEventAny(t *testing.T) {
	type args struct {
		callbacks []ReleaseEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single ReleaseEventHandleFunc",
			args: args{
				[]ReleaseEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ReleaseEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple ReleaseEventHandleFuncs",
			args: args{
				[]ReleaseEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ReleaseEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ReleaseEvent) error {
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
			g.SetOnReleaseEventAny(func(deliveryID string, eventName string, event *github.ReleaseEvent) error {
				return nil
			})
			g.SetOnReleaseEventAny(tt.args.callbacks...)
			if len(g.onReleaseEvent[ReleaseEventAnyAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onReleaseEvent[ReleaseEventAnyAction]), tt.want)
			}
		})
	}
}

func TestHandleReleaseEventAny(t *testing.T) {

	action := "*"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.ReleaseEvent
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
				eventName:  "release",

				event: &github.ReleaseEvent{Action: &action},

				fail: false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "release",

				event: &github.ReleaseEvent{Action: &action},

				fail: true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "release",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnReleaseEventAny(func(deliveryID string, eventName string, event *github.ReleaseEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleReleaseEventAny(tt.args.deliveryID, tt.args.deliveryID, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("TestHandleReleaseEventAny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnReleaseEventPublished(t *testing.T) {
	type args struct {
		callbacks []ReleaseEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single ReleaseEventHandleFunc",
			args: args{
				callbacks: []ReleaseEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ReleaseEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple ReleaseEventHandleFunc",
			args: args{
				callbacks: []ReleaseEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ReleaseEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ReleaseEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnReleaseEventPublished(tt.args.callbacks...)
			if len(g.onReleaseEvent[ReleaseEventPublishedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onReleaseEvent[ReleaseEventPublishedAction]))
			}
		})
	}
}

func TestSetOnReleaseEventPublished(t *testing.T) {
	type args struct {
		callbacks []ReleaseEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single ReleaseEventHandleFunc",
			args: args{
				[]ReleaseEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ReleaseEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple ReleaseEventHandleFuncs",
			args: args{
				[]ReleaseEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ReleaseEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ReleaseEvent) error {
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
			g.SetOnReleaseEventPublished(func(deliveryID string, eventName string, event *github.ReleaseEvent) error {
				return nil
			})
			g.SetOnReleaseEventPublished(tt.args.callbacks...)
			if len(g.onReleaseEvent[ReleaseEventPublishedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onReleaseEvent[ReleaseEventPublishedAction]), tt.want)
			}
		})
	}
}

func TestHandleReleaseEventPublished(t *testing.T) {
	action := ReleaseEventPublishedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.ReleaseEvent
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
				eventName:  "release",
				event:      &github.ReleaseEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "release",
				event:      &github.ReleaseEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "release",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "release",
				event:      &github.ReleaseEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "release",
				event:      &github.ReleaseEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "release",
				event:      &github.ReleaseEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnReleaseEventPublished(func(deliveryID string, eventName string, event *github.ReleaseEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleReleaseEventPublished(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleReleaseEventPublished() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnReleaseEventUnpublished(t *testing.T) {
	type args struct {
		callbacks []ReleaseEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single ReleaseEventHandleFunc",
			args: args{
				callbacks: []ReleaseEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ReleaseEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple ReleaseEventHandleFunc",
			args: args{
				callbacks: []ReleaseEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ReleaseEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ReleaseEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnReleaseEventUnpublished(tt.args.callbacks...)
			if len(g.onReleaseEvent[ReleaseEventUnpublishedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onReleaseEvent[ReleaseEventUnpublishedAction]))
			}
		})
	}
}

func TestSetOnReleaseEventUnpublished(t *testing.T) {
	type args struct {
		callbacks []ReleaseEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single ReleaseEventHandleFunc",
			args: args{
				[]ReleaseEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ReleaseEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple ReleaseEventHandleFuncs",
			args: args{
				[]ReleaseEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ReleaseEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ReleaseEvent) error {
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
			g.SetOnReleaseEventUnpublished(func(deliveryID string, eventName string, event *github.ReleaseEvent) error {
				return nil
			})
			g.SetOnReleaseEventUnpublished(tt.args.callbacks...)
			if len(g.onReleaseEvent[ReleaseEventUnpublishedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onReleaseEvent[ReleaseEventUnpublishedAction]), tt.want)
			}
		})
	}
}

func TestHandleReleaseEventUnpublished(t *testing.T) {
	action := ReleaseEventUnpublishedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.ReleaseEvent
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
				eventName:  "release",
				event:      &github.ReleaseEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "release",
				event:      &github.ReleaseEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "release",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "release",
				event:      &github.ReleaseEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "release",
				event:      &github.ReleaseEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "release",
				event:      &github.ReleaseEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnReleaseEventUnpublished(func(deliveryID string, eventName string, event *github.ReleaseEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleReleaseEventUnpublished(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleReleaseEventUnpublished() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnReleaseEventCreated(t *testing.T) {
	type args struct {
		callbacks []ReleaseEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single ReleaseEventHandleFunc",
			args: args{
				callbacks: []ReleaseEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ReleaseEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple ReleaseEventHandleFunc",
			args: args{
				callbacks: []ReleaseEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ReleaseEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ReleaseEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnReleaseEventCreated(tt.args.callbacks...)
			if len(g.onReleaseEvent[ReleaseEventCreatedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onReleaseEvent[ReleaseEventCreatedAction]))
			}
		})
	}
}

func TestSetOnReleaseEventCreated(t *testing.T) {
	type args struct {
		callbacks []ReleaseEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single ReleaseEventHandleFunc",
			args: args{
				[]ReleaseEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ReleaseEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple ReleaseEventHandleFuncs",
			args: args{
				[]ReleaseEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ReleaseEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ReleaseEvent) error {
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
			g.SetOnReleaseEventCreated(func(deliveryID string, eventName string, event *github.ReleaseEvent) error {
				return nil
			})
			g.SetOnReleaseEventCreated(tt.args.callbacks...)
			if len(g.onReleaseEvent[ReleaseEventCreatedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onReleaseEvent[ReleaseEventCreatedAction]), tt.want)
			}
		})
	}
}

func TestHandleReleaseEventCreated(t *testing.T) {
	action := ReleaseEventCreatedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.ReleaseEvent
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
				eventName:  "release",
				event:      &github.ReleaseEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "release",
				event:      &github.ReleaseEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "release",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "release",
				event:      &github.ReleaseEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "release",
				event:      &github.ReleaseEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "release",
				event:      &github.ReleaseEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnReleaseEventCreated(func(deliveryID string, eventName string, event *github.ReleaseEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleReleaseEventCreated(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleReleaseEventCreated() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnReleaseEventEdited(t *testing.T) {
	type args struct {
		callbacks []ReleaseEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single ReleaseEventHandleFunc",
			args: args{
				callbacks: []ReleaseEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ReleaseEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple ReleaseEventHandleFunc",
			args: args{
				callbacks: []ReleaseEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ReleaseEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ReleaseEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnReleaseEventEdited(tt.args.callbacks...)
			if len(g.onReleaseEvent[ReleaseEventEditedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onReleaseEvent[ReleaseEventEditedAction]))
			}
		})
	}
}

func TestSetOnReleaseEventEdited(t *testing.T) {
	type args struct {
		callbacks []ReleaseEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single ReleaseEventHandleFunc",
			args: args{
				[]ReleaseEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ReleaseEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple ReleaseEventHandleFuncs",
			args: args{
				[]ReleaseEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ReleaseEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ReleaseEvent) error {
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
			g.SetOnReleaseEventEdited(func(deliveryID string, eventName string, event *github.ReleaseEvent) error {
				return nil
			})
			g.SetOnReleaseEventEdited(tt.args.callbacks...)
			if len(g.onReleaseEvent[ReleaseEventEditedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onReleaseEvent[ReleaseEventEditedAction]), tt.want)
			}
		})
	}
}

func TestHandleReleaseEventEdited(t *testing.T) {
	action := ReleaseEventEditedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.ReleaseEvent
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
				eventName:  "release",
				event:      &github.ReleaseEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "release",
				event:      &github.ReleaseEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "release",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "release",
				event:      &github.ReleaseEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "release",
				event:      &github.ReleaseEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "release",
				event:      &github.ReleaseEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnReleaseEventEdited(func(deliveryID string, eventName string, event *github.ReleaseEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleReleaseEventEdited(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleReleaseEventEdited() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnReleaseEventDeleted(t *testing.T) {
	type args struct {
		callbacks []ReleaseEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single ReleaseEventHandleFunc",
			args: args{
				callbacks: []ReleaseEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ReleaseEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple ReleaseEventHandleFunc",
			args: args{
				callbacks: []ReleaseEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ReleaseEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ReleaseEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnReleaseEventDeleted(tt.args.callbacks...)
			if len(g.onReleaseEvent[ReleaseEventDeletedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onReleaseEvent[ReleaseEventDeletedAction]))
			}
		})
	}
}

func TestSetOnReleaseEventDeleted(t *testing.T) {
	type args struct {
		callbacks []ReleaseEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single ReleaseEventHandleFunc",
			args: args{
				[]ReleaseEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ReleaseEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple ReleaseEventHandleFuncs",
			args: args{
				[]ReleaseEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ReleaseEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ReleaseEvent) error {
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
			g.SetOnReleaseEventDeleted(func(deliveryID string, eventName string, event *github.ReleaseEvent) error {
				return nil
			})
			g.SetOnReleaseEventDeleted(tt.args.callbacks...)
			if len(g.onReleaseEvent[ReleaseEventDeletedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onReleaseEvent[ReleaseEventDeletedAction]), tt.want)
			}
		})
	}
}

func TestHandleReleaseEventDeleted(t *testing.T) {
	action := ReleaseEventDeletedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.ReleaseEvent
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
				eventName:  "release",
				event:      &github.ReleaseEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "release",
				event:      &github.ReleaseEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "release",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "release",
				event:      &github.ReleaseEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "release",
				event:      &github.ReleaseEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "release",
				event:      &github.ReleaseEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnReleaseEventDeleted(func(deliveryID string, eventName string, event *github.ReleaseEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleReleaseEventDeleted(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleReleaseEventDeleted() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnReleaseEventPreReleased(t *testing.T) {
	type args struct {
		callbacks []ReleaseEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single ReleaseEventHandleFunc",
			args: args{
				callbacks: []ReleaseEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ReleaseEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple ReleaseEventHandleFunc",
			args: args{
				callbacks: []ReleaseEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ReleaseEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ReleaseEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnReleaseEventPreReleased(tt.args.callbacks...)
			if len(g.onReleaseEvent[ReleaseEventPreReleasedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onReleaseEvent[ReleaseEventPreReleasedAction]))
			}
		})
	}
}

func TestSetOnReleaseEventPreReleased(t *testing.T) {
	type args struct {
		callbacks []ReleaseEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single ReleaseEventHandleFunc",
			args: args{
				[]ReleaseEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ReleaseEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple ReleaseEventHandleFuncs",
			args: args{
				[]ReleaseEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ReleaseEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ReleaseEvent) error {
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
			g.SetOnReleaseEventPreReleased(func(deliveryID string, eventName string, event *github.ReleaseEvent) error {
				return nil
			})
			g.SetOnReleaseEventPreReleased(tt.args.callbacks...)
			if len(g.onReleaseEvent[ReleaseEventPreReleasedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onReleaseEvent[ReleaseEventPreReleasedAction]), tt.want)
			}
		})
	}
}

func TestHandleReleaseEventPreReleased(t *testing.T) {
	action := ReleaseEventPreReleasedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.ReleaseEvent
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
				eventName:  "release",
				event:      &github.ReleaseEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "release",
				event:      &github.ReleaseEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "release",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "release",
				event:      &github.ReleaseEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "release",
				event:      &github.ReleaseEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "release",
				event:      &github.ReleaseEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnReleaseEventPreReleased(func(deliveryID string, eventName string, event *github.ReleaseEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleReleaseEventPreReleased(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleReleaseEventPreReleased() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnReleaseEventReleased(t *testing.T) {
	type args struct {
		callbacks []ReleaseEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single ReleaseEventHandleFunc",
			args: args{
				callbacks: []ReleaseEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ReleaseEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple ReleaseEventHandleFunc",
			args: args{
				callbacks: []ReleaseEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ReleaseEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ReleaseEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnReleaseEventReleased(tt.args.callbacks...)
			if len(g.onReleaseEvent[ReleaseEventReleasedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onReleaseEvent[ReleaseEventReleasedAction]))
			}
		})
	}
}

func TestSetOnReleaseEventReleased(t *testing.T) {
	type args struct {
		callbacks []ReleaseEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single ReleaseEventHandleFunc",
			args: args{
				[]ReleaseEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ReleaseEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple ReleaseEventHandleFuncs",
			args: args{
				[]ReleaseEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ReleaseEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ReleaseEvent) error {
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
			g.SetOnReleaseEventReleased(func(deliveryID string, eventName string, event *github.ReleaseEvent) error {
				return nil
			})
			g.SetOnReleaseEventReleased(tt.args.callbacks...)
			if len(g.onReleaseEvent[ReleaseEventReleasedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onReleaseEvent[ReleaseEventReleasedAction]), tt.want)
			}
		})
	}
}

func TestHandleReleaseEventReleased(t *testing.T) {
	action := ReleaseEventReleasedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.ReleaseEvent
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
				eventName:  "release",
				event:      &github.ReleaseEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "release",
				event:      &github.ReleaseEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "release",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "release",
				event:      &github.ReleaseEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "release",
				event:      &github.ReleaseEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "release",
				event:      &github.ReleaseEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnReleaseEventReleased(func(deliveryID string, eventName string, event *github.ReleaseEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleReleaseEventReleased(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleReleaseEventReleased() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
