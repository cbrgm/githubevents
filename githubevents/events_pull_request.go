package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"fmt"
	"github.com/google/go-github/v43/github"
	"golang.org/x/sync/errgroup"
)

// PullRequestEventHandleFunc represents a callback function triggered on github.PullRequestEvent.
// deliveryID (type: string) is the unique webhook delivery ID.
// eventName (type: string) is the name of the event.
// event (type: *github.PullRequestEvent) is the webhook payload.
type PullRequestEventHandleFunc func(deliveryID string, eventName string, event *github.PullRequestEvent) error

// OnPullRequestEventAssigned registers callbacks listening to events of type github.PullRequestEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnPullRequestEventAssigned must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnPullRequestEventAssigned(callbacks ...PullRequestEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "assigned"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestEvent == nil {
		g.onPullRequestEvent = make(map[string][]PullRequestEventHandleFunc)
	}
	g.onPullRequestEvent[action] = append(g.onPullRequestEvent[action], callbacks...)
}

// SetOnPullRequestEventAssigned registers callbacks listening to events of type github.PullRequestEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnPullRequestEventAssignedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnPullRequestEventAssigned(callbacks ...PullRequestEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "assigned"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestEvent == nil {
		g.onPullRequestEvent = make(map[string][]PullRequestEventHandleFunc)
	}
	g.onPullRequestEvent[action] = callbacks
}

func (g *EventHandler) handlePullRequestEventAssigned(deliveryID string, eventName string, event *github.PullRequestEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "assigned"
	if action != *event.Action {
		return fmt.Errorf(
			"handlePullRequestEventAssigned() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handlePullRequestEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onPullRequestEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onPullRequestEvent[action] {
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

// OnPullRequestEventAutoMergeDisabled registers callbacks listening to events of type github.PullRequestEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnPullRequestEventAutoMergeDisabled must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnPullRequestEventAutoMergeDisabled(callbacks ...PullRequestEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "auto_merge_disabled"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestEvent == nil {
		g.onPullRequestEvent = make(map[string][]PullRequestEventHandleFunc)
	}
	g.onPullRequestEvent[action] = append(g.onPullRequestEvent[action], callbacks...)
}

// SetOnPullRequestEventAutoMergeDisabled registers callbacks listening to events of type github.PullRequestEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnPullRequestEventAutoMergeDisabledAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnPullRequestEventAutoMergeDisabled(callbacks ...PullRequestEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "auto_merge_disabled"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestEvent == nil {
		g.onPullRequestEvent = make(map[string][]PullRequestEventHandleFunc)
	}
	g.onPullRequestEvent[action] = callbacks
}

func (g *EventHandler) handlePullRequestEventAutoMergeDisabled(deliveryID string, eventName string, event *github.PullRequestEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "auto_merge_disabled"
	if action != *event.Action {
		return fmt.Errorf(
			"handlePullRequestEventAutoMergeDisabled() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handlePullRequestEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onPullRequestEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onPullRequestEvent[action] {
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

// OnPullRequestEventAutoMergeEnabled registers callbacks listening to events of type github.PullRequestEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnPullRequestEventAutoMergeEnabled must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnPullRequestEventAutoMergeEnabled(callbacks ...PullRequestEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "auto_merge_enabled"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestEvent == nil {
		g.onPullRequestEvent = make(map[string][]PullRequestEventHandleFunc)
	}
	g.onPullRequestEvent[action] = append(g.onPullRequestEvent[action], callbacks...)
}

// SetOnPullRequestEventAutoMergeEnabled registers callbacks listening to events of type github.PullRequestEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnPullRequestEventAutoMergeEnabledAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnPullRequestEventAutoMergeEnabled(callbacks ...PullRequestEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "auto_merge_enabled"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestEvent == nil {
		g.onPullRequestEvent = make(map[string][]PullRequestEventHandleFunc)
	}
	g.onPullRequestEvent[action] = callbacks
}

