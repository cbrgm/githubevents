// Copyright 2022 The GithubEvents Authors. All rights reserved.
// Use of this source code is governed by the MIT License
// that can be found in the LICENSE file.

package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"fmt"
	"github.com/google/go-github/v55/github"
	"golang.org/x/sync/errgroup"
	"net/http"
	"sync"
)

// Actions are used to identify registered callbacks.
const (
	// EventAnyAction is used to identify callbacks listening to all events.
	EventAnyAction = "*"
)

// EventHandler represents a Github webhook handler.
type EventHandler struct {
	// settings

	// WebhookSecret is the GitHub Webhook secret token.
	WebhookSecret string

	// handleFuncs

	onBranchProtectionRuleEvent         map[string][]BranchProtectionRuleEventHandleFunc
	onCheckRunEvent                     map[string][]CheckRunEventHandleFunc
	onCheckSuiteEvent                   map[string][]CheckSuiteEventHandleFunc
	onCommitCommentEvent                map[string][]CommitCommentEventHandleFunc
	onCreateEvent                       map[string][]CreateEventHandleFunc
	onDeleteEvent                       map[string][]DeleteEventHandleFunc
	onDeployKeyEvent                    map[string][]DeployKeyEventHandleFunc
	onDeploymentEvent                   map[string][]DeploymentEventHandleFunc
	onDeploymentStatusEvent             map[string][]DeploymentStatusEventHandleFunc
	onDiscussionEvent                   map[string][]DiscussionEventHandleFunc
	onForkEvent                         map[string][]ForkEventHandleFunc
	onGitHubAppAuthorizationEvent       map[string][]GitHubAppAuthorizationEventHandleFunc
	onGollumEvent                       map[string][]GollumEventHandleFunc
	onInstallationEvent                 map[string][]InstallationEventHandleFunc
	onInstallationRepositoriesEvent     map[string][]InstallationRepositoriesEventHandleFunc
	onIssueCommentEvent                 map[string][]IssueCommentEventHandleFunc
	onIssuesEvent                       map[string][]IssuesEventHandleFunc
	onLabelEvent                        map[string][]LabelEventHandleFunc
	onMarketplacePurchaseEvent          map[string][]MarketplacePurchaseEventHandleFunc
	onMemberEvent                       map[string][]MemberEventHandleFunc
	onMembershipEvent                   map[string][]MembershipEventHandleFunc
	onMetaEvent                         map[string][]MetaEventHandleFunc
	onMilestoneEvent                    map[string][]MilestoneEventHandleFunc
	onOrganizationEvent                 map[string][]OrganizationEventHandleFunc
	onOrgBlockEvent                     map[string][]OrgBlockEventHandleFunc
	onPackageEvent                      map[string][]PackageEventHandleFunc
	onPageBuildEvent                    map[string][]PageBuildEventHandleFunc
	onPingEvent                         map[string][]PingEventHandleFunc
	onProjectEvent                      map[string][]ProjectEventHandleFunc
	onProjectCardEvent                  map[string][]ProjectCardEventHandleFunc
	onProjectColumnEvent                map[string][]ProjectColumnEventHandleFunc
	onPublicEvent                       map[string][]PublicEventHandleFunc
	onPullRequestEvent                  map[string][]PullRequestEventHandleFunc
	onPullRequestReviewEvent            map[string][]PullRequestReviewEventHandleFunc
	onPullRequestReviewCommentEvent     map[string][]PullRequestReviewCommentEventHandleFunc
	onPushEvent                         map[string][]PushEventHandleFunc
	onReleaseEvent                      map[string][]ReleaseEventHandleFunc
	onRepositoryDispatchEvent           map[string][]RepositoryDispatchEventHandleFunc
	onRepositoryEvent                   map[string][]RepositoryEventHandleFunc
	onRepositoryVulnerabilityAlertEvent map[string][]RepositoryVulnerabilityAlertEventHandleFunc
	onStarEvent                         map[string][]StarEventHandleFunc
	onStatusEvent                       map[string][]StatusEventHandleFunc
	onTeamEvent                         map[string][]TeamEventHandleFunc
	onTeamAddEvent                      map[string][]TeamAddEventHandleFunc
	onWatchEvent                        map[string][]WatchEventHandleFunc
	onWorkflowJobEvent                  map[string][]WorkflowJobEventHandleFunc
	onWorkflowDispatchEvent             map[string][]WorkflowDispatchEventHandleFunc
	onWorkflowRunEvent                  map[string][]WorkflowRunEventHandleFunc

	onBeforeAny map[string][]EventHandleFunc
	onAfterAny  map[string][]EventHandleFunc
	onErrorAny  map[string][]ErrorEventHandleFunc

	// internal
	mu sync.RWMutex
}

// New returns EventHandler
// webhookSecret is the GitHub Webhook secret token.
// If your webhook does not contain a secret token, you can set nil.
// This is intended for local development purposes only and all webhooks should ideally set up a secret token.
func New(webhookSecret string) *EventHandler {
	return &EventHandler{
		WebhookSecret: webhookSecret,
	}
}

