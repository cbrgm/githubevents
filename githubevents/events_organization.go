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
	// OrganizationEvent is the event name of github.OrganizationEvent's
	OrganizationEvent = "organization"

	// OrganizationEventAnyAction is used to identify callbacks
	// listening to all events of type github.OrganizationEvent
	OrganizationEventAnyAction = "*"

	// OrganizationEventDeletedAction is used to identify callbacks
	// listening to events of type github.OrganizationEvent and action "deleted"
	OrganizationEventDeletedAction = "deleted"

	// OrganizationEventRenamedAction is used to identify callbacks
	// listening to events of type github.OrganizationEvent and action "renamed"
	OrganizationEventRenamedAction = "renamed"

	// OrganizationEventMemberAddedAction is used to identify callbacks
	// listening to events of type github.OrganizationEvent and action "member_added"
	OrganizationEventMemberAddedAction = "member_added"

	// OrganizationEventMemberRemovedAction is used to identify callbacks
	// listening to events of type github.OrganizationEvent and action "member_removed"
	OrganizationEventMemberRemovedAction = "member_removed"

	// OrganizationEventMemberInvitedAction is used to identify callbacks
	// listening to events of type github.OrganizationEvent and action "member_invited"
	OrganizationEventMemberInvitedAction = "member_invited"
)

// OrganizationEventHandleFunc represents a callback function triggered on github.OrganizationEvent's.
// 'deliveryID' (type: string) is the unique webhook delivery ID.
// 'eventName' (type: string) is the name of the event.
// 'event' (type: *github.OrganizationEvent) is the webhook payload.
type OrganizationEventHandleFunc func(deliveryID string, eventName string, event *github.OrganizationEvent) error

// OnOrganizationEventDeleted registers callbacks listening to events of type github.OrganizationEvent and action 'deleted'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnOrganizationEventDeleted must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#organization
func (g *EventHandler) OnOrganizationEventDeleted(callbacks ...OrganizationEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onOrganizationEvent == nil {
		g.onOrganizationEvent = make(map[string][]OrganizationEventHandleFunc)
	}
	g.onOrganizationEvent[OrganizationEventDeletedAction] = append(
		g.onOrganizationEvent[OrganizationEventDeletedAction],
		callbacks...,
	)
}

// SetOnOrganizationEventDeleted registers callbacks listening to events of type github.OrganizationEvent and action 'deleted'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnOrganizationEventDeletedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#organization
func (g *EventHandler) SetOnOrganizationEventDeleted(callbacks ...OrganizationEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onOrganizationEvent == nil {
		g.onOrganizationEvent = make(map[string][]OrganizationEventHandleFunc)
	}
	g.onOrganizationEvent[OrganizationEventDeletedAction] = callbacks
}

