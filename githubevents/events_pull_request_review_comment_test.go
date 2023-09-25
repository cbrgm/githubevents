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

func TestOnPullRequestReviewCommentEventAny(t *testing.T) {
	type args struct {
		callbacks []PullRequestReviewCommentEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single PullRequestReviewCommentEventHandleFunc",
			args: args{
				[]PullRequestReviewCommentEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple PullRequestReviewCommentEventHandleFuncs",
			args: args{
				[]PullRequestReviewCommentEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnPullRequestReviewCommentEventAny(tt.args.callbacks...)
			if len(g.onPullRequestReviewCommentEvent[PullRequestReviewCommentEventAnyAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onPullRequestReviewCommentEvent[PullRequestReviewCommentEventAnyAction]))
			}
		})
	}
}

func TestSetOnPullRequestReviewCommentEventAny(t *testing.T) {
	type args struct {
		callbacks []PullRequestReviewCommentEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single PullRequestReviewCommentEventHandleFunc",
			args: args{
				[]PullRequestReviewCommentEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple PullRequestReviewCommentEventHandleFuncs",
			args: args{
				[]PullRequestReviewCommentEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
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
			g.SetOnPullRequestReviewCommentEventAny(func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
				return nil
			})
			g.SetOnPullRequestReviewCommentEventAny(tt.args.callbacks...)
			if len(g.onPullRequestReviewCommentEvent[PullRequestReviewCommentEventAnyAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onPullRequestReviewCommentEvent[PullRequestReviewCommentEventAnyAction]), tt.want)
			}
		})
	}
}

func TestHandlePullRequestReviewCommentEventAny(t *testing.T) {

	action := "*"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.PullRequestReviewCommentEvent
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
				eventName:  "pull_request_review_comment",

				event: &github.PullRequestReviewCommentEvent{Action: &action},

				fail: false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request_review_comment",

				event: &github.PullRequestReviewCommentEvent{Action: &action},

				fail: true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request_review_comment",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnPullRequestReviewCommentEventAny(func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handlePullRequestReviewCommentEventAny(tt.args.deliveryID, tt.args.deliveryID, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("TestHandlePullRequestReviewCommentEventAny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnPullRequestReviewCommentEventCreated(t *testing.T) {
	type args struct {
		callbacks []PullRequestReviewCommentEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single PullRequestReviewCommentEventHandleFunc",
			args: args{
				callbacks: []PullRequestReviewCommentEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple PullRequestReviewCommentEventHandleFunc",
			args: args{
				callbacks: []PullRequestReviewCommentEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnPullRequestReviewCommentEventCreated(tt.args.callbacks...)
			if len(g.onPullRequestReviewCommentEvent[PullRequestReviewCommentEventCreatedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onPullRequestReviewCommentEvent[PullRequestReviewCommentEventCreatedAction]))
			}
		})
	}
}

func TestSetOnPullRequestReviewCommentEventCreated(t *testing.T) {
	type args struct {
		callbacks []PullRequestReviewCommentEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single PullRequestReviewCommentEventHandleFunc",
			args: args{
				[]PullRequestReviewCommentEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple PullRequestReviewCommentEventHandleFuncs",
			args: args{
				[]PullRequestReviewCommentEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
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
			g.SetOnPullRequestReviewCommentEventCreated(func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
				return nil
			})
			g.SetOnPullRequestReviewCommentEventCreated(tt.args.callbacks...)
			if len(g.onPullRequestReviewCommentEvent[PullRequestReviewCommentEventCreatedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onPullRequestReviewCommentEvent[PullRequestReviewCommentEventCreatedAction]), tt.want)
			}
		})
	}
}

