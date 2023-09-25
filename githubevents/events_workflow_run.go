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
	// WorkflowRunEvent is the event name of github.WorkflowRunEvent's
	WorkflowRunEvent = "workflow_run"

	// WorkflowRunEventAnyAction is used to identify callbacks
	// listening to all events of type github.WorkflowRunEvent
	WorkflowRunEventAnyAction = "*"

	// WorkflowRunEventRequestedAction is used to identify callbacks
	// listening to events of type github.WorkflowRunEvent and action "requested"
	WorkflowRunEventRequestedAction = "requested"

	// WorkflowRunEventCompletedAction is used to identify callbacks
	// listening to events of type github.WorkflowRunEvent and action "completed"
	WorkflowRunEventCompletedAction = "completed"
)

// WorkflowRunEventHandleFunc represents a callback function triggered on github.WorkflowRunEvent's.
// 'deliveryID' (type: string) is the unique webhook delivery ID.
// 'eventName' (type: string) is the name of the event.
// 'event' (type: *github.WorkflowRunEvent) is the webhook payload.
type WorkflowRunEventHandleFunc func(deliveryID string, eventName string, event *github.WorkflowRunEvent) error

// OnWorkflowRunEventRequested registers callbacks listening to events of type github.WorkflowRunEvent and action 'requested'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnWorkflowRunEventRequested must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#workflow_run
func (g *EventHandler) OnWorkflowRunEventRequested(callbacks ...WorkflowRunEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onWorkflowRunEvent == nil {
		g.onWorkflowRunEvent = make(map[string][]WorkflowRunEventHandleFunc)
	}
	g.onWorkflowRunEvent[WorkflowRunEventRequestedAction] = append(
		g.onWorkflowRunEvent[WorkflowRunEventRequestedAction],
		callbacks...,
	)
}

// SetOnWorkflowRunEventRequested registers callbacks listening to events of type github.WorkflowRunEvent and action 'requested'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnWorkflowRunEventRequestedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#workflow_run
func (g *EventHandler) SetOnWorkflowRunEventRequested(callbacks ...WorkflowRunEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onWorkflowRunEvent == nil {
		g.onWorkflowRunEvent = make(map[string][]WorkflowRunEventHandleFunc)
	}
	g.onWorkflowRunEvent[WorkflowRunEventRequestedAction] = callbacks
}

func (g *EventHandler) handleWorkflowRunEventRequested(deliveryID string, eventName string, event *github.WorkflowRunEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if WorkflowRunEventRequestedAction != *event.Action {
		return fmt.Errorf(
			"handleWorkflowRunEventRequested() called with wrong action, want %s, got %s",
			WorkflowRunEventRequestedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		WorkflowRunEventRequestedAction,
		WorkflowRunEventAnyAction,
	} {
		if _, ok := g.onWorkflowRunEvent[action]; ok {
			for _, h := range g.onWorkflowRunEvent[action] {
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

// OnWorkflowRunEventCompleted registers callbacks listening to events of type github.WorkflowRunEvent and action 'completed'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnWorkflowRunEventCompleted must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#workflow_run
func (g *EventHandler) OnWorkflowRunEventCompleted(callbacks ...WorkflowRunEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onWorkflowRunEvent == nil {
		g.onWorkflowRunEvent = make(map[string][]WorkflowRunEventHandleFunc)
	}
	g.onWorkflowRunEvent[WorkflowRunEventCompletedAction] = append(
		g.onWorkflowRunEvent[WorkflowRunEventCompletedAction],
		callbacks...,
	)
}

// SetOnWorkflowRunEventCompleted registers callbacks listening to events of type github.WorkflowRunEvent and action 'completed'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnWorkflowRunEventCompletedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#workflow_run
func (g *EventHandler) SetOnWorkflowRunEventCompleted(callbacks ...WorkflowRunEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onWorkflowRunEvent == nil {
		g.onWorkflowRunEvent = make(map[string][]WorkflowRunEventHandleFunc)
	}
	g.onWorkflowRunEvent[WorkflowRunEventCompletedAction] = callbacks
}

func (g *EventHandler) handleWorkflowRunEventCompleted(deliveryID string, eventName string, event *github.WorkflowRunEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if WorkflowRunEventCompletedAction != *event.Action {
		return fmt.Errorf(
			"handleWorkflowRunEventCompleted() called with wrong action, want %s, got %s",
			WorkflowRunEventCompletedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		WorkflowRunEventCompletedAction,
		WorkflowRunEventAnyAction,
	} {
		if _, ok := g.onWorkflowRunEvent[action]; ok {
			for _, h := range g.onWorkflowRunEvent[action] {
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

// OnWorkflowRunEventAny registers callbacks listening to any events of type github.WorkflowRunEvent
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnWorkflowRunEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#workflow_run
func (g *EventHandler) OnWorkflowRunEventAny(callbacks ...WorkflowRunEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onWorkflowRunEvent == nil {
		g.onWorkflowRunEvent = make(map[string][]WorkflowRunEventHandleFunc)
	}
	g.onWorkflowRunEvent[WorkflowRunEventAnyAction] = append(
		g.onWorkflowRunEvent[WorkflowRunEventAnyAction],
		callbacks...,
	)
}

// SetOnWorkflowRunEventAny registers callbacks listening to any events of type github.WorkflowRunEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnWorkflowRunEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#workflow_run
func (g *EventHandler) SetOnWorkflowRunEventAny(callbacks ...WorkflowRunEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onWorkflowRunEvent == nil {
		g.onWorkflowRunEvent = make(map[string][]WorkflowRunEventHandleFunc)
	}
	g.onWorkflowRunEvent[WorkflowRunEventAnyAction] = callbacks
}

func (g *EventHandler) handleWorkflowRunEventAny(deliveryID string, eventName string, event *github.WorkflowRunEvent) error {
	if event == nil {
		return fmt.Errorf("event was empty or nil")
	}
	if _, ok := g.onWorkflowRunEvent[WorkflowRunEventAnyAction]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onWorkflowRunEvent[WorkflowRunEventAnyAction] {
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

// WorkflowRunEvent handles github.WorkflowRunEvent.
//
// Callbacks are executed in the following order:
//
// 1) All callbacks registered with OnBeforeAny are executed in parallel.
// 2) All callbacks registered with OnWorkflowRunEvent... are executed in parallel in case the Event has actions.
// 3) All callbacks registered with OnAfterAny are executed in parallel.
//
// on any error all callbacks registered with OnError are executed in parallel.
func (g *EventHandler) WorkflowRunEvent(deliveryID string, eventName string, event *github.WorkflowRunEvent) error {

	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	action := *event.Action

	err := g.handleBeforeAny(deliveryID, eventName, event)
	if err != nil {
		return g.handleError(deliveryID, eventName, event, err)
	}

	switch action {

	case WorkflowRunEventRequestedAction:
		err := g.handleWorkflowRunEventRequested(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case WorkflowRunEventCompletedAction:
		err := g.handleWorkflowRunEventCompleted(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	default:
		err := g.handleWorkflowRunEventAny(deliveryID, eventName, event)
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
