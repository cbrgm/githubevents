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
	// ProjectCardEvent is the event name of github.ProjectCardEvent's
	ProjectCardEvent = "project_card"

	// ProjectCardEventAnyAction is used to identify callbacks
	// listening to all events of type github.ProjectCardEvent
	ProjectCardEventAnyAction = "*"

	// ProjectCardEventCreatedAction is used to identify callbacks
	// listening to events of type github.ProjectCardEvent and action "created"
	ProjectCardEventCreatedAction = "created"

	// ProjectCardEventEditedAction is used to identify callbacks
	// listening to events of type github.ProjectCardEvent and action "edited"
	ProjectCardEventEditedAction = "edited"

	// ProjectCardEventConvertedAction is used to identify callbacks
	// listening to events of type github.ProjectCardEvent and action "converted"
	ProjectCardEventConvertedAction = "converted"

	// ProjectCardEventMovedAction is used to identify callbacks
	// listening to events of type github.ProjectCardEvent and action "moved"
	ProjectCardEventMovedAction = "moved"

	// ProjectCardEventDeletedAction is used to identify callbacks
	// listening to events of type github.ProjectCardEvent and action "deleted"
	ProjectCardEventDeletedAction = "deleted"
)

// ProjectCardEventHandleFunc represents a callback function triggered on github.ProjectCardEvent's.
// 'deliveryID' (type: string) is the unique webhook delivery ID.
// 'eventName' (type: string) is the name of the event.
// 'event' (type: *github.ProjectCardEvent) is the webhook payload.
type ProjectCardEventHandleFunc func(deliveryID string, eventName string, event *github.ProjectCardEvent) error

// OnProjectCardEventCreated registers callbacks listening to events of type github.ProjectCardEvent and action 'created'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnProjectCardEventCreated must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#project_card
func (g *EventHandler) OnProjectCardEventCreated(callbacks ...ProjectCardEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectCardEvent == nil {
		g.onProjectCardEvent = make(map[string][]ProjectCardEventHandleFunc)
	}
	g.onProjectCardEvent[ProjectCardEventCreatedAction] = append(
		g.onProjectCardEvent[ProjectCardEventCreatedAction],
		callbacks...,
	)
}

// SetOnProjectCardEventCreated registers callbacks listening to events of type github.ProjectCardEvent and action 'created'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnProjectCardEventCreatedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#project_card
func (g *EventHandler) SetOnProjectCardEventCreated(callbacks ...ProjectCardEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectCardEvent == nil {
		g.onProjectCardEvent = make(map[string][]ProjectCardEventHandleFunc)
	}
	g.onProjectCardEvent[ProjectCardEventCreatedAction] = callbacks
}

