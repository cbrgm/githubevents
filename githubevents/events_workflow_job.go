package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"fmt"
	"github.com/google/go-github/v43/github"
	"golang.org/x/sync/errgroup"
)

// WorkflowJobEventHandleFunc represents a callback function triggered on github.WorkflowJobEvent.
// deliveryID (type: string) is the unique webhook delivery ID.
// eventName (type: string) is the name of the event.
// event (type: *github.WorkflowJobEvent) is the webhook payload.
type WorkflowJobEventHandleFunc func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error

// OnWorkflowJobEventQueued registers callbacks listening to events of type github.WorkflowJobEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnWorkflowJobEventQueued must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnWorkflowJobEventQueued(callbacks ...WorkflowJobEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "queued"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onWorkflowJobEvent == nil {
		g.onWorkflowJobEvent = make(map[string][]WorkflowJobEventHandleFunc)
	}
	g.onWorkflowJobEvent[action] = append(g.onWorkflowJobEvent[action], callbacks...)
}

// SetOnWorkflowJobEventQueued registers callbacks listening to events of type github.WorkflowJobEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnWorkflowJobEventQueuedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnWorkflowJobEventQueued(callbacks ...WorkflowJobEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "queued"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onWorkflowJobEvent == nil {
		g.onWorkflowJobEvent = make(map[string][]WorkflowJobEventHandleFunc)
	}
	g.onWorkflowJobEvent[action] = callbacks
}

func (g *EventHandler) handleWorkflowJobEventQueued(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "queued"
	if action != *event.Action {
		return fmt.Errorf(
			"handleWorkflowJobEventQueued() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleWorkflowJobEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onWorkflowJobEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onWorkflowJobEvent[action] {
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

// OnWorkflowJobEventInProgress registers callbacks listening to events of type github.WorkflowJobEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnWorkflowJobEventInProgress must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnWorkflowJobEventInProgress(callbacks ...WorkflowJobEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "in_progress"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onWorkflowJobEvent == nil {
		g.onWorkflowJobEvent = make(map[string][]WorkflowJobEventHandleFunc)
	}
	g.onWorkflowJobEvent[action] = append(g.onWorkflowJobEvent[action], callbacks...)
}

// SetOnWorkflowJobEventInProgress registers callbacks listening to events of type github.WorkflowJobEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnWorkflowJobEventInProgressAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnWorkflowJobEventInProgress(callbacks ...WorkflowJobEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "in_progress"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onWorkflowJobEvent == nil {
		g.onWorkflowJobEvent = make(map[string][]WorkflowJobEventHandleFunc)
	}
	g.onWorkflowJobEvent[action] = callbacks
}

func (g *EventHandler) handleWorkflowJobEventInProgress(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "in_progress"
	if action != *event.Action {
		return fmt.Errorf(
			"handleWorkflowJobEventInProgress() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleWorkflowJobEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onWorkflowJobEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onWorkflowJobEvent[action] {
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

// OnWorkflowJobEventCompleted registers callbacks listening to events of type github.WorkflowJobEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnWorkflowJobEventCompleted must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnWorkflowJobEventCompleted(callbacks ...WorkflowJobEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "completed"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onWorkflowJobEvent == nil {
		g.onWorkflowJobEvent = make(map[string][]WorkflowJobEventHandleFunc)
	}
	g.onWorkflowJobEvent[action] = append(g.onWorkflowJobEvent[action], callbacks...)
}

// SetOnWorkflowJobEventCompleted registers callbacks listening to events of type github.WorkflowJobEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnWorkflowJobEventCompletedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnWorkflowJobEventCompleted(callbacks ...WorkflowJobEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "completed"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onWorkflowJobEvent == nil {
		g.onWorkflowJobEvent = make(map[string][]WorkflowJobEventHandleFunc)
	}
	g.onWorkflowJobEvent[action] = callbacks
}

func (g *EventHandler) handleWorkflowJobEventCompleted(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "completed"
	if action != *event.Action {
		return fmt.Errorf(
			"handleWorkflowJobEventCompleted() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleWorkflowJobEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onWorkflowJobEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onWorkflowJobEvent[action] {
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

// OnWorkflowJobEventAny registers callbacks listening to events of type github.WorkflowJobEvent
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnWorkflowJobEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnWorkflowJobEventAny(callbacks ...WorkflowJobEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const any = "*"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onWorkflowJobEvent == nil {
		g.onWorkflowJobEvent = make(map[string][]WorkflowJobEventHandleFunc)
	}
	g.onWorkflowJobEvent[any] = append(g.onWorkflowJobEvent[any], callbacks...)
}

// SetOnWorkflowJobEventAny registers callbacks listening to events of type github.WorkflowJobEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnWorkflowJobEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnWorkflowJobEventAny(callbacks ...WorkflowJobEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const any = "*"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onWorkflowJobEvent == nil {
		g.onWorkflowJobEvent = make(map[string][]WorkflowJobEventHandleFunc)
	}
	g.onWorkflowJobEvent[any] = callbacks
}

func (g *EventHandler) handleWorkflowJobEventAny(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
	if event == nil {
		return fmt.Errorf("event was empty or nil")
	}
	const any = "*"
	if _, ok := g.onWorkflowJobEvent[any]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onWorkflowJobEvent[any] {
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

// WorkflowJobEvent handles github.WorkflowJobEvent
//
// Callbacks are executed in the following order:
//
// 1) All callbacks registered with OnBeforeAny are executed in parallel.
// 2) All callbacks registered with OnWorkflowJobEventAny are executed in parallel.
// 3) Optional: All callbacks registered with OnWorkflowJobEvent... are executed in parallel in case the Event has actions.
// 4) All callbacks registered with OnAfterAny are executed in parallel.
//
// on any error all callbacks registered with OnError are executed in parallel.
func (g *EventHandler) WorkflowJobEvent(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {

	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	action := *event.Action

	err := g.handleBeforeAny(deliveryID, eventName, event)
	if err != nil {
		return g.handleError(deliveryID, eventName, event, err)
	}

	switch action {

	case "queued":
		err := g.handleWorkflowJobEventQueued(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "in_progress":
		err := g.handleWorkflowJobEventInProgress(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "completed":
		err := g.handleWorkflowJobEventCompleted(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	default:
		err := g.handleWorkflowJobEventAny(deliveryID, eventName, event)
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
