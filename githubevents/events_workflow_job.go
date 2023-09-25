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
	// WorkflowJobEvent is the event name of github.WorkflowJobEvent's
	WorkflowJobEvent = "workflow_job"

	// WorkflowJobEventAnyAction is used to identify callbacks
	// listening to all events of type github.WorkflowJobEvent
	WorkflowJobEventAnyAction = "*"

	// WorkflowJobEventQueuedAction is used to identify callbacks
	// listening to events of type github.WorkflowJobEvent and action "queued"
	WorkflowJobEventQueuedAction = "queued"

	// WorkflowJobEventInProgressAction is used to identify callbacks
	// listening to events of type github.WorkflowJobEvent and action "in_progress"
	WorkflowJobEventInProgressAction = "in_progress"

	// WorkflowJobEventCompletedAction is used to identify callbacks
	// listening to events of type github.WorkflowJobEvent and action "completed"
	WorkflowJobEventCompletedAction = "completed"
)

// WorkflowJobEventHandleFunc represents a callback function triggered on github.WorkflowJobEvent's.
// 'deliveryID' (type: string) is the unique webhook delivery ID.
// 'eventName' (type: string) is the name of the event.
// 'event' (type: *github.WorkflowJobEvent) is the webhook payload.
type WorkflowJobEventHandleFunc func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error

// OnWorkflowJobEventQueued registers callbacks listening to events of type github.WorkflowJobEvent and action 'queued'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnWorkflowJobEventQueued must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#workflow_job
func (g *EventHandler) OnWorkflowJobEventQueued(callbacks ...WorkflowJobEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onWorkflowJobEvent == nil {
		g.onWorkflowJobEvent = make(map[string][]WorkflowJobEventHandleFunc)
	}
	g.onWorkflowJobEvent[WorkflowJobEventQueuedAction] = append(
		g.onWorkflowJobEvent[WorkflowJobEventQueuedAction],
		callbacks...,
	)
}

// SetOnWorkflowJobEventQueued registers callbacks listening to events of type github.WorkflowJobEvent and action 'queued'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnWorkflowJobEventQueuedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#workflow_job
func (g *EventHandler) SetOnWorkflowJobEventQueued(callbacks ...WorkflowJobEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onWorkflowJobEvent == nil {
		g.onWorkflowJobEvent = make(map[string][]WorkflowJobEventHandleFunc)
	}
	g.onWorkflowJobEvent[WorkflowJobEventQueuedAction] = callbacks
}

