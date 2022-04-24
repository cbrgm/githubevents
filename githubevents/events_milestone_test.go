package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"errors"
	"github.com/google/go-github/v43/github"
	"testing"
)

func TestOnMilestoneEventAny(t *testing.T) {
	type args struct {
		callbacks []MilestoneEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single MilestoneEventHandleFunc",
			args: args{
				[]MilestoneEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MilestoneEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple MilestoneEventHandleFuncs",
			args: args{
				[]MilestoneEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MilestoneEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.MilestoneEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnMilestoneEventAny(tt.args.callbacks...)
			if len(g.onMilestoneEvent["*"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onMilestoneEvent["*"]))
			}
		})
	}
}

func TestSetOnMilestoneEventAny(t *testing.T) {
	type args struct {
		callbacks []MilestoneEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single MilestoneEventHandleFunc",
			args: args{
				[]MilestoneEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MilestoneEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple MilestoneEventHandleFuncs",
			args: args{
				[]MilestoneEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MilestoneEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.MilestoneEvent) error {
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
			g.SetOnMilestoneEventAny(func(deliveryID string, eventName string, event *github.MilestoneEvent) error {
				return nil
			})
			g.SetOnMilestoneEventAny(tt.args.callbacks...)
			if len(g.onMilestoneEvent["*"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onMilestoneEvent["*"]), tt.want)
			}
		})
	}
}

