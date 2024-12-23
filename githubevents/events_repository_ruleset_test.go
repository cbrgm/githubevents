// Copyright 2022 The GithubEvents Authors. All rights reserved.
// Use of this source code is governed by the MIT License
// that can be found in the LICENSE file.

package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"errors"
	"github.com/google/go-github/v68/github"
	"sync"
	"testing"
)

func TestOnRepositoryRulesetEventAny(t *testing.T) {
	type args struct {
		callbacks []RepositoryRulesetEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single RepositoryRulesetEventHandleFunc",
			args: args{
				[]RepositoryRulesetEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryRulesetEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple RepositoryRulesetEventHandleFuncs",
			args: args{
				[]RepositoryRulesetEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryRulesetEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.RepositoryRulesetEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnRepositoryRulesetEventAny(tt.args.callbacks...)
			if len(g.onRepositoryRulesetEvent[RepositoryRulesetEventAnyAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onRepositoryRulesetEvent[RepositoryRulesetEventAnyAction]))
			}
		})
	}
}

func TestSetOnRepositoryRulesetEventAny(t *testing.T) {
	type args struct {
		callbacks []RepositoryRulesetEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single RepositoryRulesetEventHandleFunc",
			args: args{
				[]RepositoryRulesetEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryRulesetEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple RepositoryRulesetEventHandleFuncs",
			args: args{
				[]RepositoryRulesetEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryRulesetEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.RepositoryRulesetEvent) error {
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
			g.SetOnRepositoryRulesetEventAny(func(deliveryID string, eventName string, event *github.RepositoryRulesetEvent) error {
				return nil
			})
			g.SetOnRepositoryRulesetEventAny(tt.args.callbacks...)
			if len(g.onRepositoryRulesetEvent[RepositoryRulesetEventAnyAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onRepositoryRulesetEvent[RepositoryRulesetEventAnyAction]), tt.want)
			}
		})
	}
}

