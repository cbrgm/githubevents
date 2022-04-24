package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"fmt"
	"github.com/google/go-github/v43/github"
	"golang.org/x/sync/errgroup"
)

// ProjectCardEventHandleFunc represents a callback function triggered on github.ProjectCardEvent.
// deliveryID (type: string) is the unique webhook delivery ID.
// eventName (type: string) is the name of the event.
// event (type: *github.ProjectCardEvent) is the webhook payload.
type ProjectCardEventHandleFunc func(deliveryID string, eventName string, event *github.ProjectCardEvent) error

// OnProjectCardEventCreated registers callbacks listening to events of type github.ProjectCardEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnProjectCardEventCreated must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnProjectCardEventCreated(callbacks ...ProjectCardEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "created"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectCardEvent == nil {
		g.onProjectCardEvent = make(map[string][]ProjectCardEventHandleFunc)
	}
	g.onProjectCardEvent[action] = append(g.onProjectCardEvent[action], callbacks...)
}

// SetOnProjectCardEventCreated registers callbacks listening to events of type github.ProjectCardEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnProjectCardEventCreatedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnProjectCardEventCreated(callbacks ...ProjectCardEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "created"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectCardEvent == nil {
		g.onProjectCardEvent = make(map[string][]ProjectCardEventHandleFunc)
	}
	g.onProjectCardEvent[action] = callbacks
}

func (g *EventHandler) handleProjectCardEventCreated(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "created"
	if action != *event.Action {
		return fmt.Errorf(
			"handleProjectCardEventCreated() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleProjectCardEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onProjectCardEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onProjectCardEvent[action] {
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

// OnProjectCardEventEdited registers callbacks listening to events of type github.ProjectCardEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnProjectCardEventEdited must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnProjectCardEventEdited(callbacks ...ProjectCardEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "edited"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectCardEvent == nil {
		g.onProjectCardEvent = make(map[string][]ProjectCardEventHandleFunc)
	}
	g.onProjectCardEvent[action] = append(g.onProjectCardEvent[action], callbacks...)
}

// SetOnProjectCardEventEdited registers callbacks listening to events of type github.ProjectCardEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnProjectCardEventEditedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnProjectCardEventEdited(callbacks ...ProjectCardEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "edited"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectCardEvent == nil {
		g.onProjectCardEvent = make(map[string][]ProjectCardEventHandleFunc)
	}
	g.onProjectCardEvent[action] = callbacks
}

func (g *EventHandler) handleProjectCardEventEdited(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "edited"
	if action != *event.Action {
		return fmt.Errorf(
			"handleProjectCardEventEdited() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleProjectCardEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onProjectCardEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onProjectCardEvent[action] {
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

// OnProjectCardEventConverted registers callbacks listening to events of type github.ProjectCardEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnProjectCardEventConverted must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnProjectCardEventConverted(callbacks ...ProjectCardEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "converted"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectCardEvent == nil {
		g.onProjectCardEvent = make(map[string][]ProjectCardEventHandleFunc)
	}
	g.onProjectCardEvent[action] = append(g.onProjectCardEvent[action], callbacks...)
}

// SetOnProjectCardEventConverted registers callbacks listening to events of type github.ProjectCardEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnProjectCardEventConvertedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnProjectCardEventConverted(callbacks ...ProjectCardEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "converted"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectCardEvent == nil {
		g.onProjectCardEvent = make(map[string][]ProjectCardEventHandleFunc)
	}
	g.onProjectCardEvent[action] = callbacks
}