func (g *EventHandler) handlePullRequestEventAutoMergeEnabled(deliveryID string, eventName string, event *github.PullRequestEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "auto_merge_enabled"
	if action != *event.Action {
		return fmt.Errorf(
			"handlePullRequestEventAutoMergeEnabled() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handlePullRequestEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onPullRequestEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onPullRequestEvent[action] {
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

// OnPullRequestEventClosed registers callbacks listening to events of type github.PullRequestEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnPullRequestEventClosed must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnPullRequestEventClosed(callbacks ...PullRequestEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "closed"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestEvent == nil {
		g.onPullRequestEvent = make(map[string][]PullRequestEventHandleFunc)
	}
	g.onPullRequestEvent[action] = append(g.onPullRequestEvent[action], callbacks...)
}

// SetOnPullRequestEventClosed registers callbacks listening to events of type github.PullRequestEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnPullRequestEventClosedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnPullRequestEventClosed(callbacks ...PullRequestEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "closed"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestEvent == nil {
		g.onPullRequestEvent = make(map[string][]PullRequestEventHandleFunc)
	}
	g.onPullRequestEvent[action] = callbacks
}

func (g *EventHandler) handlePullRequestEventClosed(deliveryID string, eventName string, event *github.PullRequestEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "closed"
	if action != *event.Action {
		return fmt.Errorf(
			"handlePullRequestEventClosed() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handlePullRequestEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onPullRequestEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onPullRequestEvent[action] {
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

// OnPullRequestEventConvertedToDraft registers callbacks listening to events of type github.PullRequestEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnPullRequestEventConvertedToDraft must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnPullRequestEventConvertedToDraft(callbacks ...PullRequestEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "converted_to_draft"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestEvent == nil {
		g.onPullRequestEvent = make(map[string][]PullRequestEventHandleFunc)
	}
	g.onPullRequestEvent[action] = append(g.onPullRequestEvent[action], callbacks...)
}

// SetOnPullRequestEventConvertedToDraft registers callbacks listening to events of type github.PullRequestEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnPullRequestEventConvertedToDraftAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnPullRequestEventConvertedToDraft(callbacks ...PullRequestEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "converted_to_draft"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestEvent == nil {
		g.onPullRequestEvent = make(map[string][]PullRequestEventHandleFunc)
	}
	g.onPullRequestEvent[action] = callbacks
}

func (g *EventHandler) handlePullRequestEventConvertedToDraft(deliveryID string, eventName string, event *github.PullRequestEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "converted_to_draft"
	if action != *event.Action {
		return fmt.Errorf(
			"handlePullRequestEventConvertedToDraft() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handlePullRequestEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onPullRequestEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onPullRequestEvent[action] {
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

// OnPullRequestEventEdited registers callbacks listening to events of type github.PullRequestEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnPullRequestEventEdited must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnPullRequestEventEdited(callbacks ...PullRequestEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "edited"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestEvent == nil {
		g.onPullRequestEvent = make(map[string][]PullRequestEventHandleFunc)
	}
	g.onPullRequestEvent[action] = append(g.onPullRequestEvent[action], callbacks...)
}

// SetOnPullRequestEventEdited registers callbacks listening to events of type github.PullRequestEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnPullRequestEventEditedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnPullRequestEventEdited(callbacks ...PullRequestEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "edited"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestEvent == nil {
		g.onPullRequestEvent = make(map[string][]PullRequestEventHandleFunc)
	}
	g.onPullRequestEvent[action] = callbacks
}

