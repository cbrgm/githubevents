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
	// CheckSuiteEvent is the event name of github.CheckSuiteEvent's
	CheckSuiteEvent = "check_suite"

	// CheckSuiteEventAnyAction is used to identify callbacks
	// listening to all events of type github.CheckSuiteEvent
	CheckSuiteEventAnyAction = "*"

	// CheckSuiteEventCompletedAction is used to identify callbacks
	// listening to events of type github.CheckSuiteEvent and action "completed"
	CheckSuiteEventCompletedAction = "completed"

	// CheckSuiteEventRequestedAction is used to identify callbacks
	// listening to events of type github.CheckSuiteEvent and action "requested"
	CheckSuiteEventRequestedAction = "requested"

	// CheckSuiteEventReRequestedAction is used to identify callbacks
	// listening to events of type github.CheckSuiteEvent and action "rerequested"
	CheckSuiteEventReRequestedAction = "rerequested"
)

// CheckSuiteEventHandleFunc represents a callback function triggered on github.CheckSuiteEvent's.
// 'deliveryID' (type: string) is the unique webhook delivery ID.
// 'eventName' (type: string) is the name of the event.
// 'event' (type: *github.CheckSuiteEvent) is the webhook payload.
type CheckSuiteEventHandleFunc func(deliveryID string, eventName string, event *github.CheckSuiteEvent) error

// OnCheckSuiteEventCompleted registers callbacks listening to events of type github.CheckSuiteEvent and action 'completed'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnCheckSuiteEventCompleted must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#check_suite
func (g *EventHandler) OnCheckSuiteEventCompleted(callbacks ...CheckSuiteEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onCheckSuiteEvent == nil {
		g.onCheckSuiteEvent = make(map[string][]CheckSuiteEventHandleFunc)
	}
	g.onCheckSuiteEvent[CheckSuiteEventCompletedAction] = append(
		g.onCheckSuiteEvent[CheckSuiteEventCompletedAction],
		callbacks...,
	)
}

// SetOnCheckSuiteEventCompleted registers callbacks listening to events of type github.CheckSuiteEvent and action 'completed'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnCheckSuiteEventCompletedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#check_suite
func (g *EventHandler) SetOnCheckSuiteEventCompleted(callbacks ...CheckSuiteEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onCheckSuiteEvent == nil {
		g.onCheckSuiteEvent = make(map[string][]CheckSuiteEventHandleFunc)
	}
	g.onCheckSuiteEvent[CheckSuiteEventCompletedAction] = callbacks
}

