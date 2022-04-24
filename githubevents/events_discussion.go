package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"fmt"
	"github.com/google/go-github/v43/github"
	"golang.org/x/sync/errgroup"
)

// DiscussionEventHandleFunc represents a callback function triggered on github.DiscussionEvent.
// deliveryID (type: string) is the unique webhook delivery ID.
// eventName (type: string) is the name of the event.
// event (type: *github.DiscussionEvent) is the webhook payload.
type DiscussionEventHandleFunc func(deliveryID string, eventName string, event *github.DiscussionEvent) error

// OnDiscussionEventCreated registers callbacks listening to events of type github.DiscussionEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnDiscussionEventCreated must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnDiscussionEventCreated(callbacks ...DiscussionEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "created"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onDiscussionEvent == nil {
		g.onDiscussionEvent = make(map[string][]DiscussionEventHandleFunc)
	}
	g.onDiscussionEvent[action] = append(g.onDiscussionEvent[action], callbacks...)
}

// SetOnDiscussionEventCreated registers callbacks listening to events of type github.DiscussionEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnDiscussionEventCreatedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnDiscussionEventCreated(callbacks ...DiscussionEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "created"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onDiscussionEvent == nil {
		g.onDiscussionEvent = make(map[string][]DiscussionEventHandleFunc)
	}
	g.onDiscussionEvent[action] = callbacks
}

func (g *EventHandler) handleDiscussionEventCreated(deliveryID string, eventName string, event *github.DiscussionEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "created"
	if action != *event.Action {
		return fmt.Errorf(
			"handleDiscussionEventCreated() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleDiscussionEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onDiscussionEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onDiscussionEvent[action] {
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

// OnDiscussionEventEdited registers callbacks listening to events of type github.DiscussionEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnDiscussionEventEdited must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnDiscussionEventEdited(callbacks ...DiscussionEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "edited"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onDiscussionEvent == nil {
		g.onDiscussionEvent = make(map[string][]DiscussionEventHandleFunc)
	}
	g.onDiscussionEvent[action] = append(g.onDiscussionEvent[action], callbacks...)
}

// SetOnDiscussionEventEdited registers callbacks listening to events of type github.DiscussionEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnDiscussionEventEditedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnDiscussionEventEdited(callbacks ...DiscussionEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "edited"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onDiscussionEvent == nil {
		g.onDiscussionEvent = make(map[string][]DiscussionEventHandleFunc)
	}
	g.onDiscussionEvent[action] = callbacks
}

func (g *EventHandler) handleDiscussionEventEdited(deliveryID string, eventName string, event *github.DiscussionEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "edited"
	if action != *event.Action {
		return fmt.Errorf(
			"handleDiscussionEventEdited() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleDiscussionEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onDiscussionEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onDiscussionEvent[action] {
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

// OnDiscussionEventDeleted registers callbacks listening to events of type github.DiscussionEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnDiscussionEventDeleted must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnDiscussionEventDeleted(callbacks ...DiscussionEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "deleted"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onDiscussionEvent == nil {
		g.onDiscussionEvent = make(map[string][]DiscussionEventHandleFunc)
	}
	g.onDiscussionEvent[action] = append(g.onDiscussionEvent[action], callbacks...)
}

// SetOnDiscussionEventDeleted registers callbacks listening to events of type github.DiscussionEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnDiscussionEventDeletedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnDiscussionEventDeleted(callbacks ...DiscussionEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "deleted"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onDiscussionEvent == nil {
		g.onDiscussionEvent = make(map[string][]DiscussionEventHandleFunc)
	}
	g.onDiscussionEvent[action] = callbacks
}

func (g *EventHandler) handleDiscussionEventDeleted(deliveryID string, eventName string, event *github.DiscussionEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "deleted"
	if action != *event.Action {
		return fmt.Errorf(
			"handleDiscussionEventDeleted() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleDiscussionEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onDiscussionEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onDiscussionEvent[action] {
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

// OnDiscussionEventPinned registers callbacks listening to events of type github.DiscussionEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnDiscussionEventPinned must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnDiscussionEventPinned(callbacks ...DiscussionEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "pinned"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onDiscussionEvent == nil {
		g.onDiscussionEvent = make(map[string][]DiscussionEventHandleFunc)
	}
	g.onDiscussionEvent[action] = append(g.onDiscussionEvent[action], callbacks...)
}

// SetOnDiscussionEventPinned registers callbacks listening to events of type github.DiscussionEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnDiscussionEventPinnedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnDiscussionEventPinned(callbacks ...DiscussionEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "pinned"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onDiscussionEvent == nil {
		g.onDiscussionEvent = make(map[string][]DiscussionEventHandleFunc)
	}
	g.onDiscussionEvent[action] = callbacks
}

