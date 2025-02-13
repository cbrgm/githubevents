// Copyright 2022 The GithubEvents Authors. All rights reserved.
// Use of this source code is governed by the MIT License
// that can be found in the LICENSE file.

package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"context"
	"errors"
	"github.com/google/go-github/v69/github"
	"go.opentelemetry.io/otel/trace/noop"
	"sync"
	"testing"
)

func TestOnMilestoneEventAny(t *testing.T) {
	type args struct {
		callbacks []MilestoneEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single MilestoneEventHandleFunc",
			args: args{
				[]MilestoneEventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple MilestoneEventHandleFuncs",
			args: args{
				[]MilestoneEventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
						return nil
					},
					func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnMilestoneEventAny(tt.args.callbacks...)
			if len(g.onMilestoneEvent[MilestoneEventAnyAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onMilestoneEvent[MilestoneEventAnyAction]))
			}
		})
	}
}

func TestSetOnMilestoneEventAny(t *testing.T) {
	type args struct {
		callbacks []MilestoneEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single MilestoneEventHandleFunc",
			args: args{
				[]MilestoneEventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple MilestoneEventHandleFuncs",
			args: args{
				[]MilestoneEventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
						return nil
					},
					func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
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
			g.SetOnMilestoneEventAny(func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
				return nil
			})
			g.SetOnMilestoneEventAny(tt.args.callbacks...)
			if len(g.onMilestoneEvent[MilestoneEventAnyAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onMilestoneEvent[MilestoneEventAnyAction]), tt.want)
			}
		})
	}
}

func TestHandleMilestoneEventAny(t *testing.T) {

	action := "*"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.MilestoneEvent
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
				eventName:  "milestone",

				event: &github.MilestoneEvent{Action: &action},

				fail: false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "milestone",

				event: &github.MilestoneEvent{Action: &action},

				fail: true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "milestone",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnMilestoneEventAny(func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleMilestoneEventAny(context.Background(), tt.args.deliveryID, tt.args.deliveryID, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("TestHandleMilestoneEventAny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnMilestoneEventCreated(t *testing.T) {
	type args struct {
		callbacks []MilestoneEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single MilestoneEventHandleFunc",
			args: args{
				callbacks: []MilestoneEventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple MilestoneEventHandleFunc",
			args: args{
				callbacks: []MilestoneEventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
						return nil
					},
					func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnMilestoneEventCreated(tt.args.callbacks...)
			if len(g.onMilestoneEvent[MilestoneEventCreatedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onMilestoneEvent[MilestoneEventCreatedAction]))
			}
		})
	}
}

func TestSetOnMilestoneEventCreated(t *testing.T) {
	type args struct {
		callbacks []MilestoneEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single MilestoneEventHandleFunc",
			args: args{
				[]MilestoneEventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple MilestoneEventHandleFuncs",
			args: args{
				[]MilestoneEventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
						return nil
					},
					func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
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
			g.SetOnMilestoneEventCreated(func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
				return nil
			})
			g.SetOnMilestoneEventCreated(tt.args.callbacks...)
			if len(g.onMilestoneEvent[MilestoneEventCreatedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onMilestoneEvent[MilestoneEventCreatedAction]), tt.want)
			}
		})
	}
}

func TestHandleMilestoneEventCreated(t *testing.T) {
	action := MilestoneEventCreatedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.MilestoneEvent
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
				eventName:  "milestone",
				event:      &github.MilestoneEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "milestone",
				event:      &github.MilestoneEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "milestone",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "milestone",
				event:      &github.MilestoneEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "milestone",
				event:      &github.MilestoneEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "milestone",
				event:      &github.MilestoneEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnMilestoneEventCreated(func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleMilestoneEventCreated(context.Background(), tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleMilestoneEventCreated() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnMilestoneEventClosed(t *testing.T) {
	type args struct {
		callbacks []MilestoneEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single MilestoneEventHandleFunc",
			args: args{
				callbacks: []MilestoneEventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple MilestoneEventHandleFunc",
			args: args{
				callbacks: []MilestoneEventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
						return nil
					},
					func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnMilestoneEventClosed(tt.args.callbacks...)
			if len(g.onMilestoneEvent[MilestoneEventClosedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onMilestoneEvent[MilestoneEventClosedAction]))
			}
		})
	}
}

func TestSetOnMilestoneEventClosed(t *testing.T) {
	type args struct {
		callbacks []MilestoneEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single MilestoneEventHandleFunc",
			args: args{
				[]MilestoneEventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple MilestoneEventHandleFuncs",
			args: args{
				[]MilestoneEventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
						return nil
					},
					func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
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
			g.SetOnMilestoneEventClosed(func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
				return nil
			})
			g.SetOnMilestoneEventClosed(tt.args.callbacks...)
			if len(g.onMilestoneEvent[MilestoneEventClosedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onMilestoneEvent[MilestoneEventClosedAction]), tt.want)
			}
		})
	}
}

