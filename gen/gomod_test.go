package main

import (
	"os"
	"path/filepath"
	"testing"
)

func writeGoMod(t *testing.T, content string) string {
	t.Helper()
	dir := t.TempDir()
	p := filepath.Join(dir, "go.mod")
	if err := os.WriteFile(p, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
	return p
}

func TestGoGithubImportPath(t *testing.T) {
	p := writeGoMod(t, `module github.com/cbrgm/githubevents/v2

go 1.25.0

require (
	github.com/google/go-github/v89 v89.0.0
	golang.org/x/sync v0.22.0
)
`)
	got, err := goGithubImportPath(p)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if want := "github.com/google/go-github/v89/github"; got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestGoGithubImportPathMissing(t *testing.T) {
	p := writeGoMod(t, "module x\n\ngo 1.25.0\n")
	if _, err := goGithubImportPath(p); err == nil {
		t.Fatal("expected error when no go-github require present")
	}
}
