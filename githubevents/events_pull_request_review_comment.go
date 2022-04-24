package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"fmt"
	"github.com/google/go-github/v43/github"
	"golang.org/x/sync/errgroup"
)

// PullRequestReviewCommentEventHandleFunc represents a callback function triggered on github.PullRequestReviewCommentEvent.
// deliveryID (type: string) is the unique webhook delivery ID.
// eventName (type: string) is the name of the event.
// event (type: *github.PullRequestReviewCommentEvent) is the webhook payload.
type PullRequestReviewCommentEventHandleFunc func(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error

// OnPullRequestReviewCommentEventCreated registers callbacks listening to events of type github.PullRequestReviewCommentEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnPullRequestReviewCommentEventCreated must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnPullRequestReviewCommentEventCreated(callbacks ...PullRequestReviewCommentEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "created"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestReviewCommentEvent == nil {
		g.onPullRequestReviewCommentEvent = make(map[string][]PullRequestReviewCommentEventHandleFunc)
	}
	g.onPullRequestReviewCommentEvent[action] = append(g.onPullRequestReviewCommentEvent[action], callbacks...)
}

// SetOnPullRequestReviewCommentEventCreated registers callbacks listening to events of type github.PullRequestReviewCommentEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnPullRequestReviewCommentEventCreatedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnPullRequestReviewCommentEventCreated(callbacks ...PullRequestReviewCommentEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "created"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestReviewCommentEvent == nil {
		g.onPullRequestReviewCommentEvent = make(map[string][]PullRequestReviewCommentEventHandleFunc)
	}
	g.onPullRequestReviewCommentEvent[action] = callbacks
}

func (g *EventHandler) handlePullRequestReviewCommentEventCreated(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "created"
	if action != *event.Action {
		return fmt.Errorf(
			"handlePullRequestReviewCommentEventCreated() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handlePullRequestReviewCommentEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onPullRequestReviewCommentEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onPullRequestReviewCommentEvent[action] {
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

// OnPullRequestReviewCommentEventEdited registers callbacks listening to events of type github.PullRequestReviewCommentEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnPullRequestReviewCommentEventEdited must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnPullRequestReviewCommentEventEdited(callbacks ...PullRequestReviewCommentEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "edited"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestReviewCommentEvent == nil {
		g.onPullRequestReviewCommentEvent = make(map[string][]PullRequestReviewCommentEventHandleFunc)
	}
	g.onPullRequestReviewCommentEvent[action] = append(g.onPullRequestReviewCommentEvent[action], callbacks...)
}

// SetOnPullRequestReviewCommentEventEdited registers callbacks listening to events of type github.PullRequestReviewCommentEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnPullRequestReviewCommentEventEditedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnPullRequestReviewCommentEventEdited(callbacks ...PullRequestReviewCommentEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "edited"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestReviewCommentEvent == nil {
		g.onPullRequestReviewCommentEvent = make(map[string][]PullRequestReviewCommentEventHandleFunc)
	}
	g.onPullRequestReviewCommentEvent[action] = callbacks
}

func (g *EventHandler) handlePullRequestReviewCommentEventEdited(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "edited"
	if action != *event.Action {
		return fmt.Errorf(
			"handlePullRequestReviewCommentEventEdited() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handlePullRequestReviewCommentEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onPullRequestReviewCommentEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onPullRequestReviewCommentEvent[action] {
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

// OnPullRequestReviewCommentEventDeleted registers callbacks listening to events of type github.PullRequestReviewCommentEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnPullRequestReviewCommentEventDeleted must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnPullRequestReviewCommentEventDeleted(callbacks ...PullRequestReviewCommentEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "deleted"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestReviewCommentEvent == nil {
		g.onPullRequestReviewCommentEvent = make(map[string][]PullRequestReviewCommentEventHandleFunc)
	}
	g.onPullRequestReviewCommentEvent[action] = append(g.onPullRequestReviewCommentEvent[action], callbacks...)
}

// SetOnPullRequestReviewCommentEventDeleted registers callbacks listening to events of type github.PullRequestReviewCommentEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnPullRequestReviewCommentEventDeletedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnPullRequestReviewCommentEventDeleted(callbacks ...PullRequestReviewCommentEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "deleted"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestReviewCommentEvent == nil {
		g.onPullRequestReviewCommentEvent = make(map[string][]PullRequestReviewCommentEventHandleFunc)
	}
	g.onPullRequestReviewCommentEvent[action] = callbacks
}

func (g *EventHandler) handlePullRequestReviewCommentEventDeleted(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "deleted"
	if action != *event.Action {
		return fmt.Errorf(
			"handlePullRequestReviewCommentEventDeleted() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handlePullRequestReviewCommentEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onPullRequestReviewCommentEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onPullRequestReviewCommentEvent[action] {
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

// OnPullRequestReviewCommentEventAny registers callbacks listening to events of type github.PullRequestReviewCommentEvent
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnPullRequestReviewCommentEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnPullRequestReviewCommentEventAny(callbacks ...PullRequestReviewCommentEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const any = "*"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestReviewCommentEvent == nil {
		g.onPullRequestReviewCommentEvent = make(map[string][]PullRequestReviewCommentEventHandleFunc)
	}
	g.onPullRequestReviewCommentEvent[any] = append(g.onPullRequestReviewCommentEvent[any], callbacks...)
}

// SetOnPullRequestReviewCommentEventAny registers callbacks listening to events of type github.PullRequestReviewCommentEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnPullRequestReviewCommentEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnPullRequestReviewCommentEventAny(callbacks ...PullRequestReviewCommentEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const any = "*"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onPullRequestReviewCommentEvent == nil {
		g.onPullRequestReviewCommentEvent = make(map[string][]PullRequestReviewCommentEventHandleFunc)
	}
	g.onPullRequestReviewCommentEvent[any] = callbacks
}

func (g *EventHandler) handlePullRequestReviewCommentEventAny(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {
	if event == nil {
		return fmt.Errorf("event was empty or nil")
	}
	const any = "*"
	if _, ok := g.onPullRequestReviewCommentEvent[any]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onPullRequestReviewCommentEvent[any] {
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

// PullRequestReviewCommentEvent handles github.PullRequestReviewCommentEvent
//
// Callbacks are executed in the following order:
//
// 1) All callbacks registered with OnBeforeAny are executed in parallel.
// 2) All callbacks registered with OnPullRequestReviewCommentEventAny are executed in parallel.
// 3) Optional: All callbacks registered with OnPullRequestReviewCommentEvent... are executed in parallel in case the Event has actions.
// 4) All callbacks registered with OnAfterAny are executed in parallel.
//
// on any error all callbacks registered with OnError are executed in parallel.
func (g *EventHandler) PullRequestReviewCommentEvent(deliveryID string, eventName string, event *github.PullRequestReviewCommentEvent) error {

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
		err := g.handlePullRequestReviewCommentEventCreated(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "edited":
		err := g.handlePullRequestReviewCommentEventEdited(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "deleted":
		err := g.handlePullRequestReviewCommentEventDeleted(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	default:
		err := g.handlePullRequestReviewCommentEventAny(deliveryID, eventName, event)
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
