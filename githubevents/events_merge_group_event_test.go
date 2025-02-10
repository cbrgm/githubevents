// Copyright 2022 The GithubEvents Authors. All rights reserved.
// Use of this source code is governed by the MIT License
// that can be found in the LICENSE file.

package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"errors"
	"github.com/google/go-github/v69/github"
	"sync"
	"testing"
)

func TestOnMergeGroupEventAny(t *testing.T) {
	type args struct {
		callbacks []MergeGroupEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single MergeGroupEventHandleFunc",
			args: args{
				[]MergeGroupEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MergeGroupEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple MergeGroupEventHandleFuncs",
			args: args{
				[]MergeGroupEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MergeGroupEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.MergeGroupEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnMergeGroupEventAny(tt.args.callbacks...)
			if len(g.onMergeGroupEvent[MergeGroupEventAnyAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onMergeGroupEvent[MergeGroupEventAnyAction]))
			}
		})
	}
}

func TestSetOnMergeGroupEventAny(t *testing.T) {
	type args struct {
		callbacks []MergeGroupEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single MergeGroupEventHandleFunc",
			args: args{
				[]MergeGroupEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MergeGroupEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple MergeGroupEventHandleFuncs",
			args: args{
				[]MergeGroupEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MergeGroupEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.MergeGroupEvent) error {
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
			g.SetOnMergeGroupEventAny(func(deliveryID string, eventName string, event *github.MergeGroupEvent) error {
				return nil
			})
			g.SetOnMergeGroupEventAny(tt.args.callbacks...)
			if len(g.onMergeGroupEvent[MergeGroupEventAnyAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onMergeGroupEvent[MergeGroupEventAnyAction]), tt.want)
			}
		})
	}
}

