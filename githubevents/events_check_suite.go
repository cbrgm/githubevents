package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"fmt"
	"github.com/google/go-github/v43/github"
	"golang.org/x/sync/errgroup"
)

// CheckSuiteEventHandleFunc represents a callback function triggered on github.CheckSuiteEvent.
// deliveryID (type: string) is the unique webhook delivery ID.
// eventName (type: string) is the name of the event.
// event (type: *github.CheckSuiteEvent) is the webhook payload.
type CheckSuiteEventHandleFunc func(deliveryID string, eventName string, event *github.CheckSuiteEvent) error

// OnCheckSuiteEventCompleted registers callbacks listening to events of type github.CheckSuiteEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnCheckSuiteEventCompleted must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnCheckSuiteEventCompleted(callbacks ...CheckSuiteEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "completed"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onCheckSuiteEvent == nil {
		g.onCheckSuiteEvent = make(map[string][]CheckSuiteEventHandleFunc)
	}
	g.onCheckSuiteEvent[action] = append(g.onCheckSuiteEvent[action], callbacks...)
}

// SetOnCheckSuiteEventCompleted registers callbacks listening to events of type github.CheckSuiteEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnCheckSuiteEventCompletedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnCheckSuiteEventCompleted(callbacks ...CheckSuiteEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "completed"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onCheckSuiteEvent == nil {
		g.onCheckSuiteEvent = make(map[string][]CheckSuiteEventHandleFunc)
	}
	g.onCheckSuiteEvent[action] = callbacks
}

func (g *EventHandler) handleCheckSuiteEventCompleted(deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "completed"
	if action != *event.Action {
		return fmt.Errorf(
			"handleCheckSuiteEventCompleted() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleCheckSuiteEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onCheckSuiteEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onCheckSuiteEvent[action] {
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

// OnCheckSuiteEventRequested registers callbacks listening to events of type github.CheckSuiteEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnCheckSuiteEventRequested must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnCheckSuiteEventRequested(callbacks ...CheckSuiteEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "requested"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onCheckSuiteEvent == nil {
		g.onCheckSuiteEvent = make(map[string][]CheckSuiteEventHandleFunc)
	}
	g.onCheckSuiteEvent[action] = append(g.onCheckSuiteEvent[action], callbacks...)
}

// SetOnCheckSuiteEventRequested registers callbacks listening to events of type github.CheckSuiteEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnCheckSuiteEventRequestedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnCheckSuiteEventRequested(callbacks ...CheckSuiteEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "requested"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onCheckSuiteEvent == nil {
		g.onCheckSuiteEvent = make(map[string][]CheckSuiteEventHandleFunc)
	}
	g.onCheckSuiteEvent[action] = callbacks
}

func (g *EventHandler) handleCheckSuiteEventRequested(deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "requested"
	if action != *event.Action {
		return fmt.Errorf(
			"handleCheckSuiteEventRequested() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleCheckSuiteEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onCheckSuiteEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onCheckSuiteEvent[action] {
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

// OnCheckSuiteEventReRequested registers callbacks listening to events of type github.CheckSuiteEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnCheckSuiteEventReRequested must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnCheckSuiteEventReRequested(callbacks ...CheckSuiteEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "rerequested"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onCheckSuiteEvent == nil {
		g.onCheckSuiteEvent = make(map[string][]CheckSuiteEventHandleFunc)
	}
	g.onCheckSuiteEvent[action] = append(g.onCheckSuiteEvent[action], callbacks...)
}

// SetOnCheckSuiteEventReRequested registers callbacks listening to events of type github.CheckSuiteEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnCheckSuiteEventReRequestedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnCheckSuiteEventReRequested(callbacks ...CheckSuiteEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "rerequested"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onCheckSuiteEvent == nil {
		g.onCheckSuiteEvent = make(map[string][]CheckSuiteEventHandleFunc)
	}
	g.onCheckSuiteEvent[action] = callbacks
}

func (g *EventHandler) handleCheckSuiteEventReRequested(deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "rerequested"
	if action != *event.Action {
		return fmt.Errorf(
			"handleCheckSuiteEventReRequested() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleCheckSuiteEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onCheckSuiteEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onCheckSuiteEvent[action] {
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

// OnCheckSuiteEventAny registers callbacks listening to events of type github.CheckSuiteEvent
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnCheckSuiteEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnCheckSuiteEventAny(callbacks ...CheckSuiteEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const any = "*"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onCheckSuiteEvent == nil {
		g.onCheckSuiteEvent = make(map[string][]CheckSuiteEventHandleFunc)
	}
	g.onCheckSuiteEvent[any] = append(g.onCheckSuiteEvent[any], callbacks...)
}

// SetOnCheckSuiteEventAny registers callbacks listening to events of type github.CheckSuiteEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnCheckSuiteEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnCheckSuiteEventAny(callbacks ...CheckSuiteEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const any = "*"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onCheckSuiteEvent == nil {
		g.onCheckSuiteEvent = make(map[string][]CheckSuiteEventHandleFunc)
	}
	g.onCheckSuiteEvent[any] = callbacks
}

func (g *EventHandler) handleCheckSuiteEventAny(deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
	if event == nil {
		return fmt.Errorf("event was empty or nil")
	}
	const any = "*"
	if _, ok := g.onCheckSuiteEvent[any]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onCheckSuiteEvent[any] {
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

// CheckSuiteEvent handles github.CheckSuiteEvent
//
// Callbacks are executed in the following order:
//
// 1) All callbacks registered by OnBeforeAny are executed in parallel.
// 2) All callbacks registered by OnCheckSuiteEventAny are executed in parallel.
// 3) Optional: All callbacks registered via OnCheckSuiteEvent... are executed in parallel in case the Event has actions.
// 4) All callbacks registered by OnAfterAny are executed in parallel.
// on any error all callbacks registered by OnError are executed in parallel.
//
func (g *EventHandler) CheckSuiteEvent(deliveryID string, eventName string, event *github.CheckSuiteEvent) error {

	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	action := *event.Action

	err := g.handleBeforeAny(deliveryID, eventName, event)
	if err != nil {
		return g.handleError(deliveryID, eventName, event, err)
	}

	switch action {

	case "completed":
		err := g.handleCheckSuiteEventCompleted(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "requested":
		err := g.handleCheckSuiteEventRequested(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "rerequested":
		err := g.handleCheckSuiteEventReRequested(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	default:
		err := g.handleCheckSuiteEventAny(deliveryID, eventName, event)
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
