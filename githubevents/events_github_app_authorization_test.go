package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"errors"
	"github.com/google/go-github/v43/github"
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
			if len(g.onGitHubAppAuthorizationEvent["*"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onGitHubAppAuthorizationEvent["*"]))
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
			if len(g.onGitHubAppAuthorizationEvent["*"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onGitHubAppAuthorizationEvent["*"]), tt.want)
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
			if len(g.onGitHubAppAuthorizationEvent["revoked"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onGitHubAppAuthorizationEvent["revoked"]))
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
			if len(g.onGitHubAppAuthorizationEvent["revoked"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onGitHubAppAuthorizationEvent["revoked"]), tt.want)
			}
		})
	}
}

func TestHandleGitHubAppAuthorizationEventRevoked(t *testing.T) {
	action := "revoked"

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
