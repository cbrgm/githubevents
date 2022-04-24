package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"errors"
	"github.com/google/go-github/v43/github"
	"testing"
)

func TestOnPullRequestReviewCommentEventAny(t *testing.T) {
	type args struct {
		callbacks []PullRequestReviewCommentEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single PullRequestReviewCommentEventHandleFunc",
			args: args{
				[]PullRequestReviewCommentEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple PullRequestReviewCommentEventHandleFuncs",
			args: args{
				[]PullRequestReviewCommentEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnPullRequestReviewCommentEventAny(tt.args.callbacks...)
			if len(g.onPullRequestReviewCommentEvent["*"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onPullRequestReviewCommentEvent["*"]))
			}
		})
	}
}

func TestSetOnPullRequestReviewCommentEventAny(t *testing.T) {
	type args struct {
		callbacks []PullRequestReviewCommentEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single PullRequestReviewCommentEventHandleFunc",
			args: args{
				[]PullRequestReviewCommentEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple PullRequestReviewCommentEventHandleFuncs",
			args: args{
				[]PullRequestReviewCommentEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
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
			g.SetOnPullRequestReviewCommentEventAny(func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
				return nil
			})
			g.SetOnPullRequestReviewCommentEventAny(tt.args.callbacks...)
			if len(g.onPullRequestReviewCommentEvent["*"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onPullRequestReviewCommentEvent["*"]), tt.want)
			}
		})
	}
}

func TestHandlePullRequestReviewCommentEventAny(t *testing.T) {

	action := "*"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.PullRequestReviewCommentEvent
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
				eventName:  "pull_request_review_comment",

				event: &github.PullRequestReviewCommentEvent{Action: &action},

				fail: false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request_review_comment",

				event: &github.PullRequestReviewCommentEvent{Action: &action},

				fail: true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request_review_comment",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnPullRequestReviewCommentEventAny(func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handlePullRequestReviewCommentEventAny(tt.args.deliveryID, tt.args.deliveryID, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("TestHandlePullRequestReviewCommentEventAny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnPullRequestReviewCommentEventCreated(t *testing.T) {
	type args struct {
		callbacks []PullRequestReviewCommentEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single PullRequestReviewCommentEventHandleFunc",
			args: args{
				callbacks: []PullRequestReviewCommentEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple PullRequestReviewCommentEventHandleFunc",
			args: args{
				callbacks: []PullRequestReviewCommentEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnPullRequestReviewCommentEventCreated(tt.args.callbacks...)
			if len(g.onPullRequestReviewCommentEvent["created"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onPullRequestReviewCommentEvent["created"]))
			}
		})
	}
}

func TestSetOnPullRequestReviewCommentEventCreated(t *testing.T) {
	type args struct {
		callbacks []PullRequestReviewCommentEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single PullRequestReviewCommentEventHandleFunc",
			args: args{
				[]PullRequestReviewCommentEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple PullRequestReviewCommentEventHandleFuncs",
			args: args{
				[]PullRequestReviewCommentEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
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
			g.SetOnPullRequestReviewCommentEventCreated(func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
				return nil
			})
			g.SetOnPullRequestReviewCommentEventCreated(tt.args.callbacks...)
			if len(g.onPullRequestReviewCommentEvent["created"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onPullRequestReviewCommentEvent["created"]), tt.want)
			}
		})
	}
}

