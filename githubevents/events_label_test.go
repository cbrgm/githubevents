package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"errors"
	"github.com/google/go-github/v43/github"
	"testing"
)

func TestOnLabelEventAny(t *testing.T) {
	type args struct {
		callbacks []LabelEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single LabelEventHandleFunc",
			args: args{
				[]LabelEventHandleFunc{
					func(deliveryID string, eventName string, event *github.LabelEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple LabelEventHandleFuncs",
			args: args{
				[]LabelEventHandleFunc{
					func(deliveryID string, eventName string, event *github.LabelEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.LabelEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnLabelEventAny(tt.args.callbacks...)
			if len(g.onLabelEvent["*"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onLabelEvent["*"]))
			}
		})
	}
}

func TestSetOnLabelEventAny(t *testing.T) {
	type args struct {
		callbacks []LabelEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single LabelEventHandleFunc",
			args: args{
				[]LabelEventHandleFunc{
					func(deliveryID string, eventName string, event *github.LabelEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple LabelEventHandleFuncs",
			args: args{
				[]LabelEventHandleFunc{
					func(deliveryID string, eventName string, event *github.LabelEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.LabelEvent) error {
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
			g.SetOnLabelEventAny(func(deliveryID string, eventName string, event *github.LabelEvent) error {
				return nil
			})
			g.SetOnLabelEventAny(tt.args.callbacks...)
			if len(g.onLabelEvent["*"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onLabelEvent["*"]), tt.want)
			}
		})
	}
}

func TestHandleLabelEventAny(t *testing.T) {

	action := "*"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.LabelEvent
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
				eventName:  "label",

				event: &github.LabelEvent{Action: &action},

				fail: false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "label",

				event: &github.LabelEvent{Action: &action},

				fail: true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "label",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnLabelEventAny(func(deliveryID string, eventName string, event *github.LabelEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleLabelEventAny(tt.args.deliveryID, tt.args.deliveryID, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("TestHandleLabelEventAny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnLabelEventCreated(t *testing.T) {
	type args struct {
		callbacks []LabelEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single LabelEventHandleFunc",
			args: args{
				callbacks: []LabelEventHandleFunc{
					func(deliveryID string, eventName string, event *github.LabelEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple LabelEventHandleFunc",
			args: args{
				callbacks: []LabelEventHandleFunc{
					func(deliveryID string, eventName string, event *github.LabelEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.LabelEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnLabelEventCreated(tt.args.callbacks...)
			if len(g.onLabelEvent["created"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onLabelEvent["created"]))
			}
		})
	}
}

func TestSetOnLabelEventCreated(t *testing.T) {
	type args struct {
		callbacks []LabelEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single LabelEventHandleFunc",
			args: args{
				[]LabelEventHandleFunc{
					func(deliveryID string, eventName string, event *github.LabelEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple LabelEventHandleFuncs",
			args: args{
				[]LabelEventHandleFunc{
					func(deliveryID string, eventName string, event *github.LabelEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.LabelEvent) error {
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
			g.SetOnLabelEventCreated(func(deliveryID string, eventName string, event *github.LabelEvent) error {
				return nil
			})
			g.SetOnLabelEventCreated(tt.args.callbacks...)
			if len(g.onLabelEvent["created"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onLabelEvent["created"]), tt.want)
			}
		})
	}
}

func TestHandleLabelEventCreated(t *testing.T) {
	action := "created"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.LabelEvent
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
				eventName:  "label",
				event:      &github.LabelEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "label",
				event:      &github.LabelEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "label",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "label",
				event:      &github.LabelEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "label",
				event:      &github.LabelEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "label",
				event:      &github.LabelEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnLabelEventCreated(func(deliveryID string, eventName string, event *github.LabelEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleLabelEventCreated(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleLabelEventCreated() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnLabelEventEdited(t *testing.T) {
	type args struct {
		callbacks []LabelEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single LabelEventHandleFunc",
			args: args{
				callbacks: []LabelEventHandleFunc{
					func(deliveryID string, eventName string, event *github.LabelEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple LabelEventHandleFunc",
			args: args{
				callbacks: []LabelEventHandleFunc{
					func(deliveryID string, eventName string, event *github.LabelEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.LabelEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnLabelEventEdited(tt.args.callbacks...)
			if len(g.onLabelEvent["edited"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onLabelEvent["edited"]))
			}
		})
	}
}

func TestSetOnLabelEventEdited(t *testing.T) {
	type args struct {
		callbacks []LabelEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single LabelEventHandleFunc",
			args: args{
				[]LabelEventHandleFunc{
					func(deliveryID string, eventName string, event *github.LabelEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple LabelEventHandleFuncs",
			args: args{
				[]LabelEventHandleFunc{
					func(deliveryID string, eventName string, event *github.LabelEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.LabelEvent) error {
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
			g.SetOnLabelEventEdited(func(deliveryID string, eventName string, event *github.LabelEvent) error {
				return nil
			})
			g.SetOnLabelEventEdited(tt.args.callbacks...)
			if len(g.onLabelEvent["edited"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onLabelEvent["edited"]), tt.want)
			}
		})
	}
}

func TestHandleLabelEventEdited(t *testing.T) {
	action := "edited"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.LabelEvent
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
				eventName:  "label",
				event:      &github.LabelEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "label",
				event:      &github.LabelEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "label",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "label",
				event:      &github.LabelEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "label",
				event:      &github.LabelEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "label",
				event:      &github.LabelEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnLabelEventEdited(func(deliveryID string, eventName string, event *github.LabelEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleLabelEventEdited(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleLabelEventEdited() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnLabelEventDeleted(t *testing.T) {
	type args struct {
		callbacks []LabelEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single LabelEventHandleFunc",
			args: args{
				callbacks: []LabelEventHandleFunc{
					func(deliveryID string, eventName string, event *github.LabelEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple LabelEventHandleFunc",
			args: args{
				callbacks: []LabelEventHandleFunc{
					func(deliveryID string, eventName string, event *github.LabelEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.LabelEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnLabelEventDeleted(tt.args.callbacks...)
			if len(g.onLabelEvent["deleted"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onLabelEvent["deleted"]))
			}
		})
	}
}

func TestSetOnLabelEventDeleted(t *testing.T) {
	type args struct {
		callbacks []LabelEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single LabelEventHandleFunc",
			args: args{
				[]LabelEventHandleFunc{
					func(deliveryID string, eventName string, event *github.LabelEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple LabelEventHandleFuncs",
			args: args{
				[]LabelEventHandleFunc{
					func(deliveryID string, eventName string, event *github.LabelEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.LabelEvent) error {
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
			g.SetOnLabelEventDeleted(func(deliveryID string, eventName string, event *github.LabelEvent) error {
				return nil
			})
			g.SetOnLabelEventDeleted(tt.args.callbacks...)
			if len(g.onLabelEvent["deleted"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onLabelEvent["deleted"]), tt.want)
			}
		})
	}
}

func TestHandleLabelEventDeleted(t *testing.T) {
	action := "deleted"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.LabelEvent
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
				eventName:  "label",
				event:      &github.LabelEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "label",
				event:      &github.LabelEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "label",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "label",
				event:      &github.LabelEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "label",
				event:      &github.LabelEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "label",
				event:      &github.LabelEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnLabelEventDeleted(func(deliveryID string, eventName string, event *github.LabelEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleLabelEventDeleted(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleLabelEventDeleted() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