func (g *EventHandler) handlePullRequestEventEdited(deliveryID string, eventName string, event *github.PullRequestEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "edited"
	if action != *event.Action {
		return fmt.Errorf(
			"handlePullRequestEventEdited() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handlePullRequestEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onPullRequestEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onPullRequestEvent[action] {
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

// OnPullRequestEventLabeled registers callbacks listening to events of type github.PullRequestEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnPullRequestEventLabeled must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnPullRequestEventLabeled(callbacks ...PullRequestEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "labeled"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestEvent == nil {
		g.onPullRequestEvent = make(map[string][]PullRequestEventHandleFunc)
	}
	g.onPullRequestEvent[action] = append(g.onPullRequestEvent[action], callbacks...)
}

// SetOnPullRequestEventLabeled registers callbacks listening to events of type github.PullRequestEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnPullRequestEventLabeledAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnPullRequestEventLabeled(callbacks ...PullRequestEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "labeled"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestEvent == nil {
		g.onPullRequestEvent = make(map[string][]PullRequestEventHandleFunc)
	}
	g.onPullRequestEvent[action] = callbacks
}

func (g *EventHandler) handlePullRequestEventLabeled(deliveryID string, eventName string, event *github.PullRequestEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "labeled"
	if action != *event.Action {
		return fmt.Errorf(
			"handlePullRequestEventLabeled() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handlePullRequestEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onPullRequestEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onPullRequestEvent[action] {
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

// OnPullRequestEventLocked registers callbacks listening to events of type github.PullRequestEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnPullRequestEventLocked must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnPullRequestEventLocked(callbacks ...PullRequestEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "locked"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestEvent == nil {
		g.onPullRequestEvent = make(map[string][]PullRequestEventHandleFunc)
	}
	g.onPullRequestEvent[action] = append(g.onPullRequestEvent[action], callbacks...)
}

// SetOnPullRequestEventLocked registers callbacks listening to events of type github.PullRequestEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnPullRequestEventLockedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnPullRequestEventLocked(callbacks ...PullRequestEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "locked"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestEvent == nil {
		g.onPullRequestEvent = make(map[string][]PullRequestEventHandleFunc)
	}
	g.onPullRequestEvent[action] = callbacks
}

func (g *EventHandler) handlePullRequestEventLocked(deliveryID string, eventName string, event *github.PullRequestEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "locked"
	if action != *event.Action {
		return fmt.Errorf(
			"handlePullRequestEventLocked() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handlePullRequestEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onPullRequestEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onPullRequestEvent[action] {
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

// OnPullRequestEventOpened registers callbacks listening to events of type github.PullRequestEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnPullRequestEventOpened must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnPullRequestEventOpened(callbacks ...PullRequestEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "opened"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestEvent == nil {
		g.onPullRequestEvent = make(map[string][]PullRequestEventHandleFunc)
	}
	g.onPullRequestEvent[action] = append(g.onPullRequestEvent[action], callbacks...)
}

// SetOnPullRequestEventOpened registers callbacks listening to events of type github.PullRequestEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnPullRequestEventOpenedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnPullRequestEventOpened(callbacks ...PullRequestEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "opened"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestEvent == nil {
		g.onPullRequestEvent = make(map[string][]PullRequestEventHandleFunc)
	}
	g.onPullRequestEvent[action] = callbacks
}

func (g *EventHandler) handlePullRequestEventOpened(deliveryID string, eventName string, event *github.PullRequestEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "opened"
	if action != *event.Action {
		return fmt.Errorf(
			"handlePullRequestEventOpened() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handlePullRequestEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onPullRequestEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onPullRequestEvent[action] {
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

// OnPullRequestEventReadyForReview registers callbacks listening to events of type github.PullRequestEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnPullRequestEventReadyForReview must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnPullRequestEventReadyForReview(callbacks ...PullRequestEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "ready_for_review"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestEvent == nil {
		g.onPullRequestEvent = make(map[string][]PullRequestEventHandleFunc)
	}
	g.onPullRequestEvent[action] = append(g.onPullRequestEvent[action], callbacks...)
}

// SetOnPullRequestEventReadyForReview registers callbacks listening to events of type github.PullRequestEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnPullRequestEventReadyForReviewAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnPullRequestEventReadyForReview(callbacks ...PullRequestEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "ready_for_review"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestEvent == nil {
		g.onPullRequestEvent = make(map[string][]PullRequestEventHandleFunc)
	}
	g.onPullRequestEvent[action] = callbacks
}

