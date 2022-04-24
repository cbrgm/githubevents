package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"fmt"
	"github.com/google/go-github/v43/github"
	"golang.org/x/sync/errgroup"
)

// MembershipEventHandleFunc represents a callback function triggered on github.MembershipEvent.
// deliveryID (type: string) is the unique webhook delivery ID.
// eventName (type: string) is the name of the event.
// event (type: *github.MembershipEvent) is the webhook payload.
type MembershipEventHandleFunc func(deliveryID string, eventName string, event *github.MembershipEvent) error

// OnMembershipEventAdded registers callbacks listening to events of type github.MembershipEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnMembershipEventAdded must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnMembershipEventAdded(callbacks ...MembershipEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "added"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onMembershipEvent == nil {
		g.onMembershipEvent = make(map[string][]MembershipEventHandleFunc)
	}
	g.onMembershipEvent[action] = append(g.onMembershipEvent[action], callbacks...)
}

// SetOnMembershipEventAdded registers callbacks listening to events of type github.MembershipEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnMembershipEventAddedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnMembershipEventAdded(callbacks ...MembershipEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "added"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onMembershipEvent == nil {
		g.onMembershipEvent = make(map[string][]MembershipEventHandleFunc)
	}
	g.onMembershipEvent[action] = callbacks
}

func (g *EventHandler) handleMembershipEventAdded(deliveryID string, eventName string, event *github.MembershipEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "added"
	if action != *event.Action {
		return fmt.Errorf(
			"handleMembershipEventAdded() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleMembershipEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onMembershipEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onMembershipEvent[action] {
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

// OnMembershipEventRemoved registers callbacks listening to events of type github.MembershipEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnMembershipEventRemoved must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnMembershipEventRemoved(callbacks ...MembershipEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "removed"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onMembershipEvent == nil {
		g.onMembershipEvent = make(map[string][]MembershipEventHandleFunc)
	}
	g.onMembershipEvent[action] = append(g.onMembershipEvent[action], callbacks...)
}

// SetOnMembershipEventRemoved registers callbacks listening to events of type github.MembershipEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnMembershipEventRemovedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnMembershipEventRemoved(callbacks ...MembershipEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "removed"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onMembershipEvent == nil {
		g.onMembershipEvent = make(map[string][]MembershipEventHandleFunc)
	}
	g.onMembershipEvent[action] = callbacks
}

func (g *EventHandler) handleMembershipEventRemoved(deliveryID string, eventName string, event *github.MembershipEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "removed"
	if action != *event.Action {
		return fmt.Errorf(
			"handleMembershipEventRemoved() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleMembershipEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onMembershipEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onMembershipEvent[action] {
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

// OnMembershipEventAny registers callbacks listening to events of type github.MembershipEvent
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnMembershipEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnMembershipEventAny(callbacks ...MembershipEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const any = "*"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onMembershipEvent == nil {
		g.onMembershipEvent = make(map[string][]MembershipEventHandleFunc)
	}
	g.onMembershipEvent[any] = append(g.onMembershipEvent[any], callbacks...)
}

// SetOnMembershipEventAny registers callbacks listening to events of type github.MembershipEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnMembershipEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnMembershipEventAny(callbacks ...MembershipEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const any = "*"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onMembershipEvent == nil {
		g.onMembershipEvent = make(map[string][]MembershipEventHandleFunc)
	}
	g.onMembershipEvent[any] = callbacks
}

func (g *EventHandler) handleMembershipEventAny(deliveryID string, eventName string, event *github.MembershipEvent) error {
	if event == nil {
		return fmt.Errorf("event was empty or nil")
	}
	const any = "*"
	if _, ok := g.onMembershipEvent[any]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onMembershipEvent[any] {
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

// MembershipEvent handles github.MembershipEvent
//
// Callbacks are executed in the following order:
//
// 1) All callbacks registered by OnBeforeAny are executed in parallel.
// 2) All callbacks registered by OnMembershipEventAny are executed in parallel.
// 3) Optional: All callbacks registered via OnMembershipEvent... are executed in parallel in case the Event has actions.
// 4) All callbacks registered by OnAfterAny are executed in parallel.
// on any error all callbacks registered by OnError are executed in parallel.
//
func (g *EventHandler) MembershipEvent(deliveryID string, eventName string, event *github.MembershipEvent) error {

	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	action := *event.Action

	err := g.handleBeforeAny(deliveryID, eventName, event)
	if err != nil {
		return g.handleError(deliveryID, eventName, event, err)
	}

	switch action {

	case "added":
		err := g.handleMembershipEventAdded(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "removed":
		err := g.handleMembershipEventRemoved(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	default:
		err := g.handleMembershipEventAny(deliveryID, eventName, event)
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
