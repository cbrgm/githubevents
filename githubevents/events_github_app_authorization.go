package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"fmt"
	"github.com/google/go-github/v43/github"
	"golang.org/x/sync/errgroup"
)

// GitHubAppAuthorizationEventHandleFunc represents a callback function triggered on github.GitHubAppAuthorizationEvent.
// deliveryID (type: string) is the unique webhook delivery ID.
// eventName (type: string) is the name of the event.
// event (type: *github.GitHubAppAuthorizationEvent) is the webhook payload.
type GitHubAppAuthorizationEventHandleFunc func(deliveryID string, eventName string, event *github.GitHubAppAuthorizationEvent) error

// OnGitHubAppAuthorizationEventRevoked registers callbacks listening to events of type github.GitHubAppAuthorizationEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnGitHubAppAuthorizationEventRevoked must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnGitHubAppAuthorizationEventRevoked(callbacks ...GitHubAppAuthorizationEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "revoked"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onGitHubAppAuthorizationEvent == nil {
		g.onGitHubAppAuthorizationEvent = make(map[string][]GitHubAppAuthorizationEventHandleFunc)
	}
	g.onGitHubAppAuthorizationEvent[action] = append(g.onGitHubAppAuthorizationEvent[action], callbacks...)
}

// SetOnGitHubAppAuthorizationEventRevoked registers callbacks listening to events of type github.GitHubAppAuthorizationEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnGitHubAppAuthorizationEventRevokedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnGitHubAppAuthorizationEventRevoked(callbacks ...GitHubAppAuthorizationEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "revoked"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onGitHubAppAuthorizationEvent == nil {
		g.onGitHubAppAuthorizationEvent = make(map[string][]GitHubAppAuthorizationEventHandleFunc)
	}
	g.onGitHubAppAuthorizationEvent[action] = callbacks
}

func (g *EventHandler) handleGitHubAppAuthorizationEventRevoked(deliveryID string, eventName string, event *github.GitHubAppAuthorizationEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "revoked"
	if action != *event.Action {
		return fmt.Errorf(
			"handleGitHubAppAuthorizationEventRevoked() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleGitHubAppAuthorizationEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onGitHubAppAuthorizationEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onGitHubAppAuthorizationEvent[action] {
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

// OnGitHubAppAuthorizationEventAny registers callbacks listening to events of type github.GitHubAppAuthorizationEvent
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnGitHubAppAuthorizationEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnGitHubAppAuthorizationEventAny(callbacks ...GitHubAppAuthorizationEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const any = "*"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onGitHubAppAuthorizationEvent == nil {
		g.onGitHubAppAuthorizationEvent = make(map[string][]GitHubAppAuthorizationEventHandleFunc)
	}
	g.onGitHubAppAuthorizationEvent[any] = append(g.onGitHubAppAuthorizationEvent[any], callbacks...)
}

// SetOnGitHubAppAuthorizationEventAny registers callbacks listening to events of type github.GitHubAppAuthorizationEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnGitHubAppAuthorizationEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnGitHubAppAuthorizationEventAny(callbacks ...GitHubAppAuthorizationEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const any = "*"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onGitHubAppAuthorizationEvent == nil {
		g.onGitHubAppAuthorizationEvent = make(map[string][]GitHubAppAuthorizationEventHandleFunc)
	}
	g.onGitHubAppAuthorizationEvent[any] = callbacks
}

func (g *EventHandler) handleGitHubAppAuthorizationEventAny(deliveryID string, eventName string, event *github.GitHubAppAuthorizationEvent) error {
	if event == nil {
		return fmt.Errorf("event was empty or nil")
	}
	const any = "*"
	if _, ok := g.onGitHubAppAuthorizationEvent[any]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onGitHubAppAuthorizationEvent[any] {
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

// GitHubAppAuthorizationEvent handles github.GitHubAppAuthorizationEvent
//
// Callbacks are executed in the following order:
//
// 1) All callbacks registered with OnBeforeAny are executed in parallel.
// 2) All callbacks registered with OnGitHubAppAuthorizationEventAny are executed in parallel.
// 3) Optional: All callbacks registered with OnGitHubAppAuthorizationEvent... are executed in parallel in case the Event has actions.
// 4) All callbacks registered with OnAfterAny are executed in parallel.
//
// on any error all callbacks registered with OnError are executed in parallel.
func (g *EventHandler) GitHubAppAuthorizationEvent(deliveryID string, eventName string, event *github.GitHubAppAuthorizationEvent) error {

	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	action := *event.Action

	err := g.handleBeforeAny(deliveryID, eventName, event)
	if err != nil {
		return g.handleError(deliveryID, eventName, event, err)
	}

	switch action {

	case "revoked":
		err := g.handleGitHubAppAuthorizationEventRevoked(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	default:
		err := g.handleGitHubAppAuthorizationEventAny(deliveryID, eventName, event)
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