func TestHandleRepositoryRulesetEventAny(t *testing.T) {

	action := "*"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.RepositoryRulesetEvent
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
				eventName:  "repository_ruleset",

				event: &github.RepositoryRulesetEvent{Action: &action},

				fail: false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "repository_ruleset",

				event: &github.RepositoryRulesetEvent{Action: &action},

				fail: true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "repository_ruleset",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnRepositoryRulesetEventAny(func(deliveryID string, eventName string, event *github.RepositoryRulesetEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleRepositoryRulesetEventAny(tt.args.deliveryID, tt.args.deliveryID, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("TestHandleRepositoryRulesetEventAny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnRepositoryRulesetEventCreated(t *testing.T) {
	type args struct {
		callbacks []RepositoryRulesetEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single RepositoryRulesetEventHandleFunc",
			args: args{
				callbacks: []RepositoryRulesetEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryRulesetEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple RepositoryRulesetEventHandleFunc",
			args: args{
				callbacks: []RepositoryRulesetEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryRulesetEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.RepositoryRulesetEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnRepositoryRulesetEventCreated(tt.args.callbacks...)
			if len(g.onRepositoryRulesetEvent[RepositoryRulesetEventCreatedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onRepositoryRulesetEvent[RepositoryRulesetEventCreatedAction]))
			}
		})
	}
}

func TestSetOnRepositoryRulesetEventCreated(t *testing.T) {
	type args struct {
		callbacks []RepositoryRulesetEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single RepositoryRulesetEventHandleFunc",
			args: args{
				[]RepositoryRulesetEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryRulesetEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple RepositoryRulesetEventHandleFuncs",
			args: args{
				[]RepositoryRulesetEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryRulesetEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.RepositoryRulesetEvent) error {
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
			g.SetOnRepositoryRulesetEventCreated(func(deliveryID string, eventName string, event *github.RepositoryRulesetEvent) error {
				return nil
			})
			g.SetOnRepositoryRulesetEventCreated(tt.args.callbacks...)
			if len(g.onRepositoryRulesetEvent[RepositoryRulesetEventCreatedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onRepositoryRulesetEvent[RepositoryRulesetEventCreatedAction]), tt.want)
			}
		})
	}
}

func TestHandleRepositoryRulesetEventCreated(t *testing.T) {
	action := RepositoryRulesetEventCreatedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.RepositoryRulesetEvent
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
				eventName:  "repository_ruleset",
				event:      &github.RepositoryRulesetEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "repository_ruleset",
				event:      &github.RepositoryRulesetEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "repository_ruleset",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "repository_ruleset",
				event:      &github.RepositoryRulesetEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "repository_ruleset",
				event:      &github.RepositoryRulesetEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "repository_ruleset",
				event:      &github.RepositoryRulesetEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnRepositoryRulesetEventCreated(func(deliveryID string, eventName string, event *github.RepositoryRulesetEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleRepositoryRulesetEventCreated(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleRepositoryRulesetEventCreated() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnRepositoryRulesetEventDeleted(t *testing.T) {
	type args struct {
		callbacks []RepositoryRulesetEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single RepositoryRulesetEventHandleFunc",
			args: args{
				callbacks: []RepositoryRulesetEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryRulesetEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple RepositoryRulesetEventHandleFunc",
			args: args{
				callbacks: []RepositoryRulesetEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryRulesetEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.RepositoryRulesetEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnRepositoryRulesetEventDeleted(tt.args.callbacks...)
			if len(g.onRepositoryRulesetEvent[RepositoryRulesetEventDeletedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onRepositoryRulesetEvent[RepositoryRulesetEventDeletedAction]))
			}
		})
	}
}

func TestSetOnRepositoryRulesetEventDeleted(t *testing.T) {
	type args struct {
		callbacks []RepositoryRulesetEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single RepositoryRulesetEventHandleFunc",
			args: args{
				[]RepositoryRulesetEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryRulesetEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple RepositoryRulesetEventHandleFuncs",
			args: args{
				[]RepositoryRulesetEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryRulesetEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.RepositoryRulesetEvent) error {
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
			g.SetOnRepositoryRulesetEventDeleted(func(deliveryID string, eventName string, event *github.RepositoryRulesetEvent) error {
				return nil
			})
			g.SetOnRepositoryRulesetEventDeleted(tt.args.callbacks...)
			if len(g.onRepositoryRulesetEvent[RepositoryRulesetEventDeletedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onRepositoryRulesetEvent[RepositoryRulesetEventDeletedAction]), tt.want)
			}
		})
	}
}

func TestHandleRepositoryRulesetEventDeleted(t *testing.T) {
	action := RepositoryRulesetEventDeletedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.RepositoryRulesetEvent
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
				eventName:  "repository_ruleset",
				event:      &github.RepositoryRulesetEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "repository_ruleset",
				event:      &github.RepositoryRulesetEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "repository_ruleset",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "repository_ruleset",
				event:      &github.RepositoryRulesetEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "repository_ruleset",
				event:      &github.RepositoryRulesetEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "repository_ruleset",
				event:      &github.RepositoryRulesetEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnRepositoryRulesetEventDeleted(func(deliveryID string, eventName string, event *github.RepositoryRulesetEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleRepositoryRulesetEventDeleted(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleRepositoryRulesetEventDeleted() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnRepositoryRulesetEventEdited(t *testing.T) {
	type args struct {
		callbacks []RepositoryRulesetEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single RepositoryRulesetEventHandleFunc",
			args: args{
				callbacks: []RepositoryRulesetEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryRulesetEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple RepositoryRulesetEventHandleFunc",
			args: args{
				callbacks: []RepositoryRulesetEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryRulesetEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.RepositoryRulesetEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnRepositoryRulesetEventEdited(tt.args.callbacks...)
			if len(g.onRepositoryRulesetEvent[RepositoryRulesetEventEditedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onRepositoryRulesetEvent[RepositoryRulesetEventEditedAction]))
			}
		})
	}
}

func TestSetOnRepositoryRulesetEventEdited(t *testing.T) {
	type args struct {
		callbacks []RepositoryRulesetEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single RepositoryRulesetEventHandleFunc",
			args: args{
				[]RepositoryRulesetEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryRulesetEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple RepositoryRulesetEventHandleFuncs",
			args: args{
				[]RepositoryRulesetEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryRulesetEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.RepositoryRulesetEvent) error {
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
			g.SetOnRepositoryRulesetEventEdited(func(deliveryID string, eventName string, event *github.RepositoryRulesetEvent) error {
				return nil
			})
			g.SetOnRepositoryRulesetEventEdited(tt.args.callbacks...)
			if len(g.onRepositoryRulesetEvent[RepositoryRulesetEventEditedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onRepositoryRulesetEvent[RepositoryRulesetEventEditedAction]), tt.want)
			}
		})
	}
}

func TestHandleRepositoryRulesetEventEdited(t *testing.T) {
	action := RepositoryRulesetEventEditedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.RepositoryRulesetEvent
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
				eventName:  "repository_ruleset",
				event:      &github.RepositoryRulesetEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "repository_ruleset",
				event:      &github.RepositoryRulesetEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "repository_ruleset",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "repository_ruleset",
				event:      &github.RepositoryRulesetEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "repository_ruleset",
				event:      &github.RepositoryRulesetEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "repository_ruleset",
				event:      &github.RepositoryRulesetEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnRepositoryRulesetEventEdited(func(deliveryID string, eventName string, event *github.RepositoryRulesetEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleRepositoryRulesetEventEdited(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleRepositoryRulesetEventEdited() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRepositoryRulesetEvent(t *testing.T) {
	type fields struct {
		handler *EventHandler
	}
	type args struct {
		deliveryID string
		eventName  string
		event      *github.RepositoryRulesetEvent
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "must trigger RepositoryRulesetEventAny with unknown event action",
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
					onRepositoryRulesetEvent: map[string][]RepositoryRulesetEventHandleFunc{
						RepositoryRulesetEventAnyAction: {
							func(deliveryID string, eventName string, event *github.RepositoryRulesetEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  RepositoryRulesetEvent,

				event: &github.RepositoryRulesetEvent{Action: ptrString("unknown")},
			},
			wantErr: false,
		},

		{
			name: "must trigger RepositoryRulesetEventCreated",
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
					onRepositoryRulesetEvent: map[string][]RepositoryRulesetEventHandleFunc{
						RepositoryRulesetEventAnyAction: {
							func(deliveryID string, eventName string, event *github.RepositoryRulesetEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						RepositoryRulesetEventCreatedAction: {
							func(deliveryID string, eventName string, event *github.RepositoryRulesetEvent) error {
								t.Logf("%s action called", RepositoryRulesetEventCreatedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "repository_ruleset",
				event:      &github.RepositoryRulesetEvent{Action: ptrString(RepositoryRulesetEventCreatedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail RepositoryRulesetEventCreated with empty action",
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
					onRepositoryRulesetEvent: map[string][]RepositoryRulesetEventHandleFunc{
						RepositoryRulesetEventAnyAction: {
							func(deliveryID string, eventName string, event *github.RepositoryRulesetEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						RepositoryRulesetEventCreatedAction: {
							func(deliveryID string, eventName string, event *github.RepositoryRulesetEvent) error {
								t.Logf("%s action called", RepositoryRulesetEventCreatedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "repository_ruleset",
				event:      &github.RepositoryRulesetEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail RepositoryRulesetEventCreated with nil action",
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
					onRepositoryRulesetEvent: map[string][]RepositoryRulesetEventHandleFunc{
						RepositoryRulesetEventAnyAction: {
							func(deliveryID string, eventName string, event *github.RepositoryRulesetEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						RepositoryRulesetEventCreatedAction: {
							func(deliveryID string, eventName string, event *github.RepositoryRulesetEvent) error {
								t.Logf("%s action called", RepositoryRulesetEventCreatedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "repository_ruleset",
				event:      &github.RepositoryRulesetEvent{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger RepositoryRulesetEventDeleted",
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
					onRepositoryRulesetEvent: map[string][]RepositoryRulesetEventHandleFunc{
						RepositoryRulesetEventAnyAction: {
							func(deliveryID string, eventName string, event *github.RepositoryRulesetEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						RepositoryRulesetEventDeletedAction: {
							func(deliveryID string, eventName string, event *github.RepositoryRulesetEvent) error {
								t.Logf("%s action called", RepositoryRulesetEventDeletedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "repository_ruleset",
				event:      &github.RepositoryRulesetEvent{Action: ptrString(RepositoryRulesetEventDeletedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail RepositoryRulesetEventDeleted with empty action",
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
					onRepositoryRulesetEvent: map[string][]RepositoryRulesetEventHandleFunc{
						RepositoryRulesetEventAnyAction: {
							func(deliveryID string, eventName string, event *github.RepositoryRulesetEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						RepositoryRulesetEventDeletedAction: {
							func(deliveryID string, eventName string, event *github.RepositoryRulesetEvent) error {
								t.Logf("%s action called", RepositoryRulesetEventDeletedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "repository_ruleset",
				event:      &github.RepositoryRulesetEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail RepositoryRulesetEventDeleted with nil action",
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
					onRepositoryRulesetEvent: map[string][]RepositoryRulesetEventHandleFunc{
						RepositoryRulesetEventAnyAction: {
							func(deliveryID string, eventName string, event *github.RepositoryRulesetEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						RepositoryRulesetEventDeletedAction: {
							func(deliveryID string, eventName string, event *github.RepositoryRulesetEvent) error {
								t.Logf("%s action called", RepositoryRulesetEventDeletedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "repository_ruleset",
				event:      &github.RepositoryRulesetEvent{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger RepositoryRulesetEventEdited",
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
					onRepositoryRulesetEvent: map[string][]RepositoryRulesetEventHandleFunc{
						RepositoryRulesetEventAnyAction: {
							func(deliveryID string, eventName string, event *github.RepositoryRulesetEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						RepositoryRulesetEventEditedAction: {
							func(deliveryID string, eventName string, event *github.RepositoryRulesetEvent) error {
								t.Logf("%s action called", RepositoryRulesetEventEditedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "repository_ruleset",
				event:      &github.RepositoryRulesetEvent{Action: ptrString(RepositoryRulesetEventEditedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail RepositoryRulesetEventEdited with empty action",
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
					onRepositoryRulesetEvent: map[string][]RepositoryRulesetEventHandleFunc{
						RepositoryRulesetEventAnyAction: {
							func(deliveryID string, eventName string, event *github.RepositoryRulesetEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						RepositoryRulesetEventEditedAction: {
							func(deliveryID string, eventName string, event *github.RepositoryRulesetEvent) error {
								t.Logf("%s action called", RepositoryRulesetEventEditedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "repository_ruleset",
				event:      &github.RepositoryRulesetEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail RepositoryRulesetEventEdited with nil action",
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
					onRepositoryRulesetEvent: map[string][]RepositoryRulesetEventHandleFunc{
						RepositoryRulesetEventAnyAction: {
							func(deliveryID string, eventName string, event *github.RepositoryRulesetEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						RepositoryRulesetEventEditedAction: {
							func(deliveryID string, eventName string, event *github.RepositoryRulesetEvent) error {
								t.Logf("%s action called", RepositoryRulesetEventEditedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "repository_ruleset",
				event:      &github.RepositoryRulesetEvent{Action: nil},
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
			if err := g.RepositoryRulesetEvent(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("RepositoryRulesetEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
