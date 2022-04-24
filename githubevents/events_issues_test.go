package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"errors"
	"github.com/google/go-github/v43/github"
	"testing"
)

func TestOnIssuesEventAny(t *testing.T) {
	type args struct {
		callbacks []IssuesEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single IssuesEventHandleFunc",
			args: args{
				[]IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple IssuesEventHandleFuncs",
			args: args{
				[]IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnIssuesEventAny(tt.args.callbacks...)
			if len(g.onIssuesEvent["*"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onIssuesEvent["*"]))
			}
		})
	}
}

func TestSetOnIssuesEventAny(t *testing.T) {
	type args struct {
		callbacks []IssuesEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single IssuesEventHandleFunc",
			args: args{
				[]IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple IssuesEventHandleFuncs",
			args: args{
				[]IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
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
			g.SetOnIssuesEventAny(func(deliveryID string, eventName string, event *github.IssuesEvent) error {
				return nil
			})
			g.SetOnIssuesEventAny(tt.args.callbacks...)
			if len(g.onIssuesEvent["*"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onIssuesEvent["*"]), tt.want)
			}
		})
	}
}

func TestHandleIssuesEventAny(t *testing.T) {

	action := "*"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.IssuesEvent
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
				eventName:  "issues",

				event: &github.IssuesEvent{Action: &action},

				fail: false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "issues",

				event: &github.IssuesEvent{Action: &action},

				fail: true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnIssuesEventAny(func(deliveryID string, eventName string, event *github.IssuesEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleIssuesEventAny(tt.args.deliveryID, tt.args.deliveryID, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("TestHandleIssuesEventAny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnIssuesEventOpened(t *testing.T) {
	type args struct {
		callbacks []IssuesEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single IssuesEventHandleFunc",
			args: args{
				callbacks: []IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple IssuesEventHandleFunc",
			args: args{
				callbacks: []IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnIssuesEventOpened(tt.args.callbacks...)
			if len(g.onIssuesEvent["opened"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onIssuesEvent["opened"]))
			}
		})
	}
}

func TestSetOnIssuesEventOpened(t *testing.T) {
	type args struct {
		callbacks []IssuesEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single IssuesEventHandleFunc",
			args: args{
				[]IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple IssuesEventHandleFuncs",
			args: args{
				[]IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
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
			g.SetOnIssuesEventOpened(func(deliveryID string, eventName string, event *github.IssuesEvent) error {
				return nil
			})
			g.SetOnIssuesEventOpened(tt.args.callbacks...)
			if len(g.onIssuesEvent["opened"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onIssuesEvent["opened"]), tt.want)
			}
		})
	}
}

func TestHandleIssuesEventOpened(t *testing.T) {
	action := "opened"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.IssuesEvent
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
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnIssuesEventOpened(func(deliveryID string, eventName string, event *github.IssuesEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleIssuesEventOpened(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleIssuesEventOpened() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnIssuesEventEdited(t *testing.T) {
	type args struct {
		callbacks []IssuesEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single IssuesEventHandleFunc",
			args: args{
				callbacks: []IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple IssuesEventHandleFunc",
			args: args{
				callbacks: []IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnIssuesEventEdited(tt.args.callbacks...)
			if len(g.onIssuesEvent["edited"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onIssuesEvent["edited"]))
			}
		})
	}
}

func TestSetOnIssuesEventEdited(t *testing.T) {
	type args struct {
		callbacks []IssuesEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single IssuesEventHandleFunc",
			args: args{
				[]IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple IssuesEventHandleFuncs",
			args: args{
				[]IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
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
			g.SetOnIssuesEventEdited(func(deliveryID string, eventName string, event *github.IssuesEvent) error {
				return nil
			})
			g.SetOnIssuesEventEdited(tt.args.callbacks...)
			if len(g.onIssuesEvent["edited"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onIssuesEvent["edited"]), tt.want)
			}
		})
	}
}

func TestHandleIssuesEventEdited(t *testing.T) {
	action := "edited"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.IssuesEvent
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
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnIssuesEventEdited(func(deliveryID string, eventName string, event *github.IssuesEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleIssuesEventEdited(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleIssuesEventEdited() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnIssuesEventDeleted(t *testing.T) {
	type args struct {
		callbacks []IssuesEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single IssuesEventHandleFunc",
			args: args{
				callbacks: []IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple IssuesEventHandleFunc",
			args: args{
				callbacks: []IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnIssuesEventDeleted(tt.args.callbacks...)
			if len(g.onIssuesEvent["deleted"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onIssuesEvent["deleted"]))
			}
		})
	}
}

func TestSetOnIssuesEventDeleted(t *testing.T) {
	type args struct {
		callbacks []IssuesEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single IssuesEventHandleFunc",
			args: args{
				[]IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple IssuesEventHandleFuncs",
			args: args{
				[]IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
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
			g.SetOnIssuesEventDeleted(func(deliveryID string, eventName string, event *github.IssuesEvent) error {
				return nil
			})
			g.SetOnIssuesEventDeleted(tt.args.callbacks...)
			if len(g.onIssuesEvent["deleted"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onIssuesEvent["deleted"]), tt.want)
			}
		})
	}
}

func TestHandleIssuesEventDeleted(t *testing.T) {
	action := "deleted"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.IssuesEvent
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
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnIssuesEventDeleted(func(deliveryID string, eventName string, event *github.IssuesEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleIssuesEventDeleted(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleIssuesEventDeleted() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnIssuesEventPinned(t *testing.T) {
	type args struct {
		callbacks []IssuesEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single IssuesEventHandleFunc",
			args: args{
				callbacks: []IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple IssuesEventHandleFunc",
			args: args{
				callbacks: []IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnIssuesEventPinned(tt.args.callbacks...)
			if len(g.onIssuesEvent["pinned"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onIssuesEvent["pinned"]))
			}
		})
	}
}

func TestSetOnIssuesEventPinned(t *testing.T) {
	type args struct {
		callbacks []IssuesEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single IssuesEventHandleFunc",
			args: args{
				[]IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple IssuesEventHandleFuncs",
			args: args{
				[]IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
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
			g.SetOnIssuesEventPinned(func(deliveryID string, eventName string, event *github.IssuesEvent) error {
				return nil
			})
			g.SetOnIssuesEventPinned(tt.args.callbacks...)
			if len(g.onIssuesEvent["pinned"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onIssuesEvent["pinned"]), tt.want)
			}
		})
	}
}

func TestHandleIssuesEventPinned(t *testing.T) {
	action := "pinned"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.IssuesEvent
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
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnIssuesEventPinned(func(deliveryID string, eventName string, event *github.IssuesEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleIssuesEventPinned(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleIssuesEventPinned() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnIssuesEventUnpinned(t *testing.T) {
	type args struct {
		callbacks []IssuesEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single IssuesEventHandleFunc",
			args: args{
				callbacks: []IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple IssuesEventHandleFunc",
			args: args{
				callbacks: []IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnIssuesEventUnpinned(tt.args.callbacks...)
			if len(g.onIssuesEvent["unpinned"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onIssuesEvent["unpinned"]))
			}
		})
	}
}

func TestSetOnIssuesEventUnpinned(t *testing.T) {
	type args struct {
		callbacks []IssuesEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single IssuesEventHandleFunc",
			args: args{
				[]IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple IssuesEventHandleFuncs",
			args: args{
				[]IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
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
			g.SetOnIssuesEventUnpinned(func(deliveryID string, eventName string, event *github.IssuesEvent) error {
				return nil
			})
			g.SetOnIssuesEventUnpinned(tt.args.callbacks...)
			if len(g.onIssuesEvent["unpinned"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onIssuesEvent["unpinned"]), tt.want)
			}
		})
	}
}

func TestHandleIssuesEventUnpinned(t *testing.T) {
	action := "unpinned"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.IssuesEvent
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
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnIssuesEventUnpinned(func(deliveryID string, eventName string, event *github.IssuesEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleIssuesEventUnpinned(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleIssuesEventUnpinned() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnIssuesEventClosed(t *testing.T) {
	type args struct {
		callbacks []IssuesEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single IssuesEventHandleFunc",
			args: args{
				callbacks: []IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple IssuesEventHandleFunc",
			args: args{
				callbacks: []IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnIssuesEventClosed(tt.args.callbacks...)
			if len(g.onIssuesEvent["closed"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onIssuesEvent["closed"]))
			}
		})
	}
}

func TestSetOnIssuesEventClosed(t *testing.T) {
	type args struct {
		callbacks []IssuesEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single IssuesEventHandleFunc",
			args: args{
				[]IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple IssuesEventHandleFuncs",
			args: args{
				[]IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
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
			g.SetOnIssuesEventClosed(func(deliveryID string, eventName string, event *github.IssuesEvent) error {
				return nil
			})
			g.SetOnIssuesEventClosed(tt.args.callbacks...)
			if len(g.onIssuesEvent["closed"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onIssuesEvent["closed"]), tt.want)
			}
		})
	}
}

func TestHandleIssuesEventClosed(t *testing.T) {
	action := "closed"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.IssuesEvent
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
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnIssuesEventClosed(func(deliveryID string, eventName string, event *github.IssuesEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleIssuesEventClosed(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleIssuesEventClosed() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnIssuesEventReopened(t *testing.T) {
	type args struct {
		callbacks []IssuesEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single IssuesEventHandleFunc",
			args: args{
				callbacks: []IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple IssuesEventHandleFunc",
			args: args{
				callbacks: []IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnIssuesEventReopened(tt.args.callbacks...)
			if len(g.onIssuesEvent["reopened"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onIssuesEvent["reopened"]))
			}
		})
	}
}

func TestSetOnIssuesEventReopened(t *testing.T) {
	type args struct {
		callbacks []IssuesEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single IssuesEventHandleFunc",
			args: args{
				[]IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple IssuesEventHandleFuncs",
			args: args{
				[]IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
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
			g.SetOnIssuesEventReopened(func(deliveryID string, eventName string, event *github.IssuesEvent) error {
				return nil
			})
			g.SetOnIssuesEventReopened(tt.args.callbacks...)
			if len(g.onIssuesEvent["reopened"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onIssuesEvent["reopened"]), tt.want)
			}
		})
	}
}

func TestHandleIssuesEventReopened(t *testing.T) {
	action := "reopened"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.IssuesEvent
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
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnIssuesEventReopened(func(deliveryID string, eventName string, event *github.IssuesEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleIssuesEventReopened(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleIssuesEventReopened() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnIssuesEventAssigned(t *testing.T) {
	type args struct {
		callbacks []IssuesEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single IssuesEventHandleFunc",
			args: args{
				callbacks: []IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple IssuesEventHandleFunc",
			args: args{
				callbacks: []IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnIssuesEventAssigned(tt.args.callbacks...)
			if len(g.onIssuesEvent["assigned"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onIssuesEvent["assigned"]))
			}
		})
	}
}

func TestSetOnIssuesEventAssigned(t *testing.T) {
	type args struct {
		callbacks []IssuesEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single IssuesEventHandleFunc",
			args: args{
				[]IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple IssuesEventHandleFuncs",
			args: args{
				[]IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
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
			g.SetOnIssuesEventAssigned(func(deliveryID string, eventName string, event *github.IssuesEvent) error {
				return nil
			})
			g.SetOnIssuesEventAssigned(tt.args.callbacks...)
			if len(g.onIssuesEvent["assigned"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onIssuesEvent["assigned"]), tt.want)
			}
		})
	}
}

func TestHandleIssuesEventAssigned(t *testing.T) {
	action := "assigned"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.IssuesEvent
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
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnIssuesEventAssigned(func(deliveryID string, eventName string, event *github.IssuesEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleIssuesEventAssigned(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleIssuesEventAssigned() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnIssuesEventUnassigned(t *testing.T) {
	type args struct {
		callbacks []IssuesEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single IssuesEventHandleFunc",
			args: args{
				callbacks: []IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple IssuesEventHandleFunc",
			args: args{
				callbacks: []IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnIssuesEventUnassigned(tt.args.callbacks...)
			if len(g.onIssuesEvent["unassigned"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onIssuesEvent["unassigned"]))
			}
		})
	}
}

func TestSetOnIssuesEventUnassigned(t *testing.T) {
	type args struct {
		callbacks []IssuesEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single IssuesEventHandleFunc",
			args: args{
				[]IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple IssuesEventHandleFuncs",
			args: args{
				[]IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
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
			g.SetOnIssuesEventUnassigned(func(deliveryID string, eventName string, event *github.IssuesEvent) error {
				return nil
			})
			g.SetOnIssuesEventUnassigned(tt.args.callbacks...)
			if len(g.onIssuesEvent["unassigned"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onIssuesEvent["unassigned"]), tt.want)
			}
		})
	}
}

func TestHandleIssuesEventUnassigned(t *testing.T) {
	action := "unassigned"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.IssuesEvent
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
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnIssuesEventUnassigned(func(deliveryID string, eventName string, event *github.IssuesEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleIssuesEventUnassigned(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleIssuesEventUnassigned() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnIssuesEventLabeled(t *testing.T) {
	type args struct {
		callbacks []IssuesEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single IssuesEventHandleFunc",
			args: args{
				callbacks: []IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple IssuesEventHandleFunc",
			args: args{
				callbacks: []IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnIssuesEventLabeled(tt.args.callbacks...)
			if len(g.onIssuesEvent["labeled"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onIssuesEvent["labeled"]))
			}
		})
	}
}

func TestSetOnIssuesEventLabeled(t *testing.T) {
	type args struct {
		callbacks []IssuesEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single IssuesEventHandleFunc",
			args: args{
				[]IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple IssuesEventHandleFuncs",
			args: args{
				[]IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
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
			g.SetOnIssuesEventLabeled(func(deliveryID string, eventName string, event *github.IssuesEvent) error {
				return nil
			})
			g.SetOnIssuesEventLabeled(tt.args.callbacks...)
			if len(g.onIssuesEvent["labeled"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onIssuesEvent["labeled"]), tt.want)
			}
		})
	}
}

func TestHandleIssuesEventLabeled(t *testing.T) {
	action := "labeled"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.IssuesEvent
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
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnIssuesEventLabeled(func(deliveryID string, eventName string, event *github.IssuesEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleIssuesEventLabeled(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleIssuesEventLabeled() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnIssuesEventUnlabeled(t *testing.T) {
	type args struct {
		callbacks []IssuesEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single IssuesEventHandleFunc",
			args: args{
				callbacks: []IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple IssuesEventHandleFunc",
			args: args{
				callbacks: []IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnIssuesEventUnlabeled(tt.args.callbacks...)
			if len(g.onIssuesEvent["unlabeled"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onIssuesEvent["unlabeled"]))
			}
		})
	}
}

func TestSetOnIssuesEventUnlabeled(t *testing.T) {
	type args struct {
		callbacks []IssuesEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single IssuesEventHandleFunc",
			args: args{
				[]IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple IssuesEventHandleFuncs",
			args: args{
				[]IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
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
			g.SetOnIssuesEventUnlabeled(func(deliveryID string, eventName string, event *github.IssuesEvent) error {
				return nil
			})
			g.SetOnIssuesEventUnlabeled(tt.args.callbacks...)
			if len(g.onIssuesEvent["unlabeled"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onIssuesEvent["unlabeled"]), tt.want)
			}
		})
	}
}

func TestHandleIssuesEventUnlabeled(t *testing.T) {
	action := "unlabeled"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.IssuesEvent
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
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnIssuesEventUnlabeled(func(deliveryID string, eventName string, event *github.IssuesEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleIssuesEventUnlabeled(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleIssuesEventUnlabeled() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnIssuesEventLocked(t *testing.T) {
	type args struct {
		callbacks []IssuesEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single IssuesEventHandleFunc",
			args: args{
				callbacks: []IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple IssuesEventHandleFunc",
			args: args{
				callbacks: []IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnIssuesEventLocked(tt.args.callbacks...)
			if len(g.onIssuesEvent["locked"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onIssuesEvent["locked"]))
			}
		})
	}
}

func TestSetOnIssuesEventLocked(t *testing.T) {
	type args struct {
		callbacks []IssuesEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single IssuesEventHandleFunc",
			args: args{
				[]IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple IssuesEventHandleFuncs",
			args: args{
				[]IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
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
			g.SetOnIssuesEventLocked(func(deliveryID string, eventName string, event *github.IssuesEvent) error {
				return nil
			})
			g.SetOnIssuesEventLocked(tt.args.callbacks...)
			if len(g.onIssuesEvent["locked"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onIssuesEvent["locked"]), tt.want)
			}
		})
	}
}

func TestHandleIssuesEventLocked(t *testing.T) {
	action := "locked"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.IssuesEvent
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
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnIssuesEventLocked(func(deliveryID string, eventName string, event *github.IssuesEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleIssuesEventLocked(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleIssuesEventLocked() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnIssuesEventUnlocked(t *testing.T) {
	type args struct {
		callbacks []IssuesEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single IssuesEventHandleFunc",
			args: args{
				callbacks: []IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple IssuesEventHandleFunc",
			args: args{
				callbacks: []IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnIssuesEventUnlocked(tt.args.callbacks...)
			if len(g.onIssuesEvent["unlocked"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onIssuesEvent["unlocked"]))
			}
		})
	}
}

func TestSetOnIssuesEventUnlocked(t *testing.T) {
	type args struct {
		callbacks []IssuesEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single IssuesEventHandleFunc",
			args: args{
				[]IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple IssuesEventHandleFuncs",
			args: args{
				[]IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
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
			g.SetOnIssuesEventUnlocked(func(deliveryID string, eventName string, event *github.IssuesEvent) error {
				return nil
			})
			g.SetOnIssuesEventUnlocked(tt.args.callbacks...)
			if len(g.onIssuesEvent["unlocked"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onIssuesEvent["unlocked"]), tt.want)
			}
		})
	}
}

func TestHandleIssuesEventUnlocked(t *testing.T) {
	action := "unlocked"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.IssuesEvent
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
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnIssuesEventUnlocked(func(deliveryID string, eventName string, event *github.IssuesEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleIssuesEventUnlocked(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleIssuesEventUnlocked() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnIssuesEventTransferred(t *testing.T) {
	type args struct {
		callbacks []IssuesEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single IssuesEventHandleFunc",
			args: args{
				callbacks: []IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple IssuesEventHandleFunc",
			args: args{
				callbacks: []IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnIssuesEventTransferred(tt.args.callbacks...)
			if len(g.onIssuesEvent["transferred"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onIssuesEvent["transferred"]))
			}
		})
	}
}

func TestSetOnIssuesEventTransferred(t *testing.T) {
	type args struct {
		callbacks []IssuesEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single IssuesEventHandleFunc",
			args: args{
				[]IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple IssuesEventHandleFuncs",
			args: args{
				[]IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
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
			g.SetOnIssuesEventTransferred(func(deliveryID string, eventName string, event *github.IssuesEvent) error {
				return nil
			})
			g.SetOnIssuesEventTransferred(tt.args.callbacks...)
			if len(g.onIssuesEvent["transferred"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onIssuesEvent["transferred"]), tt.want)
			}
		})
	}
}

func TestHandleIssuesEventTransferred(t *testing.T) {
	action := "transferred"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.IssuesEvent
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
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnIssuesEventTransferred(func(deliveryID string, eventName string, event *github.IssuesEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleIssuesEventTransferred(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleIssuesEventTransferred() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnIssuesEventMilestoned(t *testing.T) {
	type args struct {
		callbacks []IssuesEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single IssuesEventHandleFunc",
			args: args{
				callbacks: []IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple IssuesEventHandleFunc",
			args: args{
				callbacks: []IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnIssuesEventMilestoned(tt.args.callbacks...)
			if len(g.onIssuesEvent["milestoned"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onIssuesEvent["milestoned"]))
			}
		})
	}
}

func TestSetOnIssuesEventMilestoned(t *testing.T) {
	type args struct {
		callbacks []IssuesEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single IssuesEventHandleFunc",
			args: args{
				[]IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple IssuesEventHandleFuncs",
			args: args{
				[]IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
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
			g.SetOnIssuesEventMilestoned(func(deliveryID string, eventName string, event *github.IssuesEvent) error {
				return nil
			})
			g.SetOnIssuesEventMilestoned(tt.args.callbacks...)
			if len(g.onIssuesEvent["milestoned"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onIssuesEvent["milestoned"]), tt.want)
			}
		})
	}
}

func TestHandleIssuesEventMilestoned(t *testing.T) {
	action := "milestoned"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.IssuesEvent
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
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnIssuesEventMilestoned(func(deliveryID string, eventName string, event *github.IssuesEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleIssuesEventMilestoned(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleIssuesEventMilestoned() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnIssuesEventDeMilestoned(t *testing.T) {
	type args struct {
		callbacks []IssuesEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single IssuesEventHandleFunc",
			args: args{
				callbacks: []IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple IssuesEventHandleFunc",
			args: args{
				callbacks: []IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnIssuesEventDeMilestoned(tt.args.callbacks...)
			if len(g.onIssuesEvent["demilestoned"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onIssuesEvent["demilestoned"]))
			}
		})
	}
}

func TestSetOnIssuesEventDeMilestoned(t *testing.T) {
	type args struct {
		callbacks []IssuesEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single IssuesEventHandleFunc",
			args: args{
				[]IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple IssuesEventHandleFuncs",
			args: args{
				[]IssuesEventHandleFunc{
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.IssuesEvent) error {
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
			g.SetOnIssuesEventDeMilestoned(func(deliveryID string, eventName string, event *github.IssuesEvent) error {
				return nil
			})
			g.SetOnIssuesEventDeMilestoned(tt.args.callbacks...)
			if len(g.onIssuesEvent["demilestoned"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onIssuesEvent["demilestoned"]), tt.want)
			}
		})
	}
}

func TestHandleIssuesEventDeMilestoned(t *testing.T) {
	action := "demilestoned"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.IssuesEvent
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
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "issues",
				event:      &github.IssuesEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnIssuesEventDeMilestoned(func(deliveryID string, eventName string, event *github.IssuesEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleIssuesEventDeMilestoned(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleIssuesEventDeMilestoned() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
