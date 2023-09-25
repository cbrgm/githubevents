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
	// MilestoneEvent is the event name of github.MilestoneEvent's
	MilestoneEvent = "milestone"

	// MilestoneEventAnyAction is used to identify callbacks
	// listening to all events of type github.MilestoneEvent
	MilestoneEventAnyAction = "*"

	// MilestoneEventCreatedAction is used to identify callbacks
	// listening to events of type github.MilestoneEvent and action "created"
	MilestoneEventCreatedAction = "created"

	// MilestoneEventClosedAction is used to identify callbacks
	// listening to events of type github.MilestoneEvent and action "closed"
	MilestoneEventClosedAction = "closed"

	// MilestoneEventOpenedAction is used to identify callbacks
	// listening to events of type github.MilestoneEvent and action "opened"
	MilestoneEventOpenedAction = "opened"

	// MilestoneEventEditedAction is used to identify callbacks
	// listening to events of type github.MilestoneEvent and action "edited"
	MilestoneEventEditedAction = "edited"

	// MilestoneEventDeletedAction is used to identify callbacks
	// listening to events of type github.MilestoneEvent and action "deleted"
	MilestoneEventDeletedAction = "deleted"
)

// MilestoneEventHandleFunc represents a callback function triggered on github.MilestoneEvent's.
// 'deliveryID' (type: string) is the unique webhook delivery ID.
// 'eventName' (type: string) is the name of the event.
// 'event' (type: *github.MilestoneEvent) is the webhook payload.
type MilestoneEventHandleFunc func(deliveryID string, eventName string, event *github.MilestoneEvent) error

// OnMilestoneEventCreated registers callbacks listening to events of type github.MilestoneEvent and action 'created'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnMilestoneEventCreated must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#milestone
func (g *EventHandler) OnMilestoneEventCreated(callbacks ...MilestoneEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onMilestoneEvent == nil {
		g.onMilestoneEvent = make(map[string][]MilestoneEventHandleFunc)
	}
	g.onMilestoneEvent[MilestoneEventCreatedAction] = append(
		g.onMilestoneEvent[MilestoneEventCreatedAction],
		callbacks...,
	)
}

// SetOnMilestoneEventCreated registers callbacks listening to events of type github.MilestoneEvent and action 'created'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnMilestoneEventCreatedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#milestone
func (g *EventHandler) SetOnMilestoneEventCreated(callbacks ...MilestoneEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onMilestoneEvent == nil {
		g.onMilestoneEvent = make(map[string][]MilestoneEventHandleFunc)
	}
	g.onMilestoneEvent[MilestoneEventCreatedAction] = callbacks
}

