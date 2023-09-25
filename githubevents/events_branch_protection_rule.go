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
	// BranchProtectionRuleEvent is the event name of github.BranchProtectionRuleEvent's
	BranchProtectionRuleEvent = "branch_protection_rule"

	// BranchProtectionRuleEventAnyAction is used to identify callbacks
	// listening to all events of type github.BranchProtectionRuleEvent
	BranchProtectionRuleEventAnyAction = "*"

	// BranchProtectionRuleEventCreatedAction is used to identify callbacks
	// listening to events of type github.BranchProtectionRuleEvent and action "created"
	BranchProtectionRuleEventCreatedAction = "created"

	// BranchProtectionRuleEventEditedAction is used to identify callbacks
	// listening to events of type github.BranchProtectionRuleEvent and action "edited"
	BranchProtectionRuleEventEditedAction = "edited"

	// BranchProtectionRuleEventDeletedAction is used to identify callbacks
	// listening to events of type github.BranchProtectionRuleEvent and action "deleted"
	BranchProtectionRuleEventDeletedAction = "deleted"
)

// BranchProtectionRuleEventHandleFunc represents a callback function triggered on github.BranchProtectionRuleEvent's.
// 'deliveryID' (type: string) is the unique webhook delivery ID.
// 'eventName' (type: string) is the name of the event.
// 'event' (type: *github.BranchProtectionRuleEvent) is the webhook payload.
type BranchProtectionRuleEventHandleFunc func(deliveryID string, eventName string, event *github.BranchProtectionRuleEvent) error

// OnBranchProtectionRuleEventCreated registers callbacks listening to events of type github.BranchProtectionRuleEvent and action 'created'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnBranchProtectionRuleEventCreated must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#branch_protection_rule
func (g *EventHandler) OnBranchProtectionRuleEventCreated(callbacks ...BranchProtectionRuleEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onBranchProtectionRuleEvent == nil {
		g.onBranchProtectionRuleEvent = make(map[string][]BranchProtectionRuleEventHandleFunc)
	}
	g.onBranchProtectionRuleEvent[BranchProtectionRuleEventCreatedAction] = append(
		g.onBranchProtectionRuleEvent[BranchProtectionRuleEventCreatedAction],
		callbacks...,
	)
}

// SetOnBranchProtectionRuleEventCreated registers callbacks listening to events of type github.BranchProtectionRuleEvent and action 'created'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnBranchProtectionRuleEventCreatedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#branch_protection_rule
func (g *EventHandler) SetOnBranchProtectionRuleEventCreated(callbacks ...BranchProtectionRuleEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onBranchProtectionRuleEvent == nil {
		g.onBranchProtectionRuleEvent = make(map[string][]BranchProtectionRuleEventHandleFunc)
	}
	g.onBranchProtectionRuleEvent[BranchProtectionRuleEventCreatedAction] = callbacks
}

