package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"errors"
	"github.com/google/go-github/v43/github"
	"testing"
)

func TestOnMetaEventAny(t *testing.T) {
	type args struct {
		callbacks []MetaEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single MetaEventHandleFunc",
			args: args{
				[]MetaEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MetaEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple MetaEventHandleFuncs",
			args: args{
				[]MetaEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MetaEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.MetaEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnMetaEventAny(tt.args.callbacks...)
			if len(g.onMetaEvent["*"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onMetaEvent["*"]))
			}
		})
	}
}

func TestSetOnMetaEventAny(t *testing.T) {
	type args struct {
		callbacks []MetaEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single MetaEventHandleFunc",
			args: args{
				[]MetaEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MetaEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple MetaEventHandleFuncs",
			args: args{
				[]MetaEventHandleFunc{
					func(deliveryID string, eventName string, event *github.MetaEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.MetaEvent) error {
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
			g.SetOnMetaEventAny(func(deliveryID string, eventName string, event *github.MetaEvent) error {
				return nil
			})
			g.SetOnMetaEventAny(tt.args.callbacks...)
			if len(g.onMetaEvent["*"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onMetaEvent["*"]), tt.want)
			}
		})
	}
}

func TestHandleMetaEventAny(t *testing.T) {

	type args struct {
		deliveryID string
		eventName  string
		event      *github.MetaEvent
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
				eventName:  "meta",

				event: &github.MetaEvent{},

				fail: false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "meta",

				event: &github.MetaEvent{},

				fail: true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "meta",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnMetaEventAny(func(deliveryID string, eventName string, event *github.MetaEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleMetaEventAny(tt.args.deliveryID, tt.args.deliveryID, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("TestHandleMetaEventAny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
