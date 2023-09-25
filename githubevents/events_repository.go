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
	// RepositoryEvent is the event name of github.RepositoryEvent's
	RepositoryEvent = "repository"

	// RepositoryEventAnyAction is used to identify callbacks
	// listening to all events of type github.RepositoryEvent
	RepositoryEventAnyAction = "*"

	// RepositoryEventCreatedAction is used to identify callbacks
	// listening to events of type github.RepositoryEvent and action "created"
	RepositoryEventCreatedAction = "created"

	// RepositoryEventDeletedAction is used to identify callbacks
	// listening to events of type github.RepositoryEvent and action "deleted"
	RepositoryEventDeletedAction = "deleted"

	// RepositoryEventArchivedAction is used to identify callbacks
	// listening to events of type github.RepositoryEvent and action "archived"
	RepositoryEventArchivedAction = "archived"

	// RepositoryEventUnarchivedAction is used to identify callbacks
	// listening to events of type github.RepositoryEvent and action "unarchived"
	RepositoryEventUnarchivedAction = "unarchived"

	// RepositoryEventEditedAction is used to identify callbacks
	// listening to events of type github.RepositoryEvent and action "edited"
	RepositoryEventEditedAction = "edited"

	// RepositoryEventRenamedAction is used to identify callbacks
	// listening to events of type github.RepositoryEvent and action "renamed"
	RepositoryEventRenamedAction = "renamed"

	// RepositoryEventTransferredAction is used to identify callbacks
	// listening to events of type github.RepositoryEvent and action "transferred"
	RepositoryEventTransferredAction = "transferred"

	// RepositoryEventPublicizedAction is used to identify callbacks
	// listening to events of type github.RepositoryEvent and action "publicized"
	RepositoryEventPublicizedAction = "publicized"

	// RepositoryEventPrivatizedAction is used to identify callbacks
	// listening to events of type github.RepositoryEvent and action "privatized"
	RepositoryEventPrivatizedAction = "privatized"
)

// RepositoryEventHandleFunc represents a callback function triggered on github.RepositoryEvent's.
// 'deliveryID' (type: string) is the unique webhook delivery ID.
// 'eventName' (type: string) is the name of the event.
// 'event' (type: *github.RepositoryEvent) is the webhook payload.
type RepositoryEventHandleFunc func(deliveryID string, eventName string, event *github.RepositoryEvent) error

// OnRepositoryEventCreated registers callbacks listening to events of type github.RepositoryEvent and action 'created'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnRepositoryEventCreated must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#repository
func (g *EventHandler) OnRepositoryEventCreated(callbacks ...RepositoryEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onRepositoryEvent == nil {
		g.onRepositoryEvent = make(map[string][]RepositoryEventHandleFunc)
	}
	g.onRepositoryEvent[RepositoryEventCreatedAction] = append(
		g.onRepositoryEvent[RepositoryEventCreatedAction],
		callbacks...,
	)
}

// SetOnRepositoryEventCreated registers callbacks listening to events of type github.RepositoryEvent and action 'created'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnRepositoryEventCreatedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#repository
func (g *EventHandler) SetOnRepositoryEventCreated(callbacks ...RepositoryEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onRepositoryEvent == nil {
		g.onRepositoryEvent = make(map[string][]RepositoryEventHandleFunc)
	}
	g.onRepositoryEvent[RepositoryEventCreatedAction] = callbacks
}

