package main

var webhookMarkdownTemplate = `

{{ range $_, $webhook := .Webhooks }}
* ***[{{ $webhook.Name }}](https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#{{ $webhook.Name }})***
{{ end }}
`
