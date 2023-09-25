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
	// GitHubAppAuthorizationEvent is the event name of github.GitHubAppAuthorizationEvent's
	GitHubAppAuthorizationEvent = "github_app_authorization"

	// GitHubAppAuthorizationEventAnyAction is used to identify callbacks
	// listening to all events of type github.GitHubAppAuthorizationEvent
	GitHubAppAuthorizationEventAnyAction = "*"

	// GitHubAppAuthorizationEventRevokedAction is used to identify callbacks
	// listening to events of type github.GitHubAppAuthorizationEvent and action "revoked"
	GitHubAppAuthorizationEventRevokedAction = "revoked"
)

// GitHubAppAuthorizationEventHandleFunc represents a callback function triggered on github.GitHubAppAuthorizationEvent's.
// 'deliveryID' (type: string) is the unique webhook delivery ID.
// 'eventName' (type: string) is the name of the event.
// 'event' (type: *github.GitHubAppAuthorizationEvent) is the webhook payload.
type GitHubAppAuthorizationEventHandleFunc func(deliveryID string, eventName string, event *github.GitHubAppAuthorizationEvent) error

// OnGitHubAppAuthorizationEventRevoked registers callbacks listening to events of type github.GitHubAppAuthorizationEvent and action 'revoked'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnGitHubAppAuthorizationEventRevoked must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#github_app_authorization
func (g *EventHandler) OnGitHubAppAuthorizationEventRevoked(callbacks ...GitHubAppAuthorizationEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onGitHubAppAuthorizationEvent == nil {
		g.onGitHubAppAuthorizationEvent = make(map[string][]GitHubAppAuthorizationEventHandleFunc)
	}
	g.onGitHubAppAuthorizationEvent[GitHubAppAuthorizationEventRevokedAction] = append(
		g.onGitHubAppAuthorizationEvent[GitHubAppAuthorizationEventRevokedAction],
		callbacks...,
	)
}

// SetOnGitHubAppAuthorizationEventRevoked registers callbacks listening to events of type github.GitHubAppAuthorizationEvent and action 'revoked'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnGitHubAppAuthorizationEventRevokedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#github_app_authorization
func (g *EventHandler) SetOnGitHubAppAuthorizationEventRevoked(callbacks ...GitHubAppAuthorizationEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onGitHubAppAuthorizationEvent == nil {
		g.onGitHubAppAuthorizationEvent = make(map[string][]GitHubAppAuthorizationEventHandleFunc)
	}
	g.onGitHubAppAuthorizationEvent[GitHubAppAuthorizationEventRevokedAction] = callbacks
}

func (g *EventHandler) handleGitHubAppAuthorizationEventRevoked(deliveryID string, eventName string, event *github.GitHubAppAuthorizationEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if GitHubAppAuthorizationEventRevokedAction != *event.Action {
		return fmt.Errorf(
			"handleGitHubAppAuthorizationEventRevoked() called with wrong action, want %s, got %s",
			GitHubAppAuthorizationEventRevokedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		GitHubAppAuthorizationEventRevokedAction,
		GitHubAppAuthorizationEventAnyAction,
	} {
		if _, ok := g.onGitHubAppAuthorizationEvent[action]; ok {
			for _, h := range g.onGitHubAppAuthorizationEvent[action] {
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

// OnGitHubAppAuthorizationEventAny registers callbacks listening to any events of type github.GitHubAppAuthorizationEvent
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnGitHubAppAuthorizationEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#github_app_authorization
func (g *EventHandler) OnGitHubAppAuthorizationEventAny(callbacks ...GitHubAppAuthorizationEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onGitHubAppAuthorizationEvent == nil {
		g.onGitHubAppAuthorizationEvent = make(map[string][]GitHubAppAuthorizationEventHandleFunc)
	}
	g.onGitHubAppAuthorizationEvent[GitHubAppAuthorizationEventAnyAction] = append(
		g.onGitHubAppAuthorizationEvent[GitHubAppAuthorizationEventAnyAction],
		callbacks...,
	)
}

// SetOnGitHubAppAuthorizationEventAny registers callbacks listening to any events of type github.GitHubAppAuthorizationEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnGitHubAppAuthorizationEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#github_app_authorization
func (g *EventHandler) SetOnGitHubAppAuthorizationEventAny(callbacks ...GitHubAppAuthorizationEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onGitHubAppAuthorizationEvent == nil {
		g.onGitHubAppAuthorizationEvent = make(map[string][]GitHubAppAuthorizationEventHandleFunc)
	}
	g.onGitHubAppAuthorizationEvent[GitHubAppAuthorizationEventAnyAction] = callbacks
}

func (g *EventHandler) handleGitHubAppAuthorizationEventAny(deliveryID string, eventName string, event *github.GitHubAppAuthorizationEvent) error {
	if event == nil {
		return fmt.Errorf("event was empty or nil")
	}
	if _, ok := g.onGitHubAppAuthorizationEvent[GitHubAppAuthorizationEventAnyAction]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onGitHubAppAuthorizationEvent[GitHubAppAuthorizationEventAnyAction] {
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

// GitHubAppAuthorizationEvent handles github.GitHubAppAuthorizationEvent.
//
// Callbacks are executed in the following order:
//
// 1) All callbacks registered with OnBeforeAny are executed in parallel.
// 2) All callbacks registered with OnGitHubAppAuthorizationEvent... are executed in parallel in case the Event has actions.
// 3) All callbacks registered with OnAfterAny are executed in parallel.
//
// on any error all callbacks registered with OnError are executed in parallel.
func (g *EventHandler) GitHubAppAuthorizationEvent(deliveryID string, eventName string, event *github.GitHubAppAuthorizationEvent) error {

	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	action := *event.Action

	err := g.handleBeforeAny(deliveryID, eventName, event)
	if err != nil {
		return g.handleError(deliveryID, eventName, event, err)
	}

	switch action {

	case GitHubAppAuthorizationEventRevokedAction:
		err := g.handleGitHubAppAuthorizationEventRevoked(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	default:
		err := g.handleGitHubAppAuthorizationEventAny(deliveryID, eventName, event)
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