func (g *EventHandler) handleDiscussionEventPinned(deliveryID string, eventName string, event *github.DiscussionEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "pinned"
	if action != *event.Action {
		return fmt.Errorf(
			"handleDiscussionEventPinned() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleDiscussionEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onDiscussionEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onDiscussionEvent[action] {
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

// OnDiscussionEventUnpinned registers callbacks listening to events of type github.DiscussionEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnDiscussionEventUnpinned must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnDiscussionEventUnpinned(callbacks ...DiscussionEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "unpinned"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onDiscussionEvent == nil {
		g.onDiscussionEvent = make(map[string][]DiscussionEventHandleFunc)
	}
	g.onDiscussionEvent[action] = append(g.onDiscussionEvent[action], callbacks...)
}

// SetOnDiscussionEventUnpinned registers callbacks listening to events of type github.DiscussionEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnDiscussionEventUnpinnedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnDiscussionEventUnpinned(callbacks ...DiscussionEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "unpinned"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onDiscussionEvent == nil {
		g.onDiscussionEvent = make(map[string][]DiscussionEventHandleFunc)
	}
	g.onDiscussionEvent[action] = callbacks
}

func (g *EventHandler) handleDiscussionEventUnpinned(deliveryID string, eventName string, event *github.DiscussionEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "unpinned"
	if action != *event.Action {
		return fmt.Errorf(
			"handleDiscussionEventUnpinned() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleDiscussionEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onDiscussionEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onDiscussionEvent[action] {
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

// OnDiscussionEventLocked registers callbacks listening to events of type github.DiscussionEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnDiscussionEventLocked must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnDiscussionEventLocked(callbacks ...DiscussionEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "locked"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onDiscussionEvent == nil {
		g.onDiscussionEvent = make(map[string][]DiscussionEventHandleFunc)
	}
	g.onDiscussionEvent[action] = append(g.onDiscussionEvent[action], callbacks...)
}

// SetOnDiscussionEventLocked registers callbacks listening to events of type github.DiscussionEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnDiscussionEventLockedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnDiscussionEventLocked(callbacks ...DiscussionEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "locked"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onDiscussionEvent == nil {
		g.onDiscussionEvent = make(map[string][]DiscussionEventHandleFunc)
	}
	g.onDiscussionEvent[action] = callbacks
}

func (g *EventHandler) handleDiscussionEventLocked(deliveryID string, eventName string, event *github.DiscussionEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "locked"
	if action != *event.Action {
		return fmt.Errorf(
			"handleDiscussionEventLocked() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleDiscussionEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onDiscussionEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onDiscussionEvent[action] {
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

// OnDiscussionEventUnlocked registers callbacks listening to events of type github.DiscussionEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnDiscussionEventUnlocked must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnDiscussionEventUnlocked(callbacks ...DiscussionEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "unlocked"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onDiscussionEvent == nil {
		g.onDiscussionEvent = make(map[string][]DiscussionEventHandleFunc)
	}
	g.onDiscussionEvent[action] = append(g.onDiscussionEvent[action], callbacks...)
}

// SetOnDiscussionEventUnlocked registers callbacks listening to events of type github.DiscussionEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnDiscussionEventUnlockedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnDiscussionEventUnlocked(callbacks ...DiscussionEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "unlocked"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onDiscussionEvent == nil {
		g.onDiscussionEvent = make(map[string][]DiscussionEventHandleFunc)
	}
	g.onDiscussionEvent[action] = callbacks
}

func (g *EventHandler) handleDiscussionEventUnlocked(deliveryID string, eventName string, event *github.DiscussionEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "unlocked"
	if action != *event.Action {
		return fmt.Errorf(
			"handleDiscussionEventUnlocked() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleDiscussionEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onDiscussionEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onDiscussionEvent[action] {
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

// OnDiscussionEventTransferred registers callbacks listening to events of type github.DiscussionEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnDiscussionEventTransferred must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnDiscussionEventTransferred(callbacks ...DiscussionEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "transferred"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onDiscussionEvent == nil {
		g.onDiscussionEvent = make(map[string][]DiscussionEventHandleFunc)
	}
	g.onDiscussionEvent[action] = append(g.onDiscussionEvent[action], callbacks...)
}

// SetOnDiscussionEventTransferred registers callbacks listening to events of type github.DiscussionEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnDiscussionEventTransferredAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnDiscussionEventTransferred(callbacks ...DiscussionEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "transferred"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onDiscussionEvent == nil {
		g.onDiscussionEvent = make(map[string][]DiscussionEventHandleFunc)
	}
	g.onDiscussionEvent[action] = callbacks
}