func (g *EventHandler) handlePullRequestEventReadyForReview(deliveryID string, eventName string, event *github.PullRequestEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "ready_for_review"
	if action != *event.Action {
		return fmt.Errorf(
			"handlePullRequestEventReadyForReview() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handlePullRequestEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onPullRequestEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onPullRequestEvent[action] {
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

// OnPullRequestEventReopened registers callbacks listening to events of type github.PullRequestEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnPullRequestEventReopened must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnPullRequestEventReopened(callbacks ...PullRequestEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "reopened"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestEvent == nil {
		g.onPullRequestEvent = make(map[string][]PullRequestEventHandleFunc)
	}
	g.onPullRequestEvent[action] = append(g.onPullRequestEvent[action], callbacks...)
}

// SetOnPullRequestEventReopened registers callbacks listening to events of type github.PullRequestEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnPullRequestEventReopenedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnPullRequestEventReopened(callbacks ...PullRequestEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "reopened"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestEvent == nil {
		g.onPullRequestEvent = make(map[string][]PullRequestEventHandleFunc)
	}
	g.onPullRequestEvent[action] = callbacks
}

func (g *EventHandler) handlePullRequestEventReopened(deliveryID string, eventName string, event *github.PullRequestEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "reopened"
	if action != *event.Action {
		return fmt.Errorf(
			"handlePullRequestEventReopened() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handlePullRequestEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onPullRequestEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onPullRequestEvent[action] {
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

// OnPullRequestEventReviewRequestRemoved registers callbacks listening to events of type github.PullRequestEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnPullRequestEventReviewRequestRemoved must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnPullRequestEventReviewRequestRemoved(callbacks ...PullRequestEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "review_request_removed"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestEvent == nil {
		g.onPullRequestEvent = make(map[string][]PullRequestEventHandleFunc)
	}
	g.onPullRequestEvent[action] = append(g.onPullRequestEvent[action], callbacks...)
}

// SetOnPullRequestEventReviewRequestRemoved registers callbacks listening to events of type github.PullRequestEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnPullRequestEventReviewRequestRemovedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnPullRequestEventReviewRequestRemoved(callbacks ...PullRequestEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "review_request_removed"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestEvent == nil {
		g.onPullRequestEvent = make(map[string][]PullRequestEventHandleFunc)
	}
	g.onPullRequestEvent[action] = callbacks
}

func (g *EventHandler) handlePullRequestEventReviewRequestRemoved(deliveryID string, eventName string, event *github.PullRequestEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "review_request_removed"
	if action != *event.Action {
		return fmt.Errorf(
			"handlePullRequestEventReviewRequestRemoved() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handlePullRequestEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onPullRequestEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onPullRequestEvent[action] {
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

// OnPullRequestEventReviewRequested registers callbacks listening to events of type github.PullRequestEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnPullRequestEventReviewRequested must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnPullRequestEventReviewRequested(callbacks ...PullRequestEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "review_requested"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestEvent == nil {
		g.onPullRequestEvent = make(map[string][]PullRequestEventHandleFunc)
	}
	g.onPullRequestEvent[action] = append(g.onPullRequestEvent[action], callbacks...)
}

// SetOnPullRequestEventReviewRequested registers callbacks listening to events of type github.PullRequestEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnPullRequestEventReviewRequestedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnPullRequestEventReviewRequested(callbacks ...PullRequestEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "review_requested"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestEvent == nil {
		g.onPullRequestEvent = make(map[string][]PullRequestEventHandleFunc)
	}
	g.onPullRequestEvent[action] = callbacks
}

