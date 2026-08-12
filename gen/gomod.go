package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
)

var goGithubModuleRE = regexp.MustCompile(`github\.com/google/go-github/v\d+`)

// goGithubImportPath reads gomodPath and returns the go-github package import
// path (module path + "/github") derived from the module's require directive.
// This keeps the go-github major version defined in exactly one place.
func goGithubImportPath(gomodPath string) (string, error) {
	data, err := os.ReadFile(gomodPath)
	if err != nil {
		return "", fmt.Errorf("read %s: %w", gomodPath, err)
	}
	matches := goGithubModuleRE.FindAllString(string(data), -1)
	uniq := map[string]struct{}{}
	for _, m := range matches {
		uniq[m] = struct{}{}
	}
	switch len(uniq) {
	case 0:
		return "", fmt.Errorf("no github.com/google/go-github require found in %s", gomodPath)
	case 1:
		return matches[0] + "/github", nil
	default:
		uniqueVersions := make([]string, 0, len(uniq))
		for v := range uniq {
			uniqueVersions = append(uniqueVersions, v)
		}
		sort.Strings(uniqueVersions)
		return "", fmt.Errorf("multiple go-github module versions found in %s: %v", gomodPath, uniqueVersions)
	}
}
