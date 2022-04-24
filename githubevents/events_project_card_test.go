package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"errors"
	"github.com/google/go-github/v43/github"
	"testing"
)

func TestOnProjectCardEventAny(t *testing.T) {
	type args struct {
		callbacks []ProjectCardEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single ProjectCardEventHandleFunc",
			args: args{
				[]ProjectCardEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple ProjectCardEventHandleFuncs",
			args: args{
				[]ProjectCardEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectCardEventAny(tt.args.callbacks...)
			if len(g.onProjectCardEvent["*"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onProjectCardEvent["*"]))
			}
		})
	}
}

func TestSetOnProjectCardEventAny(t *testing.T) {
	type args struct {
		callbacks []ProjectCardEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single ProjectCardEventHandleFunc",
			args: args{
				[]ProjectCardEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple ProjectCardEventHandleFuncs",
			args: args{
				[]ProjectCardEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
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
			g.SetOnProjectCardEventAny(func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
				return nil
			})
			g.SetOnProjectCardEventAny(tt.args.callbacks...)
			if len(g.onProjectCardEvent["*"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onProjectCardEvent["*"]), tt.want)
			}
		})
	}
}

func TestHandleProjectCardEventAny(t *testing.T) {

	action := "*"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.ProjectCardEvent
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
				eventName:  "project_card",

				event: &github.ProjectCardEvent{Action: &action},

				fail: false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "project_card",

				event: &github.ProjectCardEvent{Action: &action},

				fail: true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "project_card",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectCardEventAny(func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleProjectCardEventAny(tt.args.deliveryID, tt.args.deliveryID, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("TestHandleProjectCardEventAny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnProjectCardEventCreated(t *testing.T) {
	type args struct {
		callbacks []ProjectCardEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single ProjectCardEventHandleFunc",
			args: args{
				callbacks: []ProjectCardEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple ProjectCardEventHandleFunc",
			args: args{
				callbacks: []ProjectCardEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectCardEventCreated(tt.args.callbacks...)
			if len(g.onProjectCardEvent["created"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onProjectCardEvent["created"]))
			}
		})
	}
}

func TestSetOnProjectCardEventCreated(t *testing.T) {
	type args struct {
		callbacks []ProjectCardEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single ProjectCardEventHandleFunc",
			args: args{
				[]ProjectCardEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple ProjectCardEventHandleFuncs",
			args: args{
				[]ProjectCardEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
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
			g.SetOnProjectCardEventCreated(func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
				return nil
			})
			g.SetOnProjectCardEventCreated(tt.args.callbacks...)
			if len(g.onProjectCardEvent["created"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onProjectCardEvent["created"]), tt.want)
			}
		})
	}
}

func TestHandleProjectCardEventCreated(t *testing.T) {
	action := "created"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.ProjectCardEvent
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
				eventName:  "project_card",
				event:      &github.ProjectCardEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "project_card",
				event:      &github.ProjectCardEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "project_card",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "project_card",
				event:      &github.ProjectCardEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "project_card",
				event:      &github.ProjectCardEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "project_card",
				event:      &github.ProjectCardEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectCardEventCreated(func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleProjectCardEventCreated(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleProjectCardEventCreated() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnProjectCardEventEdited(t *testing.T) {
	type args struct {
		callbacks []ProjectCardEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single ProjectCardEventHandleFunc",
			args: args{
				callbacks: []ProjectCardEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple ProjectCardEventHandleFunc",
			args: args{
				callbacks: []ProjectCardEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectCardEventEdited(tt.args.callbacks...)
			if len(g.onProjectCardEvent["edited"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onProjectCardEvent["edited"]))
			}
		})
	}
}

func TestSetOnProjectCardEventEdited(t *testing.T) {
	type args struct {
		callbacks []ProjectCardEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single ProjectCardEventHandleFunc",
			args: args{
				[]ProjectCardEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple ProjectCardEventHandleFuncs",
			args: args{
				[]ProjectCardEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
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
			g.SetOnProjectCardEventEdited(func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
				return nil
			})
			g.SetOnProjectCardEventEdited(tt.args.callbacks...)
			if len(g.onProjectCardEvent["edited"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onProjectCardEvent["edited"]), tt.want)
			}
		})
	}
}

