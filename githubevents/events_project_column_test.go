package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"errors"
	"github.com/google/go-github/v43/github"
	"testing"
)

func TestOnProjectColumnEventAny(t *testing.T) {
	type args struct {
		callbacks []ProjectColumnEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single ProjectColumnEventHandleFunc",
			args: args{
				[]ProjectColumnEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple ProjectColumnEventHandleFuncs",
			args: args{
				[]ProjectColumnEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectColumnEventAny(tt.args.callbacks...)
			if len(g.onProjectColumnEvent["*"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onProjectColumnEvent["*"]))
			}
		})
	}
}

func TestSetOnProjectColumnEventAny(t *testing.T) {
	type args struct {
		callbacks []ProjectColumnEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single ProjectColumnEventHandleFunc",
			args: args{
				[]ProjectColumnEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple ProjectColumnEventHandleFuncs",
			args: args{
				[]ProjectColumnEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
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
			g.SetOnProjectColumnEventAny(func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
				return nil
			})
			g.SetOnProjectColumnEventAny(tt.args.callbacks...)
			if len(g.onProjectColumnEvent["*"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onProjectColumnEvent["*"]), tt.want)
			}
		})
	}
}

func TestHandleProjectColumnEventAny(t *testing.T) {

	action := "*"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.ProjectColumnEvent
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
				eventName:  "project_column",

				event: &github.ProjectColumnEvent{Action: &action},

				fail: false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "project_column",

				event: &github.ProjectColumnEvent{Action: &action},

				fail: true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "project_column",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectColumnEventAny(func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleProjectColumnEventAny(tt.args.deliveryID, tt.args.deliveryID, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("TestHandleProjectColumnEventAny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnProjectColumnEventCreated(t *testing.T) {
	type args struct {
		callbacks []ProjectColumnEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single ProjectColumnEventHandleFunc",
			args: args{
				callbacks: []ProjectColumnEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple ProjectColumnEventHandleFunc",
			args: args{
				callbacks: []ProjectColumnEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectColumnEventCreated(tt.args.callbacks...)
			if len(g.onProjectColumnEvent["created"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onProjectColumnEvent["created"]))
			}
		})
	}
}

func TestSetOnProjectColumnEventCreated(t *testing.T) {
	type args struct {
		callbacks []ProjectColumnEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single ProjectColumnEventHandleFunc",
			args: args{
				[]ProjectColumnEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple ProjectColumnEventHandleFuncs",
			args: args{
				[]ProjectColumnEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
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
			g.SetOnProjectColumnEventCreated(func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
				return nil
			})
			g.SetOnProjectColumnEventCreated(tt.args.callbacks...)
			if len(g.onProjectColumnEvent["created"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onProjectColumnEvent["created"]), tt.want)
			}
		})
	}
}

func TestHandleProjectColumnEventCreated(t *testing.T) {
	action := "created"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.ProjectColumnEvent
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
				eventName:  "project_column",
				event:      &github.ProjectColumnEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "project_column",
				event:      &github.ProjectColumnEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "project_column",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "project_column",
				event:      &github.ProjectColumnEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "project_column",
				event:      &github.ProjectColumnEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "project_column",
				event:      &github.ProjectColumnEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectColumnEventCreated(func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleProjectColumnEventCreated(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleProjectColumnEventCreated() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnProjectColumnEventEdited(t *testing.T) {
	type args struct {
		callbacks []ProjectColumnEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single ProjectColumnEventHandleFunc",
			args: args{
				callbacks: []ProjectColumnEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple ProjectColumnEventHandleFunc",
			args: args{
				callbacks: []ProjectColumnEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectColumnEventEdited(tt.args.callbacks...)
			if len(g.onProjectColumnEvent["edited"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onProjectColumnEvent["edited"]))
			}
		})
	}
}

func TestSetOnProjectColumnEventEdited(t *testing.T) {
	type args struct {
		callbacks []ProjectColumnEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single ProjectColumnEventHandleFunc",
			args: args{
				[]ProjectColumnEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple ProjectColumnEventHandleFuncs",
			args: args{
				[]ProjectColumnEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
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
			g.SetOnProjectColumnEventEdited(func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
				return nil
			})
			g.SetOnProjectColumnEventEdited(tt.args.callbacks...)
			if len(g.onProjectColumnEvent["edited"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onProjectColumnEvent["edited"]), tt.want)
			}
		})
	}
}

