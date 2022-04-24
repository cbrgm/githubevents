package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"errors"
	"github.com/google/go-github/v43/github"
	"testing"
)

func TestOnBranchProtectionRuleEventAny(t *testing.T) {
	type args struct {
		callbacks []BranchProtectionRuleEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single BranchProtectionRuleEventHandleFunc",
			args: args{
				[]BranchProtectionRuleEventHandleFunc{
					func(deliveryID string, eventName string, event *github.BranchProtectionRuleEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple BranchProtectionRuleEventHandleFuncs",
			args: args{
				[]BranchProtectionRuleEventHandleFunc{
					func(deliveryID string, eventName string, event *github.BranchProtectionRuleEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.BranchProtectionRuleEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnBranchProtectionRuleEventAny(tt.args.callbacks...)
			if len(g.onBranchProtectionRuleEvent["*"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onBranchProtectionRuleEvent["*"]))
			}
		})
	}
}

func TestSetOnBranchProtectionRuleEventAny(t *testing.T) {
	type args struct {
		callbacks []BranchProtectionRuleEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single BranchProtectionRuleEventHandleFunc",
			args: args{
				[]BranchProtectionRuleEventHandleFunc{
					func(deliveryID string, eventName string, event *github.BranchProtectionRuleEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple BranchProtectionRuleEventHandleFuncs",
			args: args{
				[]BranchProtectionRuleEventHandleFunc{
					func(deliveryID string, eventName string, event *github.BranchProtectionRuleEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.BranchProtectionRuleEvent) error {
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
			g.SetOnBranchProtectionRuleEventAny(func(deliveryID string, eventName string, event *github.BranchProtectionRuleEvent) error {
				return nil
			})
			g.SetOnBranchProtectionRuleEventAny(tt.args.callbacks...)
			if len(g.onBranchProtectionRuleEvent["*"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onBranchProtectionRuleEvent["*"]), tt.want)
			}
		})
	}
}

func TestHandleBranchProtectionRuleEventAny(t *testing.T) {

	action := "*"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.BranchProtectionRuleEvent
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
				eventName:  "branch_protection_rule",

				event: &github.BranchProtectionRuleEvent{Action: &action},

				fail: false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "branch_protection_rule",

				event: &github.BranchProtectionRuleEvent{Action: &action},

				fail: true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "branch_protection_rule",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnBranchProtectionRuleEventAny(func(deliveryID string, eventName string, event *github.BranchProtectionRuleEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleBranchProtectionRuleEventAny(tt.args.deliveryID, tt.args.deliveryID, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("TestHandleBranchProtectionRuleEventAny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnBranchProtectionRuleEventCreated(t *testing.T) {
	type args struct {
		callbacks []BranchProtectionRuleEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single BranchProtectionRuleEventHandleFunc",
			args: args{
				callbacks: []BranchProtectionRuleEventHandleFunc{
					func(deliveryID string, eventName string, event *github.BranchProtectionRuleEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple BranchProtectionRuleEventHandleFunc",
			args: args{
				callbacks: []BranchProtectionRuleEventHandleFunc{
					func(deliveryID string, eventName string, event *github.BranchProtectionRuleEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.BranchProtectionRuleEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnBranchProtectionRuleEventCreated(tt.args.callbacks...)
			if len(g.onBranchProtectionRuleEvent["created"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onBranchProtectionRuleEvent["created"]))
			}
		})
	}
}

func TestSetOnBranchProtectionRuleEventCreated(t *testing.T) {
	type args struct {
		callbacks []BranchProtectionRuleEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single BranchProtectionRuleEventHandleFunc",
			args: args{
				[]BranchProtectionRuleEventHandleFunc{
					func(deliveryID string, eventName string, event *github.BranchProtectionRuleEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple BranchProtectionRuleEventHandleFuncs",
			args: args{
				[]BranchProtectionRuleEventHandleFunc{
					func(deliveryID string, eventName string, event *github.BranchProtectionRuleEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.BranchProtectionRuleEvent) error {
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
			g.SetOnBranchProtectionRuleEventCreated(func(deliveryID string, eventName string, event *github.BranchProtectionRuleEvent) error {
				return nil
			})
			g.SetOnBranchProtectionRuleEventCreated(tt.args.callbacks...)
			if len(g.onBranchProtectionRuleEvent["created"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onBranchProtectionRuleEvent["created"]), tt.want)
			}
		})
	}
}

func TestHandleBranchProtectionRuleEventCreated(t *testing.T) {
	action := "created"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.BranchProtectionRuleEvent
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
				eventName:  "branch_protection_rule",
				event:      &github.BranchProtectionRuleEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "branch_protection_rule",
				event:      &github.BranchProtectionRuleEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "branch_protection_rule",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "branch_protection_rule",
				event:      &github.BranchProtectionRuleEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "branch_protection_rule",
				event:      &github.BranchProtectionRuleEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "branch_protection_rule",
				event:      &github.BranchProtectionRuleEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnBranchProtectionRuleEventCreated(func(deliveryID string, eventName string, event *github.BranchProtectionRuleEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleBranchProtectionRuleEventCreated(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleBranchProtectionRuleEventCreated() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnBranchProtectionRuleEventEdited(t *testing.T) {
	type args struct {
		callbacks []BranchProtectionRuleEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single BranchProtectionRuleEventHandleFunc",
			args: args{
				callbacks: []BranchProtectionRuleEventHandleFunc{
					func(deliveryID string, eventName string, event *github.BranchProtectionRuleEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple BranchProtectionRuleEventHandleFunc",
			args: args{
				callbacks: []BranchProtectionRuleEventHandleFunc{
					func(deliveryID string, eventName string, event *github.BranchProtectionRuleEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.BranchProtectionRuleEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnBranchProtectionRuleEventEdited(tt.args.callbacks...)
			if len(g.onBranchProtectionRuleEvent["edited"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onBranchProtectionRuleEvent["edited"]))
			}
		})
	}
}

func TestSetOnBranchProtectionRuleEventEdited(t *testing.T) {
	type args struct {
		callbacks []BranchProtectionRuleEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single BranchProtectionRuleEventHandleFunc",
			args: args{
				[]BranchProtectionRuleEventHandleFunc{
					func(deliveryID string, eventName string, event *github.BranchProtectionRuleEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple BranchProtectionRuleEventHandleFuncs",
			args: args{
				[]BranchProtectionRuleEventHandleFunc{
					func(deliveryID string, eventName string, event *github.BranchProtectionRuleEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.BranchProtectionRuleEvent) error {
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
			g.SetOnBranchProtectionRuleEventEdited(func(deliveryID string, eventName string, event *github.BranchProtectionRuleEvent) error {
				return nil
			})
			g.SetOnBranchProtectionRuleEventEdited(tt.args.callbacks...)
			if len(g.onBranchProtectionRuleEvent["edited"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onBranchProtectionRuleEvent["edited"]), tt.want)
			}
		})
	}
}

func TestHandleBranchProtectionRuleEventEdited(t *testing.T) {
	action := "edited"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.BranchProtectionRuleEvent
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
				eventName:  "branch_protection_rule",
				event:      &github.BranchProtectionRuleEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "branch_protection_rule",
				event:      &github.BranchProtectionRuleEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "branch_protection_rule",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "branch_protection_rule",
				event:      &github.BranchProtectionRuleEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "branch_protection_rule",
				event:      &github.BranchProtectionRuleEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "branch_protection_rule",
				event:      &github.BranchProtectionRuleEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnBranchProtectionRuleEventEdited(func(deliveryID string, eventName string, event *github.BranchProtectionRuleEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleBranchProtectionRuleEventEdited(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleBranchProtectionRuleEventEdited() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnBranchProtectionRuleEventDeleted(t *testing.T) {
	type args struct {
		callbacks []BranchProtectionRuleEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single BranchProtectionRuleEventHandleFunc",
			args: args{
				callbacks: []BranchProtectionRuleEventHandleFunc{
					func(deliveryID string, eventName string, event *github.BranchProtectionRuleEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple BranchProtectionRuleEventHandleFunc",
			args: args{
				callbacks: []BranchProtectionRuleEventHandleFunc{
					func(deliveryID string, eventName string, event *github.BranchProtectionRuleEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.BranchProtectionRuleEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnBranchProtectionRuleEventDeleted(tt.args.callbacks...)
			if len(g.onBranchProtectionRuleEvent["deleted"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onBranchProtectionRuleEvent["deleted"]))
			}
		})
	}
}

func TestSetOnBranchProtectionRuleEventDeleted(t *testing.T) {
	type args struct {
		callbacks []BranchProtectionRuleEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single BranchProtectionRuleEventHandleFunc",
			args: args{
				[]BranchProtectionRuleEventHandleFunc{
					func(deliveryID string, eventName string, event *github.BranchProtectionRuleEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple BranchProtectionRuleEventHandleFuncs",
			args: args{
				[]BranchProtectionRuleEventHandleFunc{
					func(deliveryID string, eventName string, event *github.BranchProtectionRuleEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.BranchProtectionRuleEvent) error {
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
			g.SetOnBranchProtectionRuleEventDeleted(func(deliveryID string, eventName string, event *github.BranchProtectionRuleEvent) error {
				return nil
			})
			g.SetOnBranchProtectionRuleEventDeleted(tt.args.callbacks...)
			if len(g.onBranchProtectionRuleEvent["deleted"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onBranchProtectionRuleEvent["deleted"]), tt.want)
			}
		})
	}
}

func TestHandleBranchProtectionRuleEventDeleted(t *testing.T) {
	action := "deleted"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.BranchProtectionRuleEvent
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
				eventName:  "branch_protection_rule",
				event:      &github.BranchProtectionRuleEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "branch_protection_rule",
				event:      &github.BranchProtectionRuleEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "branch_protection_rule",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "branch_protection_rule",
				event:      &github.BranchProtectionRuleEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "branch_protection_rule",
				event:      &github.BranchProtectionRuleEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "branch_protection_rule",
				event:      &github.BranchProtectionRuleEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnBranchProtectionRuleEventDeleted(func(deliveryID string, eventName string, event *github.BranchProtectionRuleEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleBranchProtectionRuleEventDeleted(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleBranchProtectionRuleEventDeleted() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
