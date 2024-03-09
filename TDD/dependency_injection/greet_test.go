package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Rafael")

	got := buffer.String()
	want := "Hello, Rafael"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
