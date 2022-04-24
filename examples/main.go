package main

import (
	"fmt"
	"github.com/cbrgm/githubevents/githubevents"
	"github.com/google/go-github/v43/github"
	"github.com/rs/zerolog/log"
)

func main() {
	handle := githubevents.New("")

	newNotifier(handle)

	action := "deleted"
	handle.BranchProtectionRuleEvent("42", "foo", &github.BranchProtectionRuleEvent{Action: &action})
}

func newNotifier(handle *githubevents.EventHandler) {

	handle.OnBeforeAny(
		func(deliveryID string, eventName string, event interface{}) error {
			log.Info().Msg("onBeforeAny 001")
			return nil
		},
		func(deliveryID string, eventName string, event interface{}) error {
			log.Info().Msg("onBeforeAny 002")
			return nil
		},
		func(deliveryID string, eventName string, event interface{}) error {
			log.Info().Msg("onBeforeAny 003")
			return nil
		},
	)

	handle.OnBranchProtectionRuleEventCreated(
		func(deliveryID string, eventName string, event *github.BranchProtectionRuleEvent) error {
			log.Info().Msg("OnBranchProtectionRuleEventCreated 001")
			fmt.Println("aaaaa")
			return nil
		},
		func(deliveryID string, eventName string, event *github.BranchProtectionRuleEvent) error {
			log.Info().Msg("OnBranchProtectionRuleEventCreated 002222")
			return nil
		},
	)

	handle.OnError(
		func(deliveryID string, eventName string, event interface{}, err error) error {
			fmt.Printf("received error %s", err)
			return err
		},
	)

	handle.OnIssueCommentCreated(
		func(deliveryID string, eventName string, event *github.IssueCommentEvent) error {
			fmt.Printf("%s made a comment!", *event.Sender.Login)
			return nil
		},
	)

	handle.OnBranchProtectionRuleEventDeleted(
		func(deliveryID string, eventName string, event *github.BranchProtectionRuleEvent) error {
			log.Info().Msg("OnBranchProtectionRuleEventDeleted 001")
			return nil
		},
	)

	handle.OnAfterAny(func(deliveryID string, eventName string, event interface{}) error {
		log.Info().Msg("onAfterAny 001")
		return nil
	})

	handle.OnError(func(deliveryID string, eventName string, event interface{}, err error) error {
		log.Info().Msg("OnError 001")
		return nil
	})

}
