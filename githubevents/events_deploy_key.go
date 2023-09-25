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
	// DeployKeyEvent is the event name of github.DeployKeyEvent's
	DeployKeyEvent = "deploy_key"

	// DeployKeyEventAnyAction is used to identify callbacks
	// listening to all events of type github.DeployKeyEvent
	DeployKeyEventAnyAction = "*"

	// DeployKeyEventCreatedAction is used to identify callbacks
	// listening to events of type github.DeployKeyEvent and action "created"
	DeployKeyEventCreatedAction = "created"

	// DeployKeyEventDeletedAction is used to identify callbacks
	// listening to events of type github.DeployKeyEvent and action "deleted"
	DeployKeyEventDeletedAction = "deleted"
)

// DeployKeyEventHandleFunc represents a callback function triggered on github.DeployKeyEvent's.
// 'deliveryID' (type: string) is the unique webhook delivery ID.
// 'eventName' (type: string) is the name of the event.
// 'event' (type: *github.DeployKeyEvent) is the webhook payload.
type DeployKeyEventHandleFunc func(deliveryID string, eventName string, event *github.DeployKeyEvent) error

// OnDeployKeyEventCreated registers callbacks listening to events of type github.DeployKeyEvent and action 'created'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnDeployKeyEventCreated must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#deploy_key
func (g *EventHandler) OnDeployKeyEventCreated(callbacks ...DeployKeyEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onDeployKeyEvent == nil {
		g.onDeployKeyEvent = make(map[string][]DeployKeyEventHandleFunc)
	}
	g.onDeployKeyEvent[DeployKeyEventCreatedAction] = append(
		g.onDeployKeyEvent[DeployKeyEventCreatedAction],
		callbacks...,
	)
}

// SetOnDeployKeyEventCreated registers callbacks listening to events of type github.DeployKeyEvent and action 'created'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnDeployKeyEventCreatedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#deploy_key
func (g *EventHandler) SetOnDeployKeyEventCreated(callbacks ...DeployKeyEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onDeployKeyEvent == nil {
		g.onDeployKeyEvent = make(map[string][]DeployKeyEventHandleFunc)
	}
	g.onDeployKeyEvent[DeployKeyEventCreatedAction] = callbacks
}

func (g *EventHandler) handleDeployKeyEventCreated(deliveryID string, eventName string, event *github.DeployKeyEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if DeployKeyEventCreatedAction != *event.Action {
		return fmt.Errorf(
			"handleDeployKeyEventCreated() called with wrong action, want %s, got %s",
			DeployKeyEventCreatedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		DeployKeyEventCreatedAction,
		DeployKeyEventAnyAction,
	} {
		if _, ok := g.onDeployKeyEvent[action]; ok {
			for _, h := range g.onDeployKeyEvent[action] {
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

// OnDeployKeyEventDeleted registers callbacks listening to events of type github.DeployKeyEvent and action 'deleted'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnDeployKeyEventDeleted must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#deploy_key
func (g *EventHandler) OnDeployKeyEventDeleted(callbacks ...DeployKeyEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onDeployKeyEvent == nil {
		g.onDeployKeyEvent = make(map[string][]DeployKeyEventHandleFunc)
	}
	g.onDeployKeyEvent[DeployKeyEventDeletedAction] = append(
		g.onDeployKeyEvent[DeployKeyEventDeletedAction],
		callbacks...,
	)
}

// SetOnDeployKeyEventDeleted registers callbacks listening to events of type github.DeployKeyEvent and action 'deleted'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnDeployKeyEventDeletedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#deploy_key
func (g *EventHandler) SetOnDeployKeyEventDeleted(callbacks ...DeployKeyEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onDeployKeyEvent == nil {
		g.onDeployKeyEvent = make(map[string][]DeployKeyEventHandleFunc)
	}
	g.onDeployKeyEvent[DeployKeyEventDeletedAction] = callbacks
}

func (g *EventHandler) handleDeployKeyEventDeleted(deliveryID string, eventName string, event *github.DeployKeyEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if DeployKeyEventDeletedAction != *event.Action {
		return fmt.Errorf(
			"handleDeployKeyEventDeleted() called with wrong action, want %s, got %s",
			DeployKeyEventDeletedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		DeployKeyEventDeletedAction,
		DeployKeyEventAnyAction,
	} {
		if _, ok := g.onDeployKeyEvent[action]; ok {
			for _, h := range g.onDeployKeyEvent[action] {
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

// OnDeployKeyEventAny registers callbacks listening to any events of type github.DeployKeyEvent
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnDeployKeyEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#deploy_key
func (g *EventHandler) OnDeployKeyEventAny(callbacks ...DeployKeyEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onDeployKeyEvent == nil {
		g.onDeployKeyEvent = make(map[string][]DeployKeyEventHandleFunc)
	}
	g.onDeployKeyEvent[DeployKeyEventAnyAction] = append(
		g.onDeployKeyEvent[DeployKeyEventAnyAction],
		callbacks...,
	)
}

// SetOnDeployKeyEventAny registers callbacks listening to any events of type github.DeployKeyEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnDeployKeyEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#deploy_key
func (g *EventHandler) SetOnDeployKeyEventAny(callbacks ...DeployKeyEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onDeployKeyEvent == nil {
		g.onDeployKeyEvent = make(map[string][]DeployKeyEventHandleFunc)
	}
	g.onDeployKeyEvent[DeployKeyEventAnyAction] = callbacks
}

func (g *EventHandler) handleDeployKeyEventAny(deliveryID string, eventName string, event *github.DeployKeyEvent) error {
	if event == nil {
		return fmt.Errorf("event was empty or nil")
	}
	if _, ok := g.onDeployKeyEvent[DeployKeyEventAnyAction]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onDeployKeyEvent[DeployKeyEventAnyAction] {
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

// DeployKeyEvent handles github.DeployKeyEvent.
//
// Callbacks are executed in the following order:
//
// 1) All callbacks registered with OnBeforeAny are executed in parallel.
// 2) All callbacks registered with OnDeployKeyEvent... are executed in parallel in case the Event has actions.
// 3) All callbacks registered with OnAfterAny are executed in parallel.
//
// on any error all callbacks registered with OnError are executed in parallel.
func (g *EventHandler) DeployKeyEvent(deliveryID string, eventName string, event *github.DeployKeyEvent) error {

	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	action := *event.Action

	err := g.handleBeforeAny(deliveryID, eventName, event)
	if err != nil {
		return g.handleError(deliveryID, eventName, event, err)
	}

	switch action {

	case DeployKeyEventCreatedAction:
		err := g.handleDeployKeyEventCreated(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case DeployKeyEventDeletedAction:
		err := g.handleDeployKeyEventDeleted(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	default:
		err := g.handleDeployKeyEventAny(deliveryID, eventName, event)
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
