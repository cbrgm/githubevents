package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"fmt"
	"github.com/google/go-github/v43/github"
	"golang.org/x/sync/errgroup"
)

// ReleaseEventHandleFunc represents a callback function triggered on github.ReleaseEvent.
// deliveryID (type: string) is the unique webhook delivery ID.
// eventName (type: string) is the name of the event.
// event (type: *github.ReleaseEvent) is the webhook payload.
type ReleaseEventHandleFunc func(deliveryID string, eventName string, event *github.ReleaseEvent) error

// OnReleaseEventPublished registers callbacks listening to events of type github.ReleaseEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnReleaseEventPublished must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnReleaseEventPublished(callbacks ...ReleaseEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "published"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onReleaseEvent == nil {
		g.onReleaseEvent = make(map[string][]ReleaseEventHandleFunc)
	}
	g.onReleaseEvent[action] = append(g.onReleaseEvent[action], callbacks...)
}

// SetOnReleaseEventPublished registers callbacks listening to events of type github.ReleaseEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnReleaseEventPublishedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnReleaseEventPublished(callbacks ...ReleaseEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "published"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onReleaseEvent == nil {
		g.onReleaseEvent = make(map[string][]ReleaseEventHandleFunc)
	}
	g.onReleaseEvent[action] = callbacks
}

func (g *EventHandler) handleReleaseEventPublished(deliveryID string, eventName string, event *github.ReleaseEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "published"
	if action != *event.Action {
		return fmt.Errorf(
			"handleReleaseEventPublished() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleReleaseEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onReleaseEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onReleaseEvent[action] {
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

// OnReleaseEventUnpublished registers callbacks listening to events of type github.ReleaseEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnReleaseEventUnpublished must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnReleaseEventUnpublished(callbacks ...ReleaseEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "unpublished"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onReleaseEvent == nil {
		g.onReleaseEvent = make(map[string][]ReleaseEventHandleFunc)
	}
	g.onReleaseEvent[action] = append(g.onReleaseEvent[action], callbacks...)
}

// SetOnReleaseEventUnpublished registers callbacks listening to events of type github.ReleaseEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnReleaseEventUnpublishedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnReleaseEventUnpublished(callbacks ...ReleaseEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "unpublished"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onReleaseEvent == nil {
		g.onReleaseEvent = make(map[string][]ReleaseEventHandleFunc)
	}
	g.onReleaseEvent[action] = callbacks
}

func (g *EventHandler) handleReleaseEventUnpublished(deliveryID string, eventName string, event *github.ReleaseEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "unpublished"
	if action != *event.Action {
		return fmt.Errorf(
			"handleReleaseEventUnpublished() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleReleaseEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onReleaseEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onReleaseEvent[action] {
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

// OnReleaseEventCreated registers callbacks listening to events of type github.ReleaseEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnReleaseEventCreated must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnReleaseEventCreated(callbacks ...ReleaseEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "created"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onReleaseEvent == nil {
		g.onReleaseEvent = make(map[string][]ReleaseEventHandleFunc)
	}
	g.onReleaseEvent[action] = append(g.onReleaseEvent[action], callbacks...)
}

// SetOnReleaseEventCreated registers callbacks listening to events of type github.ReleaseEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnReleaseEventCreatedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnReleaseEventCreated(callbacks ...ReleaseEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "created"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onReleaseEvent == nil {
		g.onReleaseEvent = make(map[string][]ReleaseEventHandleFunc)
	}
	g.onReleaseEvent[action] = callbacks
}

func (g *EventHandler) handleReleaseEventCreated(deliveryID string, eventName string, event *github.ReleaseEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "created"
	if action != *event.Action {
		return fmt.Errorf(
			"handleReleaseEventCreated() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleReleaseEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onReleaseEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onReleaseEvent[action] {
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

// OnReleaseEventEdited registers callbacks listening to events of type github.ReleaseEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnReleaseEventEdited must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnReleaseEventEdited(callbacks ...ReleaseEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "edited"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onReleaseEvent == nil {
		g.onReleaseEvent = make(map[string][]ReleaseEventHandleFunc)
	}
	g.onReleaseEvent[action] = append(g.onReleaseEvent[action], callbacks...)
}

// SetOnReleaseEventEdited registers callbacks listening to events of type github.ReleaseEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnReleaseEventEditedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnReleaseEventEdited(callbacks ...ReleaseEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "edited"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onReleaseEvent == nil {
		g.onReleaseEvent = make(map[string][]ReleaseEventHandleFunc)
	}
	g.onReleaseEvent[action] = callbacks
}

func (g *EventHandler) handleReleaseEventEdited(deliveryID string, eventName string, event *github.ReleaseEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "edited"
	if action != *event.Action {
		return fmt.Errorf(
			"handleReleaseEventEdited() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleReleaseEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onReleaseEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onReleaseEvent[action] {
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

// OnReleaseEventDeleted registers callbacks listening to events of type github.ReleaseEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnReleaseEventDeleted must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnReleaseEventDeleted(callbacks ...ReleaseEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "deleted"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onReleaseEvent == nil {
		g.onReleaseEvent = make(map[string][]ReleaseEventHandleFunc)
	}
	g.onReleaseEvent[action] = append(g.onReleaseEvent[action], callbacks...)
}

// SetOnReleaseEventDeleted registers callbacks listening to events of type github.ReleaseEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnReleaseEventDeletedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnReleaseEventDeleted(callbacks ...ReleaseEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "deleted"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onReleaseEvent == nil {
		g.onReleaseEvent = make(map[string][]ReleaseEventHandleFunc)
	}
	g.onReleaseEvent[action] = callbacks
}

