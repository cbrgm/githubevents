package main

// TemplateParameters represents template parameters
type TemplateParameters struct {
	Webhooks []GithubWebhooks
}

// GithubWebhooks represents a Github webhook event type parameters
type GithubWebhooks struct {
	// Event is the name of the event taken from
	// https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads
	Event string
	// Name is the event name taken from
	// // https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads
	Name string
	// Actions is a list of webhook actions taken from
	// https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads
	// must be empty when no actions are defined.
	Actions []Action
	// HasActions is used for generation and describes whether the event
	// has actions or not
	HasActions bool
}

// Action represents a webhook event action
type Action struct {
	// Handler is the name of the handleFunc to be generated
	Handler string
	// Action is the name of the action type taken from:
	// https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads
	Action string
}

var params = TemplateParameters{
	Webhooks: []GithubWebhooks{
		{
			Event: "BranchProtectionRuleEvent",
			Name:  "branch_protection_rule",
			Actions: []Action{
				{
					Handler: "BranchProtectionRuleEventCreated",
					Action:  "created",
				},
				{
					Handler: "BranchProtectionRuleEventEdited",
					Action:  "edited",
				},
				{
					Handler: "BranchProtectionRuleEventDeleted",
					Action:  "deleted",
				},
			},
			HasActions: true,
		},
		{
			Event: "CheckRunEvent",
			Name:  "check_run",
			Actions: []Action{
				{
					Handler: "CheckRunEventCreated",
					Action:  "created",
				},
				{
					Handler: "CheckRunEventCompleted",
					Action:  "completed",
				},
				{
					Handler: "CheckRunEventReRequested",
					Action:  "rerequested",
				},
				{
					Handler: "CheckRunEventRequestAction",
					Action:  "requested_action",
				},
			},
			HasActions: true,
		},
		{
			Event: "CheckSuiteEvent",
			Name:  "check_suite",
			Actions: []Action{
				{
					Handler: "CheckSuiteEventCompleted",
					Action:  "completed",
				},
				{
					Handler: "CheckSuiteEventRequested",
					Action:  "requested",
				},
				{
					Handler: "CheckSuiteEventReRequested",
					Action:  "rerequested",
				},
			},
			HasActions: true,
		},
		{
			Event: "CommitCommentEvent",
			Name:  "commit_comment",
			Actions: []Action{
				{
					Handler: "CommitCommentEventCreated",
					Action:  "created",
				},
			},
			HasActions: true,
		},
		{
			Event:      "CreateEvent",
			Name:       "create",
			HasActions: false,
		},
		{
			Event:      "DeleteEvent",
			Name:       "delete",
			HasActions: false,
		},
		{
			Event: "DeployKeyEvent",
			Name:  "deploy_key",
			Actions: []Action{
				{
					Handler: "DeployKeyEventCreated",
					Action:  "created",
				},
				{
					Handler: "DeployKeyEventDeleted",
					Action:  "deleted",
				},
			},
			HasActions: true,
		},
		{
			Event:      "DeploymentEvent",
			Name:       "deployment",
			HasActions: false,
		},
		{
			Event:      "DeploymentStatusEvent",
			Name:       "deployment_status",
			HasActions: false,
		},
		{
			Event: "DiscussionEvent",
			Name:  "discussion",
			Actions: []Action{
				{
					Handler: "DiscussionEventCreated",
					Action:  "created",
				},
				{
					Handler: "DiscussionEventEdited",
					Action:  "edited",
				},
				{
					Handler: "DiscussionEventDeleted",
					Action:  "deleted",
				},
				{
					Handler: "DiscussionEventPinned",
					Action:  "pinned",
				},
				{
					Handler: "DiscussionEventUnpinned",
					Action:  "unpinned",
				},
				{
					Handler: "DiscussionEventLocked",
					Action:  "locked",
				},
				{
					Handler: "DiscussionEventUnlocked",
					Action:  "unlocked",
				},
				{
					Handler: "DiscussionEventTransferred",
					Action:  "transferred",
				},
				{
					Handler: "DiscussionEventCategoryChanged",
					Action:  "category_changed",
				},
				{
					Handler: "DiscussionEventAnswered",
					Action:  "answered",
				},
				{
					Handler: "DiscussionEventUnanswered",
					Action:  "unanswered",
				},
				{
					Handler: "DiscussionEventLabeled",
					Action:  "labeled",
				},
				{
					Handler: "DiscussionEventUnlabeled",
					Action:  "unlabeled",
				},
			},
			HasActions: true,
		},
		{
			Event:      "ForkEvent",
			Name:       "fork",
			HasActions: false,
		},
		{
			Event: "GitHubAppAuthorizationEvent",
			Name:  "github_app_authorization",
			Actions: []Action{
				{
					Handler: "GitHubAppAuthorizationEventRevoked",
					Action:  "revoked",
				},
			},
			HasActions: true,
		},
		{
			Event:      "GollumEvent",
			Name:       "gollum",
			HasActions: false,
		},
		{
			Event: "InstallationEvent",
			Name:  "installation",
			Actions: []Action{
				{
					Handler: "InstallationEventCreated",
					Action:  "created",
				},
				{
					Handler: "InstallationEventDeleted",
					Action:  "deleted",
				},
				{
					Handler: "InstallationEventEventSuspend",
					Action:  "suspend",
				},
				{
					Handler: "InstallationEventEventUnsuspend",
					Action:  "unsuspend",
				},
				{
					Handler: "InstallationEventNewPermissionsAccepted",
					Action:  "new_permissions_accepted",
				},
			},
			HasActions: true,
		},
		{
			Event: "InstallationRepositoriesEvent",
			Name:  "installation_repositories",
			Actions: []Action{
				{
					Handler: "InstallationRepositoriesEventAdded",
					Action:  "added",
				},
				{
					Handler: "InstallationRepositoriesEventRemoved",
					Action:  "removed",
				},
			},
			HasActions: true,
		},
		{
			Event: "IssueCommentEvent",
			Name:  "issue_comment",
			Actions: []Action{
				{
					Handler: "IssueCommentCreated",
					Action:  "created",
				},
				{
					Handler: "IssueCommentEdited",
					Action:  "edited",
				},
				{
					Handler: "IssueCommentDeleted",
					Action:  "deleted",
				},
			},
			HasActions: true,
		},
		{
			Event: "IssuesEvent",
			Name:  "issues",
			Actions: []Action{
				{
					Handler: "IssuesEventOpened",
					Action:  "opened",
				},
				{
					Handler: "IssuesEventEdited",
					Action:  "edited",
				},
				{
					Handler: "IssuesEventDeleted",
					Action:  "deleted",
				},
				{
					Handler: "IssuesEventPinned",
					Action:  "pinned",
				},
				{
					Handler: "IssuesEventUnpinned",
					Action:  "unpinned",
				},
				{
					Handler: "IssuesEventClosed",
					Action:  "closed",
				},
				{
					Handler: "IssuesEventReopened",
					Action:  "reopened",
				},
				{
					Handler: "IssuesEventAssigned",
					Action:  "assigned",
				},
				{
					Handler: "IssuesEventUnassigned",
					Action:  "unassigned",
				},
				{
					Handler: "IssuesEventLabeled",
					Action:  "labeled",
				},
				{
					Handler: "IssuesEventUnlabeled",
					Action:  "unlabeled",
				},
				{
					Handler: "IssuesEventLocked",
					Action:  "locked",
				},
				{
					Handler: "IssuesEventUnlocked",
					Action:  "unlocked",
				},
				{
					Handler: "IssuesEventTransferred",
					Action:  "transferred",
				},
				{
					Handler: "IssuesEventMilestoned",
					Action:  "milestoned",
				},
				{
					Handler: "IssuesEventDeMilestoned",
					Action:  "demilestoned",
				},
			},
			HasActions: true,
		},
		{
			Event: "LabelEvent",
			Name:  "label",
			Actions: []Action{
				{
					Handler: "LabelEventCreated",
					Action:  "created",
				},
				{
					Handler: "LabelEventEdited",
					Action:  "edited",
				},
				{
					Handler: "LabelEventDeleted",
					Action:  "deleted",
				},
			},
			HasActions: true,
		},
		{
			Event: "MarketplacePurchaseEvent",
			Name:  "marketplace_purchase",
			Actions: []Action{
				{
					Handler: "MarketplacePurchaseEventPurchased",
					Action:  "purchased",
				},
				{
					Handler: "MarketplacePurchaseEventPendingChange",
					Action:  "pending_change",
				},
				{
					Handler: "MarketplacePurchaseEventPendingChangeCancelled",
					Action:  "pending_change_cancelled",
				},
				{
					Handler: "MarketplacePurchaseEventChanged",
					Action:  "changed",
				},
				{
					Handler: "MarketplacePurchaseEventCancelled",
					Action:  "cancelled",
				},
			},
			HasActions: true,
		},
		{
			Event: "MemberEvent",
			Name:  "member",
			Actions: []Action{
				{
					Handler: "MemberEventAdded",
					Action:  "added",
				},
				{
					Handler: "MemberEventRemoved",
					Action:  "removed",
				},
				{
					Handler: "MemberEventEdited",
					Action:  "edited",
				},
			},
			HasActions: true,
		},
		{
			Event: "MembershipEvent",
			Name:  "membership",
			Actions: []Action{
				{
					Handler: "MembershipEventAdded",
					Action:  "added",
				},
				{
					Handler: "MembershipEventRemoved",
					Action:  "removed",
				},
			},
			HasActions: true,
		},
		{
			Event: "MergeGroupEvent",
			Name:  "merge_group_event",
			Actions: []Action{
				{
					Handler: "MergeGroupEventChecksRequested",
					Action:  "checks_requested",
				},
				{
					Handler: "MergeGroupEventDestroyed",
					Action:  "destroyed",
				},
			},
			HasActions: true,
		},
		{
			Event:      "MetaEvent",
			Name:       "meta",
			HasActions: false,
		},
		{
			Event: "MilestoneEvent",
			Name:  "milestone",
			Actions: []Action{
				{
					Handler: "MilestoneEventCreated",
					Action:  "created",
				},
				{
					Handler: "MilestoneEventClosed",
					Action:  "closed",
				},
				{
					Handler: "MilestoneEventOpened",
					Action:  "opened",
				},
				{
					Handler: "MilestoneEventEdited",
					Action:  "edited",
				},
				{
					Handler: "MilestoneEventDeleted",
					Action:  "deleted",
				},
			},
			HasActions: true,
		},
		{
			Event: "OrganizationEvent",
			Name:  "organization",
			Actions: []Action{
				{
					Handler: "OrganizationEventDeleted",
					Action:  "deleted",
				},
				{
					Handler: "OrganizationEventRenamed",
					Action:  "renamed",
				},
				{
					Handler: "OrganizationEventMemberAdded",
					Action:  "member_added",
				},
				{
					Handler: "OrganizationEventMemberRemoved",
					Action:  "member_removed",
				},
				{
					Handler: "OrganizationEventMemberInvited",
					Action:  "member_invited",
				},
			},
			HasActions: true,
		},
		{
			Event: "OrgBlockEvent",
			Name:  "org_block",
			Actions: []Action{
				{
					Handler: "OrgBlockEventBlocked",
					Action:  "blocked",
				},
				{
					Handler: "OrgBlockEventUnblocked",
					Action:  "unblocked",
				},
			},
			HasActions: true,
		},
		{
			Event: "PackageEvent",
			Name:  "package",
			Actions: []Action{
				{
					Handler: "PackageEventPublished",
					Action:  "published",
				},
				{
					Handler: "PackageEventUpdated",
					Action:  "updated",
				},
			},
			HasActions: true,
		},
		{
			Event:      "PageBuildEvent",
			Name:       "page_build",
			HasActions: false,
		},
		{
			Event:      "PingEvent",
			Name:       "ping",
			HasActions: false,
		},
		{
			Event: "ProjectV2Event",
			Name:  "project_v2",
			Actions: []Action{
				{
					Handler: "ProjectEventCreated",
					Action:  "created",
				},
				{
					Handler: "ProjectEventEdited",
					Action:  "edited",
				},
				{
					Handler: "ProjectEventClosed",
					Action:  "closed",
				},
				{
					Handler: "ProjectEventReopened",
					Action:  "reopened",
				},
				{
					Handler: "ProjectEventDeleted",
					Action:  "deleted",
				},
			},
			HasActions: true,
		},
		{
			Event: "ProjectV2ItemEvent",
			Name:  "project_v2_item",
			Actions: []Action{
				{
					Handler: "ProjectItemEventCreated",
					Action:  "created",
				},
				{
					Handler: "ProjectItemEventEdited",
					Action:  "edited",
				},
				{
					Handler: "ProjectItemEventClosed",
					Action:  "closed",
				},
				{
					Handler: "ProjectItemEventReopened",
					Action:  "reopened",
				},
				{
					Handler: "ProjectItemEventDeleted",
					Action:  "deleted",
				},
				{
					Handler: "ProjectItemEventConverted",
					Action:  "converted",
				},
				{
					Handler: "ProjectItemEventRestored",
					Action:  "restored",
				},
			},
			HasActions: true,
		},
		{
			Event:      "PublicEvent",
			Name:       "public",
			HasActions: false,
		},
		{
			Event: "PullRequestEvent",
			Name:  "pull_request",
			Actions: []Action{
				{
					Handler: "PullRequestEventAssigned",
					Action:  "assigned",
				},
				{
					Handler: "PullRequestEventAutoMergeDisabled",
					Action:  "auto_merge_disabled",
				},
				{
					Handler: "PullRequestEventAutoMergeEnabled",
					Action:  "auto_merge_enabled",
				},
				{
					Handler: "PullRequestEventClosed",
					Action:  "closed",
				},
				{
					Handler: "PullRequestEventConvertedToDraft",
					Action:  "converted_to_draft",
				},
				{
					Handler: "PullRequestEventEdited",
					Action:  "edited",
				},
				{
					Handler: "PullRequestEventLabeled",
					Action:  "labeled",
				},
				{
					Handler: "PullRequestEventLocked",
					Action:  "locked",
				},
				{
					Handler: "PullRequestEventOpened",
					Action:  "opened",
				},
				{
					Handler: "PullRequestEventReadyForReview",
					Action:  "ready_for_review",
				},
				{
					Handler: "PullRequestEventReopened",
					Action:  "reopened",
				},
				{
					Handler: "PullRequestEventReviewRequestRemoved",
					Action:  "review_request_removed",
				},
				{
					Handler: "PullRequestEventReviewRequested",
					Action:  "review_requested",
				},
				{
					Handler: "PullRequestEventSynchronize",
					Action:  "synchronize",
				},
				{
					Handler: "PullRequestEventUnassigned",
					Action:  "unassigned",
				},
				{
					Handler: "PullRequestEventUnlabeled",
					Action:  "unlabeled",
				},
				{
					Handler: "PullRequestEventUnlocked",
					Action:  "unlocked",
				},
			},
			HasActions: true,
		},
		{
			Event: "PullRequestReviewEvent",
			Name:  "pull_request_review",
			Actions: []Action{
				{
					Handler: "PullRequestReviewEventSubmitted",
					Action:  "submitted",
				},
				{
					Handler: "PullRequestReviewEventEdited",
					Action:  "edited",
				},
				{
					Handler: "PullRequestReviewEventDismissed",
					Action:  "dismissed",
				},
			},
			HasActions: true,
		},
		{
			Event: "PullRequestReviewCommentEvent",
			Name:  "pull_request_review_comment",
			Actions: []Action{
				{
					Handler: "PullRequestReviewCommentEventCreated",
					Action:  "created",
				},
				{
					Handler: "PullRequestReviewCommentEventEdited",
					Action:  "edited",
				},
				{
					Handler: "PullRequestReviewCommentEventDeleted",
					Action:  "deleted",
				},
			},
			HasActions: true,
		},
		{
			Event:      "PushEvent",
			Name:       "push",
			HasActions: false,
		},
		{
			Event: "ReleaseEvent",
			Name:  "release",
			Actions: []Action{
				{
					Handler: "ReleaseEventPublished",
					Action:  "published",
				},
				{
					Handler: "ReleaseEventUnpublished",
					Action:  "unpublished",
				},
				{
					Handler: "ReleaseEventCreated",
					Action:  "created",
				},
				{
					Handler: "ReleaseEventEdited",
					Action:  "edited",
				},
				{
					Handler: "ReleaseEventDeleted",
					Action:  "deleted",
				},
				{
					Handler: "ReleaseEventPreReleased",
					Action:  "prereleased",
				},
				{
					Handler: "ReleaseEventReleased",
					Action:  "released",
				},
			},
			HasActions: true,
		},
		{
			Event:      "RepositoryDispatchEvent",
			Name:       "repository_dispatch",
			HasActions: false,
		},
		{
			Event: "RepositoryRulesetEvent",
			Name:  "repository_ruleset",
			Actions: []Action{
				{
					Handler: "RepositoryRulesetEventCreated",
					Action:  "created",
				},
				{
					Handler: "RepositoryRulesetEventDeleted",
					Action:  "deleted",
				},
				{
					Handler: "RepositoryRulesetEventEdited",
					Action:  "edited",
				},
			},
			HasActions: true,
		},
		{
			Event: "RepositoryEvent",
			Name:  "repository",
			Actions: []Action{
				{
					Handler: "RepositoryEventCreated",
					Action:  "created",
				},
				{
					Handler: "RepositoryEventDeleted",
					Action:  "deleted",
				},
				{
					Handler: "RepositoryEventArchived",
					Action:  "archived",
				},
				{
					Handler: "RepositoryEventUnarchived",
					Action:  "unarchived",
				},
				{
					Handler: "RepositoryEventEdited",
					Action:  "edited",
				},
				{
					Handler: "RepositoryEventRenamed",
					Action:  "renamed",
				},
				{
					Handler: "RepositoryEventTransferred",
					Action:  "transferred",
				},
				{
					Handler: "RepositoryEventPublicized",
					Action:  "publicized",
				},
				{
					Handler: "RepositoryEventPrivatized",
					Action:  "privatized",
				},
			},
			HasActions: true,
		},
		{
			Event: "RepositoryVulnerabilityAlertEvent",
			Name:  "repository_vulnerability_alert",
			Actions: []Action{
				{
					Handler: "RepositoryVulnerabilityAlertEventCreate",
					Action:  "create",
				},
				{
					Handler: "RepositoryVulnerabilityAlertEventDismiss",
					Action:  "dismiss",
				},
				{
					Handler: "RepositoryVulnerabilityAlertEventResolve",
					Action:  "resolve",
				},
			},
			HasActions: true,
		},
		{
			Event: "StarEvent",
			Name:  "star",
			Actions: []Action{
				{
					Handler: "StarEventCreated",
					Action:  "created",
				},
				{
					Handler: "StarEventDeleted",
					Action:  "deleted",
				},
			},
			HasActions: true,
		},
		{
			Event:      "StatusEvent",
			Name:       "status",
			HasActions: false,
		},
		{
			Event: "TeamEvent",
			Name:  "team",
			Actions: []Action{
				{
					Handler: "TeamEventCreated",
					Action:  "created",
				},
				{
					Handler: "TeamEventDeleted",
					Action:  "deleted",
				},
				{
					Handler: "TeamEventEdited",
					Action:  "edited",
				},
				{
					Handler: "TeamEventAddedToRepository",
					Action:  "added_to_repository",
				},
				{
					Handler: "TeamEventRemovedFromRepository",
					Action:  "removed_from_repository",
				},
			},
			HasActions: true,
		},
		{
			Event:      "TeamAddEvent",
			Name:       "team_add",
			HasActions: false,
		},
		{
			Event:      "WatchEvent",
			Name:       "watch",
			HasActions: false,
		},
		{
			Event: "WorkflowJobEvent",
			Name:  "workflow_job",
			Actions: []Action{
				{
					Handler: "WorkflowJobEventQueued",
					Action:  "queued",
				},
				{
					Handler: "WorkflowJobEventInProgress",
					Action:  "in_progress",
				},
				{
					Handler: "WorkflowJobEventCompleted",
					Action:  "completed",
				},
			},
			HasActions: true,
		},
		{
			Event:      "WorkflowDispatchEvent",
			Name:       "workflow_dispatch",
			HasActions: false,
		},
		{
			Event: "WorkflowRunEvent",
			Name:  "workflow_run",
			Actions: []Action{
				{
					Handler: "WorkflowRunEventRequested",
					Action:  "requested",
				},
				{
					Handler: "WorkflowRunEventCompleted",
					Action:  "completed",
				},
			},
			HasActions: true,
		},
	},
}
