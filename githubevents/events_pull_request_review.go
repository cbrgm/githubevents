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
	// PullRequestReviewEvent is the event name of github.PullRequestReviewEvent's
	PullRequestReviewEvent = "pull_request_review"

	// PullRequestReviewEventAnyAction is used to identify callbacks
	// listening to all events of type github.PullRequestReviewEvent
	PullRequestReviewEventAnyAction = "*"

	// PullRequestReviewEventSubmittedAction is used to identify callbacks
	// listening to events of type github.PullRequestReviewEvent and action "submitted"
	PullRequestReviewEventSubmittedAction = "submitted"

	// PullRequestReviewEventEditedAction is used to identify callbacks
	// listening to events of type github.PullRequestReviewEvent and action "edited"
	PullRequestReviewEventEditedAction = "edited"

	// PullRequestReviewEventDismissedAction is used to identify callbacks
	// listening to events of type github.PullRequestReviewEvent and action "dismissed"
	PullRequestReviewEventDismissedAction = "dismissed"
)

// PullRequestReviewEventHandleFunc represents a callback function triggered on github.PullRequestReviewEvent's.
// 'deliveryID' (type: string) is the unique webhook delivery ID.
// 'eventName' (type: string) is the name of the event.
// 'event' (type: *github.PullRequestReviewEvent) is the webhook payload.
type PullRequestReviewEventHandleFunc func(deliveryID string, eventName string, event *github.PullRequestReviewEvent) error

// OnPullRequestReviewEventSubmitted registers callbacks listening to events of type github.PullRequestReviewEvent and action 'submitted'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnPullRequestReviewEventSubmitted must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#pull_request_review
func (g *EventHandler) OnPullRequestReviewEventSubmitted(callbacks ...PullRequestReviewEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestReviewEvent == nil {
		g.onPullRequestReviewEvent = make(map[string][]PullRequestReviewEventHandleFunc)
	}
	g.onPullRequestReviewEvent[PullRequestReviewEventSubmittedAction] = append(
		g.onPullRequestReviewEvent[PullRequestReviewEventSubmittedAction],
		callbacks...,
	)
}

// SetOnPullRequestReviewEventSubmitted registers callbacks listening to events of type github.PullRequestReviewEvent and action 'submitted'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnPullRequestReviewEventSubmittedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#pull_request_review
func (g *EventHandler) SetOnPullRequestReviewEventSubmitted(callbacks ...PullRequestReviewEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestReviewEvent == nil {
		g.onPullRequestReviewEvent = make(map[string][]PullRequestReviewEventHandleFunc)
	}
	g.onPullRequestReviewEvent[PullRequestReviewEventSubmittedAction] = callbacks
}

