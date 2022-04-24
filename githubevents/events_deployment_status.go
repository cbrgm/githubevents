package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"fmt"
	"github.com/google/go-github/v43/github"
	"golang.org/x/sync/errgroup"
)

// DeploymentStatusEventHandleFunc represents a callback function triggered on github.DeploymentStatusEvent.
// deliveryID (type: string) is the unique webhook delivery ID.
// eventName (type: string) is the name of the event.
// event (type: *github.DeploymentStatusEvent) is the webhook payload.
type DeploymentStatusEventHandleFunc func(deliveryID string, eventName string, event *github.DeploymentStatusEvent) error

// OnDeploymentStatusEventAny registers callbacks listening to events of type github.DeploymentStatusEvent
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnDeploymentStatusEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnDeploymentStatusEventAny(callbacks ...DeploymentStatusEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const any = "*"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onDeploymentStatusEvent == nil {
		g.onDeploymentStatusEvent = make(map[string][]DeploymentStatusEventHandleFunc)
	}
	g.onDeploymentStatusEvent[any] = append(g.onDeploymentStatusEvent[any], callbacks...)
}

// SetOnDeploymentStatusEventAny registers callbacks listening to events of type github.DeploymentStatusEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnDeploymentStatusEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnDeploymentStatusEventAny(callbacks ...DeploymentStatusEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const any = "*"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onDeploymentStatusEvent == nil {
		g.onDeploymentStatusEvent = make(map[string][]DeploymentStatusEventHandleFunc)
	}
	g.onDeploymentStatusEvent[any] = callbacks
}

func (g *EventHandler) handleDeploymentStatusEventAny(deliveryID string, eventName string, event *github.DeploymentStatusEvent) error {
	if event == nil {
		return fmt.Errorf("event was empty or nil")
	}
	const any = "*"
	if _, ok := g.onDeploymentStatusEvent[any]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onDeploymentStatusEvent[any] {
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

// DeploymentStatusEvent handles github.DeploymentStatusEvent
//
// Callbacks are executed in the following order:
//
// 1) All callbacks registered with OnBeforeAny are executed in parallel.
// 2) All callbacks registered with OnDeploymentStatusEventAny are executed in parallel.
// 3) Optional: All callbacks registered with OnDeploymentStatusEvent... are executed in parallel in case the Event has actions.
// 4) All callbacks registered with OnAfterAny are executed in parallel.
//
// on any error all callbacks registered with OnError are executed in parallel.
func (g *EventHandler) DeploymentStatusEvent(deliveryID string, eventName string, event *github.DeploymentStatusEvent) error {

	if event == nil {
		return fmt.Errorf("event action was empty or nil")
	}

	err := g.handleBeforeAny(deliveryID, eventName, event)
	if err != nil {
		return g.handleError(deliveryID, eventName, event, err)
	}

	err = g.handleDeploymentStatusEventAny(deliveryID, eventName, event)
	if err != nil {
		return g.handleError(deliveryID, eventName, event, err)
	}

	err = g.handleAfterAny(deliveryID, eventName, event)
	if err != nil {
		return g.handleError(deliveryID, eventName, event, err)
	}
	return nil
}
