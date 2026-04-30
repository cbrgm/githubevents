// Copyright 2022 The GithubEvents Authors. All rights reserved.
// Use of this source code is governed by the MIT License
// that can be found in the LICENSE file.

// Package githubcompat is the single point at which this module pins to a
// specific major version of github.com/google/go-github. The rest of the
// repository imports the symbols it needs through this package so that a
// future major-version bump of go-github only requires updating the import
// path below (assuming the upstream type and function names remain stable).
//
// This package is intentionally kept tiny: it only re-exports the function
// values and event types that the generated webhook handlers reference. All
// re-exports use Go type aliases (type X = gogithub.X) so external users that
// already import go-github directly can keep passing *github.<Event> values
// from go-github into this library's callbacks without any wrapping.
//
// External callers normally do not need to reference githubcompat: they keep
// importing github.com/google/go-github/.../github directly for their own
// types. The package is exported so consumers can opt into the same pinned
// alias if they want to avoid hard-coding a specific go-github major version.
package githubcompat

import (
	gogithub "github.com/google/go-github/v85/github"
)

// Webhook helper functions re-exported from go-github. These are plain
// function values so the call sites in generated code (e.g.
// github.ValidatePayload(...)) remain unchanged.
var (
	ValidatePayload = gogithub.ValidatePayload
	ParseWebHook    = gogithub.ParseWebHook
	WebHookType     = gogithub.WebHookType
	DeliveryID      = gogithub.DeliveryID
)

// Event type aliases. These MUST be Go type aliases (type X = Y) rather than
// new named types so that *gogithub.IssueCommentEvent and
// *githubcompat.IssueCommentEvent are the exact same type, allowing callbacks
// declared with go-github's types to satisfy this library's HandleFunc types.
//
// The list mirrors the GithubWebhooks.Event values in gen/template_params.go.
type (
	BranchProtectionRuleEvent         = gogithub.BranchProtectionRuleEvent
	CheckRunEvent                     = gogithub.CheckRunEvent
	CheckSuiteEvent                   = gogithub.CheckSuiteEvent
	CommitCommentEvent                = gogithub.CommitCommentEvent
	CreateEvent                       = gogithub.CreateEvent
	CustomPropertyEvent               = gogithub.CustomPropertyEvent
	CustomPropertyValuesEvent         = gogithub.CustomPropertyValuesEvent
	DeleteEvent                       = gogithub.DeleteEvent
	DeployKeyEvent                    = gogithub.DeployKeyEvent
	DeploymentEvent                   = gogithub.DeploymentEvent
	DeploymentStatusEvent             = gogithub.DeploymentStatusEvent
	DiscussionEvent                   = gogithub.DiscussionEvent
	ForkEvent                         = gogithub.ForkEvent
	GitHubAppAuthorizationEvent       = gogithub.GitHubAppAuthorizationEvent
	GollumEvent                       = gogithub.GollumEvent
	InstallationEvent                 = gogithub.InstallationEvent
	InstallationRepositoriesEvent     = gogithub.InstallationRepositoriesEvent
	IssueCommentEvent                 = gogithub.IssueCommentEvent
	IssuesEvent                       = gogithub.IssuesEvent
	LabelEvent                        = gogithub.LabelEvent
	MarketplacePurchaseEvent          = gogithub.MarketplacePurchaseEvent
	MemberEvent                       = gogithub.MemberEvent
	MembershipEvent                   = gogithub.MembershipEvent
	MergeGroupEvent                   = gogithub.MergeGroupEvent
	MetaEvent                         = gogithub.MetaEvent
	MilestoneEvent                    = gogithub.MilestoneEvent
	OrganizationEvent                 = gogithub.OrganizationEvent
	OrgBlockEvent                     = gogithub.OrgBlockEvent
	PackageEvent                      = gogithub.PackageEvent
	PageBuildEvent                    = gogithub.PageBuildEvent
	PingEvent                         = gogithub.PingEvent
	ProjectV2Event                    = gogithub.ProjectV2Event
	ProjectV2ItemEvent                = gogithub.ProjectV2ItemEvent
	PublicEvent                       = gogithub.PublicEvent
	PullRequestEvent                  = gogithub.PullRequestEvent
	PullRequestReviewEvent            = gogithub.PullRequestReviewEvent
	PullRequestReviewCommentEvent     = gogithub.PullRequestReviewCommentEvent
	PushEvent                         = gogithub.PushEvent
	ReleaseEvent                      = gogithub.ReleaseEvent
	RepositoryDispatchEvent           = gogithub.RepositoryDispatchEvent
	RepositoryRulesetEvent            = gogithub.RepositoryRulesetEvent
	RepositoryEvent                   = gogithub.RepositoryEvent
	RepositoryVulnerabilityAlertEvent = gogithub.RepositoryVulnerabilityAlertEvent
	StarEvent                         = gogithub.StarEvent
	StatusEvent                       = gogithub.StatusEvent
	TeamEvent                         = gogithub.TeamEvent
	TeamAddEvent                      = gogithub.TeamAddEvent
	WatchEvent                        = gogithub.WatchEvent
	WorkflowJobEvent                  = gogithub.WorkflowJobEvent
	WorkflowDispatchEvent             = gogithub.WorkflowDispatchEvent
	WorkflowRunEvent                  = gogithub.WorkflowRunEvent
)
