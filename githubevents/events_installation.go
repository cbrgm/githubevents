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
	// InstallationEvent is the event name of github.InstallationEvent's
	InstallationEvent = "installation"

	// InstallationEventAnyAction is used to identify callbacks
	// listening to all events of type github.InstallationEvent
	InstallationEventAnyAction = "*"

	// InstallationEventCreatedAction is used to identify callbacks
	// listening to events of type github.InstallationEvent and action "created"
	InstallationEventCreatedAction = "created"

	// InstallationEventDeletedAction is used to identify callbacks
	// listening to events of type github.InstallationEvent and action "deleted"
	InstallationEventDeletedAction = "deleted"

	// InstallationEventEventSuspendAction is used to identify callbacks
	// listening to events of type github.InstallationEvent and action "suspend"
	InstallationEventEventSuspendAction = "suspend"

	// InstallationEventEventUnsuspendAction is used to identify callbacks
	// listening to events of type github.InstallationEvent and action "unsuspend"
	InstallationEventEventUnsuspendAction = "unsuspend"

	// InstallationEventNewPermissionsAcceptedAction is used to identify callbacks
	// listening to events of type github.InstallationEvent and action "new_permissions_accepted"
	InstallationEventNewPermissionsAcceptedAction = "new_permissions_accepted"
)

// InstallationEventHandleFunc represents a callback function triggered on github.InstallationEvent's.
// 'deliveryID' (type: string) is the unique webhook delivery ID.
// 'eventName' (type: string) is the name of the event.
// 'event' (type: *github.InstallationEvent) is the webhook payload.
type InstallationEventHandleFunc func(deliveryID string, eventName string, event *github.InstallationEvent) error

// OnInstallationEventCreated registers callbacks listening to events of type github.InstallationEvent and action 'created'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnInstallationEventCreated must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#installation
func (g *EventHandler) OnInstallationEventCreated(callbacks ...InstallationEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onInstallationEvent == nil {
		g.onInstallationEvent = make(map[string][]InstallationEventHandleFunc)
	}
	g.onInstallationEvent[InstallationEventCreatedAction] = append(
		g.onInstallationEvent[InstallationEventCreatedAction],
		callbacks...,
	)
}

// SetOnInstallationEventCreated registers callbacks listening to events of type github.InstallationEvent and action 'created'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnInstallationEventCreatedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#installation
func (g *EventHandler) SetOnInstallationEventCreated(callbacks ...InstallationEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onInstallationEvent == nil {
		g.onInstallationEvent = make(map[string][]InstallationEventHandleFunc)
	}
	g.onInstallationEvent[InstallationEventCreatedAction] = callbacks
}