func (g *EventHandler) handleProjectCardEventConverted(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "converted"
	if action != *event.Action {
		return fmt.Errorf(
			"handleProjectCardEventConverted() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleProjectCardEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onProjectCardEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onProjectCardEvent[action] {
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

// OnProjectCardEventMoved registers callbacks listening to events of type github.ProjectCardEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnProjectCardEventMoved must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnProjectCardEventMoved(callbacks ...ProjectCardEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "moved"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectCardEvent == nil {
		g.onProjectCardEvent = make(map[string][]ProjectCardEventHandleFunc)
	}
	g.onProjectCardEvent[action] = append(g.onProjectCardEvent[action], callbacks...)
}

// SetOnProjectCardEventMoved registers callbacks listening to events of type github.ProjectCardEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnProjectCardEventMovedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnProjectCardEventMoved(callbacks ...ProjectCardEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "moved"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectCardEvent == nil {
		g.onProjectCardEvent = make(map[string][]ProjectCardEventHandleFunc)
	}
	g.onProjectCardEvent[action] = callbacks
}

func (g *EventHandler) handleProjectCardEventMoved(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "moved"
	if action != *event.Action {
		return fmt.Errorf(
			"handleProjectCardEventMoved() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleProjectCardEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onProjectCardEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onProjectCardEvent[action] {
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

// OnProjectCardEventDeleted registers callbacks listening to events of type github.ProjectCardEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnProjectCardEventDeleted must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnProjectCardEventDeleted(callbacks ...ProjectCardEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "deleted"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectCardEvent == nil {
		g.onProjectCardEvent = make(map[string][]ProjectCardEventHandleFunc)
	}
	g.onProjectCardEvent[action] = append(g.onProjectCardEvent[action], callbacks...)
}

// SetOnProjectCardEventDeleted registers callbacks listening to events of type github.ProjectCardEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnProjectCardEventDeletedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnProjectCardEventDeleted(callbacks ...ProjectCardEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "deleted"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectCardEvent == nil {
		g.onProjectCardEvent = make(map[string][]ProjectCardEventHandleFunc)
	}
	g.onProjectCardEvent[action] = callbacks
}

func (g *EventHandler) handleProjectCardEventDeleted(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "deleted"
	if action != *event.Action {
		return fmt.Errorf(
			"handleProjectCardEventDeleted() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleProjectCardEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onProjectCardEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onProjectCardEvent[action] {
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

// OnProjectCardEventAny registers callbacks listening to events of type github.ProjectCardEvent
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnProjectCardEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnProjectCardEventAny(callbacks ...ProjectCardEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const any = "*"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectCardEvent == nil {
		g.onProjectCardEvent = make(map[string][]ProjectCardEventHandleFunc)
	}
	g.onProjectCardEvent[any] = append(g.onProjectCardEvent[any], callbacks...)
}

// SetOnProjectCardEventAny registers callbacks listening to events of type github.ProjectCardEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnProjectCardEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnProjectCardEventAny(callbacks ...ProjectCardEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const any = "*"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onProjectCardEvent == nil {
		g.onProjectCardEvent = make(map[string][]ProjectCardEventHandleFunc)
	}
	g.onProjectCardEvent[any] = callbacks
}

func (g *EventHandler) handleProjectCardEventAny(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
	if event == nil {
		return fmt.Errorf("event was empty or nil")
	}
	const any = "*"
	if _, ok := g.onProjectCardEvent[any]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onProjectCardEvent[any] {
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

// ProjectCardEvent handles github.ProjectCardEvent
//
// Callbacks are executed in the following order:
//
// 1) All callbacks registered by OnBeforeAny are executed in parallel.
// 2) All callbacks registered by OnProjectCardEventAny are executed in parallel.
// 3) Optional: All callbacks registered via OnProjectCardEvent... are executed in parallel in case the Event has actions.
// 4) All callbacks registered by OnAfterAny are executed in parallel.
// on any error all callbacks registered by OnError are executed in parallel.
//
func (g *EventHandler) ProjectCardEvent(deliveryID string, eventName string, event *github.ProjectCardEvent) error {

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
		err := g.handleProjectCardEventCreated(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "edited":
		err := g.handleProjectCardEventEdited(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "converted":
		err := g.handleProjectCardEventConverted(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "moved":
		err := g.handleProjectCardEventMoved(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "deleted":
		err := g.handleProjectCardEventDeleted(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	default:
		err := g.handleProjectCardEventAny(deliveryID, eventName, event)
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