func (g *EventHandler) handlePullRequestEventReviewRequested(deliveryID string, eventName string, event *github.PullRequestEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "review_requested"
	if action != *event.Action {
		return fmt.Errorf(
			"handlePullRequestEventReviewRequested() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handlePullRequestEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onPullRequestEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onPullRequestEvent[action] {
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

// OnPullRequestEventSynchronize registers callbacks listening to events of type github.PullRequestEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnPullRequestEventSynchronize must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnPullRequestEventSynchronize(callbacks ...PullRequestEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "synchronize"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestEvent == nil {
		g.onPullRequestEvent = make(map[string][]PullRequestEventHandleFunc)
	}
	g.onPullRequestEvent[action] = append(g.onPullRequestEvent[action], callbacks...)
}

// SetOnPullRequestEventSynchronize registers callbacks listening to events of type github.PullRequestEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnPullRequestEventSynchronizeAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnPullRequestEventSynchronize(callbacks ...PullRequestEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "synchronize"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestEvent == nil {
		g.onPullRequestEvent = make(map[string][]PullRequestEventHandleFunc)
	}
	g.onPullRequestEvent[action] = callbacks
}

func (g *EventHandler) handlePullRequestEventSynchronize(deliveryID string, eventName string, event *github.PullRequestEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "synchronize"
	if action != *event.Action {
		return fmt.Errorf(
			"handlePullRequestEventSynchronize() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handlePullRequestEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onPullRequestEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onPullRequestEvent[action] {
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

// OnPullRequestEventUnassigned registers callbacks listening to events of type github.PullRequestEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnPullRequestEventUnassigned must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnPullRequestEventUnassigned(callbacks ...PullRequestEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "unassigned"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestEvent == nil {
		g.onPullRequestEvent = make(map[string][]PullRequestEventHandleFunc)
	}
	g.onPullRequestEvent[action] = append(g.onPullRequestEvent[action], callbacks...)
}

// SetOnPullRequestEventUnassigned registers callbacks listening to events of type github.PullRequestEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnPullRequestEventUnassignedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnPullRequestEventUnassigned(callbacks ...PullRequestEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "unassigned"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestEvent == nil {
		g.onPullRequestEvent = make(map[string][]PullRequestEventHandleFunc)
	}
	g.onPullRequestEvent[action] = callbacks
}

func (g *EventHandler) handlePullRequestEventUnassigned(deliveryID string, eventName string, event *github.PullRequestEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "unassigned"
	if action != *event.Action {
		return fmt.Errorf(
			"handlePullRequestEventUnassigned() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handlePullRequestEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onPullRequestEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onPullRequestEvent[action] {
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

// OnPullRequestEventUnlabeled registers callbacks listening to events of type github.PullRequestEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnPullRequestEventUnlabeled must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnPullRequestEventUnlabeled(callbacks ...PullRequestEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "unlabeled"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestEvent == nil {
		g.onPullRequestEvent = make(map[string][]PullRequestEventHandleFunc)
	}
	g.onPullRequestEvent[action] = append(g.onPullRequestEvent[action], callbacks...)
}

// SetOnPullRequestEventUnlabeled registers callbacks listening to events of type github.PullRequestEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnPullRequestEventUnlabeledAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnPullRequestEventUnlabeled(callbacks ...PullRequestEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "unlabeled"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestEvent == nil {
		g.onPullRequestEvent = make(map[string][]PullRequestEventHandleFunc)
	}
	g.onPullRequestEvent[action] = callbacks
}

func (g *EventHandler) handlePullRequestEventUnlabeled(deliveryID string, eventName string, event *github.PullRequestEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "unlabeled"
	if action != *event.Action {
		return fmt.Errorf(
			"handlePullRequestEventUnlabeled() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handlePullRequestEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onPullRequestEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onPullRequestEvent[action] {
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

// OnPullRequestEventUnlocked registers callbacks listening to events of type github.PullRequestEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnPullRequestEventUnlocked must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnPullRequestEventUnlocked(callbacks ...PullRequestEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "unlocked"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestEvent == nil {
		g.onPullRequestEvent = make(map[string][]PullRequestEventHandleFunc)
	}
	g.onPullRequestEvent[action] = append(g.onPullRequestEvent[action], callbacks...)
}

// SetOnPullRequestEventUnlocked registers callbacks listening to events of type github.PullRequestEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnPullRequestEventUnlockedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnPullRequestEventUnlocked(callbacks ...PullRequestEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "unlocked"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestEvent == nil {
		g.onPullRequestEvent = make(map[string][]PullRequestEventHandleFunc)
	}
	g.onPullRequestEvent[action] = callbacks
}

