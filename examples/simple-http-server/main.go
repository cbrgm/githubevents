package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cbrgm/githubevents/v2/githubevents"
	"github.com/google/go-github/v78/github"
)

func main() {
	handle := githubevents.New("")

	handle.OnIssueCommentCreated(func(ctx context.Context, deliveryID, eventName string, event *github.IssueCommentEvent) error {
		fmt.Printf("%s has commented on issue %d", *event.Sender.Login, *event.Issue.ID)
		return nil
	})

	http.HandleFunc("/hook", func(w http.ResponseWriter, r *http.Request) {
		err := handle.HandleEventRequest(r)
		if err != nil {
			fmt.Println("error")
		}
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
