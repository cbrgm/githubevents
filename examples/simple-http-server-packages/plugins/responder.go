package plugins

import (
	"context"
	"fmt"

	"github.com/cbrgm/githubevents/v2/githubevents"
	"github.com/google/go-github/v70/github"
)

func NewResponder(msg string) githubevents.IssueCommentEventHandleFunc {
	// do some configuration here
	// ...
	return func(ctx context.Context, deliveryID string, eventName string, event *github.IssueCommentEvent) error {
		fmt.Printf("commenting %s", msg)
		return nil
	}
}
