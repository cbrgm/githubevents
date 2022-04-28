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

func TestOnPushEventAny(t *testing.T) {
	type args struct {
		callbacks []PushEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single PushEventHandleFunc",
			args: args{
				[]PushEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PushEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple PushEventHandleFuncs",
			args: args{
				[]PushEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PushEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.PushEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnPushEventAny(tt.args.callbacks...)
			if len(g.onPushEvent[PushEventAnyAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onPushEvent[PushEventAnyAction]))
			}
		})
	}
}

func TestSetOnPushEventAny(t *testing.T) {
	type args struct {
		callbacks []PushEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single PushEventHandleFunc",
			args: args{
				[]PushEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PushEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple PushEventHandleFuncs",
			args: args{
				[]PushEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PushEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.PushEvent) error {
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
			g.SetOnPushEventAny(func(deliveryID string, eventName string, event *github.PushEvent) error {
				return nil
			})
			g.SetOnPushEventAny(tt.args.callbacks...)
			if len(g.onPushEvent[PushEventAnyAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onPushEvent[PushEventAnyAction]), tt.want)
			}
		})
	}
}

func TestHandlePushEventAny(t *testing.T) {

	type args struct {
		deliveryID string
		eventName  string
		event      *github.PushEvent
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
				eventName:  "push",

				event: &github.PushEvent{},

				fail: false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "push",

				event: &github.PushEvent{},

				fail: true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "push",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnPushEventAny(func(deliveryID string, eventName string, event *github.PushEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handlePushEventAny(tt.args.deliveryID, tt.args.deliveryID, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("TestHandlePushEventAny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
