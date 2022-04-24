package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"fmt"
	"github.com/google/go-github/v43/github"
	"golang.org/x/sync/errgroup"
)

// RepositoryDispatchEventHandleFunc represents a callback function triggered on github.RepositoryDispatchEvent.
// deliveryID (type: string) is the unique webhook delivery ID.
// eventName (type: string) is the name of the event.
// event (type: *github.RepositoryDispatchEvent) is the webhook payload.
type RepositoryDispatchEventHandleFunc func(deliveryID string, eventName string, event *github.RepositoryDispatchEvent) error

// OnRepositoryDispatchEventAny registers callbacks listening to events of type github.RepositoryDispatchEvent
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnRepositoryDispatchEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnRepositoryDispatchEventAny(callbacks ...RepositoryDispatchEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const any = "*"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onRepositoryDispatchEvent == nil {
		g.onRepositoryDispatchEvent = make(map[string][]RepositoryDispatchEventHandleFunc)
	}
	g.onRepositoryDispatchEvent[any] = append(g.onRepositoryDispatchEvent[any], callbacks...)
}

// SetOnRepositoryDispatchEventAny registers callbacks listening to events of type github.RepositoryDispatchEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnRepositoryDispatchEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnRepositoryDispatchEventAny(callbacks ...RepositoryDispatchEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const any = "*"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onRepositoryDispatchEvent == nil {
		g.onRepositoryDispatchEvent = make(map[string][]RepositoryDispatchEventHandleFunc)
	}
	g.onRepositoryDispatchEvent[any] = callbacks
}

func (g *EventHandler) handleRepositoryDispatchEventAny(deliveryID string, eventName string, event *github.RepositoryDispatchEvent) error {
	if event == nil {
		return fmt.Errorf("event was empty or nil")
	}
	const any = "*"
	if _, ok := g.onRepositoryDispatchEvent[any]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onRepositoryDispatchEvent[any] {
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

// RepositoryDispatchEvent handles github.RepositoryDispatchEvent
//
// Callbacks are executed in the following order:
//
// 1) All callbacks registered by OnBeforeAny are executed in parallel.
// 2) All callbacks registered by OnRepositoryDispatchEventAny are executed in parallel.
// 3) Optional: All callbacks registered via OnRepositoryDispatchEvent... are executed in parallel in case the Event has actions.
// 4) All callbacks registered by OnAfterAny are executed in parallel.
// on any error all callbacks registered by OnError are executed in parallel.
//
func (g *EventHandler) RepositoryDispatchEvent(deliveryID string, eventName string, event *github.RepositoryDispatchEvent) error {

	if event == nil {
		return fmt.Errorf("event action was empty or nil")
	}

	err := g.handleBeforeAny(deliveryID, eventName, event)
	if err != nil {
		return g.handleError(deliveryID, eventName, event, err)
	}

	err = g.handleRepositoryDispatchEventAny(deliveryID, eventName, event)
	if err != nil {
		return g.handleError(deliveryID, eventName, event, err)
	}

	err = g.handleAfterAny(deliveryID, eventName, event)
	if err != nil {
		return g.handleError(deliveryID, eventName, event, err)
	}
	return nil
}