func (g *EventHandler) handleOrganizationEventDeleted(deliveryID string, eventName string, event *github.OrganizationEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if OrganizationEventDeletedAction != *event.Action {
		return fmt.Errorf(
			"handleOrganizationEventDeleted() called with wrong action, want %s, got %s",
			OrganizationEventDeletedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		OrganizationEventDeletedAction,
		OrganizationEventAnyAction,
	} {
		if _, ok := g.onOrganizationEvent[action]; ok {
			for _, h := range g.onOrganizationEvent[action] {
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

// OnOrganizationEventRenamed registers callbacks listening to events of type github.OrganizationEvent and action 'renamed'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnOrganizationEventRenamed must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#organization
func (g *EventHandler) OnOrganizationEventRenamed(callbacks ...OrganizationEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onOrganizationEvent == nil {
		g.onOrganizationEvent = make(map[string][]OrganizationEventHandleFunc)
	}
	g.onOrganizationEvent[OrganizationEventRenamedAction] = append(
		g.onOrganizationEvent[OrganizationEventRenamedAction],
		callbacks...,
	)
}

// SetOnOrganizationEventRenamed registers callbacks listening to events of type github.OrganizationEvent and action 'renamed'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnOrganizationEventRenamedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#organization
func (g *EventHandler) SetOnOrganizationEventRenamed(callbacks ...OrganizationEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onOrganizationEvent == nil {
		g.onOrganizationEvent = make(map[string][]OrganizationEventHandleFunc)
	}
	g.onOrganizationEvent[OrganizationEventRenamedAction] = callbacks
}

func (g *EventHandler) handleOrganizationEventRenamed(deliveryID string, eventName string, event *github.OrganizationEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if OrganizationEventRenamedAction != *event.Action {
		return fmt.Errorf(
			"handleOrganizationEventRenamed() called with wrong action, want %s, got %s",
			OrganizationEventRenamedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		OrganizationEventRenamedAction,
		OrganizationEventAnyAction,
	} {
		if _, ok := g.onOrganizationEvent[action]; ok {
			for _, h := range g.onOrganizationEvent[action] {
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

// OnOrganizationEventMemberAdded registers callbacks listening to events of type github.OrganizationEvent and action 'member_added'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnOrganizationEventMemberAdded must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#organization
func (g *EventHandler) OnOrganizationEventMemberAdded(callbacks ...OrganizationEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onOrganizationEvent == nil {
		g.onOrganizationEvent = make(map[string][]OrganizationEventHandleFunc)
	}
	g.onOrganizationEvent[OrganizationEventMemberAddedAction] = append(
		g.onOrganizationEvent[OrganizationEventMemberAddedAction],
		callbacks...,
	)
}

// SetOnOrganizationEventMemberAdded registers callbacks listening to events of type github.OrganizationEvent and action 'member_added'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnOrganizationEventMemberAddedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#organization
func (g *EventHandler) SetOnOrganizationEventMemberAdded(callbacks ...OrganizationEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onOrganizationEvent == nil {
		g.onOrganizationEvent = make(map[string][]OrganizationEventHandleFunc)
	}
	g.onOrganizationEvent[OrganizationEventMemberAddedAction] = callbacks
}

func (g *EventHandler) handleOrganizationEventMemberAdded(deliveryID string, eventName string, event *github.OrganizationEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if OrganizationEventMemberAddedAction != *event.Action {
		return fmt.Errorf(
			"handleOrganizationEventMemberAdded() called with wrong action, want %s, got %s",
			OrganizationEventMemberAddedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		OrganizationEventMemberAddedAction,
		OrganizationEventAnyAction,
	} {
		if _, ok := g.onOrganizationEvent[action]; ok {
			for _, h := range g.onOrganizationEvent[action] {
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

// OnOrganizationEventMemberRemoved registers callbacks listening to events of type github.OrganizationEvent and action 'member_removed'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnOrganizationEventMemberRemoved must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#organization
func (g *EventHandler) OnOrganizationEventMemberRemoved(callbacks ...OrganizationEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onOrganizationEvent == nil {
		g.onOrganizationEvent = make(map[string][]OrganizationEventHandleFunc)
	}
	g.onOrganizationEvent[OrganizationEventMemberRemovedAction] = append(
		g.onOrganizationEvent[OrganizationEventMemberRemovedAction],
		callbacks...,
	)
}

// SetOnOrganizationEventMemberRemoved registers callbacks listening to events of type github.OrganizationEvent and action 'member_removed'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnOrganizationEventMemberRemovedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#organization
func (g *EventHandler) SetOnOrganizationEventMemberRemoved(callbacks ...OrganizationEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onOrganizationEvent == nil {
		g.onOrganizationEvent = make(map[string][]OrganizationEventHandleFunc)
	}
	g.onOrganizationEvent[OrganizationEventMemberRemovedAction] = callbacks
}

func (g *EventHandler) handleOrganizationEventMemberRemoved(deliveryID string, eventName string, event *github.OrganizationEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if OrganizationEventMemberRemovedAction != *event.Action {
		return fmt.Errorf(
			"handleOrganizationEventMemberRemoved() called with wrong action, want %s, got %s",
			OrganizationEventMemberRemovedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		OrganizationEventMemberRemovedAction,
		OrganizationEventAnyAction,
	} {
		if _, ok := g.onOrganizationEvent[action]; ok {
			for _, h := range g.onOrganizationEvent[action] {
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

// OnOrganizationEventMemberInvited registers callbacks listening to events of type github.OrganizationEvent and action 'member_invited'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnOrganizationEventMemberInvited must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#organization
func (g *EventHandler) OnOrganizationEventMemberInvited(callbacks ...OrganizationEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onOrganizationEvent == nil {
		g.onOrganizationEvent = make(map[string][]OrganizationEventHandleFunc)
	}
	g.onOrganizationEvent[OrganizationEventMemberInvitedAction] = append(
		g.onOrganizationEvent[OrganizationEventMemberInvitedAction],
		callbacks...,
	)
}

// SetOnOrganizationEventMemberInvited registers callbacks listening to events of type github.OrganizationEvent and action 'member_invited'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnOrganizationEventMemberInvitedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#organization
func (g *EventHandler) SetOnOrganizationEventMemberInvited(callbacks ...OrganizationEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onOrganizationEvent == nil {
		g.onOrganizationEvent = make(map[string][]OrganizationEventHandleFunc)
	}
	g.onOrganizationEvent[OrganizationEventMemberInvitedAction] = callbacks
}

func (g *EventHandler) handleOrganizationEventMemberInvited(deliveryID string, eventName string, event *github.OrganizationEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if OrganizationEventMemberInvitedAction != *event.Action {
		return fmt.Errorf(
			"handleOrganizationEventMemberInvited() called with wrong action, want %s, got %s",
			OrganizationEventMemberInvitedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		OrganizationEventMemberInvitedAction,
		OrganizationEventAnyAction,
	} {
		if _, ok := g.onOrganizationEvent[action]; ok {
			for _, h := range g.onOrganizationEvent[action] {
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

// OnOrganizationEventAny registers callbacks listening to any events of type github.OrganizationEvent
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnOrganizationEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#organization
func (g *EventHandler) OnOrganizationEventAny(callbacks ...OrganizationEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onOrganizationEvent == nil {
		g.onOrganizationEvent = make(map[string][]OrganizationEventHandleFunc)
	}
	g.onOrganizationEvent[OrganizationEventAnyAction] = append(
		g.onOrganizationEvent[OrganizationEventAnyAction],
		callbacks...,
	)
}

// SetOnOrganizationEventAny registers callbacks listening to any events of type github.OrganizationEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnOrganizationEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#organization
func (g *EventHandler) SetOnOrganizationEventAny(callbacks ...OrganizationEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onOrganizationEvent == nil {
		g.onOrganizationEvent = make(map[string][]OrganizationEventHandleFunc)
	}
	g.onOrganizationEvent[OrganizationEventAnyAction] = callbacks
}

func (g *EventHandler) handleOrganizationEventAny(deliveryID string, eventName string, event *github.OrganizationEvent) error {
	if event == nil {
		return fmt.Errorf("event was empty or nil")
	}
	if _, ok := g.onOrganizationEvent[OrganizationEventAnyAction]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onOrganizationEvent[OrganizationEventAnyAction] {
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

// OrganizationEvent handles github.OrganizationEvent.
//
// Callbacks are executed in the following order:
//
// 1) All callbacks registered with OnBeforeAny are executed in parallel.
// 2) All callbacks registered with OnOrganizationEvent... are executed in parallel in case the Event has actions.
// 3) All callbacks registered with OnAfterAny are executed in parallel.
//
// on any error all callbacks registered with OnError are executed in parallel.
func (g *EventHandler) OrganizationEvent(deliveryID string, eventName string, event *github.OrganizationEvent) error {

	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	action := *event.Action

	err := g.handleBeforeAny(deliveryID, eventName, event)
	if err != nil {
		return g.handleError(deliveryID, eventName, event, err)
	}

	switch action {

	case OrganizationEventDeletedAction:
		err := g.handleOrganizationEventDeleted(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case OrganizationEventRenamedAction:
		err := g.handleOrganizationEventRenamed(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case OrganizationEventMemberAddedAction:
		err := g.handleOrganizationEventMemberAdded(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case OrganizationEventMemberRemovedAction:
		err := g.handleOrganizationEventMemberRemoved(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case OrganizationEventMemberInvitedAction:
		err := g.handleOrganizationEventMemberInvited(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	default:
		err := g.handleOrganizationEventAny(deliveryID, eventName, event)
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
