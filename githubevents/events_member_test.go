package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"errors"
	"github.com/google/go-github/v43/github"
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
			if len(g.onMemberEvent["*"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onMemberEvent["*"]))
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
			if len(g.onMemberEvent["*"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onMemberEvent["*"]), tt.want)
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
			if len(g.onMemberEvent["added"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onMemberEvent["added"]))
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
			if len(g.onMemberEvent["added"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onMemberEvent["added"]), tt.want)
			}
		})
	}
}

func TestHandleMemberEventAdded(t *testing.T) {
	action := "added"

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
			if len(g.onMemberEvent["removed"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onMemberEvent["removed"]))
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
			if len(g.onMemberEvent["removed"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onMemberEvent["removed"]), tt.want)
			}
		})
	}
}

func TestHandleMemberEventRemoved(t *testing.T) {
	action := "removed"

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
			if len(g.onMemberEvent["edited"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onMemberEvent["edited"]))
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
			if len(g.onMemberEvent["edited"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onMemberEvent["edited"]), tt.want)
			}
		})
	}
}

func TestHandleMemberEventEdited(t *testing.T) {
	action := "edited"

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
