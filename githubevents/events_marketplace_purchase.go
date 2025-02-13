// Copyright 2022 The GithubEvents Authors. All rights reserved.
// Use of this source code is governed by the MIT License
// that can be found in the LICENSE file.

package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"context"
	"fmt"
	"github.com/google/go-github/v69/github"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/sync/errgroup"
)

// Actions are used to identify registered callbacks.
const (
	// MarketplacePurchaseEvent is the event name of github.MarketplacePurchaseEvent's
	MarketplacePurchaseEvent = "marketplace_purchase"

	// MarketplacePurchaseEventAnyAction is used to identify callbacks
	// listening to all events of type github.MarketplacePurchaseEvent
	MarketplacePurchaseEventAnyAction = "*"

	// MarketplacePurchaseEventPurchasedAction is used to identify callbacks
	// listening to events of type github.MarketplacePurchaseEvent and action "purchased"
	MarketplacePurchaseEventPurchasedAction = "purchased"

	// MarketplacePurchaseEventPendingChangeAction is used to identify callbacks
	// listening to events of type github.MarketplacePurchaseEvent and action "pending_change"
	MarketplacePurchaseEventPendingChangeAction = "pending_change"

	// MarketplacePurchaseEventPendingChangeCancelledAction is used to identify callbacks
	// listening to events of type github.MarketplacePurchaseEvent and action "pending_change_cancelled"
	MarketplacePurchaseEventPendingChangeCancelledAction = "pending_change_cancelled"

	// MarketplacePurchaseEventChangedAction is used to identify callbacks
	// listening to events of type github.MarketplacePurchaseEvent and action "changed"
	MarketplacePurchaseEventChangedAction = "changed"

	// MarketplacePurchaseEventCancelledAction is used to identify callbacks
	// listening to events of type github.MarketplacePurchaseEvent and action "cancelled"
	MarketplacePurchaseEventCancelledAction = "cancelled"
)

// MarketplacePurchaseEventHandleFunc represents a callback function triggered on github.MarketplacePurchaseEvent's.
// 'deliveryID' (type: string) is the unique webhook delivery ID.
// 'eventName' (type: string) is the name of the event.
// 'event' (type: *github.MarketplacePurchaseEvent) is the webhook payload.
type MarketplacePurchaseEventHandleFunc func(ctx context.Context, deliveryID string, eventName string, event *github.MarketplacePurchaseEvent) error

// OnMarketplacePurchaseEventPurchased registers callbacks listening to events of type github.MarketplacePurchaseEvent and action 'purchased'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnMarketplacePurchaseEventPurchased must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#marketplace_purchase
func (g *EventHandler) OnMarketplacePurchaseEventPurchased(callbacks ...MarketplacePurchaseEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onMarketplacePurchaseEvent == nil {
		g.onMarketplacePurchaseEvent = make(map[string][]MarketplacePurchaseEventHandleFunc)
	}
	g.onMarketplacePurchaseEvent[MarketplacePurchaseEventPurchasedAction] = append(
		g.onMarketplacePurchaseEvent[MarketplacePurchaseEventPurchasedAction],
		callbacks...,
	)
}

// SetOnMarketplacePurchaseEventPurchased registers callbacks listening to events of type github.MarketplacePurchaseEvent and action 'purchased'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnMarketplacePurchaseEventPurchasedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#marketplace_purchase
func (g *EventHandler) SetOnMarketplacePurchaseEventPurchased(callbacks ...MarketplacePurchaseEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onMarketplacePurchaseEvent == nil {
		g.onMarketplacePurchaseEvent = make(map[string][]MarketplacePurchaseEventHandleFunc)
	}
	g.onMarketplacePurchaseEvent[MarketplacePurchaseEventPurchasedAction] = callbacks
}

