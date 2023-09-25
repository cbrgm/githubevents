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

func TestOnIssueCommentEventAny(t *testing.T) {
	type args struct {
		callbacks []IssueCommentEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single IssueCommentEventHandleFunc",
			args: args{
				[]IssueCommentEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssueCommentEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple IssueCommentEventHandleFuncs",
			args: args{
				[]IssueCommentEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssueCommentEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.IssueCommentEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnIssueCommentEventAny(tt.args.callbacks...)
			if len(g.onIssueCommentEvent[IssueCommentEventAnyAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onIssueCommentEvent[IssueCommentEventAnyAction]))
			}
		})
	}
}

func TestSetOnIssueCommentEventAny(t *testing.T) {
	type args struct {
		callbacks []IssueCommentEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single IssueCommentEventHandleFunc",
			args: args{
				[]IssueCommentEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssueCommentEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple IssueCommentEventHandleFuncs",
			args: args{
				[]IssueCommentEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssueCommentEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.IssueCommentEvent) error {
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
			g.SetOnIssueCommentEventAny(func(deliveryID string, eventName string, event *github.IssueCommentEvent) error {
				return nil
			})
			g.SetOnIssueCommentEventAny(tt.args.callbacks...)
			if len(g.onIssueCommentEvent[IssueCommentEventAnyAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onIssueCommentEvent[IssueCommentEventAnyAction]), tt.want)
			}
		})
	}
}

