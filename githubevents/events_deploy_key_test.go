package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"errors"
	"github.com/google/go-github/v43/github"
	"testing"
)

func TestOnDeployKeyEventAny(t *testing.T) {
	type args struct {
		callbacks []DeployKeyEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single DeployKeyEventHandleFunc",
			args: args{
				[]DeployKeyEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DeployKeyEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple DeployKeyEventHandleFuncs",
			args: args{
				[]DeployKeyEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DeployKeyEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.DeployKeyEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnDeployKeyEventAny(tt.args.callbacks...)
			if len(g.onDeployKeyEvent["*"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onDeployKeyEvent["*"]))
			}
		})
	}
}

func TestSetOnDeployKeyEventAny(t *testing.T) {
	type args struct {
		callbacks []DeployKeyEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single DeployKeyEventHandleFunc",
			args: args{
				[]DeployKeyEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DeployKeyEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple DeployKeyEventHandleFuncs",
			args: args{
				[]DeployKeyEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DeployKeyEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.DeployKeyEvent) error {
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
			g.SetOnDeployKeyEventAny(func(deliveryID string, eventName string, event *github.DeployKeyEvent) error {
				return nil
			})
			g.SetOnDeployKeyEventAny(tt.args.callbacks...)
			if len(g.onDeployKeyEvent["*"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onDeployKeyEvent["*"]), tt.want)
			}
		})
	}
}

func TestHandleDeployKeyEventAny(t *testing.T) {

	action := "*"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.DeployKeyEvent
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
				eventName:  "deploy_key",

				event: &github.DeployKeyEvent{Action: &action},

				fail: false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "deploy_key",

				event: &github.DeployKeyEvent{Action: &action},

				fail: true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "deploy_key",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnDeployKeyEventAny(func(deliveryID string, eventName string, event *github.DeployKeyEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleDeployKeyEventAny(tt.args.deliveryID, tt.args.deliveryID, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("TestHandleDeployKeyEventAny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnDeployKeyEventCreated(t *testing.T) {
	type args struct {
		callbacks []DeployKeyEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single DeployKeyEventHandleFunc",
			args: args{
				callbacks: []DeployKeyEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DeployKeyEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple DeployKeyEventHandleFunc",
			args: args{
				callbacks: []DeployKeyEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DeployKeyEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.DeployKeyEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnDeployKeyEventCreated(tt.args.callbacks...)
			if len(g.onDeployKeyEvent["created"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onDeployKeyEvent["created"]))
			}
		})
	}
}

func TestSetOnDeployKeyEventCreated(t *testing.T) {
	type args struct {
		callbacks []DeployKeyEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single DeployKeyEventHandleFunc",
			args: args{
				[]DeployKeyEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DeployKeyEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple DeployKeyEventHandleFuncs",
			args: args{
				[]DeployKeyEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DeployKeyEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.DeployKeyEvent) error {
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
			g.SetOnDeployKeyEventCreated(func(deliveryID string, eventName string, event *github.DeployKeyEvent) error {
				return nil
			})
			g.SetOnDeployKeyEventCreated(tt.args.callbacks...)
			if len(g.onDeployKeyEvent["created"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onDeployKeyEvent["created"]), tt.want)
			}
		})
	}
}

func TestHandleDeployKeyEventCreated(t *testing.T) {
	action := "created"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.DeployKeyEvent
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
				eventName:  "deploy_key",
				event:      &github.DeployKeyEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "deploy_key",
				event:      &github.DeployKeyEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "deploy_key",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "deploy_key",
				event:      &github.DeployKeyEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "deploy_key",
				event:      &github.DeployKeyEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "deploy_key",
				event:      &github.DeployKeyEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnDeployKeyEventCreated(func(deliveryID string, eventName string, event *github.DeployKeyEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleDeployKeyEventCreated(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleDeployKeyEventCreated() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnDeployKeyEventDeleted(t *testing.T) {
	type args struct {
		callbacks []DeployKeyEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single DeployKeyEventHandleFunc",
			args: args{
				callbacks: []DeployKeyEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DeployKeyEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple DeployKeyEventHandleFunc",
			args: args{
				callbacks: []DeployKeyEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DeployKeyEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.DeployKeyEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnDeployKeyEventDeleted(tt.args.callbacks...)
			if len(g.onDeployKeyEvent["deleted"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onDeployKeyEvent["deleted"]))
			}
		})
	}
}

func TestSetOnDeployKeyEventDeleted(t *testing.T) {
	type args struct {
		callbacks []DeployKeyEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single DeployKeyEventHandleFunc",
			args: args{
				[]DeployKeyEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DeployKeyEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple DeployKeyEventHandleFuncs",
			args: args{
				[]DeployKeyEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DeployKeyEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.DeployKeyEvent) error {
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
			g.SetOnDeployKeyEventDeleted(func(deliveryID string, eventName string, event *github.DeployKeyEvent) error {
				return nil
			})
			g.SetOnDeployKeyEventDeleted(tt.args.callbacks...)
			if len(g.onDeployKeyEvent["deleted"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onDeployKeyEvent["deleted"]), tt.want)
			}
		})
	}
}

func TestHandleDeployKeyEventDeleted(t *testing.T) {
	action := "deleted"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.DeployKeyEvent
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
				eventName:  "deploy_key",
				event:      &github.DeployKeyEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "deploy_key",
				event:      &github.DeployKeyEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "deploy_key",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "deploy_key",
				event:      &github.DeployKeyEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "deploy_key",
				event:      &github.DeployKeyEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "deploy_key",
				event:      &github.DeployKeyEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnDeployKeyEventDeleted(func(deliveryID string, eventName string, event *github.DeployKeyEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleDeployKeyEventDeleted(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleDeployKeyEventDeleted() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