func (g *EventHandler) handleDiscussionEventTransferred(deliveryID string, eventName string, event *github.DiscussionEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "transferred"
	if action != *event.Action {
		return fmt.Errorf(
			"handleDiscussionEventTransferred() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleDiscussionEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onDiscussionEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onDiscussionEvent[action] {
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

// OnDiscussionEventCategoryChanged registers callbacks listening to events of type github.DiscussionEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnDiscussionEventCategoryChanged must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnDiscussionEventCategoryChanged(callbacks ...DiscussionEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "category_changed"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onDiscussionEvent == nil {
		g.onDiscussionEvent = make(map[string][]DiscussionEventHandleFunc)
	}
	g.onDiscussionEvent[action] = append(g.onDiscussionEvent[action], callbacks...)
}

// SetOnDiscussionEventCategoryChanged registers callbacks listening to events of type github.DiscussionEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnDiscussionEventCategoryChangedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnDiscussionEventCategoryChanged(callbacks ...DiscussionEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "category_changed"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onDiscussionEvent == nil {
		g.onDiscussionEvent = make(map[string][]DiscussionEventHandleFunc)
	}
	g.onDiscussionEvent[action] = callbacks
}

func (g *EventHandler) handleDiscussionEventCategoryChanged(deliveryID string, eventName string, event *github.DiscussionEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "category_changed"
	if action != *event.Action {
		return fmt.Errorf(
			"handleDiscussionEventCategoryChanged() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleDiscussionEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onDiscussionEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onDiscussionEvent[action] {
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

// OnDiscussionEventAnswered registers callbacks listening to events of type github.DiscussionEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnDiscussionEventAnswered must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnDiscussionEventAnswered(callbacks ...DiscussionEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "answered"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onDiscussionEvent == nil {
		g.onDiscussionEvent = make(map[string][]DiscussionEventHandleFunc)
	}
	g.onDiscussionEvent[action] = append(g.onDiscussionEvent[action], callbacks...)
}

// SetOnDiscussionEventAnswered registers callbacks listening to events of type github.DiscussionEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnDiscussionEventAnsweredAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnDiscussionEventAnswered(callbacks ...DiscussionEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "answered"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onDiscussionEvent == nil {
		g.onDiscussionEvent = make(map[string][]DiscussionEventHandleFunc)
	}
	g.onDiscussionEvent[action] = callbacks
}

func (g *EventHandler) handleDiscussionEventAnswered(deliveryID string, eventName string, event *github.DiscussionEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "answered"
	if action != *event.Action {
		return fmt.Errorf(
			"handleDiscussionEventAnswered() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleDiscussionEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onDiscussionEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onDiscussionEvent[action] {
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

// OnDiscussionEventUnanswered registers callbacks listening to events of type github.DiscussionEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnDiscussionEventUnanswered must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnDiscussionEventUnanswered(callbacks ...DiscussionEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "unanswered"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onDiscussionEvent == nil {
		g.onDiscussionEvent = make(map[string][]DiscussionEventHandleFunc)
	}
	g.onDiscussionEvent[action] = append(g.onDiscussionEvent[action], callbacks...)
}

// SetOnDiscussionEventUnanswered registers callbacks listening to events of type github.DiscussionEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnDiscussionEventUnansweredAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnDiscussionEventUnanswered(callbacks ...DiscussionEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "unanswered"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onDiscussionEvent == nil {
		g.onDiscussionEvent = make(map[string][]DiscussionEventHandleFunc)
	}
	g.onDiscussionEvent[action] = callbacks
}

func (g *EventHandler) handleDiscussionEventUnanswered(deliveryID string, eventName string, event *github.DiscussionEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "unanswered"
	if action != *event.Action {
		return fmt.Errorf(
			"handleDiscussionEventUnanswered() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleDiscussionEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onDiscussionEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onDiscussionEvent[action] {
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

// OnDiscussionEventLabeled registers callbacks listening to events of type github.DiscussionEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnDiscussionEventLabeled must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnDiscussionEventLabeled(callbacks ...DiscussionEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "labeled"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onDiscussionEvent == nil {
		g.onDiscussionEvent = make(map[string][]DiscussionEventHandleFunc)
	}
	g.onDiscussionEvent[action] = append(g.onDiscussionEvent[action], callbacks...)
}

// SetOnDiscussionEventLabeled registers callbacks listening to events of type github.DiscussionEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnDiscussionEventLabeledAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnDiscussionEventLabeled(callbacks ...DiscussionEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "labeled"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onDiscussionEvent == nil {
		g.onDiscussionEvent = make(map[string][]DiscussionEventHandleFunc)
	}
	g.onDiscussionEvent[action] = callbacks
}

