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
	// IssueCommentEvent is the event name of github.IssueCommentEvent's
	IssueCommentEvent = "issue_comment"

	// IssueCommentEventAnyAction is used to identify callbacks
	// listening to all events of type github.IssueCommentEvent
	IssueCommentEventAnyAction = "*"

	// IssueCommentCreatedAction is used to identify callbacks
	// listening to events of type github.IssueCommentEvent and action "created"
	IssueCommentCreatedAction = "created"

	// IssueCommentEditedAction is used to identify callbacks
	// listening to events of type github.IssueCommentEvent and action "edited"
	IssueCommentEditedAction = "edited"

	// IssueCommentDeletedAction is used to identify callbacks
	// listening to events of type github.IssueCommentEvent and action "deleted"
	IssueCommentDeletedAction = "deleted"
)

// IssueCommentEventHandleFunc represents a callback function triggered on github.IssueCommentEvent's.
// 'deliveryID' (type: string) is the unique webhook delivery ID.
// 'eventName' (type: string) is the name of the event.
// 'event' (type: *github.IssueCommentEvent) is the webhook payload.
type IssueCommentEventHandleFunc func(deliveryID string, eventName string, event *github.IssueCommentEvent) error

// OnIssueCommentCreated registers callbacks listening to events of type github.IssueCommentEvent and action 'created'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnIssueCommentCreated must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#issue_comment
func (g *EventHandler) OnIssueCommentCreated(callbacks ...IssueCommentEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssueCommentEvent == nil {
		g.onIssueCommentEvent = make(map[string][]IssueCommentEventHandleFunc)
	}
	g.onIssueCommentEvent[IssueCommentCreatedAction] = append(
		g.onIssueCommentEvent[IssueCommentCreatedAction],
		callbacks...,
	)
}

// SetOnIssueCommentCreated registers callbacks listening to events of type github.IssueCommentEvent and action 'created'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnIssueCommentCreatedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#issue_comment
func (g *EventHandler) SetOnIssueCommentCreated(callbacks ...IssueCommentEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssueCommentEvent == nil {
		g.onIssueCommentEvent = make(map[string][]IssueCommentEventHandleFunc)
	}
	g.onIssueCommentEvent[IssueCommentCreatedAction] = callbacks
}

