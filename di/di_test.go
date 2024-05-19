package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "世界")
	got := buffer.String()
	want := "Hello, 世界"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
