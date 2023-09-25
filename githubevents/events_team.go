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
	// TeamEvent is the event name of github.TeamEvent's
	TeamEvent = "team"

	// TeamEventAnyAction is used to identify callbacks
	// listening to all events of type github.TeamEvent
	TeamEventAnyAction = "*"

	// TeamEventCreatedAction is used to identify callbacks
	// listening to events of type github.TeamEvent and action "created"
	TeamEventCreatedAction = "created"

	// TeamEventDeletedAction is used to identify callbacks
	// listening to events of type github.TeamEvent and action "deleted"
	TeamEventDeletedAction = "deleted"

	// TeamEventEditedAction is used to identify callbacks
	// listening to events of type github.TeamEvent and action "edited"
	TeamEventEditedAction = "edited"

	// TeamEventAddedToRepositoryAction is used to identify callbacks
	// listening to events of type github.TeamEvent and action "added_to_repository"
	TeamEventAddedToRepositoryAction = "added_to_repository"

	// TeamEventRemovedFromRepositoryAction is used to identify callbacks
	// listening to events of type github.TeamEvent and action "removed_from_repository"
	TeamEventRemovedFromRepositoryAction = "removed_from_repository"
)

// TeamEventHandleFunc represents a callback function triggered on github.TeamEvent's.
// 'deliveryID' (type: string) is the unique webhook delivery ID.
// 'eventName' (type: string) is the name of the event.
// 'event' (type: *github.TeamEvent) is the webhook payload.
type TeamEventHandleFunc func(deliveryID string, eventName string, event *github.TeamEvent) error

// OnTeamEventCreated registers callbacks listening to events of type github.TeamEvent and action 'created'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnTeamEventCreated must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#team
func (g *EventHandler) OnTeamEventCreated(callbacks ...TeamEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onTeamEvent == nil {
		g.onTeamEvent = make(map[string][]TeamEventHandleFunc)
	}
	g.onTeamEvent[TeamEventCreatedAction] = append(
		g.onTeamEvent[TeamEventCreatedAction],
		callbacks...,
	)
}

// SetOnTeamEventCreated registers callbacks listening to events of type github.TeamEvent and action 'created'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnTeamEventCreatedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#team
func (g *EventHandler) SetOnTeamEventCreated(callbacks ...TeamEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onTeamEvent == nil {
		g.onTeamEvent = make(map[string][]TeamEventHandleFunc)
	}
	g.onTeamEvent[TeamEventCreatedAction] = callbacks
}