func TestHandleProjectColumnEventEdited(t *testing.T) {
	action := "edited"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.ProjectColumnEvent
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
				eventName:  "project_column",
				event:      &github.ProjectColumnEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "project_column",
				event:      &github.ProjectColumnEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "project_column",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "project_column",
				event:      &github.ProjectColumnEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "project_column",
				event:      &github.ProjectColumnEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "project_column",
				event:      &github.ProjectColumnEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectColumnEventEdited(func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleProjectColumnEventEdited(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleProjectColumnEventEdited() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnProjectColumnEventMoved(t *testing.T) {
	type args struct {
		callbacks []ProjectColumnEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single ProjectColumnEventHandleFunc",
			args: args{
				callbacks: []ProjectColumnEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple ProjectColumnEventHandleFunc",
			args: args{
				callbacks: []ProjectColumnEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectColumnEventMoved(tt.args.callbacks...)
			if len(g.onProjectColumnEvent["moved"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onProjectColumnEvent["moved"]))
			}
		})
	}
}

func TestSetOnProjectColumnEventMoved(t *testing.T) {
	type args struct {
		callbacks []ProjectColumnEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single ProjectColumnEventHandleFunc",
			args: args{
				[]ProjectColumnEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple ProjectColumnEventHandleFuncs",
			args: args{
				[]ProjectColumnEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
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
			g.SetOnProjectColumnEventMoved(func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
				return nil
			})
			g.SetOnProjectColumnEventMoved(tt.args.callbacks...)
			if len(g.onProjectColumnEvent["moved"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onProjectColumnEvent["moved"]), tt.want)
			}
		})
	}
}

func TestHandleProjectColumnEventMoved(t *testing.T) {
	action := "moved"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.ProjectColumnEvent
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
				eventName:  "project_column",
				event:      &github.ProjectColumnEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "project_column",
				event:      &github.ProjectColumnEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "project_column",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "project_column",
				event:      &github.ProjectColumnEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "project_column",
				event:      &github.ProjectColumnEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "project_column",
				event:      &github.ProjectColumnEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectColumnEventMoved(func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleProjectColumnEventMoved(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleProjectColumnEventMoved() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnProjectColumnEventDeleted(t *testing.T) {
	type args struct {
		callbacks []ProjectColumnEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single ProjectColumnEventHandleFunc",
			args: args{
				callbacks: []ProjectColumnEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple ProjectColumnEventHandleFunc",
			args: args{
				callbacks: []ProjectColumnEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectColumnEventDeleted(tt.args.callbacks...)
			if len(g.onProjectColumnEvent["deleted"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onProjectColumnEvent["deleted"]))
			}
		})
	}
}

func TestSetOnProjectColumnEventDeleted(t *testing.T) {
	type args struct {
		callbacks []ProjectColumnEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single ProjectColumnEventHandleFunc",
			args: args{
				[]ProjectColumnEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple ProjectColumnEventHandleFuncs",
			args: args{
				[]ProjectColumnEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
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
			g.SetOnProjectColumnEventDeleted(func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
				return nil
			})
			g.SetOnProjectColumnEventDeleted(tt.args.callbacks...)
			if len(g.onProjectColumnEvent["deleted"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onProjectColumnEvent["deleted"]), tt.want)
			}
		})
	}
}

func TestHandleProjectColumnEventDeleted(t *testing.T) {
	action := "deleted"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.ProjectColumnEvent
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
				eventName:  "project_column",
				event:      &github.ProjectColumnEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "project_column",
				event:      &github.ProjectColumnEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "project_column",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "project_column",
				event:      &github.ProjectColumnEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "project_column",
				event:      &github.ProjectColumnEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "project_column",
				event:      &github.ProjectColumnEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectColumnEventDeleted(func(deliveryID string, eventName string, event *github.ProjectColumnEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleProjectColumnEventDeleted(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleProjectColumnEventDeleted() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
