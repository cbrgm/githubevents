package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"fmt"
	"github.com/google/go-github/v43/github"
	"golang.org/x/sync/errgroup"
)

// OrgBlockEventHandleFunc represents a callback function triggered on github.OrgBlockEvent.
// deliveryID (type: string) is the unique webhook delivery ID.
// eventName (type: string) is the name of the event.
// event (type: *github.OrgBlockEvent) is the webhook payload.
type OrgBlockEventHandleFunc func(deliveryID string, eventName string, event *github.OrgBlockEvent) error

// OnOrgBlockEventBlocked registers callbacks listening to events of type github.OrgBlockEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnOrgBlockEventBlocked must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnOrgBlockEventBlocked(callbacks ...OrgBlockEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "blocked"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onOrgBlockEvent == nil {
		g.onOrgBlockEvent = make(map[string][]OrgBlockEventHandleFunc)
	}
	g.onOrgBlockEvent[action] = append(g.onOrgBlockEvent[action], callbacks...)
}

// SetOnOrgBlockEventBlocked registers callbacks listening to events of type github.OrgBlockEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnOrgBlockEventBlockedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnOrgBlockEventBlocked(callbacks ...OrgBlockEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "blocked"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onOrgBlockEvent == nil {
		g.onOrgBlockEvent = make(map[string][]OrgBlockEventHandleFunc)
	}
	g.onOrgBlockEvent[action] = callbacks
}

func (g *EventHandler) handleOrgBlockEventBlocked(deliveryID string, eventName string, event *github.OrgBlockEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "blocked"
	if action != *event.Action {
		return fmt.Errorf(
			"handleOrgBlockEventBlocked() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleOrgBlockEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onOrgBlockEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onOrgBlockEvent[action] {
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

// OnOrgBlockEventUnblocked registers callbacks listening to events of type github.OrgBlockEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnOrgBlockEventUnblocked must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnOrgBlockEventUnblocked(callbacks ...OrgBlockEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "unblocked"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onOrgBlockEvent == nil {
		g.onOrgBlockEvent = make(map[string][]OrgBlockEventHandleFunc)
	}
	g.onOrgBlockEvent[action] = append(g.onOrgBlockEvent[action], callbacks...)
}

// SetOnOrgBlockEventUnblocked registers callbacks listening to events of type github.OrgBlockEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnOrgBlockEventUnblockedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnOrgBlockEventUnblocked(callbacks ...OrgBlockEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "unblocked"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onOrgBlockEvent == nil {
		g.onOrgBlockEvent = make(map[string][]OrgBlockEventHandleFunc)
	}
	g.onOrgBlockEvent[action] = callbacks
}

func (g *EventHandler) handleOrgBlockEventUnblocked(deliveryID string, eventName string, event *github.OrgBlockEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "unblocked"
	if action != *event.Action {
		return fmt.Errorf(
			"handleOrgBlockEventUnblocked() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleOrgBlockEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onOrgBlockEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onOrgBlockEvent[action] {
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

// OnOrgBlockEventAny registers callbacks listening to events of type github.OrgBlockEvent
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnOrgBlockEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnOrgBlockEventAny(callbacks ...OrgBlockEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const any = "*"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onOrgBlockEvent == nil {
		g.onOrgBlockEvent = make(map[string][]OrgBlockEventHandleFunc)
	}
	g.onOrgBlockEvent[any] = append(g.onOrgBlockEvent[any], callbacks...)
}

// SetOnOrgBlockEventAny registers callbacks listening to events of type github.OrgBlockEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnOrgBlockEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnOrgBlockEventAny(callbacks ...OrgBlockEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const any = "*"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onOrgBlockEvent == nil {
		g.onOrgBlockEvent = make(map[string][]OrgBlockEventHandleFunc)
	}
	g.onOrgBlockEvent[any] = callbacks
}

func (g *EventHandler) handleOrgBlockEventAny(deliveryID string, eventName string, event *github.OrgBlockEvent) error {
	if event == nil {
		return fmt.Errorf("event was empty or nil")
	}
	const any = "*"
	if _, ok := g.onOrgBlockEvent[any]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onOrgBlockEvent[any] {
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

// OrgBlockEvent handles github.OrgBlockEvent
//
// Callbacks are executed in the following order:
//
// 1) All callbacks registered with OnBeforeAny are executed in parallel.
// 2) All callbacks registered with OnOrgBlockEventAny are executed in parallel.
// 3) Optional: All callbacks registered with OnOrgBlockEvent... are executed in parallel in case the Event has actions.
// 4) All callbacks registered with OnAfterAny are executed in parallel.
//
// on any error all callbacks registered with OnError are executed in parallel.
func (g *EventHandler) OrgBlockEvent(deliveryID string, eventName string, event *github.OrgBlockEvent) error {

	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	action := *event.Action

	err := g.handleBeforeAny(deliveryID, eventName, event)
	if err != nil {
		return g.handleError(deliveryID, eventName, event, err)
	}

	switch action {

	case "blocked":
		err := g.handleOrgBlockEventBlocked(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "unblocked":
		err := g.handleOrgBlockEventUnblocked(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	default:
		err := g.handleOrgBlockEventAny(deliveryID, eventName, event)
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
