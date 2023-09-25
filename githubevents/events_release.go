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
	// ReleaseEvent is the event name of github.ReleaseEvent's
	ReleaseEvent = "release"

	// ReleaseEventAnyAction is used to identify callbacks
	// listening to all events of type github.ReleaseEvent
	ReleaseEventAnyAction = "*"

	// ReleaseEventPublishedAction is used to identify callbacks
	// listening to events of type github.ReleaseEvent and action "published"
	ReleaseEventPublishedAction = "published"

	// ReleaseEventUnpublishedAction is used to identify callbacks
	// listening to events of type github.ReleaseEvent and action "unpublished"
	ReleaseEventUnpublishedAction = "unpublished"

	// ReleaseEventCreatedAction is used to identify callbacks
	// listening to events of type github.ReleaseEvent and action "created"
	ReleaseEventCreatedAction = "created"

	// ReleaseEventEditedAction is used to identify callbacks
	// listening to events of type github.ReleaseEvent and action "edited"
	ReleaseEventEditedAction = "edited"

	// ReleaseEventDeletedAction is used to identify callbacks
	// listening to events of type github.ReleaseEvent and action "deleted"
	ReleaseEventDeletedAction = "deleted"

	// ReleaseEventPreReleasedAction is used to identify callbacks
	// listening to events of type github.ReleaseEvent and action "prereleased"
	ReleaseEventPreReleasedAction = "prereleased"

	// ReleaseEventReleasedAction is used to identify callbacks
	// listening to events of type github.ReleaseEvent and action "released"
	ReleaseEventReleasedAction = "released"
)

// ReleaseEventHandleFunc represents a callback function triggered on github.ReleaseEvent's.
// 'deliveryID' (type: string) is the unique webhook delivery ID.
// 'eventName' (type: string) is the name of the event.
// 'event' (type: *github.ReleaseEvent) is the webhook payload.
type ReleaseEventHandleFunc func(deliveryID string, eventName string, event *github.ReleaseEvent) error

// OnReleaseEventPublished registers callbacks listening to events of type github.ReleaseEvent and action 'published'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnReleaseEventPublished must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#release
func (g *EventHandler) OnReleaseEventPublished(callbacks ...ReleaseEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onReleaseEvent == nil {
		g.onReleaseEvent = make(map[string][]ReleaseEventHandleFunc)
	}
	g.onReleaseEvent[ReleaseEventPublishedAction] = append(
		g.onReleaseEvent[ReleaseEventPublishedAction],
		callbacks...,
	)
}

// SetOnReleaseEventPublished registers callbacks listening to events of type github.ReleaseEvent and action 'published'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnReleaseEventPublishedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#release
func (g *EventHandler) SetOnReleaseEventPublished(callbacks ...ReleaseEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onReleaseEvent == nil {
		g.onReleaseEvent = make(map[string][]ReleaseEventHandleFunc)
	}
	g.onReleaseEvent[ReleaseEventPublishedAction] = callbacks
}

