package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"fmt"
	"github.com/google/go-github/v43/github"
	"golang.org/x/sync/errgroup"
)

// OrganizationEventHandleFunc represents a callback function triggered on github.OrganizationEvent.
// deliveryID (type: string) is the unique webhook delivery ID.
// eventName (type: string) is the name of the event.
// event (type: *github.OrganizationEvent) is the webhook payload.
type OrganizationEventHandleFunc func(deliveryID string, eventName string, event *github.OrganizationEvent) error

// OnOrganizationEventDeleted registers callbacks listening to events of type github.OrganizationEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnOrganizationEventDeleted must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnOrganizationEventDeleted(callbacks ...OrganizationEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "deleted"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onOrganizationEvent == nil {
		g.onOrganizationEvent = make(map[string][]OrganizationEventHandleFunc)
	}
	g.onOrganizationEvent[action] = append(g.onOrganizationEvent[action], callbacks...)
}

// SetOnOrganizationEventDeleted registers callbacks listening to events of type github.OrganizationEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnOrganizationEventDeletedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnOrganizationEventDeleted(callbacks ...OrganizationEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "deleted"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onOrganizationEvent == nil {
		g.onOrganizationEvent = make(map[string][]OrganizationEventHandleFunc)
	}
	g.onOrganizationEvent[action] = callbacks
}

func (g *EventHandler) handleOrganizationEventDeleted(deliveryID string, eventName string, event *github.OrganizationEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "deleted"
	if action != *event.Action {
		return fmt.Errorf(
			"handleOrganizationEventDeleted() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleOrganizationEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onOrganizationEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onOrganizationEvent[action] {
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

// OnOrganizationEventRenamed registers callbacks listening to events of type github.OrganizationEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnOrganizationEventRenamed must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnOrganizationEventRenamed(callbacks ...OrganizationEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "renamed"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onOrganizationEvent == nil {
		g.onOrganizationEvent = make(map[string][]OrganizationEventHandleFunc)
	}
	g.onOrganizationEvent[action] = append(g.onOrganizationEvent[action], callbacks...)
}

// SetOnOrganizationEventRenamed registers callbacks listening to events of type github.OrganizationEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnOrganizationEventRenamedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnOrganizationEventRenamed(callbacks ...OrganizationEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "renamed"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onOrganizationEvent == nil {
		g.onOrganizationEvent = make(map[string][]OrganizationEventHandleFunc)
	}
	g.onOrganizationEvent[action] = callbacks
}

func (g *EventHandler) handleOrganizationEventRenamed(deliveryID string, eventName string, event *github.OrganizationEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "renamed"
	if action != *event.Action {
		return fmt.Errorf(
			"handleOrganizationEventRenamed() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleOrganizationEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onOrganizationEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onOrganizationEvent[action] {
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

// OnOrganizationEventMemberAdded registers callbacks listening to events of type github.OrganizationEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnOrganizationEventMemberAdded must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnOrganizationEventMemberAdded(callbacks ...OrganizationEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "member_added"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onOrganizationEvent == nil {
		g.onOrganizationEvent = make(map[string][]OrganizationEventHandleFunc)
	}
	g.onOrganizationEvent[action] = append(g.onOrganizationEvent[action], callbacks...)
}

// SetOnOrganizationEventMemberAdded registers callbacks listening to events of type github.OrganizationEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnOrganizationEventMemberAddedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnOrganizationEventMemberAdded(callbacks ...OrganizationEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "member_added"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onOrganizationEvent == nil {
		g.onOrganizationEvent = make(map[string][]OrganizationEventHandleFunc)
	}
	g.onOrganizationEvent[action] = callbacks
}

func (g *EventHandler) handleOrganizationEventMemberAdded(deliveryID string, eventName string, event *github.OrganizationEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "member_added"
	if action != *event.Action {
		return fmt.Errorf(
			"handleOrganizationEventMemberAdded() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleOrganizationEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onOrganizationEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onOrganizationEvent[action] {
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

// OnOrganizationEventMemberRemoved registers callbacks listening to events of type github.OrganizationEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnOrganizationEventMemberRemoved must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnOrganizationEventMemberRemoved(callbacks ...OrganizationEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "member_removed"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onOrganizationEvent == nil {
		g.onOrganizationEvent = make(map[string][]OrganizationEventHandleFunc)
	}
	g.onOrganizationEvent[action] = append(g.onOrganizationEvent[action], callbacks...)
}

// SetOnOrganizationEventMemberRemoved registers callbacks listening to events of type github.OrganizationEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnOrganizationEventMemberRemovedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnOrganizationEventMemberRemoved(callbacks ...OrganizationEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "member_removed"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onOrganizationEvent == nil {
		g.onOrganizationEvent = make(map[string][]OrganizationEventHandleFunc)
	}
	g.onOrganizationEvent[action] = callbacks
}

