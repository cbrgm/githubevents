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

func TestOnOrgBlockEventAny(t *testing.T) {
	type args struct {
		callbacks []OrgBlockEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single OrgBlockEventHandleFunc",
			args: args{
				[]OrgBlockEventHandleFunc{
					func(deliveryID string, eventName string, event *github.OrgBlockEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple OrgBlockEventHandleFuncs",
			args: args{
				[]OrgBlockEventHandleFunc{
					func(deliveryID string, eventName string, event *github.OrgBlockEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.OrgBlockEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnOrgBlockEventAny(tt.args.callbacks...)
			if len(g.onOrgBlockEvent[OrgBlockEventAnyAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onOrgBlockEvent[OrgBlockEventAnyAction]))
			}
		})
	}
}

func TestSetOnOrgBlockEventAny(t *testing.T) {
	type args struct {
		callbacks []OrgBlockEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single OrgBlockEventHandleFunc",
			args: args{
				[]OrgBlockEventHandleFunc{
					func(deliveryID string, eventName string, event *github.OrgBlockEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple OrgBlockEventHandleFuncs",
			args: args{
				[]OrgBlockEventHandleFunc{
					func(deliveryID string, eventName string, event *github.OrgBlockEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.OrgBlockEvent) error {
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
			g.SetOnOrgBlockEventAny(func(deliveryID string, eventName string, event *github.OrgBlockEvent) error {
				return nil
			})
			g.SetOnOrgBlockEventAny(tt.args.callbacks...)
			if len(g.onOrgBlockEvent[OrgBlockEventAnyAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onOrgBlockEvent[OrgBlockEventAnyAction]), tt.want)
			}
		})
	}
}

func TestHandleOrgBlockEventAny(t *testing.T) {

	action := "*"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.OrgBlockEvent
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
				eventName:  "org_block",

				event: &github.OrgBlockEvent{Action: &action},

				fail: false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "org_block",

				event: &github.OrgBlockEvent{Action: &action},

				fail: true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "org_block",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnOrgBlockEventAny(func(deliveryID string, eventName string, event *github.OrgBlockEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleOrgBlockEventAny(tt.args.deliveryID, tt.args.deliveryID, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("TestHandleOrgBlockEventAny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnOrgBlockEventBlocked(t *testing.T) {
	type args struct {
		callbacks []OrgBlockEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single OrgBlockEventHandleFunc",
			args: args{
				callbacks: []OrgBlockEventHandleFunc{
					func(deliveryID string, eventName string, event *github.OrgBlockEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple OrgBlockEventHandleFunc",
			args: args{
				callbacks: []OrgBlockEventHandleFunc{
					func(deliveryID string, eventName string, event *github.OrgBlockEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.OrgBlockEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnOrgBlockEventBlocked(tt.args.callbacks...)
			if len(g.onOrgBlockEvent[OrgBlockEventBlockedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onOrgBlockEvent[OrgBlockEventBlockedAction]))
			}
		})
	}
}

func TestSetOnOrgBlockEventBlocked(t *testing.T) {
	type args struct {
		callbacks []OrgBlockEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single OrgBlockEventHandleFunc",
			args: args{
				[]OrgBlockEventHandleFunc{
					func(deliveryID string, eventName string, event *github.OrgBlockEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple OrgBlockEventHandleFuncs",
			args: args{
				[]OrgBlockEventHandleFunc{
					func(deliveryID string, eventName string, event *github.OrgBlockEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.OrgBlockEvent) error {
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
			g.SetOnOrgBlockEventBlocked(func(deliveryID string, eventName string, event *github.OrgBlockEvent) error {
				return nil
			})
			g.SetOnOrgBlockEventBlocked(tt.args.callbacks...)
			if len(g.onOrgBlockEvent[OrgBlockEventBlockedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onOrgBlockEvent[OrgBlockEventBlockedAction]), tt.want)
			}
		})
	}
}

func TestHandleOrgBlockEventBlocked(t *testing.T) {
	action := OrgBlockEventBlockedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.OrgBlockEvent
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
				eventName:  "org_block",
				event:      &github.OrgBlockEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "org_block",
				event:      &github.OrgBlockEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "org_block",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "org_block",
				event:      &github.OrgBlockEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "org_block",
				event:      &github.OrgBlockEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "org_block",
				event:      &github.OrgBlockEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnOrgBlockEventBlocked(func(deliveryID string, eventName string, event *github.OrgBlockEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleOrgBlockEventBlocked(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleOrgBlockEventBlocked() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnOrgBlockEventUnblocked(t *testing.T) {
	type args struct {
		callbacks []OrgBlockEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single OrgBlockEventHandleFunc",
			args: args{
				callbacks: []OrgBlockEventHandleFunc{
					func(deliveryID string, eventName string, event *github.OrgBlockEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple OrgBlockEventHandleFunc",
			args: args{
				callbacks: []OrgBlockEventHandleFunc{
					func(deliveryID string, eventName string, event *github.OrgBlockEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.OrgBlockEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnOrgBlockEventUnblocked(tt.args.callbacks...)
			if len(g.onOrgBlockEvent[OrgBlockEventUnblockedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onOrgBlockEvent[OrgBlockEventUnblockedAction]))
			}
		})
	}
}

func TestSetOnOrgBlockEventUnblocked(t *testing.T) {
	type args struct {
		callbacks []OrgBlockEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single OrgBlockEventHandleFunc",
			args: args{
				[]OrgBlockEventHandleFunc{
					func(deliveryID string, eventName string, event *github.OrgBlockEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple OrgBlockEventHandleFuncs",
			args: args{
				[]OrgBlockEventHandleFunc{
					func(deliveryID string, eventName string, event *github.OrgBlockEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.OrgBlockEvent) error {
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
			g.SetOnOrgBlockEventUnblocked(func(deliveryID string, eventName string, event *github.OrgBlockEvent) error {
				return nil
			})
			g.SetOnOrgBlockEventUnblocked(tt.args.callbacks...)
			if len(g.onOrgBlockEvent[OrgBlockEventUnblockedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onOrgBlockEvent[OrgBlockEventUnblockedAction]), tt.want)
			}
		})
	}
}

func TestHandleOrgBlockEventUnblocked(t *testing.T) {
	action := OrgBlockEventUnblockedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.OrgBlockEvent
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
				eventName:  "org_block",
				event:      &github.OrgBlockEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "org_block",
				event:      &github.OrgBlockEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "org_block",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "org_block",
				event:      &github.OrgBlockEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "org_block",
				event:      &github.OrgBlockEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "org_block",
				event:      &github.OrgBlockEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnOrgBlockEventUnblocked(func(deliveryID string, eventName string, event *github.OrgBlockEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleOrgBlockEventUnblocked(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleOrgBlockEventUnblocked() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOrgBlockEvent(t *testing.T) {
	type fields struct {
		handler *EventHandler
	}
	type args struct {
		deliveryID string
		eventName  string
		event      *github.OrgBlockEvent
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "must trigger OrgBlockEventAny with unknown event action",
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
					onOrgBlockEvent: map[string][]OrgBlockEventHandleFunc{
						OrgBlockEventAnyAction: {
							func(deliveryID string, eventName string, event *github.OrgBlockEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  OrgBlockEvent,

				event: &github.OrgBlockEvent{Action: ptrString("unknown")},
			},
			wantErr: false,
		},

		{
			name: "must trigger OrgBlockEventBlocked",
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
					onOrgBlockEvent: map[string][]OrgBlockEventHandleFunc{
						OrgBlockEventAnyAction: {
							func(deliveryID string, eventName string, event *github.OrgBlockEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						OrgBlockEventBlockedAction: {
							func(deliveryID string, eventName string, event *github.OrgBlockEvent) error {
								t.Logf("%s action called", OrgBlockEventBlockedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "org_block",
				event:      &github.OrgBlockEvent{Action: ptrString(OrgBlockEventBlockedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail OrgBlockEventBlocked with empty action",
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
					onOrgBlockEvent: map[string][]OrgBlockEventHandleFunc{
						OrgBlockEventAnyAction: {
							func(deliveryID string, eventName string, event *github.OrgBlockEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						OrgBlockEventBlockedAction: {
							func(deliveryID string, eventName string, event *github.OrgBlockEvent) error {
								t.Logf("%s action called", OrgBlockEventBlockedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "org_block",
				event:      &github.OrgBlockEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail OrgBlockEventBlocked with nil action",
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
					onOrgBlockEvent: map[string][]OrgBlockEventHandleFunc{
						OrgBlockEventAnyAction: {
							func(deliveryID string, eventName string, event *github.OrgBlockEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						OrgBlockEventBlockedAction: {
							func(deliveryID string, eventName string, event *github.OrgBlockEvent) error {
								t.Logf("%s action called", OrgBlockEventBlockedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "org_block",
				event:      &github.OrgBlockEvent{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger OrgBlockEventUnblocked",
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
					onOrgBlockEvent: map[string][]OrgBlockEventHandleFunc{
						OrgBlockEventAnyAction: {
							func(deliveryID string, eventName string, event *github.OrgBlockEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						OrgBlockEventUnblockedAction: {
							func(deliveryID string, eventName string, event *github.OrgBlockEvent) error {
								t.Logf("%s action called", OrgBlockEventUnblockedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "org_block",
				event:      &github.OrgBlockEvent{Action: ptrString(OrgBlockEventUnblockedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail OrgBlockEventUnblocked with empty action",
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
					onOrgBlockEvent: map[string][]OrgBlockEventHandleFunc{
						OrgBlockEventAnyAction: {
							func(deliveryID string, eventName string, event *github.OrgBlockEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						OrgBlockEventUnblockedAction: {
							func(deliveryID string, eventName string, event *github.OrgBlockEvent) error {
								t.Logf("%s action called", OrgBlockEventUnblockedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "org_block",
				event:      &github.OrgBlockEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail OrgBlockEventUnblocked with nil action",
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
					onOrgBlockEvent: map[string][]OrgBlockEventHandleFunc{
						OrgBlockEventAnyAction: {
							func(deliveryID string, eventName string, event *github.OrgBlockEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						OrgBlockEventUnblockedAction: {
							func(deliveryID string, eventName string, event *github.OrgBlockEvent) error {
								t.Logf("%s action called", OrgBlockEventUnblockedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "org_block",
				event:      &github.OrgBlockEvent{Action: nil},
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
			if err := g.OrgBlockEvent(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("OrgBlockEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
