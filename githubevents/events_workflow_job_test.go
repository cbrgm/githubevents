package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"errors"
	"github.com/google/go-github/v43/github"
	"testing"
)

func TestOnWorkflowJobEventAny(t *testing.T) {
	type args struct {
		callbacks []WorkflowJobEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single WorkflowJobEventHandleFunc",
			args: args{
				[]WorkflowJobEventHandleFunc{
					func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple WorkflowJobEventHandleFuncs",
			args: args{
				[]WorkflowJobEventHandleFunc{
					func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnWorkflowJobEventAny(tt.args.callbacks...)
			if len(g.onWorkflowJobEvent["*"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onWorkflowJobEvent["*"]))
			}
		})
	}
}

func TestSetOnWorkflowJobEventAny(t *testing.T) {
	type args struct {
		callbacks []WorkflowJobEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single WorkflowJobEventHandleFunc",
			args: args{
				[]WorkflowJobEventHandleFunc{
					func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple WorkflowJobEventHandleFuncs",
			args: args{
				[]WorkflowJobEventHandleFunc{
					func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
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
			g.SetOnWorkflowJobEventAny(func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
				return nil
			})
			g.SetOnWorkflowJobEventAny(tt.args.callbacks...)
			if len(g.onWorkflowJobEvent["*"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onWorkflowJobEvent["*"]), tt.want)
			}
		})
	}
}

func TestHandleWorkflowJobEventAny(t *testing.T) {

	action := "*"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.WorkflowJobEvent
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
				eventName:  "workflow_job",

				event: &github.WorkflowJobEvent{Action: &action},

				fail: false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "workflow_job",

				event: &github.WorkflowJobEvent{Action: &action},

				fail: true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "workflow_job",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnWorkflowJobEventAny(func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleWorkflowJobEventAny(tt.args.deliveryID, tt.args.deliveryID, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("TestHandleWorkflowJobEventAny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnWorkflowJobEventQueued(t *testing.T) {
	type args struct {
		callbacks []WorkflowJobEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single WorkflowJobEventHandleFunc",
			args: args{
				callbacks: []WorkflowJobEventHandleFunc{
					func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple WorkflowJobEventHandleFunc",
			args: args{
				callbacks: []WorkflowJobEventHandleFunc{
					func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnWorkflowJobEventQueued(tt.args.callbacks...)
			if len(g.onWorkflowJobEvent["queued"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onWorkflowJobEvent["queued"]))
			}
		})
	}
}

func TestSetOnWorkflowJobEventQueued(t *testing.T) {
	type args struct {
		callbacks []WorkflowJobEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single WorkflowJobEventHandleFunc",
			args: args{
				[]WorkflowJobEventHandleFunc{
					func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple WorkflowJobEventHandleFuncs",
			args: args{
				[]WorkflowJobEventHandleFunc{
					func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
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
			g.SetOnWorkflowJobEventQueued(func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
				return nil
			})
			g.SetOnWorkflowJobEventQueued(tt.args.callbacks...)
			if len(g.onWorkflowJobEvent["queued"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onWorkflowJobEvent["queued"]), tt.want)
			}
		})
	}
}

func TestHandleWorkflowJobEventQueued(t *testing.T) {
	action := "queued"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.WorkflowJobEvent
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
				eventName:  "workflow_job",
				event:      &github.WorkflowJobEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "workflow_job",
				event:      &github.WorkflowJobEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "workflow_job",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "workflow_job",
				event:      &github.WorkflowJobEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "workflow_job",
				event:      &github.WorkflowJobEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "workflow_job",
				event:      &github.WorkflowJobEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnWorkflowJobEventQueued(func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleWorkflowJobEventQueued(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleWorkflowJobEventQueued() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnWorkflowJobEventInProgress(t *testing.T) {
	type args struct {
		callbacks []WorkflowJobEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single WorkflowJobEventHandleFunc",
			args: args{
				callbacks: []WorkflowJobEventHandleFunc{
					func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple WorkflowJobEventHandleFunc",
			args: args{
				callbacks: []WorkflowJobEventHandleFunc{
					func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnWorkflowJobEventInProgress(tt.args.callbacks...)
			if len(g.onWorkflowJobEvent["in_progress"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onWorkflowJobEvent["in_progress"]))
			}
		})
	}
}

func TestSetOnWorkflowJobEventInProgress(t *testing.T) {
	type args struct {
		callbacks []WorkflowJobEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single WorkflowJobEventHandleFunc",
			args: args{
				[]WorkflowJobEventHandleFunc{
					func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple WorkflowJobEventHandleFuncs",
			args: args{
				[]WorkflowJobEventHandleFunc{
					func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
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
			g.SetOnWorkflowJobEventInProgress(func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
				return nil
			})
			g.SetOnWorkflowJobEventInProgress(tt.args.callbacks...)
			if len(g.onWorkflowJobEvent["in_progress"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onWorkflowJobEvent["in_progress"]), tt.want)
			}
		})
	}
}

func TestHandleWorkflowJobEventInProgress(t *testing.T) {
	action := "in_progress"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.WorkflowJobEvent
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
				eventName:  "workflow_job",
				event:      &github.WorkflowJobEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "workflow_job",
				event:      &github.WorkflowJobEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "workflow_job",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "workflow_job",
				event:      &github.WorkflowJobEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "workflow_job",
				event:      &github.WorkflowJobEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "workflow_job",
				event:      &github.WorkflowJobEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnWorkflowJobEventInProgress(func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleWorkflowJobEventInProgress(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleWorkflowJobEventInProgress() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnWorkflowJobEventCompleted(t *testing.T) {
	type args struct {
		callbacks []WorkflowJobEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single WorkflowJobEventHandleFunc",
			args: args{
				callbacks: []WorkflowJobEventHandleFunc{
					func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple WorkflowJobEventHandleFunc",
			args: args{
				callbacks: []WorkflowJobEventHandleFunc{
					func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnWorkflowJobEventCompleted(tt.args.callbacks...)
			if len(g.onWorkflowJobEvent["completed"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onWorkflowJobEvent["completed"]))
			}
		})
	}
}

func TestSetOnWorkflowJobEventCompleted(t *testing.T) {
	type args struct {
		callbacks []WorkflowJobEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single WorkflowJobEventHandleFunc",
			args: args{
				[]WorkflowJobEventHandleFunc{
					func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple WorkflowJobEventHandleFuncs",
			args: args{
				[]WorkflowJobEventHandleFunc{
					func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
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
			g.SetOnWorkflowJobEventCompleted(func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
				return nil
			})
			g.SetOnWorkflowJobEventCompleted(tt.args.callbacks...)
			if len(g.onWorkflowJobEvent["completed"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onWorkflowJobEvent["completed"]), tt.want)
			}
		})
	}
}

func TestHandleWorkflowJobEventCompleted(t *testing.T) {
	action := "completed"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.WorkflowJobEvent
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
				eventName:  "workflow_job",
				event:      &github.WorkflowJobEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "workflow_job",
				event:      &github.WorkflowJobEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "workflow_job",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "workflow_job",
				event:      &github.WorkflowJobEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "workflow_job",
				event:      &github.WorkflowJobEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "workflow_job",
				event:      &github.WorkflowJobEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnWorkflowJobEventCompleted(func(deliveryID string, eventName string, event *github.WorkflowJobEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleWorkflowJobEventCompleted(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleWorkflowJobEventCompleted() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
