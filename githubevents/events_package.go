package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"fmt"
	"github.com/google/go-github/v43/github"
	"golang.org/x/sync/errgroup"
)

// PackageEventHandleFunc represents a callback function triggered on github.PackageEvent.
// deliveryID (type: string) is the unique webhook delivery ID.
// eventName (type: string) is the name of the event.
// event (type: *github.PackageEvent) is the webhook payload.
type PackageEventHandleFunc func(deliveryID string, eventName string, event *github.PackageEvent) error

// OnPackageEventPublished registers callbacks listening to events of type github.PackageEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnPackageEventPublished must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnPackageEventPublished(callbacks ...PackageEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "published"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPackageEvent == nil {
		g.onPackageEvent = make(map[string][]PackageEventHandleFunc)
	}
	g.onPackageEvent[action] = append(g.onPackageEvent[action], callbacks...)
}

// SetOnPackageEventPublished registers callbacks listening to events of type github.PackageEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnPackageEventPublishedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnPackageEventPublished(callbacks ...PackageEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "published"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPackageEvent == nil {
		g.onPackageEvent = make(map[string][]PackageEventHandleFunc)
	}
	g.onPackageEvent[action] = callbacks
}

func (g *EventHandler) handlePackageEventPublished(deliveryID string, eventName string, event *github.PackageEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "published"
	if action != *event.Action {
		return fmt.Errorf(
			"handlePackageEventPublished() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handlePackageEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onPackageEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onPackageEvent[action] {
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

// OnPackageEventUpdated registers callbacks listening to events of type github.PackageEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnPackageEventUpdated must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnPackageEventUpdated(callbacks ...PackageEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "updated"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPackageEvent == nil {
		g.onPackageEvent = make(map[string][]PackageEventHandleFunc)
	}
	g.onPackageEvent[action] = append(g.onPackageEvent[action], callbacks...)
}

// SetOnPackageEventUpdated registers callbacks listening to events of type github.PackageEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnPackageEventUpdatedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnPackageEventUpdated(callbacks ...PackageEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "updated"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPackageEvent == nil {
		g.onPackageEvent = make(map[string][]PackageEventHandleFunc)
	}
	g.onPackageEvent[action] = callbacks
}

func (g *EventHandler) handlePackageEventUpdated(deliveryID string, eventName string, event *github.PackageEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "updated"
	if action != *event.Action {
		return fmt.Errorf(
			"handlePackageEventUpdated() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handlePackageEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onPackageEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onPackageEvent[action] {
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

// OnPackageEventAny registers callbacks listening to events of type github.PackageEvent
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnPackageEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnPackageEventAny(callbacks ...PackageEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const any = "*"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPackageEvent == nil {
		g.onPackageEvent = make(map[string][]PackageEventHandleFunc)
	}
	g.onPackageEvent[any] = append(g.onPackageEvent[any], callbacks...)
}

// SetOnPackageEventAny registers callbacks listening to events of type github.PackageEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnPackageEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnPackageEventAny(callbacks ...PackageEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const any = "*"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPackageEvent == nil {
		g.onPackageEvent = make(map[string][]PackageEventHandleFunc)
	}
	g.onPackageEvent[any] = callbacks
}

func (g *EventHandler) handlePackageEventAny(deliveryID string, eventName string, event *github.PackageEvent) error {
	if event == nil {
		return fmt.Errorf("event was empty or nil")
	}
	const any = "*"
	if _, ok := g.onPackageEvent[any]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onPackageEvent[any] {
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

// PackageEvent handles github.PackageEvent
//
// Callbacks are executed in the following order:
//
// 1) All callbacks registered with OnBeforeAny are executed in parallel.
// 2) All callbacks registered with OnPackageEventAny are executed in parallel.
// 3) Optional: All callbacks registered with OnPackageEvent... are executed in parallel in case the Event has actions.
// 4) All callbacks registered with OnAfterAny are executed in parallel.
//
// on any error all callbacks registered with OnError are executed in parallel.
func (g *EventHandler) PackageEvent(deliveryID string, eventName string, event *github.PackageEvent) error {

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
		err := g.handlePackageEventPublished(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "updated":
		err := g.handlePackageEventUpdated(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	default:
		err := g.handlePackageEventAny(deliveryID, eventName, event)
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
