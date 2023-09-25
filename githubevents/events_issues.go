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
	// IssuesEvent is the event name of github.IssuesEvent's
	IssuesEvent = "issues"

	// IssuesEventAnyAction is used to identify callbacks
	// listening to all events of type github.IssuesEvent
	IssuesEventAnyAction = "*"

	// IssuesEventOpenedAction is used to identify callbacks
	// listening to events of type github.IssuesEvent and action "opened"
	IssuesEventOpenedAction = "opened"

	// IssuesEventEditedAction is used to identify callbacks
	// listening to events of type github.IssuesEvent and action "edited"
	IssuesEventEditedAction = "edited"

	// IssuesEventDeletedAction is used to identify callbacks
	// listening to events of type github.IssuesEvent and action "deleted"
	IssuesEventDeletedAction = "deleted"

	// IssuesEventPinnedAction is used to identify callbacks
	// listening to events of type github.IssuesEvent and action "pinned"
	IssuesEventPinnedAction = "pinned"

	// IssuesEventUnpinnedAction is used to identify callbacks
	// listening to events of type github.IssuesEvent and action "unpinned"
	IssuesEventUnpinnedAction = "unpinned"

	// IssuesEventClosedAction is used to identify callbacks
	// listening to events of type github.IssuesEvent and action "closed"
	IssuesEventClosedAction = "closed"

	// IssuesEventReopenedAction is used to identify callbacks
	// listening to events of type github.IssuesEvent and action "reopened"
	IssuesEventReopenedAction = "reopened"

	// IssuesEventAssignedAction is used to identify callbacks
	// listening to events of type github.IssuesEvent and action "assigned"
	IssuesEventAssignedAction = "assigned"

	// IssuesEventUnassignedAction is used to identify callbacks
	// listening to events of type github.IssuesEvent and action "unassigned"
	IssuesEventUnassignedAction = "unassigned"

	// IssuesEventLabeledAction is used to identify callbacks
	// listening to events of type github.IssuesEvent and action "labeled"
	IssuesEventLabeledAction = "labeled"

	// IssuesEventUnlabeledAction is used to identify callbacks
	// listening to events of type github.IssuesEvent and action "unlabeled"
	IssuesEventUnlabeledAction = "unlabeled"

	// IssuesEventLockedAction is used to identify callbacks
	// listening to events of type github.IssuesEvent and action "locked"
	IssuesEventLockedAction = "locked"

	// IssuesEventUnlockedAction is used to identify callbacks
	// listening to events of type github.IssuesEvent and action "unlocked"
	IssuesEventUnlockedAction = "unlocked"

	// IssuesEventTransferredAction is used to identify callbacks
	// listening to events of type github.IssuesEvent and action "transferred"
	IssuesEventTransferredAction = "transferred"

	// IssuesEventMilestonedAction is used to identify callbacks
	// listening to events of type github.IssuesEvent and action "milestoned"
	IssuesEventMilestonedAction = "milestoned"

	// IssuesEventDeMilestonedAction is used to identify callbacks
	// listening to events of type github.IssuesEvent and action "demilestoned"
	IssuesEventDeMilestonedAction = "demilestoned"
)

// IssuesEventHandleFunc represents a callback function triggered on github.IssuesEvent's.
// 'deliveryID' (type: string) is the unique webhook delivery ID.
// 'eventName' (type: string) is the name of the event.
// 'event' (type: *github.IssuesEvent) is the webhook payload.
type IssuesEventHandleFunc func(deliveryID string, eventName string, event *github.IssuesEvent) error

// OnIssuesEventOpened registers callbacks listening to events of type github.IssuesEvent and action 'opened'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnIssuesEventOpened must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#issues
func (g *EventHandler) OnIssuesEventOpened(callbacks ...IssuesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssuesEvent == nil {
		g.onIssuesEvent = make(map[string][]IssuesEventHandleFunc)
	}
	g.onIssuesEvent[IssuesEventOpenedAction] = append(
		g.onIssuesEvent[IssuesEventOpenedAction],
		callbacks...,
	)
}

// SetOnIssuesEventOpened registers callbacks listening to events of type github.IssuesEvent and action 'opened'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnIssuesEventOpenedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#issues
func (g *EventHandler) SetOnIssuesEventOpened(callbacks ...IssuesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssuesEvent == nil {
		g.onIssuesEvent = make(map[string][]IssuesEventHandleFunc)
	}
	g.onIssuesEvent[IssuesEventOpenedAction] = callbacks
}