func (g *EventHandler) handlePullRequestReviewEventSubmitted(deliveryID string, eventName string, event *github.PullRequestReviewEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if PullRequestReviewEventSubmittedAction != *event.Action {
		return fmt.Errorf(
			"handlePullRequestReviewEventSubmitted() called with wrong action, want %s, got %s",
			PullRequestReviewEventSubmittedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		PullRequestReviewEventSubmittedAction,
		PullRequestReviewEventAnyAction,
	} {
		if _, ok := g.onPullRequestReviewEvent[action]; ok {
			for _, h := range g.onPullRequestReviewEvent[action] {
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

// OnPullRequestReviewEventEdited registers callbacks listening to events of type github.PullRequestReviewEvent and action 'edited'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnPullRequestReviewEventEdited must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#pull_request_review
func (g *EventHandler) OnPullRequestReviewEventEdited(callbacks ...PullRequestReviewEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestReviewEvent == nil {
		g.onPullRequestReviewEvent = make(map[string][]PullRequestReviewEventHandleFunc)
	}
	g.onPullRequestReviewEvent[PullRequestReviewEventEditedAction] = append(
		g.onPullRequestReviewEvent[PullRequestReviewEventEditedAction],
		callbacks...,
	)
}

// SetOnPullRequestReviewEventEdited registers callbacks listening to events of type github.PullRequestReviewEvent and action 'edited'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnPullRequestReviewEventEditedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#pull_request_review
func (g *EventHandler) SetOnPullRequestReviewEventEdited(callbacks ...PullRequestReviewEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestReviewEvent == nil {
		g.onPullRequestReviewEvent = make(map[string][]PullRequestReviewEventHandleFunc)
	}
	g.onPullRequestReviewEvent[PullRequestReviewEventEditedAction] = callbacks
}

func (g *EventHandler) handlePullRequestReviewEventEdited(deliveryID string, eventName string, event *github.PullRequestReviewEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if PullRequestReviewEventEditedAction != *event.Action {
		return fmt.Errorf(
			"handlePullRequestReviewEventEdited() called with wrong action, want %s, got %s",
			PullRequestReviewEventEditedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		PullRequestReviewEventEditedAction,
		PullRequestReviewEventAnyAction,
	} {
		if _, ok := g.onPullRequestReviewEvent[action]; ok {
			for _, h := range g.onPullRequestReviewEvent[action] {
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

// OnPullRequestReviewEventDismissed registers callbacks listening to events of type github.PullRequestReviewEvent and action 'dismissed'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnPullRequestReviewEventDismissed must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#pull_request_review
func (g *EventHandler) OnPullRequestReviewEventDismissed(callbacks ...PullRequestReviewEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestReviewEvent == nil {
		g.onPullRequestReviewEvent = make(map[string][]PullRequestReviewEventHandleFunc)
	}
	g.onPullRequestReviewEvent[PullRequestReviewEventDismissedAction] = append(
		g.onPullRequestReviewEvent[PullRequestReviewEventDismissedAction],
		callbacks...,
	)
}

// SetOnPullRequestReviewEventDismissed registers callbacks listening to events of type github.PullRequestReviewEvent and action 'dismissed'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnPullRequestReviewEventDismissedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#pull_request_review
func (g *EventHandler) SetOnPullRequestReviewEventDismissed(callbacks ...PullRequestReviewEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestReviewEvent == nil {
		g.onPullRequestReviewEvent = make(map[string][]PullRequestReviewEventHandleFunc)
	}
	g.onPullRequestReviewEvent[PullRequestReviewEventDismissedAction] = callbacks
}

func (g *EventHandler) handlePullRequestReviewEventDismissed(deliveryID string, eventName string, event *github.PullRequestReviewEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if PullRequestReviewEventDismissedAction != *event.Action {
		return fmt.Errorf(
			"handlePullRequestReviewEventDismissed() called with wrong action, want %s, got %s",
			PullRequestReviewEventDismissedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		PullRequestReviewEventDismissedAction,
		PullRequestReviewEventAnyAction,
	} {
		if _, ok := g.onPullRequestReviewEvent[action]; ok {
			for _, h := range g.onPullRequestReviewEvent[action] {
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

// OnPullRequestReviewEventAny registers callbacks listening to any events of type github.PullRequestReviewEvent
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnPullRequestReviewEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#pull_request_review
func (g *EventHandler) OnPullRequestReviewEventAny(callbacks ...PullRequestReviewEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestReviewEvent == nil {
		g.onPullRequestReviewEvent = make(map[string][]PullRequestReviewEventHandleFunc)
	}
	g.onPullRequestReviewEvent[PullRequestReviewEventAnyAction] = append(
		g.onPullRequestReviewEvent[PullRequestReviewEventAnyAction],
		callbacks...,
	)
}

// SetOnPullRequestReviewEventAny registers callbacks listening to any events of type github.PullRequestReviewEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnPullRequestReviewEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#pull_request_review
func (g *EventHandler) SetOnPullRequestReviewEventAny(callbacks ...PullRequestReviewEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestReviewEvent == nil {
		g.onPullRequestReviewEvent = make(map[string][]PullRequestReviewEventHandleFunc)
	}
	g.onPullRequestReviewEvent[PullRequestReviewEventAnyAction] = callbacks
}

func (g *EventHandler) handlePullRequestReviewEventAny(deliveryID string, eventName string, event *github.PullRequestReviewEvent) error {
	if event == nil {
		return fmt.Errorf("event was empty or nil")
	}
	if _, ok := g.onPullRequestReviewEvent[PullRequestReviewEventAnyAction]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onPullRequestReviewEvent[PullRequestReviewEventAnyAction] {
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

// PullRequestReviewEvent handles github.PullRequestReviewEvent.
//
// Callbacks are executed in the following order:
//
// 1) All callbacks registered with OnBeforeAny are executed in parallel.
// 2) All callbacks registered with OnPullRequestReviewEvent... are executed in parallel in case the Event has actions.
// 3) All callbacks registered with OnAfterAny are executed in parallel.
//
// on any error all callbacks registered with OnError are executed in parallel.
func (g *EventHandler) PullRequestReviewEvent(deliveryID string, eventName string, event *github.PullRequestReviewEvent) error {

	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	action := *event.Action

	err := g.handleBeforeAny(deliveryID, eventName, event)
	if err != nil {
		return g.handleError(deliveryID, eventName, event, err)
	}

	switch action {

	case PullRequestReviewEventSubmittedAction:
		err := g.handlePullRequestReviewEventSubmitted(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case PullRequestReviewEventEditedAction:
		err := g.handlePullRequestReviewEventEdited(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case PullRequestReviewEventDismissedAction:
		err := g.handlePullRequestReviewEventDismissed(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	default:
		err := g.handlePullRequestReviewEventAny(deliveryID, eventName, event)
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
