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
	// ProjectV2ItemEvent is the event name of github.ProjectV2ItemEvent's
	ProjectV2ItemEvent = "project_v2_item"

	// ProjectV2ItemEventAnyAction is used to identify callbacks
	// listening to all events of type github.ProjectV2ItemEvent
	ProjectV2ItemEventAnyAction = "*"

	// ProjectItemEventCreatedAction is used to identify callbacks
	// listening to events of type github.ProjectV2ItemEvent and action "created"
	ProjectItemEventCreatedAction = "created"

	// ProjectItemEventEditedAction is used to identify callbacks
	// listening to events of type github.ProjectV2ItemEvent and action "edited"
	ProjectItemEventEditedAction = "edited"

	// ProjectItemEventClosedAction is used to identify callbacks
	// listening to events of type github.ProjectV2ItemEvent and action "closed"
	ProjectItemEventClosedAction = "closed"

	// ProjectItemEventReopenedAction is used to identify callbacks
	// listening to events of type github.ProjectV2ItemEvent and action "reopened"
	ProjectItemEventReopenedAction = "reopened"

	// ProjectItemEventDeletedAction is used to identify callbacks
	// listening to events of type github.ProjectV2ItemEvent and action "deleted"
	ProjectItemEventDeletedAction = "deleted"

	// ProjectItemEventConvertedAction is used to identify callbacks
	// listening to events of type github.ProjectV2ItemEvent and action "converted"
	ProjectItemEventConvertedAction = "converted"

	// ProjectItemEventRestoredAction is used to identify callbacks
	// listening to events of type github.ProjectV2ItemEvent and action "restored"
	ProjectItemEventRestoredAction = "restored"
)

