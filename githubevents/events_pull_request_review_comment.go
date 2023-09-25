// Copyright 2022 The GithubEvents Authors. All rights reserved.
// Use of this source code is governed by the MIT License
// that can be found in the LICENSE file.

package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"fmt"
	"github.com/google/go-github/v55/github"
	"golang.org/x/sync/errgroup"
)

// Actions are used to identify registered callbacks.
const (
	// PullRequestReviewCommentEvent is the event name of github.PullRequestReviewCommentEvent's
	PullRequestReviewCommentEvent = "pull_request_review_comment"

	// PullRequestReviewCommentEventAnyAction is used to identify callbacks
	// listening to all events of type github.PullRequestReviewCommentEvent
	PullRequestReviewCommentEventAnyAction = "*"

	// PullRequestReviewCommentEventCreatedAction is used to identify callbacks
	// listening to events of type github.PullRequestReviewCommentEvent and action "created"
	PullRequestReviewCommentEventCreatedAction = "created"

	// PullRequestReviewCommentEventEditedAction is used to identify callbacks
	// listening to events of type github.PullRequestReviewCommentEvent and action "edited"
	PullRequestReviewCommentEventEditedAction = "edited"

	// PullRequestReviewCommentEventDeletedAction is used to identify callbacks
	// listening to events of type github.PullRequestReviewCommentEvent and action "deleted"
	PullRequestReviewCommentEventDeletedAction = "deleted"
)

// PullRequestReviewCommentEventHandleFunc represents a callback function triggered on github.PullRequestReviewCommentEvent's.
// 'deliveryID' (type: string) is the unique webhook delivery ID.
// 'eventName' (type: string) is the name of the event.
// 'event' (type: *github.PullRequestReviewCommentEvent) is the webhook payload.
type PullRequestReviewCommentEventHandleFunc func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error

// OnPullRequestReviewCommentEventCreated registers callbacks listening to events of type github.PullRequestReviewCommentEvent and action 'created'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnPullRequestReviewCommentEventCreated must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#pull_request_review_comment
func (g *EventHandler) OnPullRequestReviewCommentEventCreated(callbacks ...PullRequestReviewCommentEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestReviewCommentEvent == nil {
		g.onPullRequestReviewCommentEvent = make(map[string][]PullRequestReviewCommentEventHandleFunc)
	}
	g.onPullRequestReviewCommentEvent[PullRequestReviewCommentEventCreatedAction] = append(
		g.onPullRequestReviewCommentEvent[PullRequestReviewCommentEventCreatedAction],
		callbacks...,
	)
}

// SetOnPullRequestReviewCommentEventCreated registers callbacks listening to events of type github.PullRequestReviewCommentEvent and action 'created'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnPullRequestReviewCommentEventCreatedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#pull_request_review_comment
func (g *EventHandler) SetOnPullRequestReviewCommentEventCreated(callbacks ...PullRequestReviewCommentEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestReviewCommentEvent == nil {
		g.onPullRequestReviewCommentEvent = make(map[string][]PullRequestReviewCommentEventHandleFunc)
	}
	g.onPullRequestReviewCommentEvent[PullRequestReviewCommentEventCreatedAction] = callbacks
}

