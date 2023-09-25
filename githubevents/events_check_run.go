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
	// CheckRunEvent is the event name of github.CheckRunEvent's
	CheckRunEvent = "check_run"

	// CheckRunEventAnyAction is used to identify callbacks
	// listening to all events of type github.CheckRunEvent
	CheckRunEventAnyAction = "*"

	// CheckRunEventCreatedAction is used to identify callbacks
	// listening to events of type github.CheckRunEvent and action "created"
	CheckRunEventCreatedAction = "created"

	// CheckRunEventCompletedAction is used to identify callbacks
	// listening to events of type github.CheckRunEvent and action "completed"
	CheckRunEventCompletedAction = "completed"

	// CheckRunEventReRequestedAction is used to identify callbacks
	// listening to events of type github.CheckRunEvent and action "rerequested"
	CheckRunEventReRequestedAction = "rerequested"

	// CheckRunEventRequestActionAction is used to identify callbacks
	// listening to events of type github.CheckRunEvent and action "requested_action"
	CheckRunEventRequestActionAction = "requested_action"
)

// CheckRunEventHandleFunc represents a callback function triggered on github.CheckRunEvent's.
// 'deliveryID' (type: string) is the unique webhook delivery ID.
// 'eventName' (type: string) is the name of the event.
// 'event' (type: *github.CheckRunEvent) is the webhook payload.
type CheckRunEventHandleFunc func(deliveryID string, eventName string, event *github.CheckRunEvent) error

// OnCheckRunEventCreated registers callbacks listening to events of type github.CheckRunEvent and action 'created'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnCheckRunEventCreated must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#check_run
func (g *EventHandler) OnCheckRunEventCreated(callbacks ...CheckRunEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onCheckRunEvent == nil {
		g.onCheckRunEvent = make(map[string][]CheckRunEventHandleFunc)
	}
	g.onCheckRunEvent[CheckRunEventCreatedAction] = append(
		g.onCheckRunEvent[CheckRunEventCreatedAction],
		callbacks...,
	)
}

// SetOnCheckRunEventCreated registers callbacks listening to events of type github.CheckRunEvent and action 'created'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnCheckRunEventCreatedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#check_run
func (g *EventHandler) SetOnCheckRunEventCreated(callbacks ...CheckRunEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onCheckRunEvent == nil {
		g.onCheckRunEvent = make(map[string][]CheckRunEventHandleFunc)
	}
	g.onCheckRunEvent[CheckRunEventCreatedAction] = callbacks
}