func (g *EventHandler) handleMarketplacePurchaseEventPurchased(ctx context.Context, deliveryID string, eventName string, event *github.MarketplacePurchaseEvent) error {
	ctx, span := g.Tracer.Start(ctx, "handleMarketplacePurchaseEventPurchased", trace.WithAttributes(
		attribute.String("deliveryID", deliveryID),
		attribute.String("event", eventName),
	))
	defer span.End()
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if MarketplacePurchaseEventPurchasedAction != *event.Action {
		err := fmt.Errorf(
			"handleMarketplacePurchaseEventPurchased() called with wrong action, want %s, got %s",
			MarketplacePurchaseEventPurchasedAction,
			*event.Action,
		)
		span.SetStatus(codes.Error, err.Error())
		return err
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		MarketplacePurchaseEventPurchasedAction,
		MarketplacePurchaseEventAnyAction,
	} {
		if _, ok := g.onMarketplacePurchaseEvent[action]; ok {
			for _, h := range g.onMarketplacePurchaseEvent[action] {
				handle := h
				eg.Go(func() error {
					err := handle(ctx, deliveryID, eventName, event)
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

// OnMarketplacePurchaseEventPendingChange registers callbacks listening to events of type github.MarketplacePurchaseEvent and action 'pending_change'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnMarketplacePurchaseEventPendingChange must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#marketplace_purchase
func (g *EventHandler) OnMarketplacePurchaseEventPendingChange(callbacks ...MarketplacePurchaseEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onMarketplacePurchaseEvent == nil {
		g.onMarketplacePurchaseEvent = make(map[string][]MarketplacePurchaseEventHandleFunc)
	}
	g.onMarketplacePurchaseEvent[MarketplacePurchaseEventPendingChangeAction] = append(
		g.onMarketplacePurchaseEvent[MarketplacePurchaseEventPendingChangeAction],
		callbacks...,
	)
}

// SetOnMarketplacePurchaseEventPendingChange registers callbacks listening to events of type github.MarketplacePurchaseEvent and action 'pending_change'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnMarketplacePurchaseEventPendingChangeAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#marketplace_purchase
func (g *EventHandler) SetOnMarketplacePurchaseEventPendingChange(callbacks ...MarketplacePurchaseEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onMarketplacePurchaseEvent == nil {
		g.onMarketplacePurchaseEvent = make(map[string][]MarketplacePurchaseEventHandleFunc)
	}
	g.onMarketplacePurchaseEvent[MarketplacePurchaseEventPendingChangeAction] = callbacks
}

func (g *EventHandler) handleMarketplacePurchaseEventPendingChange(ctx context.Context, deliveryID string, eventName string, event *github.MarketplacePurchaseEvent) error {
	ctx, span := g.Tracer.Start(ctx, "handleMarketplacePurchaseEventPendingChange", trace.WithAttributes(
		attribute.String("deliveryID", deliveryID),
		attribute.String("event", eventName),
	))
	defer span.End()
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if MarketplacePurchaseEventPendingChangeAction != *event.Action {
		err := fmt.Errorf(
			"handleMarketplacePurchaseEventPendingChange() called with wrong action, want %s, got %s",
			MarketplacePurchaseEventPendingChangeAction,
			*event.Action,
		)
		span.SetStatus(codes.Error, err.Error())
		return err
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		MarketplacePurchaseEventPendingChangeAction,
		MarketplacePurchaseEventAnyAction,
	} {
		if _, ok := g.onMarketplacePurchaseEvent[action]; ok {
			for _, h := range g.onMarketplacePurchaseEvent[action] {
				handle := h
				eg.Go(func() error {
					err := handle(ctx, deliveryID, eventName, event)
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

// OnMarketplacePurchaseEventPendingChangeCancelled registers callbacks listening to events of type github.MarketplacePurchaseEvent and action 'pending_change_cancelled'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnMarketplacePurchaseEventPendingChangeCancelled must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#marketplace_purchase
func (g *EventHandler) OnMarketplacePurchaseEventPendingChangeCancelled(callbacks ...MarketplacePurchaseEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onMarketplacePurchaseEvent == nil {
		g.onMarketplacePurchaseEvent = make(map[string][]MarketplacePurchaseEventHandleFunc)
	}
	g.onMarketplacePurchaseEvent[MarketplacePurchaseEventPendingChangeCancelledAction] = append(
		g.onMarketplacePurchaseEvent[MarketplacePurchaseEventPendingChangeCancelledAction],
		callbacks...,
	)
}

// SetOnMarketplacePurchaseEventPendingChangeCancelled registers callbacks listening to events of type github.MarketplacePurchaseEvent and action 'pending_change_cancelled'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnMarketplacePurchaseEventPendingChangeCancelledAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#marketplace_purchase
func (g *EventHandler) SetOnMarketplacePurchaseEventPendingChangeCancelled(callbacks ...MarketplacePurchaseEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onMarketplacePurchaseEvent == nil {
		g.onMarketplacePurchaseEvent = make(map[string][]MarketplacePurchaseEventHandleFunc)
	}
	g.onMarketplacePurchaseEvent[MarketplacePurchaseEventPendingChangeCancelledAction] = callbacks
}

func (g *EventHandler) handleMarketplacePurchaseEventPendingChangeCancelled(ctx context.Context, deliveryID string, eventName string, event *github.MarketplacePurchaseEvent) error {
	ctx, span := g.Tracer.Start(ctx, "handleMarketplacePurchaseEventPendingChangeCancelled", trace.WithAttributes(
		attribute.String("deliveryID", deliveryID),
		attribute.String("event", eventName),
	))
	defer span.End()
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if MarketplacePurchaseEventPendingChangeCancelledAction != *event.Action {
		err := fmt.Errorf(
			"handleMarketplacePurchaseEventPendingChangeCancelled() called with wrong action, want %s, got %s",
			MarketplacePurchaseEventPendingChangeCancelledAction,
			*event.Action,
		)
		span.SetStatus(codes.Error, err.Error())
		return err
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		MarketplacePurchaseEventPendingChangeCancelledAction,
		MarketplacePurchaseEventAnyAction,
	} {
		if _, ok := g.onMarketplacePurchaseEvent[action]; ok {
			for _, h := range g.onMarketplacePurchaseEvent[action] {
				handle := h
				eg.Go(func() error {
					err := handle(ctx, deliveryID, eventName, event)
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

// OnMarketplacePurchaseEventChanged registers callbacks listening to events of type github.MarketplacePurchaseEvent and action 'changed'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnMarketplacePurchaseEventChanged must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#marketplace_purchase
func (g *EventHandler) OnMarketplacePurchaseEventChanged(callbacks ...MarketplacePurchaseEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onMarketplacePurchaseEvent == nil {
		g.onMarketplacePurchaseEvent = make(map[string][]MarketplacePurchaseEventHandleFunc)
	}
	g.onMarketplacePurchaseEvent[MarketplacePurchaseEventChangedAction] = append(
		g.onMarketplacePurchaseEvent[MarketplacePurchaseEventChangedAction],
		callbacks...,
	)
}

// SetOnMarketplacePurchaseEventChanged registers callbacks listening to events of type github.MarketplacePurchaseEvent and action 'changed'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnMarketplacePurchaseEventChangedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#marketplace_purchase
func (g *EventHandler) SetOnMarketplacePurchaseEventChanged(callbacks ...MarketplacePurchaseEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onMarketplacePurchaseEvent == nil {
		g.onMarketplacePurchaseEvent = make(map[string][]MarketplacePurchaseEventHandleFunc)
	}
	g.onMarketplacePurchaseEvent[MarketplacePurchaseEventChangedAction] = callbacks
}

func (g *EventHandler) handleMarketplacePurchaseEventChanged(ctx context.Context, deliveryID string, eventName string, event *github.MarketplacePurchaseEvent) error {
	ctx, span := g.Tracer.Start(ctx, "handleMarketplacePurchaseEventChanged", trace.WithAttributes(
		attribute.String("deliveryID", deliveryID),
		attribute.String("event", eventName),
	))
	defer span.End()
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if MarketplacePurchaseEventChangedAction != *event.Action {
		err := fmt.Errorf(
			"handleMarketplacePurchaseEventChanged() called with wrong action, want %s, got %s",
			MarketplacePurchaseEventChangedAction,
			*event.Action,
		)
		span.SetStatus(codes.Error, err.Error())
		return err
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		MarketplacePurchaseEventChangedAction,
		MarketplacePurchaseEventAnyAction,
	} {
		if _, ok := g.onMarketplacePurchaseEvent[action]; ok {
			for _, h := range g.onMarketplacePurchaseEvent[action] {
				handle := h
				eg.Go(func() error {
					err := handle(ctx, deliveryID, eventName, event)
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

// OnMarketplacePurchaseEventCancelled registers callbacks listening to events of type github.MarketplacePurchaseEvent and action 'cancelled'.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnMarketplacePurchaseEventCancelled must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#marketplace_purchase
func (g *EventHandler) OnMarketplacePurchaseEventCancelled(callbacks ...MarketplacePurchaseEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onMarketplacePurchaseEvent == nil {
		g.onMarketplacePurchaseEvent = make(map[string][]MarketplacePurchaseEventHandleFunc)
	}
	g.onMarketplacePurchaseEvent[MarketplacePurchaseEventCancelledAction] = append(
		g.onMarketplacePurchaseEvent[MarketplacePurchaseEventCancelledAction],
		callbacks...,
	)
}

// SetOnMarketplacePurchaseEventCancelled registers callbacks listening to events of type github.MarketplacePurchaseEvent and action 'cancelled'
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnMarketplacePurchaseEventCancelledAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#marketplace_purchase
func (g *EventHandler) SetOnMarketplacePurchaseEventCancelled(callbacks ...MarketplacePurchaseEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onMarketplacePurchaseEvent == nil {
		g.onMarketplacePurchaseEvent = make(map[string][]MarketplacePurchaseEventHandleFunc)
	}
	g.onMarketplacePurchaseEvent[MarketplacePurchaseEventCancelledAction] = callbacks
}

func (g *EventHandler) handleMarketplacePurchaseEventCancelled(ctx context.Context, deliveryID string, eventName string, event *github.MarketplacePurchaseEvent) error {
	ctx, span := g.Tracer.Start(ctx, "handleMarketplacePurchaseEventCancelled", trace.WithAttributes(
		attribute.String("deliveryID", deliveryID),
		attribute.String("event", eventName),
	))
	defer span.End()
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	if MarketplacePurchaseEventCancelledAction != *event.Action {
		err := fmt.Errorf(
			"handleMarketplacePurchaseEventCancelled() called with wrong action, want %s, got %s",
			MarketplacePurchaseEventCancelledAction,
			*event.Action,
		)
		span.SetStatus(codes.Error, err.Error())
		return err
	}
	eg := new(errgroup.Group)
	for _, action := range []string{
		MarketplacePurchaseEventCancelledAction,
		MarketplacePurchaseEventAnyAction,
	} {
		if _, ok := g.onMarketplacePurchaseEvent[action]; ok {
			for _, h := range g.onMarketplacePurchaseEvent[action] {
				handle := h
				eg.Go(func() error {
					err := handle(ctx, deliveryID, eventName, event)
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

// OnMarketplacePurchaseEventAny registers callbacks listening to any events of type github.MarketplacePurchaseEvent
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnMarketplacePurchaseEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#marketplace_purchase
func (g *EventHandler) OnMarketplacePurchaseEventAny(callbacks ...MarketplacePurchaseEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onMarketplacePurchaseEvent == nil {
		g.onMarketplacePurchaseEvent = make(map[string][]MarketplacePurchaseEventHandleFunc)
	}
	g.onMarketplacePurchaseEvent[MarketplacePurchaseEventAnyAction] = append(
		g.onMarketplacePurchaseEvent[MarketplacePurchaseEventAnyAction],
		callbacks...,
	)
}

// SetOnMarketplacePurchaseEventAny registers callbacks listening to any events of type github.MarketplacePurchaseEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnMarketplacePurchaseEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#marketplace_purchase
func (g *EventHandler) SetOnMarketplacePurchaseEventAny(callbacks ...MarketplacePurchaseEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onMarketplacePurchaseEvent == nil {
		g.onMarketplacePurchaseEvent = make(map[string][]MarketplacePurchaseEventHandleFunc)
	}
	g.onMarketplacePurchaseEvent[MarketplacePurchaseEventAnyAction] = callbacks
}

func (g *EventHandler) handleMarketplacePurchaseEventAny(ctx context.Context, deliveryID string, eventName string, event *github.MarketplacePurchaseEvent) error {
	ctx, span := g.Tracer.Start(ctx, "handleMarketplacePurchaseEventAny", trace.WithAttributes(
		attribute.String("deliveryID", deliveryID),
		attribute.String("event", eventName),
	))
	defer span.End()
	if event == nil {
		err := fmt.Errorf("event was empty or nil")
		span.SetStatus(codes.Error, err.Error())
		return err
	}
	if _, ok := g.onMarketplacePurchaseEvent[MarketplacePurchaseEventAnyAction]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onMarketplacePurchaseEvent[MarketplacePurchaseEventAnyAction] {
		handle := h
		eg.Go(func() error {
			err := handle(ctx, deliveryID, eventName, event)
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

// MarketplacePurchaseEvent handles github.MarketplacePurchaseEvent.
//
// Callbacks are executed in the following order:
//
// 1) All callbacks registered with OnBeforeAny are executed in parallel.
// 2) All callbacks registered with OnMarketplacePurchaseEvent... are executed in parallel in case the Event has actions.
// 3) All callbacks registered with OnAfterAny are executed in parallel.
//
// on any error all callbacks registered with OnError are executed in parallel.
func (g *EventHandler) MarketplacePurchaseEvent(ctx context.Context, deliveryID string, eventName string, event *github.MarketplacePurchaseEvent) error {
	ctx, span := g.Tracer.Start(ctx, "MarketplacePurchaseEvent", trace.WithAttributes(
		attribute.String("deliveryID", deliveryID),
		attribute.String("event", eventName),
	))
	defer span.End()

	if event == nil || event.Action == nil || *event.Action == "" {
		err := fmt.Errorf("event action was empty or nil")
		span.SetStatus(codes.Error, err.Error())
		return err
	}
	action := *event.Action

	err := g.handleBeforeAny(ctx, deliveryID, eventName, event)
	if err != nil {
		return g.handleError(ctx, deliveryID, eventName, event, err)
	}

	switch action {

	case MarketplacePurchaseEventPurchasedAction:
		err := g.handleMarketplacePurchaseEventPurchased(ctx, deliveryID, eventName, event)
		if err != nil {
			return g.handleError(ctx, deliveryID, eventName, event, err)
		}

	case MarketplacePurchaseEventPendingChangeAction:
		err := g.handleMarketplacePurchaseEventPendingChange(ctx, deliveryID, eventName, event)
		if err != nil {
			return g.handleError(ctx, deliveryID, eventName, event, err)
		}

	case MarketplacePurchaseEventPendingChangeCancelledAction:
		err := g.handleMarketplacePurchaseEventPendingChangeCancelled(ctx, deliveryID, eventName, event)
		if err != nil {
			return g.handleError(ctx, deliveryID, eventName, event, err)
		}

	case MarketplacePurchaseEventChangedAction:
		err := g.handleMarketplacePurchaseEventChanged(ctx, deliveryID, eventName, event)
		if err != nil {
			return g.handleError(ctx, deliveryID, eventName, event, err)
		}

	case MarketplacePurchaseEventCancelledAction:
		err := g.handleMarketplacePurchaseEventCancelled(ctx, deliveryID, eventName, event)
		if err != nil {
			return g.handleError(ctx, deliveryID, eventName, event, err)
		}

	default:
		err := g.handleMarketplacePurchaseEventAny(ctx, deliveryID, eventName, event)
		if err != nil {
			return g.handleError(ctx, deliveryID, eventName, event, err)
		}
	}

	err = g.handleAfterAny(ctx, deliveryID, eventName, event)
	if err != nil {
		return g.handleError(ctx, deliveryID, eventName, event, err)
	}
	return nil
}