func TestHandleProjectCardEventEdited(t *testing.T) {
	action := "edited"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.ProjectCardEvent
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
				eventName:  "project_card",
				event:      &github.ProjectCardEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "project_card",
				event:      &github.ProjectCardEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "project_card",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "project_card",
				event:      &github.ProjectCardEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "project_card",
				event:      &github.ProjectCardEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "project_card",
				event:      &github.ProjectCardEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectCardEventEdited(func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleProjectCardEventEdited(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleProjectCardEventEdited() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnProjectCardEventConverted(t *testing.T) {
	type args struct {
		callbacks []ProjectCardEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single ProjectCardEventHandleFunc",
			args: args{
				callbacks: []ProjectCardEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple ProjectCardEventHandleFunc",
			args: args{
				callbacks: []ProjectCardEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectCardEventConverted(tt.args.callbacks...)
			if len(g.onProjectCardEvent["converted"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onProjectCardEvent["converted"]))
			}
		})
	}
}

func TestSetOnProjectCardEventConverted(t *testing.T) {
	type args struct {
		callbacks []ProjectCardEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single ProjectCardEventHandleFunc",
			args: args{
				[]ProjectCardEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple ProjectCardEventHandleFuncs",
			args: args{
				[]ProjectCardEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
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
			g.SetOnProjectCardEventConverted(func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
				return nil
			})
			g.SetOnProjectCardEventConverted(tt.args.callbacks...)
			if len(g.onProjectCardEvent["converted"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onProjectCardEvent["converted"]), tt.want)
			}
		})
	}
}

func TestHandleProjectCardEventConverted(t *testing.T) {
	action := "converted"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.ProjectCardEvent
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
				eventName:  "project_card",
				event:      &github.ProjectCardEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "project_card",
				event:      &github.ProjectCardEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "project_card",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "project_card",
				event:      &github.ProjectCardEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "project_card",
				event:      &github.ProjectCardEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "project_card",
				event:      &github.ProjectCardEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectCardEventConverted(func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleProjectCardEventConverted(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleProjectCardEventConverted() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnProjectCardEventMoved(t *testing.T) {
	type args struct {
		callbacks []ProjectCardEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single ProjectCardEventHandleFunc",
			args: args{
				callbacks: []ProjectCardEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple ProjectCardEventHandleFunc",
			args: args{
				callbacks: []ProjectCardEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectCardEventMoved(tt.args.callbacks...)
			if len(g.onProjectCardEvent["moved"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onProjectCardEvent["moved"]))
			}
		})
	}
}

func TestSetOnProjectCardEventMoved(t *testing.T) {
	type args struct {
		callbacks []ProjectCardEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single ProjectCardEventHandleFunc",
			args: args{
				[]ProjectCardEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple ProjectCardEventHandleFuncs",
			args: args{
				[]ProjectCardEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
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
			g.SetOnProjectCardEventMoved(func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
				return nil
			})
			g.SetOnProjectCardEventMoved(tt.args.callbacks...)
			if len(g.onProjectCardEvent["moved"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onProjectCardEvent["moved"]), tt.want)
			}
		})
	}
}

func TestHandleProjectCardEventMoved(t *testing.T) {
	action := "moved"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.ProjectCardEvent
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
				eventName:  "project_card",
				event:      &github.ProjectCardEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "project_card",
				event:      &github.ProjectCardEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "project_card",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "project_card",
				event:      &github.ProjectCardEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "project_card",
				event:      &github.ProjectCardEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "project_card",
				event:      &github.ProjectCardEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectCardEventMoved(func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleProjectCardEventMoved(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleProjectCardEventMoved() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnProjectCardEventDeleted(t *testing.T) {
	type args struct {
		callbacks []ProjectCardEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single ProjectCardEventHandleFunc",
			args: args{
				callbacks: []ProjectCardEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple ProjectCardEventHandleFunc",
			args: args{
				callbacks: []ProjectCardEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectCardEventDeleted(tt.args.callbacks...)
			if len(g.onProjectCardEvent["deleted"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onProjectCardEvent["deleted"]))
			}
		})
	}
}

func TestSetOnProjectCardEventDeleted(t *testing.T) {
	type args struct {
		callbacks []ProjectCardEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single ProjectCardEventHandleFunc",
			args: args{
				[]ProjectCardEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple ProjectCardEventHandleFuncs",
			args: args{
				[]ProjectCardEventHandleFunc{
					func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
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
			g.SetOnProjectCardEventDeleted(func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
				return nil
			})
			g.SetOnProjectCardEventDeleted(tt.args.callbacks...)
			if len(g.onProjectCardEvent["deleted"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onProjectCardEvent["deleted"]), tt.want)
			}
		})
	}
}

func TestHandleProjectCardEventDeleted(t *testing.T) {
	action := "deleted"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.ProjectCardEvent
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
				eventName:  "project_card",
				event:      &github.ProjectCardEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "project_card",
				event:      &github.ProjectCardEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "project_card",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "project_card",
				event:      &github.ProjectCardEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "project_card",
				event:      &github.ProjectCardEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "project_card",
				event:      &github.ProjectCardEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectCardEventDeleted(func(deliveryID string, eventName string, event *github.ProjectCardEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleProjectCardEventDeleted(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleProjectCardEventDeleted() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