func (g *EventHandler) handleDiscussionEventLabeled(deliveryID string, eventName string, event *github.DiscussionEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "labeled"
	if action != *event.Action {
		return fmt.Errorf(
			"handleDiscussionEventLabeled() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleDiscussionEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onDiscussionEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onDiscussionEvent[action] {
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

// OnDiscussionEventUnlabeled registers callbacks listening to events of type github.DiscussionEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnDiscussionEventUnlabeled must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnDiscussionEventUnlabeled(callbacks ...DiscussionEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "unlabeled"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onDiscussionEvent == nil {
		g.onDiscussionEvent = make(map[string][]DiscussionEventHandleFunc)
	}
	g.onDiscussionEvent[action] = append(g.onDiscussionEvent[action], callbacks...)
}

// SetOnDiscussionEventUnlabeled registers callbacks listening to events of type github.DiscussionEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnDiscussionEventUnlabeledAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnDiscussionEventUnlabeled(callbacks ...DiscussionEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "unlabeled"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onDiscussionEvent == nil {
		g.onDiscussionEvent = make(map[string][]DiscussionEventHandleFunc)
	}
	g.onDiscussionEvent[action] = callbacks
}

func (g *EventHandler) handleDiscussionEventUnlabeled(deliveryID string, eventName string, event *github.DiscussionEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "unlabeled"
	if action != *event.Action {
		return fmt.Errorf(
			"handleDiscussionEventUnlabeled() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleDiscussionEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onDiscussionEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onDiscussionEvent[action] {
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

// OnDiscussionEventAny registers callbacks listening to events of type github.DiscussionEvent
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnDiscussionEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnDiscussionEventAny(callbacks ...DiscussionEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const any = "*"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onDiscussionEvent == nil {
		g.onDiscussionEvent = make(map[string][]DiscussionEventHandleFunc)
	}
	g.onDiscussionEvent[any] = append(g.onDiscussionEvent[any], callbacks...)
}

// SetOnDiscussionEventAny registers callbacks listening to events of type github.DiscussionEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnDiscussionEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnDiscussionEventAny(callbacks ...DiscussionEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const any = "*"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onDiscussionEvent == nil {
		g.onDiscussionEvent = make(map[string][]DiscussionEventHandleFunc)
	}
	g.onDiscussionEvent[any] = callbacks
}

func (g *EventHandler) handleDiscussionEventAny(deliveryID string, eventName string, event *github.DiscussionEvent) error {
	if event == nil {
		return fmt.Errorf("event was empty or nil")
	}
	const any = "*"
	if _, ok := g.onDiscussionEvent[any]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onDiscussionEvent[any] {
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

// DiscussionEvent handles github.DiscussionEvent
//
// Callbacks are executed in the following order:
//
// 1) All callbacks registered with OnBeforeAny are executed in parallel.
// 2) All callbacks registered with OnDiscussionEventAny are executed in parallel.
// 3) Optional: All callbacks registered with OnDiscussionEvent... are executed in parallel in case the Event has actions.
// 4) All callbacks registered with OnAfterAny are executed in parallel.
//
// on any error all callbacks registered with OnError are executed in parallel.
func (g *EventHandler) DiscussionEvent(deliveryID string, eventName string, event *github.DiscussionEvent) error {

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
		err := g.handleDiscussionEventCreated(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "edited":
		err := g.handleDiscussionEventEdited(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "deleted":
		err := g.handleDiscussionEventDeleted(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "pinned":
		err := g.handleDiscussionEventPinned(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "unpinned":
		err := g.handleDiscussionEventUnpinned(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "locked":
		err := g.handleDiscussionEventLocked(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "unlocked":
		err := g.handleDiscussionEventUnlocked(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "transferred":
		err := g.handleDiscussionEventTransferred(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "category_changed":
		err := g.handleDiscussionEventCategoryChanged(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "answered":
		err := g.handleDiscussionEventAnswered(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "unanswered":
		err := g.handleDiscussionEventUnanswered(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "labeled":
		err := g.handleDiscussionEventLabeled(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "unlabeled":
		err := g.handleDiscussionEventUnlabeled(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	default:
		err := g.handleDiscussionEventAny(deliveryID, eventName, event)
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
