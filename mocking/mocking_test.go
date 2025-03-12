package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}

	Countdown(buffer)

	got := buffer.String()
	want := "8"

	if got != want {
		t.Errorf("got %v, but want %v", got, want)
	}
} 