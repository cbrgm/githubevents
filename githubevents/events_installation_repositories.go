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
	// InstallationRepositoriesEvent is the event name of github.InstallationRepositoriesEvent's
	InstallationRepositoriesEvent = "installation_repositories"

	// InstallationRepositoriesEventAnyAction is used to identify callbacks
	// listening to all events of type github.InstallationRepositoriesEvent
	InstallationRepositoriesEventAnyAction = "*"

	// InstallationRepositoriesEventAddedAction is used to identify callbacks
	// listening to events of type github.InstallationRepositoriesEvent and action "added"
	InstallationRepositoriesEventAddedAction = "added"

	// InstallationRepositoriesEventRemovedAction is used to identify callbacks
	// listening to events of type github.InstallationRepositoriesEvent and action "removed"
	InstallationRepositoriesEventRemovedAction = "removed"
)

// InstallationRepositoriesEventHandleFunc represents a callback function triggered on github.InstallationRepositoriesEvent's.
// 'deliveryID' (type: string) is the unique webhook delivery ID.
// 'eventName' (type: string) is the name of the event.
// 'event' (type: *github.InstallationRepositoriesEvent) is the webhook payload.
type InstallationRepositoriesEventHandleFunc func(deliveryID string, eventName string, event *github.InstallationRepositoriesEvent) error

// OnInstallationRepositoriesEventAdded registers callbacks listening to events of type github.InstallationRepositoriesEvent and action 'added'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnInstallationRepositoriesEventAdded must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#installation_repositories
func (g *EventHandler) OnInstallationRepositoriesEventAdded(callbacks ...InstallationRepositoriesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onInstallationRepositoriesEvent == nil {
		g.onInstallationRepositoriesEvent = make(map[string][]InstallationRepositoriesEventHandleFunc)
	}
	g.onInstallationRepositoriesEvent[InstallationRepositoriesEventAddedAction] = append(
		g.onInstallationRepositoriesEvent[InstallationRepositoriesEventAddedAction],
		callbacks...,
	)
}

// SetOnInstallationRepositoriesEventAdded registers callbacks listening to events of type github.InstallationRepositoriesEvent and action 'added'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnInstallationRepositoriesEventAddedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#installation_repositories
func (g *EventHandler) SetOnInstallationRepositoriesEventAdded(callbacks ...InstallationRepositoriesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onInstallationRepositoriesEvent == nil {
		g.onInstallationRepositoriesEvent = make(map[string][]InstallationRepositoriesEventHandleFunc)
	}
	g.onInstallationRepositoriesEvent[InstallationRepositoriesEventAddedAction] = callbacks
}

