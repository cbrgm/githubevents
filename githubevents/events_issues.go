package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"fmt"
	"github.com/google/go-github/v43/github"
	"golang.org/x/sync/errgroup"
)

// IssuesEventHandleFunc represents a callback function triggered on github.IssuesEvent.
// deliveryID (type: string) is the unique webhook delivery ID.
// eventName (type: string) is the name of the event.
// event (type: *github.IssuesEvent) is the webhook payload.
type IssuesEventHandleFunc func(deliveryID string, eventName string, event *github.IssuesEvent) error

// OnIssuesEventOpened registers callbacks listening to events of type github.IssuesEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnIssuesEventOpened must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnIssuesEventOpened(callbacks ...IssuesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "opened"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssuesEvent == nil {
		g.onIssuesEvent = make(map[string][]IssuesEventHandleFunc)
	}
	g.onIssuesEvent[action] = append(g.onIssuesEvent[action], callbacks...)
}

// SetOnIssuesEventOpened registers callbacks listening to events of type github.IssuesEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnIssuesEventOpenedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnIssuesEventOpened(callbacks ...IssuesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "opened"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssuesEvent == nil {
		g.onIssuesEvent = make(map[string][]IssuesEventHandleFunc)
	}
	g.onIssuesEvent[action] = callbacks
}

func (g *EventHandler) handleIssuesEventOpened(deliveryID string, eventName string, event *github.IssuesEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "opened"
	if action != *event.Action {
		return fmt.Errorf(
			"handleIssuesEventOpened() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleIssuesEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onIssuesEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onIssuesEvent[action] {
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

// OnIssuesEventEdited registers callbacks listening to events of type github.IssuesEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnIssuesEventEdited must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnIssuesEventEdited(callbacks ...IssuesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "edited"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssuesEvent == nil {
		g.onIssuesEvent = make(map[string][]IssuesEventHandleFunc)
	}
	g.onIssuesEvent[action] = append(g.onIssuesEvent[action], callbacks...)
}

// SetOnIssuesEventEdited registers callbacks listening to events of type github.IssuesEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnIssuesEventEditedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnIssuesEventEdited(callbacks ...IssuesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "edited"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssuesEvent == nil {
		g.onIssuesEvent = make(map[string][]IssuesEventHandleFunc)
	}
	g.onIssuesEvent[action] = callbacks
}

func (g *EventHandler) handleIssuesEventEdited(deliveryID string, eventName string, event *github.IssuesEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "edited"
	if action != *event.Action {
		return fmt.Errorf(
			"handleIssuesEventEdited() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleIssuesEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onIssuesEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onIssuesEvent[action] {
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

// OnIssuesEventDeleted registers callbacks listening to events of type github.IssuesEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnIssuesEventDeleted must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnIssuesEventDeleted(callbacks ...IssuesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "deleted"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssuesEvent == nil {
		g.onIssuesEvent = make(map[string][]IssuesEventHandleFunc)
	}
	g.onIssuesEvent[action] = append(g.onIssuesEvent[action], callbacks...)
}

// SetOnIssuesEventDeleted registers callbacks listening to events of type github.IssuesEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnIssuesEventDeletedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnIssuesEventDeleted(callbacks ...IssuesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "deleted"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssuesEvent == nil {
		g.onIssuesEvent = make(map[string][]IssuesEventHandleFunc)
	}
	g.onIssuesEvent[action] = callbacks
}

func (g *EventHandler) handleIssuesEventDeleted(deliveryID string, eventName string, event *github.IssuesEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "deleted"
	if action != *event.Action {
		return fmt.Errorf(
			"handleIssuesEventDeleted() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleIssuesEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onIssuesEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onIssuesEvent[action] {
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

// OnIssuesEventPinned registers callbacks listening to events of type github.IssuesEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnIssuesEventPinned must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnIssuesEventPinned(callbacks ...IssuesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "pinned"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssuesEvent == nil {
		g.onIssuesEvent = make(map[string][]IssuesEventHandleFunc)
	}
	g.onIssuesEvent[action] = append(g.onIssuesEvent[action], callbacks...)
}

// SetOnIssuesEventPinned registers callbacks listening to events of type github.IssuesEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnIssuesEventPinnedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnIssuesEventPinned(callbacks ...IssuesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "pinned"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssuesEvent == nil {
		g.onIssuesEvent = make(map[string][]IssuesEventHandleFunc)
	}
	g.onIssuesEvent[action] = callbacks
}

