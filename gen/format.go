package main

import (
	"bytes"
	gofmt "go/format"
	"io"
)

// format formats a template using go/format.
func format(in io.Reader) (io.Reader, error) {
	src, err := io.ReadAll(in)
	if err != nil {
		return nil, err
	}
	out, err := gofmt.Source(src)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(out), nil
}
