// Copyright 2022 The GithubEvents Authors. All rights reserved.
// Use of this source code is governed by the MIT License
// that can be found in the LICENSE file.

package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"fmt"
	"github.com/google/go-github/v60/github"
	"golang.org/x/sync/errgroup"
)

// Actions are used to identify registered callbacks.
const (
	// MergeGroupEvent is the event name of github.MergeGroupEvent's
	MergeGroupEvent = "merge_group_event"

	// MergeGroupEventAnyAction is used to identify callbacks
	// listening to all events of type github.MergeGroupEvent
	MergeGroupEventAnyAction = "*"

	// MergeGroupEventChecksRequestedAction is used to identify callbacks
	// listening to events of type github.MergeGroupEvent and action "checks_requested"
	MergeGroupEventChecksRequestedAction = "checks_requested"

	// MergeGroupEventDestroyedAction is used to identify callbacks
	// listening to events of type github.MergeGroupEvent and action "destroyed"
	MergeGroupEventDestroyedAction = "destroyed"
)

// MergeGroupEventHandleFunc represents a callback function triggered on github.MergeGroupEvent's.
// 'deliveryID' (type: string) is the unique webhook delivery ID.
// 'eventName' (type: string) is the name of the event.
// 'event' (type: *github.MergeGroupEvent) is the webhook payload.
type MergeGroupEventHandleFunc func(deliveryID string, eventName string, event *github.MergeGroupEvent) error

// OnMergeGroupEventChecksRequested registers callbacks listening to events of type github.MergeGroupEvent and action 'checks_requested'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnMergeGroupEventChecksRequested must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#merge_group_event
func (g *EventHandler) OnMergeGroupEventChecksRequested(callbacks ...MergeGroupEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onMergeGroupEvent == nil {
		g.onMergeGroupEvent = make(map[string][]MergeGroupEventHandleFunc)
	}
	g.onMergeGroupEvent[MergeGroupEventChecksRequestedAction] = append(
		g.onMergeGroupEvent[MergeGroupEventChecksRequestedAction],
		callbacks...,
	)
}

// SetOnMergeGroupEventChecksRequested registers callbacks listening to events of type github.MergeGroupEvent and action 'checks_requested'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnMergeGroupEventChecksRequestedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#merge_group_event
func (g *EventHandler) SetOnMergeGroupEventChecksRequested(callbacks ...MergeGroupEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onMergeGroupEvent == nil {
		g.onMergeGroupEvent = make(map[string][]MergeGroupEventHandleFunc)
	}
	g.onMergeGroupEvent[MergeGroupEventChecksRequestedAction] = callbacks
}

func (g *EventHandler) handleMergeGroupEventChecksRequested(deliveryID string, eventName string, event *github.MergeGroupEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if MergeGroupEventChecksRequestedAction != *event.Action {
		return fmt.Errorf(
			"handleMergeGroupEventChecksRequested() called with wrong action, want %s, got %s",
			MergeGroupEventChecksRequestedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		MergeGroupEventChecksRequestedAction,
		MergeGroupEventAnyAction,
	} {
		if _, ok := g.onMergeGroupEvent[action]; ok {
			for _, h := range g.onMergeGroupEvent[action] {
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

// OnMergeGroupEventDestroyed registers callbacks listening to events of type github.MergeGroupEvent and action 'destroyed'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnMergeGroupEventDestroyed must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#merge_group_event
func (g *EventHandler) OnMergeGroupEventDestroyed(callbacks ...MergeGroupEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onMergeGroupEvent == nil {
		g.onMergeGroupEvent = make(map[string][]MergeGroupEventHandleFunc)
	}
	g.onMergeGroupEvent[MergeGroupEventDestroyedAction] = append(
		g.onMergeGroupEvent[MergeGroupEventDestroyedAction],
		callbacks...,
	)
}

// SetOnMergeGroupEventDestroyed registers callbacks listening to events of type github.MergeGroupEvent and action 'destroyed'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnMergeGroupEventDestroyedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#merge_group_event
func (g *EventHandler) SetOnMergeGroupEventDestroyed(callbacks ...MergeGroupEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onMergeGroupEvent == nil {
		g.onMergeGroupEvent = make(map[string][]MergeGroupEventHandleFunc)
	}
	g.onMergeGroupEvent[MergeGroupEventDestroyedAction] = callbacks
}

func (g *EventHandler) handleMergeGroupEventDestroyed(deliveryID string, eventName string, event *github.MergeGroupEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if MergeGroupEventDestroyedAction != *event.Action {
		return fmt.Errorf(
			"handleMergeGroupEventDestroyed() called with wrong action, want %s, got %s",
			MergeGroupEventDestroyedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		MergeGroupEventDestroyedAction,
		MergeGroupEventAnyAction,
	} {
		if _, ok := g.onMergeGroupEvent[action]; ok {
			for _, h := range g.onMergeGroupEvent[action] {
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

// OnMergeGroupEventAny registers callbacks listening to any events of type github.MergeGroupEvent
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnMergeGroupEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#merge_group_event
func (g *EventHandler) OnMergeGroupEventAny(callbacks ...MergeGroupEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onMergeGroupEvent == nil {
		g.onMergeGroupEvent = make(map[string][]MergeGroupEventHandleFunc)
	}
	g.onMergeGroupEvent[MergeGroupEventAnyAction] = append(
		g.onMergeGroupEvent[MergeGroupEventAnyAction],
		callbacks...,
	)
}

// SetOnMergeGroupEventAny registers callbacks listening to any events of type github.MergeGroupEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnMergeGroupEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#merge_group_event
func (g *EventHandler) SetOnMergeGroupEventAny(callbacks ...MergeGroupEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onMergeGroupEvent == nil {
		g.onMergeGroupEvent = make(map[string][]MergeGroupEventHandleFunc)
	}
	g.onMergeGroupEvent[MergeGroupEventAnyAction] = callbacks
}

func (g *EventHandler) handleMergeGroupEventAny(deliveryID string, eventName string, event *github.MergeGroupEvent) error {
	if event == nil {
		return fmt.Errorf("event was empty or nil")
	}
	if _, ok := g.onMergeGroupEvent[MergeGroupEventAnyAction]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onMergeGroupEvent[MergeGroupEventAnyAction] {
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

// MergeGroupEvent handles github.MergeGroupEvent.
//
// Callbacks are executed in the following order:
//
// 1) All callbacks registered with OnBeforeAny are executed in parallel.
// 2) All callbacks registered with OnMergeGroupEvent... are executed in parallel in case the Event has actions.
// 3) All callbacks registered with OnAfterAny are executed in parallel.
//
// on any error all callbacks registered with OnError are executed in parallel.
func (g *EventHandler) MergeGroupEvent(deliveryID string, eventName string, event *github.MergeGroupEvent) error {

	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	action := *event.Action

	err := g.handleBeforeAny(deliveryID, eventName, event)
	if err != nil {
		return g.handleError(deliveryID, eventName, event, err)
	}

	switch action {

	case MergeGroupEventChecksRequestedAction:
		err := g.handleMergeGroupEventChecksRequested(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case MergeGroupEventDestroyedAction:
		err := g.handleMergeGroupEventDestroyed(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	default:
		err := g.handleMergeGroupEventAny(deliveryID, eventName, event)
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