func TestHandleMilestoneEventAny(t *testing.T) {

	action := "*"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.MilestoneEvent
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
				eventName:  "milestone",

				event: &github.MilestoneEvent{Action: &action},

				fail: false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "milestone",

				event: &github.MilestoneEvent{Action: &action},

				fail: true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "milestone",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnMilestoneEventAny(func(deliveryID string, eventName string, event *github.MilestoneEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleMilestoneEventAny(tt.args.deliveryID, tt.args.deliveryID, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("TestHandleMilestoneEventAny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnMilestoneEventCreated(t *testing.T) {
	type args struct {
		callbacks []MilestoneEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single MilestoneEventHandleFunc",
			args: args{
				callbacks: []MilestoneEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MilestoneEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple MilestoneEventHandleFunc",
			args: args{
				callbacks: []MilestoneEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MilestoneEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.MilestoneEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnMilestoneEventCreated(tt.args.callbacks...)
			if len(g.onMilestoneEvent["created"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onMilestoneEvent["created"]))
			}
		})
	}
}

func TestSetOnMilestoneEventCreated(t *testing.T) {
	type args struct {
		callbacks []MilestoneEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single MilestoneEventHandleFunc",
			args: args{
				[]MilestoneEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MilestoneEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple MilestoneEventHandleFuncs",
			args: args{
				[]MilestoneEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MilestoneEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.MilestoneEvent) error {
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
			g.SetOnMilestoneEventCreated(func(deliveryID string, eventName string, event *github.MilestoneEvent) error {
				return nil
			})
			g.SetOnMilestoneEventCreated(tt.args.callbacks...)
			if len(g.onMilestoneEvent["created"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onMilestoneEvent["created"]), tt.want)
			}
		})
	}
}

func TestHandleMilestoneEventCreated(t *testing.T) {
	action := "created"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.MilestoneEvent
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
				eventName:  "milestone",
				event:      &github.MilestoneEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "milestone",
				event:      &github.MilestoneEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "milestone",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "milestone",
				event:      &github.MilestoneEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "milestone",
				event:      &github.MilestoneEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "milestone",
				event:      &github.MilestoneEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnMilestoneEventCreated(func(deliveryID string, eventName string, event *github.MilestoneEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleMilestoneEventCreated(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleMilestoneEventCreated() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnMilestoneEventClosed(t *testing.T) {
	type args struct {
		callbacks []MilestoneEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single MilestoneEventHandleFunc",
			args: args{
				callbacks: []MilestoneEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MilestoneEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple MilestoneEventHandleFunc",
			args: args{
				callbacks: []MilestoneEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MilestoneEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.MilestoneEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnMilestoneEventClosed(tt.args.callbacks...)
			if len(g.onMilestoneEvent["closed"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onMilestoneEvent["closed"]))
			}
		})
	}
}

func TestSetOnMilestoneEventClosed(t *testing.T) {
	type args struct {
		callbacks []MilestoneEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single MilestoneEventHandleFunc",
			args: args{
				[]MilestoneEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MilestoneEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple MilestoneEventHandleFuncs",
			args: args{
				[]MilestoneEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MilestoneEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.MilestoneEvent) error {
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
			g.SetOnMilestoneEventClosed(func(deliveryID string, eventName string, event *github.MilestoneEvent) error {
				return nil
			})
			g.SetOnMilestoneEventClosed(tt.args.callbacks...)
			if len(g.onMilestoneEvent["closed"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onMilestoneEvent["closed"]), tt.want)
			}
		})
	}
}

func TestHandleMilestoneEventClosed(t *testing.T) {
	action := "closed"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.MilestoneEvent
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
				eventName:  "milestone",
				event:      &github.MilestoneEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "milestone",
				event:      &github.MilestoneEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "milestone",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "milestone",
				event:      &github.MilestoneEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "milestone",
				event:      &github.MilestoneEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "milestone",
				event:      &github.MilestoneEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnMilestoneEventClosed(func(deliveryID string, eventName string, event *github.MilestoneEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleMilestoneEventClosed(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleMilestoneEventClosed() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnMilestoneEventOpened(t *testing.T) {
	type args struct {
		callbacks []MilestoneEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single MilestoneEventHandleFunc",
			args: args{
				callbacks: []MilestoneEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MilestoneEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple MilestoneEventHandleFunc",
			args: args{
				callbacks: []MilestoneEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MilestoneEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.MilestoneEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnMilestoneEventOpened(tt.args.callbacks...)
			if len(g.onMilestoneEvent["opened"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onMilestoneEvent["opened"]))
			}
		})
	}
}

func TestSetOnMilestoneEventOpened(t *testing.T) {
	type args struct {
		callbacks []MilestoneEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single MilestoneEventHandleFunc",
			args: args{
				[]MilestoneEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MilestoneEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple MilestoneEventHandleFuncs",
			args: args{
				[]MilestoneEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MilestoneEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.MilestoneEvent) error {
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
			g.SetOnMilestoneEventOpened(func(deliveryID string, eventName string, event *github.MilestoneEvent) error {
				return nil
			})
			g.SetOnMilestoneEventOpened(tt.args.callbacks...)
			if len(g.onMilestoneEvent["opened"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onMilestoneEvent["opened"]), tt.want)
			}
		})
	}
}

func TestHandleMilestoneEventOpened(t *testing.T) {
	action := "opened"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.MilestoneEvent
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
				eventName:  "milestone",
				event:      &github.MilestoneEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "milestone",
				event:      &github.MilestoneEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "milestone",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "milestone",
				event:      &github.MilestoneEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "milestone",
				event:      &github.MilestoneEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "milestone",
				event:      &github.MilestoneEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnMilestoneEventOpened(func(deliveryID string, eventName string, event *github.MilestoneEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleMilestoneEventOpened(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleMilestoneEventOpened() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnMilestoneEventEdited(t *testing.T) {
	type args struct {
		callbacks []MilestoneEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single MilestoneEventHandleFunc",
			args: args{
				callbacks: []MilestoneEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MilestoneEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple MilestoneEventHandleFunc",
			args: args{
				callbacks: []MilestoneEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MilestoneEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.MilestoneEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnMilestoneEventEdited(tt.args.callbacks...)
			if len(g.onMilestoneEvent["edited"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onMilestoneEvent["edited"]))
			}
		})
	}
}

func TestSetOnMilestoneEventEdited(t *testing.T) {
	type args struct {
		callbacks []MilestoneEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single MilestoneEventHandleFunc",
			args: args{
				[]MilestoneEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MilestoneEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple MilestoneEventHandleFuncs",
			args: args{
				[]MilestoneEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MilestoneEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.MilestoneEvent) error {
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
			g.SetOnMilestoneEventEdited(func(deliveryID string, eventName string, event *github.MilestoneEvent) error {
				return nil
			})
			g.SetOnMilestoneEventEdited(tt.args.callbacks...)
			if len(g.onMilestoneEvent["edited"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onMilestoneEvent["edited"]), tt.want)
			}
		})
	}
}

func TestHandleMilestoneEventEdited(t *testing.T) {
	action := "edited"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.MilestoneEvent
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
				eventName:  "milestone",
				event:      &github.MilestoneEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "milestone",
				event:      &github.MilestoneEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "milestone",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "milestone",
				event:      &github.MilestoneEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "milestone",
				event:      &github.MilestoneEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "milestone",
				event:      &github.MilestoneEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnMilestoneEventEdited(func(deliveryID string, eventName string, event *github.MilestoneEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleMilestoneEventEdited(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleMilestoneEventEdited() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnMilestoneEventDeleted(t *testing.T) {
	type args struct {
		callbacks []MilestoneEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single MilestoneEventHandleFunc",
			args: args{
				callbacks: []MilestoneEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MilestoneEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple MilestoneEventHandleFunc",
			args: args{
				callbacks: []MilestoneEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MilestoneEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.MilestoneEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnMilestoneEventDeleted(tt.args.callbacks...)
			if len(g.onMilestoneEvent["deleted"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onMilestoneEvent["deleted"]))
			}
		})
	}
}

func TestSetOnMilestoneEventDeleted(t *testing.T) {
	type args struct {
		callbacks []MilestoneEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single MilestoneEventHandleFunc",
			args: args{
				[]MilestoneEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MilestoneEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple MilestoneEventHandleFuncs",
			args: args{
				[]MilestoneEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MilestoneEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.MilestoneEvent) error {
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
			g.SetOnMilestoneEventDeleted(func(deliveryID string, eventName string, event *github.MilestoneEvent) error {
				return nil
			})
			g.SetOnMilestoneEventDeleted(tt.args.callbacks...)
			if len(g.onMilestoneEvent["deleted"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onMilestoneEvent["deleted"]), tt.want)
			}
		})
	}
}

func TestHandleMilestoneEventDeleted(t *testing.T) {
	action := "deleted"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.MilestoneEvent
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
				eventName:  "milestone",
				event:      &github.MilestoneEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "milestone",
				event:      &github.MilestoneEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "milestone",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "milestone",
				event:      &github.MilestoneEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "milestone",
				event:      &github.MilestoneEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "milestone",
				event:      &github.MilestoneEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnMilestoneEventDeleted(func(deliveryID string, eventName string, event *github.MilestoneEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleMilestoneEventDeleted(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleMilestoneEventDeleted() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
