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
	// OrgBlockEvent is the event name of github.OrgBlockEvent's
	OrgBlockEvent = "org_block"

	// OrgBlockEventAnyAction is used to identify callbacks
	// listening to all events of type github.OrgBlockEvent
	OrgBlockEventAnyAction = "*"

	// OrgBlockEventBlockedAction is used to identify callbacks
	// listening to events of type github.OrgBlockEvent and action "blocked"
	OrgBlockEventBlockedAction = "blocked"

	// OrgBlockEventUnblockedAction is used to identify callbacks
	// listening to events of type github.OrgBlockEvent and action "unblocked"
	OrgBlockEventUnblockedAction = "unblocked"
)

// OrgBlockEventHandleFunc represents a callback function triggered on github.OrgBlockEvent's.
// 'deliveryID' (type: string) is the unique webhook delivery ID.
// 'eventName' (type: string) is the name of the event.
// 'event' (type: *github.OrgBlockEvent) is the webhook payload.
type OrgBlockEventHandleFunc func(deliveryID string, eventName string, event *github.OrgBlockEvent) error

// OnOrgBlockEventBlocked registers callbacks listening to events of type github.OrgBlockEvent and action 'blocked'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnOrgBlockEventBlocked must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#org_block
func (g *EventHandler) OnOrgBlockEventBlocked(callbacks ...OrgBlockEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onOrgBlockEvent == nil {
		g.onOrgBlockEvent = make(map[string][]OrgBlockEventHandleFunc)
	}
	g.onOrgBlockEvent[OrgBlockEventBlockedAction] = append(
		g.onOrgBlockEvent[OrgBlockEventBlockedAction],
		callbacks...,
	)
}

// SetOnOrgBlockEventBlocked registers callbacks listening to events of type github.OrgBlockEvent and action 'blocked'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnOrgBlockEventBlockedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#org_block
func (g *EventHandler) SetOnOrgBlockEventBlocked(callbacks ...OrgBlockEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onOrgBlockEvent == nil {
		g.onOrgBlockEvent = make(map[string][]OrgBlockEventHandleFunc)
	}
	g.onOrgBlockEvent[OrgBlockEventBlockedAction] = callbacks
}

func (g *EventHandler) handleOrgBlockEventBlocked(deliveryID string, eventName string, event *github.OrgBlockEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if OrgBlockEventBlockedAction != *event.Action {
		return fmt.Errorf(
			"handleOrgBlockEventBlocked() called with wrong action, want %s, got %s",
			OrgBlockEventBlockedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		OrgBlockEventBlockedAction,
		OrgBlockEventAnyAction,
	} {
		if _, ok := g.onOrgBlockEvent[action]; ok {
			for _, h := range g.onOrgBlockEvent[action] {
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

// OnOrgBlockEventUnblocked registers callbacks listening to events of type github.OrgBlockEvent and action 'unblocked'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnOrgBlockEventUnblocked must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#org_block
func (g *EventHandler) OnOrgBlockEventUnblocked(callbacks ...OrgBlockEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onOrgBlockEvent == nil {
		g.onOrgBlockEvent = make(map[string][]OrgBlockEventHandleFunc)
	}
	g.onOrgBlockEvent[OrgBlockEventUnblockedAction] = append(
		g.onOrgBlockEvent[OrgBlockEventUnblockedAction],
		callbacks...,
	)
}

// SetOnOrgBlockEventUnblocked registers callbacks listening to events of type github.OrgBlockEvent and action 'unblocked'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnOrgBlockEventUnblockedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#org_block
func (g *EventHandler) SetOnOrgBlockEventUnblocked(callbacks ...OrgBlockEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onOrgBlockEvent == nil {
		g.onOrgBlockEvent = make(map[string][]OrgBlockEventHandleFunc)
	}
	g.onOrgBlockEvent[OrgBlockEventUnblockedAction] = callbacks
}

func (g *EventHandler) handleOrgBlockEventUnblocked(deliveryID string, eventName string, event *github.OrgBlockEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if OrgBlockEventUnblockedAction != *event.Action {
		return fmt.Errorf(
			"handleOrgBlockEventUnblocked() called with wrong action, want %s, got %s",
			OrgBlockEventUnblockedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		OrgBlockEventUnblockedAction,
		OrgBlockEventAnyAction,
	} {
		if _, ok := g.onOrgBlockEvent[action]; ok {
			for _, h := range g.onOrgBlockEvent[action] {
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

// OnOrgBlockEventAny registers callbacks listening to any events of type github.OrgBlockEvent
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnOrgBlockEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#org_block
func (g *EventHandler) OnOrgBlockEventAny(callbacks ...OrgBlockEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onOrgBlockEvent == nil {
		g.onOrgBlockEvent = make(map[string][]OrgBlockEventHandleFunc)
	}
	g.onOrgBlockEvent[OrgBlockEventAnyAction] = append(
		g.onOrgBlockEvent[OrgBlockEventAnyAction],
		callbacks...,
	)
}

// SetOnOrgBlockEventAny registers callbacks listening to any events of type github.OrgBlockEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnOrgBlockEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#org_block
func (g *EventHandler) SetOnOrgBlockEventAny(callbacks ...OrgBlockEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onOrgBlockEvent == nil {
		g.onOrgBlockEvent = make(map[string][]OrgBlockEventHandleFunc)
	}
	g.onOrgBlockEvent[OrgBlockEventAnyAction] = callbacks
}

func (g *EventHandler) handleOrgBlockEventAny(deliveryID string, eventName string, event *github.OrgBlockEvent) error {
	if event == nil {
		return fmt.Errorf("event was empty or nil")
	}
	if _, ok := g.onOrgBlockEvent[OrgBlockEventAnyAction]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onOrgBlockEvent[OrgBlockEventAnyAction] {
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

// OrgBlockEvent handles github.OrgBlockEvent.
//
// Callbacks are executed in the following order:
//
// 1) All callbacks registered with OnBeforeAny are executed in parallel.
// 2) All callbacks registered with OnOrgBlockEvent... are executed in parallel in case the Event has actions.
// 3) All callbacks registered with OnAfterAny are executed in parallel.
//
// on any error all callbacks registered with OnError are executed in parallel.
func (g *EventHandler) OrgBlockEvent(deliveryID string, eventName string, event *github.OrgBlockEvent) error {

	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	action := *event.Action

	err := g.handleBeforeAny(deliveryID, eventName, event)
	if err != nil {
		return g.handleError(deliveryID, eventName, event, err)
	}

	switch action {

	case OrgBlockEventBlockedAction:
		err := g.handleOrgBlockEventBlocked(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case OrgBlockEventUnblockedAction:
		err := g.handleOrgBlockEventUnblocked(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	default:
		err := g.handleOrgBlockEventAny(deliveryID, eventName, event)
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
