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

func TestOnInstallationRepositoriesEventAny(t *testing.T) {
	type args struct {
		callbacks []InstallationRepositoriesEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single InstallationRepositoriesEventHandleFunc",
			args: args{
				[]InstallationRepositoriesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.InstallationRepositoriesEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple InstallationRepositoriesEventHandleFuncs",
			args: args{
				[]InstallationRepositoriesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.InstallationRepositoriesEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.InstallationRepositoriesEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnInstallationRepositoriesEventAny(tt.args.callbacks...)
			if len(g.onInstallationRepositoriesEvent[InstallationRepositoriesEventAnyAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onInstallationRepositoriesEvent[InstallationRepositoriesEventAnyAction]))
			}
		})
	}
}

func TestSetOnInstallationRepositoriesEventAny(t *testing.T) {
	type args struct {
		callbacks []InstallationRepositoriesEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single InstallationRepositoriesEventHandleFunc",
			args: args{
				[]InstallationRepositoriesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.InstallationRepositoriesEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple InstallationRepositoriesEventHandleFuncs",
			args: args{
				[]InstallationRepositoriesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.InstallationRepositoriesEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.InstallationRepositoriesEvent) error {
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
			g.SetOnInstallationRepositoriesEventAny(func(deliveryID string, eventName string, event *github.InstallationRepositoriesEvent) error {
				return nil
			})
			g.SetOnInstallationRepositoriesEventAny(tt.args.callbacks...)
			if len(g.onInstallationRepositoriesEvent[InstallationRepositoriesEventAnyAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onInstallationRepositoriesEvent[InstallationRepositoriesEventAnyAction]), tt.want)
			}
		})
	}
}

func TestHandleInstallationRepositoriesEventAny(t *testing.T) {

	action := "*"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.InstallationRepositoriesEvent
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
				eventName:  "installation_repositories",

				event: &github.InstallationRepositoriesEvent{Action: &action},

				fail: false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "installation_repositories",

				event: &github.InstallationRepositoriesEvent{Action: &action},

				fail: true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "installation_repositories",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnInstallationRepositoriesEventAny(func(deliveryID string, eventName string, event *github.InstallationRepositoriesEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleInstallationRepositoriesEventAny(tt.args.deliveryID, tt.args.deliveryID, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("TestHandleInstallationRepositoriesEventAny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnInstallationRepositoriesEventAdded(t *testing.T) {
	type args struct {
		callbacks []InstallationRepositoriesEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single InstallationRepositoriesEventHandleFunc",
			args: args{
				callbacks: []InstallationRepositoriesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.InstallationRepositoriesEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple InstallationRepositoriesEventHandleFunc",
			args: args{
				callbacks: []InstallationRepositoriesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.InstallationRepositoriesEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.InstallationRepositoriesEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnInstallationRepositoriesEventAdded(tt.args.callbacks...)
			if len(g.onInstallationRepositoriesEvent[InstallationRepositoriesEventAddedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onInstallationRepositoriesEvent[InstallationRepositoriesEventAddedAction]))
			}
		})
	}
}

func TestSetOnInstallationRepositoriesEventAdded(t *testing.T) {
	type args struct {
		callbacks []InstallationRepositoriesEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single InstallationRepositoriesEventHandleFunc",
			args: args{
				[]InstallationRepositoriesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.InstallationRepositoriesEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple InstallationRepositoriesEventHandleFuncs",
			args: args{
				[]InstallationRepositoriesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.InstallationRepositoriesEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.InstallationRepositoriesEvent) error {
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
			g.SetOnInstallationRepositoriesEventAdded(func(deliveryID string, eventName string, event *github.InstallationRepositoriesEvent) error {
				return nil
			})
			g.SetOnInstallationRepositoriesEventAdded(tt.args.callbacks...)
			if len(g.onInstallationRepositoriesEvent[InstallationRepositoriesEventAddedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onInstallationRepositoriesEvent[InstallationRepositoriesEventAddedAction]), tt.want)
			}
		})
	}
}

func TestHandleInstallationRepositoriesEventAdded(t *testing.T) {
	action := InstallationRepositoriesEventAddedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.InstallationRepositoriesEvent
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
				eventName:  "installation_repositories",
				event:      &github.InstallationRepositoriesEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "installation_repositories",
				event:      &github.InstallationRepositoriesEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "installation_repositories",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "installation_repositories",
				event:      &github.InstallationRepositoriesEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "installation_repositories",
				event:      &github.InstallationRepositoriesEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "installation_repositories",
				event:      &github.InstallationRepositoriesEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnInstallationRepositoriesEventAdded(func(deliveryID string, eventName string, event *github.InstallationRepositoriesEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleInstallationRepositoriesEventAdded(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleInstallationRepositoriesEventAdded() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnInstallationRepositoriesEventRemoved(t *testing.T) {
	type args struct {
		callbacks []InstallationRepositoriesEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single InstallationRepositoriesEventHandleFunc",
			args: args{
				callbacks: []InstallationRepositoriesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.InstallationRepositoriesEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple InstallationRepositoriesEventHandleFunc",
			args: args{
				callbacks: []InstallationRepositoriesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.InstallationRepositoriesEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.InstallationRepositoriesEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnInstallationRepositoriesEventRemoved(tt.args.callbacks...)
			if len(g.onInstallationRepositoriesEvent[InstallationRepositoriesEventRemovedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onInstallationRepositoriesEvent[InstallationRepositoriesEventRemovedAction]))
			}
		})
	}
}

func TestSetOnInstallationRepositoriesEventRemoved(t *testing.T) {
	type args struct {
		callbacks []InstallationRepositoriesEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single InstallationRepositoriesEventHandleFunc",
			args: args{
				[]InstallationRepositoriesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.InstallationRepositoriesEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple InstallationRepositoriesEventHandleFuncs",
			args: args{
				[]InstallationRepositoriesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.InstallationRepositoriesEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.InstallationRepositoriesEvent) error {
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
			g.SetOnInstallationRepositoriesEventRemoved(func(deliveryID string, eventName string, event *github.InstallationRepositoriesEvent) error {
				return nil
			})
			g.SetOnInstallationRepositoriesEventRemoved(tt.args.callbacks...)
			if len(g.onInstallationRepositoriesEvent[InstallationRepositoriesEventRemovedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onInstallationRepositoriesEvent[InstallationRepositoriesEventRemovedAction]), tt.want)
			}
		})
	}
}

func TestHandleInstallationRepositoriesEventRemoved(t *testing.T) {
	action := InstallationRepositoriesEventRemovedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.InstallationRepositoriesEvent
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
				eventName:  "installation_repositories",
				event:      &github.InstallationRepositoriesEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "installation_repositories",
				event:      &github.InstallationRepositoriesEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "installation_repositories",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "installation_repositories",
				event:      &github.InstallationRepositoriesEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "installation_repositories",
				event:      &github.InstallationRepositoriesEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "installation_repositories",
				event:      &github.InstallationRepositoriesEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnInstallationRepositoriesEventRemoved(func(deliveryID string, eventName string, event *github.InstallationRepositoriesEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleInstallationRepositoriesEventRemoved(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleInstallationRepositoriesEventRemoved() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