func (g *EventHandler) handlePullRequestReviewCommentEventCreated(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if PullRequestReviewCommentEventCreatedAction != *event.Action {
		return fmt.Errorf(
			"handlePullRequestReviewCommentEventCreated() called with wrong action, want %s, got %s",
			PullRequestReviewCommentEventCreatedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		PullRequestReviewCommentEventCreatedAction,
		PullRequestReviewCommentEventAnyAction,
	} {
		if _, ok := g.onPullRequestReviewCommentEvent[action]; ok {
			for _, h := range g.onPullRequestReviewCommentEvent[action] {
				handle := h
				eg.Go(func() error {
					err := handle(deliveryID, eventName, event)
					if err != nil {
						return err
					}
					return nil
				})
			}
		}
	}
	if err := eg.Wait(); err != nil {
		return err
	}
	return nil
}

// OnPullRequestReviewCommentEventEdited registers callbacks listening to events of type github.PullRequestReviewCommentEvent and action 'edited'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnPullRequestReviewCommentEventEdited must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#pull_request_review_comment
func (g *EventHandler) OnPullRequestReviewCommentEventEdited(callbacks ...PullRequestReviewCommentEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestReviewCommentEvent == nil {
		g.onPullRequestReviewCommentEvent = make(map[string][]PullRequestReviewCommentEventHandleFunc)
	}
	g.onPullRequestReviewCommentEvent[PullRequestReviewCommentEventEditedAction] = append(
		g.onPullRequestReviewCommentEvent[PullRequestReviewCommentEventEditedAction],
		callbacks...,
	)
}

// SetOnPullRequestReviewCommentEventEdited registers callbacks listening to events of type github.PullRequestReviewCommentEvent and action 'edited'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnPullRequestReviewCommentEventEditedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#pull_request_review_comment
func (g *EventHandler) SetOnPullRequestReviewCommentEventEdited(callbacks ...PullRequestReviewCommentEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestReviewCommentEvent == nil {
		g.onPullRequestReviewCommentEvent = make(map[string][]PullRequestReviewCommentEventHandleFunc)
	}
	g.onPullRequestReviewCommentEvent[PullRequestReviewCommentEventEditedAction] = callbacks
}

func (g *EventHandler) handlePullRequestReviewCommentEventEdited(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if PullRequestReviewCommentEventEditedAction != *event.Action {
		return fmt.Errorf(
			"handlePullRequestReviewCommentEventEdited() called with wrong action, want %s, got %s",
			PullRequestReviewCommentEventEditedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		PullRequestReviewCommentEventEditedAction,
		PullRequestReviewCommentEventAnyAction,
	} {
		if _, ok := g.onPullRequestReviewCommentEvent[action]; ok {
			for _, h := range g.onPullRequestReviewCommentEvent[action] {
				handle := h
				eg.Go(func() error {
					err := handle(deliveryID, eventName, event)
					if err != nil {
						return err
					}
					return nil
				})
			}
		}
	}
	if err := eg.Wait(); err != nil {
		return err
	}
	return nil
}

// OnPullRequestReviewCommentEventDeleted registers callbacks listening to events of type github.PullRequestReviewCommentEvent and action 'deleted'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnPullRequestReviewCommentEventDeleted must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#pull_request_review_comment
func (g *EventHandler) OnPullRequestReviewCommentEventDeleted(callbacks ...PullRequestReviewCommentEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestReviewCommentEvent == nil {
		g.onPullRequestReviewCommentEvent = make(map[string][]PullRequestReviewCommentEventHandleFunc)
	}
	g.onPullRequestReviewCommentEvent[PullRequestReviewCommentEventDeletedAction] = append(
		g.onPullRequestReviewCommentEvent[PullRequestReviewCommentEventDeletedAction],
		callbacks...,
	)
}

// SetOnPullRequestReviewCommentEventDeleted registers callbacks listening to events of type github.PullRequestReviewCommentEvent and action 'deleted'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnPullRequestReviewCommentEventDeletedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#pull_request_review_comment
func (g *EventHandler) SetOnPullRequestReviewCommentEventDeleted(callbacks ...PullRequestReviewCommentEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestReviewCommentEvent == nil {
		g.onPullRequestReviewCommentEvent = make(map[string][]PullRequestReviewCommentEventHandleFunc)
	}
	g.onPullRequestReviewCommentEvent[PullRequestReviewCommentEventDeletedAction] = callbacks
}

func (g *EventHandler) handlePullRequestReviewCommentEventDeleted(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if PullRequestReviewCommentEventDeletedAction != *event.Action {
		return fmt.Errorf(
			"handlePullRequestReviewCommentEventDeleted() called with wrong action, want %s, got %s",
			PullRequestReviewCommentEventDeletedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		PullRequestReviewCommentEventDeletedAction,
		PullRequestReviewCommentEventAnyAction,
	} {
		if _, ok := g.onPullRequestReviewCommentEvent[action]; ok {
			for _, h := range g.onPullRequestReviewCommentEvent[action] {
				handle := h
				eg.Go(func() error {
					err := handle(deliveryID, eventName, event)
					if err != nil {
						return err
					}
					return nil
				})
			}
		}
	}
	if err := eg.Wait(); err != nil {
		return err
	}
	return nil
}

// OnPullRequestReviewCommentEventAny registers callbacks listening to any events of type github.PullRequestReviewCommentEvent
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnPullRequestReviewCommentEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#pull_request_review_comment
func (g *EventHandler) OnPullRequestReviewCommentEventAny(callbacks ...PullRequestReviewCommentEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestReviewCommentEvent == nil {
		g.onPullRequestReviewCommentEvent = make(map[string][]PullRequestReviewCommentEventHandleFunc)
	}
	g.onPullRequestReviewCommentEvent[PullRequestReviewCommentEventAnyAction] = append(
		g.onPullRequestReviewCommentEvent[PullRequestReviewCommentEventAnyAction],
		callbacks...,
	)
}

// SetOnPullRequestReviewCommentEventAny registers callbacks listening to any events of type github.PullRequestReviewCommentEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnPullRequestReviewCommentEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#pull_request_review_comment
func (g *EventHandler) SetOnPullRequestReviewCommentEventAny(callbacks ...PullRequestReviewCommentEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestReviewCommentEvent == nil {
		g.onPullRequestReviewCommentEvent = make(map[string][]PullRequestReviewCommentEventHandleFunc)
	}
	g.onPullRequestReviewCommentEvent[PullRequestReviewCommentEventAnyAction] = callbacks
}

func (g *EventHandler) handlePullRequestReviewCommentEventAny(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
	if event == nil {
		return fmt.Errorf("event was empty or nil")
	}
	if _, ok := g.onPullRequestReviewCommentEvent[PullRequestReviewCommentEventAnyAction]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onPullRequestReviewCommentEvent[PullRequestReviewCommentEventAnyAction] {
		handle := h
		eg.Go(func() error {
			err := handle(deliveryID, eventName, event)
			if err != nil {
				return err
			}
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		return err
	}
	return nil
}

// PullRequestReviewCommentEvent handles github.PullRequestReviewCommentEvent.
//
// Callbacks are executed in the following order:
//
// 1) All callbacks registered with OnBeforeAny are executed in parallel.
// 2) All callbacks registered with OnPullRequestReviewCommentEvent... are executed in parallel in case the Event has actions.
// 3) All callbacks registered with OnAfterAny are executed in parallel.
//
// on any error all callbacks registered with OnError are executed in parallel.
func (g *EventHandler) PullRequestReviewCommentEvent(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {

	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	action := *event.Action

	err := g.handleBeforeAny(deliveryID, eventName, event)
	if err != nil {
		return g.handleError(deliveryID, eventName, event, err)
	}

	switch action {

	case PullRequestReviewCommentEventCreatedAction:
		err := g.handlePullRequestReviewCommentEventCreated(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case PullRequestReviewCommentEventEditedAction:
		err := g.handlePullRequestReviewCommentEventEdited(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case PullRequestReviewCommentEventDeletedAction:
		err := g.handlePullRequestReviewCommentEventDeleted(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	default:
		err := g.handlePullRequestReviewCommentEventAny(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}
	}

	err = g.handleAfterAny(deliveryID, eventName, event)
	if err != nil {
		return g.handleError(deliveryID, eventName, event, err)
	}
	return nil
}
