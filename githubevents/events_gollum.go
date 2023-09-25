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
	// GollumEvent is the event name of github.GollumEvent's
	GollumEvent = "gollum"

	// GollumEventAnyAction is used to identify callbacks
	// listening to all events of type github.GollumEvent
	GollumEventAnyAction = "*"
)

// GollumEventHandleFunc represents a callback function triggered on github.GollumEvent's.
// 'deliveryID' (type: string) is the unique webhook delivery ID.
// 'eventName' (type: string) is the name of the event.
// 'event' (type: *github.GollumEvent) is the webhook payload.
type GollumEventHandleFunc func(deliveryID string, eventName string, event *github.GollumEvent) error

// OnGollumEventAny registers callbacks listening to any events of type github.GollumEvent
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnGollumEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#gollum
func (g *EventHandler) OnGollumEventAny(callbacks ...GollumEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onGollumEvent == nil {
		g.onGollumEvent = make(map[string][]GollumEventHandleFunc)
	}
	g.onGollumEvent[GollumEventAnyAction] = append(
		g.onGollumEvent[GollumEventAnyAction],
		callbacks...,
	)
}

// SetOnGollumEventAny registers callbacks listening to any events of type github.GollumEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnGollumEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#gollum
func (g *EventHandler) SetOnGollumEventAny(callbacks ...GollumEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onGollumEvent == nil {
		g.onGollumEvent = make(map[string][]GollumEventHandleFunc)
	}
	g.onGollumEvent[GollumEventAnyAction] = callbacks
}

func (g *EventHandler) handleGollumEventAny(deliveryID string, eventName string, event *github.GollumEvent) error {
	if event == nil {
		return fmt.Errorf("event was empty or nil")
	}
	if _, ok := g.onGollumEvent[GollumEventAnyAction]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onGollumEvent[GollumEventAnyAction] {
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

// GollumEvent handles github.GollumEvent.
//
// Callbacks are executed in the following order:
//
// 1) All callbacks registered with OnBeforeAny are executed in parallel.
// 2) All callbacks registered with OnGollumEvent... are executed in parallel in case the Event has actions.
// 3) All callbacks registered with OnAfterAny are executed in parallel.
//
// on any error all callbacks registered with OnError are executed in parallel.
func (g *EventHandler) GollumEvent(deliveryID string, eventName string, event *github.GollumEvent) error {

	if event == nil {
		return fmt.Errorf("event action was empty or nil")
	}

	err := g.handleBeforeAny(deliveryID, eventName, event)
	if err != nil {
		return g.handleError(deliveryID, eventName, event, err)
	}

	err = g.handleGollumEventAny(deliveryID, eventName, event)
	if err != nil {
		return g.handleError(deliveryID, eventName, event, err)
	}

	err = g.handleAfterAny(deliveryID, eventName, event)
	if err != nil {
		return g.handleError(deliveryID, eventName, event, err)
	}
	return nil
}
