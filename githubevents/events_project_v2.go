// Copyright 2022 The GithubEvents Authors. All rights reserved.
// Use of this source code is governed by the MIT License
// that can be found in the LICENSE file.

package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"fmt"
	"github.com/google/go-github/v68/github"
	"golang.org/x/sync/errgroup"
)

// Actions are used to identify registered callbacks.
const (
	// ProjectV2Event is the event name of github.ProjectV2Event's
	ProjectV2Event = "project_v2"

	// ProjectV2EventAnyAction is used to identify callbacks
	// listening to all events of type github.ProjectV2Event
	ProjectV2EventAnyAction = "*"

	// ProjectEventCreatedAction is used to identify callbacks
	// listening to events of type github.ProjectV2Event and action "created"
	ProjectEventCreatedAction = "created"

	// ProjectEventEditedAction is used to identify callbacks
	// listening to events of type github.ProjectV2Event and action "edited"
	ProjectEventEditedAction = "edited"

	// ProjectEventClosedAction is used to identify callbacks
	// listening to events of type github.ProjectV2Event and action "closed"
	ProjectEventClosedAction = "closed"

	// ProjectEventReopenedAction is used to identify callbacks
	// listening to events of type github.ProjectV2Event and action "reopened"
	ProjectEventReopenedAction = "reopened"

	// ProjectEventDeletedAction is used to identify callbacks
	// listening to events of type github.ProjectV2Event and action "deleted"
	ProjectEventDeletedAction = "deleted"
)

// ProjectV2EventHandleFunc represents a callback function triggered on github.ProjectV2Event's.
// 'deliveryID' (type: string) is the unique webhook delivery ID.
// 'eventName' (type: string) is the name of the event.
// 'event' (type: *github.ProjectV2Event) is the webhook payload.
type ProjectV2EventHandleFunc func(deliveryID string, eventName string, event *github.ProjectV2Event) error

// OnProjectEventCreated registers callbacks listening to events of type github.ProjectV2Event and action 'created'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnProjectEventCreated must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#project_v2
func (g *EventHandler) OnProjectEventCreated(callbacks ...ProjectV2EventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectV2Event == nil {
		g.onProjectV2Event = make(map[string][]ProjectV2EventHandleFunc)
	}
	g.onProjectV2Event[ProjectEventCreatedAction] = append(
		g.onProjectV2Event[ProjectEventCreatedAction],
		callbacks...,
	)
}

// SetOnProjectEventCreated registers callbacks listening to events of type github.ProjectV2Event and action 'created'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnProjectEventCreatedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#project_v2
func (g *EventHandler) SetOnProjectEventCreated(callbacks ...ProjectV2EventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectV2Event == nil {
		g.onProjectV2Event = make(map[string][]ProjectV2EventHandleFunc)
	}
	g.onProjectV2Event[ProjectEventCreatedAction] = callbacks
}

