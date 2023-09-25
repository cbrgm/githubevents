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
	// MembershipEvent is the event name of github.MembershipEvent's
	MembershipEvent = "membership"

	// MembershipEventAnyAction is used to identify callbacks
	// listening to all events of type github.MembershipEvent
	MembershipEventAnyAction = "*"

	// MembershipEventAddedAction is used to identify callbacks
	// listening to events of type github.MembershipEvent and action "added"
	MembershipEventAddedAction = "added"

	// MembershipEventRemovedAction is used to identify callbacks
	// listening to events of type github.MembershipEvent and action "removed"
	MembershipEventRemovedAction = "removed"
)

// MembershipEventHandleFunc represents a callback function triggered on github.MembershipEvent's.
// 'deliveryID' (type: string) is the unique webhook delivery ID.
// 'eventName' (type: string) is the name of the event.
// 'event' (type: *github.MembershipEvent) is the webhook payload.
type MembershipEventHandleFunc func(deliveryID string, eventName string, event *github.MembershipEvent) error

// OnMembershipEventAdded registers callbacks listening to events of type github.MembershipEvent and action 'added'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnMembershipEventAdded must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#membership
func (g *EventHandler) OnMembershipEventAdded(callbacks ...MembershipEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onMembershipEvent == nil {
		g.onMembershipEvent = make(map[string][]MembershipEventHandleFunc)
	}
	g.onMembershipEvent[MembershipEventAddedAction] = append(
		g.onMembershipEvent[MembershipEventAddedAction],
		callbacks...,
	)
}

// SetOnMembershipEventAdded registers callbacks listening to events of type github.MembershipEvent and action 'added'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnMembershipEventAddedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#membership
func (g *EventHandler) SetOnMembershipEventAdded(callbacks ...MembershipEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onMembershipEvent == nil {
		g.onMembershipEvent = make(map[string][]MembershipEventHandleFunc)
	}
	g.onMembershipEvent[MembershipEventAddedAction] = callbacks
}

func (g *EventHandler) handleMembershipEventAdded(deliveryID string, eventName string, event *github.MembershipEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if MembershipEventAddedAction != *event.Action {
		return fmt.Errorf(
			"handleMembershipEventAdded() called with wrong action, want %s, got %s",
			MembershipEventAddedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		MembershipEventAddedAction,
		MembershipEventAnyAction,
	} {
		if _, ok := g.onMembershipEvent[action]; ok {
			for _, h := range g.onMembershipEvent[action] {
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

// OnMembershipEventRemoved registers callbacks listening to events of type github.MembershipEvent and action 'removed'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnMembershipEventRemoved must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#membership
func (g *EventHandler) OnMembershipEventRemoved(callbacks ...MembershipEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onMembershipEvent == nil {
		g.onMembershipEvent = make(map[string][]MembershipEventHandleFunc)
	}
	g.onMembershipEvent[MembershipEventRemovedAction] = append(
		g.onMembershipEvent[MembershipEventRemovedAction],
		callbacks...,
	)
}

// SetOnMembershipEventRemoved registers callbacks listening to events of type github.MembershipEvent and action 'removed'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnMembershipEventRemovedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#membership
func (g *EventHandler) SetOnMembershipEventRemoved(callbacks ...MembershipEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onMembershipEvent == nil {
		g.onMembershipEvent = make(map[string][]MembershipEventHandleFunc)
	}
	g.onMembershipEvent[MembershipEventRemovedAction] = callbacks
}

func (g *EventHandler) handleMembershipEventRemoved(deliveryID string, eventName string, event *github.MembershipEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if MembershipEventRemovedAction != *event.Action {
		return fmt.Errorf(
			"handleMembershipEventRemoved() called with wrong action, want %s, got %s",
			MembershipEventRemovedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		MembershipEventRemovedAction,
		MembershipEventAnyAction,
	} {
		if _, ok := g.onMembershipEvent[action]; ok {
			for _, h := range g.onMembershipEvent[action] {
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

// OnMembershipEventAny registers callbacks listening to any events of type github.MembershipEvent
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnMembershipEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#membership
func (g *EventHandler) OnMembershipEventAny(callbacks ...MembershipEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onMembershipEvent == nil {
		g.onMembershipEvent = make(map[string][]MembershipEventHandleFunc)
	}
	g.onMembershipEvent[MembershipEventAnyAction] = append(
		g.onMembershipEvent[MembershipEventAnyAction],
		callbacks...,
	)
}

// SetOnMembershipEventAny registers callbacks listening to any events of type github.MembershipEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnMembershipEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#membership
func (g *EventHandler) SetOnMembershipEventAny(callbacks ...MembershipEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onMembershipEvent == nil {
		g.onMembershipEvent = make(map[string][]MembershipEventHandleFunc)
	}
	g.onMembershipEvent[MembershipEventAnyAction] = callbacks
}

func (g *EventHandler) handleMembershipEventAny(deliveryID string, eventName string, event *github.MembershipEvent) error {
	if event == nil {
		return fmt.Errorf("event was empty or nil")
	}
	if _, ok := g.onMembershipEvent[MembershipEventAnyAction]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onMembershipEvent[MembershipEventAnyAction] {
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

// MembershipEvent handles github.MembershipEvent.
//
// Callbacks are executed in the following order:
//
// 1) All callbacks registered with OnBeforeAny are executed in parallel.
// 2) All callbacks registered with OnMembershipEvent... are executed in parallel in case the Event has actions.
// 3) All callbacks registered with OnAfterAny are executed in parallel.
//
// on any error all callbacks registered with OnError are executed in parallel.
func (g *EventHandler) MembershipEvent(deliveryID string, eventName string, event *github.MembershipEvent) error {

	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	action := *event.Action

	err := g.handleBeforeAny(deliveryID, eventName, event)
	if err != nil {
		return g.handleError(deliveryID, eventName, event, err)
	}

	switch action {

	case MembershipEventAddedAction:
		err := g.handleMembershipEventAdded(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case MembershipEventRemovedAction:
		err := g.handleMembershipEventRemoved(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	default:
		err := g.handleMembershipEventAny(deliveryID, eventName, event)
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
