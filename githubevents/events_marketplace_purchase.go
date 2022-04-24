package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"fmt"
	"github.com/google/go-github/v43/github"
	"golang.org/x/sync/errgroup"
)

// MarketplacePurchaseEventHandleFunc represents a callback function triggered on github.MarketplacePurchaseEvent.
// deliveryID (type: string) is the unique webhook delivery ID.
// eventName (type: string) is the name of the event.
// event (type: *github.MarketplacePurchaseEvent) is the webhook payload.
type MarketplacePurchaseEventHandleFunc func(deliveryID string, eventName string, event *github.MarketplacePurchaseEvent) error

// OnMarketplacePurchaseEventPurchased registers callbacks listening to events of type github.MarketplacePurchaseEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnMarketplacePurchaseEventPurchased must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnMarketplacePurchaseEventPurchased(callbacks ...MarketplacePurchaseEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "purchased"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onMarketplacePurchaseEvent == nil {
		g.onMarketplacePurchaseEvent = make(map[string][]MarketplacePurchaseEventHandleFunc)
	}
	g.onMarketplacePurchaseEvent[action] = append(g.onMarketplacePurchaseEvent[action], callbacks...)
}

// SetOnMarketplacePurchaseEventPurchased registers callbacks listening to events of type github.MarketplacePurchaseEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnMarketplacePurchaseEventPurchasedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnMarketplacePurchaseEventPurchased(callbacks ...MarketplacePurchaseEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "purchased"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onMarketplacePurchaseEvent == nil {
		g.onMarketplacePurchaseEvent = make(map[string][]MarketplacePurchaseEventHandleFunc)
	}
	g.onMarketplacePurchaseEvent[action] = callbacks
}

func (g *EventHandler) handleMarketplacePurchaseEventPurchased(deliveryID string, eventName string, event *github.MarketplacePurchaseEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "purchased"
	if action != *event.Action {
		return fmt.Errorf(
			"handleMarketplacePurchaseEventPurchased() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleMarketplacePurchaseEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onMarketplacePurchaseEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onMarketplacePurchaseEvent[action] {
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

// OnMarketplacePurchaseEventPendingChange registers callbacks listening to events of type github.MarketplacePurchaseEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnMarketplacePurchaseEventPendingChange must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnMarketplacePurchaseEventPendingChange(callbacks ...MarketplacePurchaseEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "pending_change"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onMarketplacePurchaseEvent == nil {
		g.onMarketplacePurchaseEvent = make(map[string][]MarketplacePurchaseEventHandleFunc)
	}
	g.onMarketplacePurchaseEvent[action] = append(g.onMarketplacePurchaseEvent[action], callbacks...)
}

// SetOnMarketplacePurchaseEventPendingChange registers callbacks listening to events of type github.MarketplacePurchaseEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnMarketplacePurchaseEventPendingChangeAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnMarketplacePurchaseEventPendingChange(callbacks ...MarketplacePurchaseEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "pending_change"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onMarketplacePurchaseEvent == nil {
		g.onMarketplacePurchaseEvent = make(map[string][]MarketplacePurchaseEventHandleFunc)
	}
	g.onMarketplacePurchaseEvent[action] = callbacks
}

func (g *EventHandler) handleMarketplacePurchaseEventPendingChange(deliveryID string, eventName string, event *github.MarketplacePurchaseEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "pending_change"
	if action != *event.Action {
		return fmt.Errorf(
			"handleMarketplacePurchaseEventPendingChange() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleMarketplacePurchaseEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onMarketplacePurchaseEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onMarketplacePurchaseEvent[action] {
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

// OnMarketplacePurchaseEventPendingChangeCancelled registers callbacks listening to events of type github.MarketplacePurchaseEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnMarketplacePurchaseEventPendingChangeCancelled must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnMarketplacePurchaseEventPendingChangeCancelled(callbacks ...MarketplacePurchaseEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "pending_change_cancelled"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onMarketplacePurchaseEvent == nil {
		g.onMarketplacePurchaseEvent = make(map[string][]MarketplacePurchaseEventHandleFunc)
	}
	g.onMarketplacePurchaseEvent[action] = append(g.onMarketplacePurchaseEvent[action], callbacks...)
}

// SetOnMarketplacePurchaseEventPendingChangeCancelled registers callbacks listening to events of type github.MarketplacePurchaseEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnMarketplacePurchaseEventPendingChangeCancelledAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnMarketplacePurchaseEventPendingChangeCancelled(callbacks ...MarketplacePurchaseEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "pending_change_cancelled"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onMarketplacePurchaseEvent == nil {
		g.onMarketplacePurchaseEvent = make(map[string][]MarketplacePurchaseEventHandleFunc)
	}
	g.onMarketplacePurchaseEvent[action] = callbacks
}

