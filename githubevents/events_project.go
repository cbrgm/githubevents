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
	// ProjectEvent is the event name of github.ProjectEvent's
	ProjectEvent = "project"

	// ProjectEventAnyAction is used to identify callbacks
	// listening to all events of type github.ProjectEvent
	ProjectEventAnyAction = "*"

	// ProjectEventCreatedAction is used to identify callbacks
	// listening to events of type github.ProjectEvent and action "created"
	ProjectEventCreatedAction = "created"

	// ProjectEventEditedAction is used to identify callbacks
	// listening to events of type github.ProjectEvent and action "edited"
	ProjectEventEditedAction = "edited"

	// ProjectEventClosedAction is used to identify callbacks
	// listening to events of type github.ProjectEvent and action "closed"
	ProjectEventClosedAction = "closed"

	// ProjectEventReopenedAction is used to identify callbacks
	// listening to events of type github.ProjectEvent and action "reopened"
	ProjectEventReopenedAction = "reopened"

	// ProjectEventDeletedAction is used to identify callbacks
	// listening to events of type github.ProjectEvent and action "deleted"
	ProjectEventDeletedAction = "deleted"
)

// ProjectEventHandleFunc represents a callback function triggered on github.ProjectEvent's.
// 'deliveryID' (type: string) is the unique webhook delivery ID.
// 'eventName' (type: string) is the name of the event.
// 'event' (type: *github.ProjectEvent) is the webhook payload.
type ProjectEventHandleFunc func(deliveryID string, eventName string, event *github.ProjectEvent) error

// OnProjectEventCreated registers callbacks listening to events of type github.ProjectEvent and action 'created'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnProjectEventCreated must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#project
func (g *EventHandler) OnProjectEventCreated(callbacks ...ProjectEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectEvent == nil {
		g.onProjectEvent = make(map[string][]ProjectEventHandleFunc)
	}
	g.onProjectEvent[ProjectEventCreatedAction] = append(
		g.onProjectEvent[ProjectEventCreatedAction],
		callbacks...,
	)
}

// SetOnProjectEventCreated registers callbacks listening to events of type github.ProjectEvent and action 'created'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnProjectEventCreatedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#project
func (g *EventHandler) SetOnProjectEventCreated(callbacks ...ProjectEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectEvent == nil {
		g.onProjectEvent = make(map[string][]ProjectEventHandleFunc)
	}
	g.onProjectEvent[ProjectEventCreatedAction] = callbacks
}

