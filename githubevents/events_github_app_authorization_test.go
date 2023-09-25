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

func TestOnGitHubAppAuthorizationEventAny(t *testing.T) {
	type args struct {
		callbacks []GitHubAppAuthorizationEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single GitHubAppAuthorizationEventHandleFunc",
			args: args{
				[]GitHubAppAuthorizationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.GitHubAppAuthorizationEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple GitHubAppAuthorizationEventHandleFuncs",
			args: args{
				[]GitHubAppAuthorizationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.GitHubAppAuthorizationEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.GitHubAppAuthorizationEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnGitHubAppAuthorizationEventAny(tt.args.callbacks...)
			if len(g.onGitHubAppAuthorizationEvent[GitHubAppAuthorizationEventAnyAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onGitHubAppAuthorizationEvent[GitHubAppAuthorizationEventAnyAction]))
			}
		})
	}
}

func TestSetOnGitHubAppAuthorizationEventAny(t *testing.T) {
	type args struct {
		callbacks []GitHubAppAuthorizationEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single GitHubAppAuthorizationEventHandleFunc",
			args: args{
				[]GitHubAppAuthorizationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.GitHubAppAuthorizationEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple GitHubAppAuthorizationEventHandleFuncs",
			args: args{
				[]GitHubAppAuthorizationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.GitHubAppAuthorizationEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.GitHubAppAuthorizationEvent) error {
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
			g.SetOnGitHubAppAuthorizationEventAny(func(deliveryID string, eventName string, event *github.GitHubAppAuthorizationEvent) error {
				return nil
			})
			g.SetOnGitHubAppAuthorizationEventAny(tt.args.callbacks...)
			if len(g.onGitHubAppAuthorizationEvent[GitHubAppAuthorizationEventAnyAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onGitHubAppAuthorizationEvent[GitHubAppAuthorizationEventAnyAction]), tt.want)
			}
		})
	}
}

