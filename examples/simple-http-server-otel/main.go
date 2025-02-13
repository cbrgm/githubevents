package main

import (
	"context"
	"fmt"
	"github.com/cbrgm/githubevents/githubevents"
	"github.com/google/go-github/v69/github"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/trace"
	"net/http"
)

func main() {
	// Create a stdout exporter
	exporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		panic(err)
	}

	// Create a trace provider with the exporter
	tp := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
	)
	defer func() { _ = tp.Shutdown(context.Background()) }()

	// Set the global tracer provider
	otel.SetTracerProvider(tp)

	handle := githubevents.New("")

	handle.OnRepositoryEventCreated(func(ctx context.Context, deliveryID string, eventName string, event *github.RepositoryEvent) error {
		fmt.Printf("%s has created respository %s", *event.Sender.Login, *event.Repo.FullName)
		return nil
	})

	http.HandleFunc("/hook", func(w http.ResponseWriter, r *http.Request) {
		err := handle.HandleEventRequest(r)
		if err != nil {
			fmt.Println("error")
		}
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