func (g *EventHandler) handleIssueCommentCreated(deliveryID string, eventName string, event *github.IssueCommentEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if IssueCommentCreatedAction != *event.Action {
		return fmt.Errorf(
			"handleIssueCommentCreated() called with wrong action, want %s, got %s",
			IssueCommentCreatedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		IssueCommentCreatedAction,
		IssueCommentEventAnyAction,
	} {
		if _, ok := g.onIssueCommentEvent[action]; ok {
			for _, h := range g.onIssueCommentEvent[action] {
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

// OnIssueCommentEdited registers callbacks listening to events of type github.IssueCommentEvent and action 'edited'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnIssueCommentEdited must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#issue_comment
func (g *EventHandler) OnIssueCommentEdited(callbacks ...IssueCommentEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssueCommentEvent == nil {
		g.onIssueCommentEvent = make(map[string][]IssueCommentEventHandleFunc)
	}
	g.onIssueCommentEvent[IssueCommentEditedAction] = append(
		g.onIssueCommentEvent[IssueCommentEditedAction],
		callbacks...,
	)
}

// SetOnIssueCommentEdited registers callbacks listening to events of type github.IssueCommentEvent and action 'edited'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnIssueCommentEditedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#issue_comment
func (g *EventHandler) SetOnIssueCommentEdited(callbacks ...IssueCommentEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssueCommentEvent == nil {
		g.onIssueCommentEvent = make(map[string][]IssueCommentEventHandleFunc)
	}
	g.onIssueCommentEvent[IssueCommentEditedAction] = callbacks
}

func (g *EventHandler) handleIssueCommentEdited(deliveryID string, eventName string, event *github.IssueCommentEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if IssueCommentEditedAction != *event.Action {
		return fmt.Errorf(
			"handleIssueCommentEdited() called with wrong action, want %s, got %s",
			IssueCommentEditedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		IssueCommentEditedAction,
		IssueCommentEventAnyAction,
	} {
		if _, ok := g.onIssueCommentEvent[action]; ok {
			for _, h := range g.onIssueCommentEvent[action] {
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

// OnIssueCommentDeleted registers callbacks listening to events of type github.IssueCommentEvent and action 'deleted'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnIssueCommentDeleted must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#issue_comment
func (g *EventHandler) OnIssueCommentDeleted(callbacks ...IssueCommentEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssueCommentEvent == nil {
		g.onIssueCommentEvent = make(map[string][]IssueCommentEventHandleFunc)
	}
	g.onIssueCommentEvent[IssueCommentDeletedAction] = append(
		g.onIssueCommentEvent[IssueCommentDeletedAction],
		callbacks...,
	)
}

// SetOnIssueCommentDeleted registers callbacks listening to events of type github.IssueCommentEvent and action 'deleted'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnIssueCommentDeletedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#issue_comment
func (g *EventHandler) SetOnIssueCommentDeleted(callbacks ...IssueCommentEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssueCommentEvent == nil {
		g.onIssueCommentEvent = make(map[string][]IssueCommentEventHandleFunc)
	}
	g.onIssueCommentEvent[IssueCommentDeletedAction] = callbacks
}

func (g *EventHandler) handleIssueCommentDeleted(deliveryID string, eventName string, event *github.IssueCommentEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if IssueCommentDeletedAction != *event.Action {
		return fmt.Errorf(
			"handleIssueCommentDeleted() called with wrong action, want %s, got %s",
			IssueCommentDeletedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		IssueCommentDeletedAction,
		IssueCommentEventAnyAction,
	} {
		if _, ok := g.onIssueCommentEvent[action]; ok {
			for _, h := range g.onIssueCommentEvent[action] {
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

// OnIssueCommentEventAny registers callbacks listening to any events of type github.IssueCommentEvent
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnIssueCommentEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#issue_comment
func (g *EventHandler) OnIssueCommentEventAny(callbacks ...IssueCommentEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssueCommentEvent == nil {
		g.onIssueCommentEvent = make(map[string][]IssueCommentEventHandleFunc)
	}
	g.onIssueCommentEvent[IssueCommentEventAnyAction] = append(
		g.onIssueCommentEvent[IssueCommentEventAnyAction],
		callbacks...,
	)
}

// SetOnIssueCommentEventAny registers callbacks listening to any events of type github.IssueCommentEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnIssueCommentEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#issue_comment
func (g *EventHandler) SetOnIssueCommentEventAny(callbacks ...IssueCommentEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssueCommentEvent == nil {
		g.onIssueCommentEvent = make(map[string][]IssueCommentEventHandleFunc)
	}
	g.onIssueCommentEvent[IssueCommentEventAnyAction] = callbacks
}

func (g *EventHandler) handleIssueCommentEventAny(deliveryID string, eventName string, event *github.IssueCommentEvent) error {
	if event == nil {
		return fmt.Errorf("event was empty or nil")
	}
	if _, ok := g.onIssueCommentEvent[IssueCommentEventAnyAction]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onIssueCommentEvent[IssueCommentEventAnyAction] {
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

// IssueCommentEvent handles github.IssueCommentEvent.
//
// Callbacks are executed in the following order:
//
// 1) All callbacks registered with OnBeforeAny are executed in parallel.
// 2) All callbacks registered with OnIssueCommentEvent... are executed in parallel in case the Event has actions.
// 3) All callbacks registered with OnAfterAny are executed in parallel.
//
// on any error all callbacks registered with OnError are executed in parallel.
func (g *EventHandler) IssueCommentEvent(deliveryID string, eventName string, event *github.IssueCommentEvent) error {

	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	action := *event.Action

	err := g.handleBeforeAny(deliveryID, eventName, event)
	if err != nil {
		return g.handleError(deliveryID, eventName, event, err)
	}

	switch action {

	case IssueCommentCreatedAction:
		err := g.handleIssueCommentCreated(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case IssueCommentEditedAction:
		err := g.handleIssueCommentEdited(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case IssueCommentDeletedAction:
		err := g.handleIssueCommentDeleted(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	default:
		err := g.handleIssueCommentEventAny(deliveryID, eventName, event)
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
