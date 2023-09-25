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

func TestOnMembershipEventAny(t *testing.T) {
	type args struct {
		callbacks []MembershipEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single MembershipEventHandleFunc",
			args: args{
				[]MembershipEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MembershipEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple MembershipEventHandleFuncs",
			args: args{
				[]MembershipEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MembershipEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.MembershipEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnMembershipEventAny(tt.args.callbacks...)
			if len(g.onMembershipEvent[MembershipEventAnyAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onMembershipEvent[MembershipEventAnyAction]))
			}
		})
	}
}

func TestSetOnMembershipEventAny(t *testing.T) {
	type args struct {
		callbacks []MembershipEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single MembershipEventHandleFunc",
			args: args{
				[]MembershipEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MembershipEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple MembershipEventHandleFuncs",
			args: args{
				[]MembershipEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MembershipEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.MembershipEvent) error {
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
			g.SetOnMembershipEventAny(func(deliveryID string, eventName string, event *github.MembershipEvent) error {
				return nil
			})
			g.SetOnMembershipEventAny(tt.args.callbacks...)
			if len(g.onMembershipEvent[MembershipEventAnyAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onMembershipEvent[MembershipEventAnyAction]), tt.want)
			}
		})
	}
}

func TestHandleMembershipEventAny(t *testing.T) {

	action := "*"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.MembershipEvent
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
				eventName:  "membership",

				event: &github.MembershipEvent{Action: &action},

				fail: false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "membership",

				event: &github.MembershipEvent{Action: &action},

				fail: true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "membership",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnMembershipEventAny(func(deliveryID string, eventName string, event *github.MembershipEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleMembershipEventAny(tt.args.deliveryID, tt.args.deliveryID, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("TestHandleMembershipEventAny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnMembershipEventAdded(t *testing.T) {
	type args struct {
		callbacks []MembershipEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single MembershipEventHandleFunc",
			args: args{
				callbacks: []MembershipEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MembershipEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple MembershipEventHandleFunc",
			args: args{
				callbacks: []MembershipEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MembershipEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.MembershipEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnMembershipEventAdded(tt.args.callbacks...)
			if len(g.onMembershipEvent[MembershipEventAddedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onMembershipEvent[MembershipEventAddedAction]))
			}
		})
	}
}

func TestSetOnMembershipEventAdded(t *testing.T) {
	type args struct {
		callbacks []MembershipEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single MembershipEventHandleFunc",
			args: args{
				[]MembershipEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MembershipEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple MembershipEventHandleFuncs",
			args: args{
				[]MembershipEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MembershipEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.MembershipEvent) error {
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
			g.SetOnMembershipEventAdded(func(deliveryID string, eventName string, event *github.MembershipEvent) error {
				return nil
			})
			g.SetOnMembershipEventAdded(tt.args.callbacks...)
			if len(g.onMembershipEvent[MembershipEventAddedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onMembershipEvent[MembershipEventAddedAction]), tt.want)
			}
		})
	}
}

func TestHandleMembershipEventAdded(t *testing.T) {
	action := MembershipEventAddedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.MembershipEvent
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
				eventName:  "membership",
				event:      &github.MembershipEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "membership",
				event:      &github.MembershipEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "membership",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "membership",
				event:      &github.MembershipEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "membership",
				event:      &github.MembershipEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "membership",
				event:      &github.MembershipEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnMembershipEventAdded(func(deliveryID string, eventName string, event *github.MembershipEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleMembershipEventAdded(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleMembershipEventAdded() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnMembershipEventRemoved(t *testing.T) {
	type args struct {
		callbacks []MembershipEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single MembershipEventHandleFunc",
			args: args{
				callbacks: []MembershipEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MembershipEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple MembershipEventHandleFunc",
			args: args{
				callbacks: []MembershipEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MembershipEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.MembershipEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnMembershipEventRemoved(tt.args.callbacks...)
			if len(g.onMembershipEvent[MembershipEventRemovedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onMembershipEvent[MembershipEventRemovedAction]))
			}
		})
	}
}

func TestSetOnMembershipEventRemoved(t *testing.T) {
	type args struct {
		callbacks []MembershipEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single MembershipEventHandleFunc",
			args: args{
				[]MembershipEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MembershipEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple MembershipEventHandleFuncs",
			args: args{
				[]MembershipEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MembershipEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.MembershipEvent) error {
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
			g.SetOnMembershipEventRemoved(func(deliveryID string, eventName string, event *github.MembershipEvent) error {
				return nil
			})
			g.SetOnMembershipEventRemoved(tt.args.callbacks...)
			if len(g.onMembershipEvent[MembershipEventRemovedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onMembershipEvent[MembershipEventRemovedAction]), tt.want)
			}
		})
	}
}

func TestHandleMembershipEventRemoved(t *testing.T) {
	action := MembershipEventRemovedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.MembershipEvent
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
				eventName:  "membership",
				event:      &github.MembershipEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "membership",
				event:      &github.MembershipEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "membership",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "membership",
				event:      &github.MembershipEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "membership",
				event:      &github.MembershipEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "membership",
				event:      &github.MembershipEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnMembershipEventRemoved(func(deliveryID string, eventName string, event *github.MembershipEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleMembershipEventRemoved(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleMembershipEventRemoved() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMembershipEvent(t *testing.T) {
	type fields struct {
		handler *EventHandler
	}
	type args struct {
		deliveryID string
		eventName  string
		event      *github.MembershipEvent
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "must trigger MembershipEventAny with unknown event action",
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
					onMembershipEvent: map[string][]MembershipEventHandleFunc{
						MembershipEventAnyAction: {
							func(deliveryID string, eventName string, event *github.MembershipEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  MembershipEvent,

				event: &github.MembershipEvent{Action: ptrString("unknown")},
			},
			wantErr: false,
		},

		{
			name: "must trigger MembershipEventAdded",
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
					onMembershipEvent: map[string][]MembershipEventHandleFunc{
						MembershipEventAnyAction: {
							func(deliveryID string, eventName string, event *github.MembershipEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						MembershipEventAddedAction: {
							func(deliveryID string, eventName string, event *github.MembershipEvent) error {
								t.Logf("%s action called", MembershipEventAddedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "membership",
				event:      &github.MembershipEvent{Action: ptrString(MembershipEventAddedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail MembershipEventAdded with empty action",
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
					onMembershipEvent: map[string][]MembershipEventHandleFunc{
						MembershipEventAnyAction: {
							func(deliveryID string, eventName string, event *github.MembershipEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						MembershipEventAddedAction: {
							func(deliveryID string, eventName string, event *github.MembershipEvent) error {
								t.Logf("%s action called", MembershipEventAddedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "membership",
				event:      &github.MembershipEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail MembershipEventAdded with nil action",
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
					onMembershipEvent: map[string][]MembershipEventHandleFunc{
						MembershipEventAnyAction: {
							func(deliveryID string, eventName string, event *github.MembershipEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						MembershipEventAddedAction: {
							func(deliveryID string, eventName string, event *github.MembershipEvent) error {
								t.Logf("%s action called", MembershipEventAddedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "membership",
				event:      &github.MembershipEvent{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger MembershipEventRemoved",
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
					onMembershipEvent: map[string][]MembershipEventHandleFunc{
						MembershipEventAnyAction: {
							func(deliveryID string, eventName string, event *github.MembershipEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						MembershipEventRemovedAction: {
							func(deliveryID string, eventName string, event *github.MembershipEvent) error {
								t.Logf("%s action called", MembershipEventRemovedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "membership",
				event:      &github.MembershipEvent{Action: ptrString(MembershipEventRemovedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail MembershipEventRemoved with empty action",
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
					onMembershipEvent: map[string][]MembershipEventHandleFunc{
						MembershipEventAnyAction: {
							func(deliveryID string, eventName string, event *github.MembershipEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						MembershipEventRemovedAction: {
							func(deliveryID string, eventName string, event *github.MembershipEvent) error {
								t.Logf("%s action called", MembershipEventRemovedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "membership",
				event:      &github.MembershipEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail MembershipEventRemoved with nil action",
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
					onMembershipEvent: map[string][]MembershipEventHandleFunc{
						MembershipEventAnyAction: {
							func(deliveryID string, eventName string, event *github.MembershipEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						MembershipEventRemovedAction: {
							func(deliveryID string, eventName string, event *github.MembershipEvent) error {
								t.Logf("%s action called", MembershipEventRemovedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "membership",
				event:      &github.MembershipEvent{Action: nil},
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
			if err := g.MembershipEvent(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("MembershipEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
