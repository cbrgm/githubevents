package main

var webhookTypesTemplate = `
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

{{ range $_, $webhook := .Webhooks }}

// Actions are used to identify registered callbacks.
const (
	// {{ $webhook.Event }} is the event name of github.{{ $webhook.Event }}'s
	{{ $webhook.Event }} = "{{ $webhook.Name }}"

	// {{ $webhook.Event }}AnyAction is used to identify callbacks
	// listening to all events of type github.{{ $webhook.Event }}
	{{ $webhook.Event }}AnyAction = "*"
	{{ range $_, $action := $webhook.Actions }}
	// {{ $action.Handler }}Action is used to identify callbacks
	// listening to events of type github.{{ $webhook.Event }} and action "{{ $action.Action }}"
	{{ $action.Handler }}Action = "{{ $action.Action }}"
	{{ end }}
)

// {{ $webhook.Event }}HandleFunc represents a callback function triggered on github.{{ $webhook.Event }}'s.
// 'deliveryID' (type: string) is the unique webhook delivery ID.
// 'eventName' (type: string) is the name of the event.
// 'event' (type: *github.{{ $webhook.Event }}) is the webhook payload.
type {{ $webhook.Event }}HandleFunc func(deliveryID string, eventName string, event *github.{{ $webhook.Event }}) error

{{ range $_, $action := $webhook.Actions }}

// On{{ $action.Handler }} registers callbacks listening to events of type github.{{ $webhook.Event }} and action '{{ $action.Action }}'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOn{{ $action.Handler }} must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#{{ $webhook.Name }}
func (g *EventHandler) On{{ $action.Handler }}(callbacks ...{{ $webhook.Event }}HandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.on{{ $webhook.Event }} == nil {
		g.on{{ $webhook.Event }} = make(map[string][]{{ $webhook.Event }}HandleFunc)
	}
	g.on{{ $webhook.Event }}[{{ $action.Handler }}Action] = append(
		g.on{{ $webhook.Event }}[{{ $action.Handler }}Action],
		callbacks...
	)
}

// SetOn{{ $action.Handler }} registers callbacks listening to events of type github.{{ $webhook.Event }} and action '{{ $action.Action }}'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, On{{ $action.Handler }}Any must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#{{ $webhook.Name }}
func (g *EventHandler) SetOn{{ $action.Handler }}(callbacks ...{{ $webhook.Event }}HandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.on{{ $webhook.Event }} == nil {
		g.on{{ $webhook.Event }} = make(map[string][]{{ $webhook.Event }}HandleFunc)
	}
	g.on{{ $webhook.Event }}[{{ $action.Handler }}Action] = callbacks
}

func (g *EventHandler) handle{{ $action.Handler }}(deliveryID string, eventName string, event *github.{{ $webhook.Event }}) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if {{ $action.Handler }}Action != *event.Action {
		return fmt.Errorf(
			"handle{{ $action.Handler }}() called with wrong action, want %s, got %s",
			{{ $action.Handler }}Action,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		{{ $action.Handler }}Action,
		{{ $webhook.Event }}AnyAction,
	} {
		if _, ok := g.on{{ $webhook.Event }}[action]; ok {
			for _, h := range g.on{{ $webhook.Event }}[action] {
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

{{ end }}

// On{{ $webhook.Event }}Any registers callbacks listening to any events of type github.{{ $webhook.Event }}
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOn{{ $webhook.Event }}Any must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#{{ $webhook.Name }}
func (g *EventHandler) On{{ $webhook.Event }}Any(callbacks ...{{ $webhook.Event }}HandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.on{{ $webhook.Event }} == nil {
		g.on{{ $webhook.Event }} = make(map[string][]{{ $webhook.Event }}HandleFunc)
	}
	g.on{{ $webhook.Event }}[{{ $webhook.Event }}AnyAction] = append(
		g.on{{ $webhook.Event }}[{{ $webhook.Event }}AnyAction],
		callbacks...,
	)
}

// SetOn{{ $webhook.Event }}Any registers callbacks listening to any events of type github.{{ $webhook.Event }}
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, On{{ $webhook.Event }}Any must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#{{ $webhook.Name }}
func (g *EventHandler) SetOn{{ $webhook.Event }}Any(callbacks ...{{ $webhook.Event }}HandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.on{{ $webhook.Event }} == nil {
		g.on{{ $webhook.Event }} = make(map[string][]{{ $webhook.Event }}HandleFunc)
	}
	g.on{{ $webhook.Event }}[{{ $webhook.Event }}AnyAction] = callbacks
}

func (g *EventHandler) handle{{ $webhook.Event }}Any(deliveryID string, eventName string, event *github.{{ $webhook.Event }}) error {
	if event == nil {
		return fmt.Errorf("event was empty or nil")
	}
	if _, ok := g.on{{ $webhook.Event }}[{{ $webhook.Event }}AnyAction]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.on{{ $webhook.Event }}[{{ $webhook.Event }}AnyAction] {
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

// {{ $webhook.Event }} handles github.{{ $webhook.Event }}.
//
// Callbacks are executed in the following order:
//
// 1) All callbacks registered with OnBeforeAny are executed in parallel.
// 2) All callbacks registered with On{{ $webhook.Event }}... are executed in parallel in case the Event has actions.
// 3) All callbacks registered with OnAfterAny are executed in parallel.
//
// on any error all callbacks registered with OnError are executed in parallel.
func (g *EventHandler) {{ $webhook.Event }}(deliveryID string, eventName string, event *github.{{ $webhook.Event }}) error {
	{{ if $webhook.HasActions }}
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	action := *event.Action
	{{ else }}
	if event == nil {
		return fmt.Errorf("event action was empty or nil")
	}
	{{ end }}

	err := g.handleBeforeAny(deliveryID, eventName, event)
	if err != nil {
		return g.handleError(deliveryID, eventName, event, err)
	}

	{{ if $webhook.HasActions }}

	switch action {
	{{ range $_, $action := $webhook.Actions }}
	case {{ $action.Handler }}Action:
		err := g.handle{{ $action.Handler }}(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}
	{{ end }}
	default:
		err := g.handle{{ $webhook.Event }}Any(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}
	}
	{{ else }}

	err = g.handle{{ $webhook.Event }}Any(deliveryID, eventName, event)
	if err != nil {
		return g.handleError(deliveryID, eventName, event, err)
	}

	{{ end }}

	err = g.handleAfterAny(deliveryID, eventName, event)
	if err != nil {
		return g.handleError(deliveryID, eventName, event, err)
	}
	return nil
}

{{ end }}
`
