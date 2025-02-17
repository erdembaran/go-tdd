package main

import "testing"

func TestSum(t *testing.T) {

	numbers := [4]int{2,4,6,8}

	got := Sum(numbers)
	want := 20

	if got != want {
		t.Errorf("got %v, but want %v", got, want)
	}

}