func TestHandlePullRequestReviewCommentEventCreated(t *testing.T) {
	action := PullRequestReviewCommentEventCreatedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.PullRequestReviewCommentEvent
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
				eventName:  "pull_request_review_comment",
				event:      &github.PullRequestReviewCommentEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request_review_comment",
				event:      &github.PullRequestReviewCommentEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request_review_comment",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request_review_comment",
				event:      &github.PullRequestReviewCommentEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request_review_comment",
				event:      &github.PullRequestReviewCommentEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request_review_comment",
				event:      &github.PullRequestReviewCommentEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnPullRequestReviewCommentEventCreated(func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handlePullRequestReviewCommentEventCreated(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handlePullRequestReviewCommentEventCreated() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnPullRequestReviewCommentEventEdited(t *testing.T) {
	type args struct {
		callbacks []PullRequestReviewCommentEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single PullRequestReviewCommentEventHandleFunc",
			args: args{
				callbacks: []PullRequestReviewCommentEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple PullRequestReviewCommentEventHandleFunc",
			args: args{
				callbacks: []PullRequestReviewCommentEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnPullRequestReviewCommentEventEdited(tt.args.callbacks...)
			if len(g.onPullRequestReviewCommentEvent[PullRequestReviewCommentEventEditedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onPullRequestReviewCommentEvent[PullRequestReviewCommentEventEditedAction]))
			}
		})
	}
}

func TestSetOnPullRequestReviewCommentEventEdited(t *testing.T) {
	type args struct {
		callbacks []PullRequestReviewCommentEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single PullRequestReviewCommentEventHandleFunc",
			args: args{
				[]PullRequestReviewCommentEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple PullRequestReviewCommentEventHandleFuncs",
			args: args{
				[]PullRequestReviewCommentEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
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
			g.SetOnPullRequestReviewCommentEventEdited(func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
				return nil
			})
			g.SetOnPullRequestReviewCommentEventEdited(tt.args.callbacks...)
			if len(g.onPullRequestReviewCommentEvent[PullRequestReviewCommentEventEditedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onPullRequestReviewCommentEvent[PullRequestReviewCommentEventEditedAction]), tt.want)
			}
		})
	}
}

func TestHandlePullRequestReviewCommentEventEdited(t *testing.T) {
	action := PullRequestReviewCommentEventEditedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.PullRequestReviewCommentEvent
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
				eventName:  "pull_request_review_comment",
				event:      &github.PullRequestReviewCommentEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request_review_comment",
				event:      &github.PullRequestReviewCommentEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request_review_comment",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request_review_comment",
				event:      &github.PullRequestReviewCommentEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request_review_comment",
				event:      &github.PullRequestReviewCommentEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request_review_comment",
				event:      &github.PullRequestReviewCommentEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnPullRequestReviewCommentEventEdited(func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handlePullRequestReviewCommentEventEdited(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handlePullRequestReviewCommentEventEdited() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnPullRequestReviewCommentEventDeleted(t *testing.T) {
	type args struct {
		callbacks []PullRequestReviewCommentEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single PullRequestReviewCommentEventHandleFunc",
			args: args{
				callbacks: []PullRequestReviewCommentEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple PullRequestReviewCommentEventHandleFunc",
			args: args{
				callbacks: []PullRequestReviewCommentEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnPullRequestReviewCommentEventDeleted(tt.args.callbacks...)
			if len(g.onPullRequestReviewCommentEvent[PullRequestReviewCommentEventDeletedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onPullRequestReviewCommentEvent[PullRequestReviewCommentEventDeletedAction]))
			}
		})
	}
}

func TestSetOnPullRequestReviewCommentEventDeleted(t *testing.T) {
	type args struct {
		callbacks []PullRequestReviewCommentEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single PullRequestReviewCommentEventHandleFunc",
			args: args{
				[]PullRequestReviewCommentEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple PullRequestReviewCommentEventHandleFuncs",
			args: args{
				[]PullRequestReviewCommentEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
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
			g.SetOnPullRequestReviewCommentEventDeleted(func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
				return nil
			})
			g.SetOnPullRequestReviewCommentEventDeleted(tt.args.callbacks...)
			if len(g.onPullRequestReviewCommentEvent[PullRequestReviewCommentEventDeletedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onPullRequestReviewCommentEvent[PullRequestReviewCommentEventDeletedAction]), tt.want)
			}
		})
	}
}