func (g *EventHandler) handleInstallationEventCreated(deliveryID string, eventName string, event *github.InstallationEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if InstallationEventCreatedAction != *event.Action {
		return fmt.Errorf(
			"handleInstallationEventCreated() called with wrong action, want %s, got %s",
			InstallationEventCreatedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		InstallationEventCreatedAction,
		InstallationEventAnyAction,
	} {
		if _, ok := g.onInstallationEvent[action]; ok {
			for _, h := range g.onInstallationEvent[action] {
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

// OnInstallationEventDeleted registers callbacks listening to events of type github.InstallationEvent and action 'deleted'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnInstallationEventDeleted must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#installation
func (g *EventHandler) OnInstallationEventDeleted(callbacks ...InstallationEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onInstallationEvent == nil {
		g.onInstallationEvent = make(map[string][]InstallationEventHandleFunc)
	}
	g.onInstallationEvent[InstallationEventDeletedAction] = append(
		g.onInstallationEvent[InstallationEventDeletedAction],
		callbacks...,
	)
}

// SetOnInstallationEventDeleted registers callbacks listening to events of type github.InstallationEvent and action 'deleted'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnInstallationEventDeletedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#installation
func (g *EventHandler) SetOnInstallationEventDeleted(callbacks ...InstallationEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onInstallationEvent == nil {
		g.onInstallationEvent = make(map[string][]InstallationEventHandleFunc)
	}
	g.onInstallationEvent[InstallationEventDeletedAction] = callbacks
}

func (g *EventHandler) handleInstallationEventDeleted(deliveryID string, eventName string, event *github.InstallationEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if InstallationEventDeletedAction != *event.Action {
		return fmt.Errorf(
			"handleInstallationEventDeleted() called with wrong action, want %s, got %s",
			InstallationEventDeletedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		InstallationEventDeletedAction,
		InstallationEventAnyAction,
	} {
		if _, ok := g.onInstallationEvent[action]; ok {
			for _, h := range g.onInstallationEvent[action] {
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

// OnInstallationEventEventSuspend registers callbacks listening to events of type github.InstallationEvent and action 'suspend'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnInstallationEventEventSuspend must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#installation
func (g *EventHandler) OnInstallationEventEventSuspend(callbacks ...InstallationEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onInstallationEvent == nil {
		g.onInstallationEvent = make(map[string][]InstallationEventHandleFunc)
	}
	g.onInstallationEvent[InstallationEventEventSuspendAction] = append(
		g.onInstallationEvent[InstallationEventEventSuspendAction],
		callbacks...,
	)
}

// SetOnInstallationEventEventSuspend registers callbacks listening to events of type github.InstallationEvent and action 'suspend'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnInstallationEventEventSuspendAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#installation
func (g *EventHandler) SetOnInstallationEventEventSuspend(callbacks ...InstallationEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onInstallationEvent == nil {
		g.onInstallationEvent = make(map[string][]InstallationEventHandleFunc)
	}
	g.onInstallationEvent[InstallationEventEventSuspendAction] = callbacks
}

func (g *EventHandler) handleInstallationEventEventSuspend(deliveryID string, eventName string, event *github.InstallationEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if InstallationEventEventSuspendAction != *event.Action {
		return fmt.Errorf(
			"handleInstallationEventEventSuspend() called with wrong action, want %s, got %s",
			InstallationEventEventSuspendAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		InstallationEventEventSuspendAction,
		InstallationEventAnyAction,
	} {
		if _, ok := g.onInstallationEvent[action]; ok {
			for _, h := range g.onInstallationEvent[action] {
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

// OnInstallationEventEventUnsuspend registers callbacks listening to events of type github.InstallationEvent and action 'unsuspend'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnInstallationEventEventUnsuspend must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#installation
func (g *EventHandler) OnInstallationEventEventUnsuspend(callbacks ...InstallationEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onInstallationEvent == nil {
		g.onInstallationEvent = make(map[string][]InstallationEventHandleFunc)
	}
	g.onInstallationEvent[InstallationEventEventUnsuspendAction] = append(
		g.onInstallationEvent[InstallationEventEventUnsuspendAction],
		callbacks...,
	)
}

// SetOnInstallationEventEventUnsuspend registers callbacks listening to events of type github.InstallationEvent and action 'unsuspend'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnInstallationEventEventUnsuspendAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#installation
func (g *EventHandler) SetOnInstallationEventEventUnsuspend(callbacks ...InstallationEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onInstallationEvent == nil {
		g.onInstallationEvent = make(map[string][]InstallationEventHandleFunc)
	}
	g.onInstallationEvent[InstallationEventEventUnsuspendAction] = callbacks
}

func (g *EventHandler) handleInstallationEventEventUnsuspend(deliveryID string, eventName string, event *github.InstallationEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if InstallationEventEventUnsuspendAction != *event.Action {
		return fmt.Errorf(
			"handleInstallationEventEventUnsuspend() called with wrong action, want %s, got %s",
			InstallationEventEventUnsuspendAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		InstallationEventEventUnsuspendAction,
		InstallationEventAnyAction,
	} {
		if _, ok := g.onInstallationEvent[action]; ok {
			for _, h := range g.onInstallationEvent[action] {
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

// OnInstallationEventNewPermissionsAccepted registers callbacks listening to events of type github.InstallationEvent and action 'new_permissions_accepted'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnInstallationEventNewPermissionsAccepted must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#installation
func (g *EventHandler) OnInstallationEventNewPermissionsAccepted(callbacks ...InstallationEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onInstallationEvent == nil {
		g.onInstallationEvent = make(map[string][]InstallationEventHandleFunc)
	}
	g.onInstallationEvent[InstallationEventNewPermissionsAcceptedAction] = append(
		g.onInstallationEvent[InstallationEventNewPermissionsAcceptedAction],
		callbacks...,
	)
}

// SetOnInstallationEventNewPermissionsAccepted registers callbacks listening to events of type github.InstallationEvent and action 'new_permissions_accepted'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnInstallationEventNewPermissionsAcceptedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#installation
func (g *EventHandler) SetOnInstallationEventNewPermissionsAccepted(callbacks ...InstallationEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onInstallationEvent == nil {
		g.onInstallationEvent = make(map[string][]InstallationEventHandleFunc)
	}
	g.onInstallationEvent[InstallationEventNewPermissionsAcceptedAction] = callbacks
}

func (g *EventHandler) handleInstallationEventNewPermissionsAccepted(deliveryID string, eventName string, event *github.InstallationEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if InstallationEventNewPermissionsAcceptedAction != *event.Action {
		return fmt.Errorf(
			"handleInstallationEventNewPermissionsAccepted() called with wrong action, want %s, got %s",
			InstallationEventNewPermissionsAcceptedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		InstallationEventNewPermissionsAcceptedAction,
		InstallationEventAnyAction,
	} {
		if _, ok := g.onInstallationEvent[action]; ok {
			for _, h := range g.onInstallationEvent[action] {
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

// OnInstallationEventAny registers callbacks listening to any events of type github.InstallationEvent
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnInstallationEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#installation
func (g *EventHandler) OnInstallationEventAny(callbacks ...InstallationEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onInstallationEvent == nil {
		g.onInstallationEvent = make(map[string][]InstallationEventHandleFunc)
	}
	g.onInstallationEvent[InstallationEventAnyAction] = append(
		g.onInstallationEvent[InstallationEventAnyAction],
		callbacks...,
	)
}

// SetOnInstallationEventAny registers callbacks listening to any events of type github.InstallationEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnInstallationEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#installation
func (g *EventHandler) SetOnInstallationEventAny(callbacks ...InstallationEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onInstallationEvent == nil {
		g.onInstallationEvent = make(map[string][]InstallationEventHandleFunc)
	}
	g.onInstallationEvent[InstallationEventAnyAction] = callbacks
}

func (g *EventHandler) handleInstallationEventAny(deliveryID string, eventName string, event *github.InstallationEvent) error {
	if event == nil {
		return fmt.Errorf("event was empty or nil")
	}
	if _, ok := g.onInstallationEvent[InstallationEventAnyAction]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onInstallationEvent[InstallationEventAnyAction] {
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

// InstallationEvent handles github.InstallationEvent.
//
// Callbacks are executed in the following order:
//
// 1) All callbacks registered with OnBeforeAny are executed in parallel.
// 2) All callbacks registered with OnInstallationEvent... are executed in parallel in case the Event has actions.
// 3) All callbacks registered with OnAfterAny are executed in parallel.
//
// on any error all callbacks registered with OnError are executed in parallel.
func (g *EventHandler) InstallationEvent(deliveryID string, eventName string, event *github.InstallationEvent) error {

	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	action := *event.Action

	err := g.handleBeforeAny(deliveryID, eventName, event)
	if err != nil {
		return g.handleError(deliveryID, eventName, event, err)
	}

	switch action {

	case InstallationEventCreatedAction:
		err := g.handleInstallationEventCreated(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case InstallationEventDeletedAction:
		err := g.handleInstallationEventDeleted(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case InstallationEventEventSuspendAction:
		err := g.handleInstallationEventEventSuspend(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case InstallationEventEventUnsuspendAction:
		err := g.handleInstallationEventEventUnsuspend(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case InstallationEventNewPermissionsAcceptedAction:
		err := g.handleInstallationEventNewPermissionsAccepted(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	default:
		err := g.handleInstallationEventAny(deliveryID, eventName, event)
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
