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
	// MemberEvent is the event name of github.MemberEvent's
	MemberEvent = "member"

	// MemberEventAnyAction is used to identify callbacks
	// listening to all events of type github.MemberEvent
	MemberEventAnyAction = "*"

	// MemberEventAddedAction is used to identify callbacks
	// listening to events of type github.MemberEvent and action "added"
	MemberEventAddedAction = "added"

	// MemberEventRemovedAction is used to identify callbacks
	// listening to events of type github.MemberEvent and action "removed"
	MemberEventRemovedAction = "removed"

	// MemberEventEditedAction is used to identify callbacks
	// listening to events of type github.MemberEvent and action "edited"
	MemberEventEditedAction = "edited"
)

// MemberEventHandleFunc represents a callback function triggered on github.MemberEvent's.
// 'deliveryID' (type: string) is the unique webhook delivery ID.
// 'eventName' (type: string) is the name of the event.
// 'event' (type: *github.MemberEvent) is the webhook payload.
type MemberEventHandleFunc func(deliveryID string, eventName string, event *github.MemberEvent) error

// OnMemberEventAdded registers callbacks listening to events of type github.MemberEvent and action 'added'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnMemberEventAdded must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#member
func (g *EventHandler) OnMemberEventAdded(callbacks ...MemberEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onMemberEvent == nil {
		g.onMemberEvent = make(map[string][]MemberEventHandleFunc)
	}
	g.onMemberEvent[MemberEventAddedAction] = append(
		g.onMemberEvent[MemberEventAddedAction],
		callbacks...,
	)
}

// SetOnMemberEventAdded registers callbacks listening to events of type github.MemberEvent and action 'added'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnMemberEventAddedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#member
func (g *EventHandler) SetOnMemberEventAdded(callbacks ...MemberEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onMemberEvent == nil {
		g.onMemberEvent = make(map[string][]MemberEventHandleFunc)
	}
	g.onMemberEvent[MemberEventAddedAction] = callbacks
}

func (g *EventHandler) handleMemberEventAdded(deliveryID string, eventName string, event *github.MemberEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if MemberEventAddedAction != *event.Action {
		return fmt.Errorf(
			"handleMemberEventAdded() called with wrong action, want %s, got %s",
			MemberEventAddedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		MemberEventAddedAction,
		MemberEventAnyAction,
	} {
		if _, ok := g.onMemberEvent[action]; ok {
			for _, h := range g.onMemberEvent[action] {
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

// OnMemberEventRemoved registers callbacks listening to events of type github.MemberEvent and action 'removed'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnMemberEventRemoved must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#member
func (g *EventHandler) OnMemberEventRemoved(callbacks ...MemberEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onMemberEvent == nil {
		g.onMemberEvent = make(map[string][]MemberEventHandleFunc)
	}
	g.onMemberEvent[MemberEventRemovedAction] = append(
		g.onMemberEvent[MemberEventRemovedAction],
		callbacks...,
	)
}

// SetOnMemberEventRemoved registers callbacks listening to events of type github.MemberEvent and action 'removed'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnMemberEventRemovedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#member
func (g *EventHandler) SetOnMemberEventRemoved(callbacks ...MemberEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onMemberEvent == nil {
		g.onMemberEvent = make(map[string][]MemberEventHandleFunc)
	}
	g.onMemberEvent[MemberEventRemovedAction] = callbacks
}

func (g *EventHandler) handleMemberEventRemoved(deliveryID string, eventName string, event *github.MemberEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if MemberEventRemovedAction != *event.Action {
		return fmt.Errorf(
			"handleMemberEventRemoved() called with wrong action, want %s, got %s",
			MemberEventRemovedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		MemberEventRemovedAction,
		MemberEventAnyAction,
	} {
		if _, ok := g.onMemberEvent[action]; ok {
			for _, h := range g.onMemberEvent[action] {
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

// OnMemberEventEdited registers callbacks listening to events of type github.MemberEvent and action 'edited'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnMemberEventEdited must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#member
func (g *EventHandler) OnMemberEventEdited(callbacks ...MemberEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onMemberEvent == nil {
		g.onMemberEvent = make(map[string][]MemberEventHandleFunc)
	}
	g.onMemberEvent[MemberEventEditedAction] = append(
		g.onMemberEvent[MemberEventEditedAction],
		callbacks...,
	)
}

// SetOnMemberEventEdited registers callbacks listening to events of type github.MemberEvent and action 'edited'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnMemberEventEditedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#member
func (g *EventHandler) SetOnMemberEventEdited(callbacks ...MemberEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onMemberEvent == nil {
		g.onMemberEvent = make(map[string][]MemberEventHandleFunc)
	}
	g.onMemberEvent[MemberEventEditedAction] = callbacks
}

func (g *EventHandler) handleMemberEventEdited(deliveryID string, eventName string, event *github.MemberEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if MemberEventEditedAction != *event.Action {
		return fmt.Errorf(
			"handleMemberEventEdited() called with wrong action, want %s, got %s",
			MemberEventEditedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		MemberEventEditedAction,
		MemberEventAnyAction,
	} {
		if _, ok := g.onMemberEvent[action]; ok {
			for _, h := range g.onMemberEvent[action] {
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

// OnMemberEventAny registers callbacks listening to any events of type github.MemberEvent
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnMemberEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#member
func (g *EventHandler) OnMemberEventAny(callbacks ...MemberEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onMemberEvent == nil {
		g.onMemberEvent = make(map[string][]MemberEventHandleFunc)
	}
	g.onMemberEvent[MemberEventAnyAction] = append(
		g.onMemberEvent[MemberEventAnyAction],
		callbacks...,
	)
}

// SetOnMemberEventAny registers callbacks listening to any events of type github.MemberEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnMemberEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#member
func (g *EventHandler) SetOnMemberEventAny(callbacks ...MemberEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onMemberEvent == nil {
		g.onMemberEvent = make(map[string][]MemberEventHandleFunc)
	}
	g.onMemberEvent[MemberEventAnyAction] = callbacks
}

func (g *EventHandler) handleMemberEventAny(deliveryID string, eventName string, event *github.MemberEvent) error {
	if event == nil {
		return fmt.Errorf("event was empty or nil")
	}
	if _, ok := g.onMemberEvent[MemberEventAnyAction]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onMemberEvent[MemberEventAnyAction] {
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

// MemberEvent handles github.MemberEvent.
//
// Callbacks are executed in the following order:
//
// 1) All callbacks registered with OnBeforeAny are executed in parallel.
// 2) All callbacks registered with OnMemberEvent... are executed in parallel in case the Event has actions.
// 3) All callbacks registered with OnAfterAny are executed in parallel.
//
// on any error all callbacks registered with OnError are executed in parallel.
func (g *EventHandler) MemberEvent(deliveryID string, eventName string, event *github.MemberEvent) error {

	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	action := *event.Action

	err := g.handleBeforeAny(deliveryID, eventName, event)
	if err != nil {
		return g.handleError(deliveryID, eventName, event, err)
	}

	switch action {

	case MemberEventAddedAction:
		err := g.handleMemberEventAdded(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case MemberEventRemovedAction:
		err := g.handleMemberEventRemoved(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case MemberEventEditedAction:
		err := g.handleMemberEventEdited(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	default:
		err := g.handleMemberEventAny(deliveryID, eventName, event)
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
