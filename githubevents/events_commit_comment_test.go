package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"errors"
	"github.com/google/go-github/v43/github"
	"testing"
)

func TestOnCommitCommentEventAny(t *testing.T) {
	type args struct {
		callbacks []CommitCommentEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single CommitCommentEventHandleFunc",
			args: args{
				[]CommitCommentEventHandleFunc{
					func(deliveryID string, eventName string, event *github.CommitCommentEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple CommitCommentEventHandleFuncs",
			args: args{
				[]CommitCommentEventHandleFunc{
					func(deliveryID string, eventName string, event *github.CommitCommentEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.CommitCommentEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnCommitCommentEventAny(tt.args.callbacks...)
			if len(g.onCommitCommentEvent["*"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onCommitCommentEvent["*"]))
			}
		})
	}
}

func TestSetOnCommitCommentEventAny(t *testing.T) {
	type args struct {
		callbacks []CommitCommentEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single CommitCommentEventHandleFunc",
			args: args{
				[]CommitCommentEventHandleFunc{
					func(deliveryID string, eventName string, event *github.CommitCommentEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple CommitCommentEventHandleFuncs",
			args: args{
				[]CommitCommentEventHandleFunc{
					func(deliveryID string, eventName string, event *github.CommitCommentEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.CommitCommentEvent) error {
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
			g.SetOnCommitCommentEventAny(func(deliveryID string, eventName string, event *github.CommitCommentEvent) error {
				return nil
			})
			g.SetOnCommitCommentEventAny(tt.args.callbacks...)
			if len(g.onCommitCommentEvent["*"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onCommitCommentEvent["*"]), tt.want)
			}
		})
	}
}

func TestHandleCommitCommentEventAny(t *testing.T) {

	action := "*"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.CommitCommentEvent
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
				eventName:  "commit_comment",

				event: &github.CommitCommentEvent{Action: &action},

				fail: false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "commit_comment",

				event: &github.CommitCommentEvent{Action: &action},

				fail: true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "commit_comment",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnCommitCommentEventAny(func(deliveryID string, eventName string, event *github.CommitCommentEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleCommitCommentEventAny(tt.args.deliveryID, tt.args.deliveryID, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("TestHandleCommitCommentEventAny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnCommitCommentEventCreated(t *testing.T) {
	type args struct {
		callbacks []CommitCommentEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single CommitCommentEventHandleFunc",
			args: args{
				callbacks: []CommitCommentEventHandleFunc{
					func(deliveryID string, eventName string, event *github.CommitCommentEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple CommitCommentEventHandleFunc",
			args: args{
				callbacks: []CommitCommentEventHandleFunc{
					func(deliveryID string, eventName string, event *github.CommitCommentEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.CommitCommentEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnCommitCommentEventCreated(tt.args.callbacks...)
			if len(g.onCommitCommentEvent["created"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onCommitCommentEvent["created"]))
			}
		})
	}
}

func TestSetOnCommitCommentEventCreated(t *testing.T) {
	type args struct {
		callbacks []CommitCommentEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single CommitCommentEventHandleFunc",
			args: args{
				[]CommitCommentEventHandleFunc{
					func(deliveryID string, eventName string, event *github.CommitCommentEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple CommitCommentEventHandleFuncs",
			args: args{
				[]CommitCommentEventHandleFunc{
					func(deliveryID string, eventName string, event *github.CommitCommentEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.CommitCommentEvent) error {
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
			g.SetOnCommitCommentEventCreated(func(deliveryID string, eventName string, event *github.CommitCommentEvent) error {
				return nil
			})
			g.SetOnCommitCommentEventCreated(tt.args.callbacks...)
			if len(g.onCommitCommentEvent["created"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onCommitCommentEvent["created"]), tt.want)
			}
		})
	}
}

func TestHandleCommitCommentEventCreated(t *testing.T) {
	action := "created"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.CommitCommentEvent
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
				eventName:  "commit_comment",
				event:      &github.CommitCommentEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "commit_comment",
				event:      &github.CommitCommentEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "commit_comment",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "commit_comment",
				event:      &github.CommitCommentEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "commit_comment",
				event:      &github.CommitCommentEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "commit_comment",
				event:      &github.CommitCommentEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnCommitCommentEventCreated(func(deliveryID string, eventName string, event *github.CommitCommentEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleCommitCommentEventCreated(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleCommitCommentEventCreated() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