func (g *EventHandler) handlePullRequestEventUnlocked(deliveryID string, eventName string, event *github.PullRequestEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "unlocked"
	if action != *event.Action {
		return fmt.Errorf(
			"handlePullRequestEventUnlocked() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handlePullRequestEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onPullRequestEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onPullRequestEvent[action] {
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

// OnPullRequestEventAny registers callbacks listening to events of type github.PullRequestEvent
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnPullRequestEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnPullRequestEventAny(callbacks ...PullRequestEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const any = "*"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestEvent == nil {
		g.onPullRequestEvent = make(map[string][]PullRequestEventHandleFunc)
	}
	g.onPullRequestEvent[any] = append(g.onPullRequestEvent[any], callbacks...)
}

// SetOnPullRequestEventAny registers callbacks listening to events of type github.PullRequestEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnPullRequestEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnPullRequestEventAny(callbacks ...PullRequestEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const any = "*"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestEvent == nil {
		g.onPullRequestEvent = make(map[string][]PullRequestEventHandleFunc)
	}
	g.onPullRequestEvent[any] = callbacks
}

func (g *EventHandler) handlePullRequestEventAny(deliveryID string, eventName string, event *github.PullRequestEvent) error {
	if event == nil {
		return fmt.Errorf("event was empty or nil")
	}
	const any = "*"
	if _, ok := g.onPullRequestEvent[any]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onPullRequestEvent[any] {
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

// PullRequestEvent handles github.PullRequestEvent
//
// Callbacks are executed in the following order:
//
// 1) All callbacks registered by OnBeforeAny are executed in parallel.
// 2) All callbacks registered by OnPullRequestEventAny are executed in parallel.
// 3) Optional: All callbacks registered via OnPullRequestEvent... are executed in parallel in case the Event has actions.
// 4) All callbacks registered by OnAfterAny are executed in parallel.
// on any error all callbacks registered by OnError are executed in parallel.
//
func (g *EventHandler) PullRequestEvent(deliveryID string, eventName string, event *github.PullRequestEvent) error {

	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	action := *event.Action

	err := g.handleBeforeAny(deliveryID, eventName, event)
	if err != nil {
		return g.handleError(deliveryID, eventName, event, err)
	}

	switch action {

	case "assigned":
		err := g.handlePullRequestEventAssigned(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "auto_merge_disabled":
		err := g.handlePullRequestEventAutoMergeDisabled(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "auto_merge_enabled":
		err := g.handlePullRequestEventAutoMergeEnabled(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "closed":
		err := g.handlePullRequestEventClosed(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "converted_to_draft":
		err := g.handlePullRequestEventConvertedToDraft(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "edited":
		err := g.handlePullRequestEventEdited(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "labeled":
		err := g.handlePullRequestEventLabeled(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "locked":
		err := g.handlePullRequestEventLocked(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "opened":
		err := g.handlePullRequestEventOpened(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "ready_for_review":
		err := g.handlePullRequestEventReadyForReview(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "reopened":
		err := g.handlePullRequestEventReopened(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "review_request_removed":
		err := g.handlePullRequestEventReviewRequestRemoved(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "review_requested":
		err := g.handlePullRequestEventReviewRequested(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "synchronize":
		err := g.handlePullRequestEventSynchronize(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "unassigned":
		err := g.handlePullRequestEventUnassigned(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "unlabeled":
		err := g.handlePullRequestEventUnlabeled(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "unlocked":
		err := g.handlePullRequestEventUnlocked(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	default:
		err := g.handlePullRequestEventAny(deliveryID, eventName, event)
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
