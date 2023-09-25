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
	// DiscussionEvent is the event name of github.DiscussionEvent's
	DiscussionEvent = "discussion"

	// DiscussionEventAnyAction is used to identify callbacks
	// listening to all events of type github.DiscussionEvent
	DiscussionEventAnyAction = "*"

	// DiscussionEventCreatedAction is used to identify callbacks
	// listening to events of type github.DiscussionEvent and action "created"
	DiscussionEventCreatedAction = "created"

	// DiscussionEventEditedAction is used to identify callbacks
	// listening to events of type github.DiscussionEvent and action "edited"
	DiscussionEventEditedAction = "edited"

	// DiscussionEventDeletedAction is used to identify callbacks
	// listening to events of type github.DiscussionEvent and action "deleted"
	DiscussionEventDeletedAction = "deleted"

	// DiscussionEventPinnedAction is used to identify callbacks
	// listening to events of type github.DiscussionEvent and action "pinned"
	DiscussionEventPinnedAction = "pinned"

	// DiscussionEventUnpinnedAction is used to identify callbacks
	// listening to events of type github.DiscussionEvent and action "unpinned"
	DiscussionEventUnpinnedAction = "unpinned"

	// DiscussionEventLockedAction is used to identify callbacks
	// listening to events of type github.DiscussionEvent and action "locked"
	DiscussionEventLockedAction = "locked"

	// DiscussionEventUnlockedAction is used to identify callbacks
	// listening to events of type github.DiscussionEvent and action "unlocked"
	DiscussionEventUnlockedAction = "unlocked"

	// DiscussionEventTransferredAction is used to identify callbacks
	// listening to events of type github.DiscussionEvent and action "transferred"
	DiscussionEventTransferredAction = "transferred"

	// DiscussionEventCategoryChangedAction is used to identify callbacks
	// listening to events of type github.DiscussionEvent and action "category_changed"
	DiscussionEventCategoryChangedAction = "category_changed"

	// DiscussionEventAnsweredAction is used to identify callbacks
	// listening to events of type github.DiscussionEvent and action "answered"
	DiscussionEventAnsweredAction = "answered"

	// DiscussionEventUnansweredAction is used to identify callbacks
	// listening to events of type github.DiscussionEvent and action "unanswered"
	DiscussionEventUnansweredAction = "unanswered"

	// DiscussionEventLabeledAction is used to identify callbacks
	// listening to events of type github.DiscussionEvent and action "labeled"
	DiscussionEventLabeledAction = "labeled"

	// DiscussionEventUnlabeledAction is used to identify callbacks
	// listening to events of type github.DiscussionEvent and action "unlabeled"
	DiscussionEventUnlabeledAction = "unlabeled"
)

// DiscussionEventHandleFunc represents a callback function triggered on github.DiscussionEvent's.
// 'deliveryID' (type: string) is the unique webhook delivery ID.
// 'eventName' (type: string) is the name of the event.
// 'event' (type: *github.DiscussionEvent) is the webhook payload.
type DiscussionEventHandleFunc func(deliveryID string, eventName string, event *github.DiscussionEvent) error

// OnDiscussionEventCreated registers callbacks listening to events of type github.DiscussionEvent and action 'created'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnDiscussionEventCreated must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#discussion
func (g *EventHandler) OnDiscussionEventCreated(callbacks ...DiscussionEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onDiscussionEvent == nil {
		g.onDiscussionEvent = make(map[string][]DiscussionEventHandleFunc)
	}
	g.onDiscussionEvent[DiscussionEventCreatedAction] = append(
		g.onDiscussionEvent[DiscussionEventCreatedAction],
		callbacks...,
	)
}

// SetOnDiscussionEventCreated registers callbacks listening to events of type github.DiscussionEvent and action 'created'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnDiscussionEventCreatedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#discussion
func (g *EventHandler) SetOnDiscussionEventCreated(callbacks ...DiscussionEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onDiscussionEvent == nil {
		g.onDiscussionEvent = make(map[string][]DiscussionEventHandleFunc)
	}
	g.onDiscussionEvent[DiscussionEventCreatedAction] = callbacks
}

