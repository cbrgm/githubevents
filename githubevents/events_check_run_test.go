package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"errors"
	"github.com/google/go-github/v43/github"
	"testing"
)

func TestOnCheckRunEventAny(t *testing.T) {
	type args struct {
		callbacks []CheckRunEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single CheckRunEventHandleFunc",
			args: args{
				[]CheckRunEventHandleFunc{
					func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple CheckRunEventHandleFuncs",
			args: args{
				[]CheckRunEventHandleFunc{
					func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnCheckRunEventAny(tt.args.callbacks...)
			if len(g.onCheckRunEvent["*"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onCheckRunEvent["*"]))
			}
		})
	}
}

func TestSetOnCheckRunEventAny(t *testing.T) {
	type args struct {
		callbacks []CheckRunEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single CheckRunEventHandleFunc",
			args: args{
				[]CheckRunEventHandleFunc{
					func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple CheckRunEventHandleFuncs",
			args: args{
				[]CheckRunEventHandleFunc{
					func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
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
			g.SetOnCheckRunEventAny(func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
				return nil
			})
			g.SetOnCheckRunEventAny(tt.args.callbacks...)
			if len(g.onCheckRunEvent["*"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onCheckRunEvent["*"]), tt.want)
			}
		})
	}
}

func TestHandleCheckRunEventAny(t *testing.T) {

	action := "*"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.CheckRunEvent
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
				eventName:  "check_run",

				event: &github.CheckRunEvent{Action: &action},

				fail: false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "check_run",

				event: &github.CheckRunEvent{Action: &action},

				fail: true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "check_run",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnCheckRunEventAny(func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleCheckRunEventAny(tt.args.deliveryID, tt.args.deliveryID, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("TestHandleCheckRunEventAny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnCheckRunEventCreated(t *testing.T) {
	type args struct {
		callbacks []CheckRunEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single CheckRunEventHandleFunc",
			args: args{
				callbacks: []CheckRunEventHandleFunc{
					func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple CheckRunEventHandleFunc",
			args: args{
				callbacks: []CheckRunEventHandleFunc{
					func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnCheckRunEventCreated(tt.args.callbacks...)
			if len(g.onCheckRunEvent["created"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onCheckRunEvent["created"]))
			}
		})
	}
}

func TestSetOnCheckRunEventCreated(t *testing.T) {
	type args struct {
		callbacks []CheckRunEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single CheckRunEventHandleFunc",
			args: args{
				[]CheckRunEventHandleFunc{
					func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple CheckRunEventHandleFuncs",
			args: args{
				[]CheckRunEventHandleFunc{
					func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
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
			g.SetOnCheckRunEventCreated(func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
				return nil
			})
			g.SetOnCheckRunEventCreated(tt.args.callbacks...)
			if len(g.onCheckRunEvent["created"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onCheckRunEvent["created"]), tt.want)
			}
		})
	}
}

func TestHandleCheckRunEventCreated(t *testing.T) {
	action := "created"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.CheckRunEvent
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
				eventName:  "check_run",
				event:      &github.CheckRunEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "check_run",
				event:      &github.CheckRunEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "check_run",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "check_run",
				event:      &github.CheckRunEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "check_run",
				event:      &github.CheckRunEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "check_run",
				event:      &github.CheckRunEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnCheckRunEventCreated(func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleCheckRunEventCreated(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleCheckRunEventCreated() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnCheckRunEventCompleted(t *testing.T) {
	type args struct {
		callbacks []CheckRunEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single CheckRunEventHandleFunc",
			args: args{
				callbacks: []CheckRunEventHandleFunc{
					func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple CheckRunEventHandleFunc",
			args: args{
				callbacks: []CheckRunEventHandleFunc{
					func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnCheckRunEventCompleted(tt.args.callbacks...)
			if len(g.onCheckRunEvent["completed"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onCheckRunEvent["completed"]))
			}
		})
	}
}

func TestSetOnCheckRunEventCompleted(t *testing.T) {
	type args struct {
		callbacks []CheckRunEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single CheckRunEventHandleFunc",
			args: args{
				[]CheckRunEventHandleFunc{
					func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple CheckRunEventHandleFuncs",
			args: args{
				[]CheckRunEventHandleFunc{
					func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
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
			g.SetOnCheckRunEventCompleted(func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
				return nil
			})
			g.SetOnCheckRunEventCompleted(tt.args.callbacks...)
			if len(g.onCheckRunEvent["completed"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onCheckRunEvent["completed"]), tt.want)
			}
		})
	}
}

func TestHandleCheckRunEventCompleted(t *testing.T) {
	action := "completed"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.CheckRunEvent
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
				eventName:  "check_run",
				event:      &github.CheckRunEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "check_run",
				event:      &github.CheckRunEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "check_run",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "check_run",
				event:      &github.CheckRunEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "check_run",
				event:      &github.CheckRunEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "check_run",
				event:      &github.CheckRunEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnCheckRunEventCompleted(func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleCheckRunEventCompleted(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleCheckRunEventCompleted() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnCheckRunEventReRequested(t *testing.T) {
	type args struct {
		callbacks []CheckRunEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single CheckRunEventHandleFunc",
			args: args{
				callbacks: []CheckRunEventHandleFunc{
					func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple CheckRunEventHandleFunc",
			args: args{
				callbacks: []CheckRunEventHandleFunc{
					func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnCheckRunEventReRequested(tt.args.callbacks...)
			if len(g.onCheckRunEvent["rerequested"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onCheckRunEvent["rerequested"]))
			}
		})
	}
}

func TestSetOnCheckRunEventReRequested(t *testing.T) {
	type args struct {
		callbacks []CheckRunEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single CheckRunEventHandleFunc",
			args: args{
				[]CheckRunEventHandleFunc{
					func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple CheckRunEventHandleFuncs",
			args: args{
				[]CheckRunEventHandleFunc{
					func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
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
			g.SetOnCheckRunEventReRequested(func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
				return nil
			})
			g.SetOnCheckRunEventReRequested(tt.args.callbacks...)
			if len(g.onCheckRunEvent["rerequested"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onCheckRunEvent["rerequested"]), tt.want)
			}
		})
	}
}

func TestHandleCheckRunEventReRequested(t *testing.T) {
	action := "rerequested"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.CheckRunEvent
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
				eventName:  "check_run",
				event:      &github.CheckRunEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "check_run",
				event:      &github.CheckRunEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "check_run",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "check_run",
				event:      &github.CheckRunEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "check_run",
				event:      &github.CheckRunEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "check_run",
				event:      &github.CheckRunEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnCheckRunEventReRequested(func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleCheckRunEventReRequested(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleCheckRunEventReRequested() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnCheckRunEventRequestAction(t *testing.T) {
	type args struct {
		callbacks []CheckRunEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single CheckRunEventHandleFunc",
			args: args{
				callbacks: []CheckRunEventHandleFunc{
					func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple CheckRunEventHandleFunc",
			args: args{
				callbacks: []CheckRunEventHandleFunc{
					func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnCheckRunEventRequestAction(tt.args.callbacks...)
			if len(g.onCheckRunEvent["requested_action"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onCheckRunEvent["requested_action"]))
			}
		})
	}
}

func TestSetOnCheckRunEventRequestAction(t *testing.T) {
	type args struct {
		callbacks []CheckRunEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single CheckRunEventHandleFunc",
			args: args{
				[]CheckRunEventHandleFunc{
					func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple CheckRunEventHandleFuncs",
			args: args{
				[]CheckRunEventHandleFunc{
					func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
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
			g.SetOnCheckRunEventRequestAction(func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
				return nil
			})
			g.SetOnCheckRunEventRequestAction(tt.args.callbacks...)
			if len(g.onCheckRunEvent["requested_action"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onCheckRunEvent["requested_action"]), tt.want)
			}
		})
	}
}

func TestHandleCheckRunEventRequestAction(t *testing.T) {
	action := "requested_action"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.CheckRunEvent
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
				eventName:  "check_run",
				event:      &github.CheckRunEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "check_run",
				event:      &github.CheckRunEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "check_run",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "check_run",
				event:      &github.CheckRunEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "check_run",
				event:      &github.CheckRunEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "check_run",
				event:      &github.CheckRunEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnCheckRunEventRequestAction(func(deliveryID string, eventName string, event *github.CheckRunEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleCheckRunEventRequestAction(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleCheckRunEventRequestAction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