func (g *EventHandler) handleMarketplacePurchaseEventPendingChangeCancelled(deliveryID string, eventName string, event *github.MarketplacePurchaseEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "pending_change_cancelled"
	if action != *event.Action {
		return fmt.Errorf(
			"handleMarketplacePurchaseEventPendingChangeCancelled() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleMarketplacePurchaseEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onMarketplacePurchaseEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onMarketplacePurchaseEvent[action] {
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

// OnMarketplacePurchaseEventChanged registers callbacks listening to events of type github.MarketplacePurchaseEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnMarketplacePurchaseEventChanged must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnMarketplacePurchaseEventChanged(callbacks ...MarketplacePurchaseEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "changed"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onMarketplacePurchaseEvent == nil {
		g.onMarketplacePurchaseEvent = make(map[string][]MarketplacePurchaseEventHandleFunc)
	}
	g.onMarketplacePurchaseEvent[action] = append(g.onMarketplacePurchaseEvent[action], callbacks...)
}

// SetOnMarketplacePurchaseEventChanged registers callbacks listening to events of type github.MarketplacePurchaseEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnMarketplacePurchaseEventChangedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnMarketplacePurchaseEventChanged(callbacks ...MarketplacePurchaseEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "changed"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onMarketplacePurchaseEvent == nil {
		g.onMarketplacePurchaseEvent = make(map[string][]MarketplacePurchaseEventHandleFunc)
	}
	g.onMarketplacePurchaseEvent[action] = callbacks
}

func (g *EventHandler) handleMarketplacePurchaseEventChanged(deliveryID string, eventName string, event *github.MarketplacePurchaseEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "changed"
	if action != *event.Action {
		return fmt.Errorf(
			"handleMarketplacePurchaseEventChanged() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleMarketplacePurchaseEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onMarketplacePurchaseEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onMarketplacePurchaseEvent[action] {
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

// OnMarketplacePurchaseEventCancelled registers callbacks listening to events of type github.MarketplacePurchaseEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnMarketplacePurchaseEventCancelled must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnMarketplacePurchaseEventCancelled(callbacks ...MarketplacePurchaseEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "cancelled"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onMarketplacePurchaseEvent == nil {
		g.onMarketplacePurchaseEvent = make(map[string][]MarketplacePurchaseEventHandleFunc)
	}
	g.onMarketplacePurchaseEvent[action] = append(g.onMarketplacePurchaseEvent[action], callbacks...)
}

// SetOnMarketplacePurchaseEventCancelled registers callbacks listening to events of type github.MarketplacePurchaseEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnMarketplacePurchaseEventCancelledAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnMarketplacePurchaseEventCancelled(callbacks ...MarketplacePurchaseEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "cancelled"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onMarketplacePurchaseEvent == nil {
		g.onMarketplacePurchaseEvent = make(map[string][]MarketplacePurchaseEventHandleFunc)
	}
	g.onMarketplacePurchaseEvent[action] = callbacks
}

func (g *EventHandler) handleMarketplacePurchaseEventCancelled(deliveryID string, eventName string, event *github.MarketplacePurchaseEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "cancelled"
	if action != *event.Action {
		return fmt.Errorf(
			"handleMarketplacePurchaseEventCancelled() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleMarketplacePurchaseEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onMarketplacePurchaseEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onMarketplacePurchaseEvent[action] {
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

// OnMarketplacePurchaseEventAny registers callbacks listening to events of type github.MarketplacePurchaseEvent
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnMarketplacePurchaseEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnMarketplacePurchaseEventAny(callbacks ...MarketplacePurchaseEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const any = "*"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onMarketplacePurchaseEvent == nil {
		g.onMarketplacePurchaseEvent = make(map[string][]MarketplacePurchaseEventHandleFunc)
	}
	g.onMarketplacePurchaseEvent[any] = append(g.onMarketplacePurchaseEvent[any], callbacks...)
}

// SetOnMarketplacePurchaseEventAny registers callbacks listening to events of type github.MarketplacePurchaseEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnMarketplacePurchaseEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnMarketplacePurchaseEventAny(callbacks ...MarketplacePurchaseEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const any = "*"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onMarketplacePurchaseEvent == nil {
		g.onMarketplacePurchaseEvent = make(map[string][]MarketplacePurchaseEventHandleFunc)
	}
	g.onMarketplacePurchaseEvent[any] = callbacks
}

func (g *EventHandler) handleMarketplacePurchaseEventAny(deliveryID string, eventName string, event *github.MarketplacePurchaseEvent) error {
	if event == nil {
		return fmt.Errorf("event was empty or nil")
	}
	const any = "*"
	if _, ok := g.onMarketplacePurchaseEvent[any]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onMarketplacePurchaseEvent[any] {
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

// MarketplacePurchaseEvent handles github.MarketplacePurchaseEvent
//
// Callbacks are executed in the following order:
//
// 1) All callbacks registered by OnBeforeAny are executed in parallel.
// 2) All callbacks registered by OnMarketplacePurchaseEventAny are executed in parallel.
// 3) Optional: All callbacks registered via OnMarketplacePurchaseEvent... are executed in parallel in case the Event has actions.
// 4) All callbacks registered by OnAfterAny are executed in parallel.
// on any error all callbacks registered by OnError are executed in parallel.
//
func (g *EventHandler) MarketplacePurchaseEvent(deliveryID string, eventName string, event *github.MarketplacePurchaseEvent) error {

	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	action := *event.Action

	err := g.handleBeforeAny(deliveryID, eventName, event)
	if err != nil {
		return g.handleError(deliveryID, eventName, event, err)
	}

	switch action {

	case "purchased":
		err := g.handleMarketplacePurchaseEventPurchased(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "pending_change":
		err := g.handleMarketplacePurchaseEventPendingChange(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "pending_change_cancelled":
		err := g.handleMarketplacePurchaseEventPendingChangeCancelled(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "changed":
		err := g.handleMarketplacePurchaseEventChanged(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "cancelled":
		err := g.handleMarketplacePurchaseEventCancelled(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	default:
		err := g.handleMarketplacePurchaseEventAny(deliveryID, eventName, event)
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