func (g *EventHandler) handleCheckRunEventCreated(deliveryID string, eventName string, event *github.CheckRunEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if CheckRunEventCreatedAction != *event.Action {
		return fmt.Errorf(
			"handleCheckRunEventCreated() called with wrong action, want %s, got %s",
			CheckRunEventCreatedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		CheckRunEventCreatedAction,
		CheckRunEventAnyAction,
	} {
		if _, ok := g.onCheckRunEvent[action]; ok {
			for _, h := range g.onCheckRunEvent[action] {
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

// OnCheckRunEventCompleted registers callbacks listening to events of type github.CheckRunEvent and action 'completed'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnCheckRunEventCompleted must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#check_run
func (g *EventHandler) OnCheckRunEventCompleted(callbacks ...CheckRunEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onCheckRunEvent == nil {
		g.onCheckRunEvent = make(map[string][]CheckRunEventHandleFunc)
	}
	g.onCheckRunEvent[CheckRunEventCompletedAction] = append(
		g.onCheckRunEvent[CheckRunEventCompletedAction],
		callbacks...,
	)
}

// SetOnCheckRunEventCompleted registers callbacks listening to events of type github.CheckRunEvent and action 'completed'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnCheckRunEventCompletedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#check_run
func (g *EventHandler) SetOnCheckRunEventCompleted(callbacks ...CheckRunEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onCheckRunEvent == nil {
		g.onCheckRunEvent = make(map[string][]CheckRunEventHandleFunc)
	}
	g.onCheckRunEvent[CheckRunEventCompletedAction] = callbacks
}

func (g *EventHandler) handleCheckRunEventCompleted(deliveryID string, eventName string, event *github.CheckRunEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if CheckRunEventCompletedAction != *event.Action {
		return fmt.Errorf(
			"handleCheckRunEventCompleted() called with wrong action, want %s, got %s",
			CheckRunEventCompletedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		CheckRunEventCompletedAction,
		CheckRunEventAnyAction,
	} {
		if _, ok := g.onCheckRunEvent[action]; ok {
			for _, h := range g.onCheckRunEvent[action] {
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

// OnCheckRunEventReRequested registers callbacks listening to events of type github.CheckRunEvent and action 'rerequested'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnCheckRunEventReRequested must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#check_run
func (g *EventHandler) OnCheckRunEventReRequested(callbacks ...CheckRunEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onCheckRunEvent == nil {
		g.onCheckRunEvent = make(map[string][]CheckRunEventHandleFunc)
	}
	g.onCheckRunEvent[CheckRunEventReRequestedAction] = append(
		g.onCheckRunEvent[CheckRunEventReRequestedAction],
		callbacks...,
	)
}

// SetOnCheckRunEventReRequested registers callbacks listening to events of type github.CheckRunEvent and action 'rerequested'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnCheckRunEventReRequestedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#check_run
func (g *EventHandler) SetOnCheckRunEventReRequested(callbacks ...CheckRunEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onCheckRunEvent == nil {
		g.onCheckRunEvent = make(map[string][]CheckRunEventHandleFunc)
	}
	g.onCheckRunEvent[CheckRunEventReRequestedAction] = callbacks
}

func (g *EventHandler) handleCheckRunEventReRequested(deliveryID string, eventName string, event *github.CheckRunEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if CheckRunEventReRequestedAction != *event.Action {
		return fmt.Errorf(
			"handleCheckRunEventReRequested() called with wrong action, want %s, got %s",
			CheckRunEventReRequestedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		CheckRunEventReRequestedAction,
		CheckRunEventAnyAction,
	} {
		if _, ok := g.onCheckRunEvent[action]; ok {
			for _, h := range g.onCheckRunEvent[action] {
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

// OnCheckRunEventRequestAction registers callbacks listening to events of type github.CheckRunEvent and action 'requested_action'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnCheckRunEventRequestAction must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#check_run
func (g *EventHandler) OnCheckRunEventRequestAction(callbacks ...CheckRunEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onCheckRunEvent == nil {
		g.onCheckRunEvent = make(map[string][]CheckRunEventHandleFunc)
	}
	g.onCheckRunEvent[CheckRunEventRequestActionAction] = append(
		g.onCheckRunEvent[CheckRunEventRequestActionAction],
		callbacks...,
	)
}

// SetOnCheckRunEventRequestAction registers callbacks listening to events of type github.CheckRunEvent and action 'requested_action'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnCheckRunEventRequestActionAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#check_run
func (g *EventHandler) SetOnCheckRunEventRequestAction(callbacks ...CheckRunEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onCheckRunEvent == nil {
		g.onCheckRunEvent = make(map[string][]CheckRunEventHandleFunc)
	}
	g.onCheckRunEvent[CheckRunEventRequestActionAction] = callbacks
}

func (g *EventHandler) handleCheckRunEventRequestAction(deliveryID string, eventName string, event *github.CheckRunEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if CheckRunEventRequestActionAction != *event.Action {
		return fmt.Errorf(
			"handleCheckRunEventRequestAction() called with wrong action, want %s, got %s",
			CheckRunEventRequestActionAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		CheckRunEventRequestActionAction,
		CheckRunEventAnyAction,
	} {
		if _, ok := g.onCheckRunEvent[action]; ok {
			for _, h := range g.onCheckRunEvent[action] {
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

// OnCheckRunEventAny registers callbacks listening to any events of type github.CheckRunEvent
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnCheckRunEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#check_run
func (g *EventHandler) OnCheckRunEventAny(callbacks ...CheckRunEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onCheckRunEvent == nil {
		g.onCheckRunEvent = make(map[string][]CheckRunEventHandleFunc)
	}
	g.onCheckRunEvent[CheckRunEventAnyAction] = append(
		g.onCheckRunEvent[CheckRunEventAnyAction],
		callbacks...,
	)
}

// SetOnCheckRunEventAny registers callbacks listening to any events of type github.CheckRunEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnCheckRunEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#check_run
func (g *EventHandler) SetOnCheckRunEventAny(callbacks ...CheckRunEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onCheckRunEvent == nil {
		g.onCheckRunEvent = make(map[string][]CheckRunEventHandleFunc)
	}
	g.onCheckRunEvent[CheckRunEventAnyAction] = callbacks
}

func (g *EventHandler) handleCheckRunEventAny(deliveryID string, eventName string, event *github.CheckRunEvent) error {
	if event == nil {
		return fmt.Errorf("event was empty or nil")
	}
	if _, ok := g.onCheckRunEvent[CheckRunEventAnyAction]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onCheckRunEvent[CheckRunEventAnyAction] {
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

// CheckRunEvent handles github.CheckRunEvent.
//
// Callbacks are executed in the following order:
//
// 1) All callbacks registered with OnBeforeAny are executed in parallel.
// 2) All callbacks registered with OnCheckRunEvent... are executed in parallel in case the Event has actions.
// 3) All callbacks registered with OnAfterAny are executed in parallel.
//
// on any error all callbacks registered with OnError are executed in parallel.
func (g *EventHandler) CheckRunEvent(deliveryID string, eventName string, event *github.CheckRunEvent) error {

	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	action := *event.Action

	err := g.handleBeforeAny(deliveryID, eventName, event)
	if err != nil {
		return g.handleError(deliveryID, eventName, event, err)
	}

	switch action {

	case CheckRunEventCreatedAction:
		err := g.handleCheckRunEventCreated(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case CheckRunEventCompletedAction:
		err := g.handleCheckRunEventCompleted(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case CheckRunEventReRequestedAction:
		err := g.handleCheckRunEventReRequested(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case CheckRunEventRequestActionAction:
		err := g.handleCheckRunEventRequestAction(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	default:
		err := g.handleCheckRunEventAny(deliveryID, eventName, event)
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
