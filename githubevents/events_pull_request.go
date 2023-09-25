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
	// PullRequestEvent is the event name of github.PullRequestEvent's
	PullRequestEvent = "pull_request"

	// PullRequestEventAnyAction is used to identify callbacks
	// listening to all events of type github.PullRequestEvent
	PullRequestEventAnyAction = "*"

	// PullRequestEventAssignedAction is used to identify callbacks
	// listening to events of type github.PullRequestEvent and action "assigned"
	PullRequestEventAssignedAction = "assigned"

	// PullRequestEventAutoMergeDisabledAction is used to identify callbacks
	// listening to events of type github.PullRequestEvent and action "auto_merge_disabled"
	PullRequestEventAutoMergeDisabledAction = "auto_merge_disabled"

	// PullRequestEventAutoMergeEnabledAction is used to identify callbacks
	// listening to events of type github.PullRequestEvent and action "auto_merge_enabled"
	PullRequestEventAutoMergeEnabledAction = "auto_merge_enabled"

	// PullRequestEventClosedAction is used to identify callbacks
	// listening to events of type github.PullRequestEvent and action "closed"
	PullRequestEventClosedAction = "closed"

	// PullRequestEventConvertedToDraftAction is used to identify callbacks
	// listening to events of type github.PullRequestEvent and action "converted_to_draft"
	PullRequestEventConvertedToDraftAction = "converted_to_draft"

	// PullRequestEventEditedAction is used to identify callbacks
	// listening to events of type github.PullRequestEvent and action "edited"
	PullRequestEventEditedAction = "edited"

	// PullRequestEventLabeledAction is used to identify callbacks
	// listening to events of type github.PullRequestEvent and action "labeled"
	PullRequestEventLabeledAction = "labeled"

	// PullRequestEventLockedAction is used to identify callbacks
	// listening to events of type github.PullRequestEvent and action "locked"
	PullRequestEventLockedAction = "locked"

	// PullRequestEventOpenedAction is used to identify callbacks
	// listening to events of type github.PullRequestEvent and action "opened"
	PullRequestEventOpenedAction = "opened"

	// PullRequestEventReadyForReviewAction is used to identify callbacks
	// listening to events of type github.PullRequestEvent and action "ready_for_review"
	PullRequestEventReadyForReviewAction = "ready_for_review"

	// PullRequestEventReopenedAction is used to identify callbacks
	// listening to events of type github.PullRequestEvent and action "reopened"
	PullRequestEventReopenedAction = "reopened"

	// PullRequestEventReviewRequestRemovedAction is used to identify callbacks
	// listening to events of type github.PullRequestEvent and action "review_request_removed"
	PullRequestEventReviewRequestRemovedAction = "review_request_removed"

	// PullRequestEventReviewRequestedAction is used to identify callbacks
	// listening to events of type github.PullRequestEvent and action "review_requested"
	PullRequestEventReviewRequestedAction = "review_requested"

	// PullRequestEventSynchronizeAction is used to identify callbacks
	// listening to events of type github.PullRequestEvent and action "synchronize"
	PullRequestEventSynchronizeAction = "synchronize"

	// PullRequestEventUnassignedAction is used to identify callbacks
	// listening to events of type github.PullRequestEvent and action "unassigned"
	PullRequestEventUnassignedAction = "unassigned"

	// PullRequestEventUnlabeledAction is used to identify callbacks
	// listening to events of type github.PullRequestEvent and action "unlabeled"
	PullRequestEventUnlabeledAction = "unlabeled"

	// PullRequestEventUnlockedAction is used to identify callbacks
	// listening to events of type github.PullRequestEvent and action "unlocked"
	PullRequestEventUnlockedAction = "unlocked"
)

// PullRequestEventHandleFunc represents a callback function triggered on github.PullRequestEvent's.
// 'deliveryID' (type: string) is the unique webhook delivery ID.
// 'eventName' (type: string) is the name of the event.
// 'event' (type: *github.PullRequestEvent) is the webhook payload.
type PullRequestEventHandleFunc func(deliveryID string, eventName string, event *github.PullRequestEvent) error

// OnPullRequestEventAssigned registers callbacks listening to events of type github.PullRequestEvent and action 'assigned'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnPullRequestEventAssigned must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#pull_request
func (g *EventHandler) OnPullRequestEventAssigned(callbacks ...PullRequestEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestEvent == nil {
		g.onPullRequestEvent = make(map[string][]PullRequestEventHandleFunc)
	}
	g.onPullRequestEvent[PullRequestEventAssignedAction] = append(
		g.onPullRequestEvent[PullRequestEventAssignedAction],
		callbacks...,
	)
}

