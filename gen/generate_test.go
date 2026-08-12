package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// run() generates events.go referencing the go.mod-derived import path.
func TestRunGeneratesEvents(t *testing.T) {
	orig, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := os.Chdir(orig); err != nil {
			t.Fatal(err)
		}
	}()

	// run() reads ./go.mod, so execute from the repo root.
	if err := os.Chdir(".."); err != nil {
		t.Fatal(err)
	}
	dir := t.TempDir()
	if err := run(dir, false); err != nil {
		t.Fatalf("run failed: %v", err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "events.go"))
	if err != nil {
		t.Fatalf("events.go not generated: %v", err)
	}
	if !strings.Contains(string(data), "google/go-github/v") {
		t.Fatal("generated events.go missing go-github import")
	}
}
