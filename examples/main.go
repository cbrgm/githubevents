package main

import (
	"fmt"
	"github.com/cbrgm/githubevents/githubevents"
	"github.com/google/go-github/v43/github"
	"github.com/rs/zerolog/log"
	"net/http"
)

func main() {
	handle := githubevents.New("")

	newNotifier(handle)

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

func newNotifier(handle *githubevents.EventHandler) {

	handle.OnBeforeAny(
		func(deliveryID string, eventName string, event interface{}) error {
			log.Info().Msg("onBeforeAny 001")
			return nil
		},
	)

	handle.OnIssueCommentCreated(func(deliveryID string, eventName string, event *github.IssueCommentEvent) error {
		fmt.Printf("%s has commented on issue %s", *event.Sender.Login, *event.Issue.ID)
		return nil
	})
}
