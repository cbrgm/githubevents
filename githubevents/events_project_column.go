package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"fmt"
	"github.com/google/go-github/v43/github"
	"golang.org/x/sync/errgroup"
)

// ProjectColumnEventHandleFunc represents a callback function triggered on github.ProjectColumnEvent.
// deliveryID (type: string) is the unique webhook delivery ID.
// eventName (type: string) is the name of the event.
// event (type: *github.ProjectColumnEvent) is the webhook payload.
type ProjectColumnEventHandleFunc func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error

// OnProjectColumnEventCreated registers callbacks listening to events of type github.ProjectColumnEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnProjectColumnEventCreated must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnProjectColumnEventCreated(callbacks ...ProjectColumnEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "created"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectColumnEvent == nil {
		g.onProjectColumnEvent = make(map[string][]ProjectColumnEventHandleFunc)
	}
	g.onProjectColumnEvent[action] = append(g.onProjectColumnEvent[action], callbacks...)
}

// SetOnProjectColumnEventCreated registers callbacks listening to events of type github.ProjectColumnEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnProjectColumnEventCreatedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnProjectColumnEventCreated(callbacks ...ProjectColumnEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "created"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectColumnEvent == nil {
		g.onProjectColumnEvent = make(map[string][]ProjectColumnEventHandleFunc)
	}
	g.onProjectColumnEvent[action] = callbacks
}

func (g *EventHandler) handleProjectColumnEventCreated(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "created"
	if action != *event.Action {
		return fmt.Errorf(
			"handleProjectColumnEventCreated() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleProjectColumnEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onProjectColumnEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onProjectColumnEvent[action] {
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

// OnProjectColumnEventEdited registers callbacks listening to events of type github.ProjectColumnEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnProjectColumnEventEdited must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnProjectColumnEventEdited(callbacks ...ProjectColumnEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "edited"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectColumnEvent == nil {
		g.onProjectColumnEvent = make(map[string][]ProjectColumnEventHandleFunc)
	}
	g.onProjectColumnEvent[action] = append(g.onProjectColumnEvent[action], callbacks...)
}

// SetOnProjectColumnEventEdited registers callbacks listening to events of type github.ProjectColumnEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnProjectColumnEventEditedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnProjectColumnEventEdited(callbacks ...ProjectColumnEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "edited"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectColumnEvent == nil {
		g.onProjectColumnEvent = make(map[string][]ProjectColumnEventHandleFunc)
	}
	g.onProjectColumnEvent[action] = callbacks
}

func (g *EventHandler) handleProjectColumnEventEdited(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "edited"
	if action != *event.Action {
		return fmt.Errorf(
			"handleProjectColumnEventEdited() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleProjectColumnEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onProjectColumnEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onProjectColumnEvent[action] {
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

// OnProjectColumnEventMoved registers callbacks listening to events of type github.ProjectColumnEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnProjectColumnEventMoved must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnProjectColumnEventMoved(callbacks ...ProjectColumnEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "moved"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectColumnEvent == nil {
		g.onProjectColumnEvent = make(map[string][]ProjectColumnEventHandleFunc)
	}
	g.onProjectColumnEvent[action] = append(g.onProjectColumnEvent[action], callbacks...)
}

// SetOnProjectColumnEventMoved registers callbacks listening to events of type github.ProjectColumnEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnProjectColumnEventMovedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnProjectColumnEventMoved(callbacks ...ProjectColumnEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "moved"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectColumnEvent == nil {
		g.onProjectColumnEvent = make(map[string][]ProjectColumnEventHandleFunc)
	}
	g.onProjectColumnEvent[action] = callbacks
}

func (g *EventHandler) handleProjectColumnEventMoved(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "moved"
	if action != *event.Action {
		return fmt.Errorf(
			"handleProjectColumnEventMoved() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleProjectColumnEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onProjectColumnEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onProjectColumnEvent[action] {
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

// OnProjectColumnEventDeleted registers callbacks listening to events of type github.ProjectColumnEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnProjectColumnEventDeleted must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnProjectColumnEventDeleted(callbacks ...ProjectColumnEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "deleted"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectColumnEvent == nil {
		g.onProjectColumnEvent = make(map[string][]ProjectColumnEventHandleFunc)
	}
	g.onProjectColumnEvent[action] = append(g.onProjectColumnEvent[action], callbacks...)
}

// SetOnProjectColumnEventDeleted registers callbacks listening to events of type github.ProjectColumnEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnProjectColumnEventDeletedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnProjectColumnEventDeleted(callbacks ...ProjectColumnEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "deleted"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectColumnEvent == nil {
		g.onProjectColumnEvent = make(map[string][]ProjectColumnEventHandleFunc)
	}
	g.onProjectColumnEvent[action] = callbacks
}

func (g *EventHandler) handleProjectColumnEventDeleted(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "deleted"
	if action != *event.Action {
		return fmt.Errorf(
			"handleProjectColumnEventDeleted() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleProjectColumnEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onProjectColumnEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onProjectColumnEvent[action] {
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

// OnProjectColumnEventAny registers callbacks listening to events of type github.ProjectColumnEvent
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnProjectColumnEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnProjectColumnEventAny(callbacks ...ProjectColumnEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const any = "*"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectColumnEvent == nil {
		g.onProjectColumnEvent = make(map[string][]ProjectColumnEventHandleFunc)
	}
	g.onProjectColumnEvent[any] = append(g.onProjectColumnEvent[any], callbacks...)
}

// SetOnProjectColumnEventAny registers callbacks listening to events of type github.ProjectColumnEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnProjectColumnEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnProjectColumnEventAny(callbacks ...ProjectColumnEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const any = "*"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectColumnEvent == nil {
		g.onProjectColumnEvent = make(map[string][]ProjectColumnEventHandleFunc)
	}
	g.onProjectColumnEvent[any] = callbacks
}

func (g *EventHandler) handleProjectColumnEventAny(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
	if event == nil {
		return fmt.Errorf("event was empty or nil")
	}
	const any = "*"
	if _, ok := g.onProjectColumnEvent[any]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onProjectColumnEvent[any] {
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

// ProjectColumnEvent handles github.ProjectColumnEvent
//
// Callbacks are executed in the following order:
//
// 1) All callbacks registered by OnBeforeAny are executed in parallel.
// 2) All callbacks registered by OnProjectColumnEventAny are executed in parallel.
// 3) Optional: All callbacks registered via OnProjectColumnEvent... are executed in parallel in case the Event has actions.
// 4) All callbacks registered by OnAfterAny are executed in parallel.
// on any error all callbacks registered by OnError are executed in parallel.
//
func (g *EventHandler) ProjectColumnEvent(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {

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
		err := g.handleProjectColumnEventCreated(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "edited":
		err := g.handleProjectColumnEventEdited(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "moved":
		err := g.handleProjectColumnEventMoved(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "deleted":
		err := g.handleProjectColumnEventDeleted(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	default:
		err := g.handleProjectColumnEventAny(deliveryID, eventName, event)
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
