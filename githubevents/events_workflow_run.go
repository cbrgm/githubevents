package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"fmt"
	"github.com/google/go-github/v43/github"
	"golang.org/x/sync/errgroup"
)

// WorkflowRunEventHandleFunc represents a callback function triggered on github.WorkflowRunEvent.
// deliveryID (type: string) is the unique webhook delivery ID.
// eventName (type: string) is the name of the event.
// event (type: *github.WorkflowRunEvent) is the webhook payload.
type WorkflowRunEventHandleFunc func(deliveryID string, eventName string, event *github.WorkflowRunEvent) error

// OnWorkflowRunEventRequested registers callbacks listening to events of type github.WorkflowRunEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnWorkflowRunEventRequested must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnWorkflowRunEventRequested(callbacks ...WorkflowRunEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "requested"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onWorkflowRunEvent == nil {
		g.onWorkflowRunEvent = make(map[string][]WorkflowRunEventHandleFunc)
	}
	g.onWorkflowRunEvent[action] = append(g.onWorkflowRunEvent[action], callbacks...)
}

// SetOnWorkflowRunEventRequested registers callbacks listening to events of type github.WorkflowRunEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnWorkflowRunEventRequestedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnWorkflowRunEventRequested(callbacks ...WorkflowRunEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "requested"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onWorkflowRunEvent == nil {
		g.onWorkflowRunEvent = make(map[string][]WorkflowRunEventHandleFunc)
	}
	g.onWorkflowRunEvent[action] = callbacks
}

func (g *EventHandler) handleWorkflowRunEventRequested(deliveryID string, eventName string, event *github.WorkflowRunEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "requested"
	if action != *event.Action {
		return fmt.Errorf(
			"handleWorkflowRunEventRequested() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleWorkflowRunEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onWorkflowRunEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onWorkflowRunEvent[action] {
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

// OnWorkflowRunEventCompleted registers callbacks listening to events of type github.WorkflowRunEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnWorkflowRunEventCompleted must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnWorkflowRunEventCompleted(callbacks ...WorkflowRunEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "completed"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onWorkflowRunEvent == nil {
		g.onWorkflowRunEvent = make(map[string][]WorkflowRunEventHandleFunc)
	}
	g.onWorkflowRunEvent[action] = append(g.onWorkflowRunEvent[action], callbacks...)
}

// SetOnWorkflowRunEventCompleted registers callbacks listening to events of type github.WorkflowRunEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnWorkflowRunEventCompletedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnWorkflowRunEventCompleted(callbacks ...WorkflowRunEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "completed"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onWorkflowRunEvent == nil {
		g.onWorkflowRunEvent = make(map[string][]WorkflowRunEventHandleFunc)
	}
	g.onWorkflowRunEvent[action] = callbacks
}

func (g *EventHandler) handleWorkflowRunEventCompleted(deliveryID string, eventName string, event *github.WorkflowRunEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "completed"
	if action != *event.Action {
		return fmt.Errorf(
			"handleWorkflowRunEventCompleted() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleWorkflowRunEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onWorkflowRunEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onWorkflowRunEvent[action] {
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

// OnWorkflowRunEventAny registers callbacks listening to events of type github.WorkflowRunEvent
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnWorkflowRunEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnWorkflowRunEventAny(callbacks ...WorkflowRunEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const any = "*"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onWorkflowRunEvent == nil {
		g.onWorkflowRunEvent = make(map[string][]WorkflowRunEventHandleFunc)
	}
	g.onWorkflowRunEvent[any] = append(g.onWorkflowRunEvent[any], callbacks...)
}

// SetOnWorkflowRunEventAny registers callbacks listening to events of type github.WorkflowRunEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnWorkflowRunEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnWorkflowRunEventAny(callbacks ...WorkflowRunEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const any = "*"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onWorkflowRunEvent == nil {
		g.onWorkflowRunEvent = make(map[string][]WorkflowRunEventHandleFunc)
	}
	g.onWorkflowRunEvent[any] = callbacks
}

func (g *EventHandler) handleWorkflowRunEventAny(deliveryID string, eventName string, event *github.WorkflowRunEvent) error {
	if event == nil {
		return fmt.Errorf("event was empty or nil")
	}
	const any = "*"
	if _, ok := g.onWorkflowRunEvent[any]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onWorkflowRunEvent[any] {
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

// WorkflowRunEvent handles github.WorkflowRunEvent
//
// Callbacks are executed in the following order:
//
// 1) All callbacks registered with OnBeforeAny are executed in parallel.
// 2) All callbacks registered with OnWorkflowRunEventAny are executed in parallel.
// 3) Optional: All callbacks registered with OnWorkflowRunEvent... are executed in parallel in case the Event has actions.
// 4) All callbacks registered with OnAfterAny are executed in parallel.
//
// on any error all callbacks registered with OnError are executed in parallel.
func (g *EventHandler) WorkflowRunEvent(deliveryID string, eventName string, event *github.WorkflowRunEvent) error {

	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	action := *event.Action

	err := g.handleBeforeAny(deliveryID, eventName, event)
	if err != nil {
		return g.handleError(deliveryID, eventName, event, err)
	}

	switch action {

	case "requested":
		err := g.handleWorkflowRunEventRequested(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "completed":
		err := g.handleWorkflowRunEventCompleted(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	default:
		err := g.handleWorkflowRunEventAny(deliveryID, eventName, event)
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
