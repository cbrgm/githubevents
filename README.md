# githubevents

**GitHub webhook events toolset for Go**

[![Go Report Card](https://goreportcard.com/badge/github.com/cbrgm/githubevents)](https://goreportcard.com/report/github.com/cbrgm/githubevents)
[![release](https://img.shields.io/github/release-pre/cbrgm/githubevents.svg)](https://github.com/cbrgm/githubevents/releases)
[![license](https://img.shields.io/badge/Docs-pkg.go.dev-blue.svg)](https://pkg.go.dev/github.com/cbrgm/githubevents/githubevents)
[![license](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/cbrgm/githubevents/blob/master/LICENSE)
![GitHub stars](https://img.shields.io/github/stars/cbrgm/githubevents.svg?label=github%20stars)

`githubevents` is a webhook events toolset for Go inspired by ***[octokit/webhooks.js](https://github.com/octokit/webhooks.js)***. 

This library makes use of [google/go-github](https://github.com/google/go-github) and provides functionality to register callbacks for Github events and their actions, so that you can easily execute your own logic in response to webhook events.

---

* [Usage](#usage)
* [API](#api)
    + [OnBeforeAny](#onbeforeany)
    + [OnAfterAny](#onafterany)
    + [OnError](#onerror)
* [Supported Webhooks Events](#supported-webhooks-events)
* [Contributing & License](#contributing---license)

## Usage

```go
import (
    "github.com/cbrgm/githubevents/githubevents"
)
```

Create a new `githubevents.EventHandler`, register callbacks and start a http server.

```go
func main() {
	// create a new event handler
	handle := githubevents.New("secretkey")
	
	// add callbacks
	handle.OnIssueCommentCreated(
            func(deliveryID string, eventName string, event *github.IssueCommentEvent) error {
                fmt.Printf("%s made a comment!", event.Sender.Login)
                return nil
            },
        )
	
	// start server
	
}
```

For more usage examples, please have a look at examples.

## API

Please refer to pkg.go.dev.

### OnBeforeAny

`OnBeforeAny` registers callbacks which are triggered before any event. Registered callbacks are executed in parallel in separate Goroutines.

```go
handle := githubevents.New("secretkey")
handle.OnBeforeAny(
    func(deliveryID string, eventName string, event interface{}) error {
        fmt.Printf("%s event received!", eventName)
        // do something
        return nil
    },
)
// ...
```

### OnAfterAny

`OnAfterAny` registers callbacks which are triggered after any event. Registered callbacks are executed in parallel in separate Goroutines.

```go
handle := githubevents.New("secretkey")
handle.OnAfterAny(
    func(deliveryID string, eventName string, event interface{}) error {
        fmt.Printf("%s event received!", eventName)
        // do something
        return nil
    },
)
// ...
```

### OnError

`OnError` registers callbacks which are triggered whenever an error occurs. These callbacks can be used for additional error handling, debugging or logging purposes. Registered callbacks are executed in parallel in separate Goroutines.

```go
handle := githubevents.New("secretkey")
handle.OnError(
	func(deliveryID string, eventName string, event interface{}, err error) error {
		fmt.Printf("received error %s", err)
		// additional error handling ...
		return err
	}, 
)
// ...
```

## Supported Webhooks Events



* ***[branch_protection_rule](https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#branch_protection_rule)***

* ***[check_run](https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#check_run)***

* ***[check_suite](https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#check_suite)***

* ***[commit_comment](https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#commit_comment)***

* ***[create](https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#create)***

* ***[delete](https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#delete)***

* ***[deploy_key](https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#deploy_key)***

* ***[deployment](https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#deployment)***

* ***[deployment_status](https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#deployment_status)***

* ***[discussion](https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#discussion)***

* ***[fork](https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#fork)***

* ***[github_app_authorization](https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#github_app_authorization)***

* ***[gollum](https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#gollum)***

* ***[installation](https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#installation)***

* ***[installation_repositories](https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#installation_repositories)***

* ***[issue_comment](https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#issue_comment)***

* ***[issues](https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#issues)***

* ***[label](https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#label)***

* ***[marketplace_purchase](https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#marketplace_purchase)***

* ***[member](https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#member)***

* ***[membership](https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#membership)***

* ***[meta](https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#meta)***

* ***[milestone](https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#milestone)***

* ***[organization](https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#organization)***

* ***[org_block](https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#org_block)***

* ***[package](https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#package)***

* ***[page_build](https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#page_build)***

* ***[ping](https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#ping)***

* ***[project](https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#project)***

* ***[project_card](https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#project_card)***

* ***[project_column](https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#project_column)***

* ***[public](https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#public)***

* ***[pull_request](https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#pull_request)***

* ***[pull_request_review](https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#pull_request_review)***

* ***[pull_request_review_comment](https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#pull_request_review_comment)***

* ***[push](https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#push)***

* ***[release](https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#release)***

* ***[repository_dispatch](https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#repository_dispatch)***

* ***[repository](https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#repository)***

* ***[repository_vulnerability_alert](https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#repository_vulnerability_alert)***

* ***[star](https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#star)***

* ***[status](https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#status)***

* ***[team](https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#team)***

* ***[team_add](https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#team_add)***

* ***[watch](https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#watch)***

* ***[workflow_job](https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#workflow_job)***

* ***[workflow_dispatch](https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#workflow_dispatch)***

* ***[workflow_run](https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#workflow_run)***


## Contributing & License

Feel free to submit changes! See
the [Contributing Guide](https://github.com/cbrgm/contributing/blob/master/CONTRIBUTING.md). This project is open-source
and is developed under the terms of the [MIT License](https://github.com/cbrgm/githubevents/blob/master/LICENSE).
