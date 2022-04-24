package main

import (
	"bytes"
	"flag"
	"io"
	"os"
	"path/filepath"
	"text/template"
)

func main() {
	outputDir := flag.String("output", "githubevents", "output directory")
	docs := flag.Bool("docs", false, "generate markdown docs")

	flag.Parse()
	if *outputDir == "" {
		panic("output directory is empty")
	}

	// when -docs is set, create a list of all supported markdown events as yaml on stdout
	// todo(cbrgm): clean this up a little bit
	if *docs {
		err := ExecuteMarkdownTemplate("", webhookMarkdownTemplate, params)
		if err != nil {
			panic(err)
		}
		return
	}

	out := filepath.Join(".", *outputDir)
	err := os.MkdirAll(out, os.ModePerm)
	if err != nil {
		panic("failed to create output directory")
	}

	// create events.go
	err = ExecuteWebhookEventTemplate(filepath.Join(out, "events"), params)
	if err != nil {
		panic(err)
	}

	// create individual files for each webhook event type
	// webhook events_*.go files and events_*_test.go files are generated
	for _, param := range params.Webhooks {
		fileName := "events_" + param.Name
		outFile := filepath.Join(out, fileName)
		err := ExecuteWebhookEventTypesTemplate(outFile, TemplateParameters{
			Webhooks: []GithubWebhooks{param},
		})
		if err != nil {
			panic(err)
		}
	}
}

func ExecuteWebhookEventTemplate(file string, data interface{}) error {
	err := ExecuteTemplate(file+".go", webhookEventTemplate, data)
	if err != nil {
		return err
	}
	return nil
}

func ExecuteWebhookEventTypesTemplate(file string, data interface{}) error {
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
func ExecuteTemplate(file string, tmpl string, data interface{}) error {
	wr := os.Stdout
	if output := file; output != "" {
		wri, err := os.Create(output)
		if err != nil {
			return err
		}
		wr = wri
		defer wr.Close()
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
func ExecuteMarkdownTemplate(file string, tmpl string, data interface{}) error {
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
