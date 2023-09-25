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

func TestOnInstallationEventAny(t *testing.T) {
	type args struct {
		callbacks []InstallationEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single InstallationEventHandleFunc",
			args: args{
				[]InstallationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.InstallationEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple InstallationEventHandleFuncs",
			args: args{
				[]InstallationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.InstallationEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.InstallationEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnInstallationEventAny(tt.args.callbacks...)
			if len(g.onInstallationEvent[InstallationEventAnyAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onInstallationEvent[InstallationEventAnyAction]))
			}
		})
	}
}

func TestSetOnInstallationEventAny(t *testing.T) {
	type args struct {
		callbacks []InstallationEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single InstallationEventHandleFunc",
			args: args{
				[]InstallationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.InstallationEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple InstallationEventHandleFuncs",
			args: args{
				[]InstallationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.InstallationEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.InstallationEvent) error {
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
			g.SetOnInstallationEventAny(func(deliveryID string, eventName string, event *github.InstallationEvent) error {
				return nil
			})
			g.SetOnInstallationEventAny(tt.args.callbacks...)
			if len(g.onInstallationEvent[InstallationEventAnyAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onInstallationEvent[InstallationEventAnyAction]), tt.want)
			}
		})
	}
}

func TestHandleInstallationEventAny(t *testing.T) {

	action := "*"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.InstallationEvent
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
				eventName:  "installation",

				event: &github.InstallationEvent{Action: &action},

				fail: false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "installation",

				event: &github.InstallationEvent{Action: &action},

				fail: true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "installation",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnInstallationEventAny(func(deliveryID string, eventName string, event *github.InstallationEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleInstallationEventAny(tt.args.deliveryID, tt.args.deliveryID, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("TestHandleInstallationEventAny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnInstallationEventCreated(t *testing.T) {
	type args struct {
		callbacks []InstallationEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single InstallationEventHandleFunc",
			args: args{
				callbacks: []InstallationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.InstallationEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple InstallationEventHandleFunc",
			args: args{
				callbacks: []InstallationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.InstallationEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.InstallationEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnInstallationEventCreated(tt.args.callbacks...)
			if len(g.onInstallationEvent[InstallationEventCreatedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onInstallationEvent[InstallationEventCreatedAction]))
			}
		})
	}
}

func TestSetOnInstallationEventCreated(t *testing.T) {
	type args struct {
		callbacks []InstallationEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single InstallationEventHandleFunc",
			args: args{
				[]InstallationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.InstallationEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple InstallationEventHandleFuncs",
			args: args{
				[]InstallationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.InstallationEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.InstallationEvent) error {
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
			g.SetOnInstallationEventCreated(func(deliveryID string, eventName string, event *github.InstallationEvent) error {
				return nil
			})
			g.SetOnInstallationEventCreated(tt.args.callbacks...)
			if len(g.onInstallationEvent[InstallationEventCreatedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onInstallationEvent[InstallationEventCreatedAction]), tt.want)
			}
		})
	}
}

func TestHandleInstallationEventCreated(t *testing.T) {
	action := InstallationEventCreatedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.InstallationEvent
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
				eventName:  "installation",
				event:      &github.InstallationEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "installation",
				event:      &github.InstallationEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "installation",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "installation",
				event:      &github.InstallationEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "installation",
				event:      &github.InstallationEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "installation",
				event:      &github.InstallationEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnInstallationEventCreated(func(deliveryID string, eventName string, event *github.InstallationEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleInstallationEventCreated(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleInstallationEventCreated() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnInstallationEventDeleted(t *testing.T) {
	type args struct {
		callbacks []InstallationEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single InstallationEventHandleFunc",
			args: args{
				callbacks: []InstallationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.InstallationEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple InstallationEventHandleFunc",
			args: args{
				callbacks: []InstallationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.InstallationEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.InstallationEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnInstallationEventDeleted(tt.args.callbacks...)
			if len(g.onInstallationEvent[InstallationEventDeletedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onInstallationEvent[InstallationEventDeletedAction]))
			}
		})
	}
}

func TestSetOnInstallationEventDeleted(t *testing.T) {
	type args struct {
		callbacks []InstallationEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single InstallationEventHandleFunc",
			args: args{
				[]InstallationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.InstallationEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple InstallationEventHandleFuncs",
			args: args{
				[]InstallationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.InstallationEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.InstallationEvent) error {
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
			g.SetOnInstallationEventDeleted(func(deliveryID string, eventName string, event *github.InstallationEvent) error {
				return nil
			})
			g.SetOnInstallationEventDeleted(tt.args.callbacks...)
			if len(g.onInstallationEvent[InstallationEventDeletedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onInstallationEvent[InstallationEventDeletedAction]), tt.want)
			}
		})
	}
}

func TestHandleInstallationEventDeleted(t *testing.T) {
	action := InstallationEventDeletedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.InstallationEvent
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
				eventName:  "installation",
				event:      &github.InstallationEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "installation",
				event:      &github.InstallationEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "installation",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "installation",
				event:      &github.InstallationEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "installation",
				event:      &github.InstallationEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "installation",
				event:      &github.InstallationEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnInstallationEventDeleted(func(deliveryID string, eventName string, event *github.InstallationEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleInstallationEventDeleted(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleInstallationEventDeleted() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnInstallationEventEventSuspend(t *testing.T) {
	type args struct {
		callbacks []InstallationEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single InstallationEventHandleFunc",
			args: args{
				callbacks: []InstallationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.InstallationEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple InstallationEventHandleFunc",
			args: args{
				callbacks: []InstallationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.InstallationEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.InstallationEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnInstallationEventEventSuspend(tt.args.callbacks...)
			if len(g.onInstallationEvent[InstallationEventEventSuspendAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onInstallationEvent[InstallationEventEventSuspendAction]))
			}
		})
	}
}

func TestSetOnInstallationEventEventSuspend(t *testing.T) {
	type args struct {
		callbacks []InstallationEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single InstallationEventHandleFunc",
			args: args{
				[]InstallationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.InstallationEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple InstallationEventHandleFuncs",
			args: args{
				[]InstallationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.InstallationEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.InstallationEvent) error {
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
			g.SetOnInstallationEventEventSuspend(func(deliveryID string, eventName string, event *github.InstallationEvent) error {
				return nil
			})
			g.SetOnInstallationEventEventSuspend(tt.args.callbacks...)
			if len(g.onInstallationEvent[InstallationEventEventSuspendAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onInstallationEvent[InstallationEventEventSuspendAction]), tt.want)
			}
		})
	}
}

func TestHandleInstallationEventEventSuspend(t *testing.T) {
	action := InstallationEventEventSuspendAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.InstallationEvent
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
				eventName:  "installation",
				event:      &github.InstallationEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "installation",
				event:      &github.InstallationEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "installation",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "installation",
				event:      &github.InstallationEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "installation",
				event:      &github.InstallationEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "installation",
				event:      &github.InstallationEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnInstallationEventEventSuspend(func(deliveryID string, eventName string, event *github.InstallationEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleInstallationEventEventSuspend(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleInstallationEventEventSuspend() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnInstallationEventEventUnsuspend(t *testing.T) {
	type args struct {
		callbacks []InstallationEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single InstallationEventHandleFunc",
			args: args{
				callbacks: []InstallationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.InstallationEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple InstallationEventHandleFunc",
			args: args{
				callbacks: []InstallationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.InstallationEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.InstallationEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnInstallationEventEventUnsuspend(tt.args.callbacks...)
			if len(g.onInstallationEvent[InstallationEventEventUnsuspendAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onInstallationEvent[InstallationEventEventUnsuspendAction]))
			}
		})
	}
}

func TestSetOnInstallationEventEventUnsuspend(t *testing.T) {
	type args struct {
		callbacks []InstallationEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single InstallationEventHandleFunc",
			args: args{
				[]InstallationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.InstallationEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple InstallationEventHandleFuncs",
			args: args{
				[]InstallationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.InstallationEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.InstallationEvent) error {
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
			g.SetOnInstallationEventEventUnsuspend(func(deliveryID string, eventName string, event *github.InstallationEvent) error {
				return nil
			})
			g.SetOnInstallationEventEventUnsuspend(tt.args.callbacks...)
			if len(g.onInstallationEvent[InstallationEventEventUnsuspendAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onInstallationEvent[InstallationEventEventUnsuspendAction]), tt.want)
			}
		})
	}
}

func TestHandleInstallationEventEventUnsuspend(t *testing.T) {
	action := InstallationEventEventUnsuspendAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.InstallationEvent
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
				eventName:  "installation",
				event:      &github.InstallationEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "installation",
				event:      &github.InstallationEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "installation",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "installation",
				event:      &github.InstallationEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "installation",
				event:      &github.InstallationEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "installation",
				event:      &github.InstallationEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnInstallationEventEventUnsuspend(func(deliveryID string, eventName string, event *github.InstallationEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleInstallationEventEventUnsuspend(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleInstallationEventEventUnsuspend() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnInstallationEventNewPermissionsAccepted(t *testing.T) {
	type args struct {
		callbacks []InstallationEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single InstallationEventHandleFunc",
			args: args{
				callbacks: []InstallationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.InstallationEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple InstallationEventHandleFunc",
			args: args{
				callbacks: []InstallationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.InstallationEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.InstallationEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnInstallationEventNewPermissionsAccepted(tt.args.callbacks...)
			if len(g.onInstallationEvent[InstallationEventNewPermissionsAcceptedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onInstallationEvent[InstallationEventNewPermissionsAcceptedAction]))
			}
		})
	}
}

func TestSetOnInstallationEventNewPermissionsAccepted(t *testing.T) {
	type args struct {
		callbacks []InstallationEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single InstallationEventHandleFunc",
			args: args{
				[]InstallationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.InstallationEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple InstallationEventHandleFuncs",
			args: args{
				[]InstallationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.InstallationEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.InstallationEvent) error {
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
			g.SetOnInstallationEventNewPermissionsAccepted(func(deliveryID string, eventName string, event *github.InstallationEvent) error {
				return nil
			})
			g.SetOnInstallationEventNewPermissionsAccepted(tt.args.callbacks...)
			if len(g.onInstallationEvent[InstallationEventNewPermissionsAcceptedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onInstallationEvent[InstallationEventNewPermissionsAcceptedAction]), tt.want)
			}
		})
	}
}

func TestHandleInstallationEventNewPermissionsAccepted(t *testing.T) {
	action := InstallationEventNewPermissionsAcceptedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.InstallationEvent
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
				eventName:  "installation",
				event:      &github.InstallationEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "installation",
				event:      &github.InstallationEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "installation",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "installation",
				event:      &github.InstallationEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "installation",
				event:      &github.InstallationEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "installation",
				event:      &github.InstallationEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnInstallationEventNewPermissionsAccepted(func(deliveryID string, eventName string, event *github.InstallationEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleInstallationEventNewPermissionsAccepted(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleInstallationEventNewPermissionsAccepted() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInstallationEvent(t *testing.T) {
	type fields struct {
		handler *EventHandler
	}
	type args struct {
		deliveryID string
		eventName  string
		event      *github.InstallationEvent
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "must trigger InstallationEventAny with unknown event action",
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
					onInstallationEvent: map[string][]InstallationEventHandleFunc{
						InstallationEventAnyAction: {
							func(deliveryID string, eventName string, event *github.InstallationEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  InstallationEvent,

				event: &github.InstallationEvent{Action: ptrString("unknown")},
			},
			wantErr: false,
		},

		{
			name: "must trigger InstallationEventCreated",
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
					onInstallationEvent: map[string][]InstallationEventHandleFunc{
						InstallationEventAnyAction: {
							func(deliveryID string, eventName string, event *github.InstallationEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						InstallationEventCreatedAction: {
							func(deliveryID string, eventName string, event *github.InstallationEvent) error {
								t.Logf("%s action called", InstallationEventCreatedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "installation",
				event:      &github.InstallationEvent{Action: ptrString(InstallationEventCreatedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail InstallationEventCreated with empty action",
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
					onInstallationEvent: map[string][]InstallationEventHandleFunc{
						InstallationEventAnyAction: {
							func(deliveryID string, eventName string, event *github.InstallationEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						InstallationEventCreatedAction: {
							func(deliveryID string, eventName string, event *github.InstallationEvent) error {
								t.Logf("%s action called", InstallationEventCreatedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "installation",
				event:      &github.InstallationEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail InstallationEventCreated with nil action",
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
					onInstallationEvent: map[string][]InstallationEventHandleFunc{
						InstallationEventAnyAction: {
							func(deliveryID string, eventName string, event *github.InstallationEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						InstallationEventCreatedAction: {
							func(deliveryID string, eventName string, event *github.InstallationEvent) error {
								t.Logf("%s action called", InstallationEventCreatedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "installation",
				event:      &github.InstallationEvent{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger InstallationEventDeleted",
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
					onInstallationEvent: map[string][]InstallationEventHandleFunc{
						InstallationEventAnyAction: {
							func(deliveryID string, eventName string, event *github.InstallationEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						InstallationEventDeletedAction: {
							func(deliveryID string, eventName string, event *github.InstallationEvent) error {
								t.Logf("%s action called", InstallationEventDeletedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "installation",
				event:      &github.InstallationEvent{Action: ptrString(InstallationEventDeletedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail InstallationEventDeleted with empty action",
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
					onInstallationEvent: map[string][]InstallationEventHandleFunc{
						InstallationEventAnyAction: {
							func(deliveryID string, eventName string, event *github.InstallationEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						InstallationEventDeletedAction: {
							func(deliveryID string, eventName string, event *github.InstallationEvent) error {
								t.Logf("%s action called", InstallationEventDeletedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "installation",
				event:      &github.InstallationEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail InstallationEventDeleted with nil action",
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
					onInstallationEvent: map[string][]InstallationEventHandleFunc{
						InstallationEventAnyAction: {
							func(deliveryID string, eventName string, event *github.InstallationEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						InstallationEventDeletedAction: {
							func(deliveryID string, eventName string, event *github.InstallationEvent) error {
								t.Logf("%s action called", InstallationEventDeletedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "installation",
				event:      &github.InstallationEvent{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger InstallationEventEventSuspend",
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
					onInstallationEvent: map[string][]InstallationEventHandleFunc{
						InstallationEventAnyAction: {
							func(deliveryID string, eventName string, event *github.InstallationEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						InstallationEventEventSuspendAction: {
							func(deliveryID string, eventName string, event *github.InstallationEvent) error {
								t.Logf("%s action called", InstallationEventEventSuspendAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "installation",
				event:      &github.InstallationEvent{Action: ptrString(InstallationEventEventSuspendAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail InstallationEventEventSuspend with empty action",
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
					onInstallationEvent: map[string][]InstallationEventHandleFunc{
						InstallationEventAnyAction: {
							func(deliveryID string, eventName string, event *github.InstallationEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						InstallationEventEventSuspendAction: {
							func(deliveryID string, eventName string, event *github.InstallationEvent) error {
								t.Logf("%s action called", InstallationEventEventSuspendAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "installation",
				event:      &github.InstallationEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail InstallationEventEventSuspend with nil action",
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
					onInstallationEvent: map[string][]InstallationEventHandleFunc{
						InstallationEventAnyAction: {
							func(deliveryID string, eventName string, event *github.InstallationEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						InstallationEventEventSuspendAction: {
							func(deliveryID string, eventName string, event *github.InstallationEvent) error {
								t.Logf("%s action called", InstallationEventEventSuspendAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "installation",
				event:      &github.InstallationEvent{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger InstallationEventEventUnsuspend",
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
					onInstallationEvent: map[string][]InstallationEventHandleFunc{
						InstallationEventAnyAction: {
							func(deliveryID string, eventName string, event *github.InstallationEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						InstallationEventEventUnsuspendAction: {
							func(deliveryID string, eventName string, event *github.InstallationEvent) error {
								t.Logf("%s action called", InstallationEventEventUnsuspendAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "installation",
				event:      &github.InstallationEvent{Action: ptrString(InstallationEventEventUnsuspendAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail InstallationEventEventUnsuspend with empty action",
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
					onInstallationEvent: map[string][]InstallationEventHandleFunc{
						InstallationEventAnyAction: {
							func(deliveryID string, eventName string, event *github.InstallationEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						InstallationEventEventUnsuspendAction: {
							func(deliveryID string, eventName string, event *github.InstallationEvent) error {
								t.Logf("%s action called", InstallationEventEventUnsuspendAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "installation",
				event:      &github.InstallationEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail InstallationEventEventUnsuspend with nil action",
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
					onInstallationEvent: map[string][]InstallationEventHandleFunc{
						InstallationEventAnyAction: {
							func(deliveryID string, eventName string, event *github.InstallationEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						InstallationEventEventUnsuspendAction: {
							func(deliveryID string, eventName string, event *github.InstallationEvent) error {
								t.Logf("%s action called", InstallationEventEventUnsuspendAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "installation",
				event:      &github.InstallationEvent{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger InstallationEventNewPermissionsAccepted",
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
					onInstallationEvent: map[string][]InstallationEventHandleFunc{
						InstallationEventAnyAction: {
							func(deliveryID string, eventName string, event *github.InstallationEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						InstallationEventNewPermissionsAcceptedAction: {
							func(deliveryID string, eventName string, event *github.InstallationEvent) error {
								t.Logf("%s action called", InstallationEventNewPermissionsAcceptedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "installation",
				event:      &github.InstallationEvent{Action: ptrString(InstallationEventNewPermissionsAcceptedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail InstallationEventNewPermissionsAccepted with empty action",
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
					onInstallationEvent: map[string][]InstallationEventHandleFunc{
						InstallationEventAnyAction: {
							func(deliveryID string, eventName string, event *github.InstallationEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						InstallationEventNewPermissionsAcceptedAction: {
							func(deliveryID string, eventName string, event *github.InstallationEvent) error {
								t.Logf("%s action called", InstallationEventNewPermissionsAcceptedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "installation",
				event:      &github.InstallationEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail InstallationEventNewPermissionsAccepted with nil action",
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
					onInstallationEvent: map[string][]InstallationEventHandleFunc{
						InstallationEventAnyAction: {
							func(deliveryID string, eventName string, event *github.InstallationEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						InstallationEventNewPermissionsAcceptedAction: {
							func(deliveryID string, eventName string, event *github.InstallationEvent) error {
								t.Logf("%s action called", InstallationEventNewPermissionsAcceptedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "installation",
				event:      &github.InstallationEvent{Action: nil},
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
			if err := g.InstallationEvent(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("InstallationEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
