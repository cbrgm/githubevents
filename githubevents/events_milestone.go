package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"fmt"
	"github.com/google/go-github/v43/github"
	"golang.org/x/sync/errgroup"
)

// MilestoneEventHandleFunc represents a callback function triggered on github.MilestoneEvent.
// deliveryID (type: string) is the unique webhook delivery ID.
// eventName (type: string) is the name of the event.
// event (type: *github.MilestoneEvent) is the webhook payload.
type MilestoneEventHandleFunc func(deliveryID string, eventName string, event *github.MilestoneEvent) error

// OnMilestoneEventCreated registers callbacks listening to events of type github.MilestoneEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnMilestoneEventCreated must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnMilestoneEventCreated(callbacks ...MilestoneEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "created"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onMilestoneEvent == nil {
		g.onMilestoneEvent = make(map[string][]MilestoneEventHandleFunc)
	}
	g.onMilestoneEvent[action] = append(g.onMilestoneEvent[action], callbacks...)
}

// SetOnMilestoneEventCreated registers callbacks listening to events of type github.MilestoneEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnMilestoneEventCreatedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnMilestoneEventCreated(callbacks ...MilestoneEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "created"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onMilestoneEvent == nil {
		g.onMilestoneEvent = make(map[string][]MilestoneEventHandleFunc)
	}
	g.onMilestoneEvent[action] = callbacks
}

func (g *EventHandler) handleMilestoneEventCreated(deliveryID string, eventName string, event *github.MilestoneEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "created"
	if action != *event.Action {
		return fmt.Errorf(
			"handleMilestoneEventCreated() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleMilestoneEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onMilestoneEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onMilestoneEvent[action] {
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

// OnMilestoneEventClosed registers callbacks listening to events of type github.MilestoneEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnMilestoneEventClosed must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnMilestoneEventClosed(callbacks ...MilestoneEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "closed"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onMilestoneEvent == nil {
		g.onMilestoneEvent = make(map[string][]MilestoneEventHandleFunc)
	}
	g.onMilestoneEvent[action] = append(g.onMilestoneEvent[action], callbacks...)
}

// SetOnMilestoneEventClosed registers callbacks listening to events of type github.MilestoneEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnMilestoneEventClosedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnMilestoneEventClosed(callbacks ...MilestoneEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "closed"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onMilestoneEvent == nil {
		g.onMilestoneEvent = make(map[string][]MilestoneEventHandleFunc)
	}
	g.onMilestoneEvent[action] = callbacks
}

func (g *EventHandler) handleMilestoneEventClosed(deliveryID string, eventName string, event *github.MilestoneEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "closed"
	if action != *event.Action {
		return fmt.Errorf(
			"handleMilestoneEventClosed() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleMilestoneEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onMilestoneEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onMilestoneEvent[action] {
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

// OnMilestoneEventOpened registers callbacks listening to events of type github.MilestoneEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnMilestoneEventOpened must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnMilestoneEventOpened(callbacks ...MilestoneEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "opened"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onMilestoneEvent == nil {
		g.onMilestoneEvent = make(map[string][]MilestoneEventHandleFunc)
	}
	g.onMilestoneEvent[action] = append(g.onMilestoneEvent[action], callbacks...)
}

// SetOnMilestoneEventOpened registers callbacks listening to events of type github.MilestoneEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnMilestoneEventOpenedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnMilestoneEventOpened(callbacks ...MilestoneEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "opened"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onMilestoneEvent == nil {
		g.onMilestoneEvent = make(map[string][]MilestoneEventHandleFunc)
	}
	g.onMilestoneEvent[action] = callbacks
}

