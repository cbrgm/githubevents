package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"fmt"
	"github.com/google/go-github/v43/github"
	"golang.org/x/sync/errgroup"
)

// CheckRunEventHandleFunc represents a callback function triggered on github.CheckRunEvent.
// deliveryID (type: string) is the unique webhook delivery ID.
// eventName (type: string) is the name of the event.
// event (type: *github.CheckRunEvent) is the webhook payload.
type CheckRunEventHandleFunc func(deliveryID string, eventName string, event *github.CheckRunEvent) error

// OnCheckRunEventCreated registers callbacks listening to events of type github.CheckRunEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnCheckRunEventCreated must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnCheckRunEventCreated(callbacks ...CheckRunEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "created"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onCheckRunEvent == nil {
		g.onCheckRunEvent = make(map[string][]CheckRunEventHandleFunc)
	}
	g.onCheckRunEvent[action] = append(g.onCheckRunEvent[action], callbacks...)
}

// SetOnCheckRunEventCreated registers callbacks listening to events of type github.CheckRunEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnCheckRunEventCreatedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnCheckRunEventCreated(callbacks ...CheckRunEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "created"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onCheckRunEvent == nil {
		g.onCheckRunEvent = make(map[string][]CheckRunEventHandleFunc)
	}
	g.onCheckRunEvent[action] = callbacks
}

func (g *EventHandler) handleCheckRunEventCreated(deliveryID string, eventName string, event *github.CheckRunEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "created"
	if action != *event.Action {
		return fmt.Errorf(
			"handleCheckRunEventCreated() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleCheckRunEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onCheckRunEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onCheckRunEvent[action] {
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

// OnCheckRunEventCompleted registers callbacks listening to events of type github.CheckRunEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnCheckRunEventCompleted must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnCheckRunEventCompleted(callbacks ...CheckRunEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "completed"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onCheckRunEvent == nil {
		g.onCheckRunEvent = make(map[string][]CheckRunEventHandleFunc)
	}
	g.onCheckRunEvent[action] = append(g.onCheckRunEvent[action], callbacks...)
}

// SetOnCheckRunEventCompleted registers callbacks listening to events of type github.CheckRunEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnCheckRunEventCompletedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnCheckRunEventCompleted(callbacks ...CheckRunEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "completed"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onCheckRunEvent == nil {
		g.onCheckRunEvent = make(map[string][]CheckRunEventHandleFunc)
	}
	g.onCheckRunEvent[action] = callbacks
}

func (g *EventHandler) handleCheckRunEventCompleted(deliveryID string, eventName string, event *github.CheckRunEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "completed"
	if action != *event.Action {
		return fmt.Errorf(
			"handleCheckRunEventCompleted() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleCheckRunEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onCheckRunEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onCheckRunEvent[action] {
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

// OnCheckRunEventReRequested registers callbacks listening to events of type github.CheckRunEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnCheckRunEventReRequested must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnCheckRunEventReRequested(callbacks ...CheckRunEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "rerequested"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onCheckRunEvent == nil {
		g.onCheckRunEvent = make(map[string][]CheckRunEventHandleFunc)
	}
	g.onCheckRunEvent[action] = append(g.onCheckRunEvent[action], callbacks...)
}

// SetOnCheckRunEventReRequested registers callbacks listening to events of type github.CheckRunEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnCheckRunEventReRequestedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnCheckRunEventReRequested(callbacks ...CheckRunEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "rerequested"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onCheckRunEvent == nil {
		g.onCheckRunEvent = make(map[string][]CheckRunEventHandleFunc)
	}
	g.onCheckRunEvent[action] = callbacks
}

func (g *EventHandler) handleCheckRunEventReRequested(deliveryID string, eventName string, event *github.CheckRunEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "rerequested"
	if action != *event.Action {
		return fmt.Errorf(
			"handleCheckRunEventReRequested() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleCheckRunEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onCheckRunEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onCheckRunEvent[action] {
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

// OnCheckRunEventRequestAction registers callbacks listening to events of type github.CheckRunEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnCheckRunEventRequestAction must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnCheckRunEventRequestAction(callbacks ...CheckRunEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "requested_action"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onCheckRunEvent == nil {
		g.onCheckRunEvent = make(map[string][]CheckRunEventHandleFunc)
	}
	g.onCheckRunEvent[action] = append(g.onCheckRunEvent[action], callbacks...)
}

// SetOnCheckRunEventRequestAction registers callbacks listening to events of type github.CheckRunEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnCheckRunEventRequestActionAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnCheckRunEventRequestAction(callbacks ...CheckRunEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "requested_action"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onCheckRunEvent == nil {
		g.onCheckRunEvent = make(map[string][]CheckRunEventHandleFunc)
	}
	g.onCheckRunEvent[action] = callbacks
}

func (g *EventHandler) handleCheckRunEventRequestAction(deliveryID string, eventName string, event *github.CheckRunEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "requested_action"
	if action != *event.Action {
		return fmt.Errorf(
			"handleCheckRunEventRequestAction() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleCheckRunEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onCheckRunEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onCheckRunEvent[action] {
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

// OnCheckRunEventAny registers callbacks listening to events of type github.CheckRunEvent
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnCheckRunEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnCheckRunEventAny(callbacks ...CheckRunEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const any = "*"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onCheckRunEvent == nil {
		g.onCheckRunEvent = make(map[string][]CheckRunEventHandleFunc)
	}
	g.onCheckRunEvent[any] = append(g.onCheckRunEvent[any], callbacks...)
}

// SetOnCheckRunEventAny registers callbacks listening to events of type github.CheckRunEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnCheckRunEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnCheckRunEventAny(callbacks ...CheckRunEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const any = "*"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onCheckRunEvent == nil {
		g.onCheckRunEvent = make(map[string][]CheckRunEventHandleFunc)
	}
	g.onCheckRunEvent[any] = callbacks
}

func (g *EventHandler) handleCheckRunEventAny(deliveryID string, eventName string, event *github.CheckRunEvent) error {
	if event == nil {
		return fmt.Errorf("event was empty or nil")
	}
	const any = "*"
	if _, ok := g.onCheckRunEvent[any]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onCheckRunEvent[any] {
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

// CheckRunEvent handles github.CheckRunEvent
//
// Callbacks are executed in the following order:
//
// 1) All callbacks registered by OnBeforeAny are executed in parallel.
// 2) All callbacks registered by OnCheckRunEventAny are executed in parallel.
// 3) Optional: All callbacks registered via OnCheckRunEvent... are executed in parallel in case the Event has actions.
// 4) All callbacks registered by OnAfterAny are executed in parallel.
// on any error all callbacks registered by OnError are executed in parallel.
//
func (g *EventHandler) CheckRunEvent(deliveryID string, eventName string, event *github.CheckRunEvent) error {

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
		err := g.handleCheckRunEventCreated(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "completed":
		err := g.handleCheckRunEventCompleted(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "rerequested":
		err := g.handleCheckRunEventReRequested(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "requested_action":
		err := g.handleCheckRunEventRequestAction(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	default:
		err := g.handleCheckRunEventAny(deliveryID, eventName, event)
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