func (g *EventHandler) handleProjectEventCreated(deliveryID string, eventName string, event *github.ProjectEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if ProjectEventCreatedAction != *event.Action {
		return fmt.Errorf(
			"handleProjectEventCreated() called with wrong action, want %s, got %s",
			ProjectEventCreatedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		ProjectEventCreatedAction,
		ProjectEventAnyAction,
	} {
		if _, ok := g.onProjectEvent[action]; ok {
			for _, h := range g.onProjectEvent[action] {
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

// OnProjectEventEdited registers callbacks listening to events of type github.ProjectEvent and action 'edited'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnProjectEventEdited must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#project
func (g *EventHandler) OnProjectEventEdited(callbacks ...ProjectEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectEvent == nil {
		g.onProjectEvent = make(map[string][]ProjectEventHandleFunc)
	}
	g.onProjectEvent[ProjectEventEditedAction] = append(
		g.onProjectEvent[ProjectEventEditedAction],
		callbacks...,
	)
}

// SetOnProjectEventEdited registers callbacks listening to events of type github.ProjectEvent and action 'edited'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnProjectEventEditedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#project
func (g *EventHandler) SetOnProjectEventEdited(callbacks ...ProjectEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectEvent == nil {
		g.onProjectEvent = make(map[string][]ProjectEventHandleFunc)
	}
	g.onProjectEvent[ProjectEventEditedAction] = callbacks
}

func (g *EventHandler) handleProjectEventEdited(deliveryID string, eventName string, event *github.ProjectEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if ProjectEventEditedAction != *event.Action {
		return fmt.Errorf(
			"handleProjectEventEdited() called with wrong action, want %s, got %s",
			ProjectEventEditedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		ProjectEventEditedAction,
		ProjectEventAnyAction,
	} {
		if _, ok := g.onProjectEvent[action]; ok {
			for _, h := range g.onProjectEvent[action] {
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

// OnProjectEventClosed registers callbacks listening to events of type github.ProjectEvent and action 'closed'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnProjectEventClosed must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#project
func (g *EventHandler) OnProjectEventClosed(callbacks ...ProjectEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectEvent == nil {
		g.onProjectEvent = make(map[string][]ProjectEventHandleFunc)
	}
	g.onProjectEvent[ProjectEventClosedAction] = append(
		g.onProjectEvent[ProjectEventClosedAction],
		callbacks...,
	)
}

// SetOnProjectEventClosed registers callbacks listening to events of type github.ProjectEvent and action 'closed'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnProjectEventClosedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#project
func (g *EventHandler) SetOnProjectEventClosed(callbacks ...ProjectEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectEvent == nil {
		g.onProjectEvent = make(map[string][]ProjectEventHandleFunc)
	}
	g.onProjectEvent[ProjectEventClosedAction] = callbacks
}

func (g *EventHandler) handleProjectEventClosed(deliveryID string, eventName string, event *github.ProjectEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if ProjectEventClosedAction != *event.Action {
		return fmt.Errorf(
			"handleProjectEventClosed() called with wrong action, want %s, got %s",
			ProjectEventClosedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		ProjectEventClosedAction,
		ProjectEventAnyAction,
	} {
		if _, ok := g.onProjectEvent[action]; ok {
			for _, h := range g.onProjectEvent[action] {
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

// OnProjectEventReopened registers callbacks listening to events of type github.ProjectEvent and action 'reopened'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnProjectEventReopened must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#project
func (g *EventHandler) OnProjectEventReopened(callbacks ...ProjectEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectEvent == nil {
		g.onProjectEvent = make(map[string][]ProjectEventHandleFunc)
	}
	g.onProjectEvent[ProjectEventReopenedAction] = append(
		g.onProjectEvent[ProjectEventReopenedAction],
		callbacks...,
	)
}

// SetOnProjectEventReopened registers callbacks listening to events of type github.ProjectEvent and action 'reopened'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnProjectEventReopenedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#project
func (g *EventHandler) SetOnProjectEventReopened(callbacks ...ProjectEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectEvent == nil {
		g.onProjectEvent = make(map[string][]ProjectEventHandleFunc)
	}
	g.onProjectEvent[ProjectEventReopenedAction] = callbacks
}

func (g *EventHandler) handleProjectEventReopened(deliveryID string, eventName string, event *github.ProjectEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if ProjectEventReopenedAction != *event.Action {
		return fmt.Errorf(
			"handleProjectEventReopened() called with wrong action, want %s, got %s",
			ProjectEventReopenedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		ProjectEventReopenedAction,
		ProjectEventAnyAction,
	} {
		if _, ok := g.onProjectEvent[action]; ok {
			for _, h := range g.onProjectEvent[action] {
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

// OnProjectEventDeleted registers callbacks listening to events of type github.ProjectEvent and action 'deleted'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnProjectEventDeleted must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#project
func (g *EventHandler) OnProjectEventDeleted(callbacks ...ProjectEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectEvent == nil {
		g.onProjectEvent = make(map[string][]ProjectEventHandleFunc)
	}
	g.onProjectEvent[ProjectEventDeletedAction] = append(
		g.onProjectEvent[ProjectEventDeletedAction],
		callbacks...,
	)
}

// SetOnProjectEventDeleted registers callbacks listening to events of type github.ProjectEvent and action 'deleted'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnProjectEventDeletedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#project
func (g *EventHandler) SetOnProjectEventDeleted(callbacks ...ProjectEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectEvent == nil {
		g.onProjectEvent = make(map[string][]ProjectEventHandleFunc)
	}
	g.onProjectEvent[ProjectEventDeletedAction] = callbacks
}

func (g *EventHandler) handleProjectEventDeleted(deliveryID string, eventName string, event *github.ProjectEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if ProjectEventDeletedAction != *event.Action {
		return fmt.Errorf(
			"handleProjectEventDeleted() called with wrong action, want %s, got %s",
			ProjectEventDeletedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		ProjectEventDeletedAction,
		ProjectEventAnyAction,
	} {
		if _, ok := g.onProjectEvent[action]; ok {
			for _, h := range g.onProjectEvent[action] {
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

// OnProjectEventAny registers callbacks listening to any events of type github.ProjectEvent
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnProjectEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#project
func (g *EventHandler) OnProjectEventAny(callbacks ...ProjectEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectEvent == nil {
		g.onProjectEvent = make(map[string][]ProjectEventHandleFunc)
	}
	g.onProjectEvent[ProjectEventAnyAction] = append(
		g.onProjectEvent[ProjectEventAnyAction],
		callbacks...,
	)
}

// SetOnProjectEventAny registers callbacks listening to any events of type github.ProjectEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnProjectEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#project
func (g *EventHandler) SetOnProjectEventAny(callbacks ...ProjectEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectEvent == nil {
		g.onProjectEvent = make(map[string][]ProjectEventHandleFunc)
	}
	g.onProjectEvent[ProjectEventAnyAction] = callbacks
}

func (g *EventHandler) handleProjectEventAny(deliveryID string, eventName string, event *github.ProjectEvent) error {
	if event == nil {
		return fmt.Errorf("event was empty or nil")
	}
	if _, ok := g.onProjectEvent[ProjectEventAnyAction]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onProjectEvent[ProjectEventAnyAction] {
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

// ProjectEvent handles github.ProjectEvent.
//
// Callbacks are executed in the following order:
//
// 1) All callbacks registered with OnBeforeAny are executed in parallel.
// 2) All callbacks registered with OnProjectEvent... are executed in parallel in case the Event has actions.
// 3) All callbacks registered with OnAfterAny are executed in parallel.
//
// on any error all callbacks registered with OnError are executed in parallel.
func (g *EventHandler) ProjectEvent(deliveryID string, eventName string, event *github.ProjectEvent) error {

	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	action := *event.Action

	err := g.handleBeforeAny(deliveryID, eventName, event)
	if err != nil {
		return g.handleError(deliveryID, eventName, event, err)
	}

	switch action {

	case ProjectEventCreatedAction:
		err := g.handleProjectEventCreated(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case ProjectEventEditedAction:
		err := g.handleProjectEventEdited(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case ProjectEventClosedAction:
		err := g.handleProjectEventClosed(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case ProjectEventReopenedAction:
		err := g.handleProjectEventReopened(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case ProjectEventDeletedAction:
		err := g.handleProjectEventDeleted(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	default:
		err := g.handleProjectEventAny(deliveryID, eventName, event)
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
