// Copyright 2022 The GithubEvents Authors. All rights reserved.
// Use of this source code is governed by the MIT License
// that can be found in the LICENSE file.

package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"context"
	"errors"
	"github.com/google/go-github/v69/github"
	"go.opentelemetry.io/otel/trace/noop"
	"sync"
	"testing"
)

func TestOnPageBuildEventAny(t *testing.T) {
	type args struct {
		callbacks []PageBuildEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single PageBuildEventHandleFunc",
			args: args{
				[]PageBuildEventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.PageBuildEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple PageBuildEventHandleFuncs",
			args: args{
				[]PageBuildEventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.PageBuildEvent) error {
						return nil
					},
					func(ctx context.Context, deliveryID string, eventName string, event *github.PageBuildEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnPageBuildEventAny(tt.args.callbacks...)
			if len(g.onPageBuildEvent[PageBuildEventAnyAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onPageBuildEvent[PageBuildEventAnyAction]))
			}
		})
	}
}

func TestSetOnPageBuildEventAny(t *testing.T) {
	type args struct {
		callbacks []PageBuildEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single PageBuildEventHandleFunc",
			args: args{
				[]PageBuildEventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.PageBuildEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple PageBuildEventHandleFuncs",
			args: args{
				[]PageBuildEventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.PageBuildEvent) error {
						return nil
					},
					func(ctx context.Context, deliveryID string, eventName string, event *github.PageBuildEvent) error {
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
			g.SetOnPageBuildEventAny(func(ctx context.Context, deliveryID string, eventName string, event *github.PageBuildEvent) error {
				return nil
			})
			g.SetOnPageBuildEventAny(tt.args.callbacks...)
			if len(g.onPageBuildEvent[PageBuildEventAnyAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onPageBuildEvent[PageBuildEventAnyAction]), tt.want)
			}
		})
	}
}

func TestHandlePageBuildEventAny(t *testing.T) {

	type args struct {
		deliveryID string
		eventName  string
		event      *github.PageBuildEvent
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
				eventName:  "page_build",

				event: &github.PageBuildEvent{},

				fail: false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "page_build",

				event: &github.PageBuildEvent{},

				fail: true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "page_build",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnPageBuildEventAny(func(ctx context.Context, deliveryID string, eventName string, event *github.PageBuildEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handlePageBuildEventAny(context.Background(), tt.args.deliveryID, tt.args.deliveryID, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("TestHandlePageBuildEventAny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPageBuildEvent(t *testing.T) {
	type fields struct {
		handler *EventHandler
	}
	type args struct {
		deliveryID string
		eventName  string
		event      *github.PageBuildEvent
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "must trigger PageBuildEventAny with unknown event action",
			fields: fields{
				handler: &EventHandler{
					WebhookSecret: "fake",
					onBeforeAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onBeforeAny called")
								return nil
							},
						},
					},
					onAfterAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onAfterAny called")
								return nil
							},
						},
					},
					onPageBuildEvent: map[string][]PageBuildEventHandleFunc{
						PageBuildEventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.PageBuildEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
					},
					Tracer: noop.Tracer{},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  PageBuildEvent,

				event: &github.PageBuildEvent{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &EventHandler{
				WebhookSecret: "fake",
				mu:            sync.RWMutex{},
				Tracer:        noop.Tracer{},
			}
			if err := g.PageBuildEvent(context.Background(), tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("PageBuildEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