func (g *EventHandler) handleIssuesEventPinned(deliveryID string, eventName string, event *github.IssuesEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "pinned"
	if action != *event.Action {
		return fmt.Errorf(
			"handleIssuesEventPinned() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleIssuesEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onIssuesEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onIssuesEvent[action] {
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

// OnIssuesEventUnpinned registers callbacks listening to events of type github.IssuesEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnIssuesEventUnpinned must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnIssuesEventUnpinned(callbacks ...IssuesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "unpinned"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssuesEvent == nil {
		g.onIssuesEvent = make(map[string][]IssuesEventHandleFunc)
	}
	g.onIssuesEvent[action] = append(g.onIssuesEvent[action], callbacks...)
}

// SetOnIssuesEventUnpinned registers callbacks listening to events of type github.IssuesEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnIssuesEventUnpinnedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnIssuesEventUnpinned(callbacks ...IssuesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "unpinned"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssuesEvent == nil {
		g.onIssuesEvent = make(map[string][]IssuesEventHandleFunc)
	}
	g.onIssuesEvent[action] = callbacks
}

func (g *EventHandler) handleIssuesEventUnpinned(deliveryID string, eventName string, event *github.IssuesEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "unpinned"
	if action != *event.Action {
		return fmt.Errorf(
			"handleIssuesEventUnpinned() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleIssuesEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onIssuesEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onIssuesEvent[action] {
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

// OnIssuesEventClosed registers callbacks listening to events of type github.IssuesEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnIssuesEventClosed must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnIssuesEventClosed(callbacks ...IssuesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "closed"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssuesEvent == nil {
		g.onIssuesEvent = make(map[string][]IssuesEventHandleFunc)
	}
	g.onIssuesEvent[action] = append(g.onIssuesEvent[action], callbacks...)
}

// SetOnIssuesEventClosed registers callbacks listening to events of type github.IssuesEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnIssuesEventClosedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnIssuesEventClosed(callbacks ...IssuesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "closed"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssuesEvent == nil {
		g.onIssuesEvent = make(map[string][]IssuesEventHandleFunc)
	}
	g.onIssuesEvent[action] = callbacks
}

func (g *EventHandler) handleIssuesEventClosed(deliveryID string, eventName string, event *github.IssuesEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "closed"
	if action != *event.Action {
		return fmt.Errorf(
			"handleIssuesEventClosed() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleIssuesEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onIssuesEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onIssuesEvent[action] {
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

// OnIssuesEventReopened registers callbacks listening to events of type github.IssuesEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnIssuesEventReopened must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnIssuesEventReopened(callbacks ...IssuesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "reopened"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssuesEvent == nil {
		g.onIssuesEvent = make(map[string][]IssuesEventHandleFunc)
	}
	g.onIssuesEvent[action] = append(g.onIssuesEvent[action], callbacks...)
}

// SetOnIssuesEventReopened registers callbacks listening to events of type github.IssuesEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnIssuesEventReopenedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnIssuesEventReopened(callbacks ...IssuesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "reopened"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssuesEvent == nil {
		g.onIssuesEvent = make(map[string][]IssuesEventHandleFunc)
	}
	g.onIssuesEvent[action] = callbacks
}

func (g *EventHandler) handleIssuesEventReopened(deliveryID string, eventName string, event *github.IssuesEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "reopened"
	if action != *event.Action {
		return fmt.Errorf(
			"handleIssuesEventReopened() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleIssuesEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onIssuesEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onIssuesEvent[action] {
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

// OnIssuesEventAssigned registers callbacks listening to events of type github.IssuesEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnIssuesEventAssigned must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnIssuesEventAssigned(callbacks ...IssuesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "assigned"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssuesEvent == nil {
		g.onIssuesEvent = make(map[string][]IssuesEventHandleFunc)
	}
	g.onIssuesEvent[action] = append(g.onIssuesEvent[action], callbacks...)
}

// SetOnIssuesEventAssigned registers callbacks listening to events of type github.IssuesEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnIssuesEventAssignedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnIssuesEventAssigned(callbacks ...IssuesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "assigned"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssuesEvent == nil {
		g.onIssuesEvent = make(map[string][]IssuesEventHandleFunc)
	}
	g.onIssuesEvent[action] = callbacks
}

