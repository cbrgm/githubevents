package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"errors"
	"github.com/google/go-github/v43/github"
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
			if len(g.onMembershipEvent["*"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onMembershipEvent["*"]))
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
			if len(g.onMembershipEvent["*"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onMembershipEvent["*"]), tt.want)
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
			if len(g.onMembershipEvent["added"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onMembershipEvent["added"]))
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
			if len(g.onMembershipEvent["added"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onMembershipEvent["added"]), tt.want)
			}
		})
	}
}

func TestHandleMembershipEventAdded(t *testing.T) {
	action := "added"

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
			if len(g.onMembershipEvent["removed"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onMembershipEvent["removed"]))
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
			if len(g.onMembershipEvent["removed"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onMembershipEvent["removed"]), tt.want)
			}
		})
	}
}

func TestHandleMembershipEventRemoved(t *testing.T) {
	action := "removed"

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
