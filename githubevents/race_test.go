// Copyright 2022 The GithubEvents Authors. All rights reserved.
// Use of this source code is governed by the MIT License
// that can be found in the LICENSE file.

package githubevents

import (
	"context"
	"sync"
	"testing"

	"github.com/google/go-github/v90/github"
)

// Registering a handler while events are being dispatched must not race on the
// internal handler maps. The read paths take an RLock snapshot; without it the
// race detector (and sometimes the runtime) flags a concurrent map read/write.
func TestNoDataRaceRegisterWhileDispatching(t *testing.T) {
	g := New("")
	event := &github.LabelEvent{Action: github.Ptr("created")}

	noop := func(ctx context.Context, deliveryID, eventName string, e *github.LabelEvent) error {
		return nil
	}

	var wg sync.WaitGroup
	wg.Add(2)

	// Writer: keeps registering handlers.
	go func() {
		defer wg.Done()
		for i := 0; i < 200; i++ {
			g.OnLabelEventCreated(noop)
			g.OnLabelEventAny(noop)
		}
	}()

	// Reader: keeps dispatching events.
	go func() {
		defer wg.Done()
		for i := 0; i < 200; i++ {
			_ = g.LabelEvent(context.Background(), "id", "label", event)
		}
	}()

	wg.Wait()
}