func TestHandlePullRequestReviewCommentEventCreated(t *testing.T) {
	action := "created"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.PullRequestReviewCommentEvent
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
				eventName:  "pull_request_review_comment",
				event:      &github.PullRequestReviewCommentEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request_review_comment",
				event:      &github.PullRequestReviewCommentEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request_review_comment",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request_review_comment",
				event:      &github.PullRequestReviewCommentEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request_review_comment",
				event:      &github.PullRequestReviewCommentEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request_review_comment",
				event:      &github.PullRequestReviewCommentEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnPullRequestReviewCommentEventCreated(func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handlePullRequestReviewCommentEventCreated(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handlePullRequestReviewCommentEventCreated() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnPullRequestReviewCommentEventEdited(t *testing.T) {
	type args struct {
		callbacks []PullRequestReviewCommentEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single PullRequestReviewCommentEventHandleFunc",
			args: args{
				callbacks: []PullRequestReviewCommentEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple PullRequestReviewCommentEventHandleFunc",
			args: args{
				callbacks: []PullRequestReviewCommentEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnPullRequestReviewCommentEventEdited(tt.args.callbacks...)
			if len(g.onPullRequestReviewCommentEvent["edited"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onPullRequestReviewCommentEvent["edited"]))
			}
		})
	}
}

func TestSetOnPullRequestReviewCommentEventEdited(t *testing.T) {
	type args struct {
		callbacks []PullRequestReviewCommentEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single PullRequestReviewCommentEventHandleFunc",
			args: args{
				[]PullRequestReviewCommentEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple PullRequestReviewCommentEventHandleFuncs",
			args: args{
				[]PullRequestReviewCommentEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
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
			g.SetOnPullRequestReviewCommentEventEdited(func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
				return nil
			})
			g.SetOnPullRequestReviewCommentEventEdited(tt.args.callbacks...)
			if len(g.onPullRequestReviewCommentEvent["edited"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onPullRequestReviewCommentEvent["edited"]), tt.want)
			}
		})
	}
}

func TestHandlePullRequestReviewCommentEventEdited(t *testing.T) {
	action := "edited"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.PullRequestReviewCommentEvent
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
				eventName:  "pull_request_review_comment",
				event:      &github.PullRequestReviewCommentEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request_review_comment",
				event:      &github.PullRequestReviewCommentEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request_review_comment",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request_review_comment",
				event:      &github.PullRequestReviewCommentEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request_review_comment",
				event:      &github.PullRequestReviewCommentEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request_review_comment",
				event:      &github.PullRequestReviewCommentEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnPullRequestReviewCommentEventEdited(func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handlePullRequestReviewCommentEventEdited(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handlePullRequestReviewCommentEventEdited() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnPullRequestReviewCommentEventDeleted(t *testing.T) {
	type args struct {
		callbacks []PullRequestReviewCommentEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single PullRequestReviewCommentEventHandleFunc",
			args: args{
				callbacks: []PullRequestReviewCommentEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple PullRequestReviewCommentEventHandleFunc",
			args: args{
				callbacks: []PullRequestReviewCommentEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnPullRequestReviewCommentEventDeleted(tt.args.callbacks...)
			if len(g.onPullRequestReviewCommentEvent["deleted"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onPullRequestReviewCommentEvent["deleted"]))
			}
		})
	}
}

func TestSetOnPullRequestReviewCommentEventDeleted(t *testing.T) {
	type args struct {
		callbacks []PullRequestReviewCommentEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single PullRequestReviewCommentEventHandleFunc",
			args: args{
				[]PullRequestReviewCommentEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple PullRequestReviewCommentEventHandleFuncs",
			args: args{
				[]PullRequestReviewCommentEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
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
			g.SetOnPullRequestReviewCommentEventDeleted(func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
				return nil
			})
			g.SetOnPullRequestReviewCommentEventDeleted(tt.args.callbacks...)
			if len(g.onPullRequestReviewCommentEvent["deleted"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onPullRequestReviewCommentEvent["deleted"]), tt.want)
			}
		})
	}
}

func TestHandlePullRequestReviewCommentEventDeleted(t *testing.T) {
	action := "deleted"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.PullRequestReviewCommentEvent
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
				eventName:  "pull_request_review_comment",
				event:      &github.PullRequestReviewCommentEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request_review_comment",
				event:      &github.PullRequestReviewCommentEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request_review_comment",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request_review_comment",
				event:      &github.PullRequestReviewCommentEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request_review_comment",
				event:      &github.PullRequestReviewCommentEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request_review_comment",
				event:      &github.PullRequestReviewCommentEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnPullRequestReviewCommentEventDeleted(func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handlePullRequestReviewCommentEventDeleted(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handlePullRequestReviewCommentEventDeleted() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
