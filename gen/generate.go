package main

import (
	"bytes"
	_ "embed"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"text/template"
)

//go:embed template_markdown.go.tmpl
var webhookMarkdownTemplate string

//go:embed template_webhook_event.go.tmpl
var webhookEventTemplate string

//go:embed template_webhook_event_tests.go.tmpl
var webhookTestsTemplate string

//go:embed template_webhook_event_types.go.tmpl
var webhookTypesTemplate string

func main() {
	outputDir := flag.String("output", "githubevents", "output directory")
	docs := flag.Bool("docs", false, "generate markdown docs")

	flag.Parse()

	if err := run(*outputDir, *docs); err != nil {
		fmt.Fprintln(os.Stderr, "generate:", err)
		os.Exit(1)
	}
}

func run(outputDir string, docs bool) error {
	if outputDir == "" {
		return fmt.Errorf("output directory is empty")
	}

	// when -docs is set, create a list of all supported markdown events as yaml on stdout
	// todo(cbrgm): clean this up a little bit
	if docs {
		return ExecuteMarkdownTemplate("", webhookMarkdownTemplate, params)
	}

	imp, err := goGithubImportPath("go.mod")
	if err != nil {
		return err
	}
	params.GoGithubImport = imp

	out := filepath.Clean(outputDir)
	if err := os.MkdirAll(out, os.ModePerm); err != nil {
		return fmt.Errorf("create output dir: %w", err)
	}

	// create events.go
	if err := ExecuteWebhookEventTemplate(filepath.Join(out, "events"), params); err != nil {
		return err
	}

	// create individual files for each webhook event type
	// webhook events_*.go files and events_*_test.go files are generated
	for _, param := range params.Webhooks {
		fileName := "events_" + param.Name
		outFile := filepath.Join(out, fileName)
		err := ExecuteWebhookEventTypesTemplate(outFile, TemplateParameters{
			GoGithubImport: imp,
			Webhooks:       []GithubWebhooks{param},
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func ExecuteWebhookEventTemplate(file string, data any) error {
	err := ExecuteTemplate(file+".go", webhookEventTemplate, data)
	if err != nil {
		return err
	}
	return nil
}

func ExecuteWebhookEventTypesTemplate(file string, data any) error {
	err := ExecuteTemplate(file+".go", webhookTypesTemplate, data)
	if err != nil {
		return err
	}
	err = ExecuteTemplate(file+"_test.go", webhookTestsTemplate, data)
	if err != nil {
		return err
	}
	return nil
}

// ExecuteWebhookEventTypesTemplate renders the named template and writes to io.Writer wr.
func ExecuteTemplate(file, tmpl string, data any) error {
	wr := os.Stdout
	if output := file; output != "" {
		wri, err := os.Create(output)
		if err != nil {
			return err
		}
		wr = wri
		defer func() {
			_ = wr.Close()
		}()
	}

	buf := new(bytes.Buffer)

	t, err := template.New("").Parse(tmpl)
	if err != nil {
		return err
	}

	err = t.ExecuteTemplate(buf, "", data)
	if err != nil {
		return err
	}

	src, err := format(buf)
	if err != nil {
		return err
	}
	_, err = io.Copy(wr, src)
	return err
}

// ExecuteMarkdownTemplate renders the named template and writes to io.Writer wr.
func ExecuteMarkdownTemplate(_, tmpl string, data any) error {
	wr := os.Stdout
	buf := new(bytes.Buffer)
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		return err
	}
	err = t.ExecuteTemplate(buf, "", data)
	if err != nil {
		return err
	}
	_, err = io.Copy(wr, buf)
	return err
}