func (g *EventHandler) handleReleaseEventDeleted(deliveryID string, eventName string, event *github.ReleaseEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "deleted"
	if action != *event.Action {
		return fmt.Errorf(
			"handleReleaseEventDeleted() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleReleaseEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onReleaseEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onReleaseEvent[action] {
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

// OnReleaseEventPreReleased registers callbacks listening to events of type github.ReleaseEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnReleaseEventPreReleased must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnReleaseEventPreReleased(callbacks ...ReleaseEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "prereleased"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onReleaseEvent == nil {
		g.onReleaseEvent = make(map[string][]ReleaseEventHandleFunc)
	}
	g.onReleaseEvent[action] = append(g.onReleaseEvent[action], callbacks...)
}

// SetOnReleaseEventPreReleased registers callbacks listening to events of type github.ReleaseEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnReleaseEventPreReleasedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnReleaseEventPreReleased(callbacks ...ReleaseEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "prereleased"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onReleaseEvent == nil {
		g.onReleaseEvent = make(map[string][]ReleaseEventHandleFunc)
	}
	g.onReleaseEvent[action] = callbacks
}

func (g *EventHandler) handleReleaseEventPreReleased(deliveryID string, eventName string, event *github.ReleaseEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "prereleased"
	if action != *event.Action {
		return fmt.Errorf(
			"handleReleaseEventPreReleased() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleReleaseEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onReleaseEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onReleaseEvent[action] {
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

// OnReleaseEventReleased registers callbacks listening to events of type github.ReleaseEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnReleaseEventReleased must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnReleaseEventReleased(callbacks ...ReleaseEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "released"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onReleaseEvent == nil {
		g.onReleaseEvent = make(map[string][]ReleaseEventHandleFunc)
	}
	g.onReleaseEvent[action] = append(g.onReleaseEvent[action], callbacks...)
}

// SetOnReleaseEventReleased registers callbacks listening to events of type github.ReleaseEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnReleaseEventReleasedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnReleaseEventReleased(callbacks ...ReleaseEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "released"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onReleaseEvent == nil {
		g.onReleaseEvent = make(map[string][]ReleaseEventHandleFunc)
	}
	g.onReleaseEvent[action] = callbacks
}

func (g *EventHandler) handleReleaseEventReleased(deliveryID string, eventName string, event *github.ReleaseEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "released"
	if action != *event.Action {
		return fmt.Errorf(
			"handleReleaseEventReleased() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleReleaseEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onReleaseEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onReleaseEvent[action] {
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

// OnReleaseEventAny registers callbacks listening to events of type github.ReleaseEvent
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnReleaseEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnReleaseEventAny(callbacks ...ReleaseEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const any = "*"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onReleaseEvent == nil {
		g.onReleaseEvent = make(map[string][]ReleaseEventHandleFunc)
	}
	g.onReleaseEvent[any] = append(g.onReleaseEvent[any], callbacks...)
}

// SetOnReleaseEventAny registers callbacks listening to events of type github.ReleaseEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnReleaseEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnReleaseEventAny(callbacks ...ReleaseEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const any = "*"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onReleaseEvent == nil {
		g.onReleaseEvent = make(map[string][]ReleaseEventHandleFunc)
	}
	g.onReleaseEvent[any] = callbacks
}

func (g *EventHandler) handleReleaseEventAny(deliveryID string, eventName string, event *github.ReleaseEvent) error {
	if event == nil {
		return fmt.Errorf("event was empty or nil")
	}
	const any = "*"
	if _, ok := g.onReleaseEvent[any]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onReleaseEvent[any] {
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

// ReleaseEvent handles github.ReleaseEvent
//
// Callbacks are executed in the following order:
//
// 1) All callbacks registered by OnBeforeAny are executed in parallel.
// 2) All callbacks registered by OnReleaseEventAny are executed in parallel.
// 3) Optional: All callbacks registered via OnReleaseEvent... are executed in parallel in case the Event has actions.
// 4) All callbacks registered by OnAfterAny are executed in parallel.
// on any error all callbacks registered by OnError are executed in parallel.
//
func (g *EventHandler) ReleaseEvent(deliveryID string, eventName string, event *github.ReleaseEvent) error {

	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	action := *event.Action

	err := g.handleBeforeAny(deliveryID, eventName, event)
	if err != nil {
		return g.handleError(deliveryID, eventName, event, err)
	}

	switch action {

	case "published":
		err := g.handleReleaseEventPublished(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "unpublished":
		err := g.handleReleaseEventUnpublished(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "created":
		err := g.handleReleaseEventCreated(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "edited":
		err := g.handleReleaseEventEdited(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "deleted":
		err := g.handleReleaseEventDeleted(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "prereleased":
		err := g.handleReleaseEventPreReleased(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "released":
		err := g.handleReleaseEventReleased(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	default:
		err := g.handleReleaseEventAny(deliveryID, eventName, event)
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
