package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"errors"
	"github.com/google/go-github/v43/github"
	"testing"
)

func TestOnPageBuildEventAny(t *testing.T) {
	type args struct {
		callbacks []PageBuildEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single PageBuildEventHandleFunc",
			args: args{
				[]PageBuildEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PageBuildEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple PageBuildEventHandleFuncs",
			args: args{
				[]PageBuildEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PageBuildEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.PageBuildEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnPageBuildEventAny(tt.args.callbacks...)
			if len(g.onPageBuildEvent["*"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onPageBuildEvent["*"]))
			}
		})
	}
}

func TestSetOnPageBuildEventAny(t *testing.T) {
	type args struct {
		callbacks []PageBuildEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single PageBuildEventHandleFunc",
			args: args{
				[]PageBuildEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PageBuildEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple PageBuildEventHandleFuncs",
			args: args{
				[]PageBuildEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PageBuildEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.PageBuildEvent) error {
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
			g.SetOnPageBuildEventAny(func(deliveryID string, eventName string, event *github.PageBuildEvent) error {
				return nil
			})
			g.SetOnPageBuildEventAny(tt.args.callbacks...)
			if len(g.onPageBuildEvent["*"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onPageBuildEvent["*"]), tt.want)
			}
		})
	}
}

func TestHandlePageBuildEventAny(t *testing.T) {

	type args struct {
		deliveryID string
		eventName  string
		event      *github.PageBuildEvent
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
				eventName:  "page_build",

				event: &github.PageBuildEvent{},

				fail: false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "page_build",

				event: &github.PageBuildEvent{},

				fail: true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "page_build",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnPageBuildEventAny(func(deliveryID string, eventName string, event *github.PageBuildEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handlePageBuildEventAny(tt.args.deliveryID, tt.args.deliveryID, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("TestHandlePageBuildEventAny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
