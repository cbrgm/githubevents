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

func TestOnCheckRunEventAny(t *testing.T) {
	type args struct {
		callbacks []CheckRunEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single CheckRunEventHandleFunc",
			args: args{
				[]CheckRunEventHandleFunc{
					func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple CheckRunEventHandleFuncs",
			args: args{
				[]CheckRunEventHandleFunc{
					func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnCheckRunEventAny(tt.args.callbacks...)
			if len(g.onCheckRunEvent[CheckRunEventAnyAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onCheckRunEvent[CheckRunEventAnyAction]))
			}
		})
	}
}

func TestSetOnCheckRunEventAny(t *testing.T) {
	type args struct {
		callbacks []CheckRunEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single CheckRunEventHandleFunc",
			args: args{
				[]CheckRunEventHandleFunc{
					func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple CheckRunEventHandleFuncs",
			args: args{
				[]CheckRunEventHandleFunc{
					func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
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
			g.SetOnCheckRunEventAny(func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
				return nil
			})
			g.SetOnCheckRunEventAny(tt.args.callbacks...)
			if len(g.onCheckRunEvent[CheckRunEventAnyAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onCheckRunEvent[CheckRunEventAnyAction]), tt.want)
			}
		})
	}
}

func TestHandleCheckRunEventAny(t *testing.T) {

	action := "*"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.CheckRunEvent
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
				eventName:  "check_run",

				event: &github.CheckRunEvent{Action: &action},

				fail: false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "check_run",

				event: &github.CheckRunEvent{Action: &action},

				fail: true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "check_run",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnCheckRunEventAny(func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleCheckRunEventAny(tt.args.deliveryID, tt.args.deliveryID, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("TestHandleCheckRunEventAny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnCheckRunEventCreated(t *testing.T) {
	type args struct {
		callbacks []CheckRunEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single CheckRunEventHandleFunc",
			args: args{
				callbacks: []CheckRunEventHandleFunc{
					func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple CheckRunEventHandleFunc",
			args: args{
				callbacks: []CheckRunEventHandleFunc{
					func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnCheckRunEventCreated(tt.args.callbacks...)
			if len(g.onCheckRunEvent[CheckRunEventCreatedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onCheckRunEvent[CheckRunEventCreatedAction]))
			}
		})
	}
}

func TestSetOnCheckRunEventCreated(t *testing.T) {
	type args struct {
		callbacks []CheckRunEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single CheckRunEventHandleFunc",
			args: args{
				[]CheckRunEventHandleFunc{
					func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple CheckRunEventHandleFuncs",
			args: args{
				[]CheckRunEventHandleFunc{
					func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
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
			g.SetOnCheckRunEventCreated(func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
				return nil
			})
			g.SetOnCheckRunEventCreated(tt.args.callbacks...)
			if len(g.onCheckRunEvent[CheckRunEventCreatedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onCheckRunEvent[CheckRunEventCreatedAction]), tt.want)
			}
		})
	}
}

func TestHandleCheckRunEventCreated(t *testing.T) {
	action := CheckRunEventCreatedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.CheckRunEvent
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
				eventName:  "check_run",
				event:      &github.CheckRunEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "check_run",
				event:      &github.CheckRunEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "check_run",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "check_run",
				event:      &github.CheckRunEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "check_run",
				event:      &github.CheckRunEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "check_run",
				event:      &github.CheckRunEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnCheckRunEventCreated(func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleCheckRunEventCreated(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleCheckRunEventCreated() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnCheckRunEventCompleted(t *testing.T) {
	type args struct {
		callbacks []CheckRunEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single CheckRunEventHandleFunc",
			args: args{
				callbacks: []CheckRunEventHandleFunc{
					func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple CheckRunEventHandleFunc",
			args: args{
				callbacks: []CheckRunEventHandleFunc{
					func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnCheckRunEventCompleted(tt.args.callbacks...)
			if len(g.onCheckRunEvent[CheckRunEventCompletedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onCheckRunEvent[CheckRunEventCompletedAction]))
			}
		})
	}
}

func TestSetOnCheckRunEventCompleted(t *testing.T) {
	type args struct {
		callbacks []CheckRunEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single CheckRunEventHandleFunc",
			args: args{
				[]CheckRunEventHandleFunc{
					func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple CheckRunEventHandleFuncs",
			args: args{
				[]CheckRunEventHandleFunc{
					func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
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
			g.SetOnCheckRunEventCompleted(func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
				return nil
			})
			g.SetOnCheckRunEventCompleted(tt.args.callbacks...)
			if len(g.onCheckRunEvent[CheckRunEventCompletedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onCheckRunEvent[CheckRunEventCompletedAction]), tt.want)
			}
		})
	}
}

func TestHandleCheckRunEventCompleted(t *testing.T) {
	action := CheckRunEventCompletedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.CheckRunEvent
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
				eventName:  "check_run",
				event:      &github.CheckRunEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "check_run",
				event:      &github.CheckRunEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "check_run",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "check_run",
				event:      &github.CheckRunEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "check_run",
				event:      &github.CheckRunEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "check_run",
				event:      &github.CheckRunEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnCheckRunEventCompleted(func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleCheckRunEventCompleted(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleCheckRunEventCompleted() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnCheckRunEventReRequested(t *testing.T) {
	type args struct {
		callbacks []CheckRunEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single CheckRunEventHandleFunc",
			args: args{
				callbacks: []CheckRunEventHandleFunc{
					func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple CheckRunEventHandleFunc",
			args: args{
				callbacks: []CheckRunEventHandleFunc{
					func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnCheckRunEventReRequested(tt.args.callbacks...)
			if len(g.onCheckRunEvent[CheckRunEventReRequestedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onCheckRunEvent[CheckRunEventReRequestedAction]))
			}
		})
	}
}

func TestSetOnCheckRunEventReRequested(t *testing.T) {
	type args struct {
		callbacks []CheckRunEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single CheckRunEventHandleFunc",
			args: args{
				[]CheckRunEventHandleFunc{
					func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple CheckRunEventHandleFuncs",
			args: args{
				[]CheckRunEventHandleFunc{
					func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
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
			g.SetOnCheckRunEventReRequested(func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
				return nil
			})
			g.SetOnCheckRunEventReRequested(tt.args.callbacks...)
			if len(g.onCheckRunEvent[CheckRunEventReRequestedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onCheckRunEvent[CheckRunEventReRequestedAction]), tt.want)
			}
		})
	}
}

func TestHandleCheckRunEventReRequested(t *testing.T) {
	action := CheckRunEventReRequestedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.CheckRunEvent
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
				eventName:  "check_run",
				event:      &github.CheckRunEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "check_run",
				event:      &github.CheckRunEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "check_run",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "check_run",
				event:      &github.CheckRunEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "check_run",
				event:      &github.CheckRunEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "check_run",
				event:      &github.CheckRunEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnCheckRunEventReRequested(func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleCheckRunEventReRequested(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleCheckRunEventReRequested() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnCheckRunEventRequestAction(t *testing.T) {
	type args struct {
		callbacks []CheckRunEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single CheckRunEventHandleFunc",
			args: args{
				callbacks: []CheckRunEventHandleFunc{
					func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple CheckRunEventHandleFunc",
			args: args{
				callbacks: []CheckRunEventHandleFunc{
					func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnCheckRunEventRequestAction(tt.args.callbacks...)
			if len(g.onCheckRunEvent[CheckRunEventRequestActionAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onCheckRunEvent[CheckRunEventRequestActionAction]))
			}
		})
	}
}

func TestSetOnCheckRunEventRequestAction(t *testing.T) {
	type args struct {
		callbacks []CheckRunEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single CheckRunEventHandleFunc",
			args: args{
				[]CheckRunEventHandleFunc{
					func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple CheckRunEventHandleFuncs",
			args: args{
				[]CheckRunEventHandleFunc{
					func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
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
			g.SetOnCheckRunEventRequestAction(func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
				return nil
			})
			g.SetOnCheckRunEventRequestAction(tt.args.callbacks...)
			if len(g.onCheckRunEvent[CheckRunEventRequestActionAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onCheckRunEvent[CheckRunEventRequestActionAction]), tt.want)
			}
		})
	}
}

func TestHandleCheckRunEventRequestAction(t *testing.T) {
	action := CheckRunEventRequestActionAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.CheckRunEvent
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
				eventName:  "check_run",
				event:      &github.CheckRunEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "check_run",
				event:      &github.CheckRunEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "check_run",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "check_run",
				event:      &github.CheckRunEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "check_run",
				event:      &github.CheckRunEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "check_run",
				event:      &github.CheckRunEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnCheckRunEventRequestAction(func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleCheckRunEventRequestAction(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleCheckRunEventRequestAction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCheckRunEvent(t *testing.T) {
	type fields struct {
		handler *EventHandler
	}
	type args struct {
		deliveryID string
		eventName  string
		event      *github.CheckRunEvent
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "must trigger CheckRunEventAny with unknown event action",
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
					onCheckRunEvent: map[string][]CheckRunEventHandleFunc{
						CheckRunEventAnyAction: {
							func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  CheckRunEvent,

				event: &github.CheckRunEvent{Action: ptrString("unknown")},
			},
			wantErr: false,
		},

		{
			name: "must trigger CheckRunEventCreated",
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
					onCheckRunEvent: map[string][]CheckRunEventHandleFunc{
						CheckRunEventAnyAction: {
							func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						CheckRunEventCreatedAction: {
							func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
								t.Logf("%s action called", CheckRunEventCreatedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "check_run",
				event:      &github.CheckRunEvent{Action: ptrString(CheckRunEventCreatedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail CheckRunEventCreated with empty action",
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
					onCheckRunEvent: map[string][]CheckRunEventHandleFunc{
						CheckRunEventAnyAction: {
							func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						CheckRunEventCreatedAction: {
							func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
								t.Logf("%s action called", CheckRunEventCreatedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "check_run",
				event:      &github.CheckRunEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail CheckRunEventCreated with nil action",
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
					onCheckRunEvent: map[string][]CheckRunEventHandleFunc{
						CheckRunEventAnyAction: {
							func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						CheckRunEventCreatedAction: {
							func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
								t.Logf("%s action called", CheckRunEventCreatedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "check_run",
				event:      &github.CheckRunEvent{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger CheckRunEventCompleted",
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
					onCheckRunEvent: map[string][]CheckRunEventHandleFunc{
						CheckRunEventAnyAction: {
							func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						CheckRunEventCompletedAction: {
							func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
								t.Logf("%s action called", CheckRunEventCompletedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "check_run",
				event:      &github.CheckRunEvent{Action: ptrString(CheckRunEventCompletedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail CheckRunEventCompleted with empty action",
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
					onCheckRunEvent: map[string][]CheckRunEventHandleFunc{
						CheckRunEventAnyAction: {
							func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						CheckRunEventCompletedAction: {
							func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
								t.Logf("%s action called", CheckRunEventCompletedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "check_run",
				event:      &github.CheckRunEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail CheckRunEventCompleted with nil action",
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
					onCheckRunEvent: map[string][]CheckRunEventHandleFunc{
						CheckRunEventAnyAction: {
							func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						CheckRunEventCompletedAction: {
							func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
								t.Logf("%s action called", CheckRunEventCompletedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "check_run",
				event:      &github.CheckRunEvent{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger CheckRunEventReRequested",
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
					onCheckRunEvent: map[string][]CheckRunEventHandleFunc{
						CheckRunEventAnyAction: {
							func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						CheckRunEventReRequestedAction: {
							func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
								t.Logf("%s action called", CheckRunEventReRequestedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "check_run",
				event:      &github.CheckRunEvent{Action: ptrString(CheckRunEventReRequestedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail CheckRunEventReRequested with empty action",
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
					onCheckRunEvent: map[string][]CheckRunEventHandleFunc{
						CheckRunEventAnyAction: {
							func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						CheckRunEventReRequestedAction: {
							func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
								t.Logf("%s action called", CheckRunEventReRequestedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "check_run",
				event:      &github.CheckRunEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail CheckRunEventReRequested with nil action",
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
					onCheckRunEvent: map[string][]CheckRunEventHandleFunc{
						CheckRunEventAnyAction: {
							func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						CheckRunEventReRequestedAction: {
							func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
								t.Logf("%s action called", CheckRunEventReRequestedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "check_run",
				event:      &github.CheckRunEvent{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger CheckRunEventRequestAction",
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
					onCheckRunEvent: map[string][]CheckRunEventHandleFunc{
						CheckRunEventAnyAction: {
							func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						CheckRunEventRequestActionAction: {
							func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
								t.Logf("%s action called", CheckRunEventRequestActionAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "check_run",
				event:      &github.CheckRunEvent{Action: ptrString(CheckRunEventRequestActionAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail CheckRunEventRequestAction with empty action",
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
					onCheckRunEvent: map[string][]CheckRunEventHandleFunc{
						CheckRunEventAnyAction: {
							func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						CheckRunEventRequestActionAction: {
							func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
								t.Logf("%s action called", CheckRunEventRequestActionAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "check_run",
				event:      &github.CheckRunEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail CheckRunEventRequestAction with nil action",
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
					onCheckRunEvent: map[string][]CheckRunEventHandleFunc{
						CheckRunEventAnyAction: {
							func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						CheckRunEventRequestActionAction: {
							func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
								t.Logf("%s action called", CheckRunEventRequestActionAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "check_run",
				event:      &github.CheckRunEvent{Action: nil},
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
			if err := g.CheckRunEvent(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("CheckRunEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
