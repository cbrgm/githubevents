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
	// ProjectColumnEvent is the event name of github.ProjectColumnEvent's
	ProjectColumnEvent = "project_column"

	// ProjectColumnEventAnyAction is used to identify callbacks
	// listening to all events of type github.ProjectColumnEvent
	ProjectColumnEventAnyAction = "*"

	// ProjectColumnEventCreatedAction is used to identify callbacks
	// listening to events of type github.ProjectColumnEvent and action "created"
	ProjectColumnEventCreatedAction = "created"

	// ProjectColumnEventEditedAction is used to identify callbacks
	// listening to events of type github.ProjectColumnEvent and action "edited"
	ProjectColumnEventEditedAction = "edited"

	// ProjectColumnEventMovedAction is used to identify callbacks
	// listening to events of type github.ProjectColumnEvent and action "moved"
	ProjectColumnEventMovedAction = "moved"

	// ProjectColumnEventDeletedAction is used to identify callbacks
	// listening to events of type github.ProjectColumnEvent and action "deleted"
	ProjectColumnEventDeletedAction = "deleted"
)

// ProjectColumnEventHandleFunc represents a callback function triggered on github.ProjectColumnEvent's.
// 'deliveryID' (type: string) is the unique webhook delivery ID.
// 'eventName' (type: string) is the name of the event.
// 'event' (type: *github.ProjectColumnEvent) is the webhook payload.
type ProjectColumnEventHandleFunc func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error

// OnProjectColumnEventCreated registers callbacks listening to events of type github.ProjectColumnEvent and action 'created'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnProjectColumnEventCreated must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#project_column
func (g *EventHandler) OnProjectColumnEventCreated(callbacks ...ProjectColumnEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectColumnEvent == nil {
		g.onProjectColumnEvent = make(map[string][]ProjectColumnEventHandleFunc)
	}
	g.onProjectColumnEvent[ProjectColumnEventCreatedAction] = append(
		g.onProjectColumnEvent[ProjectColumnEventCreatedAction],
		callbacks...,
	)
}

// SetOnProjectColumnEventCreated registers callbacks listening to events of type github.ProjectColumnEvent and action 'created'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnProjectColumnEventCreatedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#project_column
func (g *EventHandler) SetOnProjectColumnEventCreated(callbacks ...ProjectColumnEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectColumnEvent == nil {
		g.onProjectColumnEvent = make(map[string][]ProjectColumnEventHandleFunc)
	}
	g.onProjectColumnEvent[ProjectColumnEventCreatedAction] = callbacks
}

