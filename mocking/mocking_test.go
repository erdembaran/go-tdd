package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}

	Countdown(buffer, spySleeper)

	got := buffer.String()

	// what you see is what you get
	want := `8
7
6
5
4
3
2
1
Go!`

	if got != want {
		t.Errorf("got %v, but want %v", got, want)
	}

	if spySleeper.Calls != 8 {
		t.Errorf("not enough calls to sleeper, want 8 got %d", spySleeper.Calls)
	}
} 


    
