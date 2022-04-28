// Copyright 2022 The GithubEvents Authors. All rights reserved.
// Use of this source code is governed by the MIT License
// that can be found in the LICENSE file.

package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"errors"
	"github.com/google/go-github/v43/github"
	"testing"
)

func TestOnMarketplacePurchaseEventAny(t *testing.T) {
	type args struct {
		callbacks []MarketplacePurchaseEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single MarketplacePurchaseEventHandleFunc",
			args: args{
				[]MarketplacePurchaseEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MarketplacePurchaseEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple MarketplacePurchaseEventHandleFuncs",
			args: args{
				[]MarketplacePurchaseEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MarketplacePurchaseEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.MarketplacePurchaseEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnMarketplacePurchaseEventAny(tt.args.callbacks...)
			if len(g.onMarketplacePurchaseEvent[MarketplacePurchaseEventAnyAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onMarketplacePurchaseEvent[MarketplacePurchaseEventAnyAction]))
			}
		})
	}
}

func TestSetOnMarketplacePurchaseEventAny(t *testing.T) {
	type args struct {
		callbacks []MarketplacePurchaseEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single MarketplacePurchaseEventHandleFunc",
			args: args{
				[]MarketplacePurchaseEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MarketplacePurchaseEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple MarketplacePurchaseEventHandleFuncs",
			args: args{
				[]MarketplacePurchaseEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MarketplacePurchaseEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.MarketplacePurchaseEvent) error {
						return nil
					},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			// add callbacks to be overwritten
			g.SetOnMarketplacePurchaseEventAny(func(deliveryID string, eventName string, event *github.MarketplacePurchaseEvent) error {
				return nil
			})
			g.SetOnMarketplacePurchaseEventAny(tt.args.callbacks...)
			if len(g.onMarketplacePurchaseEvent[MarketplacePurchaseEventAnyAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onMarketplacePurchaseEvent[MarketplacePurchaseEventAnyAction]), tt.want)
			}
		})
	}
}

