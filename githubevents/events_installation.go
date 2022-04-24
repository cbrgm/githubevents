package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"fmt"
	"github.com/google/go-github/v43/github"
	"golang.org/x/sync/errgroup"
)

// InstallationEventHandleFunc represents a callback function triggered on github.InstallationEvent.
// deliveryID (type: string) is the unique webhook delivery ID.
// eventName (type: string) is the name of the event.
// event (type: *github.InstallationEvent) is the webhook payload.
type InstallationEventHandleFunc func(deliveryID string, eventName string, event *github.InstallationEvent) error

// OnInstallationEventCreated registers callbacks listening to events of type github.InstallationEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnInstallationEventCreated must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnInstallationEventCreated(callbacks ...InstallationEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "created"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onInstallationEvent == nil {
		g.onInstallationEvent = make(map[string][]InstallationEventHandleFunc)
	}
	g.onInstallationEvent[action] = append(g.onInstallationEvent[action], callbacks...)
}

// SetOnInstallationEventCreated registers callbacks listening to events of type github.InstallationEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnInstallationEventCreatedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnInstallationEventCreated(callbacks ...InstallationEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "created"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onInstallationEvent == nil {
		g.onInstallationEvent = make(map[string][]InstallationEventHandleFunc)
	}
	g.onInstallationEvent[action] = callbacks
}

func (g *EventHandler) handleInstallationEventCreated(deliveryID string, eventName string, event *github.InstallationEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "created"
	if action != *event.Action {
		return fmt.Errorf(
			"handleInstallationEventCreated() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleInstallationEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onInstallationEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onInstallationEvent[action] {
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

// OnInstallationEventDeleted registers callbacks listening to events of type github.InstallationEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnInstallationEventDeleted must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnInstallationEventDeleted(callbacks ...InstallationEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "deleted"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onInstallationEvent == nil {
		g.onInstallationEvent = make(map[string][]InstallationEventHandleFunc)
	}
	g.onInstallationEvent[action] = append(g.onInstallationEvent[action], callbacks...)
}

// SetOnInstallationEventDeleted registers callbacks listening to events of type github.InstallationEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnInstallationEventDeletedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnInstallationEventDeleted(callbacks ...InstallationEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "deleted"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onInstallationEvent == nil {
		g.onInstallationEvent = make(map[string][]InstallationEventHandleFunc)
	}
	g.onInstallationEvent[action] = callbacks
}

func (g *EventHandler) handleInstallationEventDeleted(deliveryID string, eventName string, event *github.InstallationEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "deleted"
	if action != *event.Action {
		return fmt.Errorf(
			"handleInstallationEventDeleted() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleInstallationEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onInstallationEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onInstallationEvent[action] {
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

// OnInstallationEventEventSuspend registers callbacks listening to events of type github.InstallationEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnInstallationEventEventSuspend must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnInstallationEventEventSuspend(callbacks ...InstallationEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "suspend"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onInstallationEvent == nil {
		g.onInstallationEvent = make(map[string][]InstallationEventHandleFunc)
	}
	g.onInstallationEvent[action] = append(g.onInstallationEvent[action], callbacks...)
}

// SetOnInstallationEventEventSuspend registers callbacks listening to events of type github.InstallationEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnInstallationEventEventSuspendAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnInstallationEventEventSuspend(callbacks ...InstallationEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "suspend"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onInstallationEvent == nil {
		g.onInstallationEvent = make(map[string][]InstallationEventHandleFunc)
	}
	g.onInstallationEvent[action] = callbacks
}

func (g *EventHandler) handleInstallationEventEventSuspend(deliveryID string, eventName string, event *github.InstallationEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "suspend"
	if action != *event.Action {
		return fmt.Errorf(
			"handleInstallationEventEventSuspend() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleInstallationEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onInstallationEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onInstallationEvent[action] {
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

// OnInstallationEventEventUnsuspend registers callbacks listening to events of type github.InstallationEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnInstallationEventEventUnsuspend must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnInstallationEventEventUnsuspend(callbacks ...InstallationEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "unsuspend"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onInstallationEvent == nil {
		g.onInstallationEvent = make(map[string][]InstallationEventHandleFunc)
	}
	g.onInstallationEvent[action] = append(g.onInstallationEvent[action], callbacks...)
}

// SetOnInstallationEventEventUnsuspend registers callbacks listening to events of type github.InstallationEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnInstallationEventEventUnsuspendAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnInstallationEventEventUnsuspend(callbacks ...InstallationEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "unsuspend"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onInstallationEvent == nil {
		g.onInstallationEvent = make(map[string][]InstallationEventHandleFunc)
	}
	g.onInstallationEvent[action] = callbacks
}