// EventHandleFunc represents a generic callback function which receives any event.
type EventHandleFunc func(deliveryID string, eventName string, event interface{}) error

// OnBeforeAny registers callbacks which are triggered before any event.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnBeforeAny must be used.
// Callbacks are executed in parallel.
func (g *EventHandler) OnBeforeAny(callbacks ...EventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onBeforeAny == nil {
		g.onBeforeAny = make(map[string][]EventHandleFunc)
	}
	g.onBeforeAny[EventAnyAction] = append(g.onBeforeAny[EventAnyAction], callbacks...)
}

// SetOnBeforeAny registers  callbacks which are triggered before any event
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnBeforeAny must be used.
// Callbacks are executed in parallel.
func (g *EventHandler) SetOnBeforeAny(callbacks ...EventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onBeforeAny == nil {
		g.onBeforeAny = make(map[string][]EventHandleFunc)
	}
	g.onBeforeAny[EventAnyAction] = callbacks
}

func (g *EventHandler) handleBeforeAny(deliveryID string, eventName string, event interface{}) error {
	if event == nil {
		return fmt.Errorf("event was empty or nil")
	}
	if _, ok := g.onBeforeAny[EventAnyAction]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onBeforeAny[EventAnyAction] {
		handle := h
		eg.Go(func() error {
			err := handle(deliveryID, eventName, event)
			if err != nil {
				return err
			}
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		return err
	}
	return nil
}

// OnAfterAny registers callbacks which are triggered after any event.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnAfterAny must be used.
// Callbacks are executed in parallel.
func (g *EventHandler) OnAfterAny(callbacks ...EventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onAfterAny == nil {
		g.onAfterAny = make(map[string][]EventHandleFunc)
	}
	g.onAfterAny[EventAnyAction] = append(g.onAfterAny[EventAnyAction], callbacks...)
}

// SetOnAfterAny registers  callbacks which are triggered after any event
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnAfterAny must be used.
// Callbacks are executed in parallel.
func (g *EventHandler) SetOnAfterAny(callbacks ...EventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onAfterAny == nil {
		g.onAfterAny = make(map[string][]EventHandleFunc)
	}
	g.onAfterAny[EventAnyAction] = callbacks
}

func (g *EventHandler) handleAfterAny(deliveryID string, eventName string, event interface{}) error {
	if event == nil {
		return fmt.Errorf("event was empty or nil")
	}
	if _, ok := g.onAfterAny[EventAnyAction]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onAfterAny[EventAnyAction] {
		handle := h
		eg.Go(func() error {
			err := handle(deliveryID, eventName, event)
			if err != nil {
				return err
			}
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		return err
	}
	return nil
}

// ErrorEventHandleFunc represents a generic callback function which receives any event and an error thrown by
// some function on a higher level.
type ErrorEventHandleFunc func(deliveryID string, eventName string, event interface{}, err error) error

// OnError registers callbacks which are triggered whenever an error occurs.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnError must be used.
// Callbacks are executed in parallel.
func (g *EventHandler) OnError(callbacks ...ErrorEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onErrorAny == nil {
		g.onErrorAny = make(map[string][]ErrorEventHandleFunc)
	}
	g.onErrorAny[EventAnyAction] = append(g.onErrorAny[EventAnyAction], callbacks...)
}

// SetOnError registers callbacks which are triggered whenever an error occurs
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnError must be used.
// Callbacks are executed in parallel.
func (g *EventHandler) SetOnError(callbacks ...ErrorEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onErrorAny == nil {
		g.onErrorAny = make(map[string][]ErrorEventHandleFunc)
	}
	g.onErrorAny[EventAnyAction] = callbacks
}

func (g *EventHandler) handleError(deliveryID string, eventName string, event interface{}, err error) error {
	if event == nil {
		return fmt.Errorf("event was empty or nil")
	}
	if _, ok := g.onErrorAny[EventAnyAction]; !ok {
		return err
	}
	eg := new(errgroup.Group)
	for _, h := range g.onErrorAny[EventAnyAction] {
		handle := h
		eg.Go(func() error {
			err := handle(deliveryID, eventName, event, err)
			if err != nil {
				return err
			}
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		return err
	}
	return nil
}

// HandleEventRequest parses a Github event from http.Request and executes registered handlers.
func (g *EventHandler) HandleEventRequest(req *http.Request) error {
	payload, err := github.ValidatePayload(req, []byte(g.WebhookSecret))
	if err != nil {
		fmt.Errorf("could not validate webhook payload: err=%s\n", err)
		return err
	}
	event, err := github.ParseWebHook(github.WebHookType(req), payload)
	if err != nil {
		fmt.Errorf("could not parse webhook: err=%s\n", err)
		return err
	}

	deliveryID := github.DeliveryID(req)
	eventName := github.WebHookType(req)

	switch event := event.(type) {

	case *github.BranchProtectionRuleEvent:
		return g.BranchProtectionRuleEvent(deliveryID, eventName, event)

	case *github.CheckRunEvent:
		return g.CheckRunEvent(deliveryID, eventName, event)

	case *github.CheckSuiteEvent:
		return g.CheckSuiteEvent(deliveryID, eventName, event)

	case *github.CommitCommentEvent:
		return g.CommitCommentEvent(deliveryID, eventName, event)

	case *github.CreateEvent:
		return g.CreateEvent(deliveryID, eventName, event)

	case *github.DeleteEvent:
		return g.DeleteEvent(deliveryID, eventName, event)

	case *github.DeployKeyEvent:
		return g.DeployKeyEvent(deliveryID, eventName, event)

	case *github.DeploymentEvent:
		return g.DeploymentEvent(deliveryID, eventName, event)

	case *github.DeploymentStatusEvent:
		return g.DeploymentStatusEvent(deliveryID, eventName, event)

	case *github.DiscussionEvent:
		return g.DiscussionEvent(deliveryID, eventName, event)

	case *github.ForkEvent:
		return g.ForkEvent(deliveryID, eventName, event)

	case *github.GitHubAppAuthorizationEvent:
		return g.GitHubAppAuthorizationEvent(deliveryID, eventName, event)

	case *github.GollumEvent:
		return g.GollumEvent(deliveryID, eventName, event)

	case *github.InstallationEvent:
		return g.InstallationEvent(deliveryID, eventName, event)

	case *github.InstallationRepositoriesEvent:
		return g.InstallationRepositoriesEvent(deliveryID, eventName, event)

	case *github.IssueCommentEvent:
		return g.IssueCommentEvent(deliveryID, eventName, event)

	case *github.IssuesEvent:
		return g.IssuesEvent(deliveryID, eventName, event)

	case *github.LabelEvent:
		return g.LabelEvent(deliveryID, eventName, event)

	case *github.MarketplacePurchaseEvent:
		return g.MarketplacePurchaseEvent(deliveryID, eventName, event)

	case *github.MemberEvent:
		return g.MemberEvent(deliveryID, eventName, event)

	case *github.MembershipEvent:
		return g.MembershipEvent(deliveryID, eventName, event)

	case *github.MetaEvent:
		return g.MetaEvent(deliveryID, eventName, event)

	case *github.MilestoneEvent:
		return g.MilestoneEvent(deliveryID, eventName, event)

	case *github.OrganizationEvent:
		return g.OrganizationEvent(deliveryID, eventName, event)

	case *github.OrgBlockEvent:
		return g.OrgBlockEvent(deliveryID, eventName, event)

	case *github.PackageEvent:
		return g.PackageEvent(deliveryID, eventName, event)

	case *github.PageBuildEvent:
		return g.PageBuildEvent(deliveryID, eventName, event)

	case *github.PingEvent:
		return g.PingEvent(deliveryID, eventName, event)

	case *github.ProjectEvent:
		return g.ProjectEvent(deliveryID, eventName, event)

	case *github.ProjectCardEvent:
		return g.ProjectCardEvent(deliveryID, eventName, event)

	case *github.ProjectColumnEvent:
		return g.ProjectColumnEvent(deliveryID, eventName, event)

	case *github.PublicEvent:
		return g.PublicEvent(deliveryID, eventName, event)

	case *github.PullRequestEvent:
		return g.PullRequestEvent(deliveryID, eventName, event)

	case *github.PullRequestReviewEvent:
		return g.PullRequestReviewEvent(deliveryID, eventName, event)

	case *github.PullRequestReviewCommentEvent:
		return g.PullRequestReviewCommentEvent(deliveryID, eventName, event)

	case *github.PushEvent:
		return g.PushEvent(deliveryID, eventName, event)

	case *github.ReleaseEvent:
		return g.ReleaseEvent(deliveryID, eventName, event)

	case *github.RepositoryDispatchEvent:
		return g.RepositoryDispatchEvent(deliveryID, eventName, event)

	case *github.RepositoryEvent:
		return g.RepositoryEvent(deliveryID, eventName, event)

	case *github.RepositoryVulnerabilityAlertEvent:
		return g.RepositoryVulnerabilityAlertEvent(deliveryID, eventName, event)

	case *github.StarEvent:
		return g.StarEvent(deliveryID, eventName, event)

	case *github.StatusEvent:
		return g.StatusEvent(deliveryID, eventName, event)

	case *github.TeamEvent:
		return g.TeamEvent(deliveryID, eventName, event)

	case *github.TeamAddEvent:
		return g.TeamAddEvent(deliveryID, eventName, event)

	case *github.WatchEvent:
		return g.WatchEvent(deliveryID, eventName, event)

	case *github.WorkflowJobEvent:
		return g.WorkflowJobEvent(deliveryID, eventName, event)

	case *github.WorkflowDispatchEvent:
		return g.WorkflowDispatchEvent(deliveryID, eventName, event)

	case *github.WorkflowRunEvent:
		return g.WorkflowRunEvent(deliveryID, eventName, event)

	}
	return nil
}

// ptrString returns a string pointer.
func ptrString(s string) *string {
	return &s
}