func (g *EventHandler) handleReleaseEventPublished(deliveryID string, eventName string, event *github.ReleaseEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if ReleaseEventPublishedAction != *event.Action {
		return fmt.Errorf(
			"handleReleaseEventPublished() called with wrong action, want %s, got %s",
			ReleaseEventPublishedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		ReleaseEventPublishedAction,
		ReleaseEventAnyAction,
	} {
		if _, ok := g.onReleaseEvent[action]; ok {
			for _, h := range g.onReleaseEvent[action] {
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

// OnReleaseEventUnpublished registers callbacks listening to events of type github.ReleaseEvent and action 'unpublished'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnReleaseEventUnpublished must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#release
func (g *EventHandler) OnReleaseEventUnpublished(callbacks ...ReleaseEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onReleaseEvent == nil {
		g.onReleaseEvent = make(map[string][]ReleaseEventHandleFunc)
	}
	g.onReleaseEvent[ReleaseEventUnpublishedAction] = append(
		g.onReleaseEvent[ReleaseEventUnpublishedAction],
		callbacks...,
	)
}

// SetOnReleaseEventUnpublished registers callbacks listening to events of type github.ReleaseEvent and action 'unpublished'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnReleaseEventUnpublishedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#release
func (g *EventHandler) SetOnReleaseEventUnpublished(callbacks ...ReleaseEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onReleaseEvent == nil {
		g.onReleaseEvent = make(map[string][]ReleaseEventHandleFunc)
	}
	g.onReleaseEvent[ReleaseEventUnpublishedAction] = callbacks
}

func (g *EventHandler) handleReleaseEventUnpublished(deliveryID string, eventName string, event *github.ReleaseEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if ReleaseEventUnpublishedAction != *event.Action {
		return fmt.Errorf(
			"handleReleaseEventUnpublished() called with wrong action, want %s, got %s",
			ReleaseEventUnpublishedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		ReleaseEventUnpublishedAction,
		ReleaseEventAnyAction,
	} {
		if _, ok := g.onReleaseEvent[action]; ok {
			for _, h := range g.onReleaseEvent[action] {
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

// OnReleaseEventCreated registers callbacks listening to events of type github.ReleaseEvent and action 'created'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnReleaseEventCreated must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#release
func (g *EventHandler) OnReleaseEventCreated(callbacks ...ReleaseEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onReleaseEvent == nil {
		g.onReleaseEvent = make(map[string][]ReleaseEventHandleFunc)
	}
	g.onReleaseEvent[ReleaseEventCreatedAction] = append(
		g.onReleaseEvent[ReleaseEventCreatedAction],
		callbacks...,
	)
}

// SetOnReleaseEventCreated registers callbacks listening to events of type github.ReleaseEvent and action 'created'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnReleaseEventCreatedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#release
func (g *EventHandler) SetOnReleaseEventCreated(callbacks ...ReleaseEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onReleaseEvent == nil {
		g.onReleaseEvent = make(map[string][]ReleaseEventHandleFunc)
	}
	g.onReleaseEvent[ReleaseEventCreatedAction] = callbacks
}

func (g *EventHandler) handleReleaseEventCreated(deliveryID string, eventName string, event *github.ReleaseEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if ReleaseEventCreatedAction != *event.Action {
		return fmt.Errorf(
			"handleReleaseEventCreated() called with wrong action, want %s, got %s",
			ReleaseEventCreatedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		ReleaseEventCreatedAction,
		ReleaseEventAnyAction,
	} {
		if _, ok := g.onReleaseEvent[action]; ok {
			for _, h := range g.onReleaseEvent[action] {
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

// OnReleaseEventEdited registers callbacks listening to events of type github.ReleaseEvent and action 'edited'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnReleaseEventEdited must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#release
func (g *EventHandler) OnReleaseEventEdited(callbacks ...ReleaseEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onReleaseEvent == nil {
		g.onReleaseEvent = make(map[string][]ReleaseEventHandleFunc)
	}
	g.onReleaseEvent[ReleaseEventEditedAction] = append(
		g.onReleaseEvent[ReleaseEventEditedAction],
		callbacks...,
	)
}

// SetOnReleaseEventEdited registers callbacks listening to events of type github.ReleaseEvent and action 'edited'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnReleaseEventEditedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#release
func (g *EventHandler) SetOnReleaseEventEdited(callbacks ...ReleaseEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onReleaseEvent == nil {
		g.onReleaseEvent = make(map[string][]ReleaseEventHandleFunc)
	}
	g.onReleaseEvent[ReleaseEventEditedAction] = callbacks
}

func (g *EventHandler) handleReleaseEventEdited(deliveryID string, eventName string, event *github.ReleaseEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if ReleaseEventEditedAction != *event.Action {
		return fmt.Errorf(
			"handleReleaseEventEdited() called with wrong action, want %s, got %s",
			ReleaseEventEditedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		ReleaseEventEditedAction,
		ReleaseEventAnyAction,
	} {
		if _, ok := g.onReleaseEvent[action]; ok {
			for _, h := range g.onReleaseEvent[action] {
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

// OnReleaseEventDeleted registers callbacks listening to events of type github.ReleaseEvent and action 'deleted'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnReleaseEventDeleted must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#release
func (g *EventHandler) OnReleaseEventDeleted(callbacks ...ReleaseEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onReleaseEvent == nil {
		g.onReleaseEvent = make(map[string][]ReleaseEventHandleFunc)
	}
	g.onReleaseEvent[ReleaseEventDeletedAction] = append(
		g.onReleaseEvent[ReleaseEventDeletedAction],
		callbacks...,
	)
}

// SetOnReleaseEventDeleted registers callbacks listening to events of type github.ReleaseEvent and action 'deleted'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnReleaseEventDeletedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#release
func (g *EventHandler) SetOnReleaseEventDeleted(callbacks ...ReleaseEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onReleaseEvent == nil {
		g.onReleaseEvent = make(map[string][]ReleaseEventHandleFunc)
	}
	g.onReleaseEvent[ReleaseEventDeletedAction] = callbacks
}

func (g *EventHandler) handleReleaseEventDeleted(deliveryID string, eventName string, event *github.ReleaseEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if ReleaseEventDeletedAction != *event.Action {
		return fmt.Errorf(
			"handleReleaseEventDeleted() called with wrong action, want %s, got %s",
			ReleaseEventDeletedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		ReleaseEventDeletedAction,
		ReleaseEventAnyAction,
	} {
		if _, ok := g.onReleaseEvent[action]; ok {
			for _, h := range g.onReleaseEvent[action] {
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

// OnReleaseEventPreReleased registers callbacks listening to events of type github.ReleaseEvent and action 'prereleased'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnReleaseEventPreReleased must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#release
func (g *EventHandler) OnReleaseEventPreReleased(callbacks ...ReleaseEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onReleaseEvent == nil {
		g.onReleaseEvent = make(map[string][]ReleaseEventHandleFunc)
	}
	g.onReleaseEvent[ReleaseEventPreReleasedAction] = append(
		g.onReleaseEvent[ReleaseEventPreReleasedAction],
		callbacks...,
	)
}

// SetOnReleaseEventPreReleased registers callbacks listening to events of type github.ReleaseEvent and action 'prereleased'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnReleaseEventPreReleasedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#release
func (g *EventHandler) SetOnReleaseEventPreReleased(callbacks ...ReleaseEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onReleaseEvent == nil {
		g.onReleaseEvent = make(map[string][]ReleaseEventHandleFunc)
	}
	g.onReleaseEvent[ReleaseEventPreReleasedAction] = callbacks
}

func (g *EventHandler) handleReleaseEventPreReleased(deliveryID string, eventName string, event *github.ReleaseEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if ReleaseEventPreReleasedAction != *event.Action {
		return fmt.Errorf(
			"handleReleaseEventPreReleased() called with wrong action, want %s, got %s",
			ReleaseEventPreReleasedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		ReleaseEventPreReleasedAction,
		ReleaseEventAnyAction,
	} {
		if _, ok := g.onReleaseEvent[action]; ok {
			for _, h := range g.onReleaseEvent[action] {
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

// OnReleaseEventReleased registers callbacks listening to events of type github.ReleaseEvent and action 'released'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnReleaseEventReleased must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#release
func (g *EventHandler) OnReleaseEventReleased(callbacks ...ReleaseEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onReleaseEvent == nil {
		g.onReleaseEvent = make(map[string][]ReleaseEventHandleFunc)
	}
	g.onReleaseEvent[ReleaseEventReleasedAction] = append(
		g.onReleaseEvent[ReleaseEventReleasedAction],
		callbacks...,
	)
}

// SetOnReleaseEventReleased registers callbacks listening to events of type github.ReleaseEvent and action 'released'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnReleaseEventReleasedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#release
func (g *EventHandler) SetOnReleaseEventReleased(callbacks ...ReleaseEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onReleaseEvent == nil {
		g.onReleaseEvent = make(map[string][]ReleaseEventHandleFunc)
	}
	g.onReleaseEvent[ReleaseEventReleasedAction] = callbacks
}

func (g *EventHandler) handleReleaseEventReleased(deliveryID string, eventName string, event *github.ReleaseEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if ReleaseEventReleasedAction != *event.Action {
		return fmt.Errorf(
			"handleReleaseEventReleased() called with wrong action, want %s, got %s",
			ReleaseEventReleasedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		ReleaseEventReleasedAction,
		ReleaseEventAnyAction,
	} {
		if _, ok := g.onReleaseEvent[action]; ok {
			for _, h := range g.onReleaseEvent[action] {
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

// OnReleaseEventAny registers callbacks listening to any events of type github.ReleaseEvent
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnReleaseEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#release
func (g *EventHandler) OnReleaseEventAny(callbacks ...ReleaseEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onReleaseEvent == nil {
		g.onReleaseEvent = make(map[string][]ReleaseEventHandleFunc)
	}
	g.onReleaseEvent[ReleaseEventAnyAction] = append(
		g.onReleaseEvent[ReleaseEventAnyAction],
		callbacks...,
	)
}

// SetOnReleaseEventAny registers callbacks listening to any events of type github.ReleaseEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnReleaseEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#release
func (g *EventHandler) SetOnReleaseEventAny(callbacks ...ReleaseEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onReleaseEvent == nil {
		g.onReleaseEvent = make(map[string][]ReleaseEventHandleFunc)
	}
	g.onReleaseEvent[ReleaseEventAnyAction] = callbacks
}

func (g *EventHandler) handleReleaseEventAny(deliveryID string, eventName string, event *github.ReleaseEvent) error {
	if event == nil {
		return fmt.Errorf("event was empty or nil")
	}
	if _, ok := g.onReleaseEvent[ReleaseEventAnyAction]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onReleaseEvent[ReleaseEventAnyAction] {
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

// ReleaseEvent handles github.ReleaseEvent.
//
// Callbacks are executed in the following order:
//
// 1) All callbacks registered with OnBeforeAny are executed in parallel.
// 2) All callbacks registered with OnReleaseEvent... are executed in parallel in case the Event has actions.
// 3) All callbacks registered with OnAfterAny are executed in parallel.
//
// on any error all callbacks registered with OnError are executed in parallel.
func (g *EventHandler) ReleaseEvent(deliveryID string, eventName string, event *github.ReleaseEvent) error {

	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	action := *event.Action

	err := g.handleBeforeAny(deliveryID, eventName, event)
	if err != nil {
		return g.handleError(deliveryID, eventName, event, err)
	}

	switch action {

	case ReleaseEventPublishedAction:
		err := g.handleReleaseEventPublished(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case ReleaseEventUnpublishedAction:
		err := g.handleReleaseEventUnpublished(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case ReleaseEventCreatedAction:
		err := g.handleReleaseEventCreated(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case ReleaseEventEditedAction:
		err := g.handleReleaseEventEdited(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case ReleaseEventDeletedAction:
		err := g.handleReleaseEventDeleted(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case ReleaseEventPreReleasedAction:
		err := g.handleReleaseEventPreReleased(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case ReleaseEventReleasedAction:
		err := g.handleReleaseEventReleased(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	default:
		err := g.handleReleaseEventAny(deliveryID, eventName, event)
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