func TestHandleIssueCommentEventAny(t *testing.T) {

	action := "*"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.IssueCommentEvent
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
				eventName:  "issue_comment",

				event: &github.IssueCommentEvent{Action: &action},

				fail: false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "issue_comment",

				event: &github.IssueCommentEvent{Action: &action},

				fail: true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "issue_comment",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnIssueCommentEventAny(func(deliveryID string, eventName string, event *github.IssueCommentEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleIssueCommentEventAny(tt.args.deliveryID, tt.args.deliveryID, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("TestHandleIssueCommentEventAny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnIssueCommentCreated(t *testing.T) {
	type args struct {
		callbacks []IssueCommentEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single IssueCommentEventHandleFunc",
			args: args{
				callbacks: []IssueCommentEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssueCommentEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple IssueCommentEventHandleFunc",
			args: args{
				callbacks: []IssueCommentEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssueCommentEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.IssueCommentEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnIssueCommentCreated(tt.args.callbacks...)
			if len(g.onIssueCommentEvent[IssueCommentCreatedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onIssueCommentEvent[IssueCommentCreatedAction]))
			}
		})
	}
}

func TestSetOnIssueCommentCreated(t *testing.T) {
	type args struct {
		callbacks []IssueCommentEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single IssueCommentEventHandleFunc",
			args: args{
				[]IssueCommentEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssueCommentEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple IssueCommentEventHandleFuncs",
			args: args{
				[]IssueCommentEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssueCommentEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.IssueCommentEvent) error {
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
			g.SetOnIssueCommentCreated(func(deliveryID string, eventName string, event *github.IssueCommentEvent) error {
				return nil
			})
			g.SetOnIssueCommentCreated(tt.args.callbacks...)
			if len(g.onIssueCommentEvent[IssueCommentCreatedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onIssueCommentEvent[IssueCommentCreatedAction]), tt.want)
			}
		})
	}
}

func TestHandleIssueCommentCreated(t *testing.T) {
	action := IssueCommentCreatedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.IssueCommentEvent
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
				eventName:  "issue_comment",
				event:      &github.IssueCommentEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "issue_comment",
				event:      &github.IssueCommentEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "issue_comment",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "issue_comment",
				event:      &github.IssueCommentEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "issue_comment",
				event:      &github.IssueCommentEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "issue_comment",
				event:      &github.IssueCommentEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnIssueCommentCreated(func(deliveryID string, eventName string, event *github.IssueCommentEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleIssueCommentCreated(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleIssueCommentCreated() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnIssueCommentEdited(t *testing.T) {
	type args struct {
		callbacks []IssueCommentEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single IssueCommentEventHandleFunc",
			args: args{
				callbacks: []IssueCommentEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssueCommentEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple IssueCommentEventHandleFunc",
			args: args{
				callbacks: []IssueCommentEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssueCommentEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.IssueCommentEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnIssueCommentEdited(tt.args.callbacks...)
			if len(g.onIssueCommentEvent[IssueCommentEditedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onIssueCommentEvent[IssueCommentEditedAction]))
			}
		})
	}
}

func TestSetOnIssueCommentEdited(t *testing.T) {
	type args struct {
		callbacks []IssueCommentEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single IssueCommentEventHandleFunc",
			args: args{
				[]IssueCommentEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssueCommentEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple IssueCommentEventHandleFuncs",
			args: args{
				[]IssueCommentEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssueCommentEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.IssueCommentEvent) error {
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
			g.SetOnIssueCommentEdited(func(deliveryID string, eventName string, event *github.IssueCommentEvent) error {
				return nil
			})
			g.SetOnIssueCommentEdited(tt.args.callbacks...)
			if len(g.onIssueCommentEvent[IssueCommentEditedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onIssueCommentEvent[IssueCommentEditedAction]), tt.want)
			}
		})
	}
}

func TestHandleIssueCommentEdited(t *testing.T) {
	action := IssueCommentEditedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.IssueCommentEvent
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
				eventName:  "issue_comment",
				event:      &github.IssueCommentEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "issue_comment",
				event:      &github.IssueCommentEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "issue_comment",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "issue_comment",
				event:      &github.IssueCommentEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "issue_comment",
				event:      &github.IssueCommentEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "issue_comment",
				event:      &github.IssueCommentEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnIssueCommentEdited(func(deliveryID string, eventName string, event *github.IssueCommentEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleIssueCommentEdited(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleIssueCommentEdited() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnIssueCommentDeleted(t *testing.T) {
	type args struct {
		callbacks []IssueCommentEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single IssueCommentEventHandleFunc",
			args: args{
				callbacks: []IssueCommentEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssueCommentEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple IssueCommentEventHandleFunc",
			args: args{
				callbacks: []IssueCommentEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssueCommentEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.IssueCommentEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnIssueCommentDeleted(tt.args.callbacks...)
			if len(g.onIssueCommentEvent[IssueCommentDeletedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onIssueCommentEvent[IssueCommentDeletedAction]))
			}
		})
	}
}

func TestSetOnIssueCommentDeleted(t *testing.T) {
	type args struct {
		callbacks []IssueCommentEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single IssueCommentEventHandleFunc",
			args: args{
				[]IssueCommentEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssueCommentEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple IssueCommentEventHandleFuncs",
			args: args{
				[]IssueCommentEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssueCommentEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.IssueCommentEvent) error {
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
			g.SetOnIssueCommentDeleted(func(deliveryID string, eventName string, event *github.IssueCommentEvent) error {
				return nil
			})
			g.SetOnIssueCommentDeleted(tt.args.callbacks...)
			if len(g.onIssueCommentEvent[IssueCommentDeletedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onIssueCommentEvent[IssueCommentDeletedAction]), tt.want)
			}
		})
	}
}

func TestHandleIssueCommentDeleted(t *testing.T) {
	action := IssueCommentDeletedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.IssueCommentEvent
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
				eventName:  "issue_comment",
				event:      &github.IssueCommentEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "issue_comment",
				event:      &github.IssueCommentEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "issue_comment",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "issue_comment",
				event:      &github.IssueCommentEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "issue_comment",
				event:      &github.IssueCommentEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "issue_comment",
				event:      &github.IssueCommentEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnIssueCommentDeleted(func(deliveryID string, eventName string, event *github.IssueCommentEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleIssueCommentDeleted(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleIssueCommentDeleted() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIssueCommentEvent(t *testing.T) {
	type fields struct {
		handler *EventHandler
	}
	type args struct {
		deliveryID string
		eventName  string
		event      *github.IssueCommentEvent
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "must trigger IssueCommentEventAny with unknown event action",
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
					onIssueCommentEvent: map[string][]IssueCommentEventHandleFunc{
						IssueCommentEventAnyAction: {
							func(deliveryID string, eventName string, event *github.IssueCommentEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  IssueCommentEvent,

				event: &github.IssueCommentEvent{Action: ptrString("unknown")},
			},
			wantErr: false,
		},

		{
			name: "must trigger IssueCommentCreated",
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
					onIssueCommentEvent: map[string][]IssueCommentEventHandleFunc{
						IssueCommentEventAnyAction: {
							func(deliveryID string, eventName string, event *github.IssueCommentEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						IssueCommentCreatedAction: {
							func(deliveryID string, eventName string, event *github.IssueCommentEvent) error {
								t.Logf("%s action called", IssueCommentCreatedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "issue_comment",
				event:      &github.IssueCommentEvent{Action: ptrString(IssueCommentCreatedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail IssueCommentCreated with empty action",
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
					onIssueCommentEvent: map[string][]IssueCommentEventHandleFunc{
						IssueCommentEventAnyAction: {
							func(deliveryID string, eventName string, event *github.IssueCommentEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						IssueCommentCreatedAction: {
							func(deliveryID string, eventName string, event *github.IssueCommentEvent) error {
								t.Logf("%s action called", IssueCommentCreatedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "issue_comment",
				event:      &github.IssueCommentEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail IssueCommentCreated with nil action",
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
					onIssueCommentEvent: map[string][]IssueCommentEventHandleFunc{
						IssueCommentEventAnyAction: {
							func(deliveryID string, eventName string, event *github.IssueCommentEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						IssueCommentCreatedAction: {
							func(deliveryID string, eventName string, event *github.IssueCommentEvent) error {
								t.Logf("%s action called", IssueCommentCreatedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "issue_comment",
				event:      &github.IssueCommentEvent{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger IssueCommentEdited",
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
					onIssueCommentEvent: map[string][]IssueCommentEventHandleFunc{
						IssueCommentEventAnyAction: {
							func(deliveryID string, eventName string, event *github.IssueCommentEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						IssueCommentEditedAction: {
							func(deliveryID string, eventName string, event *github.IssueCommentEvent) error {
								t.Logf("%s action called", IssueCommentEditedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "issue_comment",
				event:      &github.IssueCommentEvent{Action: ptrString(IssueCommentEditedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail IssueCommentEdited with empty action",
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
					onIssueCommentEvent: map[string][]IssueCommentEventHandleFunc{
						IssueCommentEventAnyAction: {
							func(deliveryID string, eventName string, event *github.IssueCommentEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						IssueCommentEditedAction: {
							func(deliveryID string, eventName string, event *github.IssueCommentEvent) error {
								t.Logf("%s action called", IssueCommentEditedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "issue_comment",
				event:      &github.IssueCommentEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail IssueCommentEdited with nil action",
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
					onIssueCommentEvent: map[string][]IssueCommentEventHandleFunc{
						IssueCommentEventAnyAction: {
							func(deliveryID string, eventName string, event *github.IssueCommentEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						IssueCommentEditedAction: {
							func(deliveryID string, eventName string, event *github.IssueCommentEvent) error {
								t.Logf("%s action called", IssueCommentEditedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "issue_comment",
				event:      &github.IssueCommentEvent{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger IssueCommentDeleted",
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
					onIssueCommentEvent: map[string][]IssueCommentEventHandleFunc{
						IssueCommentEventAnyAction: {
							func(deliveryID string, eventName string, event *github.IssueCommentEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						IssueCommentDeletedAction: {
							func(deliveryID string, eventName string, event *github.IssueCommentEvent) error {
								t.Logf("%s action called", IssueCommentDeletedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "issue_comment",
				event:      &github.IssueCommentEvent{Action: ptrString(IssueCommentDeletedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail IssueCommentDeleted with empty action",
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
					onIssueCommentEvent: map[string][]IssueCommentEventHandleFunc{
						IssueCommentEventAnyAction: {
							func(deliveryID string, eventName string, event *github.IssueCommentEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						IssueCommentDeletedAction: {
							func(deliveryID string, eventName string, event *github.IssueCommentEvent) error {
								t.Logf("%s action called", IssueCommentDeletedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "issue_comment",
				event:      &github.IssueCommentEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail IssueCommentDeleted with nil action",
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
					onIssueCommentEvent: map[string][]IssueCommentEventHandleFunc{
						IssueCommentEventAnyAction: {
							func(deliveryID string, eventName string, event *github.IssueCommentEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						IssueCommentDeletedAction: {
							func(deliveryID string, eventName string, event *github.IssueCommentEvent) error {
								t.Logf("%s action called", IssueCommentDeletedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "issue_comment",
				event:      &github.IssueCommentEvent{Action: nil},
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
			if err := g.IssueCommentEvent(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("IssueCommentEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