func TestHandleMergeGroupEventAny(t *testing.T) {

	action := "*"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.MergeGroupEvent
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
				eventName:  "merge_group_event",

				event: &github.MergeGroupEvent{Action: &action},

				fail: false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "merge_group_event",

				event: &github.MergeGroupEvent{Action: &action},

				fail: true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "merge_group_event",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnMergeGroupEventAny(func(deliveryID string, eventName string, event *github.MergeGroupEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleMergeGroupEventAny(tt.args.deliveryID, tt.args.deliveryID, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("TestHandleMergeGroupEventAny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnMergeGroupEventChecksRequested(t *testing.T) {
	type args struct {
		callbacks []MergeGroupEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single MergeGroupEventHandleFunc",
			args: args{
				callbacks: []MergeGroupEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MergeGroupEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple MergeGroupEventHandleFunc",
			args: args{
				callbacks: []MergeGroupEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MergeGroupEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.MergeGroupEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnMergeGroupEventChecksRequested(tt.args.callbacks...)
			if len(g.onMergeGroupEvent[MergeGroupEventChecksRequestedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onMergeGroupEvent[MergeGroupEventChecksRequestedAction]))
			}
		})
	}
}

func TestSetOnMergeGroupEventChecksRequested(t *testing.T) {
	type args struct {
		callbacks []MergeGroupEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single MergeGroupEventHandleFunc",
			args: args{
				[]MergeGroupEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MergeGroupEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple MergeGroupEventHandleFuncs",
			args: args{
				[]MergeGroupEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MergeGroupEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.MergeGroupEvent) error {
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
			g.SetOnMergeGroupEventChecksRequested(func(deliveryID string, eventName string, event *github.MergeGroupEvent) error {
				return nil
			})
			g.SetOnMergeGroupEventChecksRequested(tt.args.callbacks...)
			if len(g.onMergeGroupEvent[MergeGroupEventChecksRequestedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onMergeGroupEvent[MergeGroupEventChecksRequestedAction]), tt.want)
			}
		})
	}
}

func TestHandleMergeGroupEventChecksRequested(t *testing.T) {
	action := MergeGroupEventChecksRequestedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.MergeGroupEvent
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
				eventName:  "merge_group_event",
				event:      &github.MergeGroupEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "merge_group_event",
				event:      &github.MergeGroupEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "merge_group_event",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "merge_group_event",
				event:      &github.MergeGroupEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "merge_group_event",
				event:      &github.MergeGroupEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "merge_group_event",
				event:      &github.MergeGroupEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnMergeGroupEventChecksRequested(func(deliveryID string, eventName string, event *github.MergeGroupEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleMergeGroupEventChecksRequested(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleMergeGroupEventChecksRequested() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnMergeGroupEventDestroyed(t *testing.T) {
	type args struct {
		callbacks []MergeGroupEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single MergeGroupEventHandleFunc",
			args: args{
				callbacks: []MergeGroupEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MergeGroupEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple MergeGroupEventHandleFunc",
			args: args{
				callbacks: []MergeGroupEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MergeGroupEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.MergeGroupEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnMergeGroupEventDestroyed(tt.args.callbacks...)
			if len(g.onMergeGroupEvent[MergeGroupEventDestroyedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onMergeGroupEvent[MergeGroupEventDestroyedAction]))
			}
		})
	}
}

func TestSetOnMergeGroupEventDestroyed(t *testing.T) {
	type args struct {
		callbacks []MergeGroupEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single MergeGroupEventHandleFunc",
			args: args{
				[]MergeGroupEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MergeGroupEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple MergeGroupEventHandleFuncs",
			args: args{
				[]MergeGroupEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MergeGroupEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.MergeGroupEvent) error {
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
			g.SetOnMergeGroupEventDestroyed(func(deliveryID string, eventName string, event *github.MergeGroupEvent) error {
				return nil
			})
			g.SetOnMergeGroupEventDestroyed(tt.args.callbacks...)
			if len(g.onMergeGroupEvent[MergeGroupEventDestroyedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onMergeGroupEvent[MergeGroupEventDestroyedAction]), tt.want)
			}
		})
	}
}

func TestHandleMergeGroupEventDestroyed(t *testing.T) {
	action := MergeGroupEventDestroyedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.MergeGroupEvent
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
				eventName:  "merge_group_event",
				event:      &github.MergeGroupEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "merge_group_event",
				event:      &github.MergeGroupEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "merge_group_event",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "merge_group_event",
				event:      &github.MergeGroupEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "merge_group_event",
				event:      &github.MergeGroupEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "merge_group_event",
				event:      &github.MergeGroupEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnMergeGroupEventDestroyed(func(deliveryID string, eventName string, event *github.MergeGroupEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleMergeGroupEventDestroyed(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleMergeGroupEventDestroyed() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMergeGroupEvent(t *testing.T) {
	type fields struct {
		handler *EventHandler
	}
	type args struct {
		deliveryID string
		eventName  string
		event      *github.MergeGroupEvent
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "must trigger MergeGroupEventAny with unknown event action",
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
					onMergeGroupEvent: map[string][]MergeGroupEventHandleFunc{
						MergeGroupEventAnyAction: {
							func(deliveryID string, eventName string, event *github.MergeGroupEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  MergeGroupEvent,

				event: &github.MergeGroupEvent{Action: ptrString("unknown")},
			},
			wantErr: false,
		},

		{
			name: "must trigger MergeGroupEventChecksRequested",
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
					onMergeGroupEvent: map[string][]MergeGroupEventHandleFunc{
						MergeGroupEventAnyAction: {
							func(deliveryID string, eventName string, event *github.MergeGroupEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						MergeGroupEventChecksRequestedAction: {
							func(deliveryID string, eventName string, event *github.MergeGroupEvent) error {
								t.Logf("%s action called", MergeGroupEventChecksRequestedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "merge_group_event",
				event:      &github.MergeGroupEvent{Action: ptrString(MergeGroupEventChecksRequestedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail MergeGroupEventChecksRequested with empty action",
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
					onMergeGroupEvent: map[string][]MergeGroupEventHandleFunc{
						MergeGroupEventAnyAction: {
							func(deliveryID string, eventName string, event *github.MergeGroupEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						MergeGroupEventChecksRequestedAction: {
							func(deliveryID string, eventName string, event *github.MergeGroupEvent) error {
								t.Logf("%s action called", MergeGroupEventChecksRequestedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "merge_group_event",
				event:      &github.MergeGroupEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail MergeGroupEventChecksRequested with nil action",
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
					onMergeGroupEvent: map[string][]MergeGroupEventHandleFunc{
						MergeGroupEventAnyAction: {
							func(deliveryID string, eventName string, event *github.MergeGroupEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						MergeGroupEventChecksRequestedAction: {
							func(deliveryID string, eventName string, event *github.MergeGroupEvent) error {
								t.Logf("%s action called", MergeGroupEventChecksRequestedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "merge_group_event",
				event:      &github.MergeGroupEvent{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger MergeGroupEventDestroyed",
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
					onMergeGroupEvent: map[string][]MergeGroupEventHandleFunc{
						MergeGroupEventAnyAction: {
							func(deliveryID string, eventName string, event *github.MergeGroupEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						MergeGroupEventDestroyedAction: {
							func(deliveryID string, eventName string, event *github.MergeGroupEvent) error {
								t.Logf("%s action called", MergeGroupEventDestroyedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "merge_group_event",
				event:      &github.MergeGroupEvent{Action: ptrString(MergeGroupEventDestroyedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail MergeGroupEventDestroyed with empty action",
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
					onMergeGroupEvent: map[string][]MergeGroupEventHandleFunc{
						MergeGroupEventAnyAction: {
							func(deliveryID string, eventName string, event *github.MergeGroupEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						MergeGroupEventDestroyedAction: {
							func(deliveryID string, eventName string, event *github.MergeGroupEvent) error {
								t.Logf("%s action called", MergeGroupEventDestroyedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "merge_group_event",
				event:      &github.MergeGroupEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail MergeGroupEventDestroyed with nil action",
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
					onMergeGroupEvent: map[string][]MergeGroupEventHandleFunc{
						MergeGroupEventAnyAction: {
							func(deliveryID string, eventName string, event *github.MergeGroupEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						MergeGroupEventDestroyedAction: {
							func(deliveryID string, eventName string, event *github.MergeGroupEvent) error {
								t.Logf("%s action called", MergeGroupEventDestroyedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "merge_group_event",
				event:      &github.MergeGroupEvent{Action: nil},
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
			if err := g.MergeGroupEvent(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("MergeGroupEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
