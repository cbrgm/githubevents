package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"fmt"
	"github.com/google/go-github/v43/github"
	"golang.org/x/sync/errgroup"
)

// PullRequestReviewEventHandleFunc represents a callback function triggered on github.PullRequestReviewEvent.
// deliveryID (type: string) is the unique webhook delivery ID.
// eventName (type: string) is the name of the event.
// event (type: *github.PullRequestReviewEvent) is the webhook payload.
type PullRequestReviewEventHandleFunc func(deliveryID string, eventName string, event *github.PullRequestReviewEvent) error

// OnPullRequestReviewEventSubmitted registers callbacks listening to events of type github.PullRequestReviewEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnPullRequestReviewEventSubmitted must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnPullRequestReviewEventSubmitted(callbacks ...PullRequestReviewEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "submitted"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestReviewEvent == nil {
		g.onPullRequestReviewEvent = make(map[string][]PullRequestReviewEventHandleFunc)
	}
	g.onPullRequestReviewEvent[action] = append(g.onPullRequestReviewEvent[action], callbacks...)
}

// SetOnPullRequestReviewEventSubmitted registers callbacks listening to events of type github.PullRequestReviewEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnPullRequestReviewEventSubmittedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnPullRequestReviewEventSubmitted(callbacks ...PullRequestReviewEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "submitted"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestReviewEvent == nil {
		g.onPullRequestReviewEvent = make(map[string][]PullRequestReviewEventHandleFunc)
	}
	g.onPullRequestReviewEvent[action] = callbacks
}

func (g *EventHandler) handlePullRequestReviewEventSubmitted(deliveryID string, eventName string, event *github.PullRequestReviewEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "submitted"
	if action != *event.Action {
		return fmt.Errorf(
			"handlePullRequestReviewEventSubmitted() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handlePullRequestReviewEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onPullRequestReviewEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onPullRequestReviewEvent[action] {
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

// OnPullRequestReviewEventEdited registers callbacks listening to events of type github.PullRequestReviewEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnPullRequestReviewEventEdited must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnPullRequestReviewEventEdited(callbacks ...PullRequestReviewEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "edited"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestReviewEvent == nil {
		g.onPullRequestReviewEvent = make(map[string][]PullRequestReviewEventHandleFunc)
	}
	g.onPullRequestReviewEvent[action] = append(g.onPullRequestReviewEvent[action], callbacks...)
}

// SetOnPullRequestReviewEventEdited registers callbacks listening to events of type github.PullRequestReviewEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnPullRequestReviewEventEditedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnPullRequestReviewEventEdited(callbacks ...PullRequestReviewEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "edited"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestReviewEvent == nil {
		g.onPullRequestReviewEvent = make(map[string][]PullRequestReviewEventHandleFunc)
	}
	g.onPullRequestReviewEvent[action] = callbacks
}

func (g *EventHandler) handlePullRequestReviewEventEdited(deliveryID string, eventName string, event *github.PullRequestReviewEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "edited"
	if action != *event.Action {
		return fmt.Errorf(
			"handlePullRequestReviewEventEdited() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handlePullRequestReviewEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onPullRequestReviewEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onPullRequestReviewEvent[action] {
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

// OnPullRequestReviewEventDismissed registers callbacks listening to events of type github.PullRequestReviewEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnPullRequestReviewEventDismissed must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnPullRequestReviewEventDismissed(callbacks ...PullRequestReviewEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "dismissed"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestReviewEvent == nil {
		g.onPullRequestReviewEvent = make(map[string][]PullRequestReviewEventHandleFunc)
	}
	g.onPullRequestReviewEvent[action] = append(g.onPullRequestReviewEvent[action], callbacks...)
}

// SetOnPullRequestReviewEventDismissed registers callbacks listening to events of type github.PullRequestReviewEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnPullRequestReviewEventDismissedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnPullRequestReviewEventDismissed(callbacks ...PullRequestReviewEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "dismissed"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestReviewEvent == nil {
		g.onPullRequestReviewEvent = make(map[string][]PullRequestReviewEventHandleFunc)
	}
	g.onPullRequestReviewEvent[action] = callbacks
}

func (g *EventHandler) handlePullRequestReviewEventDismissed(deliveryID string, eventName string, event *github.PullRequestReviewEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "dismissed"
	if action != *event.Action {
		return fmt.Errorf(
			"handlePullRequestReviewEventDismissed() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handlePullRequestReviewEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onPullRequestReviewEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onPullRequestReviewEvent[action] {
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

// OnPullRequestReviewEventAny registers callbacks listening to events of type github.PullRequestReviewEvent
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnPullRequestReviewEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnPullRequestReviewEventAny(callbacks ...PullRequestReviewEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const any = "*"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestReviewEvent == nil {
		g.onPullRequestReviewEvent = make(map[string][]PullRequestReviewEventHandleFunc)
	}
	g.onPullRequestReviewEvent[any] = append(g.onPullRequestReviewEvent[any], callbacks...)
}

// SetOnPullRequestReviewEventAny registers callbacks listening to events of type github.PullRequestReviewEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnPullRequestReviewEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnPullRequestReviewEventAny(callbacks ...PullRequestReviewEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const any = "*"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestReviewEvent == nil {
		g.onPullRequestReviewEvent = make(map[string][]PullRequestReviewEventHandleFunc)
	}
	g.onPullRequestReviewEvent[any] = callbacks
}

func (g *EventHandler) handlePullRequestReviewEventAny(deliveryID string, eventName string, event *github.PullRequestReviewEvent) error {
	if event == nil {
		return fmt.Errorf("event was empty or nil")
	}
	const any = "*"
	if _, ok := g.onPullRequestReviewEvent[any]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onPullRequestReviewEvent[any] {
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

// PullRequestReviewEvent handles github.PullRequestReviewEvent
//
// Callbacks are executed in the following order:
//
// 1) All callbacks registered with OnBeforeAny are executed in parallel.
// 2) All callbacks registered with OnPullRequestReviewEventAny are executed in parallel.
// 3) Optional: All callbacks registered with OnPullRequestReviewEvent... are executed in parallel in case the Event has actions.
// 4) All callbacks registered with OnAfterAny are executed in parallel.
//
// on any error all callbacks registered with OnError are executed in parallel.
func (g *EventHandler) PullRequestReviewEvent(deliveryID string, eventName string, event *github.PullRequestReviewEvent) error {

	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	action := *event.Action

	err := g.handleBeforeAny(deliveryID, eventName, event)
	if err != nil {
		return g.handleError(deliveryID, eventName, event, err)
	}

	switch action {

	case "submitted":
		err := g.handlePullRequestReviewEventSubmitted(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "edited":
		err := g.handlePullRequestReviewEventEdited(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "dismissed":
		err := g.handlePullRequestReviewEventDismissed(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	default:
		err := g.handlePullRequestReviewEventAny(deliveryID, eventName, event)
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