func (g *EventHandler) handleMilestoneEventOpened(deliveryID string, eventName string, event *github.MilestoneEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "opened"
	if action != *event.Action {
		return fmt.Errorf(
			"handleMilestoneEventOpened() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleMilestoneEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onMilestoneEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onMilestoneEvent[action] {
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

// OnMilestoneEventEdited registers callbacks listening to events of type github.MilestoneEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnMilestoneEventEdited must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnMilestoneEventEdited(callbacks ...MilestoneEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "edited"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onMilestoneEvent == nil {
		g.onMilestoneEvent = make(map[string][]MilestoneEventHandleFunc)
	}
	g.onMilestoneEvent[action] = append(g.onMilestoneEvent[action], callbacks...)
}

// SetOnMilestoneEventEdited registers callbacks listening to events of type github.MilestoneEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnMilestoneEventEditedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnMilestoneEventEdited(callbacks ...MilestoneEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "edited"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onMilestoneEvent == nil {
		g.onMilestoneEvent = make(map[string][]MilestoneEventHandleFunc)
	}
	g.onMilestoneEvent[action] = callbacks
}

func (g *EventHandler) handleMilestoneEventEdited(deliveryID string, eventName string, event *github.MilestoneEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "edited"
	if action != *event.Action {
		return fmt.Errorf(
			"handleMilestoneEventEdited() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleMilestoneEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onMilestoneEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onMilestoneEvent[action] {
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

// OnMilestoneEventDeleted registers callbacks listening to events of type github.MilestoneEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnMilestoneEventDeleted must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnMilestoneEventDeleted(callbacks ...MilestoneEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "deleted"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onMilestoneEvent == nil {
		g.onMilestoneEvent = make(map[string][]MilestoneEventHandleFunc)
	}
	g.onMilestoneEvent[action] = append(g.onMilestoneEvent[action], callbacks...)
}

// SetOnMilestoneEventDeleted registers callbacks listening to events of type github.MilestoneEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnMilestoneEventDeletedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnMilestoneEventDeleted(callbacks ...MilestoneEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "deleted"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onMilestoneEvent == nil {
		g.onMilestoneEvent = make(map[string][]MilestoneEventHandleFunc)
	}
	g.onMilestoneEvent[action] = callbacks
}

func (g *EventHandler) handleMilestoneEventDeleted(deliveryID string, eventName string, event *github.MilestoneEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "deleted"
	if action != *event.Action {
		return fmt.Errorf(
			"handleMilestoneEventDeleted() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleMilestoneEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onMilestoneEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onMilestoneEvent[action] {
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

// OnMilestoneEventAny registers callbacks listening to events of type github.MilestoneEvent
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnMilestoneEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnMilestoneEventAny(callbacks ...MilestoneEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const any = "*"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onMilestoneEvent == nil {
		g.onMilestoneEvent = make(map[string][]MilestoneEventHandleFunc)
	}
	g.onMilestoneEvent[any] = append(g.onMilestoneEvent[any], callbacks...)
}

// SetOnMilestoneEventAny registers callbacks listening to events of type github.MilestoneEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnMilestoneEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnMilestoneEventAny(callbacks ...MilestoneEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const any = "*"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onMilestoneEvent == nil {
		g.onMilestoneEvent = make(map[string][]MilestoneEventHandleFunc)
	}
	g.onMilestoneEvent[any] = callbacks
}

func (g *EventHandler) handleMilestoneEventAny(deliveryID string, eventName string, event *github.MilestoneEvent) error {
	if event == nil {
		return fmt.Errorf("event was empty or nil")
	}
	const any = "*"
	if _, ok := g.onMilestoneEvent[any]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onMilestoneEvent[any] {
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

// MilestoneEvent handles github.MilestoneEvent
//
// Callbacks are executed in the following order:
//
// 1) All callbacks registered by OnBeforeAny are executed in parallel.
// 2) All callbacks registered by OnMilestoneEventAny are executed in parallel.
// 3) Optional: All callbacks registered via OnMilestoneEvent... are executed in parallel in case the Event has actions.
// 4) All callbacks registered by OnAfterAny are executed in parallel.
// on any error all callbacks registered by OnError are executed in parallel.
//
func (g *EventHandler) MilestoneEvent(deliveryID string, eventName string, event *github.MilestoneEvent) error {

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
		err := g.handleMilestoneEventCreated(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "closed":
		err := g.handleMilestoneEventClosed(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "opened":
		err := g.handleMilestoneEventOpened(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "edited":
		err := g.handleMilestoneEventEdited(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "deleted":
		err := g.handleMilestoneEventDeleted(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	default:
		err := g.handleMilestoneEventAny(deliveryID, eventName, event)
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
