package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"errors"
	"github.com/google/go-github/v43/github"
	"testing"
)

func TestOnProjectEventAny(t *testing.T) {
	type args struct {
		callbacks []ProjectEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single ProjectEventHandleFunc",
			args: args{
				[]ProjectEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple ProjectEventHandleFuncs",
			args: args{
				[]ProjectEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectEventAny(tt.args.callbacks...)
			if len(g.onProjectEvent["*"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onProjectEvent["*"]))
			}
		})
	}
}

func TestSetOnProjectEventAny(t *testing.T) {
	type args struct {
		callbacks []ProjectEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single ProjectEventHandleFunc",
			args: args{
				[]ProjectEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple ProjectEventHandleFuncs",
			args: args{
				[]ProjectEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectEvent) error {
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
			g.SetOnProjectEventAny(func(deliveryID string, eventName string, event *github.ProjectEvent) error {
				return nil
			})
			g.SetOnProjectEventAny(tt.args.callbacks...)
			if len(g.onProjectEvent["*"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onProjectEvent["*"]), tt.want)
			}
		})
	}
}

func TestHandleProjectEventAny(t *testing.T) {

	action := "*"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.ProjectEvent
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
				eventName:  "project",

				event: &github.ProjectEvent{Action: &action},

				fail: false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "project",

				event: &github.ProjectEvent{Action: &action},

				fail: true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "project",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectEventAny(func(deliveryID string, eventName string, event *github.ProjectEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleProjectEventAny(tt.args.deliveryID, tt.args.deliveryID, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("TestHandleProjectEventAny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnProjectEventCreated(t *testing.T) {
	type args struct {
		callbacks []ProjectEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single ProjectEventHandleFunc",
			args: args{
				callbacks: []ProjectEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple ProjectEventHandleFunc",
			args: args{
				callbacks: []ProjectEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectEventCreated(tt.args.callbacks...)
			if len(g.onProjectEvent["created"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onProjectEvent["created"]))
			}
		})
	}
}

func TestSetOnProjectEventCreated(t *testing.T) {
	type args struct {
		callbacks []ProjectEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single ProjectEventHandleFunc",
			args: args{
				[]ProjectEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple ProjectEventHandleFuncs",
			args: args{
				[]ProjectEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectEvent) error {
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
			g.SetOnProjectEventCreated(func(deliveryID string, eventName string, event *github.ProjectEvent) error {
				return nil
			})
			g.SetOnProjectEventCreated(tt.args.callbacks...)
			if len(g.onProjectEvent["created"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onProjectEvent["created"]), tt.want)
			}
		})
	}
}

func TestHandleProjectEventCreated(t *testing.T) {
	action := "created"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.ProjectEvent
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
				eventName:  "project",
				event:      &github.ProjectEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "project",
				event:      &github.ProjectEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "project",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "project",
				event:      &github.ProjectEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "project",
				event:      &github.ProjectEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "project",
				event:      &github.ProjectEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectEventCreated(func(deliveryID string, eventName string, event *github.ProjectEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleProjectEventCreated(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleProjectEventCreated() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnProjectEventEdited(t *testing.T) {
	type args struct {
		callbacks []ProjectEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single ProjectEventHandleFunc",
			args: args{
				callbacks: []ProjectEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple ProjectEventHandleFunc",
			args: args{
				callbacks: []ProjectEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectEventEdited(tt.args.callbacks...)
			if len(g.onProjectEvent["edited"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onProjectEvent["edited"]))
			}
		})
	}
}

func TestSetOnProjectEventEdited(t *testing.T) {
	type args struct {
		callbacks []ProjectEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single ProjectEventHandleFunc",
			args: args{
				[]ProjectEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple ProjectEventHandleFuncs",
			args: args{
				[]ProjectEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectEvent) error {
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
			g.SetOnProjectEventEdited(func(deliveryID string, eventName string, event *github.ProjectEvent) error {
				return nil
			})
			g.SetOnProjectEventEdited(tt.args.callbacks...)
			if len(g.onProjectEvent["edited"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onProjectEvent["edited"]), tt.want)
			}
		})
	}
}

func TestHandleProjectEventEdited(t *testing.T) {
	action := "edited"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.ProjectEvent
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
				eventName:  "project",
				event:      &github.ProjectEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "project",
				event:      &github.ProjectEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "project",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "project",
				event:      &github.ProjectEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "project",
				event:      &github.ProjectEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "project",
				event:      &github.ProjectEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectEventEdited(func(deliveryID string, eventName string, event *github.ProjectEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleProjectEventEdited(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleProjectEventEdited() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnProjectEventClosed(t *testing.T) {
	type args struct {
		callbacks []ProjectEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single ProjectEventHandleFunc",
			args: args{
				callbacks: []ProjectEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple ProjectEventHandleFunc",
			args: args{
				callbacks: []ProjectEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectEventClosed(tt.args.callbacks...)
			if len(g.onProjectEvent["closed"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onProjectEvent["closed"]))
			}
		})
	}
}

func TestSetOnProjectEventClosed(t *testing.T) {
	type args struct {
		callbacks []ProjectEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single ProjectEventHandleFunc",
			args: args{
				[]ProjectEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple ProjectEventHandleFuncs",
			args: args{
				[]ProjectEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectEvent) error {
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
			g.SetOnProjectEventClosed(func(deliveryID string, eventName string, event *github.ProjectEvent) error {
				return nil
			})
			g.SetOnProjectEventClosed(tt.args.callbacks...)
			if len(g.onProjectEvent["closed"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onProjectEvent["closed"]), tt.want)
			}
		})
	}
}

func TestHandleProjectEventClosed(t *testing.T) {
	action := "closed"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.ProjectEvent
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
				eventName:  "project",
				event:      &github.ProjectEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "project",
				event:      &github.ProjectEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "project",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "project",
				event:      &github.ProjectEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "project",
				event:      &github.ProjectEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "project",
				event:      &github.ProjectEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectEventClosed(func(deliveryID string, eventName string, event *github.ProjectEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleProjectEventClosed(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleProjectEventClosed() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnProjectEventReopened(t *testing.T) {
	type args struct {
		callbacks []ProjectEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single ProjectEventHandleFunc",
			args: args{
				callbacks: []ProjectEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple ProjectEventHandleFunc",
			args: args{
				callbacks: []ProjectEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectEventReopened(tt.args.callbacks...)
			if len(g.onProjectEvent["reopened"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onProjectEvent["reopened"]))
			}
		})
	}
}

func TestSetOnProjectEventReopened(t *testing.T) {
	type args struct {
		callbacks []ProjectEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single ProjectEventHandleFunc",
			args: args{
				[]ProjectEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple ProjectEventHandleFuncs",
			args: args{
				[]ProjectEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectEvent) error {
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
			g.SetOnProjectEventReopened(func(deliveryID string, eventName string, event *github.ProjectEvent) error {
				return nil
			})
			g.SetOnProjectEventReopened(tt.args.callbacks...)
			if len(g.onProjectEvent["reopened"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onProjectEvent["reopened"]), tt.want)
			}
		})
	}
}

func TestHandleProjectEventReopened(t *testing.T) {
	action := "reopened"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.ProjectEvent
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
				eventName:  "project",
				event:      &github.ProjectEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "project",
				event:      &github.ProjectEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "project",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "project",
				event:      &github.ProjectEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "project",
				event:      &github.ProjectEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "project",
				event:      &github.ProjectEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectEventReopened(func(deliveryID string, eventName string, event *github.ProjectEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleProjectEventReopened(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleProjectEventReopened() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnProjectEventDeleted(t *testing.T) {
	type args struct {
		callbacks []ProjectEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single ProjectEventHandleFunc",
			args: args{
				callbacks: []ProjectEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple ProjectEventHandleFunc",
			args: args{
				callbacks: []ProjectEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectEventDeleted(tt.args.callbacks...)
			if len(g.onProjectEvent["deleted"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onProjectEvent["deleted"]))
			}
		})
	}
}

func TestSetOnProjectEventDeleted(t *testing.T) {
	type args struct {
		callbacks []ProjectEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single ProjectEventHandleFunc",
			args: args{
				[]ProjectEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple ProjectEventHandleFuncs",
			args: args{
				[]ProjectEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectEvent) error {
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
			g.SetOnProjectEventDeleted(func(deliveryID string, eventName string, event *github.ProjectEvent) error {
				return nil
			})
			g.SetOnProjectEventDeleted(tt.args.callbacks...)
			if len(g.onProjectEvent["deleted"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onProjectEvent["deleted"]), tt.want)
			}
		})
	}
}

func TestHandleProjectEventDeleted(t *testing.T) {
	action := "deleted"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.ProjectEvent
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
				eventName:  "project",
				event:      &github.ProjectEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "project",
				event:      &github.ProjectEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "project",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "project",
				event:      &github.ProjectEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "project",
				event:      &github.ProjectEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "project",
				event:      &github.ProjectEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectEventDeleted(func(deliveryID string, eventName string, event *github.ProjectEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleProjectEventDeleted(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleProjectEventDeleted() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
