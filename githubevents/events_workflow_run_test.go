package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"errors"
	"github.com/google/go-github/v43/github"
	"testing"
)

func TestOnWorkflowRunEventAny(t *testing.T) {
	type args struct {
		callbacks []WorkflowRunEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single WorkflowRunEventHandleFunc",
			args: args{
				[]WorkflowRunEventHandleFunc{
					func(deliveryID string, eventName string, event *github.WorkflowRunEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple WorkflowRunEventHandleFuncs",
			args: args{
				[]WorkflowRunEventHandleFunc{
					func(deliveryID string, eventName string, event *github.WorkflowRunEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.WorkflowRunEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnWorkflowRunEventAny(tt.args.callbacks...)
			if len(g.onWorkflowRunEvent["*"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onWorkflowRunEvent["*"]))
			}
		})
	}
}

func TestSetOnWorkflowRunEventAny(t *testing.T) {
	type args struct {
		callbacks []WorkflowRunEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single WorkflowRunEventHandleFunc",
			args: args{
				[]WorkflowRunEventHandleFunc{
					func(deliveryID string, eventName string, event *github.WorkflowRunEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple WorkflowRunEventHandleFuncs",
			args: args{
				[]WorkflowRunEventHandleFunc{
					func(deliveryID string, eventName string, event *github.WorkflowRunEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.WorkflowRunEvent) error {
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
			g.SetOnWorkflowRunEventAny(func(deliveryID string, eventName string, event *github.WorkflowRunEvent) error {
				return nil
			})
			g.SetOnWorkflowRunEventAny(tt.args.callbacks...)
			if len(g.onWorkflowRunEvent["*"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onWorkflowRunEvent["*"]), tt.want)
			}
		})
	}
}

func TestHandleWorkflowRunEventAny(t *testing.T) {

	action := "*"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.WorkflowRunEvent
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
				eventName:  "workflow_run",

				event: &github.WorkflowRunEvent{Action: &action},

				fail: false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "workflow_run",

				event: &github.WorkflowRunEvent{Action: &action},

				fail: true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "workflow_run",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnWorkflowRunEventAny(func(deliveryID string, eventName string, event *github.WorkflowRunEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleWorkflowRunEventAny(tt.args.deliveryID, tt.args.deliveryID, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("TestHandleWorkflowRunEventAny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnWorkflowRunEventRequested(t *testing.T) {
	type args struct {
		callbacks []WorkflowRunEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single WorkflowRunEventHandleFunc",
			args: args{
				callbacks: []WorkflowRunEventHandleFunc{
					func(deliveryID string, eventName string, event *github.WorkflowRunEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple WorkflowRunEventHandleFunc",
			args: args{
				callbacks: []WorkflowRunEventHandleFunc{
					func(deliveryID string, eventName string, event *github.WorkflowRunEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.WorkflowRunEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnWorkflowRunEventRequested(tt.args.callbacks...)
			if len(g.onWorkflowRunEvent["requested"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onWorkflowRunEvent["requested"]))
			}
		})
	}
}

func TestSetOnWorkflowRunEventRequested(t *testing.T) {
	type args struct {
		callbacks []WorkflowRunEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single WorkflowRunEventHandleFunc",
			args: args{
				[]WorkflowRunEventHandleFunc{
					func(deliveryID string, eventName string, event *github.WorkflowRunEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple WorkflowRunEventHandleFuncs",
			args: args{
				[]WorkflowRunEventHandleFunc{
					func(deliveryID string, eventName string, event *github.WorkflowRunEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.WorkflowRunEvent) error {
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
			g.SetOnWorkflowRunEventRequested(func(deliveryID string, eventName string, event *github.WorkflowRunEvent) error {
				return nil
			})
			g.SetOnWorkflowRunEventRequested(tt.args.callbacks...)
			if len(g.onWorkflowRunEvent["requested"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onWorkflowRunEvent["requested"]), tt.want)
			}
		})
	}
}

func TestHandleWorkflowRunEventRequested(t *testing.T) {
	action := "requested"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.WorkflowRunEvent
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
				eventName:  "workflow_run",
				event:      &github.WorkflowRunEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "workflow_run",
				event:      &github.WorkflowRunEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "workflow_run",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "workflow_run",
				event:      &github.WorkflowRunEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "workflow_run",
				event:      &github.WorkflowRunEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "workflow_run",
				event:      &github.WorkflowRunEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnWorkflowRunEventRequested(func(deliveryID string, eventName string, event *github.WorkflowRunEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleWorkflowRunEventRequested(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleWorkflowRunEventRequested() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnWorkflowRunEventCompleted(t *testing.T) {
	type args struct {
		callbacks []WorkflowRunEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single WorkflowRunEventHandleFunc",
			args: args{
				callbacks: []WorkflowRunEventHandleFunc{
					func(deliveryID string, eventName string, event *github.WorkflowRunEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple WorkflowRunEventHandleFunc",
			args: args{
				callbacks: []WorkflowRunEventHandleFunc{
					func(deliveryID string, eventName string, event *github.WorkflowRunEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.WorkflowRunEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnWorkflowRunEventCompleted(tt.args.callbacks...)
			if len(g.onWorkflowRunEvent["completed"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onWorkflowRunEvent["completed"]))
			}
		})
	}
}

func TestSetOnWorkflowRunEventCompleted(t *testing.T) {
	type args struct {
		callbacks []WorkflowRunEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single WorkflowRunEventHandleFunc",
			args: args{
				[]WorkflowRunEventHandleFunc{
					func(deliveryID string, eventName string, event *github.WorkflowRunEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple WorkflowRunEventHandleFuncs",
			args: args{
				[]WorkflowRunEventHandleFunc{
					func(deliveryID string, eventName string, event *github.WorkflowRunEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.WorkflowRunEvent) error {
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
			g.SetOnWorkflowRunEventCompleted(func(deliveryID string, eventName string, event *github.WorkflowRunEvent) error {
				return nil
			})
			g.SetOnWorkflowRunEventCompleted(tt.args.callbacks...)
			if len(g.onWorkflowRunEvent["completed"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onWorkflowRunEvent["completed"]), tt.want)
			}
		})
	}
}

func TestHandleWorkflowRunEventCompleted(t *testing.T) {
	action := "completed"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.WorkflowRunEvent
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
				eventName:  "workflow_run",
				event:      &github.WorkflowRunEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "workflow_run",
				event:      &github.WorkflowRunEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "workflow_run",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "workflow_run",
				event:      &github.WorkflowRunEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "workflow_run",
				event:      &github.WorkflowRunEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "workflow_run",
				event:      &github.WorkflowRunEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnWorkflowRunEventCompleted(func(deliveryID string, eventName string, event *github.WorkflowRunEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleWorkflowRunEventCompleted(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleWorkflowRunEventCompleted() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