func (g *EventHandler) handleInstallationRepositoriesEventAdded(deliveryID string, eventName string, event *github.InstallationRepositoriesEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if InstallationRepositoriesEventAddedAction != *event.Action {
		return fmt.Errorf(
			"handleInstallationRepositoriesEventAdded() called with wrong action, want %s, got %s",
			InstallationRepositoriesEventAddedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		InstallationRepositoriesEventAddedAction,
		InstallationRepositoriesEventAnyAction,
	} {
		if _, ok := g.onInstallationRepositoriesEvent[action]; ok {
			for _, h := range g.onInstallationRepositoriesEvent[action] {
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

// OnInstallationRepositoriesEventRemoved registers callbacks listening to events of type github.InstallationRepositoriesEvent and action 'removed'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnInstallationRepositoriesEventRemoved must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#installation_repositories
func (g *EventHandler) OnInstallationRepositoriesEventRemoved(callbacks ...InstallationRepositoriesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onInstallationRepositoriesEvent == nil {
		g.onInstallationRepositoriesEvent = make(map[string][]InstallationRepositoriesEventHandleFunc)
	}
	g.onInstallationRepositoriesEvent[InstallationRepositoriesEventRemovedAction] = append(
		g.onInstallationRepositoriesEvent[InstallationRepositoriesEventRemovedAction],
		callbacks...,
	)
}

// SetOnInstallationRepositoriesEventRemoved registers callbacks listening to events of type github.InstallationRepositoriesEvent and action 'removed'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnInstallationRepositoriesEventRemovedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#installation_repositories
func (g *EventHandler) SetOnInstallationRepositoriesEventRemoved(callbacks ...InstallationRepositoriesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onInstallationRepositoriesEvent == nil {
		g.onInstallationRepositoriesEvent = make(map[string][]InstallationRepositoriesEventHandleFunc)
	}
	g.onInstallationRepositoriesEvent[InstallationRepositoriesEventRemovedAction] = callbacks
}

func (g *EventHandler) handleInstallationRepositoriesEventRemoved(deliveryID string, eventName string, event *github.InstallationRepositoriesEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if InstallationRepositoriesEventRemovedAction != *event.Action {
		return fmt.Errorf(
			"handleInstallationRepositoriesEventRemoved() called with wrong action, want %s, got %s",
			InstallationRepositoriesEventRemovedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		InstallationRepositoriesEventRemovedAction,
		InstallationRepositoriesEventAnyAction,
	} {
		if _, ok := g.onInstallationRepositoriesEvent[action]; ok {
			for _, h := range g.onInstallationRepositoriesEvent[action] {
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

// OnInstallationRepositoriesEventAny registers callbacks listening to any events of type github.InstallationRepositoriesEvent
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnInstallationRepositoriesEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#installation_repositories
func (g *EventHandler) OnInstallationRepositoriesEventAny(callbacks ...InstallationRepositoriesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onInstallationRepositoriesEvent == nil {
		g.onInstallationRepositoriesEvent = make(map[string][]InstallationRepositoriesEventHandleFunc)
	}
	g.onInstallationRepositoriesEvent[InstallationRepositoriesEventAnyAction] = append(
		g.onInstallationRepositoriesEvent[InstallationRepositoriesEventAnyAction],
		callbacks...,
	)
}

// SetOnInstallationRepositoriesEventAny registers callbacks listening to any events of type github.InstallationRepositoriesEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnInstallationRepositoriesEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#installation_repositories
func (g *EventHandler) SetOnInstallationRepositoriesEventAny(callbacks ...InstallationRepositoriesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onInstallationRepositoriesEvent == nil {
		g.onInstallationRepositoriesEvent = make(map[string][]InstallationRepositoriesEventHandleFunc)
	}
	g.onInstallationRepositoriesEvent[InstallationRepositoriesEventAnyAction] = callbacks
}

func (g *EventHandler) handleInstallationRepositoriesEventAny(deliveryID string, eventName string, event *github.InstallationRepositoriesEvent) error {
	if event == nil {
		return fmt.Errorf("event was empty or nil")
	}
	if _, ok := g.onInstallationRepositoriesEvent[InstallationRepositoriesEventAnyAction]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onInstallationRepositoriesEvent[InstallationRepositoriesEventAnyAction] {
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

// InstallationRepositoriesEvent handles github.InstallationRepositoriesEvent.
//
// Callbacks are executed in the following order:
//
// 1) All callbacks registered with OnBeforeAny are executed in parallel.
// 2) All callbacks registered with OnInstallationRepositoriesEvent... are executed in parallel in case the Event has actions.
// 3) All callbacks registered with OnAfterAny are executed in parallel.
//
// on any error all callbacks registered with OnError are executed in parallel.
func (g *EventHandler) InstallationRepositoriesEvent(deliveryID string, eventName string, event *github.InstallationRepositoriesEvent) error {

	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	action := *event.Action

	err := g.handleBeforeAny(deliveryID, eventName, event)
	if err != nil {
		return g.handleError(deliveryID, eventName, event, err)
	}

	switch action {

	case InstallationRepositoriesEventAddedAction:
		err := g.handleInstallationRepositoriesEventAdded(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case InstallationRepositoriesEventRemovedAction:
		err := g.handleInstallationRepositoriesEventRemoved(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	default:
		err := g.handleInstallationRepositoriesEventAny(deliveryID, eventName, event)
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