func (g *EventHandler) handleWorkflowJobEventQueued(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if WorkflowJobEventQueuedAction != *event.Action {
		return fmt.Errorf(
			"handleWorkflowJobEventQueued() called with wrong action, want %s, got %s",
			WorkflowJobEventQueuedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		WorkflowJobEventQueuedAction,
		WorkflowJobEventAnyAction,
	} {
		if _, ok := g.onWorkflowJobEvent[action]; ok {
			for _, h := range g.onWorkflowJobEvent[action] {
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

// OnWorkflowJobEventInProgress registers callbacks listening to events of type github.WorkflowJobEvent and action 'in_progress'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnWorkflowJobEventInProgress must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#workflow_job
func (g *EventHandler) OnWorkflowJobEventInProgress(callbacks ...WorkflowJobEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onWorkflowJobEvent == nil {
		g.onWorkflowJobEvent = make(map[string][]WorkflowJobEventHandleFunc)
	}
	g.onWorkflowJobEvent[WorkflowJobEventInProgressAction] = append(
		g.onWorkflowJobEvent[WorkflowJobEventInProgressAction],
		callbacks...,
	)
}

// SetOnWorkflowJobEventInProgress registers callbacks listening to events of type github.WorkflowJobEvent and action 'in_progress'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnWorkflowJobEventInProgressAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#workflow_job
func (g *EventHandler) SetOnWorkflowJobEventInProgress(callbacks ...WorkflowJobEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onWorkflowJobEvent == nil {
		g.onWorkflowJobEvent = make(map[string][]WorkflowJobEventHandleFunc)
	}
	g.onWorkflowJobEvent[WorkflowJobEventInProgressAction] = callbacks
}

func (g *EventHandler) handleWorkflowJobEventInProgress(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if WorkflowJobEventInProgressAction != *event.Action {
		return fmt.Errorf(
			"handleWorkflowJobEventInProgress() called with wrong action, want %s, got %s",
			WorkflowJobEventInProgressAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		WorkflowJobEventInProgressAction,
		WorkflowJobEventAnyAction,
	} {
		if _, ok := g.onWorkflowJobEvent[action]; ok {
			for _, h := range g.onWorkflowJobEvent[action] {
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

// OnWorkflowJobEventCompleted registers callbacks listening to events of type github.WorkflowJobEvent and action 'completed'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnWorkflowJobEventCompleted must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#workflow_job
func (g *EventHandler) OnWorkflowJobEventCompleted(callbacks ...WorkflowJobEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onWorkflowJobEvent == nil {
		g.onWorkflowJobEvent = make(map[string][]WorkflowJobEventHandleFunc)
	}
	g.onWorkflowJobEvent[WorkflowJobEventCompletedAction] = append(
		g.onWorkflowJobEvent[WorkflowJobEventCompletedAction],
		callbacks...,
	)
}

// SetOnWorkflowJobEventCompleted registers callbacks listening to events of type github.WorkflowJobEvent and action 'completed'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnWorkflowJobEventCompletedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#workflow_job
func (g *EventHandler) SetOnWorkflowJobEventCompleted(callbacks ...WorkflowJobEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onWorkflowJobEvent == nil {
		g.onWorkflowJobEvent = make(map[string][]WorkflowJobEventHandleFunc)
	}
	g.onWorkflowJobEvent[WorkflowJobEventCompletedAction] = callbacks
}

func (g *EventHandler) handleWorkflowJobEventCompleted(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if WorkflowJobEventCompletedAction != *event.Action {
		return fmt.Errorf(
			"handleWorkflowJobEventCompleted() called with wrong action, want %s, got %s",
			WorkflowJobEventCompletedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		WorkflowJobEventCompletedAction,
		WorkflowJobEventAnyAction,
	} {
		if _, ok := g.onWorkflowJobEvent[action]; ok {
			for _, h := range g.onWorkflowJobEvent[action] {
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

// OnWorkflowJobEventAny registers callbacks listening to any events of type github.WorkflowJobEvent
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnWorkflowJobEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#workflow_job
func (g *EventHandler) OnWorkflowJobEventAny(callbacks ...WorkflowJobEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onWorkflowJobEvent == nil {
		g.onWorkflowJobEvent = make(map[string][]WorkflowJobEventHandleFunc)
	}
	g.onWorkflowJobEvent[WorkflowJobEventAnyAction] = append(
		g.onWorkflowJobEvent[WorkflowJobEventAnyAction],
		callbacks...,
	)
}

// SetOnWorkflowJobEventAny registers callbacks listening to any events of type github.WorkflowJobEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnWorkflowJobEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#workflow_job
func (g *EventHandler) SetOnWorkflowJobEventAny(callbacks ...WorkflowJobEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onWorkflowJobEvent == nil {
		g.onWorkflowJobEvent = make(map[string][]WorkflowJobEventHandleFunc)
	}
	g.onWorkflowJobEvent[WorkflowJobEventAnyAction] = callbacks
}

func (g *EventHandler) handleWorkflowJobEventAny(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
	if event == nil {
		return fmt.Errorf("event was empty or nil")
	}
	if _, ok := g.onWorkflowJobEvent[WorkflowJobEventAnyAction]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onWorkflowJobEvent[WorkflowJobEventAnyAction] {
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

// WorkflowJobEvent handles github.WorkflowJobEvent.
//
// Callbacks are executed in the following order:
//
// 1) All callbacks registered with OnBeforeAny are executed in parallel.
// 2) All callbacks registered with OnWorkflowJobEvent... are executed in parallel in case the Event has actions.
// 3) All callbacks registered with OnAfterAny are executed in parallel.
//
// on any error all callbacks registered with OnError are executed in parallel.
func (g *EventHandler) WorkflowJobEvent(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {

	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	action := *event.Action

	err := g.handleBeforeAny(deliveryID, eventName, event)
	if err != nil {
		return g.handleError(deliveryID, eventName, event, err)
	}

	switch action {

	case WorkflowJobEventQueuedAction:
		err := g.handleWorkflowJobEventQueued(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case WorkflowJobEventInProgressAction:
		err := g.handleWorkflowJobEventInProgress(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case WorkflowJobEventCompletedAction:
		err := g.handleWorkflowJobEventCompleted(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	default:
		err := g.handleWorkflowJobEventAny(deliveryID, eventName, event)
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
