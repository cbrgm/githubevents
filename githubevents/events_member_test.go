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

func TestOnMemberEventAny(t *testing.T) {
	type args struct {
		callbacks []MemberEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single MemberEventHandleFunc",
			args: args{
				[]MemberEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MemberEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple MemberEventHandleFuncs",
			args: args{
				[]MemberEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MemberEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.MemberEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnMemberEventAny(tt.args.callbacks...)
			if len(g.onMemberEvent[MemberEventAnyAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onMemberEvent[MemberEventAnyAction]))
			}
		})
	}
}

func TestSetOnMemberEventAny(t *testing.T) {
	type args struct {
		callbacks []MemberEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single MemberEventHandleFunc",
			args: args{
				[]MemberEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MemberEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple MemberEventHandleFuncs",
			args: args{
				[]MemberEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MemberEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.MemberEvent) error {
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
			g.SetOnMemberEventAny(func(deliveryID string, eventName string, event *github.MemberEvent) error {
				return nil
			})
			g.SetOnMemberEventAny(tt.args.callbacks...)
			if len(g.onMemberEvent[MemberEventAnyAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onMemberEvent[MemberEventAnyAction]), tt.want)
			}
		})
	}
}

func TestHandleMemberEventAny(t *testing.T) {

	action := "*"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.MemberEvent
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
				eventName:  "member",

				event: &github.MemberEvent{Action: &action},

				fail: false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "member",

				event: &github.MemberEvent{Action: &action},

				fail: true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "member",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnMemberEventAny(func(deliveryID string, eventName string, event *github.MemberEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleMemberEventAny(tt.args.deliveryID, tt.args.deliveryID, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("TestHandleMemberEventAny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnMemberEventAdded(t *testing.T) {
	type args struct {
		callbacks []MemberEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single MemberEventHandleFunc",
			args: args{
				callbacks: []MemberEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MemberEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple MemberEventHandleFunc",
			args: args{
				callbacks: []MemberEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MemberEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.MemberEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnMemberEventAdded(tt.args.callbacks...)
			if len(g.onMemberEvent[MemberEventAddedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onMemberEvent[MemberEventAddedAction]))
			}
		})
	}
}

func TestSetOnMemberEventAdded(t *testing.T) {
	type args struct {
		callbacks []MemberEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single MemberEventHandleFunc",
			args: args{
				[]MemberEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MemberEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple MemberEventHandleFuncs",
			args: args{
				[]MemberEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MemberEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.MemberEvent) error {
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
			g.SetOnMemberEventAdded(func(deliveryID string, eventName string, event *github.MemberEvent) error {
				return nil
			})
			g.SetOnMemberEventAdded(tt.args.callbacks...)
			if len(g.onMemberEvent[MemberEventAddedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onMemberEvent[MemberEventAddedAction]), tt.want)
			}
		})
	}
}

func TestHandleMemberEventAdded(t *testing.T) {
	action := MemberEventAddedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.MemberEvent
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
				eventName:  "member",
				event:      &github.MemberEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "member",
				event:      &github.MemberEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "member",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "member",
				event:      &github.MemberEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "member",
				event:      &github.MemberEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "member",
				event:      &github.MemberEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnMemberEventAdded(func(deliveryID string, eventName string, event *github.MemberEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleMemberEventAdded(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleMemberEventAdded() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnMemberEventRemoved(t *testing.T) {
	type args struct {
		callbacks []MemberEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single MemberEventHandleFunc",
			args: args{
				callbacks: []MemberEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MemberEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple MemberEventHandleFunc",
			args: args{
				callbacks: []MemberEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MemberEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.MemberEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnMemberEventRemoved(tt.args.callbacks...)
			if len(g.onMemberEvent[MemberEventRemovedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onMemberEvent[MemberEventRemovedAction]))
			}
		})
	}
}

func TestSetOnMemberEventRemoved(t *testing.T) {
	type args struct {
		callbacks []MemberEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single MemberEventHandleFunc",
			args: args{
				[]MemberEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MemberEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple MemberEventHandleFuncs",
			args: args{
				[]MemberEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MemberEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.MemberEvent) error {
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
			g.SetOnMemberEventRemoved(func(deliveryID string, eventName string, event *github.MemberEvent) error {
				return nil
			})
			g.SetOnMemberEventRemoved(tt.args.callbacks...)
			if len(g.onMemberEvent[MemberEventRemovedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onMemberEvent[MemberEventRemovedAction]), tt.want)
			}
		})
	}
}

func TestHandleMemberEventRemoved(t *testing.T) {
	action := MemberEventRemovedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.MemberEvent
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
				eventName:  "member",
				event:      &github.MemberEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "member",
				event:      &github.MemberEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "member",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "member",
				event:      &github.MemberEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "member",
				event:      &github.MemberEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "member",
				event:      &github.MemberEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnMemberEventRemoved(func(deliveryID string, eventName string, event *github.MemberEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleMemberEventRemoved(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleMemberEventRemoved() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnMemberEventEdited(t *testing.T) {
	type args struct {
		callbacks []MemberEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single MemberEventHandleFunc",
			args: args{
				callbacks: []MemberEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MemberEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple MemberEventHandleFunc",
			args: args{
				callbacks: []MemberEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MemberEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.MemberEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnMemberEventEdited(tt.args.callbacks...)
			if len(g.onMemberEvent[MemberEventEditedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onMemberEvent[MemberEventEditedAction]))
			}
		})
	}
}

func TestSetOnMemberEventEdited(t *testing.T) {
	type args struct {
		callbacks []MemberEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single MemberEventHandleFunc",
			args: args{
				[]MemberEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MemberEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple MemberEventHandleFuncs",
			args: args{
				[]MemberEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MemberEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.MemberEvent) error {
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
			g.SetOnMemberEventEdited(func(deliveryID string, eventName string, event *github.MemberEvent) error {
				return nil
			})
			g.SetOnMemberEventEdited(tt.args.callbacks...)
			if len(g.onMemberEvent[MemberEventEditedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onMemberEvent[MemberEventEditedAction]), tt.want)
			}
		})
	}
}

func TestHandleMemberEventEdited(t *testing.T) {
	action := MemberEventEditedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.MemberEvent
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
				eventName:  "member",
				event:      &github.MemberEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "member",
				event:      &github.MemberEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "member",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "member",
				event:      &github.MemberEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "member",
				event:      &github.MemberEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "member",
				event:      &github.MemberEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnMemberEventEdited(func(deliveryID string, eventName string, event *github.MemberEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleMemberEventEdited(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleMemberEventEdited() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMemberEvent(t *testing.T) {
	type fields struct {
		handler *EventHandler
	}
	type args struct {
		deliveryID string
		eventName  string
		event      *github.MemberEvent
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "must trigger MemberEventAny with unknown event action",
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
					onMemberEvent: map[string][]MemberEventHandleFunc{
						MemberEventAnyAction: {
							func(deliveryID string, eventName string, event *github.MemberEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  MemberEvent,

				event: &github.MemberEvent{Action: ptrString("unknown")},
			},
			wantErr: false,
		},

		{
			name: "must trigger MemberEventAdded",
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
					onMemberEvent: map[string][]MemberEventHandleFunc{
						MemberEventAnyAction: {
							func(deliveryID string, eventName string, event *github.MemberEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						MemberEventAddedAction: {
							func(deliveryID string, eventName string, event *github.MemberEvent) error {
								t.Logf("%s action called", MemberEventAddedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "member",
				event:      &github.MemberEvent{Action: ptrString(MemberEventAddedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail MemberEventAdded with empty action",
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
					onMemberEvent: map[string][]MemberEventHandleFunc{
						MemberEventAnyAction: {
							func(deliveryID string, eventName string, event *github.MemberEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						MemberEventAddedAction: {
							func(deliveryID string, eventName string, event *github.MemberEvent) error {
								t.Logf("%s action called", MemberEventAddedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "member",
				event:      &github.MemberEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail MemberEventAdded with nil action",
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
					onMemberEvent: map[string][]MemberEventHandleFunc{
						MemberEventAnyAction: {
							func(deliveryID string, eventName string, event *github.MemberEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						MemberEventAddedAction: {
							func(deliveryID string, eventName string, event *github.MemberEvent) error {
								t.Logf("%s action called", MemberEventAddedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "member",
				event:      &github.MemberEvent{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger MemberEventRemoved",
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
					onMemberEvent: map[string][]MemberEventHandleFunc{
						MemberEventAnyAction: {
							func(deliveryID string, eventName string, event *github.MemberEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						MemberEventRemovedAction: {
							func(deliveryID string, eventName string, event *github.MemberEvent) error {
								t.Logf("%s action called", MemberEventRemovedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "member",
				event:      &github.MemberEvent{Action: ptrString(MemberEventRemovedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail MemberEventRemoved with empty action",
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
					onMemberEvent: map[string][]MemberEventHandleFunc{
						MemberEventAnyAction: {
							func(deliveryID string, eventName string, event *github.MemberEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						MemberEventRemovedAction: {
							func(deliveryID string, eventName string, event *github.MemberEvent) error {
								t.Logf("%s action called", MemberEventRemovedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "member",
				event:      &github.MemberEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail MemberEventRemoved with nil action",
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
					onMemberEvent: map[string][]MemberEventHandleFunc{
						MemberEventAnyAction: {
							func(deliveryID string, eventName string, event *github.MemberEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						MemberEventRemovedAction: {
							func(deliveryID string, eventName string, event *github.MemberEvent) error {
								t.Logf("%s action called", MemberEventRemovedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "member",
				event:      &github.MemberEvent{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger MemberEventEdited",
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
					onMemberEvent: map[string][]MemberEventHandleFunc{
						MemberEventAnyAction: {
							func(deliveryID string, eventName string, event *github.MemberEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						MemberEventEditedAction: {
							func(deliveryID string, eventName string, event *github.MemberEvent) error {
								t.Logf("%s action called", MemberEventEditedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "member",
				event:      &github.MemberEvent{Action: ptrString(MemberEventEditedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail MemberEventEdited with empty action",
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
					onMemberEvent: map[string][]MemberEventHandleFunc{
						MemberEventAnyAction: {
							func(deliveryID string, eventName string, event *github.MemberEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						MemberEventEditedAction: {
							func(deliveryID string, eventName string, event *github.MemberEvent) error {
								t.Logf("%s action called", MemberEventEditedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "member",
				event:      &github.MemberEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail MemberEventEdited with nil action",
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
					onMemberEvent: map[string][]MemberEventHandleFunc{
						MemberEventAnyAction: {
							func(deliveryID string, eventName string, event *github.MemberEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						MemberEventEditedAction: {
							func(deliveryID string, eventName string, event *github.MemberEvent) error {
								t.Logf("%s action called", MemberEventEditedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "member",
				event:      &github.MemberEvent{Action: nil},
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
			if err := g.MemberEvent(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("MemberEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