func TestHandlePullRequestReviewCommentEventDeleted(t *testing.T) {
	action := PullRequestReviewCommentEventDeletedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.PullRequestReviewCommentEvent
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
				eventName:  "pull_request_review_comment",
				event:      &github.PullRequestReviewCommentEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request_review_comment",
				event:      &github.PullRequestReviewCommentEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request_review_comment",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request_review_comment",
				event:      &github.PullRequestReviewCommentEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request_review_comment",
				event:      &github.PullRequestReviewCommentEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request_review_comment",
				event:      &github.PullRequestReviewCommentEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnPullRequestReviewCommentEventDeleted(func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handlePullRequestReviewCommentEventDeleted(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handlePullRequestReviewCommentEventDeleted() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPullRequestReviewCommentEvent(t *testing.T) {
	type fields struct {
		handler *EventHandler
	}
	type args struct {
		deliveryID string
		eventName  string
		event      *github.PullRequestReviewCommentEvent
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "must trigger PullRequestReviewCommentEventAny with unknown event action",
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
					onPullRequestReviewCommentEvent: map[string][]PullRequestReviewCommentEventHandleFunc{
						PullRequestReviewCommentEventAnyAction: {
							func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  PullRequestReviewCommentEvent,

				event: &github.PullRequestReviewCommentEvent{Action: ptrString("unknown")},
			},
			wantErr: false,
		},

		{
			name: "must trigger PullRequestReviewCommentEventCreated",
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
					onPullRequestReviewCommentEvent: map[string][]PullRequestReviewCommentEventHandleFunc{
						PullRequestReviewCommentEventAnyAction: {
							func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						PullRequestReviewCommentEventCreatedAction: {
							func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
								t.Logf("%s action called", PullRequestReviewCommentEventCreatedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "pull_request_review_comment",
				event:      &github.PullRequestReviewCommentEvent{Action: ptrString(PullRequestReviewCommentEventCreatedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail PullRequestReviewCommentEventCreated with empty action",
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
					onPullRequestReviewCommentEvent: map[string][]PullRequestReviewCommentEventHandleFunc{
						PullRequestReviewCommentEventAnyAction: {
							func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						PullRequestReviewCommentEventCreatedAction: {
							func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
								t.Logf("%s action called", PullRequestReviewCommentEventCreatedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "pull_request_review_comment",
				event:      &github.PullRequestReviewCommentEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail PullRequestReviewCommentEventCreated with nil action",
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
					onPullRequestReviewCommentEvent: map[string][]PullRequestReviewCommentEventHandleFunc{
						PullRequestReviewCommentEventAnyAction: {
							func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						PullRequestReviewCommentEventCreatedAction: {
							func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
								t.Logf("%s action called", PullRequestReviewCommentEventCreatedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "pull_request_review_comment",
				event:      &github.PullRequestReviewCommentEvent{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger PullRequestReviewCommentEventEdited",
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
					onPullRequestReviewCommentEvent: map[string][]PullRequestReviewCommentEventHandleFunc{
						PullRequestReviewCommentEventAnyAction: {
							func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						PullRequestReviewCommentEventEditedAction: {
							func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
								t.Logf("%s action called", PullRequestReviewCommentEventEditedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "pull_request_review_comment",
				event:      &github.PullRequestReviewCommentEvent{Action: ptrString(PullRequestReviewCommentEventEditedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail PullRequestReviewCommentEventEdited with empty action",
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
					onPullRequestReviewCommentEvent: map[string][]PullRequestReviewCommentEventHandleFunc{
						PullRequestReviewCommentEventAnyAction: {
							func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						PullRequestReviewCommentEventEditedAction: {
							func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
								t.Logf("%s action called", PullRequestReviewCommentEventEditedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "pull_request_review_comment",
				event:      &github.PullRequestReviewCommentEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail PullRequestReviewCommentEventEdited with nil action",
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
					onPullRequestReviewCommentEvent: map[string][]PullRequestReviewCommentEventHandleFunc{
						PullRequestReviewCommentEventAnyAction: {
							func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						PullRequestReviewCommentEventEditedAction: {
							func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
								t.Logf("%s action called", PullRequestReviewCommentEventEditedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "pull_request_review_comment",
				event:      &github.PullRequestReviewCommentEvent{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger PullRequestReviewCommentEventDeleted",
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
					onPullRequestReviewCommentEvent: map[string][]PullRequestReviewCommentEventHandleFunc{
						PullRequestReviewCommentEventAnyAction: {
							func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						PullRequestReviewCommentEventDeletedAction: {
							func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
								t.Logf("%s action called", PullRequestReviewCommentEventDeletedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "pull_request_review_comment",
				event:      &github.PullRequestReviewCommentEvent{Action: ptrString(PullRequestReviewCommentEventDeletedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail PullRequestReviewCommentEventDeleted with empty action",
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
					onPullRequestReviewCommentEvent: map[string][]PullRequestReviewCommentEventHandleFunc{
						PullRequestReviewCommentEventAnyAction: {
							func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						PullRequestReviewCommentEventDeletedAction: {
							func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
								t.Logf("%s action called", PullRequestReviewCommentEventDeletedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "pull_request_review_comment",
				event:      &github.PullRequestReviewCommentEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail PullRequestReviewCommentEventDeleted with nil action",
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
					onPullRequestReviewCommentEvent: map[string][]PullRequestReviewCommentEventHandleFunc{
						PullRequestReviewCommentEventAnyAction: {
							func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						PullRequestReviewCommentEventDeletedAction: {
							func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
								t.Logf("%s action called", PullRequestReviewCommentEventDeletedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "pull_request_review_comment",
				event:      &github.PullRequestReviewCommentEvent{Action: nil},
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
			if err := g.PullRequestReviewCommentEvent(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("PullRequestReviewCommentEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
