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
	// LabelEvent is the event name of github.LabelEvent's
	LabelEvent = "label"

	// LabelEventAnyAction is used to identify callbacks
	// listening to all events of type github.LabelEvent
	LabelEventAnyAction = "*"

	// LabelEventCreatedAction is used to identify callbacks
	// listening to events of type github.LabelEvent and action "created"
	LabelEventCreatedAction = "created"

	// LabelEventEditedAction is used to identify callbacks
	// listening to events of type github.LabelEvent and action "edited"
	LabelEventEditedAction = "edited"

	// LabelEventDeletedAction is used to identify callbacks
	// listening to events of type github.LabelEvent and action "deleted"
	LabelEventDeletedAction = "deleted"
)

// LabelEventHandleFunc represents a callback function triggered on github.LabelEvent's.
// 'deliveryID' (type: string) is the unique webhook delivery ID.
// 'eventName' (type: string) is the name of the event.
// 'event' (type: *github.LabelEvent) is the webhook payload.
type LabelEventHandleFunc func(deliveryID string, eventName string, event *github.LabelEvent) error

// OnLabelEventCreated registers callbacks listening to events of type github.LabelEvent and action 'created'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnLabelEventCreated must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#label
func (g *EventHandler) OnLabelEventCreated(callbacks ...LabelEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onLabelEvent == nil {
		g.onLabelEvent = make(map[string][]LabelEventHandleFunc)
	}
	g.onLabelEvent[LabelEventCreatedAction] = append(
		g.onLabelEvent[LabelEventCreatedAction],
		callbacks...,
	)
}

// SetOnLabelEventCreated registers callbacks listening to events of type github.LabelEvent and action 'created'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnLabelEventCreatedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#label
func (g *EventHandler) SetOnLabelEventCreated(callbacks ...LabelEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onLabelEvent == nil {
		g.onLabelEvent = make(map[string][]LabelEventHandleFunc)
	}
	g.onLabelEvent[LabelEventCreatedAction] = callbacks
}

func (g *EventHandler) handleLabelEventCreated(deliveryID string, eventName string, event *github.LabelEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if LabelEventCreatedAction != *event.Action {
		return fmt.Errorf(
			"handleLabelEventCreated() called with wrong action, want %s, got %s",
			LabelEventCreatedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		LabelEventCreatedAction,
		LabelEventAnyAction,
	} {
		if _, ok := g.onLabelEvent[action]; ok {
			for _, h := range g.onLabelEvent[action] {
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

// OnLabelEventEdited registers callbacks listening to events of type github.LabelEvent and action 'edited'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnLabelEventEdited must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#label
func (g *EventHandler) OnLabelEventEdited(callbacks ...LabelEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onLabelEvent == nil {
		g.onLabelEvent = make(map[string][]LabelEventHandleFunc)
	}
	g.onLabelEvent[LabelEventEditedAction] = append(
		g.onLabelEvent[LabelEventEditedAction],
		callbacks...,
	)
}

// SetOnLabelEventEdited registers callbacks listening to events of type github.LabelEvent and action 'edited'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnLabelEventEditedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#label
func (g *EventHandler) SetOnLabelEventEdited(callbacks ...LabelEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onLabelEvent == nil {
		g.onLabelEvent = make(map[string][]LabelEventHandleFunc)
	}
	g.onLabelEvent[LabelEventEditedAction] = callbacks
}

func (g *EventHandler) handleLabelEventEdited(deliveryID string, eventName string, event *github.LabelEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if LabelEventEditedAction != *event.Action {
		return fmt.Errorf(
			"handleLabelEventEdited() called with wrong action, want %s, got %s",
			LabelEventEditedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		LabelEventEditedAction,
		LabelEventAnyAction,
	} {
		if _, ok := g.onLabelEvent[action]; ok {
			for _, h := range g.onLabelEvent[action] {
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

// OnLabelEventDeleted registers callbacks listening to events of type github.LabelEvent and action 'deleted'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnLabelEventDeleted must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#label
func (g *EventHandler) OnLabelEventDeleted(callbacks ...LabelEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onLabelEvent == nil {
		g.onLabelEvent = make(map[string][]LabelEventHandleFunc)
	}
	g.onLabelEvent[LabelEventDeletedAction] = append(
		g.onLabelEvent[LabelEventDeletedAction],
		callbacks...,
	)
}

// SetOnLabelEventDeleted registers callbacks listening to events of type github.LabelEvent and action 'deleted'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnLabelEventDeletedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#label
func (g *EventHandler) SetOnLabelEventDeleted(callbacks ...LabelEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onLabelEvent == nil {
		g.onLabelEvent = make(map[string][]LabelEventHandleFunc)
	}
	g.onLabelEvent[LabelEventDeletedAction] = callbacks
}

func (g *EventHandler) handleLabelEventDeleted(deliveryID string, eventName string, event *github.LabelEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if LabelEventDeletedAction != *event.Action {
		return fmt.Errorf(
			"handleLabelEventDeleted() called with wrong action, want %s, got %s",
			LabelEventDeletedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		LabelEventDeletedAction,
		LabelEventAnyAction,
	} {
		if _, ok := g.onLabelEvent[action]; ok {
			for _, h := range g.onLabelEvent[action] {
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

// OnLabelEventAny registers callbacks listening to any events of type github.LabelEvent
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnLabelEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#label
func (g *EventHandler) OnLabelEventAny(callbacks ...LabelEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onLabelEvent == nil {
		g.onLabelEvent = make(map[string][]LabelEventHandleFunc)
	}
	g.onLabelEvent[LabelEventAnyAction] = append(
		g.onLabelEvent[LabelEventAnyAction],
		callbacks...,
	)
}

// SetOnLabelEventAny registers callbacks listening to any events of type github.LabelEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnLabelEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#label
func (g *EventHandler) SetOnLabelEventAny(callbacks ...LabelEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onLabelEvent == nil {
		g.onLabelEvent = make(map[string][]LabelEventHandleFunc)
	}
	g.onLabelEvent[LabelEventAnyAction] = callbacks
}

func (g *EventHandler) handleLabelEventAny(deliveryID string, eventName string, event *github.LabelEvent) error {
	if event == nil {
		return fmt.Errorf("event was empty or nil")
	}
	if _, ok := g.onLabelEvent[LabelEventAnyAction]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onLabelEvent[LabelEventAnyAction] {
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

// LabelEvent handles github.LabelEvent.
//
// Callbacks are executed in the following order:
//
// 1) All callbacks registered with OnBeforeAny are executed in parallel.
// 2) All callbacks registered with OnLabelEvent... are executed in parallel in case the Event has actions.
// 3) All callbacks registered with OnAfterAny are executed in parallel.
//
// on any error all callbacks registered with OnError are executed in parallel.
func (g *EventHandler) LabelEvent(deliveryID string, eventName string, event *github.LabelEvent) error {

	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	action := *event.Action

	err := g.handleBeforeAny(deliveryID, eventName, event)
	if err != nil {
		return g.handleError(deliveryID, eventName, event, err)
	}

	switch action {

	case LabelEventCreatedAction:
		err := g.handleLabelEventCreated(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case LabelEventEditedAction:
		err := g.handleLabelEventEdited(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case LabelEventDeletedAction:
		err := g.handleLabelEventDeleted(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	default:
		err := g.handleLabelEventAny(deliveryID, eventName, event)
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