func (g *EventHandler) handleIssuesEventAssigned(deliveryID string, eventName string, event *github.IssuesEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "assigned"
	if action != *event.Action {
		return fmt.Errorf(
			"handleIssuesEventAssigned() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleIssuesEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onIssuesEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onIssuesEvent[action] {
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

// OnIssuesEventUnassigned registers callbacks listening to events of type github.IssuesEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnIssuesEventUnassigned must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnIssuesEventUnassigned(callbacks ...IssuesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "unassigned"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssuesEvent == nil {
		g.onIssuesEvent = make(map[string][]IssuesEventHandleFunc)
	}
	g.onIssuesEvent[action] = append(g.onIssuesEvent[action], callbacks...)
}

// SetOnIssuesEventUnassigned registers callbacks listening to events of type github.IssuesEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnIssuesEventUnassignedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnIssuesEventUnassigned(callbacks ...IssuesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "unassigned"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssuesEvent == nil {
		g.onIssuesEvent = make(map[string][]IssuesEventHandleFunc)
	}
	g.onIssuesEvent[action] = callbacks
}

func (g *EventHandler) handleIssuesEventUnassigned(deliveryID string, eventName string, event *github.IssuesEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "unassigned"
	if action != *event.Action {
		return fmt.Errorf(
			"handleIssuesEventUnassigned() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleIssuesEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onIssuesEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onIssuesEvent[action] {
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

// OnIssuesEventLabeled registers callbacks listening to events of type github.IssuesEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnIssuesEventLabeled must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnIssuesEventLabeled(callbacks ...IssuesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "labeled"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssuesEvent == nil {
		g.onIssuesEvent = make(map[string][]IssuesEventHandleFunc)
	}
	g.onIssuesEvent[action] = append(g.onIssuesEvent[action], callbacks...)
}

// SetOnIssuesEventLabeled registers callbacks listening to events of type github.IssuesEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnIssuesEventLabeledAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnIssuesEventLabeled(callbacks ...IssuesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "labeled"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssuesEvent == nil {
		g.onIssuesEvent = make(map[string][]IssuesEventHandleFunc)
	}
	g.onIssuesEvent[action] = callbacks
}

func (g *EventHandler) handleIssuesEventLabeled(deliveryID string, eventName string, event *github.IssuesEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "labeled"
	if action != *event.Action {
		return fmt.Errorf(
			"handleIssuesEventLabeled() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleIssuesEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onIssuesEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onIssuesEvent[action] {
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

// OnIssuesEventUnlabeled registers callbacks listening to events of type github.IssuesEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnIssuesEventUnlabeled must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnIssuesEventUnlabeled(callbacks ...IssuesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "unlabeled"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssuesEvent == nil {
		g.onIssuesEvent = make(map[string][]IssuesEventHandleFunc)
	}
	g.onIssuesEvent[action] = append(g.onIssuesEvent[action], callbacks...)
}

// SetOnIssuesEventUnlabeled registers callbacks listening to events of type github.IssuesEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnIssuesEventUnlabeledAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnIssuesEventUnlabeled(callbacks ...IssuesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "unlabeled"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssuesEvent == nil {
		g.onIssuesEvent = make(map[string][]IssuesEventHandleFunc)
	}
	g.onIssuesEvent[action] = callbacks
}

