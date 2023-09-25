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
	// CommitCommentEvent is the event name of github.CommitCommentEvent's
	CommitCommentEvent = "commit_comment"

	// CommitCommentEventAnyAction is used to identify callbacks
	// listening to all events of type github.CommitCommentEvent
	CommitCommentEventAnyAction = "*"

	// CommitCommentEventCreatedAction is used to identify callbacks
	// listening to events of type github.CommitCommentEvent and action "created"
	CommitCommentEventCreatedAction = "created"
)

// CommitCommentEventHandleFunc represents a callback function triggered on github.CommitCommentEvent's.
// 'deliveryID' (type: string) is the unique webhook delivery ID.
// 'eventName' (type: string) is the name of the event.
// 'event' (type: *github.CommitCommentEvent) is the webhook payload.
type CommitCommentEventHandleFunc func(deliveryID string, eventName string, event *github.CommitCommentEvent) error

// OnCommitCommentEventCreated registers callbacks listening to events of type github.CommitCommentEvent and action 'created'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnCommitCommentEventCreated must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#commit_comment
func (g *EventHandler) OnCommitCommentEventCreated(callbacks ...CommitCommentEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onCommitCommentEvent == nil {
		g.onCommitCommentEvent = make(map[string][]CommitCommentEventHandleFunc)
	}
	g.onCommitCommentEvent[CommitCommentEventCreatedAction] = append(
		g.onCommitCommentEvent[CommitCommentEventCreatedAction],
		callbacks...,
	)
}

// SetOnCommitCommentEventCreated registers callbacks listening to events of type github.CommitCommentEvent and action 'created'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnCommitCommentEventCreatedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#commit_comment
func (g *EventHandler) SetOnCommitCommentEventCreated(callbacks ...CommitCommentEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onCommitCommentEvent == nil {
		g.onCommitCommentEvent = make(map[string][]CommitCommentEventHandleFunc)
	}
	g.onCommitCommentEvent[CommitCommentEventCreatedAction] = callbacks
}

func (g *EventHandler) handleCommitCommentEventCreated(deliveryID string, eventName string, event *github.CommitCommentEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if CommitCommentEventCreatedAction != *event.Action {
		return fmt.Errorf(
			"handleCommitCommentEventCreated() called with wrong action, want %s, got %s",
			CommitCommentEventCreatedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		CommitCommentEventCreatedAction,
		CommitCommentEventAnyAction,
	} {
		if _, ok := g.onCommitCommentEvent[action]; ok {
			for _, h := range g.onCommitCommentEvent[action] {
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

// OnCommitCommentEventAny registers callbacks listening to any events of type github.CommitCommentEvent
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnCommitCommentEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#commit_comment
func (g *EventHandler) OnCommitCommentEventAny(callbacks ...CommitCommentEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onCommitCommentEvent == nil {
		g.onCommitCommentEvent = make(map[string][]CommitCommentEventHandleFunc)
	}
	g.onCommitCommentEvent[CommitCommentEventAnyAction] = append(
		g.onCommitCommentEvent[CommitCommentEventAnyAction],
		callbacks...,
	)
}

// SetOnCommitCommentEventAny registers callbacks listening to any events of type github.CommitCommentEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnCommitCommentEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#commit_comment
func (g *EventHandler) SetOnCommitCommentEventAny(callbacks ...CommitCommentEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onCommitCommentEvent == nil {
		g.onCommitCommentEvent = make(map[string][]CommitCommentEventHandleFunc)
	}
	g.onCommitCommentEvent[CommitCommentEventAnyAction] = callbacks
}

func (g *EventHandler) handleCommitCommentEventAny(deliveryID string, eventName string, event *github.CommitCommentEvent) error {
	if event == nil {
		return fmt.Errorf("event was empty or nil")
	}
	if _, ok := g.onCommitCommentEvent[CommitCommentEventAnyAction]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onCommitCommentEvent[CommitCommentEventAnyAction] {
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

// CommitCommentEvent handles github.CommitCommentEvent.
//
// Callbacks are executed in the following order:
//
// 1) All callbacks registered with OnBeforeAny are executed in parallel.
// 2) All callbacks registered with OnCommitCommentEvent... are executed in parallel in case the Event has actions.
// 3) All callbacks registered with OnAfterAny are executed in parallel.
//
// on any error all callbacks registered with OnError are executed in parallel.
func (g *EventHandler) CommitCommentEvent(deliveryID string, eventName string, event *github.CommitCommentEvent) error {

	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	action := *event.Action

	err := g.handleBeforeAny(deliveryID, eventName, event)
	if err != nil {
		return g.handleError(deliveryID, eventName, event, err)
	}

	switch action {

	case CommitCommentEventCreatedAction:
		err := g.handleCommitCommentEventCreated(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	default:
		err := g.handleCommitCommentEventAny(deliveryID, eventName, event)
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