func (g *EventHandler) handleOrganizationEventMemberRemoved(deliveryID string, eventName string, event *github.OrganizationEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "member_removed"
	if action != *event.Action {
		return fmt.Errorf(
			"handleOrganizationEventMemberRemoved() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleOrganizationEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onOrganizationEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onOrganizationEvent[action] {
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

// OnOrganizationEventMemberInvited registers callbacks listening to events of type github.OrganizationEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnOrganizationEventMemberInvited must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnOrganizationEventMemberInvited(callbacks ...OrganizationEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "member_invited"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onOrganizationEvent == nil {
		g.onOrganizationEvent = make(map[string][]OrganizationEventHandleFunc)
	}
	g.onOrganizationEvent[action] = append(g.onOrganizationEvent[action], callbacks...)
}

// SetOnOrganizationEventMemberInvited registers callbacks listening to events of type github.OrganizationEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnOrganizationEventMemberInvitedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnOrganizationEventMemberInvited(callbacks ...OrganizationEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "member_invited"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onOrganizationEvent == nil {
		g.onOrganizationEvent = make(map[string][]OrganizationEventHandleFunc)
	}
	g.onOrganizationEvent[action] = callbacks
}

func (g *EventHandler) handleOrganizationEventMemberInvited(deliveryID string, eventName string, event *github.OrganizationEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "member_invited"
	if action != *event.Action {
		return fmt.Errorf(
			"handleOrganizationEventMemberInvited() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleOrganizationEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onOrganizationEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onOrganizationEvent[action] {
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

// OnOrganizationEventAny registers callbacks listening to events of type github.OrganizationEvent
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnOrganizationEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnOrganizationEventAny(callbacks ...OrganizationEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const any = "*"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onOrganizationEvent == nil {
		g.onOrganizationEvent = make(map[string][]OrganizationEventHandleFunc)
	}
	g.onOrganizationEvent[any] = append(g.onOrganizationEvent[any], callbacks...)
}

// SetOnOrganizationEventAny registers callbacks listening to events of type github.OrganizationEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnOrganizationEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnOrganizationEventAny(callbacks ...OrganizationEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const any = "*"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onOrganizationEvent == nil {
		g.onOrganizationEvent = make(map[string][]OrganizationEventHandleFunc)
	}
	g.onOrganizationEvent[any] = callbacks
}

func (g *EventHandler) handleOrganizationEventAny(deliveryID string, eventName string, event *github.OrganizationEvent) error {
	if event == nil {
		return fmt.Errorf("event was empty or nil")
	}
	const any = "*"
	if _, ok := g.onOrganizationEvent[any]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onOrganizationEvent[any] {
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

// OrganizationEvent handles github.OrganizationEvent
//
// Callbacks are executed in the following order:
//
// 1) All callbacks registered with OnBeforeAny are executed in parallel.
// 2) All callbacks registered with OnOrganizationEventAny are executed in parallel.
// 3) Optional: All callbacks registered with OnOrganizationEvent... are executed in parallel in case the Event has actions.
// 4) All callbacks registered with OnAfterAny are executed in parallel.
//
// on any error all callbacks registered with OnError are executed in parallel.
func (g *EventHandler) OrganizationEvent(deliveryID string, eventName string, event *github.OrganizationEvent) error {

	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	action := *event.Action

	err := g.handleBeforeAny(deliveryID, eventName, event)
	if err != nil {
		return g.handleError(deliveryID, eventName, event, err)
	}

	switch action {

	case "deleted":
		err := g.handleOrganizationEventDeleted(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "renamed":
		err := g.handleOrganizationEventRenamed(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "member_added":
		err := g.handleOrganizationEventMemberAdded(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "member_removed":
		err := g.handleOrganizationEventMemberRemoved(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "member_invited":
		err := g.handleOrganizationEventMemberInvited(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	default:
		err := g.handleOrganizationEventAny(deliveryID, eventName, event)
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