func TestHandleMarketplacePurchaseEventAny(t *testing.T) {

	action := "*"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.MarketplacePurchaseEvent
		fail       bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "must pass",
			args: args{
				deliveryID: "42",
				eventName:  "marketplace_purchase",

				event: &github.MarketplacePurchaseEvent{Action: &action},

				fail: false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "marketplace_purchase",

				event: &github.MarketplacePurchaseEvent{Action: &action},

				fail: true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "marketplace_purchase",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnMarketplacePurchaseEventAny(func(deliveryID string, eventName string, event *github.MarketplacePurchaseEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleMarketplacePurchaseEventAny(tt.args.deliveryID, tt.args.deliveryID, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("TestHandleMarketplacePurchaseEventAny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnMarketplacePurchaseEventPurchased(t *testing.T) {
	type args struct {
		callbacks []MarketplacePurchaseEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single MarketplacePurchaseEventHandleFunc",
			args: args{
				callbacks: []MarketplacePurchaseEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MarketplacePurchaseEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple MarketplacePurchaseEventHandleFunc",
			args: args{
				callbacks: []MarketplacePurchaseEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MarketplacePurchaseEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.MarketplacePurchaseEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnMarketplacePurchaseEventPurchased(tt.args.callbacks...)
			if len(g.onMarketplacePurchaseEvent[MarketplacePurchaseEventPurchasedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onMarketplacePurchaseEvent[MarketplacePurchaseEventPurchasedAction]))
			}
		})
	}
}

func TestSetOnMarketplacePurchaseEventPurchased(t *testing.T) {
	type args struct {
		callbacks []MarketplacePurchaseEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single MarketplacePurchaseEventHandleFunc",
			args: args{
				[]MarketplacePurchaseEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MarketplacePurchaseEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple MarketplacePurchaseEventHandleFuncs",
			args: args{
				[]MarketplacePurchaseEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MarketplacePurchaseEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.MarketplacePurchaseEvent) error {
						return nil
					},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			// add callbacks to be overwritten
			g.SetOnMarketplacePurchaseEventPurchased(func(deliveryID string, eventName string, event *github.MarketplacePurchaseEvent) error {
				return nil
			})
			g.SetOnMarketplacePurchaseEventPurchased(tt.args.callbacks...)
			if len(g.onMarketplacePurchaseEvent[MarketplacePurchaseEventPurchasedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onMarketplacePurchaseEvent[MarketplacePurchaseEventPurchasedAction]), tt.want)
			}
		})
	}
}

func TestHandleMarketplacePurchaseEventPurchased(t *testing.T) {
	action := MarketplacePurchaseEventPurchasedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.MarketplacePurchaseEvent
		fail       bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "must pass",
			args: args{
				deliveryID: "42",
				eventName:  "marketplace_purchase",
				event:      &github.MarketplacePurchaseEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "marketplace_purchase",
				event:      &github.MarketplacePurchaseEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "marketplace_purchase",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "marketplace_purchase",
				event:      &github.MarketplacePurchaseEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "marketplace_purchase",
				event:      &github.MarketplacePurchaseEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "marketplace_purchase",
				event:      &github.MarketplacePurchaseEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnMarketplacePurchaseEventPurchased(func(deliveryID string, eventName string, event *github.MarketplacePurchaseEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleMarketplacePurchaseEventPurchased(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleMarketplacePurchaseEventPurchased() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnMarketplacePurchaseEventPendingChange(t *testing.T) {
	type args struct {
		callbacks []MarketplacePurchaseEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single MarketplacePurchaseEventHandleFunc",
			args: args{
				callbacks: []MarketplacePurchaseEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MarketplacePurchaseEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple MarketplacePurchaseEventHandleFunc",
			args: args{
				callbacks: []MarketplacePurchaseEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MarketplacePurchaseEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.MarketplacePurchaseEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnMarketplacePurchaseEventPendingChange(tt.args.callbacks...)
			if len(g.onMarketplacePurchaseEvent[MarketplacePurchaseEventPendingChangeAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onMarketplacePurchaseEvent[MarketplacePurchaseEventPendingChangeAction]))
			}
		})
	}
}

func TestSetOnMarketplacePurchaseEventPendingChange(t *testing.T) {
	type args struct {
		callbacks []MarketplacePurchaseEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single MarketplacePurchaseEventHandleFunc",
			args: args{
				[]MarketplacePurchaseEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MarketplacePurchaseEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple MarketplacePurchaseEventHandleFuncs",
			args: args{
				[]MarketplacePurchaseEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MarketplacePurchaseEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.MarketplacePurchaseEvent) error {
						return nil
					},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			// add callbacks to be overwritten
			g.SetOnMarketplacePurchaseEventPendingChange(func(deliveryID string, eventName string, event *github.MarketplacePurchaseEvent) error {
				return nil
			})
			g.SetOnMarketplacePurchaseEventPendingChange(tt.args.callbacks...)
			if len(g.onMarketplacePurchaseEvent[MarketplacePurchaseEventPendingChangeAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onMarketplacePurchaseEvent[MarketplacePurchaseEventPendingChangeAction]), tt.want)
			}
		})
	}
}

func TestHandleMarketplacePurchaseEventPendingChange(t *testing.T) {
	action := MarketplacePurchaseEventPendingChangeAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.MarketplacePurchaseEvent
		fail       bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "must pass",
			args: args{
				deliveryID: "42",
				eventName:  "marketplace_purchase",
				event:      &github.MarketplacePurchaseEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "marketplace_purchase",
				event:      &github.MarketplacePurchaseEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "marketplace_purchase",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "marketplace_purchase",
				event:      &github.MarketplacePurchaseEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "marketplace_purchase",
				event:      &github.MarketplacePurchaseEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "marketplace_purchase",
				event:      &github.MarketplacePurchaseEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnMarketplacePurchaseEventPendingChange(func(deliveryID string, eventName string, event *github.MarketplacePurchaseEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleMarketplacePurchaseEventPendingChange(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleMarketplacePurchaseEventPendingChange() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnMarketplacePurchaseEventPendingChangeCancelled(t *testing.T) {
	type args struct {
		callbacks []MarketplacePurchaseEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single MarketplacePurchaseEventHandleFunc",
			args: args{
				callbacks: []MarketplacePurchaseEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MarketplacePurchaseEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple MarketplacePurchaseEventHandleFunc",
			args: args{
				callbacks: []MarketplacePurchaseEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MarketplacePurchaseEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.MarketplacePurchaseEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnMarketplacePurchaseEventPendingChangeCancelled(tt.args.callbacks...)
			if len(g.onMarketplacePurchaseEvent[MarketplacePurchaseEventPendingChangeCancelledAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onMarketplacePurchaseEvent[MarketplacePurchaseEventPendingChangeCancelledAction]))
			}
		})
	}
}

func TestSetOnMarketplacePurchaseEventPendingChangeCancelled(t *testing.T) {
	type args struct {
		callbacks []MarketplacePurchaseEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single MarketplacePurchaseEventHandleFunc",
			args: args{
				[]MarketplacePurchaseEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MarketplacePurchaseEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple MarketplacePurchaseEventHandleFuncs",
			args: args{
				[]MarketplacePurchaseEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MarketplacePurchaseEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.MarketplacePurchaseEvent) error {
						return nil
					},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			// add callbacks to be overwritten
			g.SetOnMarketplacePurchaseEventPendingChangeCancelled(func(deliveryID string, eventName string, event *github.MarketplacePurchaseEvent) error {
				return nil
			})
			g.SetOnMarketplacePurchaseEventPendingChangeCancelled(tt.args.callbacks...)
			if len(g.onMarketplacePurchaseEvent[MarketplacePurchaseEventPendingChangeCancelledAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onMarketplacePurchaseEvent[MarketplacePurchaseEventPendingChangeCancelledAction]), tt.want)
			}
		})
	}
}

func TestHandleMarketplacePurchaseEventPendingChangeCancelled(t *testing.T) {
	action := MarketplacePurchaseEventPendingChangeCancelledAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.MarketplacePurchaseEvent
		fail       bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "must pass",
			args: args{
				deliveryID: "42",
				eventName:  "marketplace_purchase",
				event:      &github.MarketplacePurchaseEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "marketplace_purchase",
				event:      &github.MarketplacePurchaseEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "marketplace_purchase",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "marketplace_purchase",
				event:      &github.MarketplacePurchaseEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "marketplace_purchase",
				event:      &github.MarketplacePurchaseEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "marketplace_purchase",
				event:      &github.MarketplacePurchaseEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnMarketplacePurchaseEventPendingChangeCancelled(func(deliveryID string, eventName string, event *github.MarketplacePurchaseEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleMarketplacePurchaseEventPendingChangeCancelled(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleMarketplacePurchaseEventPendingChangeCancelled() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnMarketplacePurchaseEventChanged(t *testing.T) {
	type args struct {
		callbacks []MarketplacePurchaseEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single MarketplacePurchaseEventHandleFunc",
			args: args{
				callbacks: []MarketplacePurchaseEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MarketplacePurchaseEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple MarketplacePurchaseEventHandleFunc",
			args: args{
				callbacks: []MarketplacePurchaseEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MarketplacePurchaseEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.MarketplacePurchaseEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnMarketplacePurchaseEventChanged(tt.args.callbacks...)
			if len(g.onMarketplacePurchaseEvent[MarketplacePurchaseEventChangedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onMarketplacePurchaseEvent[MarketplacePurchaseEventChangedAction]))
			}
		})
	}
}

func TestSetOnMarketplacePurchaseEventChanged(t *testing.T) {
	type args struct {
		callbacks []MarketplacePurchaseEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single MarketplacePurchaseEventHandleFunc",
			args: args{
				[]MarketplacePurchaseEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MarketplacePurchaseEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple MarketplacePurchaseEventHandleFuncs",
			args: args{
				[]MarketplacePurchaseEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MarketplacePurchaseEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.MarketplacePurchaseEvent) error {
						return nil
					},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			// add callbacks to be overwritten
			g.SetOnMarketplacePurchaseEventChanged(func(deliveryID string, eventName string, event *github.MarketplacePurchaseEvent) error {
				return nil
			})
			g.SetOnMarketplacePurchaseEventChanged(tt.args.callbacks...)
			if len(g.onMarketplacePurchaseEvent[MarketplacePurchaseEventChangedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onMarketplacePurchaseEvent[MarketplacePurchaseEventChangedAction]), tt.want)
			}
		})
	}
}

func TestHandleMarketplacePurchaseEventChanged(t *testing.T) {
	action := MarketplacePurchaseEventChangedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.MarketplacePurchaseEvent
		fail       bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "must pass",
			args: args{
				deliveryID: "42",
				eventName:  "marketplace_purchase",
				event:      &github.MarketplacePurchaseEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "marketplace_purchase",
				event:      &github.MarketplacePurchaseEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "marketplace_purchase",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "marketplace_purchase",
				event:      &github.MarketplacePurchaseEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "marketplace_purchase",
				event:      &github.MarketplacePurchaseEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "marketplace_purchase",
				event:      &github.MarketplacePurchaseEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnMarketplacePurchaseEventChanged(func(deliveryID string, eventName string, event *github.MarketplacePurchaseEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleMarketplacePurchaseEventChanged(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleMarketplacePurchaseEventChanged() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnMarketplacePurchaseEventCancelled(t *testing.T) {
	type args struct {
		callbacks []MarketplacePurchaseEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single MarketplacePurchaseEventHandleFunc",
			args: args{
				callbacks: []MarketplacePurchaseEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MarketplacePurchaseEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple MarketplacePurchaseEventHandleFunc",
			args: args{
				callbacks: []MarketplacePurchaseEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MarketplacePurchaseEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.MarketplacePurchaseEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnMarketplacePurchaseEventCancelled(tt.args.callbacks...)
			if len(g.onMarketplacePurchaseEvent[MarketplacePurchaseEventCancelledAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onMarketplacePurchaseEvent[MarketplacePurchaseEventCancelledAction]))
			}
		})
	}
}

func TestSetOnMarketplacePurchaseEventCancelled(t *testing.T) {
	type args struct {
		callbacks []MarketplacePurchaseEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single MarketplacePurchaseEventHandleFunc",
			args: args{
				[]MarketplacePurchaseEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MarketplacePurchaseEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple MarketplacePurchaseEventHandleFuncs",
			args: args{
				[]MarketplacePurchaseEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MarketplacePurchaseEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.MarketplacePurchaseEvent) error {
						return nil
					},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			// add callbacks to be overwritten
			g.SetOnMarketplacePurchaseEventCancelled(func(deliveryID string, eventName string, event *github.MarketplacePurchaseEvent) error {
				return nil
			})
			g.SetOnMarketplacePurchaseEventCancelled(tt.args.callbacks...)
			if len(g.onMarketplacePurchaseEvent[MarketplacePurchaseEventCancelledAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onMarketplacePurchaseEvent[MarketplacePurchaseEventCancelledAction]), tt.want)
			}
		})
	}
}

func TestHandleMarketplacePurchaseEventCancelled(t *testing.T) {
	action := MarketplacePurchaseEventCancelledAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.MarketplacePurchaseEvent
		fail       bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "must pass",
			args: args{
				deliveryID: "42",
				eventName:  "marketplace_purchase",
				event:      &github.MarketplacePurchaseEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "marketplace_purchase",
				event:      &github.MarketplacePurchaseEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "marketplace_purchase",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "marketplace_purchase",
				event:      &github.MarketplacePurchaseEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "marketplace_purchase",
				event:      &github.MarketplacePurchaseEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "marketplace_purchase",
				event:      &github.MarketplacePurchaseEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnMarketplacePurchaseEventCancelled(func(deliveryID string, eventName string, event *github.MarketplacePurchaseEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleMarketplacePurchaseEventCancelled(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleMarketplacePurchaseEventCancelled() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
