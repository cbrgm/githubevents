package plugins

import (
	"context"
	"fmt"

	"github.com/cbrgm/githubevents/githubevents"
	"github.com/google/go-github/v69/github"
)

func NewResponder(msg string) githubevents.IssueCommentEventHandleFunc {
	// do some configuration here
	// ...
	return func(ctx context.Context, deliveryID string, eventName string, event *github.IssueCommentEvent) error {
		fmt.Printf("commenting %s", msg)
		return nil
	}
}