func TestHandleMilestoneEventClosed(t *testing.T) {
	action := MilestoneEventClosedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.MilestoneEvent
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
				eventName:  "milestone",
				event:      &github.MilestoneEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "milestone",
				event:      &github.MilestoneEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "milestone",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "milestone",
				event:      &github.MilestoneEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "milestone",
				event:      &github.MilestoneEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "milestone",
				event:      &github.MilestoneEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnMilestoneEventClosed(func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleMilestoneEventClosed(context.Background(), tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleMilestoneEventClosed() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnMilestoneEventOpened(t *testing.T) {
	type args struct {
		callbacks []MilestoneEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single MilestoneEventHandleFunc",
			args: args{
				callbacks: []MilestoneEventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple MilestoneEventHandleFunc",
			args: args{
				callbacks: []MilestoneEventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
						return nil
					},
					func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnMilestoneEventOpened(tt.args.callbacks...)
			if len(g.onMilestoneEvent[MilestoneEventOpenedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onMilestoneEvent[MilestoneEventOpenedAction]))
			}
		})
	}
}

func TestSetOnMilestoneEventOpened(t *testing.T) {
	type args struct {
		callbacks []MilestoneEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single MilestoneEventHandleFunc",
			args: args{
				[]MilestoneEventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple MilestoneEventHandleFuncs",
			args: args{
				[]MilestoneEventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
						return nil
					},
					func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
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
			g.SetOnMilestoneEventOpened(func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
				return nil
			})
			g.SetOnMilestoneEventOpened(tt.args.callbacks...)
			if len(g.onMilestoneEvent[MilestoneEventOpenedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onMilestoneEvent[MilestoneEventOpenedAction]), tt.want)
			}
		})
	}
}

func TestHandleMilestoneEventOpened(t *testing.T) {
	action := MilestoneEventOpenedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.MilestoneEvent
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
				eventName:  "milestone",
				event:      &github.MilestoneEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "milestone",
				event:      &github.MilestoneEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "milestone",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "milestone",
				event:      &github.MilestoneEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "milestone",
				event:      &github.MilestoneEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "milestone",
				event:      &github.MilestoneEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnMilestoneEventOpened(func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleMilestoneEventOpened(context.Background(), tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleMilestoneEventOpened() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnMilestoneEventEdited(t *testing.T) {
	type args struct {
		callbacks []MilestoneEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single MilestoneEventHandleFunc",
			args: args{
				callbacks: []MilestoneEventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple MilestoneEventHandleFunc",
			args: args{
				callbacks: []MilestoneEventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
						return nil
					},
					func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnMilestoneEventEdited(tt.args.callbacks...)
			if len(g.onMilestoneEvent[MilestoneEventEditedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onMilestoneEvent[MilestoneEventEditedAction]))
			}
		})
	}
}

func TestSetOnMilestoneEventEdited(t *testing.T) {
	type args struct {
		callbacks []MilestoneEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single MilestoneEventHandleFunc",
			args: args{
				[]MilestoneEventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple MilestoneEventHandleFuncs",
			args: args{
				[]MilestoneEventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
						return nil
					},
					func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
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
			g.SetOnMilestoneEventEdited(func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
				return nil
			})
			g.SetOnMilestoneEventEdited(tt.args.callbacks...)
			if len(g.onMilestoneEvent[MilestoneEventEditedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onMilestoneEvent[MilestoneEventEditedAction]), tt.want)
			}
		})
	}
}

