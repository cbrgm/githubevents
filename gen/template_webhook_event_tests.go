package main

var webhookTestsTemplate = `
// Copyright 2022 The GithubEvents Authors. All rights reserved.
// Use of this source code is governed by the MIT License
// that can be found in the LICENSE file.

package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"errors"
	"github.com/google/go-github/v55/github"
	"testing"
	"sync"
)

{{ range $_, $webhook := .Webhooks }}

func TestOn{{ $webhook.Event }}Any(t *testing.T) {
	type args struct {
		callbacks []{{ $webhook.Event }}HandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single {{ $webhook.Event }}HandleFunc",
			args: args{
				[]{{ $webhook.Event }}HandleFunc{
					func(deliveryID string, eventName string, event *github.{{ $webhook.Event }}) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple {{ $webhook.Event }}HandleFuncs",
			args: args{
				[]{{ $webhook.Event }}HandleFunc{
					func(deliveryID string, eventName string, event *github.{{ $webhook.Event }}) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.{{ $webhook.Event }}) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.On{{ $webhook.Event }}Any(tt.args.callbacks...)
			if len(g.on{{ $webhook.Event }}[{{ $webhook.Event }}AnyAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.on{{ $webhook.Event }}[{{ $webhook.Event }}AnyAction]))
			}
		})
	}
}

func TestSetOn{{ $webhook.Event }}Any(t *testing.T) {
	type args struct {
		callbacks []{{ $webhook.Event }}HandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single {{ $webhook.Event }}HandleFunc",
			args: args{
				[]{{ $webhook.Event }}HandleFunc{
					func(deliveryID string, eventName string, event *github.{{ $webhook.Event }}) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple {{ $webhook.Event }}HandleFuncs",
			args: args{
				[]{{ $webhook.Event }}HandleFunc{
					func(deliveryID string, eventName string, event *github.{{ $webhook.Event }}) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.{{ $webhook.Event }}) error {
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
			g.SetOn{{ $webhook.Event }}Any(func(deliveryID string, eventName string, event *github.{{ $webhook.Event }}) error {
				return nil
			})
			g.SetOn{{ $webhook.Event }}Any(tt.args.callbacks...)
			if len(g.on{{ $webhook.Event }}[{{ $webhook.Event }}AnyAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.on{{ $webhook.Event }}[{{ $webhook.Event }}AnyAction]), tt.want)
			}
		})
	}
}

func TestHandle{{ $webhook.Event }}Any(t *testing.T) {
	{{ if $webhook.HasActions }}
	action := "*"
	{{ end }}

	type args struct {
		deliveryID string
		eventName  string
		event      *github.{{ $webhook.Event }}
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
				eventName:  "{{ $webhook.Name }}",
				{{ if $webhook.HasActions }}
				event:    &github.{{ $webhook.Event }}{Action: &action},
				{{ else }}
				event:    &github.{{ $webhook.Event }}{},
				{{ end }}
				fail:     false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "{{ $webhook.Name }}",
				{{ if $webhook.HasActions }}
				event:    &github.{{ $webhook.Event }}{Action: &action},
				{{ else }}
				event:    &github.{{ $webhook.Event }}{},
				{{ end }}
				fail:     true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "{{ $webhook.Name }}",
				event:    nil,
				fail:     false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.On{{ $webhook.Event }}Any(func(deliveryID string, eventName string, event *github.{{ $webhook.Event }}) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handle{{ $webhook.Event }}Any(tt.args.deliveryID, tt.args.deliveryID, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("TestHandle{{ $webhook.Event }}Any() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

{{ range $_, $action := $webhook.Actions }}

func TestOn{{ $action.Handler }}(t *testing.T) {
	type args struct {
		callbacks []{{ $webhook.Event }}HandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single {{ $webhook.Event }}HandleFunc",
			args: args{
				callbacks: []{{ $webhook.Event }}HandleFunc{
					func(deliveryID string, eventName string, event *github.{{ $webhook.Event }}) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple {{ $webhook.Event }}HandleFunc",
			args: args{
				callbacks: []{{ $webhook.Event }}HandleFunc{
					func(deliveryID string, eventName string, event *github.{{ $webhook.Event }}) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.{{ $webhook.Event }}) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.On{{ $action.Handler }}(tt.args.callbacks...)
			if len(g.on{{ $webhook.Event }}[{{ $action.Handler }}Action]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.on{{ $webhook.Event }}[{{ $action.Handler }}Action]))
			}
		})
	}
}

func TestSetOn{{ $action.Handler }}(t *testing.T) {
	type args struct {
		callbacks []{{ $webhook.Event }}HandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single {{ $webhook.Event }}HandleFunc",
			args: args{
				[]{{ $webhook.Event }}HandleFunc{
					func(deliveryID string, eventName string, event *github.{{ $webhook.Event }}) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple {{ $webhook.Event }}HandleFuncs",
			args: args{
				[]{{ $webhook.Event }}HandleFunc{
					func(deliveryID string, eventName string, event *github.{{ $webhook.Event }}) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.{{ $webhook.Event }}) error {
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
			g.SetOn{{ $action.Handler }}(func(deliveryID string, eventName string, event *github.{{ $webhook.Event }}) error {
				return nil
			})
			g.SetOn{{ $action.Handler }}(tt.args.callbacks...)
			if len(g.on{{ $webhook.Event }}[{{ $action.Handler }}Action]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.on{{ $webhook.Event }}[{{ $action.Handler }}Action]), tt.want)
			}
		})
	}
}

func TestHandle{{ $action.Handler }}(t *testing.T) {
	action := {{ $action.Handler }}Action

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.{{ $webhook.Event }}
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
				eventName:  "{{ $webhook.Name }}",
				event:      &github.{{ $webhook.Event }}{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "{{ $webhook.Name }}",
				event:      &github.{{ $webhook.Event }}{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "{{ $webhook.Name }}",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "{{ $webhook.Name }}",
				event:      &github.{{ $webhook.Event }}{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "{{ $webhook.Name }}",
				event:      &github.{{ $webhook.Event }}{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "{{ $webhook.Name }}",
				event:      &github.{{ $webhook.Event }}{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.On{{ $action.Handler }}(func(deliveryID string, eventName string, event *github.{{ $webhook.Event }}) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handle{{ $action.Handler }}(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handle{{ $action.Handler }}() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

{{ end }}

func Test{{ $webhook.Event }}(t *testing.T) {
	type fields struct {
		handler *EventHandler
	}
	type args struct {
		deliveryID string
		eventName  string
		event      *github.{{ $webhook.Event }}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "must trigger {{ $webhook.Event }}Any with unknown event action",
			fields: fields{
				handler: &EventHandler{
					WebhookSecret: "fake",
					onBeforeAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(deliveryID string, eventName string, event any) error {
								t.Log("onBeforeAny called")
								return nil
							},
						},
					},
					onAfterAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(deliveryID string, eventName string, event any) error {
								t.Log("onAfterAny called")
								return nil
							},
						},
					},
					on{{ $webhook.Event }}: map[string][]{{ $webhook.Event }}HandleFunc{
						{{ $webhook.Event }}AnyAction: {
							func(deliveryID string, eventName string, event *github.{{ $webhook.Event }}) error {
								t.Log("onAny action called")
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  {{ $webhook.Event }},
				{{ if $webhook.HasActions }}
				event:      &github.{{ $webhook.Event }}{Action: ptrString("unknown")},
				{{ else }}
				event:      &github.{{ $webhook.Event }}{},
				{{ end }}
			},
			wantErr: false,
		},
		{{ if $webhook.HasActions }}
		{{ range $_, $action := $webhook.Actions }}
		{
			name: "must trigger {{ $action.Handler }}",
			fields: fields{
				handler: &EventHandler{
					WebhookSecret: "fake",
					onBeforeAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(deliveryID string, eventName string, event any) error {
								t.Log("onBeforeAny called")
								return nil
							},
						},
					},
					onAfterAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(deliveryID string, eventName string, event any) error {
								t.Log("onAfterAny called")
								return nil
							},
						},
					},
					on{{ $webhook.Event }}: map[string][]{{ $webhook.Event }}HandleFunc{
						{{ $webhook.Event }}AnyAction: {
							func(deliveryID string, eventName string, event *github.{{ $webhook.Event }}) error {
								t.Log("onAny action called")
								return nil
							},
						},
						{{ $action.Handler }}Action: {
							func(deliveryID string, eventName string, event *github.{{ $webhook.Event }}) error {
								t.Logf("%s action called", {{ $action.Handler }}Action)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "{{ $webhook.Name }}",
				event:      &github.{{ $webhook.Event }}{Action: ptrString({{ $action.Handler }}Action)},
			},
			wantErr: false,
		},
		{
			name: "must fail {{ $action.Handler }} with empty action",
			fields: fields{
				handler: &EventHandler{
					WebhookSecret: "fake",
					onBeforeAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(deliveryID string, eventName string, event any) error {
								t.Log("onBeforeAny called")
								return nil
							},
						},
					},
					onAfterAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(deliveryID string, eventName string, event any) error {
								t.Log("onAfterAny called")
								return nil
							},
						},
					},
					on{{ $webhook.Event }}: map[string][]{{ $webhook.Event }}HandleFunc{
						{{ $webhook.Event }}AnyAction: {
							func(deliveryID string, eventName string, event *github.{{ $webhook.Event }}) error {
								t.Log("onAny action called")
								return nil
							},
						},
						{{ $action.Handler }}Action: {
							func(deliveryID string, eventName string, event *github.{{ $webhook.Event }}) error {
								t.Logf("%s action called", {{ $action.Handler }}Action)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "{{ $webhook.Name }}",
				event:      &github.{{ $webhook.Event }}{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail {{ $action.Handler }} with nil action",
			fields: fields{
				handler: &EventHandler{
					WebhookSecret: "fake",
					onBeforeAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(deliveryID string, eventName string, event any) error {
								t.Log("onBeforeAny called")
								return nil
							},
						},
					},
					onAfterAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(deliveryID string, eventName string, event any) error {
								t.Log("onAfterAny called")
								return nil
							},
						},
					},
					on{{ $webhook.Event }}: map[string][]{{ $webhook.Event }}HandleFunc{
						{{ $webhook.Event }}AnyAction: {
							func(deliveryID string, eventName string, event *github.{{ $webhook.Event }}) error {
								t.Log("onAny action called")
								return nil
							},
						},
						{{ $action.Handler }}Action: {
							func(deliveryID string, eventName string, event *github.{{ $webhook.Event }}) error {
								t.Logf("%s action called", {{ $action.Handler }}Action)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "{{ $webhook.Name }}",
				event:      &github.{{ $webhook.Event }}{Action: nil},
			},
			wantErr: true,
		},
		{{ end }}
		{{ end }}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &EventHandler{
				WebhookSecret: "fake",
				mu:            sync.RWMutex{},
			}
			if err := g.{{ $webhook.Event }}(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("{{ $webhook.Event }}() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

{{ end }}
`