// ProjectV2ItemEventHandleFunc represents a callback function triggered on github.ProjectV2ItemEvent's.
// 'deliveryID' (type: string) is the unique webhook delivery ID.
// 'eventName' (type: string) is the name of the event.
// 'event' (type: *github.ProjectV2ItemEvent) is the webhook payload.
type ProjectV2ItemEventHandleFunc func(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error

// OnProjectItemEventCreated registers callbacks listening to events of type github.ProjectV2ItemEvent and action 'created'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnProjectItemEventCreated must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#project_v2_item
func (g *EventHandler) OnProjectItemEventCreated(callbacks ...ProjectV2ItemEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectV2ItemEvent == nil {
		g.onProjectV2ItemEvent = make(map[string][]ProjectV2ItemEventHandleFunc)
	}
	g.onProjectV2ItemEvent[ProjectItemEventCreatedAction] = append(
		g.onProjectV2ItemEvent[ProjectItemEventCreatedAction],
		callbacks...,
	)
}

// SetOnProjectItemEventCreated registers callbacks listening to events of type github.ProjectV2ItemEvent and action 'created'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnProjectItemEventCreatedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#project_v2_item
func (g *EventHandler) SetOnProjectItemEventCreated(callbacks ...ProjectV2ItemEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectV2ItemEvent == nil {
		g.onProjectV2ItemEvent = make(map[string][]ProjectV2ItemEventHandleFunc)
	}
	g.onProjectV2ItemEvent[ProjectItemEventCreatedAction] = callbacks
}

func (g *EventHandler) handleProjectItemEventCreated(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if ProjectItemEventCreatedAction != *event.Action {
		return fmt.Errorf(
			"handleProjectItemEventCreated() called with wrong action, want %s, got %s",
			ProjectItemEventCreatedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		ProjectItemEventCreatedAction,
		ProjectV2ItemEventAnyAction,
	} {
		if _, ok := g.onProjectV2ItemEvent[action]; ok {
			for _, h := range g.onProjectV2ItemEvent[action] {
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

// OnProjectItemEventEdited registers callbacks listening to events of type github.ProjectV2ItemEvent and action 'edited'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnProjectItemEventEdited must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#project_v2_item
func (g *EventHandler) OnProjectItemEventEdited(callbacks ...ProjectV2ItemEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectV2ItemEvent == nil {
		g.onProjectV2ItemEvent = make(map[string][]ProjectV2ItemEventHandleFunc)
	}
	g.onProjectV2ItemEvent[ProjectItemEventEditedAction] = append(
		g.onProjectV2ItemEvent[ProjectItemEventEditedAction],
		callbacks...,
	)
}

// SetOnProjectItemEventEdited registers callbacks listening to events of type github.ProjectV2ItemEvent and action 'edited'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnProjectItemEventEditedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#project_v2_item
func (g *EventHandler) SetOnProjectItemEventEdited(callbacks ...ProjectV2ItemEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectV2ItemEvent == nil {
		g.onProjectV2ItemEvent = make(map[string][]ProjectV2ItemEventHandleFunc)
	}
	g.onProjectV2ItemEvent[ProjectItemEventEditedAction] = callbacks
}

func (g *EventHandler) handleProjectItemEventEdited(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if ProjectItemEventEditedAction != *event.Action {
		return fmt.Errorf(
			"handleProjectItemEventEdited() called with wrong action, want %s, got %s",
			ProjectItemEventEditedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		ProjectItemEventEditedAction,
		ProjectV2ItemEventAnyAction,
	} {
		if _, ok := g.onProjectV2ItemEvent[action]; ok {
			for _, h := range g.onProjectV2ItemEvent[action] {
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

// OnProjectItemEventClosed registers callbacks listening to events of type github.ProjectV2ItemEvent and action 'closed'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnProjectItemEventClosed must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#project_v2_item
func (g *EventHandler) OnProjectItemEventClosed(callbacks ...ProjectV2ItemEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectV2ItemEvent == nil {
		g.onProjectV2ItemEvent = make(map[string][]ProjectV2ItemEventHandleFunc)
	}
	g.onProjectV2ItemEvent[ProjectItemEventClosedAction] = append(
		g.onProjectV2ItemEvent[ProjectItemEventClosedAction],
		callbacks...,
	)
}

// SetOnProjectItemEventClosed registers callbacks listening to events of type github.ProjectV2ItemEvent and action 'closed'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnProjectItemEventClosedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#project_v2_item
func (g *EventHandler) SetOnProjectItemEventClosed(callbacks ...ProjectV2ItemEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectV2ItemEvent == nil {
		g.onProjectV2ItemEvent = make(map[string][]ProjectV2ItemEventHandleFunc)
	}
	g.onProjectV2ItemEvent[ProjectItemEventClosedAction] = callbacks
}

func (g *EventHandler) handleProjectItemEventClosed(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if ProjectItemEventClosedAction != *event.Action {
		return fmt.Errorf(
			"handleProjectItemEventClosed() called with wrong action, want %s, got %s",
			ProjectItemEventClosedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		ProjectItemEventClosedAction,
		ProjectV2ItemEventAnyAction,
	} {
		if _, ok := g.onProjectV2ItemEvent[action]; ok {
			for _, h := range g.onProjectV2ItemEvent[action] {
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

// OnProjectItemEventReopened registers callbacks listening to events of type github.ProjectV2ItemEvent and action 'reopened'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnProjectItemEventReopened must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#project_v2_item
func (g *EventHandler) OnProjectItemEventReopened(callbacks ...ProjectV2ItemEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectV2ItemEvent == nil {
		g.onProjectV2ItemEvent = make(map[string][]ProjectV2ItemEventHandleFunc)
	}
	g.onProjectV2ItemEvent[ProjectItemEventReopenedAction] = append(
		g.onProjectV2ItemEvent[ProjectItemEventReopenedAction],
		callbacks...,
	)
}

// SetOnProjectItemEventReopened registers callbacks listening to events of type github.ProjectV2ItemEvent and action 'reopened'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnProjectItemEventReopenedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#project_v2_item
func (g *EventHandler) SetOnProjectItemEventReopened(callbacks ...ProjectV2ItemEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectV2ItemEvent == nil {
		g.onProjectV2ItemEvent = make(map[string][]ProjectV2ItemEventHandleFunc)
	}
	g.onProjectV2ItemEvent[ProjectItemEventReopenedAction] = callbacks
}

func (g *EventHandler) handleProjectItemEventReopened(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if ProjectItemEventReopenedAction != *event.Action {
		return fmt.Errorf(
			"handleProjectItemEventReopened() called with wrong action, want %s, got %s",
			ProjectItemEventReopenedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		ProjectItemEventReopenedAction,
		ProjectV2ItemEventAnyAction,
	} {
		if _, ok := g.onProjectV2ItemEvent[action]; ok {
			for _, h := range g.onProjectV2ItemEvent[action] {
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

// OnProjectItemEventDeleted registers callbacks listening to events of type github.ProjectV2ItemEvent and action 'deleted'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnProjectItemEventDeleted must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#project_v2_item
func (g *EventHandler) OnProjectItemEventDeleted(callbacks ...ProjectV2ItemEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectV2ItemEvent == nil {
		g.onProjectV2ItemEvent = make(map[string][]ProjectV2ItemEventHandleFunc)
	}
	g.onProjectV2ItemEvent[ProjectItemEventDeletedAction] = append(
		g.onProjectV2ItemEvent[ProjectItemEventDeletedAction],
		callbacks...,
	)
}

// SetOnProjectItemEventDeleted registers callbacks listening to events of type github.ProjectV2ItemEvent and action 'deleted'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnProjectItemEventDeletedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#project_v2_item
func (g *EventHandler) SetOnProjectItemEventDeleted(callbacks ...ProjectV2ItemEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectV2ItemEvent == nil {
		g.onProjectV2ItemEvent = make(map[string][]ProjectV2ItemEventHandleFunc)
	}
	g.onProjectV2ItemEvent[ProjectItemEventDeletedAction] = callbacks
}

func (g *EventHandler) handleProjectItemEventDeleted(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if ProjectItemEventDeletedAction != *event.Action {
		return fmt.Errorf(
			"handleProjectItemEventDeleted() called with wrong action, want %s, got %s",
			ProjectItemEventDeletedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		ProjectItemEventDeletedAction,
		ProjectV2ItemEventAnyAction,
	} {
		if _, ok := g.onProjectV2ItemEvent[action]; ok {
			for _, h := range g.onProjectV2ItemEvent[action] {
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

// OnProjectItemEventConverted registers callbacks listening to events of type github.ProjectV2ItemEvent and action 'converted'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnProjectItemEventConverted must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#project_v2_item
func (g *EventHandler) OnProjectItemEventConverted(callbacks ...ProjectV2ItemEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectV2ItemEvent == nil {
		g.onProjectV2ItemEvent = make(map[string][]ProjectV2ItemEventHandleFunc)
	}
	g.onProjectV2ItemEvent[ProjectItemEventConvertedAction] = append(
		g.onProjectV2ItemEvent[ProjectItemEventConvertedAction],
		callbacks...,
	)
}

// SetOnProjectItemEventConverted registers callbacks listening to events of type github.ProjectV2ItemEvent and action 'converted'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnProjectItemEventConvertedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#project_v2_item
func (g *EventHandler) SetOnProjectItemEventConverted(callbacks ...ProjectV2ItemEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectV2ItemEvent == nil {
		g.onProjectV2ItemEvent = make(map[string][]ProjectV2ItemEventHandleFunc)
	}
	g.onProjectV2ItemEvent[ProjectItemEventConvertedAction] = callbacks
}

func (g *EventHandler) handleProjectItemEventConverted(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if ProjectItemEventConvertedAction != *event.Action {
		return fmt.Errorf(
			"handleProjectItemEventConverted() called with wrong action, want %s, got %s",
			ProjectItemEventConvertedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		ProjectItemEventConvertedAction,
		ProjectV2ItemEventAnyAction,
	} {
		if _, ok := g.onProjectV2ItemEvent[action]; ok {
			for _, h := range g.onProjectV2ItemEvent[action] {
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

// OnProjectItemEventRestored registers callbacks listening to events of type github.ProjectV2ItemEvent and action 'restored'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnProjectItemEventRestored must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#project_v2_item
func (g *EventHandler) OnProjectItemEventRestored(callbacks ...ProjectV2ItemEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectV2ItemEvent == nil {
		g.onProjectV2ItemEvent = make(map[string][]ProjectV2ItemEventHandleFunc)
	}
	g.onProjectV2ItemEvent[ProjectItemEventRestoredAction] = append(
		g.onProjectV2ItemEvent[ProjectItemEventRestoredAction],
		callbacks...,
	)
}

// SetOnProjectItemEventRestored registers callbacks listening to events of type github.ProjectV2ItemEvent and action 'restored'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnProjectItemEventRestoredAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#project_v2_item
func (g *EventHandler) SetOnProjectItemEventRestored(callbacks ...ProjectV2ItemEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectV2ItemEvent == nil {
		g.onProjectV2ItemEvent = make(map[string][]ProjectV2ItemEventHandleFunc)
	}
	g.onProjectV2ItemEvent[ProjectItemEventRestoredAction] = callbacks
}

func (g *EventHandler) handleProjectItemEventRestored(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if ProjectItemEventRestoredAction != *event.Action {
		return fmt.Errorf(
			"handleProjectItemEventRestored() called with wrong action, want %s, got %s",
			ProjectItemEventRestoredAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		ProjectItemEventRestoredAction,
		ProjectV2ItemEventAnyAction,
	} {
		if _, ok := g.onProjectV2ItemEvent[action]; ok {
			for _, h := range g.onProjectV2ItemEvent[action] {
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

// OnProjectV2ItemEventAny registers callbacks listening to any events of type github.ProjectV2ItemEvent
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnProjectV2ItemEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#project_v2_item
func (g *EventHandler) OnProjectV2ItemEventAny(callbacks ...ProjectV2ItemEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectV2ItemEvent == nil {
		g.onProjectV2ItemEvent = make(map[string][]ProjectV2ItemEventHandleFunc)
	}
	g.onProjectV2ItemEvent[ProjectV2ItemEventAnyAction] = append(
		g.onProjectV2ItemEvent[ProjectV2ItemEventAnyAction],
		callbacks...,
	)
}

// SetOnProjectV2ItemEventAny registers callbacks listening to any events of type github.ProjectV2ItemEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnProjectV2ItemEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#project_v2_item
func (g *EventHandler) SetOnProjectV2ItemEventAny(callbacks ...ProjectV2ItemEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectV2ItemEvent == nil {
		g.onProjectV2ItemEvent = make(map[string][]ProjectV2ItemEventHandleFunc)
	}
	g.onProjectV2ItemEvent[ProjectV2ItemEventAnyAction] = callbacks
}

func (g *EventHandler) handleProjectV2ItemEventAny(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {
	if event == nil {
		return fmt.Errorf("event was empty or nil")
	}
	if _, ok := g.onProjectV2ItemEvent[ProjectV2ItemEventAnyAction]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onProjectV2ItemEvent[ProjectV2ItemEventAnyAction] {
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

// ProjectV2ItemEvent handles github.ProjectV2ItemEvent.
//
// Callbacks are executed in the following order:
//
// 1) All callbacks registered with OnBeforeAny are executed in parallel.
// 2) All callbacks registered with OnProjectV2ItemEvent... are executed in parallel in case the Event has actions.
// 3) All callbacks registered with OnAfterAny are executed in parallel.
//
// on any error all callbacks registered with OnError are executed in parallel.
func (g *EventHandler) ProjectV2ItemEvent(deliveryID string, eventName string, event *github.ProjectV2ItemEvent) error {

	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	action := *event.Action

	err := g.handleBeforeAny(deliveryID, eventName, event)
	if err != nil {
		return g.handleError(deliveryID, eventName, event, err)
	}

	switch action {

	case ProjectItemEventCreatedAction:
		err := g.handleProjectItemEventCreated(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case ProjectItemEventEditedAction:
		err := g.handleProjectItemEventEdited(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case ProjectItemEventClosedAction:
		err := g.handleProjectItemEventClosed(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case ProjectItemEventReopenedAction:
		err := g.handleProjectItemEventReopened(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case ProjectItemEventDeletedAction:
		err := g.handleProjectItemEventDeleted(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case ProjectItemEventConvertedAction:
		err := g.handleProjectItemEventConverted(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case ProjectItemEventRestoredAction:
		err := g.handleProjectItemEventRestored(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	default:
		err := g.handleProjectV2ItemEventAny(deliveryID, eventName, event)
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
