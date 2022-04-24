package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"fmt"
	"github.com/google/go-github/v43/github"
	"golang.org/x/sync/errgroup"
)

// ProjectEventHandleFunc represents a callback function triggered on github.ProjectEvent.
// deliveryID (type: string) is the unique webhook delivery ID.
// eventName (type: string) is the name of the event.
// event (type: *github.ProjectEvent) is the webhook payload.
type ProjectEventHandleFunc func(deliveryID string, eventName string, event *github.ProjectEvent) error

// OnProjectEventCreated registers callbacks listening to events of type github.ProjectEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnProjectEventCreated must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnProjectEventCreated(callbacks ...ProjectEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "created"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectEvent == nil {
		g.onProjectEvent = make(map[string][]ProjectEventHandleFunc)
	}
	g.onProjectEvent[action] = append(g.onProjectEvent[action], callbacks...)
}

// SetOnProjectEventCreated registers callbacks listening to events of type github.ProjectEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnProjectEventCreatedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnProjectEventCreated(callbacks ...ProjectEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "created"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectEvent == nil {
		g.onProjectEvent = make(map[string][]ProjectEventHandleFunc)
	}
	g.onProjectEvent[action] = callbacks
}

func (g *EventHandler) handleProjectEventCreated(deliveryID string, eventName string, event *github.ProjectEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "created"
	if action != *event.Action {
		return fmt.Errorf(
			"handleProjectEventCreated() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleProjectEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onProjectEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onProjectEvent[action] {
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

// OnProjectEventEdited registers callbacks listening to events of type github.ProjectEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnProjectEventEdited must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnProjectEventEdited(callbacks ...ProjectEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "edited"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectEvent == nil {
		g.onProjectEvent = make(map[string][]ProjectEventHandleFunc)
	}
	g.onProjectEvent[action] = append(g.onProjectEvent[action], callbacks...)
}

// SetOnProjectEventEdited registers callbacks listening to events of type github.ProjectEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnProjectEventEditedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnProjectEventEdited(callbacks ...ProjectEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "edited"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectEvent == nil {
		g.onProjectEvent = make(map[string][]ProjectEventHandleFunc)
	}
	g.onProjectEvent[action] = callbacks
}

func (g *EventHandler) handleProjectEventEdited(deliveryID string, eventName string, event *github.ProjectEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "edited"
	if action != *event.Action {
		return fmt.Errorf(
			"handleProjectEventEdited() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleProjectEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onProjectEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onProjectEvent[action] {
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

// OnProjectEventClosed registers callbacks listening to events of type github.ProjectEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnProjectEventClosed must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnProjectEventClosed(callbacks ...ProjectEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "closed"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectEvent == nil {
		g.onProjectEvent = make(map[string][]ProjectEventHandleFunc)
	}
	g.onProjectEvent[action] = append(g.onProjectEvent[action], callbacks...)
}

// SetOnProjectEventClosed registers callbacks listening to events of type github.ProjectEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnProjectEventClosedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnProjectEventClosed(callbacks ...ProjectEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "closed"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectEvent == nil {
		g.onProjectEvent = make(map[string][]ProjectEventHandleFunc)
	}
	g.onProjectEvent[action] = callbacks
}

func (g *EventHandler) handleProjectEventClosed(deliveryID string, eventName string, event *github.ProjectEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "closed"
	if action != *event.Action {
		return fmt.Errorf(
			"handleProjectEventClosed() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleProjectEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onProjectEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onProjectEvent[action] {
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

// OnProjectEventReopened registers callbacks listening to events of type github.ProjectEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnProjectEventReopened must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnProjectEventReopened(callbacks ...ProjectEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "reopened"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectEvent == nil {
		g.onProjectEvent = make(map[string][]ProjectEventHandleFunc)
	}
	g.onProjectEvent[action] = append(g.onProjectEvent[action], callbacks...)
}

// SetOnProjectEventReopened registers callbacks listening to events of type github.ProjectEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnProjectEventReopenedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnProjectEventReopened(callbacks ...ProjectEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "reopened"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectEvent == nil {
		g.onProjectEvent = make(map[string][]ProjectEventHandleFunc)
	}
	g.onProjectEvent[action] = callbacks
}

func (g *EventHandler) handleProjectEventReopened(deliveryID string, eventName string, event *github.ProjectEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "reopened"
	if action != *event.Action {
		return fmt.Errorf(
			"handleProjectEventReopened() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleProjectEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onProjectEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onProjectEvent[action] {
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

// OnProjectEventDeleted registers callbacks listening to events of type github.ProjectEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnProjectEventDeleted must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnProjectEventDeleted(callbacks ...ProjectEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "deleted"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectEvent == nil {
		g.onProjectEvent = make(map[string][]ProjectEventHandleFunc)
	}
	g.onProjectEvent[action] = append(g.onProjectEvent[action], callbacks...)
}

// SetOnProjectEventDeleted registers callbacks listening to events of type github.ProjectEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnProjectEventDeletedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnProjectEventDeleted(callbacks ...ProjectEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "deleted"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectEvent == nil {
		g.onProjectEvent = make(map[string][]ProjectEventHandleFunc)
	}
	g.onProjectEvent[action] = callbacks
}

func (g *EventHandler) handleProjectEventDeleted(deliveryID string, eventName string, event *github.ProjectEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "deleted"
	if action != *event.Action {
		return fmt.Errorf(
			"handleProjectEventDeleted() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleProjectEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onProjectEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onProjectEvent[action] {
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

// OnProjectEventAny registers callbacks listening to events of type github.ProjectEvent
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnProjectEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnProjectEventAny(callbacks ...ProjectEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const any = "*"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectEvent == nil {
		g.onProjectEvent = make(map[string][]ProjectEventHandleFunc)
	}
	g.onProjectEvent[any] = append(g.onProjectEvent[any], callbacks...)
}

// SetOnProjectEventAny registers callbacks listening to events of type github.ProjectEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnProjectEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnProjectEventAny(callbacks ...ProjectEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const any = "*"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectEvent == nil {
		g.onProjectEvent = make(map[string][]ProjectEventHandleFunc)
	}
	g.onProjectEvent[any] = callbacks
}

func (g *EventHandler) handleProjectEventAny(deliveryID string, eventName string, event *github.ProjectEvent) error {
	if event == nil {
		return fmt.Errorf("event was empty or nil")
	}
	const any = "*"
	if _, ok := g.onProjectEvent[any]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onProjectEvent[any] {
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

// ProjectEvent handles github.ProjectEvent
//
// Callbacks are executed in the following order:
//
// 1) All callbacks registered by OnBeforeAny are executed in parallel.
// 2) All callbacks registered by OnProjectEventAny are executed in parallel.
// 3) Optional: All callbacks registered via OnProjectEvent... are executed in parallel in case the Event has actions.
// 4) All callbacks registered by OnAfterAny are executed in parallel.
// on any error all callbacks registered by OnError are executed in parallel.
//
func (g *EventHandler) ProjectEvent(deliveryID string, eventName string, event *github.ProjectEvent) error {

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
		err := g.handleProjectEventCreated(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "edited":
		err := g.handleProjectEventEdited(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "closed":
		err := g.handleProjectEventClosed(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "reopened":
		err := g.handleProjectEventReopened(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "deleted":
		err := g.handleProjectEventDeleted(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	default:
		err := g.handleProjectEventAny(deliveryID, eventName, event)
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