func (g *EventHandler) handleRepositoryEventCreated(deliveryID string, eventName string, event *github.RepositoryEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if RepositoryEventCreatedAction != *event.Action {
		return fmt.Errorf(
			"handleRepositoryEventCreated() called with wrong action, want %s, got %s",
			RepositoryEventCreatedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		RepositoryEventCreatedAction,
		RepositoryEventAnyAction,
	} {
		if _, ok := g.onRepositoryEvent[action]; ok {
			for _, h := range g.onRepositoryEvent[action] {
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

// OnRepositoryEventDeleted registers callbacks listening to events of type github.RepositoryEvent and action 'deleted'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnRepositoryEventDeleted must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#repository
func (g *EventHandler) OnRepositoryEventDeleted(callbacks ...RepositoryEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onRepositoryEvent == nil {
		g.onRepositoryEvent = make(map[string][]RepositoryEventHandleFunc)
	}
	g.onRepositoryEvent[RepositoryEventDeletedAction] = append(
		g.onRepositoryEvent[RepositoryEventDeletedAction],
		callbacks...,
	)
}

// SetOnRepositoryEventDeleted registers callbacks listening to events of type github.RepositoryEvent and action 'deleted'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnRepositoryEventDeletedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#repository
func (g *EventHandler) SetOnRepositoryEventDeleted(callbacks ...RepositoryEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onRepositoryEvent == nil {
		g.onRepositoryEvent = make(map[string][]RepositoryEventHandleFunc)
	}
	g.onRepositoryEvent[RepositoryEventDeletedAction] = callbacks
}

func (g *EventHandler) handleRepositoryEventDeleted(deliveryID string, eventName string, event *github.RepositoryEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if RepositoryEventDeletedAction != *event.Action {
		return fmt.Errorf(
			"handleRepositoryEventDeleted() called with wrong action, want %s, got %s",
			RepositoryEventDeletedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		RepositoryEventDeletedAction,
		RepositoryEventAnyAction,
	} {
		if _, ok := g.onRepositoryEvent[action]; ok {
			for _, h := range g.onRepositoryEvent[action] {
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

// OnRepositoryEventArchived registers callbacks listening to events of type github.RepositoryEvent and action 'archived'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnRepositoryEventArchived must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#repository
func (g *EventHandler) OnRepositoryEventArchived(callbacks ...RepositoryEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onRepositoryEvent == nil {
		g.onRepositoryEvent = make(map[string][]RepositoryEventHandleFunc)
	}
	g.onRepositoryEvent[RepositoryEventArchivedAction] = append(
		g.onRepositoryEvent[RepositoryEventArchivedAction],
		callbacks...,
	)
}

// SetOnRepositoryEventArchived registers callbacks listening to events of type github.RepositoryEvent and action 'archived'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnRepositoryEventArchivedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#repository
func (g *EventHandler) SetOnRepositoryEventArchived(callbacks ...RepositoryEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onRepositoryEvent == nil {
		g.onRepositoryEvent = make(map[string][]RepositoryEventHandleFunc)
	}
	g.onRepositoryEvent[RepositoryEventArchivedAction] = callbacks
}

func (g *EventHandler) handleRepositoryEventArchived(deliveryID string, eventName string, event *github.RepositoryEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if RepositoryEventArchivedAction != *event.Action {
		return fmt.Errorf(
			"handleRepositoryEventArchived() called with wrong action, want %s, got %s",
			RepositoryEventArchivedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		RepositoryEventArchivedAction,
		RepositoryEventAnyAction,
	} {
		if _, ok := g.onRepositoryEvent[action]; ok {
			for _, h := range g.onRepositoryEvent[action] {
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

// OnRepositoryEventUnarchived registers callbacks listening to events of type github.RepositoryEvent and action 'unarchived'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnRepositoryEventUnarchived must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#repository
func (g *EventHandler) OnRepositoryEventUnarchived(callbacks ...RepositoryEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onRepositoryEvent == nil {
		g.onRepositoryEvent = make(map[string][]RepositoryEventHandleFunc)
	}
	g.onRepositoryEvent[RepositoryEventUnarchivedAction] = append(
		g.onRepositoryEvent[RepositoryEventUnarchivedAction],
		callbacks...,
	)
}

// SetOnRepositoryEventUnarchived registers callbacks listening to events of type github.RepositoryEvent and action 'unarchived'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnRepositoryEventUnarchivedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#repository
func (g *EventHandler) SetOnRepositoryEventUnarchived(callbacks ...RepositoryEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onRepositoryEvent == nil {
		g.onRepositoryEvent = make(map[string][]RepositoryEventHandleFunc)
	}
	g.onRepositoryEvent[RepositoryEventUnarchivedAction] = callbacks
}

func (g *EventHandler) handleRepositoryEventUnarchived(deliveryID string, eventName string, event *github.RepositoryEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if RepositoryEventUnarchivedAction != *event.Action {
		return fmt.Errorf(
			"handleRepositoryEventUnarchived() called with wrong action, want %s, got %s",
			RepositoryEventUnarchivedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		RepositoryEventUnarchivedAction,
		RepositoryEventAnyAction,
	} {
		if _, ok := g.onRepositoryEvent[action]; ok {
			for _, h := range g.onRepositoryEvent[action] {
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

// OnRepositoryEventEdited registers callbacks listening to events of type github.RepositoryEvent and action 'edited'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnRepositoryEventEdited must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#repository
func (g *EventHandler) OnRepositoryEventEdited(callbacks ...RepositoryEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onRepositoryEvent == nil {
		g.onRepositoryEvent = make(map[string][]RepositoryEventHandleFunc)
	}
	g.onRepositoryEvent[RepositoryEventEditedAction] = append(
		g.onRepositoryEvent[RepositoryEventEditedAction],
		callbacks...,
	)
}

// SetOnRepositoryEventEdited registers callbacks listening to events of type github.RepositoryEvent and action 'edited'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnRepositoryEventEditedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#repository
func (g *EventHandler) SetOnRepositoryEventEdited(callbacks ...RepositoryEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onRepositoryEvent == nil {
		g.onRepositoryEvent = make(map[string][]RepositoryEventHandleFunc)
	}
	g.onRepositoryEvent[RepositoryEventEditedAction] = callbacks
}

func (g *EventHandler) handleRepositoryEventEdited(deliveryID string, eventName string, event *github.RepositoryEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if RepositoryEventEditedAction != *event.Action {
		return fmt.Errorf(
			"handleRepositoryEventEdited() called with wrong action, want %s, got %s",
			RepositoryEventEditedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		RepositoryEventEditedAction,
		RepositoryEventAnyAction,
	} {
		if _, ok := g.onRepositoryEvent[action]; ok {
			for _, h := range g.onRepositoryEvent[action] {
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

// OnRepositoryEventRenamed registers callbacks listening to events of type github.RepositoryEvent and action 'renamed'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnRepositoryEventRenamed must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#repository
func (g *EventHandler) OnRepositoryEventRenamed(callbacks ...RepositoryEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onRepositoryEvent == nil {
		g.onRepositoryEvent = make(map[string][]RepositoryEventHandleFunc)
	}
	g.onRepositoryEvent[RepositoryEventRenamedAction] = append(
		g.onRepositoryEvent[RepositoryEventRenamedAction],
		callbacks...,
	)
}

// SetOnRepositoryEventRenamed registers callbacks listening to events of type github.RepositoryEvent and action 'renamed'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnRepositoryEventRenamedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#repository
func (g *EventHandler) SetOnRepositoryEventRenamed(callbacks ...RepositoryEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onRepositoryEvent == nil {
		g.onRepositoryEvent = make(map[string][]RepositoryEventHandleFunc)
	}
	g.onRepositoryEvent[RepositoryEventRenamedAction] = callbacks
}

func (g *EventHandler) handleRepositoryEventRenamed(deliveryID string, eventName string, event *github.RepositoryEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if RepositoryEventRenamedAction != *event.Action {
		return fmt.Errorf(
			"handleRepositoryEventRenamed() called with wrong action, want %s, got %s",
			RepositoryEventRenamedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		RepositoryEventRenamedAction,
		RepositoryEventAnyAction,
	} {
		if _, ok := g.onRepositoryEvent[action]; ok {
			for _, h := range g.onRepositoryEvent[action] {
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

// OnRepositoryEventTransferred registers callbacks listening to events of type github.RepositoryEvent and action 'transferred'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnRepositoryEventTransferred must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#repository
func (g *EventHandler) OnRepositoryEventTransferred(callbacks ...RepositoryEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onRepositoryEvent == nil {
		g.onRepositoryEvent = make(map[string][]RepositoryEventHandleFunc)
	}
	g.onRepositoryEvent[RepositoryEventTransferredAction] = append(
		g.onRepositoryEvent[RepositoryEventTransferredAction],
		callbacks...,
	)
}

// SetOnRepositoryEventTransferred registers callbacks listening to events of type github.RepositoryEvent and action 'transferred'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnRepositoryEventTransferredAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#repository
func (g *EventHandler) SetOnRepositoryEventTransferred(callbacks ...RepositoryEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onRepositoryEvent == nil {
		g.onRepositoryEvent = make(map[string][]RepositoryEventHandleFunc)
	}
	g.onRepositoryEvent[RepositoryEventTransferredAction] = callbacks
}

func (g *EventHandler) handleRepositoryEventTransferred(deliveryID string, eventName string, event *github.RepositoryEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if RepositoryEventTransferredAction != *event.Action {
		return fmt.Errorf(
			"handleRepositoryEventTransferred() called with wrong action, want %s, got %s",
			RepositoryEventTransferredAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		RepositoryEventTransferredAction,
		RepositoryEventAnyAction,
	} {
		if _, ok := g.onRepositoryEvent[action]; ok {
			for _, h := range g.onRepositoryEvent[action] {
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

// OnRepositoryEventPublicized registers callbacks listening to events of type github.RepositoryEvent and action 'publicized'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnRepositoryEventPublicized must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#repository
func (g *EventHandler) OnRepositoryEventPublicized(callbacks ...RepositoryEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onRepositoryEvent == nil {
		g.onRepositoryEvent = make(map[string][]RepositoryEventHandleFunc)
	}
	g.onRepositoryEvent[RepositoryEventPublicizedAction] = append(
		g.onRepositoryEvent[RepositoryEventPublicizedAction],
		callbacks...,
	)
}

// SetOnRepositoryEventPublicized registers callbacks listening to events of type github.RepositoryEvent and action 'publicized'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnRepositoryEventPublicizedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#repository
func (g *EventHandler) SetOnRepositoryEventPublicized(callbacks ...RepositoryEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onRepositoryEvent == nil {
		g.onRepositoryEvent = make(map[string][]RepositoryEventHandleFunc)
	}
	g.onRepositoryEvent[RepositoryEventPublicizedAction] = callbacks
}

func (g *EventHandler) handleRepositoryEventPublicized(deliveryID string, eventName string, event *github.RepositoryEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if RepositoryEventPublicizedAction != *event.Action {
		return fmt.Errorf(
			"handleRepositoryEventPublicized() called with wrong action, want %s, got %s",
			RepositoryEventPublicizedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		RepositoryEventPublicizedAction,
		RepositoryEventAnyAction,
	} {
		if _, ok := g.onRepositoryEvent[action]; ok {
			for _, h := range g.onRepositoryEvent[action] {
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

// OnRepositoryEventPrivatized registers callbacks listening to events of type github.RepositoryEvent and action 'privatized'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnRepositoryEventPrivatized must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#repository
func (g *EventHandler) OnRepositoryEventPrivatized(callbacks ...RepositoryEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onRepositoryEvent == nil {
		g.onRepositoryEvent = make(map[string][]RepositoryEventHandleFunc)
	}
	g.onRepositoryEvent[RepositoryEventPrivatizedAction] = append(
		g.onRepositoryEvent[RepositoryEventPrivatizedAction],
		callbacks...,
	)
}

// SetOnRepositoryEventPrivatized registers callbacks listening to events of type github.RepositoryEvent and action 'privatized'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnRepositoryEventPrivatizedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#repository
func (g *EventHandler) SetOnRepositoryEventPrivatized(callbacks ...RepositoryEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onRepositoryEvent == nil {
		g.onRepositoryEvent = make(map[string][]RepositoryEventHandleFunc)
	}
	g.onRepositoryEvent[RepositoryEventPrivatizedAction] = callbacks
}

func (g *EventHandler) handleRepositoryEventPrivatized(deliveryID string, eventName string, event *github.RepositoryEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if RepositoryEventPrivatizedAction != *event.Action {
		return fmt.Errorf(
			"handleRepositoryEventPrivatized() called with wrong action, want %s, got %s",
			RepositoryEventPrivatizedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		RepositoryEventPrivatizedAction,
		RepositoryEventAnyAction,
	} {
		if _, ok := g.onRepositoryEvent[action]; ok {
			for _, h := range g.onRepositoryEvent[action] {
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

// OnRepositoryEventAny registers callbacks listening to any events of type github.RepositoryEvent
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnRepositoryEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#repository
func (g *EventHandler) OnRepositoryEventAny(callbacks ...RepositoryEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onRepositoryEvent == nil {
		g.onRepositoryEvent = make(map[string][]RepositoryEventHandleFunc)
	}
	g.onRepositoryEvent[RepositoryEventAnyAction] = append(
		g.onRepositoryEvent[RepositoryEventAnyAction],
		callbacks...,
	)
}

// SetOnRepositoryEventAny registers callbacks listening to any events of type github.RepositoryEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnRepositoryEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#repository
func (g *EventHandler) SetOnRepositoryEventAny(callbacks ...RepositoryEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onRepositoryEvent == nil {
		g.onRepositoryEvent = make(map[string][]RepositoryEventHandleFunc)
	}
	g.onRepositoryEvent[RepositoryEventAnyAction] = callbacks
}

func (g *EventHandler) handleRepositoryEventAny(deliveryID string, eventName string, event *github.RepositoryEvent) error {
	if event == nil {
		return fmt.Errorf("event was empty or nil")
	}
	if _, ok := g.onRepositoryEvent[RepositoryEventAnyAction]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onRepositoryEvent[RepositoryEventAnyAction] {
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

// RepositoryEvent handles github.RepositoryEvent.
//
// Callbacks are executed in the following order:
//
// 1) All callbacks registered with OnBeforeAny are executed in parallel.
// 2) All callbacks registered with OnRepositoryEvent... are executed in parallel in case the Event has actions.
// 3) All callbacks registered with OnAfterAny are executed in parallel.
//
// on any error all callbacks registered with OnError are executed in parallel.
func (g *EventHandler) RepositoryEvent(deliveryID string, eventName string, event *github.RepositoryEvent) error {

	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	action := *event.Action

	err := g.handleBeforeAny(deliveryID, eventName, event)
	if err != nil {
		return g.handleError(deliveryID, eventName, event, err)
	}

	switch action {

	case RepositoryEventCreatedAction:
		err := g.handleRepositoryEventCreated(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case RepositoryEventDeletedAction:
		err := g.handleRepositoryEventDeleted(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case RepositoryEventArchivedAction:
		err := g.handleRepositoryEventArchived(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case RepositoryEventUnarchivedAction:
		err := g.handleRepositoryEventUnarchived(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case RepositoryEventEditedAction:
		err := g.handleRepositoryEventEdited(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case RepositoryEventRenamedAction:
		err := g.handleRepositoryEventRenamed(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case RepositoryEventTransferredAction:
		err := g.handleRepositoryEventTransferred(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case RepositoryEventPublicizedAction:
		err := g.handleRepositoryEventPublicized(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case RepositoryEventPrivatizedAction:
		err := g.handleRepositoryEventPrivatized(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	default:
		err := g.handleRepositoryEventAny(deliveryID, eventName, event)
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
