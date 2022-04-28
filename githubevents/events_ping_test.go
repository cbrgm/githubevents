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

func TestOnPingEventAny(t *testing.T) {
	type args struct {
		callbacks []PingEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single PingEventHandleFunc",
			args: args{
				[]PingEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PingEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple PingEventHandleFuncs",
			args: args{
				[]PingEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PingEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.PingEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnPingEventAny(tt.args.callbacks...)
			if len(g.onPingEvent[PingEventAnyAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onPingEvent[PingEventAnyAction]))
			}
		})
	}
}

func TestSetOnPingEventAny(t *testing.T) {
	type args struct {
		callbacks []PingEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single PingEventHandleFunc",
			args: args{
				[]PingEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PingEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple PingEventHandleFuncs",
			args: args{
				[]PingEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PingEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.PingEvent) error {
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
			g.SetOnPingEventAny(func(deliveryID string, eventName string, event *github.PingEvent) error {
				return nil
			})
			g.SetOnPingEventAny(tt.args.callbacks...)
			if len(g.onPingEvent[PingEventAnyAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onPingEvent[PingEventAnyAction]), tt.want)
			}
		})
	}
}

func TestHandlePingEventAny(t *testing.T) {

	type args struct {
		deliveryID string
		eventName  string
		event      *github.PingEvent
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
				eventName:  "ping",

				event: &github.PingEvent{},

				fail: false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "ping",

				event: &github.PingEvent{},

				fail: true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "ping",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnPingEventAny(func(deliveryID string, eventName string, event *github.PingEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handlePingEventAny(tt.args.deliveryID, tt.args.deliveryID, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("TestHandlePingEventAny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
