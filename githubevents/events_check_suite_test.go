package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"errors"
	"github.com/google/go-github/v43/github"
	"testing"
)

func TestOnCheckSuiteEventAny(t *testing.T) {
	type args struct {
		callbacks []CheckSuiteEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single CheckSuiteEventHandleFunc",
			args: args{
				[]CheckSuiteEventHandleFunc{
					func(deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple CheckSuiteEventHandleFuncs",
			args: args{
				[]CheckSuiteEventHandleFunc{
					func(deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnCheckSuiteEventAny(tt.args.callbacks...)
			if len(g.onCheckSuiteEvent["*"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onCheckSuiteEvent["*"]))
			}
		})
	}
}

func TestSetOnCheckSuiteEventAny(t *testing.T) {
	type args struct {
		callbacks []CheckSuiteEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single CheckSuiteEventHandleFunc",
			args: args{
				[]CheckSuiteEventHandleFunc{
					func(deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple CheckSuiteEventHandleFuncs",
			args: args{
				[]CheckSuiteEventHandleFunc{
					func(deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
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
			g.SetOnCheckSuiteEventAny(func(deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
				return nil
			})
			g.SetOnCheckSuiteEventAny(tt.args.callbacks...)
			if len(g.onCheckSuiteEvent["*"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onCheckSuiteEvent["*"]), tt.want)
			}
		})
	}
}

func TestHandleCheckSuiteEventAny(t *testing.T) {

	action := "*"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.CheckSuiteEvent
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
				eventName:  "check_suite",

				event: &github.CheckSuiteEvent{Action: &action},

				fail: false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "check_suite",

				event: &github.CheckSuiteEvent{Action: &action},

				fail: true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "check_suite",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnCheckSuiteEventAny(func(deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleCheckSuiteEventAny(tt.args.deliveryID, tt.args.deliveryID, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("TestHandleCheckSuiteEventAny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnCheckSuiteEventCompleted(t *testing.T) {
	type args struct {
		callbacks []CheckSuiteEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single CheckSuiteEventHandleFunc",
			args: args{
				callbacks: []CheckSuiteEventHandleFunc{
					func(deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple CheckSuiteEventHandleFunc",
			args: args{
				callbacks: []CheckSuiteEventHandleFunc{
					func(deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnCheckSuiteEventCompleted(tt.args.callbacks...)
			if len(g.onCheckSuiteEvent["completed"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onCheckSuiteEvent["completed"]))
			}
		})
	}
}

func TestSetOnCheckSuiteEventCompleted(t *testing.T) {
	type args struct {
		callbacks []CheckSuiteEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single CheckSuiteEventHandleFunc",
			args: args{
				[]CheckSuiteEventHandleFunc{
					func(deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple CheckSuiteEventHandleFuncs",
			args: args{
				[]CheckSuiteEventHandleFunc{
					func(deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
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
			g.SetOnCheckSuiteEventCompleted(func(deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
				return nil
			})
			g.SetOnCheckSuiteEventCompleted(tt.args.callbacks...)
			if len(g.onCheckSuiteEvent["completed"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onCheckSuiteEvent["completed"]), tt.want)
			}
		})
	}
}

func TestHandleCheckSuiteEventCompleted(t *testing.T) {
	action := "completed"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.CheckSuiteEvent
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
				eventName:  "check_suite",
				event:      &github.CheckSuiteEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "check_suite",
				event:      &github.CheckSuiteEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "check_suite",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "check_suite",
				event:      &github.CheckSuiteEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "check_suite",
				event:      &github.CheckSuiteEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "check_suite",
				event:      &github.CheckSuiteEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnCheckSuiteEventCompleted(func(deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleCheckSuiteEventCompleted(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleCheckSuiteEventCompleted() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnCheckSuiteEventRequested(t *testing.T) {
	type args struct {
		callbacks []CheckSuiteEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single CheckSuiteEventHandleFunc",
			args: args{
				callbacks: []CheckSuiteEventHandleFunc{
					func(deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple CheckSuiteEventHandleFunc",
			args: args{
				callbacks: []CheckSuiteEventHandleFunc{
					func(deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnCheckSuiteEventRequested(tt.args.callbacks...)
			if len(g.onCheckSuiteEvent["requested"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onCheckSuiteEvent["requested"]))
			}
		})
	}
}

func TestSetOnCheckSuiteEventRequested(t *testing.T) {
	type args struct {
		callbacks []CheckSuiteEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single CheckSuiteEventHandleFunc",
			args: args{
				[]CheckSuiteEventHandleFunc{
					func(deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple CheckSuiteEventHandleFuncs",
			args: args{
				[]CheckSuiteEventHandleFunc{
					func(deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
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
			g.SetOnCheckSuiteEventRequested(func(deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
				return nil
			})
			g.SetOnCheckSuiteEventRequested(tt.args.callbacks...)
			if len(g.onCheckSuiteEvent["requested"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onCheckSuiteEvent["requested"]), tt.want)
			}
		})
	}
}

func TestHandleCheckSuiteEventRequested(t *testing.T) {
	action := "requested"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.CheckSuiteEvent
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
				eventName:  "check_suite",
				event:      &github.CheckSuiteEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "check_suite",
				event:      &github.CheckSuiteEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "check_suite",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "check_suite",
				event:      &github.CheckSuiteEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "check_suite",
				event:      &github.CheckSuiteEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "check_suite",
				event:      &github.CheckSuiteEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnCheckSuiteEventRequested(func(deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleCheckSuiteEventRequested(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleCheckSuiteEventRequested() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnCheckSuiteEventReRequested(t *testing.T) {
	type args struct {
		callbacks []CheckSuiteEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single CheckSuiteEventHandleFunc",
			args: args{
				callbacks: []CheckSuiteEventHandleFunc{
					func(deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple CheckSuiteEventHandleFunc",
			args: args{
				callbacks: []CheckSuiteEventHandleFunc{
					func(deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnCheckSuiteEventReRequested(tt.args.callbacks...)
			if len(g.onCheckSuiteEvent["rerequested"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onCheckSuiteEvent["rerequested"]))
			}
		})
	}
}

func TestSetOnCheckSuiteEventReRequested(t *testing.T) {
	type args struct {
		callbacks []CheckSuiteEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single CheckSuiteEventHandleFunc",
			args: args{
				[]CheckSuiteEventHandleFunc{
					func(deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple CheckSuiteEventHandleFuncs",
			args: args{
				[]CheckSuiteEventHandleFunc{
					func(deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
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
			g.SetOnCheckSuiteEventReRequested(func(deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
				return nil
			})
			g.SetOnCheckSuiteEventReRequested(tt.args.callbacks...)
			if len(g.onCheckSuiteEvent["rerequested"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onCheckSuiteEvent["rerequested"]), tt.want)
			}
		})
	}
}

func TestHandleCheckSuiteEventReRequested(t *testing.T) {
	action := "rerequested"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.CheckSuiteEvent
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
				eventName:  "check_suite",
				event:      &github.CheckSuiteEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "check_suite",
				event:      &github.CheckSuiteEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "check_suite",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "check_suite",
				event:      &github.CheckSuiteEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "check_suite",
				event:      &github.CheckSuiteEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "check_suite",
				event:      &github.CheckSuiteEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnCheckSuiteEventReRequested(func(deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleCheckSuiteEventReRequested(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleCheckSuiteEventReRequested() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