func (g *EventHandler) handleMilestoneEventCreated(deliveryID string, eventName string, event *github.MilestoneEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if MilestoneEventCreatedAction != *event.Action {
		return fmt.Errorf(
			"handleMilestoneEventCreated() called with wrong action, want %s, got %s",
			MilestoneEventCreatedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		MilestoneEventCreatedAction,
		MilestoneEventAnyAction,
	} {
		if _, ok := g.onMilestoneEvent[action]; ok {
			for _, h := range g.onMilestoneEvent[action] {
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

// OnMilestoneEventClosed registers callbacks listening to events of type github.MilestoneEvent and action 'closed'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnMilestoneEventClosed must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#milestone
func (g *EventHandler) OnMilestoneEventClosed(callbacks ...MilestoneEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onMilestoneEvent == nil {
		g.onMilestoneEvent = make(map[string][]MilestoneEventHandleFunc)
	}
	g.onMilestoneEvent[MilestoneEventClosedAction] = append(
		g.onMilestoneEvent[MilestoneEventClosedAction],
		callbacks...,
	)
}

// SetOnMilestoneEventClosed registers callbacks listening to events of type github.MilestoneEvent and action 'closed'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnMilestoneEventClosedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#milestone
func (g *EventHandler) SetOnMilestoneEventClosed(callbacks ...MilestoneEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onMilestoneEvent == nil {
		g.onMilestoneEvent = make(map[string][]MilestoneEventHandleFunc)
	}
	g.onMilestoneEvent[MilestoneEventClosedAction] = callbacks
}

func (g *EventHandler) handleMilestoneEventClosed(deliveryID string, eventName string, event *github.MilestoneEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if MilestoneEventClosedAction != *event.Action {
		return fmt.Errorf(
			"handleMilestoneEventClosed() called with wrong action, want %s, got %s",
			MilestoneEventClosedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		MilestoneEventClosedAction,
		MilestoneEventAnyAction,
	} {
		if _, ok := g.onMilestoneEvent[action]; ok {
			for _, h := range g.onMilestoneEvent[action] {
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

// OnMilestoneEventOpened registers callbacks listening to events of type github.MilestoneEvent and action 'opened'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnMilestoneEventOpened must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#milestone
func (g *EventHandler) OnMilestoneEventOpened(callbacks ...MilestoneEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onMilestoneEvent == nil {
		g.onMilestoneEvent = make(map[string][]MilestoneEventHandleFunc)
	}
	g.onMilestoneEvent[MilestoneEventOpenedAction] = append(
		g.onMilestoneEvent[MilestoneEventOpenedAction],
		callbacks...,
	)
}

// SetOnMilestoneEventOpened registers callbacks listening to events of type github.MilestoneEvent and action 'opened'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnMilestoneEventOpenedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#milestone
func (g *EventHandler) SetOnMilestoneEventOpened(callbacks ...MilestoneEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onMilestoneEvent == nil {
		g.onMilestoneEvent = make(map[string][]MilestoneEventHandleFunc)
	}
	g.onMilestoneEvent[MilestoneEventOpenedAction] = callbacks
}

func (g *EventHandler) handleMilestoneEventOpened(deliveryID string, eventName string, event *github.MilestoneEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if MilestoneEventOpenedAction != *event.Action {
		return fmt.Errorf(
			"handleMilestoneEventOpened() called with wrong action, want %s, got %s",
			MilestoneEventOpenedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		MilestoneEventOpenedAction,
		MilestoneEventAnyAction,
	} {
		if _, ok := g.onMilestoneEvent[action]; ok {
			for _, h := range g.onMilestoneEvent[action] {
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

// OnMilestoneEventEdited registers callbacks listening to events of type github.MilestoneEvent and action 'edited'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnMilestoneEventEdited must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#milestone
func (g *EventHandler) OnMilestoneEventEdited(callbacks ...MilestoneEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onMilestoneEvent == nil {
		g.onMilestoneEvent = make(map[string][]MilestoneEventHandleFunc)
	}
	g.onMilestoneEvent[MilestoneEventEditedAction] = append(
		g.onMilestoneEvent[MilestoneEventEditedAction],
		callbacks...,
	)
}

// SetOnMilestoneEventEdited registers callbacks listening to events of type github.MilestoneEvent and action 'edited'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnMilestoneEventEditedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#milestone
func (g *EventHandler) SetOnMilestoneEventEdited(callbacks ...MilestoneEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onMilestoneEvent == nil {
		g.onMilestoneEvent = make(map[string][]MilestoneEventHandleFunc)
	}
	g.onMilestoneEvent[MilestoneEventEditedAction] = callbacks
}

func (g *EventHandler) handleMilestoneEventEdited(deliveryID string, eventName string, event *github.MilestoneEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if MilestoneEventEditedAction != *event.Action {
		return fmt.Errorf(
			"handleMilestoneEventEdited() called with wrong action, want %s, got %s",
			MilestoneEventEditedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		MilestoneEventEditedAction,
		MilestoneEventAnyAction,
	} {
		if _, ok := g.onMilestoneEvent[action]; ok {
			for _, h := range g.onMilestoneEvent[action] {
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

// OnMilestoneEventDeleted registers callbacks listening to events of type github.MilestoneEvent and action 'deleted'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnMilestoneEventDeleted must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#milestone
func (g *EventHandler) OnMilestoneEventDeleted(callbacks ...MilestoneEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onMilestoneEvent == nil {
		g.onMilestoneEvent = make(map[string][]MilestoneEventHandleFunc)
	}
	g.onMilestoneEvent[MilestoneEventDeletedAction] = append(
		g.onMilestoneEvent[MilestoneEventDeletedAction],
		callbacks...,
	)
}

// SetOnMilestoneEventDeleted registers callbacks listening to events of type github.MilestoneEvent and action 'deleted'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnMilestoneEventDeletedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#milestone
func (g *EventHandler) SetOnMilestoneEventDeleted(callbacks ...MilestoneEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onMilestoneEvent == nil {
		g.onMilestoneEvent = make(map[string][]MilestoneEventHandleFunc)
	}
	g.onMilestoneEvent[MilestoneEventDeletedAction] = callbacks
}

func (g *EventHandler) handleMilestoneEventDeleted(deliveryID string, eventName string, event *github.MilestoneEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if MilestoneEventDeletedAction != *event.Action {
		return fmt.Errorf(
			"handleMilestoneEventDeleted() called with wrong action, want %s, got %s",
			MilestoneEventDeletedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		MilestoneEventDeletedAction,
		MilestoneEventAnyAction,
	} {
		if _, ok := g.onMilestoneEvent[action]; ok {
			for _, h := range g.onMilestoneEvent[action] {
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

// OnMilestoneEventAny registers callbacks listening to any events of type github.MilestoneEvent
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnMilestoneEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#milestone
func (g *EventHandler) OnMilestoneEventAny(callbacks ...MilestoneEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onMilestoneEvent == nil {
		g.onMilestoneEvent = make(map[string][]MilestoneEventHandleFunc)
	}
	g.onMilestoneEvent[MilestoneEventAnyAction] = append(
		g.onMilestoneEvent[MilestoneEventAnyAction],
		callbacks...,
	)
}

// SetOnMilestoneEventAny registers callbacks listening to any events of type github.MilestoneEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnMilestoneEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#milestone
func (g *EventHandler) SetOnMilestoneEventAny(callbacks ...MilestoneEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onMilestoneEvent == nil {
		g.onMilestoneEvent = make(map[string][]MilestoneEventHandleFunc)
	}
	g.onMilestoneEvent[MilestoneEventAnyAction] = callbacks
}

func (g *EventHandler) handleMilestoneEventAny(deliveryID string, eventName string, event *github.MilestoneEvent) error {
	if event == nil {
		return fmt.Errorf("event was empty or nil")
	}
	if _, ok := g.onMilestoneEvent[MilestoneEventAnyAction]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onMilestoneEvent[MilestoneEventAnyAction] {
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

// MilestoneEvent handles github.MilestoneEvent.
//
// Callbacks are executed in the following order:
//
// 1) All callbacks registered with OnBeforeAny are executed in parallel.
// 2) All callbacks registered with OnMilestoneEvent... are executed in parallel in case the Event has actions.
// 3) All callbacks registered with OnAfterAny are executed in parallel.
//
// on any error all callbacks registered with OnError are executed in parallel.
func (g *EventHandler) MilestoneEvent(deliveryID string, eventName string, event *github.MilestoneEvent) error {

	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	action := *event.Action

	err := g.handleBeforeAny(deliveryID, eventName, event)
	if err != nil {
		return g.handleError(deliveryID, eventName, event, err)
	}

	switch action {

	case MilestoneEventCreatedAction:
		err := g.handleMilestoneEventCreated(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case MilestoneEventClosedAction:
		err := g.handleMilestoneEventClosed(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case MilestoneEventOpenedAction:
		err := g.handleMilestoneEventOpened(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case MilestoneEventEditedAction:
		err := g.handleMilestoneEventEdited(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case MilestoneEventDeletedAction:
		err := g.handleMilestoneEventDeleted(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	default:
		err := g.handleMilestoneEventAny(deliveryID, eventName, event)
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