func TestHandleMilestoneEventEdited(t *testing.T) {
	action := MilestoneEventEditedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.MilestoneEvent
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
				eventName:  "milestone",
				event:      &github.MilestoneEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "milestone",
				event:      &github.MilestoneEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "milestone",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "milestone",
				event:      &github.MilestoneEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "milestone",
				event:      &github.MilestoneEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "milestone",
				event:      &github.MilestoneEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnMilestoneEventEdited(func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleMilestoneEventEdited(context.Background(), tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleMilestoneEventEdited() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnMilestoneEventDeleted(t *testing.T) {
	type args struct {
		callbacks []MilestoneEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single MilestoneEventHandleFunc",
			args: args{
				callbacks: []MilestoneEventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple MilestoneEventHandleFunc",
			args: args{
				callbacks: []MilestoneEventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
						return nil
					},
					func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnMilestoneEventDeleted(tt.args.callbacks...)
			if len(g.onMilestoneEvent[MilestoneEventDeletedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onMilestoneEvent[MilestoneEventDeletedAction]))
			}
		})
	}
}

func TestSetOnMilestoneEventDeleted(t *testing.T) {
	type args struct {
		callbacks []MilestoneEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single MilestoneEventHandleFunc",
			args: args{
				[]MilestoneEventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple MilestoneEventHandleFuncs",
			args: args{
				[]MilestoneEventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
						return nil
					},
					func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
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
			g.SetOnMilestoneEventDeleted(func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
				return nil
			})
			g.SetOnMilestoneEventDeleted(tt.args.callbacks...)
			if len(g.onMilestoneEvent[MilestoneEventDeletedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onMilestoneEvent[MilestoneEventDeletedAction]), tt.want)
			}
		})
	}
}

