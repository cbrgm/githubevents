package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"errors"
	"github.com/google/go-github/v43/github"
	"testing"
)

func TestOnOrganizationEventAny(t *testing.T) {
	type args struct {
		callbacks []OrganizationEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single OrganizationEventHandleFunc",
			args: args{
				[]OrganizationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple OrganizationEventHandleFuncs",
			args: args{
				[]OrganizationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnOrganizationEventAny(tt.args.callbacks...)
			if len(g.onOrganizationEvent["*"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onOrganizationEvent["*"]))
			}
		})
	}
}

func TestSetOnOrganizationEventAny(t *testing.T) {
	type args struct {
		callbacks []OrganizationEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single OrganizationEventHandleFunc",
			args: args{
				[]OrganizationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple OrganizationEventHandleFuncs",
			args: args{
				[]OrganizationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
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
			g.SetOnOrganizationEventAny(func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
				return nil
			})
			g.SetOnOrganizationEventAny(tt.args.callbacks...)
			if len(g.onOrganizationEvent["*"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onOrganizationEvent["*"]), tt.want)
			}
		})
	}
}

func TestHandleOrganizationEventAny(t *testing.T) {

	action := "*"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.OrganizationEvent
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
				eventName:  "organization",

				event: &github.OrganizationEvent{Action: &action},

				fail: false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "organization",

				event: &github.OrganizationEvent{Action: &action},

				fail: true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "organization",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnOrganizationEventAny(func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleOrganizationEventAny(tt.args.deliveryID, tt.args.deliveryID, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("TestHandleOrganizationEventAny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnOrganizationEventDeleted(t *testing.T) {
	type args struct {
		callbacks []OrganizationEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single OrganizationEventHandleFunc",
			args: args{
				callbacks: []OrganizationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple OrganizationEventHandleFunc",
			args: args{
				callbacks: []OrganizationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnOrganizationEventDeleted(tt.args.callbacks...)
			if len(g.onOrganizationEvent["deleted"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onOrganizationEvent["deleted"]))
			}
		})
	}
}

func TestSetOnOrganizationEventDeleted(t *testing.T) {
	type args struct {
		callbacks []OrganizationEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single OrganizationEventHandleFunc",
			args: args{
				[]OrganizationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple OrganizationEventHandleFuncs",
			args: args{
				[]OrganizationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
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
			g.SetOnOrganizationEventDeleted(func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
				return nil
			})
			g.SetOnOrganizationEventDeleted(tt.args.callbacks...)
			if len(g.onOrganizationEvent["deleted"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onOrganizationEvent["deleted"]), tt.want)
			}
		})
	}
}

func TestHandleOrganizationEventDeleted(t *testing.T) {
	action := "deleted"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.OrganizationEvent
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
				eventName:  "organization",
				event:      &github.OrganizationEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "organization",
				event:      &github.OrganizationEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "organization",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "organization",
				event:      &github.OrganizationEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "organization",
				event:      &github.OrganizationEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "organization",
				event:      &github.OrganizationEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnOrganizationEventDeleted(func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleOrganizationEventDeleted(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleOrganizationEventDeleted() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnOrganizationEventRenamed(t *testing.T) {
	type args struct {
		callbacks []OrganizationEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single OrganizationEventHandleFunc",
			args: args{
				callbacks: []OrganizationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple OrganizationEventHandleFunc",
			args: args{
				callbacks: []OrganizationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnOrganizationEventRenamed(tt.args.callbacks...)
			if len(g.onOrganizationEvent["renamed"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onOrganizationEvent["renamed"]))
			}
		})
	}
}

func TestSetOnOrganizationEventRenamed(t *testing.T) {
	type args struct {
		callbacks []OrganizationEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single OrganizationEventHandleFunc",
			args: args{
				[]OrganizationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple OrganizationEventHandleFuncs",
			args: args{
				[]OrganizationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
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
			g.SetOnOrganizationEventRenamed(func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
				return nil
			})
			g.SetOnOrganizationEventRenamed(tt.args.callbacks...)
			if len(g.onOrganizationEvent["renamed"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onOrganizationEvent["renamed"]), tt.want)
			}
		})
	}
}