// SetOnPullRequestEventAssigned registers callbacks listening to events of type github.PullRequestEvent and action 'assigned'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnPullRequestEventAssignedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#pull_request
func (g *EventHandler) SetOnPullRequestEventAssigned(callbacks ...PullRequestEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestEvent == nil {
		g.onPullRequestEvent = make(map[string][]PullRequestEventHandleFunc)
	}
	g.onPullRequestEvent[PullRequestEventAssignedAction] = callbacks
}

func (g *EventHandler) handlePullRequestEventAssigned(deliveryID string, eventName string, event *github.PullRequestEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if PullRequestEventAssignedAction != *event.Action {
		return fmt.Errorf(
			"handlePullRequestEventAssigned() called with wrong action, want %s, got %s",
			PullRequestEventAssignedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		PullRequestEventAssignedAction,
		PullRequestEventAnyAction,
	} {
		if _, ok := g.onPullRequestEvent[action]; ok {
			for _, h := range g.onPullRequestEvent[action] {
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

// OnPullRequestEventAutoMergeDisabled registers callbacks listening to events of type github.PullRequestEvent and action 'auto_merge_disabled'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnPullRequestEventAutoMergeDisabled must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#pull_request
func (g *EventHandler) OnPullRequestEventAutoMergeDisabled(callbacks ...PullRequestEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestEvent == nil {
		g.onPullRequestEvent = make(map[string][]PullRequestEventHandleFunc)
	}
	g.onPullRequestEvent[PullRequestEventAutoMergeDisabledAction] = append(
		g.onPullRequestEvent[PullRequestEventAutoMergeDisabledAction],
		callbacks...,
	)
}

// SetOnPullRequestEventAutoMergeDisabled registers callbacks listening to events of type github.PullRequestEvent and action 'auto_merge_disabled'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnPullRequestEventAutoMergeDisabledAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#pull_request
func (g *EventHandler) SetOnPullRequestEventAutoMergeDisabled(callbacks ...PullRequestEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestEvent == nil {
		g.onPullRequestEvent = make(map[string][]PullRequestEventHandleFunc)
	}
	g.onPullRequestEvent[PullRequestEventAutoMergeDisabledAction] = callbacks
}

func (g *EventHandler) handlePullRequestEventAutoMergeDisabled(deliveryID string, eventName string, event *github.PullRequestEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if PullRequestEventAutoMergeDisabledAction != *event.Action {
		return fmt.Errorf(
			"handlePullRequestEventAutoMergeDisabled() called with wrong action, want %s, got %s",
			PullRequestEventAutoMergeDisabledAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		PullRequestEventAutoMergeDisabledAction,
		PullRequestEventAnyAction,
	} {
		if _, ok := g.onPullRequestEvent[action]; ok {
			for _, h := range g.onPullRequestEvent[action] {
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

// OnPullRequestEventAutoMergeEnabled registers callbacks listening to events of type github.PullRequestEvent and action 'auto_merge_enabled'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnPullRequestEventAutoMergeEnabled must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#pull_request
func (g *EventHandler) OnPullRequestEventAutoMergeEnabled(callbacks ...PullRequestEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestEvent == nil {
		g.onPullRequestEvent = make(map[string][]PullRequestEventHandleFunc)
	}
	g.onPullRequestEvent[PullRequestEventAutoMergeEnabledAction] = append(
		g.onPullRequestEvent[PullRequestEventAutoMergeEnabledAction],
		callbacks...,
	)
}

// SetOnPullRequestEventAutoMergeEnabled registers callbacks listening to events of type github.PullRequestEvent and action 'auto_merge_enabled'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnPullRequestEventAutoMergeEnabledAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#pull_request
func (g *EventHandler) SetOnPullRequestEventAutoMergeEnabled(callbacks ...PullRequestEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestEvent == nil {
		g.onPullRequestEvent = make(map[string][]PullRequestEventHandleFunc)
	}
	g.onPullRequestEvent[PullRequestEventAutoMergeEnabledAction] = callbacks
}

func (g *EventHandler) handlePullRequestEventAutoMergeEnabled(deliveryID string, eventName string, event *github.PullRequestEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if PullRequestEventAutoMergeEnabledAction != *event.Action {
		return fmt.Errorf(
			"handlePullRequestEventAutoMergeEnabled() called with wrong action, want %s, got %s",
			PullRequestEventAutoMergeEnabledAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		PullRequestEventAutoMergeEnabledAction,
		PullRequestEventAnyAction,
	} {
		if _, ok := g.onPullRequestEvent[action]; ok {
			for _, h := range g.onPullRequestEvent[action] {
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

// OnPullRequestEventClosed registers callbacks listening to events of type github.PullRequestEvent and action 'closed'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnPullRequestEventClosed must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#pull_request
func (g *EventHandler) OnPullRequestEventClosed(callbacks ...PullRequestEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestEvent == nil {
		g.onPullRequestEvent = make(map[string][]PullRequestEventHandleFunc)
	}
	g.onPullRequestEvent[PullRequestEventClosedAction] = append(
		g.onPullRequestEvent[PullRequestEventClosedAction],
		callbacks...,
	)
}

// SetOnPullRequestEventClosed registers callbacks listening to events of type github.PullRequestEvent and action 'closed'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnPullRequestEventClosedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#pull_request
func (g *EventHandler) SetOnPullRequestEventClosed(callbacks ...PullRequestEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestEvent == nil {
		g.onPullRequestEvent = make(map[string][]PullRequestEventHandleFunc)
	}
	g.onPullRequestEvent[PullRequestEventClosedAction] = callbacks
}

func (g *EventHandler) handlePullRequestEventClosed(deliveryID string, eventName string, event *github.PullRequestEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if PullRequestEventClosedAction != *event.Action {
		return fmt.Errorf(
			"handlePullRequestEventClosed() called with wrong action, want %s, got %s",
			PullRequestEventClosedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		PullRequestEventClosedAction,
		PullRequestEventAnyAction,
	} {
		if _, ok := g.onPullRequestEvent[action]; ok {
			for _, h := range g.onPullRequestEvent[action] {
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

// OnPullRequestEventConvertedToDraft registers callbacks listening to events of type github.PullRequestEvent and action 'converted_to_draft'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnPullRequestEventConvertedToDraft must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#pull_request
func (g *EventHandler) OnPullRequestEventConvertedToDraft(callbacks ...PullRequestEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestEvent == nil {
		g.onPullRequestEvent = make(map[string][]PullRequestEventHandleFunc)
	}
	g.onPullRequestEvent[PullRequestEventConvertedToDraftAction] = append(
		g.onPullRequestEvent[PullRequestEventConvertedToDraftAction],
		callbacks...,
	)
}

// SetOnPullRequestEventConvertedToDraft registers callbacks listening to events of type github.PullRequestEvent and action 'converted_to_draft'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnPullRequestEventConvertedToDraftAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#pull_request
func (g *EventHandler) SetOnPullRequestEventConvertedToDraft(callbacks ...PullRequestEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestEvent == nil {
		g.onPullRequestEvent = make(map[string][]PullRequestEventHandleFunc)
	}
	g.onPullRequestEvent[PullRequestEventConvertedToDraftAction] = callbacks
}

func (g *EventHandler) handlePullRequestEventConvertedToDraft(deliveryID string, eventName string, event *github.PullRequestEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if PullRequestEventConvertedToDraftAction != *event.Action {
		return fmt.Errorf(
			"handlePullRequestEventConvertedToDraft() called with wrong action, want %s, got %s",
			PullRequestEventConvertedToDraftAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		PullRequestEventConvertedToDraftAction,
		PullRequestEventAnyAction,
	} {
		if _, ok := g.onPullRequestEvent[action]; ok {
			for _, h := range g.onPullRequestEvent[action] {
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

// OnPullRequestEventEdited registers callbacks listening to events of type github.PullRequestEvent and action 'edited'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnPullRequestEventEdited must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#pull_request
func (g *EventHandler) OnPullRequestEventEdited(callbacks ...PullRequestEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestEvent == nil {
		g.onPullRequestEvent = make(map[string][]PullRequestEventHandleFunc)
	}
	g.onPullRequestEvent[PullRequestEventEditedAction] = append(
		g.onPullRequestEvent[PullRequestEventEditedAction],
		callbacks...,
	)
}

// SetOnPullRequestEventEdited registers callbacks listening to events of type github.PullRequestEvent and action 'edited'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnPullRequestEventEditedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#pull_request
func (g *EventHandler) SetOnPullRequestEventEdited(callbacks ...PullRequestEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestEvent == nil {
		g.onPullRequestEvent = make(map[string][]PullRequestEventHandleFunc)
	}
	g.onPullRequestEvent[PullRequestEventEditedAction] = callbacks
}

func (g *EventHandler) handlePullRequestEventEdited(deliveryID string, eventName string, event *github.PullRequestEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if PullRequestEventEditedAction != *event.Action {
		return fmt.Errorf(
			"handlePullRequestEventEdited() called with wrong action, want %s, got %s",
			PullRequestEventEditedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		PullRequestEventEditedAction,
		PullRequestEventAnyAction,
	} {
		if _, ok := g.onPullRequestEvent[action]; ok {
			for _, h := range g.onPullRequestEvent[action] {
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

// OnPullRequestEventLabeled registers callbacks listening to events of type github.PullRequestEvent and action 'labeled'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnPullRequestEventLabeled must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#pull_request
func (g *EventHandler) OnPullRequestEventLabeled(callbacks ...PullRequestEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestEvent == nil {
		g.onPullRequestEvent = make(map[string][]PullRequestEventHandleFunc)
	}
	g.onPullRequestEvent[PullRequestEventLabeledAction] = append(
		g.onPullRequestEvent[PullRequestEventLabeledAction],
		callbacks...,
	)
}

// SetOnPullRequestEventLabeled registers callbacks listening to events of type github.PullRequestEvent and action 'labeled'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnPullRequestEventLabeledAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#pull_request
func (g *EventHandler) SetOnPullRequestEventLabeled(callbacks ...PullRequestEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestEvent == nil {
		g.onPullRequestEvent = make(map[string][]PullRequestEventHandleFunc)
	}
	g.onPullRequestEvent[PullRequestEventLabeledAction] = callbacks
}

func (g *EventHandler) handlePullRequestEventLabeled(deliveryID string, eventName string, event *github.PullRequestEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if PullRequestEventLabeledAction != *event.Action {
		return fmt.Errorf(
			"handlePullRequestEventLabeled() called with wrong action, want %s, got %s",
			PullRequestEventLabeledAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		PullRequestEventLabeledAction,
		PullRequestEventAnyAction,
	} {
		if _, ok := g.onPullRequestEvent[action]; ok {
			for _, h := range g.onPullRequestEvent[action] {
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

// OnPullRequestEventLocked registers callbacks listening to events of type github.PullRequestEvent and action 'locked'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnPullRequestEventLocked must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#pull_request
func (g *EventHandler) OnPullRequestEventLocked(callbacks ...PullRequestEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestEvent == nil {
		g.onPullRequestEvent = make(map[string][]PullRequestEventHandleFunc)
	}
	g.onPullRequestEvent[PullRequestEventLockedAction] = append(
		g.onPullRequestEvent[PullRequestEventLockedAction],
		callbacks...,
	)
}

// SetOnPullRequestEventLocked registers callbacks listening to events of type github.PullRequestEvent and action 'locked'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnPullRequestEventLockedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#pull_request
func (g *EventHandler) SetOnPullRequestEventLocked(callbacks ...PullRequestEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestEvent == nil {
		g.onPullRequestEvent = make(map[string][]PullRequestEventHandleFunc)
	}
	g.onPullRequestEvent[PullRequestEventLockedAction] = callbacks
}

func (g *EventHandler) handlePullRequestEventLocked(deliveryID string, eventName string, event *github.PullRequestEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if PullRequestEventLockedAction != *event.Action {
		return fmt.Errorf(
			"handlePullRequestEventLocked() called with wrong action, want %s, got %s",
			PullRequestEventLockedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		PullRequestEventLockedAction,
		PullRequestEventAnyAction,
	} {
		if _, ok := g.onPullRequestEvent[action]; ok {
			for _, h := range g.onPullRequestEvent[action] {
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

// OnPullRequestEventOpened registers callbacks listening to events of type github.PullRequestEvent and action 'opened'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnPullRequestEventOpened must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#pull_request
func (g *EventHandler) OnPullRequestEventOpened(callbacks ...PullRequestEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestEvent == nil {
		g.onPullRequestEvent = make(map[string][]PullRequestEventHandleFunc)
	}
	g.onPullRequestEvent[PullRequestEventOpenedAction] = append(
		g.onPullRequestEvent[PullRequestEventOpenedAction],
		callbacks...,
	)
}

// SetOnPullRequestEventOpened registers callbacks listening to events of type github.PullRequestEvent and action 'opened'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnPullRequestEventOpenedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#pull_request
func (g *EventHandler) SetOnPullRequestEventOpened(callbacks ...PullRequestEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestEvent == nil {
		g.onPullRequestEvent = make(map[string][]PullRequestEventHandleFunc)
	}
	g.onPullRequestEvent[PullRequestEventOpenedAction] = callbacks
}

func (g *EventHandler) handlePullRequestEventOpened(deliveryID string, eventName string, event *github.PullRequestEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if PullRequestEventOpenedAction != *event.Action {
		return fmt.Errorf(
			"handlePullRequestEventOpened() called with wrong action, want %s, got %s",
			PullRequestEventOpenedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		PullRequestEventOpenedAction,
		PullRequestEventAnyAction,
	} {
		if _, ok := g.onPullRequestEvent[action]; ok {
			for _, h := range g.onPullRequestEvent[action] {
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

// OnPullRequestEventReadyForReview registers callbacks listening to events of type github.PullRequestEvent and action 'ready_for_review'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnPullRequestEventReadyForReview must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#pull_request
func (g *EventHandler) OnPullRequestEventReadyForReview(callbacks ...PullRequestEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestEvent == nil {
		g.onPullRequestEvent = make(map[string][]PullRequestEventHandleFunc)
	}
	g.onPullRequestEvent[PullRequestEventReadyForReviewAction] = append(
		g.onPullRequestEvent[PullRequestEventReadyForReviewAction],
		callbacks...,
	)
}

// SetOnPullRequestEventReadyForReview registers callbacks listening to events of type github.PullRequestEvent and action 'ready_for_review'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnPullRequestEventReadyForReviewAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#pull_request
func (g *EventHandler) SetOnPullRequestEventReadyForReview(callbacks ...PullRequestEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestEvent == nil {
		g.onPullRequestEvent = make(map[string][]PullRequestEventHandleFunc)
	}
	g.onPullRequestEvent[PullRequestEventReadyForReviewAction] = callbacks
}

func (g *EventHandler) handlePullRequestEventReadyForReview(deliveryID string, eventName string, event *github.PullRequestEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if PullRequestEventReadyForReviewAction != *event.Action {
		return fmt.Errorf(
			"handlePullRequestEventReadyForReview() called with wrong action, want %s, got %s",
			PullRequestEventReadyForReviewAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		PullRequestEventReadyForReviewAction,
		PullRequestEventAnyAction,
	} {
		if _, ok := g.onPullRequestEvent[action]; ok {
			for _, h := range g.onPullRequestEvent[action] {
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

// OnPullRequestEventReopened registers callbacks listening to events of type github.PullRequestEvent and action 'reopened'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnPullRequestEventReopened must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#pull_request
func (g *EventHandler) OnPullRequestEventReopened(callbacks ...PullRequestEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestEvent == nil {
		g.onPullRequestEvent = make(map[string][]PullRequestEventHandleFunc)
	}
	g.onPullRequestEvent[PullRequestEventReopenedAction] = append(
		g.onPullRequestEvent[PullRequestEventReopenedAction],
		callbacks...,
	)
}

// SetOnPullRequestEventReopened registers callbacks listening to events of type github.PullRequestEvent and action 'reopened'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnPullRequestEventReopenedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#pull_request
func (g *EventHandler) SetOnPullRequestEventReopened(callbacks ...PullRequestEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestEvent == nil {
		g.onPullRequestEvent = make(map[string][]PullRequestEventHandleFunc)
	}
	g.onPullRequestEvent[PullRequestEventReopenedAction] = callbacks
}

func (g *EventHandler) handlePullRequestEventReopened(deliveryID string, eventName string, event *github.PullRequestEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if PullRequestEventReopenedAction != *event.Action {
		return fmt.Errorf(
			"handlePullRequestEventReopened() called with wrong action, want %s, got %s",
			PullRequestEventReopenedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		PullRequestEventReopenedAction,
		PullRequestEventAnyAction,
	} {
		if _, ok := g.onPullRequestEvent[action]; ok {
			for _, h := range g.onPullRequestEvent[action] {
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

// OnPullRequestEventReviewRequestRemoved registers callbacks listening to events of type github.PullRequestEvent and action 'review_request_removed'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnPullRequestEventReviewRequestRemoved must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#pull_request
func (g *EventHandler) OnPullRequestEventReviewRequestRemoved(callbacks ...PullRequestEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestEvent == nil {
		g.onPullRequestEvent = make(map[string][]PullRequestEventHandleFunc)
	}
	g.onPullRequestEvent[PullRequestEventReviewRequestRemovedAction] = append(
		g.onPullRequestEvent[PullRequestEventReviewRequestRemovedAction],
		callbacks...,
	)
}

// SetOnPullRequestEventReviewRequestRemoved registers callbacks listening to events of type github.PullRequestEvent and action 'review_request_removed'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnPullRequestEventReviewRequestRemovedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#pull_request
func (g *EventHandler) SetOnPullRequestEventReviewRequestRemoved(callbacks ...PullRequestEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestEvent == nil {
		g.onPullRequestEvent = make(map[string][]PullRequestEventHandleFunc)
	}
	g.onPullRequestEvent[PullRequestEventReviewRequestRemovedAction] = callbacks
}

func (g *EventHandler) handlePullRequestEventReviewRequestRemoved(deliveryID string, eventName string, event *github.PullRequestEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if PullRequestEventReviewRequestRemovedAction != *event.Action {
		return fmt.Errorf(
			"handlePullRequestEventReviewRequestRemoved() called with wrong action, want %s, got %s",
			PullRequestEventReviewRequestRemovedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		PullRequestEventReviewRequestRemovedAction,
		PullRequestEventAnyAction,
	} {
		if _, ok := g.onPullRequestEvent[action]; ok {
			for _, h := range g.onPullRequestEvent[action] {
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

// OnPullRequestEventReviewRequested registers callbacks listening to events of type github.PullRequestEvent and action 'review_requested'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnPullRequestEventReviewRequested must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#pull_request
func (g *EventHandler) OnPullRequestEventReviewRequested(callbacks ...PullRequestEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestEvent == nil {
		g.onPullRequestEvent = make(map[string][]PullRequestEventHandleFunc)
	}
	g.onPullRequestEvent[PullRequestEventReviewRequestedAction] = append(
		g.onPullRequestEvent[PullRequestEventReviewRequestedAction],
		callbacks...,
	)
}

// SetOnPullRequestEventReviewRequested registers callbacks listening to events of type github.PullRequestEvent and action 'review_requested'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnPullRequestEventReviewRequestedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#pull_request
func (g *EventHandler) SetOnPullRequestEventReviewRequested(callbacks ...PullRequestEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestEvent == nil {
		g.onPullRequestEvent = make(map[string][]PullRequestEventHandleFunc)
	}
	g.onPullRequestEvent[PullRequestEventReviewRequestedAction] = callbacks
}

func (g *EventHandler) handlePullRequestEventReviewRequested(deliveryID string, eventName string, event *github.PullRequestEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if PullRequestEventReviewRequestedAction != *event.Action {
		return fmt.Errorf(
			"handlePullRequestEventReviewRequested() called with wrong action, want %s, got %s",
			PullRequestEventReviewRequestedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		PullRequestEventReviewRequestedAction,
		PullRequestEventAnyAction,
	} {
		if _, ok := g.onPullRequestEvent[action]; ok {
			for _, h := range g.onPullRequestEvent[action] {
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

// OnPullRequestEventSynchronize registers callbacks listening to events of type github.PullRequestEvent and action 'synchronize'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnPullRequestEventSynchronize must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#pull_request
func (g *EventHandler) OnPullRequestEventSynchronize(callbacks ...PullRequestEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestEvent == nil {
		g.onPullRequestEvent = make(map[string][]PullRequestEventHandleFunc)
	}
	g.onPullRequestEvent[PullRequestEventSynchronizeAction] = append(
		g.onPullRequestEvent[PullRequestEventSynchronizeAction],
		callbacks...,
	)
}

// SetOnPullRequestEventSynchronize registers callbacks listening to events of type github.PullRequestEvent and action 'synchronize'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnPullRequestEventSynchronizeAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#pull_request
func (g *EventHandler) SetOnPullRequestEventSynchronize(callbacks ...PullRequestEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestEvent == nil {
		g.onPullRequestEvent = make(map[string][]PullRequestEventHandleFunc)
	}
	g.onPullRequestEvent[PullRequestEventSynchronizeAction] = callbacks
}

func (g *EventHandler) handlePullRequestEventSynchronize(deliveryID string, eventName string, event *github.PullRequestEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if PullRequestEventSynchronizeAction != *event.Action {
		return fmt.Errorf(
			"handlePullRequestEventSynchronize() called with wrong action, want %s, got %s",
			PullRequestEventSynchronizeAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		PullRequestEventSynchronizeAction,
		PullRequestEventAnyAction,
	} {
		if _, ok := g.onPullRequestEvent[action]; ok {
			for _, h := range g.onPullRequestEvent[action] {
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

// OnPullRequestEventUnassigned registers callbacks listening to events of type github.PullRequestEvent and action 'unassigned'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnPullRequestEventUnassigned must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#pull_request
func (g *EventHandler) OnPullRequestEventUnassigned(callbacks ...PullRequestEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestEvent == nil {
		g.onPullRequestEvent = make(map[string][]PullRequestEventHandleFunc)
	}
	g.onPullRequestEvent[PullRequestEventUnassignedAction] = append(
		g.onPullRequestEvent[PullRequestEventUnassignedAction],
		callbacks...,
	)
}

// SetOnPullRequestEventUnassigned registers callbacks listening to events of type github.PullRequestEvent and action 'unassigned'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnPullRequestEventUnassignedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#pull_request
func (g *EventHandler) SetOnPullRequestEventUnassigned(callbacks ...PullRequestEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestEvent == nil {
		g.onPullRequestEvent = make(map[string][]PullRequestEventHandleFunc)
	}
	g.onPullRequestEvent[PullRequestEventUnassignedAction] = callbacks
}

func (g *EventHandler) handlePullRequestEventUnassigned(deliveryID string, eventName string, event *github.PullRequestEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if PullRequestEventUnassignedAction != *event.Action {
		return fmt.Errorf(
			"handlePullRequestEventUnassigned() called with wrong action, want %s, got %s",
			PullRequestEventUnassignedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		PullRequestEventUnassignedAction,
		PullRequestEventAnyAction,
	} {
		if _, ok := g.onPullRequestEvent[action]; ok {
			for _, h := range g.onPullRequestEvent[action] {
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

// OnPullRequestEventUnlabeled registers callbacks listening to events of type github.PullRequestEvent and action 'unlabeled'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnPullRequestEventUnlabeled must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#pull_request
func (g *EventHandler) OnPullRequestEventUnlabeled(callbacks ...PullRequestEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestEvent == nil {
		g.onPullRequestEvent = make(map[string][]PullRequestEventHandleFunc)
	}
	g.onPullRequestEvent[PullRequestEventUnlabeledAction] = append(
		g.onPullRequestEvent[PullRequestEventUnlabeledAction],
		callbacks...,
	)
}

// SetOnPullRequestEventUnlabeled registers callbacks listening to events of type github.PullRequestEvent and action 'unlabeled'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnPullRequestEventUnlabeledAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#pull_request
func (g *EventHandler) SetOnPullRequestEventUnlabeled(callbacks ...PullRequestEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestEvent == nil {
		g.onPullRequestEvent = make(map[string][]PullRequestEventHandleFunc)
	}
	g.onPullRequestEvent[PullRequestEventUnlabeledAction] = callbacks
}

func (g *EventHandler) handlePullRequestEventUnlabeled(deliveryID string, eventName string, event *github.PullRequestEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if PullRequestEventUnlabeledAction != *event.Action {
		return fmt.Errorf(
			"handlePullRequestEventUnlabeled() called with wrong action, want %s, got %s",
			PullRequestEventUnlabeledAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		PullRequestEventUnlabeledAction,
		PullRequestEventAnyAction,
	} {
		if _, ok := g.onPullRequestEvent[action]; ok {
			for _, h := range g.onPullRequestEvent[action] {
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

// OnPullRequestEventUnlocked registers callbacks listening to events of type github.PullRequestEvent and action 'unlocked'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnPullRequestEventUnlocked must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#pull_request
func (g *EventHandler) OnPullRequestEventUnlocked(callbacks ...PullRequestEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestEvent == nil {
		g.onPullRequestEvent = make(map[string][]PullRequestEventHandleFunc)
	}
	g.onPullRequestEvent[PullRequestEventUnlockedAction] = append(
		g.onPullRequestEvent[PullRequestEventUnlockedAction],
		callbacks...,
	)
}

// SetOnPullRequestEventUnlocked registers callbacks listening to events of type github.PullRequestEvent and action 'unlocked'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnPullRequestEventUnlockedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#pull_request
func (g *EventHandler) SetOnPullRequestEventUnlocked(callbacks ...PullRequestEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestEvent == nil {
		g.onPullRequestEvent = make(map[string][]PullRequestEventHandleFunc)
	}
	g.onPullRequestEvent[PullRequestEventUnlockedAction] = callbacks
}

func (g *EventHandler) handlePullRequestEventUnlocked(deliveryID string, eventName string, event *github.PullRequestEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if PullRequestEventUnlockedAction != *event.Action {
		return fmt.Errorf(
			"handlePullRequestEventUnlocked() called with wrong action, want %s, got %s",
			PullRequestEventUnlockedAction,
			*event.Action,
		)
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		PullRequestEventUnlockedAction,
		PullRequestEventAnyAction,
	} {
		if _, ok := g.onPullRequestEvent[action]; ok {
			for _, h := range g.onPullRequestEvent[action] {
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

// OnPullRequestEventAny registers callbacks listening to any events of type github.PullRequestEvent
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnPullRequestEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#pull_request
func (g *EventHandler) OnPullRequestEventAny(callbacks ...PullRequestEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestEvent == nil {
		g.onPullRequestEvent = make(map[string][]PullRequestEventHandleFunc)
	}
	g.onPullRequestEvent[PullRequestEventAnyAction] = append(
		g.onPullRequestEvent[PullRequestEventAnyAction],
		callbacks...,
	)
}

// SetOnPullRequestEventAny registers callbacks listening to any events of type github.PullRequestEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnPullRequestEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#pull_request
func (g *EventHandler) SetOnPullRequestEventAny(callbacks ...PullRequestEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestEvent == nil {
		g.onPullRequestEvent = make(map[string][]PullRequestEventHandleFunc)
	}
	g.onPullRequestEvent[PullRequestEventAnyAction] = callbacks
}

func (g *EventHandler) handlePullRequestEventAny(deliveryID string, eventName string, event *github.PullRequestEvent) error {
	if event == nil {
		return fmt.Errorf("event was empty or nil")
	}
	if _, ok := g.onPullRequestEvent[PullRequestEventAnyAction]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onPullRequestEvent[PullRequestEventAnyAction] {
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

// PullRequestEvent handles github.PullRequestEvent.
//
// Callbacks are executed in the following order:
//
// 1) All callbacks registered with OnBeforeAny are executed in parallel.
// 2) All callbacks registered with OnPullRequestEvent... are executed in parallel in case the Event has actions.
// 3) All callbacks registered with OnAfterAny are executed in parallel.
//
// on any error all callbacks registered with OnError are executed in parallel.
func (g *EventHandler) PullRequestEvent(deliveryID string, eventName string, event *github.PullRequestEvent) error {

	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	action := *event.Action

	err := g.handleBeforeAny(deliveryID, eventName, event)
	if err != nil {
		return g.handleError(deliveryID, eventName, event, err)
	}

	switch action {

	case PullRequestEventAssignedAction:
		err := g.handlePullRequestEventAssigned(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case PullRequestEventAutoMergeDisabledAction:
		err := g.handlePullRequestEventAutoMergeDisabled(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case PullRequestEventAutoMergeEnabledAction:
		err := g.handlePullRequestEventAutoMergeEnabled(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case PullRequestEventClosedAction:
		err := g.handlePullRequestEventClosed(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case PullRequestEventConvertedToDraftAction:
		err := g.handlePullRequestEventConvertedToDraft(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case PullRequestEventEditedAction:
		err := g.handlePullRequestEventEdited(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case PullRequestEventLabeledAction:
		err := g.handlePullRequestEventLabeled(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case PullRequestEventLockedAction:
		err := g.handlePullRequestEventLocked(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case PullRequestEventOpenedAction:
		err := g.handlePullRequestEventOpened(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case PullRequestEventReadyForReviewAction:
		err := g.handlePullRequestEventReadyForReview(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case PullRequestEventReopenedAction:
		err := g.handlePullRequestEventReopened(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case PullRequestEventReviewRequestRemovedAction:
		err := g.handlePullRequestEventReviewRequestRemoved(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case PullRequestEventReviewRequestedAction:
		err := g.handlePullRequestEventReviewRequested(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case PullRequestEventSynchronizeAction:
		err := g.handlePullRequestEventSynchronize(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case PullRequestEventUnassignedAction:
		err := g.handlePullRequestEventUnassigned(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case PullRequestEventUnlabeledAction:
		err := g.handlePullRequestEventUnlabeled(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case PullRequestEventUnlockedAction:
		err := g.handlePullRequestEventUnlocked(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	default:
		err := g.handlePullRequestEventAny(deliveryID, eventName, event)
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
