// Copyright 2022 The GithubEvents Authors. All rights reserved.
// Use of this source code is governed by the MIT License
// that can be found in the LICENSE file.

package githubevents

import (
	"context"
	"fmt"

	"golang.org/x/sync/errgroup"
)

// dispatch runs every callback across the given groups in parallel, recovers
// from panics as errors, and returns the first non-nil error. It is the single
// implementation of the concurrency semantics shared by all generated handlers.
func dispatch[T any, F ~func(ctx context.Context, deliveryID, eventName string, event T) error](
	ctx context.Context,
	deliveryID string,
	eventName string,
	event T,
	groups ...[]F,
) error {
	eg := new(errgroup.Group)
	for _, group := range groups {
		for _, h := range group {
			eg.Go(func() (err error) {
				defer func() {
					if r := recover(); r != nil {
						err = fmt.Errorf("recovered from panic: %v", r)
					}
				}()
				return h(ctx, deliveryID, eventName, event)
			})
		}
	}
	return eg.Wait()
}
