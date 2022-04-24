package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"fmt"
	"github.com/google/go-github/v43/github"
	"golang.org/x/sync/errgroup"
)

// BranchProtectionRuleEventHandleFunc represents a callback function triggered on github.BranchProtectionRuleEvent.
// deliveryID (type: string) is the unique webhook delivery ID.
// eventName (type: string) is the name of the event.
// event (type: *github.BranchProtectionRuleEvent) is the webhook payload.
type BranchProtectionRuleEventHandleFunc func(deliveryID string, eventName string, event *github.BranchProtectionRuleEvent) error

// OnBranchProtectionRuleEventCreated registers callbacks listening to events of type github.BranchProtectionRuleEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnBranchProtectionRuleEventCreated must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnBranchProtectionRuleEventCreated(callbacks ...BranchProtectionRuleEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "created"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onBranchProtectionRuleEvent == nil {
		g.onBranchProtectionRuleEvent = make(map[string][]BranchProtectionRuleEventHandleFunc)
	}
	g.onBranchProtectionRuleEvent[action] = append(g.onBranchProtectionRuleEvent[action], callbacks...)
}

// SetOnBranchProtectionRuleEventCreated registers callbacks listening to events of type github.BranchProtectionRuleEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnBranchProtectionRuleEventCreatedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnBranchProtectionRuleEventCreated(callbacks ...BranchProtectionRuleEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "created"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onBranchProtectionRuleEvent == nil {
		g.onBranchProtectionRuleEvent = make(map[string][]BranchProtectionRuleEventHandleFunc)
	}
	g.onBranchProtectionRuleEvent[action] = callbacks
}

func (g *EventHandler) handleBranchProtectionRuleEventCreated(deliveryID string, eventName string, event *github.BranchProtectionRuleEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "created"
	if action != *event.Action {
		return fmt.Errorf(
			"handleBranchProtectionRuleEventCreated() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleBranchProtectionRuleEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onBranchProtectionRuleEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onBranchProtectionRuleEvent[action] {
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

// OnBranchProtectionRuleEventEdited registers callbacks listening to events of type github.BranchProtectionRuleEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnBranchProtectionRuleEventEdited must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnBranchProtectionRuleEventEdited(callbacks ...BranchProtectionRuleEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "edited"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onBranchProtectionRuleEvent == nil {
		g.onBranchProtectionRuleEvent = make(map[string][]BranchProtectionRuleEventHandleFunc)
	}
	g.onBranchProtectionRuleEvent[action] = append(g.onBranchProtectionRuleEvent[action], callbacks...)
}

// SetOnBranchProtectionRuleEventEdited registers callbacks listening to events of type github.BranchProtectionRuleEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnBranchProtectionRuleEventEditedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnBranchProtectionRuleEventEdited(callbacks ...BranchProtectionRuleEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "edited"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onBranchProtectionRuleEvent == nil {
		g.onBranchProtectionRuleEvent = make(map[string][]BranchProtectionRuleEventHandleFunc)
	}
	g.onBranchProtectionRuleEvent[action] = callbacks
}

func (g *EventHandler) handleBranchProtectionRuleEventEdited(deliveryID string, eventName string, event *github.BranchProtectionRuleEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "edited"
	if action != *event.Action {
		return fmt.Errorf(
			"handleBranchProtectionRuleEventEdited() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleBranchProtectionRuleEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onBranchProtectionRuleEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onBranchProtectionRuleEvent[action] {
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

// OnBranchProtectionRuleEventDeleted registers callbacks listening to events of type github.BranchProtectionRuleEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnBranchProtectionRuleEventDeleted must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnBranchProtectionRuleEventDeleted(callbacks ...BranchProtectionRuleEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "deleted"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onBranchProtectionRuleEvent == nil {
		g.onBranchProtectionRuleEvent = make(map[string][]BranchProtectionRuleEventHandleFunc)
	}
	g.onBranchProtectionRuleEvent[action] = append(g.onBranchProtectionRuleEvent[action], callbacks...)
}

// SetOnBranchProtectionRuleEventDeleted registers callbacks listening to events of type github.BranchProtectionRuleEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnBranchProtectionRuleEventDeletedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnBranchProtectionRuleEventDeleted(callbacks ...BranchProtectionRuleEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "deleted"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onBranchProtectionRuleEvent == nil {
		g.onBranchProtectionRuleEvent = make(map[string][]BranchProtectionRuleEventHandleFunc)
	}
	g.onBranchProtectionRuleEvent[action] = callbacks
}

func (g *EventHandler) handleBranchProtectionRuleEventDeleted(deliveryID string, eventName string, event *github.BranchProtectionRuleEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "deleted"
	if action != *event.Action {
		return fmt.Errorf(
			"handleBranchProtectionRuleEventDeleted() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleBranchProtectionRuleEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onBranchProtectionRuleEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onBranchProtectionRuleEvent[action] {
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

// OnBranchProtectionRuleEventAny registers callbacks listening to events of type github.BranchProtectionRuleEvent
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnBranchProtectionRuleEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnBranchProtectionRuleEventAny(callbacks ...BranchProtectionRuleEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const any = "*"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onBranchProtectionRuleEvent == nil {
		g.onBranchProtectionRuleEvent = make(map[string][]BranchProtectionRuleEventHandleFunc)
	}
	g.onBranchProtectionRuleEvent[any] = append(g.onBranchProtectionRuleEvent[any], callbacks...)
}

// SetOnBranchProtectionRuleEventAny registers callbacks listening to events of type github.BranchProtectionRuleEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnBranchProtectionRuleEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnBranchProtectionRuleEventAny(callbacks ...BranchProtectionRuleEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const any = "*"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onBranchProtectionRuleEvent == nil {
		g.onBranchProtectionRuleEvent = make(map[string][]BranchProtectionRuleEventHandleFunc)
	}
	g.onBranchProtectionRuleEvent[any] = callbacks
}

func (g *EventHandler) handleBranchProtectionRuleEventAny(deliveryID string, eventName string, event *github.BranchProtectionRuleEvent) error {
	if event == nil {
		return fmt.Errorf("event was empty or nil")
	}
	const any = "*"
	if _, ok := g.onBranchProtectionRuleEvent[any]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onBranchProtectionRuleEvent[any] {
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

// BranchProtectionRuleEvent handles github.BranchProtectionRuleEvent
//
// Callbacks are executed in the following order:
//
// 1) All callbacks registered with OnBeforeAny are executed in parallel.
// 2) All callbacks registered with OnBranchProtectionRuleEventAny are executed in parallel.
// 3) Optional: All callbacks registered with OnBranchProtectionRuleEvent... are executed in parallel in case the Event has actions.
// 4) All callbacks registered with OnAfterAny are executed in parallel.
//
// on any error all callbacks registered with OnError are executed in parallel.
func (g *EventHandler) BranchProtectionRuleEvent(deliveryID string, eventName string, event *github.BranchProtectionRuleEvent) error {

	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	action := *event.Action

	err := g.handleBeforeAny(deliveryID, eventName, event)
	if err != nil {
		return g.handleError(deliveryID, eventName, event, err)
	}

	switch action {

	case "created":
		err := g.handleBranchProtectionRuleEventCreated(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "edited":
		err := g.handleBranchProtectionRuleEventEdited(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "deleted":
		err := g.handleBranchProtectionRuleEventDeleted(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	default:
		err := g.handleBranchProtectionRuleEventAny(deliveryID, eventName, event)
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