func (g *EventHandler) handleProjectCardEventCreated(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if ProjectCardEventCreatedAction != *event.Action {
		return fmt.Errorf(
			"handleProjectCardEventCreated() called with wrong action, want %s, got %s",
			ProjectCardEventCreatedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		ProjectCardEventCreatedAction,
		ProjectCardEventAnyAction,
	} {
		if _, ok := g.onProjectCardEvent[action]; ok {
			for _, h := range g.onProjectCardEvent[action] {
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

// OnProjectCardEventEdited registers callbacks listening to events of type github.ProjectCardEvent and action 'edited'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnProjectCardEventEdited must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#project_card
func (g *EventHandler) OnProjectCardEventEdited(callbacks ...ProjectCardEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectCardEvent == nil {
		g.onProjectCardEvent = make(map[string][]ProjectCardEventHandleFunc)
	}
	g.onProjectCardEvent[ProjectCardEventEditedAction] = append(
		g.onProjectCardEvent[ProjectCardEventEditedAction],
		callbacks...,
	)
}

// SetOnProjectCardEventEdited registers callbacks listening to events of type github.ProjectCardEvent and action 'edited'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnProjectCardEventEditedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#project_card
func (g *EventHandler) SetOnProjectCardEventEdited(callbacks ...ProjectCardEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectCardEvent == nil {
		g.onProjectCardEvent = make(map[string][]ProjectCardEventHandleFunc)
	}
	g.onProjectCardEvent[ProjectCardEventEditedAction] = callbacks
}

func (g *EventHandler) handleProjectCardEventEdited(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if ProjectCardEventEditedAction != *event.Action {
		return fmt.Errorf(
			"handleProjectCardEventEdited() called with wrong action, want %s, got %s",
			ProjectCardEventEditedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		ProjectCardEventEditedAction,
		ProjectCardEventAnyAction,
	} {
		if _, ok := g.onProjectCardEvent[action]; ok {
			for _, h := range g.onProjectCardEvent[action] {
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

// OnProjectCardEventConverted registers callbacks listening to events of type github.ProjectCardEvent and action 'converted'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnProjectCardEventConverted must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#project_card
func (g *EventHandler) OnProjectCardEventConverted(callbacks ...ProjectCardEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectCardEvent == nil {
		g.onProjectCardEvent = make(map[string][]ProjectCardEventHandleFunc)
	}
	g.onProjectCardEvent[ProjectCardEventConvertedAction] = append(
		g.onProjectCardEvent[ProjectCardEventConvertedAction],
		callbacks...,
	)
}

// SetOnProjectCardEventConverted registers callbacks listening to events of type github.ProjectCardEvent and action 'converted'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnProjectCardEventConvertedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#project_card
func (g *EventHandler) SetOnProjectCardEventConverted(callbacks ...ProjectCardEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectCardEvent == nil {
		g.onProjectCardEvent = make(map[string][]ProjectCardEventHandleFunc)
	}
	g.onProjectCardEvent[ProjectCardEventConvertedAction] = callbacks
}

func (g *EventHandler) handleProjectCardEventConverted(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if ProjectCardEventConvertedAction != *event.Action {
		return fmt.Errorf(
			"handleProjectCardEventConverted() called with wrong action, want %s, got %s",
			ProjectCardEventConvertedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		ProjectCardEventConvertedAction,
		ProjectCardEventAnyAction,
	} {
		if _, ok := g.onProjectCardEvent[action]; ok {
			for _, h := range g.onProjectCardEvent[action] {
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

// OnProjectCardEventMoved registers callbacks listening to events of type github.ProjectCardEvent and action 'moved'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnProjectCardEventMoved must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#project_card
func (g *EventHandler) OnProjectCardEventMoved(callbacks ...ProjectCardEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectCardEvent == nil {
		g.onProjectCardEvent = make(map[string][]ProjectCardEventHandleFunc)
	}
	g.onProjectCardEvent[ProjectCardEventMovedAction] = append(
		g.onProjectCardEvent[ProjectCardEventMovedAction],
		callbacks...,
	)
}

// SetOnProjectCardEventMoved registers callbacks listening to events of type github.ProjectCardEvent and action 'moved'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnProjectCardEventMovedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#project_card
func (g *EventHandler) SetOnProjectCardEventMoved(callbacks ...ProjectCardEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectCardEvent == nil {
		g.onProjectCardEvent = make(map[string][]ProjectCardEventHandleFunc)
	}
	g.onProjectCardEvent[ProjectCardEventMovedAction] = callbacks
}

func (g *EventHandler) handleProjectCardEventMoved(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if ProjectCardEventMovedAction != *event.Action {
		return fmt.Errorf(
			"handleProjectCardEventMoved() called with wrong action, want %s, got %s",
			ProjectCardEventMovedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		ProjectCardEventMovedAction,
		ProjectCardEventAnyAction,
	} {
		if _, ok := g.onProjectCardEvent[action]; ok {
			for _, h := range g.onProjectCardEvent[action] {
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

// OnProjectCardEventDeleted registers callbacks listening to events of type github.ProjectCardEvent and action 'deleted'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnProjectCardEventDeleted must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#project_card
func (g *EventHandler) OnProjectCardEventDeleted(callbacks ...ProjectCardEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectCardEvent == nil {
		g.onProjectCardEvent = make(map[string][]ProjectCardEventHandleFunc)
	}
	g.onProjectCardEvent[ProjectCardEventDeletedAction] = append(
		g.onProjectCardEvent[ProjectCardEventDeletedAction],
		callbacks...,
	)
}

// SetOnProjectCardEventDeleted registers callbacks listening to events of type github.ProjectCardEvent and action 'deleted'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnProjectCardEventDeletedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#project_card
func (g *EventHandler) SetOnProjectCardEventDeleted(callbacks ...ProjectCardEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectCardEvent == nil {
		g.onProjectCardEvent = make(map[string][]ProjectCardEventHandleFunc)
	}
	g.onProjectCardEvent[ProjectCardEventDeletedAction] = callbacks
}

func (g *EventHandler) handleProjectCardEventDeleted(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if ProjectCardEventDeletedAction != *event.Action {
		return fmt.Errorf(
			"handleProjectCardEventDeleted() called with wrong action, want %s, got %s",
			ProjectCardEventDeletedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		ProjectCardEventDeletedAction,
		ProjectCardEventAnyAction,
	} {
		if _, ok := g.onProjectCardEvent[action]; ok {
			for _, h := range g.onProjectCardEvent[action] {
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

// OnProjectCardEventAny registers callbacks listening to any events of type github.ProjectCardEvent
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnProjectCardEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#project_card
func (g *EventHandler) OnProjectCardEventAny(callbacks ...ProjectCardEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectCardEvent == nil {
		g.onProjectCardEvent = make(map[string][]ProjectCardEventHandleFunc)
	}
	g.onProjectCardEvent[ProjectCardEventAnyAction] = append(
		g.onProjectCardEvent[ProjectCardEventAnyAction],
		callbacks...,
	)
}

// SetOnProjectCardEventAny registers callbacks listening to any events of type github.ProjectCardEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnProjectCardEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#project_card
func (g *EventHandler) SetOnProjectCardEventAny(callbacks ...ProjectCardEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectCardEvent == nil {
		g.onProjectCardEvent = make(map[string][]ProjectCardEventHandleFunc)
	}
	g.onProjectCardEvent[ProjectCardEventAnyAction] = callbacks
}

func (g *EventHandler) handleProjectCardEventAny(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
	if event == nil {
		return fmt.Errorf("event was empty or nil")
	}
	if _, ok := g.onProjectCardEvent[ProjectCardEventAnyAction]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onProjectCardEvent[ProjectCardEventAnyAction] {
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

// ProjectCardEvent handles github.ProjectCardEvent.
//
// Callbacks are executed in the following order:
//
// 1) All callbacks registered with OnBeforeAny are executed in parallel.
// 2) All callbacks registered with OnProjectCardEvent... are executed in parallel in case the Event has actions.
// 3) All callbacks registered with OnAfterAny are executed in parallel.
//
// on any error all callbacks registered with OnError are executed in parallel.
func (g *EventHandler) ProjectCardEvent(deliveryID string, eventName string, event *github.ProjectCardEvent) error {

	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	action := *event.Action

	err := g.handleBeforeAny(deliveryID, eventName, event)
	if err != nil {
		return g.handleError(deliveryID, eventName, event, err)
	}

	switch action {

	case ProjectCardEventCreatedAction:
		err := g.handleProjectCardEventCreated(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case ProjectCardEventEditedAction:
		err := g.handleProjectCardEventEdited(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case ProjectCardEventConvertedAction:
		err := g.handleProjectCardEventConverted(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case ProjectCardEventMovedAction:
		err := g.handleProjectCardEventMoved(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case ProjectCardEventDeletedAction:
		err := g.handleProjectCardEventDeleted(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	default:
		err := g.handleProjectCardEventAny(deliveryID, eventName, event)
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