func (g *EventHandler) handleIssuesEventOpened(deliveryID string, eventName string, event *github.IssuesEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if IssuesEventOpenedAction != *event.Action {
		return fmt.Errorf(
			"handleIssuesEventOpened() called with wrong action, want %s, got %s",
			IssuesEventOpenedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		IssuesEventOpenedAction,
		IssuesEventAnyAction,
	} {
		if _, ok := g.onIssuesEvent[action]; ok {
			for _, h := range g.onIssuesEvent[action] {
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

// OnIssuesEventEdited registers callbacks listening to events of type github.IssuesEvent and action 'edited'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnIssuesEventEdited must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#issues
func (g *EventHandler) OnIssuesEventEdited(callbacks ...IssuesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssuesEvent == nil {
		g.onIssuesEvent = make(map[string][]IssuesEventHandleFunc)
	}
	g.onIssuesEvent[IssuesEventEditedAction] = append(
		g.onIssuesEvent[IssuesEventEditedAction],
		callbacks...,
	)
}

// SetOnIssuesEventEdited registers callbacks listening to events of type github.IssuesEvent and action 'edited'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnIssuesEventEditedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#issues
func (g *EventHandler) SetOnIssuesEventEdited(callbacks ...IssuesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssuesEvent == nil {
		g.onIssuesEvent = make(map[string][]IssuesEventHandleFunc)
	}
	g.onIssuesEvent[IssuesEventEditedAction] = callbacks
}

func (g *EventHandler) handleIssuesEventEdited(deliveryID string, eventName string, event *github.IssuesEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if IssuesEventEditedAction != *event.Action {
		return fmt.Errorf(
			"handleIssuesEventEdited() called with wrong action, want %s, got %s",
			IssuesEventEditedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		IssuesEventEditedAction,
		IssuesEventAnyAction,
	} {
		if _, ok := g.onIssuesEvent[action]; ok {
			for _, h := range g.onIssuesEvent[action] {
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

// OnIssuesEventDeleted registers callbacks listening to events of type github.IssuesEvent and action 'deleted'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnIssuesEventDeleted must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#issues
func (g *EventHandler) OnIssuesEventDeleted(callbacks ...IssuesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssuesEvent == nil {
		g.onIssuesEvent = make(map[string][]IssuesEventHandleFunc)
	}
	g.onIssuesEvent[IssuesEventDeletedAction] = append(
		g.onIssuesEvent[IssuesEventDeletedAction],
		callbacks...,
	)
}

// SetOnIssuesEventDeleted registers callbacks listening to events of type github.IssuesEvent and action 'deleted'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnIssuesEventDeletedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#issues
func (g *EventHandler) SetOnIssuesEventDeleted(callbacks ...IssuesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssuesEvent == nil {
		g.onIssuesEvent = make(map[string][]IssuesEventHandleFunc)
	}
	g.onIssuesEvent[IssuesEventDeletedAction] = callbacks
}

func (g *EventHandler) handleIssuesEventDeleted(deliveryID string, eventName string, event *github.IssuesEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if IssuesEventDeletedAction != *event.Action {
		return fmt.Errorf(
			"handleIssuesEventDeleted() called with wrong action, want %s, got %s",
			IssuesEventDeletedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		IssuesEventDeletedAction,
		IssuesEventAnyAction,
	} {
		if _, ok := g.onIssuesEvent[action]; ok {
			for _, h := range g.onIssuesEvent[action] {
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

// OnIssuesEventPinned registers callbacks listening to events of type github.IssuesEvent and action 'pinned'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnIssuesEventPinned must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#issues
func (g *EventHandler) OnIssuesEventPinned(callbacks ...IssuesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssuesEvent == nil {
		g.onIssuesEvent = make(map[string][]IssuesEventHandleFunc)
	}
	g.onIssuesEvent[IssuesEventPinnedAction] = append(
		g.onIssuesEvent[IssuesEventPinnedAction],
		callbacks...,
	)
}

// SetOnIssuesEventPinned registers callbacks listening to events of type github.IssuesEvent and action 'pinned'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnIssuesEventPinnedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#issues
func (g *EventHandler) SetOnIssuesEventPinned(callbacks ...IssuesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssuesEvent == nil {
		g.onIssuesEvent = make(map[string][]IssuesEventHandleFunc)
	}
	g.onIssuesEvent[IssuesEventPinnedAction] = callbacks
}

func (g *EventHandler) handleIssuesEventPinned(deliveryID string, eventName string, event *github.IssuesEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if IssuesEventPinnedAction != *event.Action {
		return fmt.Errorf(
			"handleIssuesEventPinned() called with wrong action, want %s, got %s",
			IssuesEventPinnedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		IssuesEventPinnedAction,
		IssuesEventAnyAction,
	} {
		if _, ok := g.onIssuesEvent[action]; ok {
			for _, h := range g.onIssuesEvent[action] {
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

// OnIssuesEventUnpinned registers callbacks listening to events of type github.IssuesEvent and action 'unpinned'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnIssuesEventUnpinned must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#issues
func (g *EventHandler) OnIssuesEventUnpinned(callbacks ...IssuesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssuesEvent == nil {
		g.onIssuesEvent = make(map[string][]IssuesEventHandleFunc)
	}
	g.onIssuesEvent[IssuesEventUnpinnedAction] = append(
		g.onIssuesEvent[IssuesEventUnpinnedAction],
		callbacks...,
	)
}

// SetOnIssuesEventUnpinned registers callbacks listening to events of type github.IssuesEvent and action 'unpinned'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnIssuesEventUnpinnedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#issues
func (g *EventHandler) SetOnIssuesEventUnpinned(callbacks ...IssuesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssuesEvent == nil {
		g.onIssuesEvent = make(map[string][]IssuesEventHandleFunc)
	}
	g.onIssuesEvent[IssuesEventUnpinnedAction] = callbacks
}

func (g *EventHandler) handleIssuesEventUnpinned(deliveryID string, eventName string, event *github.IssuesEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if IssuesEventUnpinnedAction != *event.Action {
		return fmt.Errorf(
			"handleIssuesEventUnpinned() called with wrong action, want %s, got %s",
			IssuesEventUnpinnedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		IssuesEventUnpinnedAction,
		IssuesEventAnyAction,
	} {
		if _, ok := g.onIssuesEvent[action]; ok {
			for _, h := range g.onIssuesEvent[action] {
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

// OnIssuesEventClosed registers callbacks listening to events of type github.IssuesEvent and action 'closed'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnIssuesEventClosed must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#issues
func (g *EventHandler) OnIssuesEventClosed(callbacks ...IssuesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssuesEvent == nil {
		g.onIssuesEvent = make(map[string][]IssuesEventHandleFunc)
	}
	g.onIssuesEvent[IssuesEventClosedAction] = append(
		g.onIssuesEvent[IssuesEventClosedAction],
		callbacks...,
	)
}

// SetOnIssuesEventClosed registers callbacks listening to events of type github.IssuesEvent and action 'closed'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnIssuesEventClosedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#issues
func (g *EventHandler) SetOnIssuesEventClosed(callbacks ...IssuesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssuesEvent == nil {
		g.onIssuesEvent = make(map[string][]IssuesEventHandleFunc)
	}
	g.onIssuesEvent[IssuesEventClosedAction] = callbacks
}

func (g *EventHandler) handleIssuesEventClosed(deliveryID string, eventName string, event *github.IssuesEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if IssuesEventClosedAction != *event.Action {
		return fmt.Errorf(
			"handleIssuesEventClosed() called with wrong action, want %s, got %s",
			IssuesEventClosedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		IssuesEventClosedAction,
		IssuesEventAnyAction,
	} {
		if _, ok := g.onIssuesEvent[action]; ok {
			for _, h := range g.onIssuesEvent[action] {
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

// OnIssuesEventReopened registers callbacks listening to events of type github.IssuesEvent and action 'reopened'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnIssuesEventReopened must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#issues
func (g *EventHandler) OnIssuesEventReopened(callbacks ...IssuesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssuesEvent == nil {
		g.onIssuesEvent = make(map[string][]IssuesEventHandleFunc)
	}
	g.onIssuesEvent[IssuesEventReopenedAction] = append(
		g.onIssuesEvent[IssuesEventReopenedAction],
		callbacks...,
	)
}

// SetOnIssuesEventReopened registers callbacks listening to events of type github.IssuesEvent and action 'reopened'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnIssuesEventReopenedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#issues
func (g *EventHandler) SetOnIssuesEventReopened(callbacks ...IssuesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssuesEvent == nil {
		g.onIssuesEvent = make(map[string][]IssuesEventHandleFunc)
	}
	g.onIssuesEvent[IssuesEventReopenedAction] = callbacks
}

func (g *EventHandler) handleIssuesEventReopened(deliveryID string, eventName string, event *github.IssuesEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if IssuesEventReopenedAction != *event.Action {
		return fmt.Errorf(
			"handleIssuesEventReopened() called with wrong action, want %s, got %s",
			IssuesEventReopenedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		IssuesEventReopenedAction,
		IssuesEventAnyAction,
	} {
		if _, ok := g.onIssuesEvent[action]; ok {
			for _, h := range g.onIssuesEvent[action] {
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

// OnIssuesEventAssigned registers callbacks listening to events of type github.IssuesEvent and action 'assigned'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnIssuesEventAssigned must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#issues
func (g *EventHandler) OnIssuesEventAssigned(callbacks ...IssuesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssuesEvent == nil {
		g.onIssuesEvent = make(map[string][]IssuesEventHandleFunc)
	}
	g.onIssuesEvent[IssuesEventAssignedAction] = append(
		g.onIssuesEvent[IssuesEventAssignedAction],
		callbacks...,
	)
}

// SetOnIssuesEventAssigned registers callbacks listening to events of type github.IssuesEvent and action 'assigned'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnIssuesEventAssignedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#issues
func (g *EventHandler) SetOnIssuesEventAssigned(callbacks ...IssuesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssuesEvent == nil {
		g.onIssuesEvent = make(map[string][]IssuesEventHandleFunc)
	}
	g.onIssuesEvent[IssuesEventAssignedAction] = callbacks
}

func (g *EventHandler) handleIssuesEventAssigned(deliveryID string, eventName string, event *github.IssuesEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if IssuesEventAssignedAction != *event.Action {
		return fmt.Errorf(
			"handleIssuesEventAssigned() called with wrong action, want %s, got %s",
			IssuesEventAssignedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		IssuesEventAssignedAction,
		IssuesEventAnyAction,
	} {
		if _, ok := g.onIssuesEvent[action]; ok {
			for _, h := range g.onIssuesEvent[action] {
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

// OnIssuesEventUnassigned registers callbacks listening to events of type github.IssuesEvent and action 'unassigned'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnIssuesEventUnassigned must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#issues
func (g *EventHandler) OnIssuesEventUnassigned(callbacks ...IssuesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssuesEvent == nil {
		g.onIssuesEvent = make(map[string][]IssuesEventHandleFunc)
	}
	g.onIssuesEvent[IssuesEventUnassignedAction] = append(
		g.onIssuesEvent[IssuesEventUnassignedAction],
		callbacks...,
	)
}

// SetOnIssuesEventUnassigned registers callbacks listening to events of type github.IssuesEvent and action 'unassigned'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnIssuesEventUnassignedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#issues
func (g *EventHandler) SetOnIssuesEventUnassigned(callbacks ...IssuesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssuesEvent == nil {
		g.onIssuesEvent = make(map[string][]IssuesEventHandleFunc)
	}
	g.onIssuesEvent[IssuesEventUnassignedAction] = callbacks
}

func (g *EventHandler) handleIssuesEventUnassigned(deliveryID string, eventName string, event *github.IssuesEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if IssuesEventUnassignedAction != *event.Action {
		return fmt.Errorf(
			"handleIssuesEventUnassigned() called with wrong action, want %s, got %s",
			IssuesEventUnassignedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		IssuesEventUnassignedAction,
		IssuesEventAnyAction,
	} {
		if _, ok := g.onIssuesEvent[action]; ok {
			for _, h := range g.onIssuesEvent[action] {
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

// OnIssuesEventLabeled registers callbacks listening to events of type github.IssuesEvent and action 'labeled'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnIssuesEventLabeled must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#issues
func (g *EventHandler) OnIssuesEventLabeled(callbacks ...IssuesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssuesEvent == nil {
		g.onIssuesEvent = make(map[string][]IssuesEventHandleFunc)
	}
	g.onIssuesEvent[IssuesEventLabeledAction] = append(
		g.onIssuesEvent[IssuesEventLabeledAction],
		callbacks...,
	)
}

// SetOnIssuesEventLabeled registers callbacks listening to events of type github.IssuesEvent and action 'labeled'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnIssuesEventLabeledAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#issues
func (g *EventHandler) SetOnIssuesEventLabeled(callbacks ...IssuesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssuesEvent == nil {
		g.onIssuesEvent = make(map[string][]IssuesEventHandleFunc)
	}
	g.onIssuesEvent[IssuesEventLabeledAction] = callbacks
}

func (g *EventHandler) handleIssuesEventLabeled(deliveryID string, eventName string, event *github.IssuesEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if IssuesEventLabeledAction != *event.Action {
		return fmt.Errorf(
			"handleIssuesEventLabeled() called with wrong action, want %s, got %s",
			IssuesEventLabeledAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		IssuesEventLabeledAction,
		IssuesEventAnyAction,
	} {
		if _, ok := g.onIssuesEvent[action]; ok {
			for _, h := range g.onIssuesEvent[action] {
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

// OnIssuesEventUnlabeled registers callbacks listening to events of type github.IssuesEvent and action 'unlabeled'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnIssuesEventUnlabeled must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#issues
func (g *EventHandler) OnIssuesEventUnlabeled(callbacks ...IssuesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssuesEvent == nil {
		g.onIssuesEvent = make(map[string][]IssuesEventHandleFunc)
	}
	g.onIssuesEvent[IssuesEventUnlabeledAction] = append(
		g.onIssuesEvent[IssuesEventUnlabeledAction],
		callbacks...,
	)
}

// SetOnIssuesEventUnlabeled registers callbacks listening to events of type github.IssuesEvent and action 'unlabeled'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnIssuesEventUnlabeledAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#issues
func (g *EventHandler) SetOnIssuesEventUnlabeled(callbacks ...IssuesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssuesEvent == nil {
		g.onIssuesEvent = make(map[string][]IssuesEventHandleFunc)
	}
	g.onIssuesEvent[IssuesEventUnlabeledAction] = callbacks
}

func (g *EventHandler) handleIssuesEventUnlabeled(deliveryID string, eventName string, event *github.IssuesEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if IssuesEventUnlabeledAction != *event.Action {
		return fmt.Errorf(
			"handleIssuesEventUnlabeled() called with wrong action, want %s, got %s",
			IssuesEventUnlabeledAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		IssuesEventUnlabeledAction,
		IssuesEventAnyAction,
	} {
		if _, ok := g.onIssuesEvent[action]; ok {
			for _, h := range g.onIssuesEvent[action] {
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

// OnIssuesEventLocked registers callbacks listening to events of type github.IssuesEvent and action 'locked'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnIssuesEventLocked must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#issues
func (g *EventHandler) OnIssuesEventLocked(callbacks ...IssuesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssuesEvent == nil {
		g.onIssuesEvent = make(map[string][]IssuesEventHandleFunc)
	}
	g.onIssuesEvent[IssuesEventLockedAction] = append(
		g.onIssuesEvent[IssuesEventLockedAction],
		callbacks...,
	)
}

// SetOnIssuesEventLocked registers callbacks listening to events of type github.IssuesEvent and action 'locked'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnIssuesEventLockedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#issues
func (g *EventHandler) SetOnIssuesEventLocked(callbacks ...IssuesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssuesEvent == nil {
		g.onIssuesEvent = make(map[string][]IssuesEventHandleFunc)
	}
	g.onIssuesEvent[IssuesEventLockedAction] = callbacks
}

func (g *EventHandler) handleIssuesEventLocked(deliveryID string, eventName string, event *github.IssuesEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if IssuesEventLockedAction != *event.Action {
		return fmt.Errorf(
			"handleIssuesEventLocked() called with wrong action, want %s, got %s",
			IssuesEventLockedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		IssuesEventLockedAction,
		IssuesEventAnyAction,
	} {
		if _, ok := g.onIssuesEvent[action]; ok {
			for _, h := range g.onIssuesEvent[action] {
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

// OnIssuesEventUnlocked registers callbacks listening to events of type github.IssuesEvent and action 'unlocked'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnIssuesEventUnlocked must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#issues
func (g *EventHandler) OnIssuesEventUnlocked(callbacks ...IssuesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssuesEvent == nil {
		g.onIssuesEvent = make(map[string][]IssuesEventHandleFunc)
	}
	g.onIssuesEvent[IssuesEventUnlockedAction] = append(
		g.onIssuesEvent[IssuesEventUnlockedAction],
		callbacks...,
	)
}

// SetOnIssuesEventUnlocked registers callbacks listening to events of type github.IssuesEvent and action 'unlocked'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnIssuesEventUnlockedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#issues
func (g *EventHandler) SetOnIssuesEventUnlocked(callbacks ...IssuesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssuesEvent == nil {
		g.onIssuesEvent = make(map[string][]IssuesEventHandleFunc)
	}
	g.onIssuesEvent[IssuesEventUnlockedAction] = callbacks
}

func (g *EventHandler) handleIssuesEventUnlocked(deliveryID string, eventName string, event *github.IssuesEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if IssuesEventUnlockedAction != *event.Action {
		return fmt.Errorf(
			"handleIssuesEventUnlocked() called with wrong action, want %s, got %s",
			IssuesEventUnlockedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		IssuesEventUnlockedAction,
		IssuesEventAnyAction,
	} {
		if _, ok := g.onIssuesEvent[action]; ok {
			for _, h := range g.onIssuesEvent[action] {
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

// OnIssuesEventTransferred registers callbacks listening to events of type github.IssuesEvent and action 'transferred'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnIssuesEventTransferred must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#issues
func (g *EventHandler) OnIssuesEventTransferred(callbacks ...IssuesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssuesEvent == nil {
		g.onIssuesEvent = make(map[string][]IssuesEventHandleFunc)
	}
	g.onIssuesEvent[IssuesEventTransferredAction] = append(
		g.onIssuesEvent[IssuesEventTransferredAction],
		callbacks...,
	)
}

// SetOnIssuesEventTransferred registers callbacks listening to events of type github.IssuesEvent and action 'transferred'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnIssuesEventTransferredAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#issues
func (g *EventHandler) SetOnIssuesEventTransferred(callbacks ...IssuesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssuesEvent == nil {
		g.onIssuesEvent = make(map[string][]IssuesEventHandleFunc)
	}
	g.onIssuesEvent[IssuesEventTransferredAction] = callbacks
}

func (g *EventHandler) handleIssuesEventTransferred(deliveryID string, eventName string, event *github.IssuesEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if IssuesEventTransferredAction != *event.Action {
		return fmt.Errorf(
			"handleIssuesEventTransferred() called with wrong action, want %s, got %s",
			IssuesEventTransferredAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		IssuesEventTransferredAction,
		IssuesEventAnyAction,
	} {
		if _, ok := g.onIssuesEvent[action]; ok {
			for _, h := range g.onIssuesEvent[action] {
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

// OnIssuesEventMilestoned registers callbacks listening to events of type github.IssuesEvent and action 'milestoned'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnIssuesEventMilestoned must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#issues
func (g *EventHandler) OnIssuesEventMilestoned(callbacks ...IssuesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssuesEvent == nil {
		g.onIssuesEvent = make(map[string][]IssuesEventHandleFunc)
	}
	g.onIssuesEvent[IssuesEventMilestonedAction] = append(
		g.onIssuesEvent[IssuesEventMilestonedAction],
		callbacks...,
	)
}

// SetOnIssuesEventMilestoned registers callbacks listening to events of type github.IssuesEvent and action 'milestoned'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnIssuesEventMilestonedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#issues
func (g *EventHandler) SetOnIssuesEventMilestoned(callbacks ...IssuesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssuesEvent == nil {
		g.onIssuesEvent = make(map[string][]IssuesEventHandleFunc)
	}
	g.onIssuesEvent[IssuesEventMilestonedAction] = callbacks
}

func (g *EventHandler) handleIssuesEventMilestoned(deliveryID string, eventName string, event *github.IssuesEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if IssuesEventMilestonedAction != *event.Action {
		return fmt.Errorf(
			"handleIssuesEventMilestoned() called with wrong action, want %s, got %s",
			IssuesEventMilestonedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		IssuesEventMilestonedAction,
		IssuesEventAnyAction,
	} {
		if _, ok := g.onIssuesEvent[action]; ok {
			for _, h := range g.onIssuesEvent[action] {
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

// OnIssuesEventDeMilestoned registers callbacks listening to events of type github.IssuesEvent and action 'demilestoned'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnIssuesEventDeMilestoned must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#issues
func (g *EventHandler) OnIssuesEventDeMilestoned(callbacks ...IssuesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssuesEvent == nil {
		g.onIssuesEvent = make(map[string][]IssuesEventHandleFunc)
	}
	g.onIssuesEvent[IssuesEventDeMilestonedAction] = append(
		g.onIssuesEvent[IssuesEventDeMilestonedAction],
		callbacks...,
	)
}

// SetOnIssuesEventDeMilestoned registers callbacks listening to events of type github.IssuesEvent and action 'demilestoned'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnIssuesEventDeMilestonedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#issues
func (g *EventHandler) SetOnIssuesEventDeMilestoned(callbacks ...IssuesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssuesEvent == nil {
		g.onIssuesEvent = make(map[string][]IssuesEventHandleFunc)
	}
	g.onIssuesEvent[IssuesEventDeMilestonedAction] = callbacks
}

func (g *EventHandler) handleIssuesEventDeMilestoned(deliveryID string, eventName string, event *github.IssuesEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if IssuesEventDeMilestonedAction != *event.Action {
		return fmt.Errorf(
			"handleIssuesEventDeMilestoned() called with wrong action, want %s, got %s",
			IssuesEventDeMilestonedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		IssuesEventDeMilestonedAction,
		IssuesEventAnyAction,
	} {
		if _, ok := g.onIssuesEvent[action]; ok {
			for _, h := range g.onIssuesEvent[action] {
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

// OnIssuesEventAny registers callbacks listening to any events of type github.IssuesEvent
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnIssuesEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#issues
func (g *EventHandler) OnIssuesEventAny(callbacks ...IssuesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssuesEvent == nil {
		g.onIssuesEvent = make(map[string][]IssuesEventHandleFunc)
	}
	g.onIssuesEvent[IssuesEventAnyAction] = append(
		g.onIssuesEvent[IssuesEventAnyAction],
		callbacks...,
	)
}

// SetOnIssuesEventAny registers callbacks listening to any events of type github.IssuesEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnIssuesEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#issues
func (g *EventHandler) SetOnIssuesEventAny(callbacks ...IssuesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssuesEvent == nil {
		g.onIssuesEvent = make(map[string][]IssuesEventHandleFunc)
	}
	g.onIssuesEvent[IssuesEventAnyAction] = callbacks
}

func (g *EventHandler) handleIssuesEventAny(deliveryID string, eventName string, event *github.IssuesEvent) error {
	if event == nil {
		return fmt.Errorf("event was empty or nil")
	}
	if _, ok := g.onIssuesEvent[IssuesEventAnyAction]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onIssuesEvent[IssuesEventAnyAction] {
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

// IssuesEvent handles github.IssuesEvent.
//
// Callbacks are executed in the following order:
//
// 1) All callbacks registered with OnBeforeAny are executed in parallel.
// 2) All callbacks registered with OnIssuesEvent... are executed in parallel in case the Event has actions.
// 3) All callbacks registered with OnAfterAny are executed in parallel.
//
// on any error all callbacks registered with OnError are executed in parallel.
func (g *EventHandler) IssuesEvent(deliveryID string, eventName string, event *github.IssuesEvent) error {

	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	action := *event.Action

	err := g.handleBeforeAny(deliveryID, eventName, event)
	if err != nil {
		return g.handleError(deliveryID, eventName, event, err)
	}

	switch action {

	case IssuesEventOpenedAction:
		err := g.handleIssuesEventOpened(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case IssuesEventEditedAction:
		err := g.handleIssuesEventEdited(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case IssuesEventDeletedAction:
		err := g.handleIssuesEventDeleted(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case IssuesEventPinnedAction:
		err := g.handleIssuesEventPinned(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case IssuesEventUnpinnedAction:
		err := g.handleIssuesEventUnpinned(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case IssuesEventClosedAction:
		err := g.handleIssuesEventClosed(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case IssuesEventReopenedAction:
		err := g.handleIssuesEventReopened(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case IssuesEventAssignedAction:
		err := g.handleIssuesEventAssigned(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case IssuesEventUnassignedAction:
		err := g.handleIssuesEventUnassigned(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case IssuesEventLabeledAction:
		err := g.handleIssuesEventLabeled(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case IssuesEventUnlabeledAction:
		err := g.handleIssuesEventUnlabeled(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case IssuesEventLockedAction:
		err := g.handleIssuesEventLocked(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case IssuesEventUnlockedAction:
		err := g.handleIssuesEventUnlocked(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case IssuesEventTransferredAction:
		err := g.handleIssuesEventTransferred(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case IssuesEventMilestonedAction:
		err := g.handleIssuesEventMilestoned(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case IssuesEventDeMilestonedAction:
		err := g.handleIssuesEventDeMilestoned(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	default:
		err := g.handleIssuesEventAny(deliveryID, eventName, event)
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
