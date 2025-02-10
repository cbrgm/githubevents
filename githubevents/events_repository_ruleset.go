// Copyright 2022 The GithubEvents Authors. All rights reserved.
// Use of this source code is governed by the MIT License
// that can be found in the LICENSE file.

package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"fmt"
	"github.com/google/go-github/v69/github"
	"golang.org/x/sync/errgroup"
)

// Actions are used to identify registered callbacks.
const (
	// RepositoryRulesetEvent is the event name of github.RepositoryRulesetEvent's
	RepositoryRulesetEvent = "repository_ruleset"

	// RepositoryRulesetEventAnyAction is used to identify callbacks
	// listening to all events of type github.RepositoryRulesetEvent
	RepositoryRulesetEventAnyAction = "*"

	// RepositoryRulesetEventCreatedAction is used to identify callbacks
	// listening to events of type github.RepositoryRulesetEvent and action "created"
	RepositoryRulesetEventCreatedAction = "created"

	// RepositoryRulesetEventDeletedAction is used to identify callbacks
	// listening to events of type github.RepositoryRulesetEvent and action "deleted"
	RepositoryRulesetEventDeletedAction = "deleted"

	// RepositoryRulesetEventEditedAction is used to identify callbacks
	// listening to events of type github.RepositoryRulesetEvent and action "edited"
	RepositoryRulesetEventEditedAction = "edited"
)

// RepositoryRulesetEventHandleFunc represents a callback function triggered on github.RepositoryRulesetEvent's.
// 'deliveryID' (type: string) is the unique webhook delivery ID.
// 'eventName' (type: string) is the name of the event.
// 'event' (type: *github.RepositoryRulesetEvent) is the webhook payload.
type RepositoryRulesetEventHandleFunc func(deliveryID string, eventName string, event *github.RepositoryRulesetEvent) error

// OnRepositoryRulesetEventCreated registers callbacks listening to events of type github.RepositoryRulesetEvent and action 'created'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnRepositoryRulesetEventCreated must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#repository_ruleset
func (g *EventHandler) OnRepositoryRulesetEventCreated(callbacks ...RepositoryRulesetEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onRepositoryRulesetEvent == nil {
		g.onRepositoryRulesetEvent = make(map[string][]RepositoryRulesetEventHandleFunc)
	}
	g.onRepositoryRulesetEvent[RepositoryRulesetEventCreatedAction] = append(
		g.onRepositoryRulesetEvent[RepositoryRulesetEventCreatedAction],
		callbacks...,
	)
}

// SetOnRepositoryRulesetEventCreated registers callbacks listening to events of type github.RepositoryRulesetEvent and action 'created'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnRepositoryRulesetEventCreatedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#repository_ruleset
func (g *EventHandler) SetOnRepositoryRulesetEventCreated(callbacks ...RepositoryRulesetEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onRepositoryRulesetEvent == nil {
		g.onRepositoryRulesetEvent = make(map[string][]RepositoryRulesetEventHandleFunc)
	}
	g.onRepositoryRulesetEvent[RepositoryRulesetEventCreatedAction] = callbacks
}