func (g *EventHandler) handleIssuesEventUnlabeled(deliveryID string, eventName string, event *github.IssuesEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "unlabeled"
	if action != *event.Action {
		return fmt.Errorf(
			"handleIssuesEventUnlabeled() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleIssuesEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onIssuesEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onIssuesEvent[action] {
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

// OnIssuesEventLocked registers callbacks listening to events of type github.IssuesEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnIssuesEventLocked must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnIssuesEventLocked(callbacks ...IssuesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "locked"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssuesEvent == nil {
		g.onIssuesEvent = make(map[string][]IssuesEventHandleFunc)
	}
	g.onIssuesEvent[action] = append(g.onIssuesEvent[action], callbacks...)
}

// SetOnIssuesEventLocked registers callbacks listening to events of type github.IssuesEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnIssuesEventLockedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnIssuesEventLocked(callbacks ...IssuesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "locked"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssuesEvent == nil {
		g.onIssuesEvent = make(map[string][]IssuesEventHandleFunc)
	}
	g.onIssuesEvent[action] = callbacks
}

func (g *EventHandler) handleIssuesEventLocked(deliveryID string, eventName string, event *github.IssuesEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "locked"
	if action != *event.Action {
		return fmt.Errorf(
			"handleIssuesEventLocked() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleIssuesEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onIssuesEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onIssuesEvent[action] {
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

// OnIssuesEventUnlocked registers callbacks listening to events of type github.IssuesEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnIssuesEventUnlocked must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnIssuesEventUnlocked(callbacks ...IssuesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "unlocked"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssuesEvent == nil {
		g.onIssuesEvent = make(map[string][]IssuesEventHandleFunc)
	}
	g.onIssuesEvent[action] = append(g.onIssuesEvent[action], callbacks...)
}

// SetOnIssuesEventUnlocked registers callbacks listening to events of type github.IssuesEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnIssuesEventUnlockedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnIssuesEventUnlocked(callbacks ...IssuesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "unlocked"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssuesEvent == nil {
		g.onIssuesEvent = make(map[string][]IssuesEventHandleFunc)
	}
	g.onIssuesEvent[action] = callbacks
}

func (g *EventHandler) handleIssuesEventUnlocked(deliveryID string, eventName string, event *github.IssuesEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "unlocked"
	if action != *event.Action {
		return fmt.Errorf(
			"handleIssuesEventUnlocked() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleIssuesEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onIssuesEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onIssuesEvent[action] {
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

// OnIssuesEventTransferred registers callbacks listening to events of type github.IssuesEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnIssuesEventTransferred must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnIssuesEventTransferred(callbacks ...IssuesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "transferred"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssuesEvent == nil {
		g.onIssuesEvent = make(map[string][]IssuesEventHandleFunc)
	}
	g.onIssuesEvent[action] = append(g.onIssuesEvent[action], callbacks...)
}

// SetOnIssuesEventTransferred registers callbacks listening to events of type github.IssuesEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnIssuesEventTransferredAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnIssuesEventTransferred(callbacks ...IssuesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "transferred"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssuesEvent == nil {
		g.onIssuesEvent = make(map[string][]IssuesEventHandleFunc)
	}
	g.onIssuesEvent[action] = callbacks
}

func (g *EventHandler) handleIssuesEventTransferred(deliveryID string, eventName string, event *github.IssuesEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "transferred"
	if action != *event.Action {
		return fmt.Errorf(
			"handleIssuesEventTransferred() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleIssuesEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onIssuesEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onIssuesEvent[action] {
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

// OnIssuesEventMilestoned registers callbacks listening to events of type github.IssuesEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnIssuesEventMilestoned must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnIssuesEventMilestoned(callbacks ...IssuesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "milestoned"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssuesEvent == nil {
		g.onIssuesEvent = make(map[string][]IssuesEventHandleFunc)
	}
	g.onIssuesEvent[action] = append(g.onIssuesEvent[action], callbacks...)
}

// SetOnIssuesEventMilestoned registers callbacks listening to events of type github.IssuesEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnIssuesEventMilestonedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnIssuesEventMilestoned(callbacks ...IssuesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "milestoned"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssuesEvent == nil {
		g.onIssuesEvent = make(map[string][]IssuesEventHandleFunc)
	}
	g.onIssuesEvent[action] = callbacks
}

