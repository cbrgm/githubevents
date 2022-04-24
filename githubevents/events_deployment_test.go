package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"errors"
	"github.com/google/go-github/v43/github"
	"testing"
)

func TestOnDeploymentEventAny(t *testing.T) {
	type args struct {
		callbacks []DeploymentEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single DeploymentEventHandleFunc",
			args: args{
				[]DeploymentEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DeploymentEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple DeploymentEventHandleFuncs",
			args: args{
				[]DeploymentEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DeploymentEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.DeploymentEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnDeploymentEventAny(tt.args.callbacks...)
			if len(g.onDeploymentEvent["*"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onDeploymentEvent["*"]))
			}
		})
	}
}

func TestSetOnDeploymentEventAny(t *testing.T) {
	type args struct {
		callbacks []DeploymentEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single DeploymentEventHandleFunc",
			args: args{
				[]DeploymentEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DeploymentEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple DeploymentEventHandleFuncs",
			args: args{
				[]DeploymentEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DeploymentEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.DeploymentEvent) error {
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
			g.SetOnDeploymentEventAny(func(deliveryID string, eventName string, event *github.DeploymentEvent) error {
				return nil
			})
			g.SetOnDeploymentEventAny(tt.args.callbacks...)
			if len(g.onDeploymentEvent["*"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onDeploymentEvent["*"]), tt.want)
			}
		})
	}
}

func TestHandleDeploymentEventAny(t *testing.T) {

	type args struct {
		deliveryID string
		eventName  string
		event      *github.DeploymentEvent
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
				eventName:  "deployment",

				event: &github.DeploymentEvent{},

				fail: false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "deployment",

				event: &github.DeploymentEvent{},

				fail: true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "deployment",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnDeploymentEventAny(func(deliveryID string, eventName string, event *github.DeploymentEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleDeploymentEventAny(tt.args.deliveryID, tt.args.deliveryID, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("TestHandleDeploymentEventAny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
