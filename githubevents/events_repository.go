package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"fmt"
	"github.com/google/go-github/v43/github"
	"golang.org/x/sync/errgroup"
)

// RepositoryEventHandleFunc represents a callback function triggered on github.RepositoryEvent.
// deliveryID (type: string) is the unique webhook delivery ID.
// eventName (type: string) is the name of the event.
// event (type: *github.RepositoryEvent) is the webhook payload.
type RepositoryEventHandleFunc func(deliveryID string, eventName string, event *github.RepositoryEvent) error

// OnRepositoryEvenCreated registers callbacks listening to events of type github.RepositoryEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnRepositoryEvenCreated must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnRepositoryEvenCreated(callbacks ...RepositoryEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "created"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onRepositoryEvent == nil {
		g.onRepositoryEvent = make(map[string][]RepositoryEventHandleFunc)
	}
	g.onRepositoryEvent[action] = append(g.onRepositoryEvent[action], callbacks...)
}

// SetOnRepositoryEvenCreated registers callbacks listening to events of type github.RepositoryEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnRepositoryEvenCreatedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnRepositoryEvenCreated(callbacks ...RepositoryEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "created"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onRepositoryEvent == nil {
		g.onRepositoryEvent = make(map[string][]RepositoryEventHandleFunc)
	}
	g.onRepositoryEvent[action] = callbacks
}

func (g *EventHandler) handleRepositoryEvenCreated(deliveryID string, eventName string, event *github.RepositoryEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "created"
	if action != *event.Action {
		return fmt.Errorf(
			"handleRepositoryEvenCreated() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleRepositoryEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onRepositoryEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onRepositoryEvent[action] {
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

// OnRepositoryEventDeleted registers callbacks listening to events of type github.RepositoryEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnRepositoryEventDeleted must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnRepositoryEventDeleted(callbacks ...RepositoryEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "deleted"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onRepositoryEvent == nil {
		g.onRepositoryEvent = make(map[string][]RepositoryEventHandleFunc)
	}
	g.onRepositoryEvent[action] = append(g.onRepositoryEvent[action], callbacks...)
}

// SetOnRepositoryEventDeleted registers callbacks listening to events of type github.RepositoryEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnRepositoryEventDeletedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnRepositoryEventDeleted(callbacks ...RepositoryEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "deleted"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onRepositoryEvent == nil {
		g.onRepositoryEvent = make(map[string][]RepositoryEventHandleFunc)
	}
	g.onRepositoryEvent[action] = callbacks
}

func (g *EventHandler) handleRepositoryEventDeleted(deliveryID string, eventName string, event *github.RepositoryEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "deleted"
	if action != *event.Action {
		return fmt.Errorf(
			"handleRepositoryEventDeleted() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleRepositoryEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onRepositoryEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onRepositoryEvent[action] {
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

// OnRepositoryEventArchived registers callbacks listening to events of type github.RepositoryEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnRepositoryEventArchived must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnRepositoryEventArchived(callbacks ...RepositoryEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "archived"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onRepositoryEvent == nil {
		g.onRepositoryEvent = make(map[string][]RepositoryEventHandleFunc)
	}
	g.onRepositoryEvent[action] = append(g.onRepositoryEvent[action], callbacks...)
}

// SetOnRepositoryEventArchived registers callbacks listening to events of type github.RepositoryEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnRepositoryEventArchivedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnRepositoryEventArchived(callbacks ...RepositoryEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "archived"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onRepositoryEvent == nil {
		g.onRepositoryEvent = make(map[string][]RepositoryEventHandleFunc)
	}
	g.onRepositoryEvent[action] = callbacks
}

func (g *EventHandler) handleRepositoryEventArchived(deliveryID string, eventName string, event *github.RepositoryEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "archived"
	if action != *event.Action {
		return fmt.Errorf(
			"handleRepositoryEventArchived() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleRepositoryEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onRepositoryEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onRepositoryEvent[action] {
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

// OnRepositoryEventUnarchived registers callbacks listening to events of type github.RepositoryEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnRepositoryEventUnarchived must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnRepositoryEventUnarchived(callbacks ...RepositoryEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "unarchived"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onRepositoryEvent == nil {
		g.onRepositoryEvent = make(map[string][]RepositoryEventHandleFunc)
	}
	g.onRepositoryEvent[action] = append(g.onRepositoryEvent[action], callbacks...)
}

// SetOnRepositoryEventUnarchived registers callbacks listening to events of type github.RepositoryEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnRepositoryEventUnarchivedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnRepositoryEventUnarchived(callbacks ...RepositoryEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "unarchived"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onRepositoryEvent == nil {
		g.onRepositoryEvent = make(map[string][]RepositoryEventHandleFunc)
	}
	g.onRepositoryEvent[action] = callbacks
}