func TestHandleOrganizationEventRenamed(t *testing.T) {
	action := "renamed"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.OrganizationEvent
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
				eventName:  "organization",
				event:      &github.OrganizationEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "organization",
				event:      &github.OrganizationEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "organization",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "organization",
				event:      &github.OrganizationEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "organization",
				event:      &github.OrganizationEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "organization",
				event:      &github.OrganizationEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnOrganizationEventRenamed(func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleOrganizationEventRenamed(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleOrganizationEventRenamed() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnOrganizationEventMemberAdded(t *testing.T) {
	type args struct {
		callbacks []OrganizationEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single OrganizationEventHandleFunc",
			args: args{
				callbacks: []OrganizationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple OrganizationEventHandleFunc",
			args: args{
				callbacks: []OrganizationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnOrganizationEventMemberAdded(tt.args.callbacks...)
			if len(g.onOrganizationEvent["member_added"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onOrganizationEvent["member_added"]))
			}
		})
	}
}

func TestSetOnOrganizationEventMemberAdded(t *testing.T) {
	type args struct {
		callbacks []OrganizationEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single OrganizationEventHandleFunc",
			args: args{
				[]OrganizationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple OrganizationEventHandleFuncs",
			args: args{
				[]OrganizationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
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
			g.SetOnOrganizationEventMemberAdded(func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
				return nil
			})
			g.SetOnOrganizationEventMemberAdded(tt.args.callbacks...)
			if len(g.onOrganizationEvent["member_added"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onOrganizationEvent["member_added"]), tt.want)
			}
		})
	}
}

func TestHandleOrganizationEventMemberAdded(t *testing.T) {
	action := "member_added"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.OrganizationEvent
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
				eventName:  "organization",
				event:      &github.OrganizationEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "organization",
				event:      &github.OrganizationEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "organization",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "organization",
				event:      &github.OrganizationEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "organization",
				event:      &github.OrganizationEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "organization",
				event:      &github.OrganizationEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnOrganizationEventMemberAdded(func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleOrganizationEventMemberAdded(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleOrganizationEventMemberAdded() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnOrganizationEventMemberRemoved(t *testing.T) {
	type args struct {
		callbacks []OrganizationEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single OrganizationEventHandleFunc",
			args: args{
				callbacks: []OrganizationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple OrganizationEventHandleFunc",
			args: args{
				callbacks: []OrganizationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnOrganizationEventMemberRemoved(tt.args.callbacks...)
			if len(g.onOrganizationEvent["member_removed"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onOrganizationEvent["member_removed"]))
			}
		})
	}
}

func TestSetOnOrganizationEventMemberRemoved(t *testing.T) {
	type args struct {
		callbacks []OrganizationEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single OrganizationEventHandleFunc",
			args: args{
				[]OrganizationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple OrganizationEventHandleFuncs",
			args: args{
				[]OrganizationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
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
			g.SetOnOrganizationEventMemberRemoved(func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
				return nil
			})
			g.SetOnOrganizationEventMemberRemoved(tt.args.callbacks...)
			if len(g.onOrganizationEvent["member_removed"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onOrganizationEvent["member_removed"]), tt.want)
			}
		})
	}
}

func TestHandleOrganizationEventMemberRemoved(t *testing.T) {
	action := "member_removed"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.OrganizationEvent
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
				eventName:  "organization",
				event:      &github.OrganizationEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "organization",
				event:      &github.OrganizationEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "organization",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "organization",
				event:      &github.OrganizationEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "organization",
				event:      &github.OrganizationEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "organization",
				event:      &github.OrganizationEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnOrganizationEventMemberRemoved(func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleOrganizationEventMemberRemoved(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleOrganizationEventMemberRemoved() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnOrganizationEventMemberInvited(t *testing.T) {
	type args struct {
		callbacks []OrganizationEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single OrganizationEventHandleFunc",
			args: args{
				callbacks: []OrganizationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple OrganizationEventHandleFunc",
			args: args{
				callbacks: []OrganizationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnOrganizationEventMemberInvited(tt.args.callbacks...)
			if len(g.onOrganizationEvent["member_invited"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onOrganizationEvent["member_invited"]))
			}
		})
	}
}

func TestSetOnOrganizationEventMemberInvited(t *testing.T) {
	type args struct {
		callbacks []OrganizationEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single OrganizationEventHandleFunc",
			args: args{
				[]OrganizationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple OrganizationEventHandleFuncs",
			args: args{
				[]OrganizationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
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
			g.SetOnOrganizationEventMemberInvited(func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
				return nil
			})
			g.SetOnOrganizationEventMemberInvited(tt.args.callbacks...)
			if len(g.onOrganizationEvent["member_invited"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onOrganizationEvent["member_invited"]), tt.want)
			}
		})
	}
}

func TestHandleOrganizationEventMemberInvited(t *testing.T) {
	action := "member_invited"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.OrganizationEvent
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
				eventName:  "organization",
				event:      &github.OrganizationEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "organization",
				event:      &github.OrganizationEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "organization",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "organization",
				event:      &github.OrganizationEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "organization",
				event:      &github.OrganizationEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "organization",
				event:      &github.OrganizationEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnOrganizationEventMemberInvited(func(deliveryID string, eventName string, event *github.OrganizationEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleOrganizationEventMemberInvited(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleOrganizationEventMemberInvited() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