func TestHandleMilestoneEventDeleted(t *testing.T) {
	action := MilestoneEventDeletedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.MilestoneEvent
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
				eventName:  "milestone",
				event:      &github.MilestoneEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "milestone",
				event:      &github.MilestoneEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "milestone",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "milestone",
				event:      &github.MilestoneEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "milestone",
				event:      &github.MilestoneEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "milestone",
				event:      &github.MilestoneEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnMilestoneEventDeleted(func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleMilestoneEventDeleted(context.Background(), tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleMilestoneEventDeleted() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMilestoneEvent(t *testing.T) {
	type fields struct {
		handler *EventHandler
	}
	type args struct {
		deliveryID string
		eventName  string
		event      *github.MilestoneEvent
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "must trigger MilestoneEventAny with unknown event action",
			fields: fields{
				handler: &EventHandler{
					WebhookSecret: "fake",
					onBeforeAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onBeforeAny called")
								return nil
							},
						},
					},
					onAfterAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onAfterAny called")
								return nil
							},
						},
					},
					onMilestoneEvent: map[string][]MilestoneEventHandleFunc{
						MilestoneEventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
					},
					Tracer: noop.Tracer{},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  MilestoneEvent,

				event: &github.MilestoneEvent{Action: ptrString("unknown")},
			},
			wantErr: false,
		},

		{
			name: "must trigger MilestoneEventCreated",
			fields: fields{
				handler: &EventHandler{
					WebhookSecret: "fake",
					onBeforeAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onBeforeAny called")
								return nil
							},
						},
					},
					onAfterAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onAfterAny called")
								return nil
							},
						},
					},
					onMilestoneEvent: map[string][]MilestoneEventHandleFunc{
						MilestoneEventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						MilestoneEventCreatedAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
								t.Logf("%s action called", MilestoneEventCreatedAction)
								return nil
							},
						},
					},
					Tracer: noop.Tracer{},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "milestone",
				event:      &github.MilestoneEvent{Action: ptrString(MilestoneEventCreatedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail MilestoneEventCreated with empty action",
			fields: fields{
				handler: &EventHandler{
					WebhookSecret: "fake",
					onBeforeAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onBeforeAny called")
								return nil
							},
						},
					},
					onAfterAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onAfterAny called")
								return nil
							},
						},
					},
					onMilestoneEvent: map[string][]MilestoneEventHandleFunc{
						MilestoneEventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						MilestoneEventCreatedAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
								t.Logf("%s action called", MilestoneEventCreatedAction)
								return nil
							},
						},
					},
					Tracer: noop.Tracer{},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "milestone",
				event:      &github.MilestoneEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail MilestoneEventCreated with nil action",
			fields: fields{
				handler: &EventHandler{
					WebhookSecret: "fake",
					onBeforeAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onBeforeAny called")
								return nil
							},
						},
					},
					onAfterAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onAfterAny called")
								return nil
							},
						},
					},
					onMilestoneEvent: map[string][]MilestoneEventHandleFunc{
						MilestoneEventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						MilestoneEventCreatedAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
								t.Logf("%s action called", MilestoneEventCreatedAction)
								return nil
							},
						},
					},
					Tracer: noop.Tracer{},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "milestone",
				event:      &github.MilestoneEvent{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger MilestoneEventClosed",
			fields: fields{
				handler: &EventHandler{
					WebhookSecret: "fake",
					onBeforeAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onBeforeAny called")
								return nil
							},
						},
					},
					onAfterAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onAfterAny called")
								return nil
							},
						},
					},
					onMilestoneEvent: map[string][]MilestoneEventHandleFunc{
						MilestoneEventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						MilestoneEventClosedAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
								t.Logf("%s action called", MilestoneEventClosedAction)
								return nil
							},
						},
					},
					Tracer: noop.Tracer{},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "milestone",
				event:      &github.MilestoneEvent{Action: ptrString(MilestoneEventClosedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail MilestoneEventClosed with empty action",
			fields: fields{
				handler: &EventHandler{
					WebhookSecret: "fake",
					onBeforeAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onBeforeAny called")
								return nil
							},
						},
					},
					onAfterAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onAfterAny called")
								return nil
							},
						},
					},
					onMilestoneEvent: map[string][]MilestoneEventHandleFunc{
						MilestoneEventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						MilestoneEventClosedAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
								t.Logf("%s action called", MilestoneEventClosedAction)
								return nil
							},
						},
					},
					Tracer: noop.Tracer{},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "milestone",
				event:      &github.MilestoneEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail MilestoneEventClosed with nil action",
			fields: fields{
				handler: &EventHandler{
					WebhookSecret: "fake",
					onBeforeAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onBeforeAny called")
								return nil
							},
						},
					},
					onAfterAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onAfterAny called")
								return nil
							},
						},
					},
					onMilestoneEvent: map[string][]MilestoneEventHandleFunc{
						MilestoneEventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						MilestoneEventClosedAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
								t.Logf("%s action called", MilestoneEventClosedAction)
								return nil
							},
						},
					},
					Tracer: noop.Tracer{},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "milestone",
				event:      &github.MilestoneEvent{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger MilestoneEventOpened",
			fields: fields{
				handler: &EventHandler{
					WebhookSecret: "fake",
					onBeforeAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onBeforeAny called")
								return nil
							},
						},
					},
					onAfterAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onAfterAny called")
								return nil
							},
						},
					},
					onMilestoneEvent: map[string][]MilestoneEventHandleFunc{
						MilestoneEventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						MilestoneEventOpenedAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
								t.Logf("%s action called", MilestoneEventOpenedAction)
								return nil
							},
						},
					},
					Tracer: noop.Tracer{},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "milestone",
				event:      &github.MilestoneEvent{Action: ptrString(MilestoneEventOpenedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail MilestoneEventOpened with empty action",
			fields: fields{
				handler: &EventHandler{
					WebhookSecret: "fake",
					onBeforeAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onBeforeAny called")
								return nil
							},
						},
					},
					onAfterAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onAfterAny called")
								return nil
							},
						},
					},
					onMilestoneEvent: map[string][]MilestoneEventHandleFunc{
						MilestoneEventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						MilestoneEventOpenedAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
								t.Logf("%s action called", MilestoneEventOpenedAction)
								return nil
							},
						},
					},
					Tracer: noop.Tracer{},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "milestone",
				event:      &github.MilestoneEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail MilestoneEventOpened with nil action",
			fields: fields{
				handler: &EventHandler{
					WebhookSecret: "fake",
					onBeforeAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onBeforeAny called")
								return nil
							},
						},
					},
					onAfterAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onAfterAny called")
								return nil
							},
						},
					},
					onMilestoneEvent: map[string][]MilestoneEventHandleFunc{
						MilestoneEventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						MilestoneEventOpenedAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
								t.Logf("%s action called", MilestoneEventOpenedAction)
								return nil
							},
						},
					},
					Tracer: noop.Tracer{},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "milestone",
				event:      &github.MilestoneEvent{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger MilestoneEventEdited",
			fields: fields{
				handler: &EventHandler{
					WebhookSecret: "fake",
					onBeforeAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onBeforeAny called")
								return nil
							},
						},
					},
					onAfterAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onAfterAny called")
								return nil
							},
						},
					},
					onMilestoneEvent: map[string][]MilestoneEventHandleFunc{
						MilestoneEventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						MilestoneEventEditedAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
								t.Logf("%s action called", MilestoneEventEditedAction)
								return nil
							},
						},
					},
					Tracer: noop.Tracer{},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "milestone",
				event:      &github.MilestoneEvent{Action: ptrString(MilestoneEventEditedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail MilestoneEventEdited with empty action",
			fields: fields{
				handler: &EventHandler{
					WebhookSecret: "fake",
					onBeforeAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onBeforeAny called")
								return nil
							},
						},
					},
					onAfterAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onAfterAny called")
								return nil
							},
						},
					},
					onMilestoneEvent: map[string][]MilestoneEventHandleFunc{
						MilestoneEventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						MilestoneEventEditedAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
								t.Logf("%s action called", MilestoneEventEditedAction)
								return nil
							},
						},
					},
					Tracer: noop.Tracer{},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "milestone",
				event:      &github.MilestoneEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail MilestoneEventEdited with nil action",
			fields: fields{
				handler: &EventHandler{
					WebhookSecret: "fake",
					onBeforeAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onBeforeAny called")
								return nil
							},
						},
					},
					onAfterAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onAfterAny called")
								return nil
							},
						},
					},
					onMilestoneEvent: map[string][]MilestoneEventHandleFunc{
						MilestoneEventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						MilestoneEventEditedAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
								t.Logf("%s action called", MilestoneEventEditedAction)
								return nil
							},
						},
					},
					Tracer: noop.Tracer{},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "milestone",
				event:      &github.MilestoneEvent{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger MilestoneEventDeleted",
			fields: fields{
				handler: &EventHandler{
					WebhookSecret: "fake",
					onBeforeAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onBeforeAny called")
								return nil
							},
						},
					},
					onAfterAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onAfterAny called")
								return nil
							},
						},
					},
					onMilestoneEvent: map[string][]MilestoneEventHandleFunc{
						MilestoneEventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						MilestoneEventDeletedAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
								t.Logf("%s action called", MilestoneEventDeletedAction)
								return nil
							},
						},
					},
					Tracer: noop.Tracer{},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "milestone",
				event:      &github.MilestoneEvent{Action: ptrString(MilestoneEventDeletedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail MilestoneEventDeleted with empty action",
			fields: fields{
				handler: &EventHandler{
					WebhookSecret: "fake",
					onBeforeAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onBeforeAny called")
								return nil
							},
						},
					},
					onAfterAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onAfterAny called")
								return nil
							},
						},
					},
					onMilestoneEvent: map[string][]MilestoneEventHandleFunc{
						MilestoneEventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						MilestoneEventDeletedAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
								t.Logf("%s action called", MilestoneEventDeletedAction)
								return nil
							},
						},
					},
					Tracer: noop.Tracer{},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "milestone",
				event:      &github.MilestoneEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail MilestoneEventDeleted with nil action",
			fields: fields{
				handler: &EventHandler{
					WebhookSecret: "fake",
					onBeforeAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onBeforeAny called")
								return nil
							},
						},
					},
					onAfterAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onAfterAny called")
								return nil
							},
						},
					},
					onMilestoneEvent: map[string][]MilestoneEventHandleFunc{
						MilestoneEventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						MilestoneEventDeletedAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.MilestoneEvent) error {
								t.Logf("%s action called", MilestoneEventDeletedAction)
								return nil
							},
						},
					},
					Tracer: noop.Tracer{},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "milestone",
				event:      &github.MilestoneEvent{Action: nil},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &EventHandler{
				WebhookSecret: "fake",
				mu:            sync.RWMutex{},
				Tracer:        noop.Tracer{},
			}
			if err := g.MilestoneEvent(context.Background(), tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("MilestoneEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
