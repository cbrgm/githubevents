package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"fmt"
	"github.com/google/go-github/v43/github"
	"golang.org/x/sync/errgroup"
)

// InstallationRepositoriesEventHandleFunc represents a callback function triggered on github.InstallationRepositoriesEvent.
// deliveryID (type: string) is the unique webhook delivery ID.
// eventName (type: string) is the name of the event.
// event (type: *github.InstallationRepositoriesEvent) is the webhook payload.
type InstallationRepositoriesEventHandleFunc func(deliveryID string, eventName string, event *github.InstallationRepositoriesEvent) error

// OnInstallationRepositoriesEventAdded registers callbacks listening to events of type github.InstallationRepositoriesEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnInstallationRepositoriesEventAdded must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnInstallationRepositoriesEventAdded(callbacks ...InstallationRepositoriesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "added"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onInstallationRepositoriesEvent == nil {
		g.onInstallationRepositoriesEvent = make(map[string][]InstallationRepositoriesEventHandleFunc)
	}
	g.onInstallationRepositoriesEvent[action] = append(g.onInstallationRepositoriesEvent[action], callbacks...)
}

// SetOnInstallationRepositoriesEventAdded registers callbacks listening to events of type github.InstallationRepositoriesEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnInstallationRepositoriesEventAddedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnInstallationRepositoriesEventAdded(callbacks ...InstallationRepositoriesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "added"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onInstallationRepositoriesEvent == nil {
		g.onInstallationRepositoriesEvent = make(map[string][]InstallationRepositoriesEventHandleFunc)
	}
	g.onInstallationRepositoriesEvent[action] = callbacks
}

func (g *EventHandler) handleInstallationRepositoriesEventAdded(deliveryID string, eventName string, event *github.InstallationRepositoriesEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "added"
	if action != *event.Action {
		return fmt.Errorf(
			"handleInstallationRepositoriesEventAdded() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleInstallationRepositoriesEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onInstallationRepositoriesEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onInstallationRepositoriesEvent[action] {
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

// OnInstallationRepositoriesEventRemoved registers callbacks listening to events of type github.InstallationRepositoriesEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnInstallationRepositoriesEventRemoved must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnInstallationRepositoriesEventRemoved(callbacks ...InstallationRepositoriesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "removed"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onInstallationRepositoriesEvent == nil {
		g.onInstallationRepositoriesEvent = make(map[string][]InstallationRepositoriesEventHandleFunc)
	}
	g.onInstallationRepositoriesEvent[action] = append(g.onInstallationRepositoriesEvent[action], callbacks...)
}

// SetOnInstallationRepositoriesEventRemoved registers callbacks listening to events of type github.InstallationRepositoriesEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnInstallationRepositoriesEventRemovedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnInstallationRepositoriesEventRemoved(callbacks ...InstallationRepositoriesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "removed"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onInstallationRepositoriesEvent == nil {
		g.onInstallationRepositoriesEvent = make(map[string][]InstallationRepositoriesEventHandleFunc)
	}
	g.onInstallationRepositoriesEvent[action] = callbacks
}

func (g *EventHandler) handleInstallationRepositoriesEventRemoved(deliveryID string, eventName string, event *github.InstallationRepositoriesEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "removed"
	if action != *event.Action {
		return fmt.Errorf(
			"handleInstallationRepositoriesEventRemoved() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleInstallationRepositoriesEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onInstallationRepositoriesEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onInstallationRepositoriesEvent[action] {
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

// OnInstallationRepositoriesEventAny registers callbacks listening to events of type github.InstallationRepositoriesEvent
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnInstallationRepositoriesEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnInstallationRepositoriesEventAny(callbacks ...InstallationRepositoriesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const any = "*"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onInstallationRepositoriesEvent == nil {
		g.onInstallationRepositoriesEvent = make(map[string][]InstallationRepositoriesEventHandleFunc)
	}
	g.onInstallationRepositoriesEvent[any] = append(g.onInstallationRepositoriesEvent[any], callbacks...)
}

// SetOnInstallationRepositoriesEventAny registers callbacks listening to events of type github.InstallationRepositoriesEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnInstallationRepositoriesEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnInstallationRepositoriesEventAny(callbacks ...InstallationRepositoriesEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const any = "*"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onInstallationRepositoriesEvent == nil {
		g.onInstallationRepositoriesEvent = make(map[string][]InstallationRepositoriesEventHandleFunc)
	}
	g.onInstallationRepositoriesEvent[any] = callbacks
}

func (g *EventHandler) handleInstallationRepositoriesEventAny(deliveryID string, eventName string, event *github.InstallationRepositoriesEvent) error {
	if event == nil {
		return fmt.Errorf("event was empty or nil")
	}
	const any = "*"
	if _, ok := g.onInstallationRepositoriesEvent[any]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onInstallationRepositoriesEvent[any] {
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

// InstallationRepositoriesEvent handles github.InstallationRepositoriesEvent
//
// Callbacks are executed in the following order:
//
// 1) All callbacks registered with OnBeforeAny are executed in parallel.
// 2) All callbacks registered with OnInstallationRepositoriesEventAny are executed in parallel.
// 3) Optional: All callbacks registered with OnInstallationRepositoriesEvent... are executed in parallel in case the Event has actions.
// 4) All callbacks registered with OnAfterAny are executed in parallel.
//
// on any error all callbacks registered with OnError are executed in parallel.
func (g *EventHandler) InstallationRepositoriesEvent(deliveryID string, eventName string, event *github.InstallationRepositoriesEvent) error {

	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	action := *event.Action

	err := g.handleBeforeAny(deliveryID, eventName, event)
	if err != nil {
		return g.handleError(deliveryID, eventName, event, err)
	}

	switch action {

	case "added":
		err := g.handleInstallationRepositoriesEventAdded(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "removed":
		err := g.handleInstallationRepositoriesEventRemoved(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	default:
		err := g.handleInstallationRepositoriesEventAny(deliveryID, eventName, event)
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