func (g *EventHandler) handleProjectColumnEventCreated(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if ProjectColumnEventCreatedAction != *event.Action {
		return fmt.Errorf(
			"handleProjectColumnEventCreated() called with wrong action, want %s, got %s",
			ProjectColumnEventCreatedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		ProjectColumnEventCreatedAction,
		ProjectColumnEventAnyAction,
	} {
		if _, ok := g.onProjectColumnEvent[action]; ok {
			for _, h := range g.onProjectColumnEvent[action] {
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

// OnProjectColumnEventEdited registers callbacks listening to events of type github.ProjectColumnEvent and action 'edited'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnProjectColumnEventEdited must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#project_column
func (g *EventHandler) OnProjectColumnEventEdited(callbacks ...ProjectColumnEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectColumnEvent == nil {
		g.onProjectColumnEvent = make(map[string][]ProjectColumnEventHandleFunc)
	}
	g.onProjectColumnEvent[ProjectColumnEventEditedAction] = append(
		g.onProjectColumnEvent[ProjectColumnEventEditedAction],
		callbacks...,
	)
}

// SetOnProjectColumnEventEdited registers callbacks listening to events of type github.ProjectColumnEvent and action 'edited'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnProjectColumnEventEditedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#project_column
func (g *EventHandler) SetOnProjectColumnEventEdited(callbacks ...ProjectColumnEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectColumnEvent == nil {
		g.onProjectColumnEvent = make(map[string][]ProjectColumnEventHandleFunc)
	}
	g.onProjectColumnEvent[ProjectColumnEventEditedAction] = callbacks
}

func (g *EventHandler) handleProjectColumnEventEdited(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if ProjectColumnEventEditedAction != *event.Action {
		return fmt.Errorf(
			"handleProjectColumnEventEdited() called with wrong action, want %s, got %s",
			ProjectColumnEventEditedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		ProjectColumnEventEditedAction,
		ProjectColumnEventAnyAction,
	} {
		if _, ok := g.onProjectColumnEvent[action]; ok {
			for _, h := range g.onProjectColumnEvent[action] {
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

// OnProjectColumnEventMoved registers callbacks listening to events of type github.ProjectColumnEvent and action 'moved'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnProjectColumnEventMoved must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#project_column
func (g *EventHandler) OnProjectColumnEventMoved(callbacks ...ProjectColumnEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectColumnEvent == nil {
		g.onProjectColumnEvent = make(map[string][]ProjectColumnEventHandleFunc)
	}
	g.onProjectColumnEvent[ProjectColumnEventMovedAction] = append(
		g.onProjectColumnEvent[ProjectColumnEventMovedAction],
		callbacks...,
	)
}

// SetOnProjectColumnEventMoved registers callbacks listening to events of type github.ProjectColumnEvent and action 'moved'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnProjectColumnEventMovedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#project_column
func (g *EventHandler) SetOnProjectColumnEventMoved(callbacks ...ProjectColumnEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectColumnEvent == nil {
		g.onProjectColumnEvent = make(map[string][]ProjectColumnEventHandleFunc)
	}
	g.onProjectColumnEvent[ProjectColumnEventMovedAction] = callbacks
}

func (g *EventHandler) handleProjectColumnEventMoved(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if ProjectColumnEventMovedAction != *event.Action {
		return fmt.Errorf(
			"handleProjectColumnEventMoved() called with wrong action, want %s, got %s",
			ProjectColumnEventMovedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		ProjectColumnEventMovedAction,
		ProjectColumnEventAnyAction,
	} {
		if _, ok := g.onProjectColumnEvent[action]; ok {
			for _, h := range g.onProjectColumnEvent[action] {
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

// OnProjectColumnEventDeleted registers callbacks listening to events of type github.ProjectColumnEvent and action 'deleted'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnProjectColumnEventDeleted must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#project_column
func (g *EventHandler) OnProjectColumnEventDeleted(callbacks ...ProjectColumnEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectColumnEvent == nil {
		g.onProjectColumnEvent = make(map[string][]ProjectColumnEventHandleFunc)
	}
	g.onProjectColumnEvent[ProjectColumnEventDeletedAction] = append(
		g.onProjectColumnEvent[ProjectColumnEventDeletedAction],
		callbacks...,
	)
}

// SetOnProjectColumnEventDeleted registers callbacks listening to events of type github.ProjectColumnEvent and action 'deleted'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnProjectColumnEventDeletedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#project_column
func (g *EventHandler) SetOnProjectColumnEventDeleted(callbacks ...ProjectColumnEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectColumnEvent == nil {
		g.onProjectColumnEvent = make(map[string][]ProjectColumnEventHandleFunc)
	}
	g.onProjectColumnEvent[ProjectColumnEventDeletedAction] = callbacks
}

func (g *EventHandler) handleProjectColumnEventDeleted(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if ProjectColumnEventDeletedAction != *event.Action {
		return fmt.Errorf(
			"handleProjectColumnEventDeleted() called with wrong action, want %s, got %s",
			ProjectColumnEventDeletedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		ProjectColumnEventDeletedAction,
		ProjectColumnEventAnyAction,
	} {
		if _, ok := g.onProjectColumnEvent[action]; ok {
			for _, h := range g.onProjectColumnEvent[action] {
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

// OnProjectColumnEventAny registers callbacks listening to any events of type github.ProjectColumnEvent
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnProjectColumnEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#project_column
func (g *EventHandler) OnProjectColumnEventAny(callbacks ...ProjectColumnEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectColumnEvent == nil {
		g.onProjectColumnEvent = make(map[string][]ProjectColumnEventHandleFunc)
	}
	g.onProjectColumnEvent[ProjectColumnEventAnyAction] = append(
		g.onProjectColumnEvent[ProjectColumnEventAnyAction],
		callbacks...,
	)
}

// SetOnProjectColumnEventAny registers callbacks listening to any events of type github.ProjectColumnEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnProjectColumnEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#project_column
func (g *EventHandler) SetOnProjectColumnEventAny(callbacks ...ProjectColumnEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectColumnEvent == nil {
		g.onProjectColumnEvent = make(map[string][]ProjectColumnEventHandleFunc)
	}
	g.onProjectColumnEvent[ProjectColumnEventAnyAction] = callbacks
}

func (g *EventHandler) handleProjectColumnEventAny(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
	if event == nil {
		return fmt.Errorf("event was empty or nil")
	}
	if _, ok := g.onProjectColumnEvent[ProjectColumnEventAnyAction]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onProjectColumnEvent[ProjectColumnEventAnyAction] {
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

// ProjectColumnEvent handles github.ProjectColumnEvent.
//
// Callbacks are executed in the following order:
//
// 1) All callbacks registered with OnBeforeAny are executed in parallel.
// 2) All callbacks registered with OnProjectColumnEvent... are executed in parallel in case the Event has actions.
// 3) All callbacks registered with OnAfterAny are executed in parallel.
//
// on any error all callbacks registered with OnError are executed in parallel.
func (g *EventHandler) ProjectColumnEvent(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {

	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	action := *event.Action

	err := g.handleBeforeAny(deliveryID, eventName, event)
	if err != nil {
		return g.handleError(deliveryID, eventName, event, err)
	}

	switch action {

	case ProjectColumnEventCreatedAction:
		err := g.handleProjectColumnEventCreated(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case ProjectColumnEventEditedAction:
		err := g.handleProjectColumnEventEdited(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case ProjectColumnEventMovedAction:
		err := g.handleProjectColumnEventMoved(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case ProjectColumnEventDeletedAction:
		err := g.handleProjectColumnEventDeleted(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	default:
		err := g.handleProjectColumnEventAny(deliveryID, eventName, event)
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
