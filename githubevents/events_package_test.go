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

func TestOnPackageEventAny(t *testing.T) {
	type args struct {
		callbacks []PackageEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single PackageEventHandleFunc",
			args: args{
				[]PackageEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PackageEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple PackageEventHandleFuncs",
			args: args{
				[]PackageEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PackageEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.PackageEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnPackageEventAny(tt.args.callbacks...)
			if len(g.onPackageEvent[PackageEventAnyAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onPackageEvent[PackageEventAnyAction]))
			}
		})
	}
}

func TestSetOnPackageEventAny(t *testing.T) {
	type args struct {
		callbacks []PackageEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single PackageEventHandleFunc",
			args: args{
				[]PackageEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PackageEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple PackageEventHandleFuncs",
			args: args{
				[]PackageEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PackageEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.PackageEvent) error {
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
			g.SetOnPackageEventAny(func(deliveryID string, eventName string, event *github.PackageEvent) error {
				return nil
			})
			g.SetOnPackageEventAny(tt.args.callbacks...)
			if len(g.onPackageEvent[PackageEventAnyAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onPackageEvent[PackageEventAnyAction]), tt.want)
			}
		})
	}
}

func TestHandlePackageEventAny(t *testing.T) {

	action := "*"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.PackageEvent
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
				eventName:  "package",

				event: &github.PackageEvent{Action: &action},

				fail: false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "package",

				event: &github.PackageEvent{Action: &action},

				fail: true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "package",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnPackageEventAny(func(deliveryID string, eventName string, event *github.PackageEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handlePackageEventAny(tt.args.deliveryID, tt.args.deliveryID, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("TestHandlePackageEventAny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnPackageEventPublished(t *testing.T) {
	type args struct {
		callbacks []PackageEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single PackageEventHandleFunc",
			args: args{
				callbacks: []PackageEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PackageEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple PackageEventHandleFunc",
			args: args{
				callbacks: []PackageEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PackageEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.PackageEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnPackageEventPublished(tt.args.callbacks...)
			if len(g.onPackageEvent[PackageEventPublishedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onPackageEvent[PackageEventPublishedAction]))
			}
		})
	}
}

func TestSetOnPackageEventPublished(t *testing.T) {
	type args struct {
		callbacks []PackageEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single PackageEventHandleFunc",
			args: args{
				[]PackageEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PackageEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple PackageEventHandleFuncs",
			args: args{
				[]PackageEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PackageEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.PackageEvent) error {
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
			g.SetOnPackageEventPublished(func(deliveryID string, eventName string, event *github.PackageEvent) error {
				return nil
			})
			g.SetOnPackageEventPublished(tt.args.callbacks...)
			if len(g.onPackageEvent[PackageEventPublishedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onPackageEvent[PackageEventPublishedAction]), tt.want)
			}
		})
	}
}

func TestHandlePackageEventPublished(t *testing.T) {
	action := PackageEventPublishedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.PackageEvent
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
				eventName:  "package",
				event:      &github.PackageEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "package",
				event:      &github.PackageEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "package",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "package",
				event:      &github.PackageEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "package",
				event:      &github.PackageEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "package",
				event:      &github.PackageEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnPackageEventPublished(func(deliveryID string, eventName string, event *github.PackageEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handlePackageEventPublished(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handlePackageEventPublished() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnPackageEventUpdated(t *testing.T) {
	type args struct {
		callbacks []PackageEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single PackageEventHandleFunc",
			args: args{
				callbacks: []PackageEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PackageEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple PackageEventHandleFunc",
			args: args{
				callbacks: []PackageEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PackageEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.PackageEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnPackageEventUpdated(tt.args.callbacks...)
			if len(g.onPackageEvent[PackageEventUpdatedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onPackageEvent[PackageEventUpdatedAction]))
			}
		})
	}
}

func TestSetOnPackageEventUpdated(t *testing.T) {
	type args struct {
		callbacks []PackageEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single PackageEventHandleFunc",
			args: args{
				[]PackageEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PackageEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple PackageEventHandleFuncs",
			args: args{
				[]PackageEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PackageEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.PackageEvent) error {
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
			g.SetOnPackageEventUpdated(func(deliveryID string, eventName string, event *github.PackageEvent) error {
				return nil
			})
			g.SetOnPackageEventUpdated(tt.args.callbacks...)
			if len(g.onPackageEvent[PackageEventUpdatedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onPackageEvent[PackageEventUpdatedAction]), tt.want)
			}
		})
	}
}

func TestHandlePackageEventUpdated(t *testing.T) {
	action := PackageEventUpdatedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.PackageEvent
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
				eventName:  "package",
				event:      &github.PackageEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "package",
				event:      &github.PackageEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "package",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "package",
				event:      &github.PackageEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "package",
				event:      &github.PackageEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "package",
				event:      &github.PackageEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnPackageEventUpdated(func(deliveryID string, eventName string, event *github.PackageEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handlePackageEventUpdated(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handlePackageEventUpdated() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