func (g *EventHandler) handleRepositoryRulesetEventCreated(deliveryID string, eventName string, event *github.RepositoryRulesetEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if RepositoryRulesetEventCreatedAction != *event.Action {
		return fmt.Errorf(
			"handleRepositoryRulesetEventCreated() called with wrong action, want %s, got %s",
			RepositoryRulesetEventCreatedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		RepositoryRulesetEventCreatedAction,
		RepositoryRulesetEventAnyAction,
	} {
		if _, ok := g.onRepositoryRulesetEvent[action]; ok {
			for _, h := range g.onRepositoryRulesetEvent[action] {
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

// OnRepositoryRulesetEventDeleted registers callbacks listening to events of type github.RepositoryRulesetEvent and action 'deleted'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnRepositoryRulesetEventDeleted must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#repository_ruleset
func (g *EventHandler) OnRepositoryRulesetEventDeleted(callbacks ...RepositoryRulesetEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onRepositoryRulesetEvent == nil {
		g.onRepositoryRulesetEvent = make(map[string][]RepositoryRulesetEventHandleFunc)
	}
	g.onRepositoryRulesetEvent[RepositoryRulesetEventDeletedAction] = append(
		g.onRepositoryRulesetEvent[RepositoryRulesetEventDeletedAction],
		callbacks...,
	)
}

// SetOnRepositoryRulesetEventDeleted registers callbacks listening to events of type github.RepositoryRulesetEvent and action 'deleted'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnRepositoryRulesetEventDeletedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#repository_ruleset
func (g *EventHandler) SetOnRepositoryRulesetEventDeleted(callbacks ...RepositoryRulesetEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onRepositoryRulesetEvent == nil {
		g.onRepositoryRulesetEvent = make(map[string][]RepositoryRulesetEventHandleFunc)
	}
	g.onRepositoryRulesetEvent[RepositoryRulesetEventDeletedAction] = callbacks
}

func (g *EventHandler) handleRepositoryRulesetEventDeleted(deliveryID string, eventName string, event *github.RepositoryRulesetEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if RepositoryRulesetEventDeletedAction != *event.Action {
		return fmt.Errorf(
			"handleRepositoryRulesetEventDeleted() called with wrong action, want %s, got %s",
			RepositoryRulesetEventDeletedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		RepositoryRulesetEventDeletedAction,
		RepositoryRulesetEventAnyAction,
	} {
		if _, ok := g.onRepositoryRulesetEvent[action]; ok {
			for _, h := range g.onRepositoryRulesetEvent[action] {
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

// OnRepositoryRulesetEventEdited registers callbacks listening to events of type github.RepositoryRulesetEvent and action 'edited'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnRepositoryRulesetEventEdited must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#repository_ruleset
func (g *EventHandler) OnRepositoryRulesetEventEdited(callbacks ...RepositoryRulesetEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onRepositoryRulesetEvent == nil {
		g.onRepositoryRulesetEvent = make(map[string][]RepositoryRulesetEventHandleFunc)
	}
	g.onRepositoryRulesetEvent[RepositoryRulesetEventEditedAction] = append(
		g.onRepositoryRulesetEvent[RepositoryRulesetEventEditedAction],
		callbacks...,
	)
}

// SetOnRepositoryRulesetEventEdited registers callbacks listening to events of type github.RepositoryRulesetEvent and action 'edited'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnRepositoryRulesetEventEditedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#repository_ruleset
func (g *EventHandler) SetOnRepositoryRulesetEventEdited(callbacks ...RepositoryRulesetEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onRepositoryRulesetEvent == nil {
		g.onRepositoryRulesetEvent = make(map[string][]RepositoryRulesetEventHandleFunc)
	}
	g.onRepositoryRulesetEvent[RepositoryRulesetEventEditedAction] = callbacks
}

func (g *EventHandler) handleRepositoryRulesetEventEdited(deliveryID string, eventName string, event *github.RepositoryRulesetEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if RepositoryRulesetEventEditedAction != *event.Action {
		return fmt.Errorf(
			"handleRepositoryRulesetEventEdited() called with wrong action, want %s, got %s",
			RepositoryRulesetEventEditedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		RepositoryRulesetEventEditedAction,
		RepositoryRulesetEventAnyAction,
	} {
		if _, ok := g.onRepositoryRulesetEvent[action]; ok {
			for _, h := range g.onRepositoryRulesetEvent[action] {
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

// OnRepositoryRulesetEventAny registers callbacks listening to any events of type github.RepositoryRulesetEvent
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnRepositoryRulesetEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#repository_ruleset
func (g *EventHandler) OnRepositoryRulesetEventAny(callbacks ...RepositoryRulesetEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onRepositoryRulesetEvent == nil {
		g.onRepositoryRulesetEvent = make(map[string][]RepositoryRulesetEventHandleFunc)
	}
	g.onRepositoryRulesetEvent[RepositoryRulesetEventAnyAction] = append(
		g.onRepositoryRulesetEvent[RepositoryRulesetEventAnyAction],
		callbacks...,
	)
}

// SetOnRepositoryRulesetEventAny registers callbacks listening to any events of type github.RepositoryRulesetEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnRepositoryRulesetEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#repository_ruleset
func (g *EventHandler) SetOnRepositoryRulesetEventAny(callbacks ...RepositoryRulesetEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onRepositoryRulesetEvent == nil {
		g.onRepositoryRulesetEvent = make(map[string][]RepositoryRulesetEventHandleFunc)
	}
	g.onRepositoryRulesetEvent[RepositoryRulesetEventAnyAction] = callbacks
}

func (g *EventHandler) handleRepositoryRulesetEventAny(deliveryID string, eventName string, event *github.RepositoryRulesetEvent) error {
	if event == nil {
		return fmt.Errorf("event was empty or nil")
	}
	if _, ok := g.onRepositoryRulesetEvent[RepositoryRulesetEventAnyAction]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onRepositoryRulesetEvent[RepositoryRulesetEventAnyAction] {
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

// RepositoryRulesetEvent handles github.RepositoryRulesetEvent.
//
// Callbacks are executed in the following order:
//
// 1) All callbacks registered with OnBeforeAny are executed in parallel.
// 2) All callbacks registered with OnRepositoryRulesetEvent... are executed in parallel in case the Event has actions.
// 3) All callbacks registered with OnAfterAny are executed in parallel.
//
// on any error all callbacks registered with OnError are executed in parallel.
func (g *EventHandler) RepositoryRulesetEvent(deliveryID string, eventName string, event *github.RepositoryRulesetEvent) error {

	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	action := *event.Action

	err := g.handleBeforeAny(deliveryID, eventName, event)
	if err != nil {
		return g.handleError(deliveryID, eventName, event, err)
	}

	switch action {

	case RepositoryRulesetEventCreatedAction:
		err := g.handleRepositoryRulesetEventCreated(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case RepositoryRulesetEventDeletedAction:
		err := g.handleRepositoryRulesetEventDeleted(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case RepositoryRulesetEventEditedAction:
		err := g.handleRepositoryRulesetEventEdited(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	default:
		err := g.handleRepositoryRulesetEventAny(deliveryID, eventName, event)
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