func (g *EventHandler) handleInstallationEventEventUnsuspend(deliveryID string, eventName string, event *github.InstallationEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "unsuspend"
	if action != *event.Action {
		return fmt.Errorf(
			"handleInstallationEventEventUnsuspend() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleInstallationEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onInstallationEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onInstallationEvent[action] {
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

// OnInstallationEventNewPermissionsAccepted registers callbacks listening to events of type github.InstallationEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnInstallationEventNewPermissionsAccepted must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnInstallationEventNewPermissionsAccepted(callbacks ...InstallationEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "new_permissions_accepted"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onInstallationEvent == nil {
		g.onInstallationEvent = make(map[string][]InstallationEventHandleFunc)
	}
	g.onInstallationEvent[action] = append(g.onInstallationEvent[action], callbacks...)
}

// SetOnInstallationEventNewPermissionsAccepted registers callbacks listening to events of type github.InstallationEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnInstallationEventNewPermissionsAcceptedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnInstallationEventNewPermissionsAccepted(callbacks ...InstallationEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "new_permissions_accepted"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onInstallationEvent == nil {
		g.onInstallationEvent = make(map[string][]InstallationEventHandleFunc)
	}
	g.onInstallationEvent[action] = callbacks
}

func (g *EventHandler) handleInstallationEventNewPermissionsAccepted(deliveryID string, eventName string, event *github.InstallationEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "new_permissions_accepted"
	if action != *event.Action {
		return fmt.Errorf(
			"handleInstallationEventNewPermissionsAccepted() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleInstallationEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onInstallationEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onInstallationEvent[action] {
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

// OnInstallationEventAny registers callbacks listening to events of type github.InstallationEvent
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnInstallationEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnInstallationEventAny(callbacks ...InstallationEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const any = "*"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onInstallationEvent == nil {
		g.onInstallationEvent = make(map[string][]InstallationEventHandleFunc)
	}
	g.onInstallationEvent[any] = append(g.onInstallationEvent[any], callbacks...)
}

// SetOnInstallationEventAny registers callbacks listening to events of type github.InstallationEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnInstallationEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnInstallationEventAny(callbacks ...InstallationEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const any = "*"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onInstallationEvent == nil {
		g.onInstallationEvent = make(map[string][]InstallationEventHandleFunc)
	}
	g.onInstallationEvent[any] = callbacks
}

func (g *EventHandler) handleInstallationEventAny(deliveryID string, eventName string, event *github.InstallationEvent) error {
	if event == nil {
		return fmt.Errorf("event was empty or nil")
	}
	const any = "*"
	if _, ok := g.onInstallationEvent[any]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onInstallationEvent[any] {
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

// InstallationEvent handles github.InstallationEvent
//
// Callbacks are executed in the following order:
//
// 1) All callbacks registered by OnBeforeAny are executed in parallel.
// 2) All callbacks registered by OnInstallationEventAny are executed in parallel.
// 3) Optional: All callbacks registered via OnInstallationEvent... are executed in parallel in case the Event has actions.
// 4) All callbacks registered by OnAfterAny are executed in parallel.
// on any error all callbacks registered by OnError are executed in parallel.
//
func (g *EventHandler) InstallationEvent(deliveryID string, eventName string, event *github.InstallationEvent) error {

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
		err := g.handleInstallationEventCreated(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "deleted":
		err := g.handleInstallationEventDeleted(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "suspend":
		err := g.handleInstallationEventEventSuspend(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "unsuspend":
		err := g.handleInstallationEventEventUnsuspend(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "new_permissions_accepted":
		err := g.handleInstallationEventNewPermissionsAccepted(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	default:
		err := g.handleInstallationEventAny(deliveryID, eventName, event)
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