func (g *EventHandler) handleTeamEventCreated(deliveryID string, eventName string, event *github.TeamEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if TeamEventCreatedAction != *event.Action {
		return fmt.Errorf(
			"handleTeamEventCreated() called with wrong action, want %s, got %s",
			TeamEventCreatedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		TeamEventCreatedAction,
		TeamEventAnyAction,
	} {
		if _, ok := g.onTeamEvent[action]; ok {
			for _, h := range g.onTeamEvent[action] {
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

// OnTeamEventDeleted registers callbacks listening to events of type github.TeamEvent and action 'deleted'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnTeamEventDeleted must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#team
func (g *EventHandler) OnTeamEventDeleted(callbacks ...TeamEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onTeamEvent == nil {
		g.onTeamEvent = make(map[string][]TeamEventHandleFunc)
	}
	g.onTeamEvent[TeamEventDeletedAction] = append(
		g.onTeamEvent[TeamEventDeletedAction],
		callbacks...,
	)
}

// SetOnTeamEventDeleted registers callbacks listening to events of type github.TeamEvent and action 'deleted'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnTeamEventDeletedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#team
func (g *EventHandler) SetOnTeamEventDeleted(callbacks ...TeamEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onTeamEvent == nil {
		g.onTeamEvent = make(map[string][]TeamEventHandleFunc)
	}
	g.onTeamEvent[TeamEventDeletedAction] = callbacks
}

func (g *EventHandler) handleTeamEventDeleted(deliveryID string, eventName string, event *github.TeamEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if TeamEventDeletedAction != *event.Action {
		return fmt.Errorf(
			"handleTeamEventDeleted() called with wrong action, want %s, got %s",
			TeamEventDeletedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		TeamEventDeletedAction,
		TeamEventAnyAction,
	} {
		if _, ok := g.onTeamEvent[action]; ok {
			for _, h := range g.onTeamEvent[action] {
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

// OnTeamEventEdited registers callbacks listening to events of type github.TeamEvent and action 'edited'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnTeamEventEdited must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#team
func (g *EventHandler) OnTeamEventEdited(callbacks ...TeamEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onTeamEvent == nil {
		g.onTeamEvent = make(map[string][]TeamEventHandleFunc)
	}
	g.onTeamEvent[TeamEventEditedAction] = append(
		g.onTeamEvent[TeamEventEditedAction],
		callbacks...,
	)
}

// SetOnTeamEventEdited registers callbacks listening to events of type github.TeamEvent and action 'edited'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnTeamEventEditedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#team
func (g *EventHandler) SetOnTeamEventEdited(callbacks ...TeamEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onTeamEvent == nil {
		g.onTeamEvent = make(map[string][]TeamEventHandleFunc)
	}
	g.onTeamEvent[TeamEventEditedAction] = callbacks
}

func (g *EventHandler) handleTeamEventEdited(deliveryID string, eventName string, event *github.TeamEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if TeamEventEditedAction != *event.Action {
		return fmt.Errorf(
			"handleTeamEventEdited() called with wrong action, want %s, got %s",
			TeamEventEditedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		TeamEventEditedAction,
		TeamEventAnyAction,
	} {
		if _, ok := g.onTeamEvent[action]; ok {
			for _, h := range g.onTeamEvent[action] {
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

// OnTeamEventAddedToRepository registers callbacks listening to events of type github.TeamEvent and action 'added_to_repository'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnTeamEventAddedToRepository must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#team
func (g *EventHandler) OnTeamEventAddedToRepository(callbacks ...TeamEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onTeamEvent == nil {
		g.onTeamEvent = make(map[string][]TeamEventHandleFunc)
	}
	g.onTeamEvent[TeamEventAddedToRepositoryAction] = append(
		g.onTeamEvent[TeamEventAddedToRepositoryAction],
		callbacks...,
	)
}

// SetOnTeamEventAddedToRepository registers callbacks listening to events of type github.TeamEvent and action 'added_to_repository'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnTeamEventAddedToRepositoryAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#team
func (g *EventHandler) SetOnTeamEventAddedToRepository(callbacks ...TeamEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onTeamEvent == nil {
		g.onTeamEvent = make(map[string][]TeamEventHandleFunc)
	}
	g.onTeamEvent[TeamEventAddedToRepositoryAction] = callbacks
}

func (g *EventHandler) handleTeamEventAddedToRepository(deliveryID string, eventName string, event *github.TeamEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if TeamEventAddedToRepositoryAction != *event.Action {
		return fmt.Errorf(
			"handleTeamEventAddedToRepository() called with wrong action, want %s, got %s",
			TeamEventAddedToRepositoryAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		TeamEventAddedToRepositoryAction,
		TeamEventAnyAction,
	} {
		if _, ok := g.onTeamEvent[action]; ok {
			for _, h := range g.onTeamEvent[action] {
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

// OnTeamEventRemovedFromRepository registers callbacks listening to events of type github.TeamEvent and action 'removed_from_repository'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnTeamEventRemovedFromRepository must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#team
func (g *EventHandler) OnTeamEventRemovedFromRepository(callbacks ...TeamEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onTeamEvent == nil {
		g.onTeamEvent = make(map[string][]TeamEventHandleFunc)
	}
	g.onTeamEvent[TeamEventRemovedFromRepositoryAction] = append(
		g.onTeamEvent[TeamEventRemovedFromRepositoryAction],
		callbacks...,
	)
}

// SetOnTeamEventRemovedFromRepository registers callbacks listening to events of type github.TeamEvent and action 'removed_from_repository'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnTeamEventRemovedFromRepositoryAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#team
func (g *EventHandler) SetOnTeamEventRemovedFromRepository(callbacks ...TeamEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onTeamEvent == nil {
		g.onTeamEvent = make(map[string][]TeamEventHandleFunc)
	}
	g.onTeamEvent[TeamEventRemovedFromRepositoryAction] = callbacks
}

func (g *EventHandler) handleTeamEventRemovedFromRepository(deliveryID string, eventName string, event *github.TeamEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if TeamEventRemovedFromRepositoryAction != *event.Action {
		return fmt.Errorf(
			"handleTeamEventRemovedFromRepository() called with wrong action, want %s, got %s",
			TeamEventRemovedFromRepositoryAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		TeamEventRemovedFromRepositoryAction,
		TeamEventAnyAction,
	} {
		if _, ok := g.onTeamEvent[action]; ok {
			for _, h := range g.onTeamEvent[action] {
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

// OnTeamEventAny registers callbacks listening to any events of type github.TeamEvent
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnTeamEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#team
func (g *EventHandler) OnTeamEventAny(callbacks ...TeamEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onTeamEvent == nil {
		g.onTeamEvent = make(map[string][]TeamEventHandleFunc)
	}
	g.onTeamEvent[TeamEventAnyAction] = append(
		g.onTeamEvent[TeamEventAnyAction],
		callbacks...,
	)
}

// SetOnTeamEventAny registers callbacks listening to any events of type github.TeamEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnTeamEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#team
func (g *EventHandler) SetOnTeamEventAny(callbacks ...TeamEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onTeamEvent == nil {
		g.onTeamEvent = make(map[string][]TeamEventHandleFunc)
	}
	g.onTeamEvent[TeamEventAnyAction] = callbacks
}

func (g *EventHandler) handleTeamEventAny(deliveryID string, eventName string, event *github.TeamEvent) error {
	if event == nil {
		return fmt.Errorf("event was empty or nil")
	}
	if _, ok := g.onTeamEvent[TeamEventAnyAction]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onTeamEvent[TeamEventAnyAction] {
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

// TeamEvent handles github.TeamEvent.
//
// Callbacks are executed in the following order:
//
// 1) All callbacks registered with OnBeforeAny are executed in parallel.
// 2) All callbacks registered with OnTeamEvent... are executed in parallel in case the Event has actions.
// 3) All callbacks registered with OnAfterAny are executed in parallel.
//
// on any error all callbacks registered with OnError are executed in parallel.
func (g *EventHandler) TeamEvent(deliveryID string, eventName string, event *github.TeamEvent) error {

	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	action := *event.Action

	err := g.handleBeforeAny(deliveryID, eventName, event)
	if err != nil {
		return g.handleError(deliveryID, eventName, event, err)
	}

	switch action {

	case TeamEventCreatedAction:
		err := g.handleTeamEventCreated(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case TeamEventDeletedAction:
		err := g.handleTeamEventDeleted(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case TeamEventEditedAction:
		err := g.handleTeamEventEdited(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case TeamEventAddedToRepositoryAction:
		err := g.handleTeamEventAddedToRepository(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case TeamEventRemovedFromRepositoryAction:
		err := g.handleTeamEventRemovedFromRepository(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	default:
		err := g.handleTeamEventAny(deliveryID, eventName, event)
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
