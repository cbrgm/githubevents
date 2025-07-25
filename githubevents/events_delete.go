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
	// DeleteEvent is the event name of github.DeleteEvent's
	DeleteEvent = "delete"

	// DeleteEventAnyAction is used to identify callbacks
	// listening to all events of type github.DeleteEvent
	DeleteEventAnyAction = "*"
)

// DeleteEventHandleFunc represents a callback function triggered on github.DeleteEvent's.
// 'deliveryID' (type: string) is the unique webhook delivery ID.
// 'eventName' (type: string) is the name of the event.
// 'event' (type: *github.DeleteEvent) is the webhook payload.
type DeleteEventHandleFunc func(ctx context.Context, deliveryID string, eventName string, event *github.DeleteEvent) error

// OnDeleteEventAny registers callbacks listening to any events of type github.DeleteEvent
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnDeleteEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#delete
func (g *EventHandler) OnDeleteEventAny(callbacks ...DeleteEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onDeleteEvent == nil {
		g.onDeleteEvent = make(map[string][]DeleteEventHandleFunc)
	}
	g.onDeleteEvent[DeleteEventAnyAction] = append(
		g.onDeleteEvent[DeleteEventAnyAction],
		callbacks...,
	)
}

// SetOnDeleteEventAny registers callbacks listening to any events of type github.DeleteEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnDeleteEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#delete
func (g *EventHandler) SetOnDeleteEventAny(callbacks ...DeleteEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onDeleteEvent == nil {
		g.onDeleteEvent = make(map[string][]DeleteEventHandleFunc)
	}
	g.onDeleteEvent[DeleteEventAnyAction] = callbacks
}

func (g *EventHandler) handleDeleteEventAny(ctx context.Context, deliveryID string, eventName string, event *github.DeleteEvent) error {
	if event == nil {
		return fmt.Errorf("event was empty or nil")
	}
	if _, ok := g.onDeleteEvent[DeleteEventAnyAction]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onDeleteEvent[DeleteEventAnyAction] {
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

// DeleteEvent handles github.DeleteEvent.
//
// Callbacks are executed in the following order:
//
// 1) All callbacks registered with OnBeforeAny are executed in parallel.
// 2) All callbacks registered with OnDeleteEvent... are executed in parallel in case the Event has actions.
// 3) All callbacks registered with OnAfterAny are executed in parallel.
//
// on any error all callbacks registered with OnError are executed in parallel.
func (g *EventHandler) DeleteEvent(ctx context.Context, deliveryID string, eventName string, event *github.DeleteEvent) error {

	if event == nil {
		return fmt.Errorf("event action was empty or nil")
	}

	err := g.handleBeforeAny(ctx, deliveryID, eventName, event)
	if err != nil {
		return g.handleError(ctx, deliveryID, eventName, event, err)
	}

	err = g.handleDeleteEventAny(ctx, deliveryID, eventName, event)
	if err != nil {
		return g.handleError(ctx, deliveryID, eventName, event, err)
	}

	err = g.handleAfterAny(ctx, deliveryID, eventName, event)
	if err != nil {
		return g.handleError(ctx, deliveryID, eventName, event, err)
	}
	return nil
}
