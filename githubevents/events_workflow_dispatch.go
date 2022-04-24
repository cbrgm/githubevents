package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"fmt"
	"github.com/google/go-github/v43/github"
	"golang.org/x/sync/errgroup"
)

// WorkflowDispatchEventHandleFunc represents a callback function triggered on github.WorkflowDispatchEvent.
// deliveryID (type: string) is the unique webhook delivery ID.
// eventName (type: string) is the name of the event.
// event (type: *github.WorkflowDispatchEvent) is the webhook payload.
type WorkflowDispatchEventHandleFunc func(deliveryID string, eventName string, event *github.WorkflowDispatchEvent) error

// OnWorkflowDispatchEventAny registers callbacks listening to events of type github.WorkflowDispatchEvent
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnWorkflowDispatchEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnWorkflowDispatchEventAny(callbacks ...WorkflowDispatchEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const any = "*"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onWorkflowDispatchEvent == nil {
		g.onWorkflowDispatchEvent = make(map[string][]WorkflowDispatchEventHandleFunc)
	}
	g.onWorkflowDispatchEvent[any] = append(g.onWorkflowDispatchEvent[any], callbacks...)
}

// SetOnWorkflowDispatchEventAny registers callbacks listening to events of type github.WorkflowDispatchEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnWorkflowDispatchEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnWorkflowDispatchEventAny(callbacks ...WorkflowDispatchEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const any = "*"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onWorkflowDispatchEvent == nil {
		g.onWorkflowDispatchEvent = make(map[string][]WorkflowDispatchEventHandleFunc)
	}
	g.onWorkflowDispatchEvent[any] = callbacks
}

func (g *EventHandler) handleWorkflowDispatchEventAny(deliveryID string, eventName string, event *github.WorkflowDispatchEvent) error {
	if event == nil {
		return fmt.Errorf("event was empty or nil")
	}
	const any = "*"
	if _, ok := g.onWorkflowDispatchEvent[any]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onWorkflowDispatchEvent[any] {
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

// WorkflowDispatchEvent handles github.WorkflowDispatchEvent
//
// Callbacks are executed in the following order:
//
// 1) All callbacks registered with OnBeforeAny are executed in parallel.
// 2) All callbacks registered with OnWorkflowDispatchEventAny are executed in parallel.
// 3) Optional: All callbacks registered with OnWorkflowDispatchEvent... are executed in parallel in case the Event has actions.
// 4) All callbacks registered with OnAfterAny are executed in parallel.
//
// on any error all callbacks registered with OnError are executed in parallel.
func (g *EventHandler) WorkflowDispatchEvent(deliveryID string, eventName string, event *github.WorkflowDispatchEvent) error {

	if event == nil {
		return fmt.Errorf("event action was empty or nil")
	}

	err := g.handleBeforeAny(deliveryID, eventName, event)
	if err != nil {
		return g.handleError(deliveryID, eventName, event, err)
	}

	err = g.handleWorkflowDispatchEventAny(deliveryID, eventName, event)
	if err != nil {
		return g.handleError(deliveryID, eventName, event, err)
	}

	err = g.handleAfterAny(deliveryID, eventName, event)
	if err != nil {
		return g.handleError(deliveryID, eventName, event, err)
	}
	return nil
}