func TestHandleGitHubAppAuthorizationEventAny(t *testing.T) {

	action := "*"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.GitHubAppAuthorizationEvent
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
				eventName:  "github_app_authorization",

				event: &github.GitHubAppAuthorizationEvent{Action: &action},

				fail: false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "github_app_authorization",

				event: &github.GitHubAppAuthorizationEvent{Action: &action},

				fail: true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "github_app_authorization",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnGitHubAppAuthorizationEventAny(func(deliveryID string, eventName string, event *github.GitHubAppAuthorizationEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleGitHubAppAuthorizationEventAny(tt.args.deliveryID, tt.args.deliveryID, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("TestHandleGitHubAppAuthorizationEventAny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnGitHubAppAuthorizationEventRevoked(t *testing.T) {
	type args struct {
		callbacks []GitHubAppAuthorizationEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single GitHubAppAuthorizationEventHandleFunc",
			args: args{
				callbacks: []GitHubAppAuthorizationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.GitHubAppAuthorizationEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple GitHubAppAuthorizationEventHandleFunc",
			args: args{
				callbacks: []GitHubAppAuthorizationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.GitHubAppAuthorizationEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.GitHubAppAuthorizationEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnGitHubAppAuthorizationEventRevoked(tt.args.callbacks...)
			if len(g.onGitHubAppAuthorizationEvent[GitHubAppAuthorizationEventRevokedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onGitHubAppAuthorizationEvent[GitHubAppAuthorizationEventRevokedAction]))
			}
		})
	}
}

func TestSetOnGitHubAppAuthorizationEventRevoked(t *testing.T) {
	type args struct {
		callbacks []GitHubAppAuthorizationEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single GitHubAppAuthorizationEventHandleFunc",
			args: args{
				[]GitHubAppAuthorizationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.GitHubAppAuthorizationEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple GitHubAppAuthorizationEventHandleFuncs",
			args: args{
				[]GitHubAppAuthorizationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.GitHubAppAuthorizationEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.GitHubAppAuthorizationEvent) error {
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
			g.SetOnGitHubAppAuthorizationEventRevoked(func(deliveryID string, eventName string, event *github.GitHubAppAuthorizationEvent) error {
				return nil
			})
			g.SetOnGitHubAppAuthorizationEventRevoked(tt.args.callbacks...)
			if len(g.onGitHubAppAuthorizationEvent[GitHubAppAuthorizationEventRevokedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onGitHubAppAuthorizationEvent[GitHubAppAuthorizationEventRevokedAction]), tt.want)
			}
		})
	}
}

func TestHandleGitHubAppAuthorizationEventRevoked(t *testing.T) {
	action := GitHubAppAuthorizationEventRevokedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.GitHubAppAuthorizationEvent
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
				eventName:  "github_app_authorization",
				event:      &github.GitHubAppAuthorizationEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "github_app_authorization",
				event:      &github.GitHubAppAuthorizationEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "github_app_authorization",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "github_app_authorization",
				event:      &github.GitHubAppAuthorizationEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "github_app_authorization",
				event:      &github.GitHubAppAuthorizationEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "github_app_authorization",
				event:      &github.GitHubAppAuthorizationEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnGitHubAppAuthorizationEventRevoked(func(deliveryID string, eventName string, event *github.GitHubAppAuthorizationEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleGitHubAppAuthorizationEventRevoked(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleGitHubAppAuthorizationEventRevoked() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGitHubAppAuthorizationEvent(t *testing.T) {
	type fields struct {
		handler *EventHandler
	}
	type args struct {
		deliveryID string
		eventName  string
		event      *github.GitHubAppAuthorizationEvent
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "must trigger GitHubAppAuthorizationEventAny with unknown event action",
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
					onGitHubAppAuthorizationEvent: map[string][]GitHubAppAuthorizationEventHandleFunc{
						GitHubAppAuthorizationEventAnyAction: {
							func(deliveryID string, eventName string, event *github.GitHubAppAuthorizationEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  GitHubAppAuthorizationEvent,

				event: &github.GitHubAppAuthorizationEvent{Action: ptrString("unknown")},
			},
			wantErr: false,
		},

		{
			name: "must trigger GitHubAppAuthorizationEventRevoked",
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
					onGitHubAppAuthorizationEvent: map[string][]GitHubAppAuthorizationEventHandleFunc{
						GitHubAppAuthorizationEventAnyAction: {
							func(deliveryID string, eventName string, event *github.GitHubAppAuthorizationEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						GitHubAppAuthorizationEventRevokedAction: {
							func(deliveryID string, eventName string, event *github.GitHubAppAuthorizationEvent) error {
								t.Logf("%s action called", GitHubAppAuthorizationEventRevokedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "github_app_authorization",
				event:      &github.GitHubAppAuthorizationEvent{Action: ptrString(GitHubAppAuthorizationEventRevokedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail GitHubAppAuthorizationEventRevoked with empty action",
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
					onGitHubAppAuthorizationEvent: map[string][]GitHubAppAuthorizationEventHandleFunc{
						GitHubAppAuthorizationEventAnyAction: {
							func(deliveryID string, eventName string, event *github.GitHubAppAuthorizationEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						GitHubAppAuthorizationEventRevokedAction: {
							func(deliveryID string, eventName string, event *github.GitHubAppAuthorizationEvent) error {
								t.Logf("%s action called", GitHubAppAuthorizationEventRevokedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "github_app_authorization",
				event:      &github.GitHubAppAuthorizationEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail GitHubAppAuthorizationEventRevoked with nil action",
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
					onGitHubAppAuthorizationEvent: map[string][]GitHubAppAuthorizationEventHandleFunc{
						GitHubAppAuthorizationEventAnyAction: {
							func(deliveryID string, eventName string, event *github.GitHubAppAuthorizationEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						GitHubAppAuthorizationEventRevokedAction: {
							func(deliveryID string, eventName string, event *github.GitHubAppAuthorizationEvent) error {
								t.Logf("%s action called", GitHubAppAuthorizationEventRevokedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "github_app_authorization",
				event:      &github.GitHubAppAuthorizationEvent{Action: nil},
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
			if err := g.GitHubAppAuthorizationEvent(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("GitHubAppAuthorizationEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
