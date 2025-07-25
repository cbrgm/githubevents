// Code generated by gen/generate.go. DO NOT EDIT.
// make edits in gen/generate.go

// Copyright 2022 The GithubEvents Authors. All rights reserved.
// Use of this source code is governed by the MIT License
// that can be found in the LICENSE file.

package githubevents

import (
	"context"
	"fmt"
	"github.com/google/go-github/v74/github"
	"golang.org/x/sync/errgroup"
)

// Actions are used to identify registered callbacks.
const (
	// StatusEvent is the event name of github.StatusEvent's
	StatusEvent = "status"

	// StatusEventAnyAction is used to identify callbacks
	// listening to all events of type github.StatusEvent
	StatusEventAnyAction = "*"
)

// StatusEventHandleFunc represents a callback function triggered on github.StatusEvent's.
// 'deliveryID' (type: string) is the unique webhook delivery ID.
// 'eventName' (type: string) is the name of the event.
// 'event' (type: *github.StatusEvent) is the webhook payload.
type StatusEventHandleFunc func(ctx context.Context, deliveryID string, eventName string, event *github.StatusEvent) error

// OnStatusEventAny registers callbacks listening to any events of type github.StatusEvent
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnStatusEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#status
func (g *EventHandler) OnStatusEventAny(callbacks ...StatusEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onStatusEvent == nil {
		g.onStatusEvent = make(map[string][]StatusEventHandleFunc)
	}
	g.onStatusEvent[StatusEventAnyAction] = append(
		g.onStatusEvent[StatusEventAnyAction],
		callbacks...,
	)
}

// SetOnStatusEventAny registers callbacks listening to any events of type github.StatusEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnStatusEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#status
func (g *EventHandler) SetOnStatusEventAny(callbacks ...StatusEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onStatusEvent == nil {
		g.onStatusEvent = make(map[string][]StatusEventHandleFunc)
	}
	g.onStatusEvent[StatusEventAnyAction] = callbacks
}

func (g *EventHandler) handleStatusEventAny(ctx context.Context, deliveryID string, eventName string, event *github.StatusEvent) error {
	if event == nil {
		return fmt.Errorf("event was empty or nil")
	}
	if _, ok := g.onStatusEvent[StatusEventAnyAction]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onStatusEvent[StatusEventAnyAction] {
		handle := h
		eg.Go(func() (err error) {
			defer func() {
				if r := recover(); r != nil {
					err = fmt.Errorf("recovered from panic: %v", r)
				}
			}()
			err = handle(ctx, deliveryID, eventName, event)
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

// StatusEvent handles github.StatusEvent.
//
// Callbacks are executed in the following order:
//
// 1) All callbacks registered with OnBeforeAny are executed in parallel.
// 2) All callbacks registered with OnStatusEvent... are executed in parallel in case the Event has actions.
// 3) All callbacks registered with OnAfterAny are executed in parallel.
//
// on any error all callbacks registered with OnError are executed in parallel.
func (g *EventHandler) StatusEvent(ctx context.Context, deliveryID string, eventName string, event *github.StatusEvent) error {

	if event == nil {
		return fmt.Errorf("event action was empty or nil")
	}

	err := g.handleBeforeAny(ctx, deliveryID, eventName, event)
	if err != nil {
		return g.handleError(ctx, deliveryID, eventName, event, err)
	}

	err = g.handleStatusEventAny(ctx, deliveryID, eventName, event)
	if err != nil {
		return g.handleError(ctx, deliveryID, eventName, event, err)
	}

	err = g.handleAfterAny(ctx, deliveryID, eventName, event)
	if err != nil {
		return g.handleError(ctx, deliveryID, eventName, event, err)
	}
	return nil
}