func (g *EventHandler) handleDiscussionEventCreated(deliveryID string, eventName string, event *github.DiscussionEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if DiscussionEventCreatedAction != *event.Action {
		return fmt.Errorf(
			"handleDiscussionEventCreated() called with wrong action, want %s, got %s",
			DiscussionEventCreatedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		DiscussionEventCreatedAction,
		DiscussionEventAnyAction,
	} {
		if _, ok := g.onDiscussionEvent[action]; ok {
			for _, h := range g.onDiscussionEvent[action] {
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

// OnDiscussionEventEdited registers callbacks listening to events of type github.DiscussionEvent and action 'edited'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnDiscussionEventEdited must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#discussion
func (g *EventHandler) OnDiscussionEventEdited(callbacks ...DiscussionEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onDiscussionEvent == nil {
		g.onDiscussionEvent = make(map[string][]DiscussionEventHandleFunc)
	}
	g.onDiscussionEvent[DiscussionEventEditedAction] = append(
		g.onDiscussionEvent[DiscussionEventEditedAction],
		callbacks...,
	)
}

// SetOnDiscussionEventEdited registers callbacks listening to events of type github.DiscussionEvent and action 'edited'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnDiscussionEventEditedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#discussion
func (g *EventHandler) SetOnDiscussionEventEdited(callbacks ...DiscussionEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onDiscussionEvent == nil {
		g.onDiscussionEvent = make(map[string][]DiscussionEventHandleFunc)
	}
	g.onDiscussionEvent[DiscussionEventEditedAction] = callbacks
}

func (g *EventHandler) handleDiscussionEventEdited(deliveryID string, eventName string, event *github.DiscussionEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if DiscussionEventEditedAction != *event.Action {
		return fmt.Errorf(
			"handleDiscussionEventEdited() called with wrong action, want %s, got %s",
			DiscussionEventEditedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		DiscussionEventEditedAction,
		DiscussionEventAnyAction,
	} {
		if _, ok := g.onDiscussionEvent[action]; ok {
			for _, h := range g.onDiscussionEvent[action] {
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

// OnDiscussionEventDeleted registers callbacks listening to events of type github.DiscussionEvent and action 'deleted'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnDiscussionEventDeleted must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#discussion
func (g *EventHandler) OnDiscussionEventDeleted(callbacks ...DiscussionEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onDiscussionEvent == nil {
		g.onDiscussionEvent = make(map[string][]DiscussionEventHandleFunc)
	}
	g.onDiscussionEvent[DiscussionEventDeletedAction] = append(
		g.onDiscussionEvent[DiscussionEventDeletedAction],
		callbacks...,
	)
}

// SetOnDiscussionEventDeleted registers callbacks listening to events of type github.DiscussionEvent and action 'deleted'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnDiscussionEventDeletedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#discussion
func (g *EventHandler) SetOnDiscussionEventDeleted(callbacks ...DiscussionEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onDiscussionEvent == nil {
		g.onDiscussionEvent = make(map[string][]DiscussionEventHandleFunc)
	}
	g.onDiscussionEvent[DiscussionEventDeletedAction] = callbacks
}

func (g *EventHandler) handleDiscussionEventDeleted(deliveryID string, eventName string, event *github.DiscussionEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if DiscussionEventDeletedAction != *event.Action {
		return fmt.Errorf(
			"handleDiscussionEventDeleted() called with wrong action, want %s, got %s",
			DiscussionEventDeletedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		DiscussionEventDeletedAction,
		DiscussionEventAnyAction,
	} {
		if _, ok := g.onDiscussionEvent[action]; ok {
			for _, h := range g.onDiscussionEvent[action] {
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

// OnDiscussionEventPinned registers callbacks listening to events of type github.DiscussionEvent and action 'pinned'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnDiscussionEventPinned must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#discussion
func (g *EventHandler) OnDiscussionEventPinned(callbacks ...DiscussionEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onDiscussionEvent == nil {
		g.onDiscussionEvent = make(map[string][]DiscussionEventHandleFunc)
	}
	g.onDiscussionEvent[DiscussionEventPinnedAction] = append(
		g.onDiscussionEvent[DiscussionEventPinnedAction],
		callbacks...,
	)
}

// SetOnDiscussionEventPinned registers callbacks listening to events of type github.DiscussionEvent and action 'pinned'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnDiscussionEventPinnedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#discussion
func (g *EventHandler) SetOnDiscussionEventPinned(callbacks ...DiscussionEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onDiscussionEvent == nil {
		g.onDiscussionEvent = make(map[string][]DiscussionEventHandleFunc)
	}
	g.onDiscussionEvent[DiscussionEventPinnedAction] = callbacks
}

func (g *EventHandler) handleDiscussionEventPinned(deliveryID string, eventName string, event *github.DiscussionEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if DiscussionEventPinnedAction != *event.Action {
		return fmt.Errorf(
			"handleDiscussionEventPinned() called with wrong action, want %s, got %s",
			DiscussionEventPinnedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		DiscussionEventPinnedAction,
		DiscussionEventAnyAction,
	} {
		if _, ok := g.onDiscussionEvent[action]; ok {
			for _, h := range g.onDiscussionEvent[action] {
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

// OnDiscussionEventUnpinned registers callbacks listening to events of type github.DiscussionEvent and action 'unpinned'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnDiscussionEventUnpinned must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#discussion
func (g *EventHandler) OnDiscussionEventUnpinned(callbacks ...DiscussionEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onDiscussionEvent == nil {
		g.onDiscussionEvent = make(map[string][]DiscussionEventHandleFunc)
	}
	g.onDiscussionEvent[DiscussionEventUnpinnedAction] = append(
		g.onDiscussionEvent[DiscussionEventUnpinnedAction],
		callbacks...,
	)
}

// SetOnDiscussionEventUnpinned registers callbacks listening to events of type github.DiscussionEvent and action 'unpinned'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnDiscussionEventUnpinnedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#discussion
func (g *EventHandler) SetOnDiscussionEventUnpinned(callbacks ...DiscussionEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onDiscussionEvent == nil {
		g.onDiscussionEvent = make(map[string][]DiscussionEventHandleFunc)
	}
	g.onDiscussionEvent[DiscussionEventUnpinnedAction] = callbacks
}

func (g *EventHandler) handleDiscussionEventUnpinned(deliveryID string, eventName string, event *github.DiscussionEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if DiscussionEventUnpinnedAction != *event.Action {
		return fmt.Errorf(
			"handleDiscussionEventUnpinned() called with wrong action, want %s, got %s",
			DiscussionEventUnpinnedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		DiscussionEventUnpinnedAction,
		DiscussionEventAnyAction,
	} {
		if _, ok := g.onDiscussionEvent[action]; ok {
			for _, h := range g.onDiscussionEvent[action] {
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

// OnDiscussionEventLocked registers callbacks listening to events of type github.DiscussionEvent and action 'locked'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnDiscussionEventLocked must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#discussion
func (g *EventHandler) OnDiscussionEventLocked(callbacks ...DiscussionEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onDiscussionEvent == nil {
		g.onDiscussionEvent = make(map[string][]DiscussionEventHandleFunc)
	}
	g.onDiscussionEvent[DiscussionEventLockedAction] = append(
		g.onDiscussionEvent[DiscussionEventLockedAction],
		callbacks...,
	)
}

// SetOnDiscussionEventLocked registers callbacks listening to events of type github.DiscussionEvent and action 'locked'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnDiscussionEventLockedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#discussion
func (g *EventHandler) SetOnDiscussionEventLocked(callbacks ...DiscussionEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onDiscussionEvent == nil {
		g.onDiscussionEvent = make(map[string][]DiscussionEventHandleFunc)
	}
	g.onDiscussionEvent[DiscussionEventLockedAction] = callbacks
}

func (g *EventHandler) handleDiscussionEventLocked(deliveryID string, eventName string, event *github.DiscussionEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if DiscussionEventLockedAction != *event.Action {
		return fmt.Errorf(
			"handleDiscussionEventLocked() called with wrong action, want %s, got %s",
			DiscussionEventLockedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		DiscussionEventLockedAction,
		DiscussionEventAnyAction,
	} {
		if _, ok := g.onDiscussionEvent[action]; ok {
			for _, h := range g.onDiscussionEvent[action] {
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

// OnDiscussionEventUnlocked registers callbacks listening to events of type github.DiscussionEvent and action 'unlocked'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnDiscussionEventUnlocked must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#discussion
func (g *EventHandler) OnDiscussionEventUnlocked(callbacks ...DiscussionEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onDiscussionEvent == nil {
		g.onDiscussionEvent = make(map[string][]DiscussionEventHandleFunc)
	}
	g.onDiscussionEvent[DiscussionEventUnlockedAction] = append(
		g.onDiscussionEvent[DiscussionEventUnlockedAction],
		callbacks...,
	)
}

// SetOnDiscussionEventUnlocked registers callbacks listening to events of type github.DiscussionEvent and action 'unlocked'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnDiscussionEventUnlockedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#discussion
func (g *EventHandler) SetOnDiscussionEventUnlocked(callbacks ...DiscussionEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onDiscussionEvent == nil {
		g.onDiscussionEvent = make(map[string][]DiscussionEventHandleFunc)
	}
	g.onDiscussionEvent[DiscussionEventUnlockedAction] = callbacks
}

func (g *EventHandler) handleDiscussionEventUnlocked(deliveryID string, eventName string, event *github.DiscussionEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if DiscussionEventUnlockedAction != *event.Action {
		return fmt.Errorf(
			"handleDiscussionEventUnlocked() called with wrong action, want %s, got %s",
			DiscussionEventUnlockedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		DiscussionEventUnlockedAction,
		DiscussionEventAnyAction,
	} {
		if _, ok := g.onDiscussionEvent[action]; ok {
			for _, h := range g.onDiscussionEvent[action] {
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

// OnDiscussionEventTransferred registers callbacks listening to events of type github.DiscussionEvent and action 'transferred'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnDiscussionEventTransferred must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#discussion
func (g *EventHandler) OnDiscussionEventTransferred(callbacks ...DiscussionEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onDiscussionEvent == nil {
		g.onDiscussionEvent = make(map[string][]DiscussionEventHandleFunc)
	}
	g.onDiscussionEvent[DiscussionEventTransferredAction] = append(
		g.onDiscussionEvent[DiscussionEventTransferredAction],
		callbacks...,
	)
}

// SetOnDiscussionEventTransferred registers callbacks listening to events of type github.DiscussionEvent and action 'transferred'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnDiscussionEventTransferredAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#discussion
func (g *EventHandler) SetOnDiscussionEventTransferred(callbacks ...DiscussionEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onDiscussionEvent == nil {
		g.onDiscussionEvent = make(map[string][]DiscussionEventHandleFunc)
	}
	g.onDiscussionEvent[DiscussionEventTransferredAction] = callbacks
}

func (g *EventHandler) handleDiscussionEventTransferred(deliveryID string, eventName string, event *github.DiscussionEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if DiscussionEventTransferredAction != *event.Action {
		return fmt.Errorf(
			"handleDiscussionEventTransferred() called with wrong action, want %s, got %s",
			DiscussionEventTransferredAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		DiscussionEventTransferredAction,
		DiscussionEventAnyAction,
	} {
		if _, ok := g.onDiscussionEvent[action]; ok {
			for _, h := range g.onDiscussionEvent[action] {
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

// OnDiscussionEventCategoryChanged registers callbacks listening to events of type github.DiscussionEvent and action 'category_changed'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnDiscussionEventCategoryChanged must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#discussion
func (g *EventHandler) OnDiscussionEventCategoryChanged(callbacks ...DiscussionEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onDiscussionEvent == nil {
		g.onDiscussionEvent = make(map[string][]DiscussionEventHandleFunc)
	}
	g.onDiscussionEvent[DiscussionEventCategoryChangedAction] = append(
		g.onDiscussionEvent[DiscussionEventCategoryChangedAction],
		callbacks...,
	)
}

// SetOnDiscussionEventCategoryChanged registers callbacks listening to events of type github.DiscussionEvent and action 'category_changed'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnDiscussionEventCategoryChangedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#discussion
func (g *EventHandler) SetOnDiscussionEventCategoryChanged(callbacks ...DiscussionEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onDiscussionEvent == nil {
		g.onDiscussionEvent = make(map[string][]DiscussionEventHandleFunc)
	}
	g.onDiscussionEvent[DiscussionEventCategoryChangedAction] = callbacks
}

func (g *EventHandler) handleDiscussionEventCategoryChanged(deliveryID string, eventName string, event *github.DiscussionEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if DiscussionEventCategoryChangedAction != *event.Action {
		return fmt.Errorf(
			"handleDiscussionEventCategoryChanged() called with wrong action, want %s, got %s",
			DiscussionEventCategoryChangedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		DiscussionEventCategoryChangedAction,
		DiscussionEventAnyAction,
	} {
		if _, ok := g.onDiscussionEvent[action]; ok {
			for _, h := range g.onDiscussionEvent[action] {
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

// OnDiscussionEventAnswered registers callbacks listening to events of type github.DiscussionEvent and action 'answered'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnDiscussionEventAnswered must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#discussion
func (g *EventHandler) OnDiscussionEventAnswered(callbacks ...DiscussionEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onDiscussionEvent == nil {
		g.onDiscussionEvent = make(map[string][]DiscussionEventHandleFunc)
	}
	g.onDiscussionEvent[DiscussionEventAnsweredAction] = append(
		g.onDiscussionEvent[DiscussionEventAnsweredAction],
		callbacks...,
	)
}

// SetOnDiscussionEventAnswered registers callbacks listening to events of type github.DiscussionEvent and action 'answered'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnDiscussionEventAnsweredAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#discussion
func (g *EventHandler) SetOnDiscussionEventAnswered(callbacks ...DiscussionEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onDiscussionEvent == nil {
		g.onDiscussionEvent = make(map[string][]DiscussionEventHandleFunc)
	}
	g.onDiscussionEvent[DiscussionEventAnsweredAction] = callbacks
}

func (g *EventHandler) handleDiscussionEventAnswered(deliveryID string, eventName string, event *github.DiscussionEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if DiscussionEventAnsweredAction != *event.Action {
		return fmt.Errorf(
			"handleDiscussionEventAnswered() called with wrong action, want %s, got %s",
			DiscussionEventAnsweredAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		DiscussionEventAnsweredAction,
		DiscussionEventAnyAction,
	} {
		if _, ok := g.onDiscussionEvent[action]; ok {
			for _, h := range g.onDiscussionEvent[action] {
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

// OnDiscussionEventUnanswered registers callbacks listening to events of type github.DiscussionEvent and action 'unanswered'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnDiscussionEventUnanswered must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#discussion
func (g *EventHandler) OnDiscussionEventUnanswered(callbacks ...DiscussionEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onDiscussionEvent == nil {
		g.onDiscussionEvent = make(map[string][]DiscussionEventHandleFunc)
	}
	g.onDiscussionEvent[DiscussionEventUnansweredAction] = append(
		g.onDiscussionEvent[DiscussionEventUnansweredAction],
		callbacks...,
	)
}

// SetOnDiscussionEventUnanswered registers callbacks listening to events of type github.DiscussionEvent and action 'unanswered'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnDiscussionEventUnansweredAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#discussion
func (g *EventHandler) SetOnDiscussionEventUnanswered(callbacks ...DiscussionEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onDiscussionEvent == nil {
		g.onDiscussionEvent = make(map[string][]DiscussionEventHandleFunc)
	}
	g.onDiscussionEvent[DiscussionEventUnansweredAction] = callbacks
}

func (g *EventHandler) handleDiscussionEventUnanswered(deliveryID string, eventName string, event *github.DiscussionEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if DiscussionEventUnansweredAction != *event.Action {
		return fmt.Errorf(
			"handleDiscussionEventUnanswered() called with wrong action, want %s, got %s",
			DiscussionEventUnansweredAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		DiscussionEventUnansweredAction,
		DiscussionEventAnyAction,
	} {
		if _, ok := g.onDiscussionEvent[action]; ok {
			for _, h := range g.onDiscussionEvent[action] {
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

// OnDiscussionEventLabeled registers callbacks listening to events of type github.DiscussionEvent and action 'labeled'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnDiscussionEventLabeled must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#discussion
func (g *EventHandler) OnDiscussionEventLabeled(callbacks ...DiscussionEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onDiscussionEvent == nil {
		g.onDiscussionEvent = make(map[string][]DiscussionEventHandleFunc)
	}
	g.onDiscussionEvent[DiscussionEventLabeledAction] = append(
		g.onDiscussionEvent[DiscussionEventLabeledAction],
		callbacks...,
	)
}

// SetOnDiscussionEventLabeled registers callbacks listening to events of type github.DiscussionEvent and action 'labeled'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnDiscussionEventLabeledAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#discussion
func (g *EventHandler) SetOnDiscussionEventLabeled(callbacks ...DiscussionEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onDiscussionEvent == nil {
		g.onDiscussionEvent = make(map[string][]DiscussionEventHandleFunc)
	}
	g.onDiscussionEvent[DiscussionEventLabeledAction] = callbacks
}

func (g *EventHandler) handleDiscussionEventLabeled(deliveryID string, eventName string, event *github.DiscussionEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if DiscussionEventLabeledAction != *event.Action {
		return fmt.Errorf(
			"handleDiscussionEventLabeled() called with wrong action, want %s, got %s",
			DiscussionEventLabeledAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		DiscussionEventLabeledAction,
		DiscussionEventAnyAction,
	} {
		if _, ok := g.onDiscussionEvent[action]; ok {
			for _, h := range g.onDiscussionEvent[action] {
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

// OnDiscussionEventUnlabeled registers callbacks listening to events of type github.DiscussionEvent and action 'unlabeled'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnDiscussionEventUnlabeled must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#discussion
func (g *EventHandler) OnDiscussionEventUnlabeled(callbacks ...DiscussionEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onDiscussionEvent == nil {
		g.onDiscussionEvent = make(map[string][]DiscussionEventHandleFunc)
	}
	g.onDiscussionEvent[DiscussionEventUnlabeledAction] = append(
		g.onDiscussionEvent[DiscussionEventUnlabeledAction],
		callbacks...,
	)
}

// SetOnDiscussionEventUnlabeled registers callbacks listening to events of type github.DiscussionEvent and action 'unlabeled'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnDiscussionEventUnlabeledAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#discussion
func (g *EventHandler) SetOnDiscussionEventUnlabeled(callbacks ...DiscussionEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onDiscussionEvent == nil {
		g.onDiscussionEvent = make(map[string][]DiscussionEventHandleFunc)
	}
	g.onDiscussionEvent[DiscussionEventUnlabeledAction] = callbacks
}

func (g *EventHandler) handleDiscussionEventUnlabeled(deliveryID string, eventName string, event *github.DiscussionEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if DiscussionEventUnlabeledAction != *event.Action {
		return fmt.Errorf(
			"handleDiscussionEventUnlabeled() called with wrong action, want %s, got %s",
			DiscussionEventUnlabeledAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		DiscussionEventUnlabeledAction,
		DiscussionEventAnyAction,
	} {
		if _, ok := g.onDiscussionEvent[action]; ok {
			for _, h := range g.onDiscussionEvent[action] {
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

// OnDiscussionEventAny registers callbacks listening to any events of type github.DiscussionEvent
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnDiscussionEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#discussion
func (g *EventHandler) OnDiscussionEventAny(callbacks ...DiscussionEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onDiscussionEvent == nil {
		g.onDiscussionEvent = make(map[string][]DiscussionEventHandleFunc)
	}
	g.onDiscussionEvent[DiscussionEventAnyAction] = append(
		g.onDiscussionEvent[DiscussionEventAnyAction],
		callbacks...,
	)
}

// SetOnDiscussionEventAny registers callbacks listening to any events of type github.DiscussionEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnDiscussionEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#discussion
func (g *EventHandler) SetOnDiscussionEventAny(callbacks ...DiscussionEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onDiscussionEvent == nil {
		g.onDiscussionEvent = make(map[string][]DiscussionEventHandleFunc)
	}
	g.onDiscussionEvent[DiscussionEventAnyAction] = callbacks
}

func (g *EventHandler) handleDiscussionEventAny(deliveryID string, eventName string, event *github.DiscussionEvent) error {
	if event == nil {
		return fmt.Errorf("event was empty or nil")
	}
	if _, ok := g.onDiscussionEvent[DiscussionEventAnyAction]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onDiscussionEvent[DiscussionEventAnyAction] {
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

// DiscussionEvent handles github.DiscussionEvent.
//
// Callbacks are executed in the following order:
//
// 1) All callbacks registered with OnBeforeAny are executed in parallel.
// 2) All callbacks registered with OnDiscussionEvent... are executed in parallel in case the Event has actions.
// 3) All callbacks registered with OnAfterAny are executed in parallel.
//
// on any error all callbacks registered with OnError are executed in parallel.
func (g *EventHandler) DiscussionEvent(deliveryID string, eventName string, event *github.DiscussionEvent) error {

	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	action := *event.Action

	err := g.handleBeforeAny(deliveryID, eventName, event)
	if err != nil {
		return g.handleError(deliveryID, eventName, event, err)
	}

	switch action {

	case DiscussionEventCreatedAction:
		err := g.handleDiscussionEventCreated(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case DiscussionEventEditedAction:
		err := g.handleDiscussionEventEdited(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case DiscussionEventDeletedAction:
		err := g.handleDiscussionEventDeleted(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case DiscussionEventPinnedAction:
		err := g.handleDiscussionEventPinned(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case DiscussionEventUnpinnedAction:
		err := g.handleDiscussionEventUnpinned(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case DiscussionEventLockedAction:
		err := g.handleDiscussionEventLocked(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case DiscussionEventUnlockedAction:
		err := g.handleDiscussionEventUnlocked(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case DiscussionEventTransferredAction:
		err := g.handleDiscussionEventTransferred(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case DiscussionEventCategoryChangedAction:
		err := g.handleDiscussionEventCategoryChanged(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case DiscussionEventAnsweredAction:
		err := g.handleDiscussionEventAnswered(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case DiscussionEventUnansweredAction:
		err := g.handleDiscussionEventUnanswered(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case DiscussionEventLabeledAction:
		err := g.handleDiscussionEventLabeled(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case DiscussionEventUnlabeledAction:
		err := g.handleDiscussionEventUnlabeled(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	default:
		err := g.handleDiscussionEventAny(deliveryID, eventName, event)
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
