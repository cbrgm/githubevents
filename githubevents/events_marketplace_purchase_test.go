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
			if len(g.onMarketplacePurchaseEvent["*"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onMarketplacePurchaseEvent["*"]))
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
			if len(g.onMarketplacePurchaseEvent["*"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onMarketplacePurchaseEvent["*"]), tt.want)
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
			if len(g.onMarketplacePurchaseEvent["purchased"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onMarketplacePurchaseEvent["purchased"]))
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
			if len(g.onMarketplacePurchaseEvent["purchased"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onMarketplacePurchaseEvent["purchased"]), tt.want)
			}
		})
	}
}

func TestHandleMarketplacePurchaseEventPurchased(t *testing.T) {
	action := "purchased"

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
			if len(g.onMarketplacePurchaseEvent["pending_change"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onMarketplacePurchaseEvent["pending_change"]))
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
			if len(g.onMarketplacePurchaseEvent["pending_change"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onMarketplacePurchaseEvent["pending_change"]), tt.want)
			}
		})
	}
}

func TestHandleMarketplacePurchaseEventPendingChange(t *testing.T) {
	action := "pending_change"

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
			if len(g.onMarketplacePurchaseEvent["pending_change_cancelled"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onMarketplacePurchaseEvent["pending_change_cancelled"]))
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
			if len(g.onMarketplacePurchaseEvent["pending_change_cancelled"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onMarketplacePurchaseEvent["pending_change_cancelled"]), tt.want)
			}
		})
	}
}

func TestHandleMarketplacePurchaseEventPendingChangeCancelled(t *testing.T) {
	action := "pending_change_cancelled"

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
			if len(g.onMarketplacePurchaseEvent["changed"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onMarketplacePurchaseEvent["changed"]))
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
			if len(g.onMarketplacePurchaseEvent["changed"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onMarketplacePurchaseEvent["changed"]), tt.want)
			}
		})
	}
}

func TestHandleMarketplacePurchaseEventChanged(t *testing.T) {
	action := "changed"

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
			if len(g.onMarketplacePurchaseEvent["cancelled"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onMarketplacePurchaseEvent["cancelled"]))
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
			if len(g.onMarketplacePurchaseEvent["cancelled"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onMarketplacePurchaseEvent["cancelled"]), tt.want)
			}
		})
	}
}

func TestHandleMarketplacePurchaseEventCancelled(t *testing.T) {
	action := "cancelled"

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