func (g *EventHandler) handleBranchProtectionRuleEventCreated(deliveryID string, eventName string, event *github.BranchProtectionRuleEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if BranchProtectionRuleEventCreatedAction != *event.Action {
		return fmt.Errorf(
			"handleBranchProtectionRuleEventCreated() called with wrong action, want %s, got %s",
			BranchProtectionRuleEventCreatedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		BranchProtectionRuleEventCreatedAction,
		BranchProtectionRuleEventAnyAction,
	} {
		if _, ok := g.onBranchProtectionRuleEvent[action]; ok {
			for _, h := range g.onBranchProtectionRuleEvent[action] {
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

// OnBranchProtectionRuleEventEdited registers callbacks listening to events of type github.BranchProtectionRuleEvent and action 'edited'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnBranchProtectionRuleEventEdited must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#branch_protection_rule
func (g *EventHandler) OnBranchProtectionRuleEventEdited(callbacks ...BranchProtectionRuleEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onBranchProtectionRuleEvent == nil {
		g.onBranchProtectionRuleEvent = make(map[string][]BranchProtectionRuleEventHandleFunc)
	}
	g.onBranchProtectionRuleEvent[BranchProtectionRuleEventEditedAction] = append(
		g.onBranchProtectionRuleEvent[BranchProtectionRuleEventEditedAction],
		callbacks...,
	)
}

// SetOnBranchProtectionRuleEventEdited registers callbacks listening to events of type github.BranchProtectionRuleEvent and action 'edited'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnBranchProtectionRuleEventEditedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#branch_protection_rule
func (g *EventHandler) SetOnBranchProtectionRuleEventEdited(callbacks ...BranchProtectionRuleEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onBranchProtectionRuleEvent == nil {
		g.onBranchProtectionRuleEvent = make(map[string][]BranchProtectionRuleEventHandleFunc)
	}
	g.onBranchProtectionRuleEvent[BranchProtectionRuleEventEditedAction] = callbacks
}

func (g *EventHandler) handleBranchProtectionRuleEventEdited(deliveryID string, eventName string, event *github.BranchProtectionRuleEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if BranchProtectionRuleEventEditedAction != *event.Action {
		return fmt.Errorf(
			"handleBranchProtectionRuleEventEdited() called with wrong action, want %s, got %s",
			BranchProtectionRuleEventEditedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		BranchProtectionRuleEventEditedAction,
		BranchProtectionRuleEventAnyAction,
	} {
		if _, ok := g.onBranchProtectionRuleEvent[action]; ok {
			for _, h := range g.onBranchProtectionRuleEvent[action] {
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

// OnBranchProtectionRuleEventDeleted registers callbacks listening to events of type github.BranchProtectionRuleEvent and action 'deleted'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnBranchProtectionRuleEventDeleted must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#branch_protection_rule
func (g *EventHandler) OnBranchProtectionRuleEventDeleted(callbacks ...BranchProtectionRuleEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onBranchProtectionRuleEvent == nil {
		g.onBranchProtectionRuleEvent = make(map[string][]BranchProtectionRuleEventHandleFunc)
	}
	g.onBranchProtectionRuleEvent[BranchProtectionRuleEventDeletedAction] = append(
		g.onBranchProtectionRuleEvent[BranchProtectionRuleEventDeletedAction],
		callbacks...,
	)
}

// SetOnBranchProtectionRuleEventDeleted registers callbacks listening to events of type github.BranchProtectionRuleEvent and action 'deleted'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnBranchProtectionRuleEventDeletedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#branch_protection_rule
func (g *EventHandler) SetOnBranchProtectionRuleEventDeleted(callbacks ...BranchProtectionRuleEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onBranchProtectionRuleEvent == nil {
		g.onBranchProtectionRuleEvent = make(map[string][]BranchProtectionRuleEventHandleFunc)
	}
	g.onBranchProtectionRuleEvent[BranchProtectionRuleEventDeletedAction] = callbacks
}

func (g *EventHandler) handleBranchProtectionRuleEventDeleted(deliveryID string, eventName string, event *github.BranchProtectionRuleEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if BranchProtectionRuleEventDeletedAction != *event.Action {
		return fmt.Errorf(
			"handleBranchProtectionRuleEventDeleted() called with wrong action, want %s, got %s",
			BranchProtectionRuleEventDeletedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		BranchProtectionRuleEventDeletedAction,
		BranchProtectionRuleEventAnyAction,
	} {
		if _, ok := g.onBranchProtectionRuleEvent[action]; ok {
			for _, h := range g.onBranchProtectionRuleEvent[action] {
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

// OnBranchProtectionRuleEventAny registers callbacks listening to any events of type github.BranchProtectionRuleEvent
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnBranchProtectionRuleEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#branch_protection_rule
func (g *EventHandler) OnBranchProtectionRuleEventAny(callbacks ...BranchProtectionRuleEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onBranchProtectionRuleEvent == nil {
		g.onBranchProtectionRuleEvent = make(map[string][]BranchProtectionRuleEventHandleFunc)
	}
	g.onBranchProtectionRuleEvent[BranchProtectionRuleEventAnyAction] = append(
		g.onBranchProtectionRuleEvent[BranchProtectionRuleEventAnyAction],
		callbacks...,
	)
}

// SetOnBranchProtectionRuleEventAny registers callbacks listening to any events of type github.BranchProtectionRuleEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnBranchProtectionRuleEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#branch_protection_rule
func (g *EventHandler) SetOnBranchProtectionRuleEventAny(callbacks ...BranchProtectionRuleEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onBranchProtectionRuleEvent == nil {
		g.onBranchProtectionRuleEvent = make(map[string][]BranchProtectionRuleEventHandleFunc)
	}
	g.onBranchProtectionRuleEvent[BranchProtectionRuleEventAnyAction] = callbacks
}

func (g *EventHandler) handleBranchProtectionRuleEventAny(deliveryID string, eventName string, event *github.BranchProtectionRuleEvent) error {
	if event == nil {
		return fmt.Errorf("event was empty or nil")
	}
	if _, ok := g.onBranchProtectionRuleEvent[BranchProtectionRuleEventAnyAction]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onBranchProtectionRuleEvent[BranchProtectionRuleEventAnyAction] {
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

// BranchProtectionRuleEvent handles github.BranchProtectionRuleEvent.
//
// Callbacks are executed in the following order:
//
// 1) All callbacks registered with OnBeforeAny are executed in parallel.
// 2) All callbacks registered with OnBranchProtectionRuleEvent... are executed in parallel in case the Event has actions.
// 3) All callbacks registered with OnAfterAny are executed in parallel.
//
// on any error all callbacks registered with OnError are executed in parallel.
func (g *EventHandler) BranchProtectionRuleEvent(deliveryID string, eventName string, event *github.BranchProtectionRuleEvent) error {

	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	action := *event.Action

	err := g.handleBeforeAny(deliveryID, eventName, event)
	if err != nil {
		return g.handleError(deliveryID, eventName, event, err)
	}

	switch action {

	case BranchProtectionRuleEventCreatedAction:
		err := g.handleBranchProtectionRuleEventCreated(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case BranchProtectionRuleEventEditedAction:
		err := g.handleBranchProtectionRuleEventEdited(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case BranchProtectionRuleEventDeletedAction:
		err := g.handleBranchProtectionRuleEventDeleted(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	default:
		err := g.handleBranchProtectionRuleEventAny(deliveryID, eventName, event)
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
