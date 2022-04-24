package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"errors"
	"github.com/google/go-github/v43/github"
	"testing"
)

func TestOnTeamEventAny(t *testing.T) {
	type args struct {
		callbacks []TeamEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single TeamEventHandleFunc",
			args: args{
				[]TeamEventHandleFunc{
					func(deliveryID string, eventName string, event *github.TeamEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple TeamEventHandleFuncs",
			args: args{
				[]TeamEventHandleFunc{
					func(deliveryID string, eventName string, event *github.TeamEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.TeamEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnTeamEventAny(tt.args.callbacks...)
			if len(g.onTeamEvent["*"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onTeamEvent["*"]))
			}
		})
	}
}

func TestSetOnTeamEventAny(t *testing.T) {
	type args struct {
		callbacks []TeamEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single TeamEventHandleFunc",
			args: args{
				[]TeamEventHandleFunc{
					func(deliveryID string, eventName string, event *github.TeamEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple TeamEventHandleFuncs",
			args: args{
				[]TeamEventHandleFunc{
					func(deliveryID string, eventName string, event *github.TeamEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.TeamEvent) error {
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
			g.SetOnTeamEventAny(func(deliveryID string, eventName string, event *github.TeamEvent) error {
				return nil
			})
			g.SetOnTeamEventAny(tt.args.callbacks...)
			if len(g.onTeamEvent["*"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onTeamEvent["*"]), tt.want)
			}
		})
	}
}

func TestHandleTeamEventAny(t *testing.T) {

	action := "*"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.TeamEvent
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
				eventName:  "team",

				event: &github.TeamEvent{Action: &action},

				fail: false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "team",

				event: &github.TeamEvent{Action: &action},

				fail: true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "team",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnTeamEventAny(func(deliveryID string, eventName string, event *github.TeamEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleTeamEventAny(tt.args.deliveryID, tt.args.deliveryID, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("TestHandleTeamEventAny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnTeamEventCreated(t *testing.T) {
	type args struct {
		callbacks []TeamEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single TeamEventHandleFunc",
			args: args{
				callbacks: []TeamEventHandleFunc{
					func(deliveryID string, eventName string, event *github.TeamEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple TeamEventHandleFunc",
			args: args{
				callbacks: []TeamEventHandleFunc{
					func(deliveryID string, eventName string, event *github.TeamEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.TeamEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnTeamEventCreated(tt.args.callbacks...)
			if len(g.onTeamEvent["created"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onTeamEvent["created"]))
			}
		})
	}
}

func TestSetOnTeamEventCreated(t *testing.T) {
	type args struct {
		callbacks []TeamEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single TeamEventHandleFunc",
			args: args{
				[]TeamEventHandleFunc{
					func(deliveryID string, eventName string, event *github.TeamEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple TeamEventHandleFuncs",
			args: args{
				[]TeamEventHandleFunc{
					func(deliveryID string, eventName string, event *github.TeamEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.TeamEvent) error {
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
			g.SetOnTeamEventCreated(func(deliveryID string, eventName string, event *github.TeamEvent) error {
				return nil
			})
			g.SetOnTeamEventCreated(tt.args.callbacks...)
			if len(g.onTeamEvent["created"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onTeamEvent["created"]), tt.want)
			}
		})
	}
}

func TestHandleTeamEventCreated(t *testing.T) {
	action := "created"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.TeamEvent
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
				eventName:  "team",
				event:      &github.TeamEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "team",
				event:      &github.TeamEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "team",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "team",
				event:      &github.TeamEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "team",
				event:      &github.TeamEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "team",
				event:      &github.TeamEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnTeamEventCreated(func(deliveryID string, eventName string, event *github.TeamEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleTeamEventCreated(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleTeamEventCreated() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnTeamEventDeleted(t *testing.T) {
	type args struct {
		callbacks []TeamEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single TeamEventHandleFunc",
			args: args{
				callbacks: []TeamEventHandleFunc{
					func(deliveryID string, eventName string, event *github.TeamEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple TeamEventHandleFunc",
			args: args{
				callbacks: []TeamEventHandleFunc{
					func(deliveryID string, eventName string, event *github.TeamEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.TeamEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnTeamEventDeleted(tt.args.callbacks...)
			if len(g.onTeamEvent["deleted"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onTeamEvent["deleted"]))
			}
		})
	}
}

func TestSetOnTeamEventDeleted(t *testing.T) {
	type args struct {
		callbacks []TeamEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single TeamEventHandleFunc",
			args: args{
				[]TeamEventHandleFunc{
					func(deliveryID string, eventName string, event *github.TeamEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple TeamEventHandleFuncs",
			args: args{
				[]TeamEventHandleFunc{
					func(deliveryID string, eventName string, event *github.TeamEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.TeamEvent) error {
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
			g.SetOnTeamEventDeleted(func(deliveryID string, eventName string, event *github.TeamEvent) error {
				return nil
			})
			g.SetOnTeamEventDeleted(tt.args.callbacks...)
			if len(g.onTeamEvent["deleted"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onTeamEvent["deleted"]), tt.want)
			}
		})
	}
}

func TestHandleTeamEventDeleted(t *testing.T) {
	action := "deleted"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.TeamEvent
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
				eventName:  "team",
				event:      &github.TeamEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "team",
				event:      &github.TeamEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "team",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "team",
				event:      &github.TeamEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "team",
				event:      &github.TeamEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "team",
				event:      &github.TeamEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnTeamEventDeleted(func(deliveryID string, eventName string, event *github.TeamEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleTeamEventDeleted(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleTeamEventDeleted() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnTeamEventEdited(t *testing.T) {
	type args struct {
		callbacks []TeamEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single TeamEventHandleFunc",
			args: args{
				callbacks: []TeamEventHandleFunc{
					func(deliveryID string, eventName string, event *github.TeamEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple TeamEventHandleFunc",
			args: args{
				callbacks: []TeamEventHandleFunc{
					func(deliveryID string, eventName string, event *github.TeamEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.TeamEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnTeamEventEdited(tt.args.callbacks...)
			if len(g.onTeamEvent["edited"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onTeamEvent["edited"]))
			}
		})
	}
}

func TestSetOnTeamEventEdited(t *testing.T) {
	type args struct {
		callbacks []TeamEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single TeamEventHandleFunc",
			args: args{
				[]TeamEventHandleFunc{
					func(deliveryID string, eventName string, event *github.TeamEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple TeamEventHandleFuncs",
			args: args{
				[]TeamEventHandleFunc{
					func(deliveryID string, eventName string, event *github.TeamEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.TeamEvent) error {
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
			g.SetOnTeamEventEdited(func(deliveryID string, eventName string, event *github.TeamEvent) error {
				return nil
			})
			g.SetOnTeamEventEdited(tt.args.callbacks...)
			if len(g.onTeamEvent["edited"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onTeamEvent["edited"]), tt.want)
			}
		})
	}
}

func TestHandleTeamEventEdited(t *testing.T) {
	action := "edited"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.TeamEvent
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
				eventName:  "team",
				event:      &github.TeamEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "team",
				event:      &github.TeamEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "team",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "team",
				event:      &github.TeamEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "team",
				event:      &github.TeamEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "team",
				event:      &github.TeamEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnTeamEventEdited(func(deliveryID string, eventName string, event *github.TeamEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleTeamEventEdited(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleTeamEventEdited() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnTeamEventAddedToRepository(t *testing.T) {
	type args struct {
		callbacks []TeamEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single TeamEventHandleFunc",
			args: args{
				callbacks: []TeamEventHandleFunc{
					func(deliveryID string, eventName string, event *github.TeamEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple TeamEventHandleFunc",
			args: args{
				callbacks: []TeamEventHandleFunc{
					func(deliveryID string, eventName string, event *github.TeamEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.TeamEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnTeamEventAddedToRepository(tt.args.callbacks...)
			if len(g.onTeamEvent["added_to_repository"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onTeamEvent["added_to_repository"]))
			}
		})
	}
}

func TestSetOnTeamEventAddedToRepository(t *testing.T) {
	type args struct {
		callbacks []TeamEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single TeamEventHandleFunc",
			args: args{
				[]TeamEventHandleFunc{
					func(deliveryID string, eventName string, event *github.TeamEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple TeamEventHandleFuncs",
			args: args{
				[]TeamEventHandleFunc{
					func(deliveryID string, eventName string, event *github.TeamEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.TeamEvent) error {
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
			g.SetOnTeamEventAddedToRepository(func(deliveryID string, eventName string, event *github.TeamEvent) error {
				return nil
			})
			g.SetOnTeamEventAddedToRepository(tt.args.callbacks...)
			if len(g.onTeamEvent["added_to_repository"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onTeamEvent["added_to_repository"]), tt.want)
			}
		})
	}
}

func TestHandleTeamEventAddedToRepository(t *testing.T) {
	action := "added_to_repository"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.TeamEvent
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
				eventName:  "team",
				event:      &github.TeamEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "team",
				event:      &github.TeamEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "team",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "team",
				event:      &github.TeamEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "team",
				event:      &github.TeamEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "team",
				event:      &github.TeamEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnTeamEventAddedToRepository(func(deliveryID string, eventName string, event *github.TeamEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleTeamEventAddedToRepository(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleTeamEventAddedToRepository() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnTeamEventRemovedFromRepository(t *testing.T) {
	type args struct {
		callbacks []TeamEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single TeamEventHandleFunc",
			args: args{
				callbacks: []TeamEventHandleFunc{
					func(deliveryID string, eventName string, event *github.TeamEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple TeamEventHandleFunc",
			args: args{
				callbacks: []TeamEventHandleFunc{
					func(deliveryID string, eventName string, event *github.TeamEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.TeamEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnTeamEventRemovedFromRepository(tt.args.callbacks...)
			if len(g.onTeamEvent["removed_from_repository"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onTeamEvent["removed_from_repository"]))
			}
		})
	}
}

func TestSetOnTeamEventRemovedFromRepository(t *testing.T) {
	type args struct {
		callbacks []TeamEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single TeamEventHandleFunc",
			args: args{
				[]TeamEventHandleFunc{
					func(deliveryID string, eventName string, event *github.TeamEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple TeamEventHandleFuncs",
			args: args{
				[]TeamEventHandleFunc{
					func(deliveryID string, eventName string, event *github.TeamEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.TeamEvent) error {
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
			g.SetOnTeamEventRemovedFromRepository(func(deliveryID string, eventName string, event *github.TeamEvent) error {
				return nil
			})
			g.SetOnTeamEventRemovedFromRepository(tt.args.callbacks...)
			if len(g.onTeamEvent["removed_from_repository"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onTeamEvent["removed_from_repository"]), tt.want)
			}
		})
	}
}

func TestHandleTeamEventRemovedFromRepository(t *testing.T) {
	action := "removed_from_repository"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.TeamEvent
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
				eventName:  "team",
				event:      &github.TeamEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "team",
				event:      &github.TeamEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "team",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "team",
				event:      &github.TeamEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "team",
				event:      &github.TeamEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "team",
				event:      &github.TeamEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnTeamEventRemovedFromRepository(func(deliveryID string, eventName string, event *github.TeamEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleTeamEventRemovedFromRepository(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleTeamEventRemovedFromRepository() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
