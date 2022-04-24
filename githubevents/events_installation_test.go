package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"errors"
	"github.com/google/go-github/v43/github"
	"testing"
)

func TestOnInstallationEventAny(t *testing.T) {
	type args struct {
		callbacks []InstallationEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single InstallationEventHandleFunc",
			args: args{
				[]InstallationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.InstallationEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple InstallationEventHandleFuncs",
			args: args{
				[]InstallationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.InstallationEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.InstallationEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnInstallationEventAny(tt.args.callbacks...)
			if len(g.onInstallationEvent["*"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onInstallationEvent["*"]))
			}
		})
	}
}

func TestSetOnInstallationEventAny(t *testing.T) {
	type args struct {
		callbacks []InstallationEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single InstallationEventHandleFunc",
			args: args{
				[]InstallationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.InstallationEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple InstallationEventHandleFuncs",
			args: args{
				[]InstallationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.InstallationEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.InstallationEvent) error {
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
			g.SetOnInstallationEventAny(func(deliveryID string, eventName string, event *github.InstallationEvent) error {
				return nil
			})
			g.SetOnInstallationEventAny(tt.args.callbacks...)
			if len(g.onInstallationEvent["*"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onInstallationEvent["*"]), tt.want)
			}
		})
	}
}

func TestHandleInstallationEventAny(t *testing.T) {

	action := "*"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.InstallationEvent
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
				eventName:  "installation",

				event: &github.InstallationEvent{Action: &action},

				fail: false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "installation",

				event: &github.InstallationEvent{Action: &action},

				fail: true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "installation",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnInstallationEventAny(func(deliveryID string, eventName string, event *github.InstallationEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleInstallationEventAny(tt.args.deliveryID, tt.args.deliveryID, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("TestHandleInstallationEventAny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnInstallationEventCreated(t *testing.T) {
	type args struct {
		callbacks []InstallationEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single InstallationEventHandleFunc",
			args: args{
				callbacks: []InstallationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.InstallationEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple InstallationEventHandleFunc",
			args: args{
				callbacks: []InstallationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.InstallationEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.InstallationEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnInstallationEventCreated(tt.args.callbacks...)
			if len(g.onInstallationEvent["created"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onInstallationEvent["created"]))
			}
		})
	}
}

func TestSetOnInstallationEventCreated(t *testing.T) {
	type args struct {
		callbacks []InstallationEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single InstallationEventHandleFunc",
			args: args{
				[]InstallationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.InstallationEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple InstallationEventHandleFuncs",
			args: args{
				[]InstallationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.InstallationEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.InstallationEvent) error {
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
			g.SetOnInstallationEventCreated(func(deliveryID string, eventName string, event *github.InstallationEvent) error {
				return nil
			})
			g.SetOnInstallationEventCreated(tt.args.callbacks...)
			if len(g.onInstallationEvent["created"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onInstallationEvent["created"]), tt.want)
			}
		})
	}
}

func TestHandleInstallationEventCreated(t *testing.T) {
	action := "created"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.InstallationEvent
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
				eventName:  "installation",
				event:      &github.InstallationEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "installation",
				event:      &github.InstallationEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "installation",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "installation",
				event:      &github.InstallationEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "installation",
				event:      &github.InstallationEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "installation",
				event:      &github.InstallationEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnInstallationEventCreated(func(deliveryID string, eventName string, event *github.InstallationEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleInstallationEventCreated(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleInstallationEventCreated() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnInstallationEventDeleted(t *testing.T) {
	type args struct {
		callbacks []InstallationEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single InstallationEventHandleFunc",
			args: args{
				callbacks: []InstallationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.InstallationEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple InstallationEventHandleFunc",
			args: args{
				callbacks: []InstallationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.InstallationEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.InstallationEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnInstallationEventDeleted(tt.args.callbacks...)
			if len(g.onInstallationEvent["deleted"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onInstallationEvent["deleted"]))
			}
		})
	}
}

func TestSetOnInstallationEventDeleted(t *testing.T) {
	type args struct {
		callbacks []InstallationEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single InstallationEventHandleFunc",
			args: args{
				[]InstallationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.InstallationEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple InstallationEventHandleFuncs",
			args: args{
				[]InstallationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.InstallationEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.InstallationEvent) error {
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
			g.SetOnInstallationEventDeleted(func(deliveryID string, eventName string, event *github.InstallationEvent) error {
				return nil
			})
			g.SetOnInstallationEventDeleted(tt.args.callbacks...)
			if len(g.onInstallationEvent["deleted"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onInstallationEvent["deleted"]), tt.want)
			}
		})
	}
}

func TestHandleInstallationEventDeleted(t *testing.T) {
	action := "deleted"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.InstallationEvent
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
				eventName:  "installation",
				event:      &github.InstallationEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "installation",
				event:      &github.InstallationEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "installation",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "installation",
				event:      &github.InstallationEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "installation",
				event:      &github.InstallationEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "installation",
				event:      &github.InstallationEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnInstallationEventDeleted(func(deliveryID string, eventName string, event *github.InstallationEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleInstallationEventDeleted(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleInstallationEventDeleted() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnInstallationEventEventSuspend(t *testing.T) {
	type args struct {
		callbacks []InstallationEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single InstallationEventHandleFunc",
			args: args{
				callbacks: []InstallationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.InstallationEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple InstallationEventHandleFunc",
			args: args{
				callbacks: []InstallationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.InstallationEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.InstallationEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnInstallationEventEventSuspend(tt.args.callbacks...)
			if len(g.onInstallationEvent["suspend"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onInstallationEvent["suspend"]))
			}
		})
	}
}

func TestSetOnInstallationEventEventSuspend(t *testing.T) {
	type args struct {
		callbacks []InstallationEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single InstallationEventHandleFunc",
			args: args{
				[]InstallationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.InstallationEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple InstallationEventHandleFuncs",
			args: args{
				[]InstallationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.InstallationEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.InstallationEvent) error {
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
			g.SetOnInstallationEventEventSuspend(func(deliveryID string, eventName string, event *github.InstallationEvent) error {
				return nil
			})
			g.SetOnInstallationEventEventSuspend(tt.args.callbacks...)
			if len(g.onInstallationEvent["suspend"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onInstallationEvent["suspend"]), tt.want)
			}
		})
	}
}

func TestHandleInstallationEventEventSuspend(t *testing.T) {
	action := "suspend"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.InstallationEvent
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
				eventName:  "installation",
				event:      &github.InstallationEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "installation",
				event:      &github.InstallationEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "installation",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "installation",
				event:      &github.InstallationEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "installation",
				event:      &github.InstallationEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "installation",
				event:      &github.InstallationEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnInstallationEventEventSuspend(func(deliveryID string, eventName string, event *github.InstallationEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleInstallationEventEventSuspend(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleInstallationEventEventSuspend() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnInstallationEventEventUnsuspend(t *testing.T) {
	type args struct {
		callbacks []InstallationEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single InstallationEventHandleFunc",
			args: args{
				callbacks: []InstallationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.InstallationEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple InstallationEventHandleFunc",
			args: args{
				callbacks: []InstallationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.InstallationEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.InstallationEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnInstallationEventEventUnsuspend(tt.args.callbacks...)
			if len(g.onInstallationEvent["unsuspend"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onInstallationEvent["unsuspend"]))
			}
		})
	}
}

func TestSetOnInstallationEventEventUnsuspend(t *testing.T) {
	type args struct {
		callbacks []InstallationEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single InstallationEventHandleFunc",
			args: args{
				[]InstallationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.InstallationEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple InstallationEventHandleFuncs",
			args: args{
				[]InstallationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.InstallationEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.InstallationEvent) error {
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
			g.SetOnInstallationEventEventUnsuspend(func(deliveryID string, eventName string, event *github.InstallationEvent) error {
				return nil
			})
			g.SetOnInstallationEventEventUnsuspend(tt.args.callbacks...)
			if len(g.onInstallationEvent["unsuspend"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onInstallationEvent["unsuspend"]), tt.want)
			}
		})
	}
}

func TestHandleInstallationEventEventUnsuspend(t *testing.T) {
	action := "unsuspend"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.InstallationEvent
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
				eventName:  "installation",
				event:      &github.InstallationEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "installation",
				event:      &github.InstallationEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "installation",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "installation",
				event:      &github.InstallationEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "installation",
				event:      &github.InstallationEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "installation",
				event:      &github.InstallationEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnInstallationEventEventUnsuspend(func(deliveryID string, eventName string, event *github.InstallationEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleInstallationEventEventUnsuspend(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleInstallationEventEventUnsuspend() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnInstallationEventNewPermissionsAccepted(t *testing.T) {
	type args struct {
		callbacks []InstallationEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single InstallationEventHandleFunc",
			args: args{
				callbacks: []InstallationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.InstallationEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple InstallationEventHandleFunc",
			args: args{
				callbacks: []InstallationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.InstallationEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.InstallationEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnInstallationEventNewPermissionsAccepted(tt.args.callbacks...)
			if len(g.onInstallationEvent["new_permissions_accepted"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onInstallationEvent["new_permissions_accepted"]))
			}
		})
	}
}

func TestSetOnInstallationEventNewPermissionsAccepted(t *testing.T) {
	type args struct {
		callbacks []InstallationEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single InstallationEventHandleFunc",
			args: args{
				[]InstallationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.InstallationEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple InstallationEventHandleFuncs",
			args: args{
				[]InstallationEventHandleFunc{
					func(deliveryID string, eventName string, event *github.InstallationEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.InstallationEvent) error {
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
			g.SetOnInstallationEventNewPermissionsAccepted(func(deliveryID string, eventName string, event *github.InstallationEvent) error {
				return nil
			})
			g.SetOnInstallationEventNewPermissionsAccepted(tt.args.callbacks...)
			if len(g.onInstallationEvent["new_permissions_accepted"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onInstallationEvent["new_permissions_accepted"]), tt.want)
			}
		})
	}
}

func TestHandleInstallationEventNewPermissionsAccepted(t *testing.T) {
	action := "new_permissions_accepted"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.InstallationEvent
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
				eventName:  "installation",
				event:      &github.InstallationEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "installation",
				event:      &github.InstallationEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "installation",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "installation",
				event:      &github.InstallationEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "installation",
				event:      &github.InstallationEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "installation",
				event:      &github.InstallationEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnInstallationEventNewPermissionsAccepted(func(deliveryID string, eventName string, event *github.InstallationEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleInstallationEventNewPermissionsAccepted(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleInstallationEventNewPermissionsAccepted() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