func (g *EventHandler) handleProjectEventCreated(deliveryID string, eventName string, event *github.ProjectV2Event) error {
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
		ProjectV2EventAnyAction,
	} {
		if _, ok := g.onProjectV2Event[action]; ok {
			for _, h := range g.onProjectV2Event[action] {
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

// OnProjectEventEdited registers callbacks listening to events of type github.ProjectV2Event and action 'edited'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnProjectEventEdited must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#project_v2
func (g *EventHandler) OnProjectEventEdited(callbacks ...ProjectV2EventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectV2Event == nil {
		g.onProjectV2Event = make(map[string][]ProjectV2EventHandleFunc)
	}
	g.onProjectV2Event[ProjectEventEditedAction] = append(
		g.onProjectV2Event[ProjectEventEditedAction],
		callbacks...,
	)
}

// SetOnProjectEventEdited registers callbacks listening to events of type github.ProjectV2Event and action 'edited'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnProjectEventEditedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#project_v2
func (g *EventHandler) SetOnProjectEventEdited(callbacks ...ProjectV2EventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectV2Event == nil {
		g.onProjectV2Event = make(map[string][]ProjectV2EventHandleFunc)
	}
	g.onProjectV2Event[ProjectEventEditedAction] = callbacks
}

func (g *EventHandler) handleProjectEventEdited(deliveryID string, eventName string, event *github.ProjectV2Event) error {
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
		ProjectV2EventAnyAction,
	} {
		if _, ok := g.onProjectV2Event[action]; ok {
			for _, h := range g.onProjectV2Event[action] {
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

// OnProjectEventClosed registers callbacks listening to events of type github.ProjectV2Event and action 'closed'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnProjectEventClosed must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#project_v2
func (g *EventHandler) OnProjectEventClosed(callbacks ...ProjectV2EventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectV2Event == nil {
		g.onProjectV2Event = make(map[string][]ProjectV2EventHandleFunc)
	}
	g.onProjectV2Event[ProjectEventClosedAction] = append(
		g.onProjectV2Event[ProjectEventClosedAction],
		callbacks...,
	)
}

// SetOnProjectEventClosed registers callbacks listening to events of type github.ProjectV2Event and action 'closed'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnProjectEventClosedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#project_v2
func (g *EventHandler) SetOnProjectEventClosed(callbacks ...ProjectV2EventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectV2Event == nil {
		g.onProjectV2Event = make(map[string][]ProjectV2EventHandleFunc)
	}
	g.onProjectV2Event[ProjectEventClosedAction] = callbacks
}

func (g *EventHandler) handleProjectEventClosed(deliveryID string, eventName string, event *github.ProjectV2Event) error {
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
		ProjectV2EventAnyAction,
	} {
		if _, ok := g.onProjectV2Event[action]; ok {
			for _, h := range g.onProjectV2Event[action] {
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

// OnProjectEventReopened registers callbacks listening to events of type github.ProjectV2Event and action 'reopened'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnProjectEventReopened must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#project_v2
func (g *EventHandler) OnProjectEventReopened(callbacks ...ProjectV2EventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectV2Event == nil {
		g.onProjectV2Event = make(map[string][]ProjectV2EventHandleFunc)
	}
	g.onProjectV2Event[ProjectEventReopenedAction] = append(
		g.onProjectV2Event[ProjectEventReopenedAction],
		callbacks...,
	)
}

// SetOnProjectEventReopened registers callbacks listening to events of type github.ProjectV2Event and action 'reopened'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnProjectEventReopenedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#project_v2
func (g *EventHandler) SetOnProjectEventReopened(callbacks ...ProjectV2EventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectV2Event == nil {
		g.onProjectV2Event = make(map[string][]ProjectV2EventHandleFunc)
	}
	g.onProjectV2Event[ProjectEventReopenedAction] = callbacks
}

func (g *EventHandler) handleProjectEventReopened(deliveryID string, eventName string, event *github.ProjectV2Event) error {
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
		ProjectV2EventAnyAction,
	} {
		if _, ok := g.onProjectV2Event[action]; ok {
			for _, h := range g.onProjectV2Event[action] {
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

// OnProjectEventDeleted registers callbacks listening to events of type github.ProjectV2Event and action 'deleted'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnProjectEventDeleted must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#project_v2
func (g *EventHandler) OnProjectEventDeleted(callbacks ...ProjectV2EventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectV2Event == nil {
		g.onProjectV2Event = make(map[string][]ProjectV2EventHandleFunc)
	}
	g.onProjectV2Event[ProjectEventDeletedAction] = append(
		g.onProjectV2Event[ProjectEventDeletedAction],
		callbacks...,
	)
}

// SetOnProjectEventDeleted registers callbacks listening to events of type github.ProjectV2Event and action 'deleted'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnProjectEventDeletedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#project_v2
func (g *EventHandler) SetOnProjectEventDeleted(callbacks ...ProjectV2EventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectV2Event == nil {
		g.onProjectV2Event = make(map[string][]ProjectV2EventHandleFunc)
	}
	g.onProjectV2Event[ProjectEventDeletedAction] = callbacks
}

func (g *EventHandler) handleProjectEventDeleted(deliveryID string, eventName string, event *github.ProjectV2Event) error {
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
		ProjectV2EventAnyAction,
	} {
		if _, ok := g.onProjectV2Event[action]; ok {
			for _, h := range g.onProjectV2Event[action] {
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

// OnProjectV2EventAny registers callbacks listening to any events of type github.ProjectV2Event
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnProjectV2EventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#project_v2
func (g *EventHandler) OnProjectV2EventAny(callbacks ...ProjectV2EventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectV2Event == nil {
		g.onProjectV2Event = make(map[string][]ProjectV2EventHandleFunc)
	}
	g.onProjectV2Event[ProjectV2EventAnyAction] = append(
		g.onProjectV2Event[ProjectV2EventAnyAction],
		callbacks...,
	)
}

// SetOnProjectV2EventAny registers callbacks listening to any events of type github.ProjectV2Event
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnProjectV2EventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#project_v2
func (g *EventHandler) SetOnProjectV2EventAny(callbacks ...ProjectV2EventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectV2Event == nil {
		g.onProjectV2Event = make(map[string][]ProjectV2EventHandleFunc)
	}
	g.onProjectV2Event[ProjectV2EventAnyAction] = callbacks
}

func (g *EventHandler) handleProjectV2EventAny(deliveryID string, eventName string, event *github.ProjectV2Event) error {
	if event == nil {
		return fmt.Errorf("event was empty or nil")
	}
	if _, ok := g.onProjectV2Event[ProjectV2EventAnyAction]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onProjectV2Event[ProjectV2EventAnyAction] {
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

// ProjectV2Event handles github.ProjectV2Event.
//
// Callbacks are executed in the following order:
//
// 1) All callbacks registered with OnBeforeAny are executed in parallel.
// 2) All callbacks registered with OnProjectV2Event... are executed in parallel in case the Event has actions.
// 3) All callbacks registered with OnAfterAny are executed in parallel.
//
// on any error all callbacks registered with OnError are executed in parallel.
func (g *EventHandler) ProjectV2Event(deliveryID string, eventName string, event *github.ProjectV2Event) error {

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
		err := g.handleProjectV2EventAny(deliveryID, eventName, event)
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