func (g *EventHandler) handleIssuesEventMilestoned(deliveryID string, eventName string, event *github.IssuesEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "milestoned"
	if action != *event.Action {
		return fmt.Errorf(
			"handleIssuesEventMilestoned() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleIssuesEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onIssuesEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onIssuesEvent[action] {
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

// OnIssuesEventDeMilestoned registers callbacks listening to events of type github.IssuesEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnIssuesEventDeMilestoned must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnIssuesEventDeMilestoned(callbacks ...IssuesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "demilestoned"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssuesEvent == nil {
		g.onIssuesEvent = make(map[string][]IssuesEventHandleFunc)
	}
	g.onIssuesEvent[action] = append(g.onIssuesEvent[action], callbacks...)
}

// SetOnIssuesEventDeMilestoned registers callbacks listening to events of type github.IssuesEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnIssuesEventDeMilestonedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnIssuesEventDeMilestoned(callbacks ...IssuesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "demilestoned"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssuesEvent == nil {
		g.onIssuesEvent = make(map[string][]IssuesEventHandleFunc)
	}
	g.onIssuesEvent[action] = callbacks
}

func (g *EventHandler) handleIssuesEventDeMilestoned(deliveryID string, eventName string, event *github.IssuesEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "demilestoned"
	if action != *event.Action {
		return fmt.Errorf(
			"handleIssuesEventDeMilestoned() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleIssuesEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onIssuesEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onIssuesEvent[action] {
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

// OnIssuesEventAny registers callbacks listening to events of type github.IssuesEvent
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnIssuesEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnIssuesEventAny(callbacks ...IssuesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const any = "*"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssuesEvent == nil {
		g.onIssuesEvent = make(map[string][]IssuesEventHandleFunc)
	}
	g.onIssuesEvent[any] = append(g.onIssuesEvent[any], callbacks...)
}

// SetOnIssuesEventAny registers callbacks listening to events of type github.IssuesEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnIssuesEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnIssuesEventAny(callbacks ...IssuesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const any = "*"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssuesEvent == nil {
		g.onIssuesEvent = make(map[string][]IssuesEventHandleFunc)
	}
	g.onIssuesEvent[any] = callbacks
}

func (g *EventHandler) handleIssuesEventAny(deliveryID string, eventName string, event *github.IssuesEvent) error {
	if event == nil {
		return fmt.Errorf("event was empty or nil")
	}
	const any = "*"
	if _, ok := g.onIssuesEvent[any]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onIssuesEvent[any] {
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

// IssuesEvent handles github.IssuesEvent
//
// Callbacks are executed in the following order:
//
// 1) All callbacks registered with OnBeforeAny are executed in parallel.
// 2) All callbacks registered with OnIssuesEventAny are executed in parallel.
// 3) Optional: All callbacks registered with OnIssuesEvent... are executed in parallel in case the Event has actions.
// 4) All callbacks registered with OnAfterAny are executed in parallel.
//
// on any error all callbacks registered with OnError are executed in parallel.
func (g *EventHandler) IssuesEvent(deliveryID string, eventName string, event *github.IssuesEvent) error {

	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	action := *event.Action

	err := g.handleBeforeAny(deliveryID, eventName, event)
	if err != nil {
		return g.handleError(deliveryID, eventName, event, err)
	}

	switch action {

	case "opened":
		err := g.handleIssuesEventOpened(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "edited":
		err := g.handleIssuesEventEdited(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "deleted":
		err := g.handleIssuesEventDeleted(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "pinned":
		err := g.handleIssuesEventPinned(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "unpinned":
		err := g.handleIssuesEventUnpinned(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "closed":
		err := g.handleIssuesEventClosed(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "reopened":
		err := g.handleIssuesEventReopened(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "assigned":
		err := g.handleIssuesEventAssigned(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "unassigned":
		err := g.handleIssuesEventUnassigned(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "labeled":
		err := g.handleIssuesEventLabeled(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "unlabeled":
		err := g.handleIssuesEventUnlabeled(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "locked":
		err := g.handleIssuesEventLocked(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "unlocked":
		err := g.handleIssuesEventUnlocked(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "transferred":
		err := g.handleIssuesEventTransferred(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "milestoned":
		err := g.handleIssuesEventMilestoned(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "demilestoned":
		err := g.handleIssuesEventDeMilestoned(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	default:
		err := g.handleIssuesEventAny(deliveryID, eventName, event)
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
