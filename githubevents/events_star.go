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
	// StarEvent is the event name of github.StarEvent's
	StarEvent = "star"

	// StarEventAnyAction is used to identify callbacks
	// listening to all events of type github.StarEvent
	StarEventAnyAction = "*"

	// StarEventCreatedAction is used to identify callbacks
	// listening to events of type github.StarEvent and action "created"
	StarEventCreatedAction = "created"

	// StarEventDeletedAction is used to identify callbacks
	// listening to events of type github.StarEvent and action "deleted"
	StarEventDeletedAction = "deleted"
)

// StarEventHandleFunc represents a callback function triggered on github.StarEvent's.
// 'deliveryID' (type: string) is the unique webhook delivery ID.
// 'eventName' (type: string) is the name of the event.
// 'event' (type: *github.StarEvent) is the webhook payload.
type StarEventHandleFunc func(deliveryID string, eventName string, event *github.StarEvent) error

// OnStarEventCreated registers callbacks listening to events of type github.StarEvent and action 'created'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnStarEventCreated must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#star
func (g *EventHandler) OnStarEventCreated(callbacks ...StarEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onStarEvent == nil {
		g.onStarEvent = make(map[string][]StarEventHandleFunc)
	}
	g.onStarEvent[StarEventCreatedAction] = append(
		g.onStarEvent[StarEventCreatedAction],
		callbacks...,
	)
}

// SetOnStarEventCreated registers callbacks listening to events of type github.StarEvent and action 'created'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnStarEventCreatedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#star
func (g *EventHandler) SetOnStarEventCreated(callbacks ...StarEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onStarEvent == nil {
		g.onStarEvent = make(map[string][]StarEventHandleFunc)
	}
	g.onStarEvent[StarEventCreatedAction] = callbacks
}

func (g *EventHandler) handleStarEventCreated(deliveryID string, eventName string, event *github.StarEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if StarEventCreatedAction != *event.Action {
		return fmt.Errorf(
			"handleStarEventCreated() called with wrong action, want %s, got %s",
			StarEventCreatedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		StarEventCreatedAction,
		StarEventAnyAction,
	} {
		if _, ok := g.onStarEvent[action]; ok {
			for _, h := range g.onStarEvent[action] {
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

// OnStarEventDeleted registers callbacks listening to events of type github.StarEvent and action 'deleted'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnStarEventDeleted must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#star
func (g *EventHandler) OnStarEventDeleted(callbacks ...StarEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onStarEvent == nil {
		g.onStarEvent = make(map[string][]StarEventHandleFunc)
	}
	g.onStarEvent[StarEventDeletedAction] = append(
		g.onStarEvent[StarEventDeletedAction],
		callbacks...,
	)
}

// SetOnStarEventDeleted registers callbacks listening to events of type github.StarEvent and action 'deleted'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnStarEventDeletedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#star
func (g *EventHandler) SetOnStarEventDeleted(callbacks ...StarEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onStarEvent == nil {
		g.onStarEvent = make(map[string][]StarEventHandleFunc)
	}
	g.onStarEvent[StarEventDeletedAction] = callbacks
}

func (g *EventHandler) handleStarEventDeleted(deliveryID string, eventName string, event *github.StarEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if StarEventDeletedAction != *event.Action {
		return fmt.Errorf(
			"handleStarEventDeleted() called with wrong action, want %s, got %s",
			StarEventDeletedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		StarEventDeletedAction,
		StarEventAnyAction,
	} {
		if _, ok := g.onStarEvent[action]; ok {
			for _, h := range g.onStarEvent[action] {
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

// OnStarEventAny registers callbacks listening to any events of type github.StarEvent
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnStarEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#star
func (g *EventHandler) OnStarEventAny(callbacks ...StarEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onStarEvent == nil {
		g.onStarEvent = make(map[string][]StarEventHandleFunc)
	}
	g.onStarEvent[StarEventAnyAction] = append(
		g.onStarEvent[StarEventAnyAction],
		callbacks...,
	)
}

// SetOnStarEventAny registers callbacks listening to any events of type github.StarEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnStarEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#star
func (g *EventHandler) SetOnStarEventAny(callbacks ...StarEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onStarEvent == nil {
		g.onStarEvent = make(map[string][]StarEventHandleFunc)
	}
	g.onStarEvent[StarEventAnyAction] = callbacks
}

func (g *EventHandler) handleStarEventAny(deliveryID string, eventName string, event *github.StarEvent) error {
	if event == nil {
		return fmt.Errorf("event was empty or nil")
	}
	if _, ok := g.onStarEvent[StarEventAnyAction]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onStarEvent[StarEventAnyAction] {
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

// StarEvent handles github.StarEvent.
//
// Callbacks are executed in the following order:
//
// 1) All callbacks registered with OnBeforeAny are executed in parallel.
// 2) All callbacks registered with OnStarEvent... are executed in parallel in case the Event has actions.
// 3) All callbacks registered with OnAfterAny are executed in parallel.
//
// on any error all callbacks registered with OnError are executed in parallel.
func (g *EventHandler) StarEvent(deliveryID string, eventName string, event *github.StarEvent) error {

	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	action := *event.Action

	err := g.handleBeforeAny(deliveryID, eventName, event)
	if err != nil {
		return g.handleError(deliveryID, eventName, event, err)
	}

	switch action {

	case StarEventCreatedAction:
		err := g.handleStarEventCreated(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case StarEventDeletedAction:
		err := g.handleStarEventDeleted(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	default:
		err := g.handleStarEventAny(deliveryID, eventName, event)
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
