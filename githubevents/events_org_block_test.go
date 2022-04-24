package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"errors"
	"github.com/google/go-github/v43/github"
	"testing"
)

func TestOnOrgBlockEventAny(t *testing.T) {
	type args struct {
		callbacks []OrgBlockEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single OrgBlockEventHandleFunc",
			args: args{
				[]OrgBlockEventHandleFunc{
					func(deliveryID string, eventName string, event *github.OrgBlockEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple OrgBlockEventHandleFuncs",
			args: args{
				[]OrgBlockEventHandleFunc{
					func(deliveryID string, eventName string, event *github.OrgBlockEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.OrgBlockEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnOrgBlockEventAny(tt.args.callbacks...)
			if len(g.onOrgBlockEvent["*"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onOrgBlockEvent["*"]))
			}
		})
	}
}

func TestSetOnOrgBlockEventAny(t *testing.T) {
	type args struct {
		callbacks []OrgBlockEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single OrgBlockEventHandleFunc",
			args: args{
				[]OrgBlockEventHandleFunc{
					func(deliveryID string, eventName string, event *github.OrgBlockEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple OrgBlockEventHandleFuncs",
			args: args{
				[]OrgBlockEventHandleFunc{
					func(deliveryID string, eventName string, event *github.OrgBlockEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.OrgBlockEvent) error {
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
			g.SetOnOrgBlockEventAny(func(deliveryID string, eventName string, event *github.OrgBlockEvent) error {
				return nil
			})
			g.SetOnOrgBlockEventAny(tt.args.callbacks...)
			if len(g.onOrgBlockEvent["*"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onOrgBlockEvent["*"]), tt.want)
			}
		})
	}
}

func TestHandleOrgBlockEventAny(t *testing.T) {

	action := "*"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.OrgBlockEvent
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
				eventName:  "org_block",

				event: &github.OrgBlockEvent{Action: &action},

				fail: false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "org_block",

				event: &github.OrgBlockEvent{Action: &action},

				fail: true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "org_block",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnOrgBlockEventAny(func(deliveryID string, eventName string, event *github.OrgBlockEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleOrgBlockEventAny(tt.args.deliveryID, tt.args.deliveryID, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("TestHandleOrgBlockEventAny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnOrgBlockEventBlocked(t *testing.T) {
	type args struct {
		callbacks []OrgBlockEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single OrgBlockEventHandleFunc",
			args: args{
				callbacks: []OrgBlockEventHandleFunc{
					func(deliveryID string, eventName string, event *github.OrgBlockEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple OrgBlockEventHandleFunc",
			args: args{
				callbacks: []OrgBlockEventHandleFunc{
					func(deliveryID string, eventName string, event *github.OrgBlockEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.OrgBlockEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnOrgBlockEventBlocked(tt.args.callbacks...)
			if len(g.onOrgBlockEvent["blocked"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onOrgBlockEvent["blocked"]))
			}
		})
	}
}

func TestSetOnOrgBlockEventBlocked(t *testing.T) {
	type args struct {
		callbacks []OrgBlockEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single OrgBlockEventHandleFunc",
			args: args{
				[]OrgBlockEventHandleFunc{
					func(deliveryID string, eventName string, event *github.OrgBlockEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple OrgBlockEventHandleFuncs",
			args: args{
				[]OrgBlockEventHandleFunc{
					func(deliveryID string, eventName string, event *github.OrgBlockEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.OrgBlockEvent) error {
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
			g.SetOnOrgBlockEventBlocked(func(deliveryID string, eventName string, event *github.OrgBlockEvent) error {
				return nil
			})
			g.SetOnOrgBlockEventBlocked(tt.args.callbacks...)
			if len(g.onOrgBlockEvent["blocked"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onOrgBlockEvent["blocked"]), tt.want)
			}
		})
	}
}

func TestHandleOrgBlockEventBlocked(t *testing.T) {
	action := "blocked"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.OrgBlockEvent
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
				eventName:  "org_block",
				event:      &github.OrgBlockEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "org_block",
				event:      &github.OrgBlockEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "org_block",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "org_block",
				event:      &github.OrgBlockEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "org_block",
				event:      &github.OrgBlockEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "org_block",
				event:      &github.OrgBlockEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnOrgBlockEventBlocked(func(deliveryID string, eventName string, event *github.OrgBlockEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleOrgBlockEventBlocked(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleOrgBlockEventBlocked() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnOrgBlockEventUnblocked(t *testing.T) {
	type args struct {
		callbacks []OrgBlockEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single OrgBlockEventHandleFunc",
			args: args{
				callbacks: []OrgBlockEventHandleFunc{
					func(deliveryID string, eventName string, event *github.OrgBlockEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple OrgBlockEventHandleFunc",
			args: args{
				callbacks: []OrgBlockEventHandleFunc{
					func(deliveryID string, eventName string, event *github.OrgBlockEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.OrgBlockEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnOrgBlockEventUnblocked(tt.args.callbacks...)
			if len(g.onOrgBlockEvent["unblocked"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onOrgBlockEvent["unblocked"]))
			}
		})
	}
}

func TestSetOnOrgBlockEventUnblocked(t *testing.T) {
	type args struct {
		callbacks []OrgBlockEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single OrgBlockEventHandleFunc",
			args: args{
				[]OrgBlockEventHandleFunc{
					func(deliveryID string, eventName string, event *github.OrgBlockEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple OrgBlockEventHandleFuncs",
			args: args{
				[]OrgBlockEventHandleFunc{
					func(deliveryID string, eventName string, event *github.OrgBlockEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.OrgBlockEvent) error {
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
			g.SetOnOrgBlockEventUnblocked(func(deliveryID string, eventName string, event *github.OrgBlockEvent) error {
				return nil
			})
			g.SetOnOrgBlockEventUnblocked(tt.args.callbacks...)
			if len(g.onOrgBlockEvent["unblocked"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onOrgBlockEvent["unblocked"]), tt.want)
			}
		})
	}
}

func TestHandleOrgBlockEventUnblocked(t *testing.T) {
	action := "unblocked"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.OrgBlockEvent
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
				eventName:  "org_block",
				event:      &github.OrgBlockEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "org_block",
				event:      &github.OrgBlockEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "org_block",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "org_block",
				event:      &github.OrgBlockEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "org_block",
				event:      &github.OrgBlockEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "org_block",
				event:      &github.OrgBlockEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnOrgBlockEventUnblocked(func(deliveryID string, eventName string, event *github.OrgBlockEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleOrgBlockEventUnblocked(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleOrgBlockEventUnblocked() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