func (g *EventHandler) handleRepositoryEventUnarchived(deliveryID string, eventName string, event *github.RepositoryEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "unarchived"
	if action != *event.Action {
		return fmt.Errorf(
			"handleRepositoryEventUnarchived() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleRepositoryEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onRepositoryEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onRepositoryEvent[action] {
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

// OnRepositoryEventEdited registers callbacks listening to events of type github.RepositoryEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnRepositoryEventEdited must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnRepositoryEventEdited(callbacks ...RepositoryEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "edited"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onRepositoryEvent == nil {
		g.onRepositoryEvent = make(map[string][]RepositoryEventHandleFunc)
	}
	g.onRepositoryEvent[action] = append(g.onRepositoryEvent[action], callbacks...)
}

// SetOnRepositoryEventEdited registers callbacks listening to events of type github.RepositoryEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnRepositoryEventEditedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnRepositoryEventEdited(callbacks ...RepositoryEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "edited"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onRepositoryEvent == nil {
		g.onRepositoryEvent = make(map[string][]RepositoryEventHandleFunc)
	}
	g.onRepositoryEvent[action] = callbacks
}

func (g *EventHandler) handleRepositoryEventEdited(deliveryID string, eventName string, event *github.RepositoryEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "edited"
	if action != *event.Action {
		return fmt.Errorf(
			"handleRepositoryEventEdited() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleRepositoryEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onRepositoryEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onRepositoryEvent[action] {
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

// OnRepositoryEventRenamed registers callbacks listening to events of type github.RepositoryEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnRepositoryEventRenamed must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnRepositoryEventRenamed(callbacks ...RepositoryEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "renamed"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onRepositoryEvent == nil {
		g.onRepositoryEvent = make(map[string][]RepositoryEventHandleFunc)
	}
	g.onRepositoryEvent[action] = append(g.onRepositoryEvent[action], callbacks...)
}

// SetOnRepositoryEventRenamed registers callbacks listening to events of type github.RepositoryEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnRepositoryEventRenamedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnRepositoryEventRenamed(callbacks ...RepositoryEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "renamed"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onRepositoryEvent == nil {
		g.onRepositoryEvent = make(map[string][]RepositoryEventHandleFunc)
	}
	g.onRepositoryEvent[action] = callbacks
}

func (g *EventHandler) handleRepositoryEventRenamed(deliveryID string, eventName string, event *github.RepositoryEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "renamed"
	if action != *event.Action {
		return fmt.Errorf(
			"handleRepositoryEventRenamed() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleRepositoryEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onRepositoryEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onRepositoryEvent[action] {
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

// OnRepositoryEventTransferred registers callbacks listening to events of type github.RepositoryEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnRepositoryEventTransferred must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnRepositoryEventTransferred(callbacks ...RepositoryEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "transferred"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onRepositoryEvent == nil {
		g.onRepositoryEvent = make(map[string][]RepositoryEventHandleFunc)
	}
	g.onRepositoryEvent[action] = append(g.onRepositoryEvent[action], callbacks...)
}

// SetOnRepositoryEventTransferred registers callbacks listening to events of type github.RepositoryEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnRepositoryEventTransferredAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnRepositoryEventTransferred(callbacks ...RepositoryEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "transferred"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onRepositoryEvent == nil {
		g.onRepositoryEvent = make(map[string][]RepositoryEventHandleFunc)
	}
	g.onRepositoryEvent[action] = callbacks
}

func (g *EventHandler) handleRepositoryEventTransferred(deliveryID string, eventName string, event *github.RepositoryEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "transferred"
	if action != *event.Action {
		return fmt.Errorf(
			"handleRepositoryEventTransferred() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleRepositoryEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onRepositoryEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onRepositoryEvent[action] {
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

// OnRepositoryEventPublicized registers callbacks listening to events of type github.RepositoryEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnRepositoryEventPublicized must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnRepositoryEventPublicized(callbacks ...RepositoryEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "publicized"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onRepositoryEvent == nil {
		g.onRepositoryEvent = make(map[string][]RepositoryEventHandleFunc)
	}
	g.onRepositoryEvent[action] = append(g.onRepositoryEvent[action], callbacks...)
}

// SetOnRepositoryEventPublicized registers callbacks listening to events of type github.RepositoryEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnRepositoryEventPublicizedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnRepositoryEventPublicized(callbacks ...RepositoryEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "publicized"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onRepositoryEvent == nil {
		g.onRepositoryEvent = make(map[string][]RepositoryEventHandleFunc)
	}
	g.onRepositoryEvent[action] = callbacks
}

