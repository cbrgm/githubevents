package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"fmt"
	"github.com/google/go-github/v43/github"
	"golang.org/x/sync/errgroup"
)

// GollumEventHandleFunc represents a callback function triggered on github.GollumEvent.
// deliveryID (type: string) is the unique webhook delivery ID.
// eventName (type: string) is the name of the event.
// event (type: *github.GollumEvent) is the webhook payload.
type GollumEventHandleFunc func(deliveryID string, eventName string, event *github.GollumEvent) error

// OnGollumEventAny registers callbacks listening to events of type github.GollumEvent
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnGollumEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnGollumEventAny(callbacks ...GollumEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const any = "*"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onGollumEvent == nil {
		g.onGollumEvent = make(map[string][]GollumEventHandleFunc)
	}
	g.onGollumEvent[any] = append(g.onGollumEvent[any], callbacks...)
}

// SetOnGollumEventAny registers callbacks listening to events of type github.GollumEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnGollumEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnGollumEventAny(callbacks ...GollumEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const any = "*"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onGollumEvent == nil {
		g.onGollumEvent = make(map[string][]GollumEventHandleFunc)
	}
	g.onGollumEvent[any] = callbacks
}

func (g *EventHandler) handleGollumEventAny(deliveryID string, eventName string, event *github.GollumEvent) error {
	if event == nil {
		return fmt.Errorf("event was empty or nil")
	}
	const any = "*"
	if _, ok := g.onGollumEvent[any]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onGollumEvent[any] {
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

// GollumEvent handles github.GollumEvent
//
// Callbacks are executed in the following order:
//
// 1) All callbacks registered by OnBeforeAny are executed in parallel.
// 2) All callbacks registered by OnGollumEventAny are executed in parallel.
// 3) Optional: All callbacks registered via OnGollumEvent... are executed in parallel in case the Event has actions.
// 4) All callbacks registered by OnAfterAny are executed in parallel.
// on any error all callbacks registered by OnError are executed in parallel.
//
func (g *EventHandler) GollumEvent(deliveryID string, eventName string, event *github.GollumEvent) error {

	if event == nil {
		return fmt.Errorf("event action was empty or nil")
	}

	err := g.handleBeforeAny(deliveryID, eventName, event)
	if err != nil {
		return g.handleError(deliveryID, eventName, event, err)
	}

	err = g.handleGollumEventAny(deliveryID, eventName, event)
	if err != nil {
		return g.handleError(deliveryID, eventName, event, err)
	}

	err = g.handleAfterAny(deliveryID, eventName, event)
	if err != nil {
		return g.handleError(deliveryID, eventName, event, err)
	}
	return nil
}