func (g *EventHandler) handleCheckSuiteEventCompleted(deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if CheckSuiteEventCompletedAction != *event.Action {
		return fmt.Errorf(
			"handleCheckSuiteEventCompleted() called with wrong action, want %s, got %s",
			CheckSuiteEventCompletedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		CheckSuiteEventCompletedAction,
		CheckSuiteEventAnyAction,
	} {
		if _, ok := g.onCheckSuiteEvent[action]; ok {
			for _, h := range g.onCheckSuiteEvent[action] {
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

// OnCheckSuiteEventRequested registers callbacks listening to events of type github.CheckSuiteEvent and action 'requested'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnCheckSuiteEventRequested must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#check_suite
func (g *EventHandler) OnCheckSuiteEventRequested(callbacks ...CheckSuiteEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onCheckSuiteEvent == nil {
		g.onCheckSuiteEvent = make(map[string][]CheckSuiteEventHandleFunc)
	}
	g.onCheckSuiteEvent[CheckSuiteEventRequestedAction] = append(
		g.onCheckSuiteEvent[CheckSuiteEventRequestedAction],
		callbacks...,
	)
}

// SetOnCheckSuiteEventRequested registers callbacks listening to events of type github.CheckSuiteEvent and action 'requested'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnCheckSuiteEventRequestedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#check_suite
func (g *EventHandler) SetOnCheckSuiteEventRequested(callbacks ...CheckSuiteEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onCheckSuiteEvent == nil {
		g.onCheckSuiteEvent = make(map[string][]CheckSuiteEventHandleFunc)
	}
	g.onCheckSuiteEvent[CheckSuiteEventRequestedAction] = callbacks
}

func (g *EventHandler) handleCheckSuiteEventRequested(deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if CheckSuiteEventRequestedAction != *event.Action {
		return fmt.Errorf(
			"handleCheckSuiteEventRequested() called with wrong action, want %s, got %s",
			CheckSuiteEventRequestedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		CheckSuiteEventRequestedAction,
		CheckSuiteEventAnyAction,
	} {
		if _, ok := g.onCheckSuiteEvent[action]; ok {
			for _, h := range g.onCheckSuiteEvent[action] {
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

// OnCheckSuiteEventReRequested registers callbacks listening to events of type github.CheckSuiteEvent and action 'rerequested'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnCheckSuiteEventReRequested must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#check_suite
func (g *EventHandler) OnCheckSuiteEventReRequested(callbacks ...CheckSuiteEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onCheckSuiteEvent == nil {
		g.onCheckSuiteEvent = make(map[string][]CheckSuiteEventHandleFunc)
	}
	g.onCheckSuiteEvent[CheckSuiteEventReRequestedAction] = append(
		g.onCheckSuiteEvent[CheckSuiteEventReRequestedAction],
		callbacks...,
	)
}

// SetOnCheckSuiteEventReRequested registers callbacks listening to events of type github.CheckSuiteEvent and action 'rerequested'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnCheckSuiteEventReRequestedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#check_suite
func (g *EventHandler) SetOnCheckSuiteEventReRequested(callbacks ...CheckSuiteEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onCheckSuiteEvent == nil {
		g.onCheckSuiteEvent = make(map[string][]CheckSuiteEventHandleFunc)
	}
	g.onCheckSuiteEvent[CheckSuiteEventReRequestedAction] = callbacks
}

func (g *EventHandler) handleCheckSuiteEventReRequested(deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if CheckSuiteEventReRequestedAction != *event.Action {
		return fmt.Errorf(
			"handleCheckSuiteEventReRequested() called with wrong action, want %s, got %s",
			CheckSuiteEventReRequestedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		CheckSuiteEventReRequestedAction,
		CheckSuiteEventAnyAction,
	} {
		if _, ok := g.onCheckSuiteEvent[action]; ok {
			for _, h := range g.onCheckSuiteEvent[action] {
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

// OnCheckSuiteEventAny registers callbacks listening to any events of type github.CheckSuiteEvent
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnCheckSuiteEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#check_suite
func (g *EventHandler) OnCheckSuiteEventAny(callbacks ...CheckSuiteEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onCheckSuiteEvent == nil {
		g.onCheckSuiteEvent = make(map[string][]CheckSuiteEventHandleFunc)
	}
	g.onCheckSuiteEvent[CheckSuiteEventAnyAction] = append(
		g.onCheckSuiteEvent[CheckSuiteEventAnyAction],
		callbacks...,
	)
}

// SetOnCheckSuiteEventAny registers callbacks listening to any events of type github.CheckSuiteEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnCheckSuiteEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#check_suite
func (g *EventHandler) SetOnCheckSuiteEventAny(callbacks ...CheckSuiteEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onCheckSuiteEvent == nil {
		g.onCheckSuiteEvent = make(map[string][]CheckSuiteEventHandleFunc)
	}
	g.onCheckSuiteEvent[CheckSuiteEventAnyAction] = callbacks
}

func (g *EventHandler) handleCheckSuiteEventAny(deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
	if event == nil {
		return fmt.Errorf("event was empty or nil")
	}
	if _, ok := g.onCheckSuiteEvent[CheckSuiteEventAnyAction]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onCheckSuiteEvent[CheckSuiteEventAnyAction] {
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

// CheckSuiteEvent handles github.CheckSuiteEvent.
//
// Callbacks are executed in the following order:
//
// 1) All callbacks registered with OnBeforeAny are executed in parallel.
// 2) All callbacks registered with OnCheckSuiteEvent... are executed in parallel in case the Event has actions.
// 3) All callbacks registered with OnAfterAny are executed in parallel.
//
// on any error all callbacks registered with OnError are executed in parallel.
func (g *EventHandler) CheckSuiteEvent(deliveryID string, eventName string, event *github.CheckSuiteEvent) error {

	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	action := *event.Action

	err := g.handleBeforeAny(deliveryID, eventName, event)
	if err != nil {
		return g.handleError(deliveryID, eventName, event, err)
	}

	switch action {

	case CheckSuiteEventCompletedAction:
		err := g.handleCheckSuiteEventCompleted(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case CheckSuiteEventRequestedAction:
		err := g.handleCheckSuiteEventRequested(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case CheckSuiteEventReRequestedAction:
		err := g.handleCheckSuiteEventReRequested(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	default:
		err := g.handleCheckSuiteEventAny(deliveryID, eventName, event)
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