func (g *EventHandler) handleRepositoryEventPublicized(deliveryID string, eventName string, event *github.RepositoryEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "publicized"
	if action != *event.Action {
		return fmt.Errorf(
			"handleRepositoryEventPublicized() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleRepositoryEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onRepositoryEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onRepositoryEvent[action] {
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

// OnRepositoryEventPrivatized registers callbacks listening to events of type github.RepositoryEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnRepositoryEventPrivatized must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnRepositoryEventPrivatized(callbacks ...RepositoryEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "privatized"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onRepositoryEvent == nil {
		g.onRepositoryEvent = make(map[string][]RepositoryEventHandleFunc)
	}
	g.onRepositoryEvent[action] = append(g.onRepositoryEvent[action], callbacks...)
}

// SetOnRepositoryEventPrivatized registers callbacks listening to events of type github.RepositoryEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnRepositoryEventPrivatizedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnRepositoryEventPrivatized(callbacks ...RepositoryEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "privatized"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onRepositoryEvent == nil {
		g.onRepositoryEvent = make(map[string][]RepositoryEventHandleFunc)
	}
	g.onRepositoryEvent[action] = callbacks
}

func (g *EventHandler) handleRepositoryEventPrivatized(deliveryID string, eventName string, event *github.RepositoryEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "privatized"
	if action != *event.Action {
		return fmt.Errorf(
			"handleRepositoryEventPrivatized() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleRepositoryEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onRepositoryEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onRepositoryEvent[action] {
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

// OnRepositoryEventAny registers callbacks listening to events of type github.RepositoryEvent
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnRepositoryEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnRepositoryEventAny(callbacks ...RepositoryEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const any = "*"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onRepositoryEvent == nil {
		g.onRepositoryEvent = make(map[string][]RepositoryEventHandleFunc)
	}
	g.onRepositoryEvent[any] = append(g.onRepositoryEvent[any], callbacks...)
}

// SetOnRepositoryEventAny registers callbacks listening to events of type github.RepositoryEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnRepositoryEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnRepositoryEventAny(callbacks ...RepositoryEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const any = "*"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onRepositoryEvent == nil {
		g.onRepositoryEvent = make(map[string][]RepositoryEventHandleFunc)
	}
	g.onRepositoryEvent[any] = callbacks
}

func (g *EventHandler) handleRepositoryEventAny(deliveryID string, eventName string, event *github.RepositoryEvent) error {
	if event == nil {
		return fmt.Errorf("event was empty or nil")
	}
	const any = "*"
	if _, ok := g.onRepositoryEvent[any]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onRepositoryEvent[any] {
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

// RepositoryEvent handles github.RepositoryEvent
//
// Callbacks are executed in the following order:
//
// 1) All callbacks registered by OnBeforeAny are executed in parallel.
// 2) All callbacks registered by OnRepositoryEventAny are executed in parallel.
// 3) Optional: All callbacks registered via OnRepositoryEvent... are executed in parallel in case the Event has actions.
// 4) All callbacks registered by OnAfterAny are executed in parallel.
// on any error all callbacks registered by OnError are executed in parallel.
//
func (g *EventHandler) RepositoryEvent(deliveryID string, eventName string, event *github.RepositoryEvent) error {

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
		err := g.handleRepositoryEvenCreated(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "deleted":
		err := g.handleRepositoryEventDeleted(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "archived":
		err := g.handleRepositoryEventArchived(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "unarchived":
		err := g.handleRepositoryEventUnarchived(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "edited":
		err := g.handleRepositoryEventEdited(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "renamed":
		err := g.handleRepositoryEventRenamed(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "transferred":
		err := g.handleRepositoryEventTransferred(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "publicized":
		err := g.handleRepositoryEventPublicized(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "privatized":
		err := g.handleRepositoryEventPrivatized(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	default:
		err := g.handleRepositoryEventAny(deliveryID, eventName, event)
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
