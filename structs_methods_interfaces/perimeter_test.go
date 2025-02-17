package main

import "testing"

func TestPerimeter(t *testing.T) {
	 rectangle := Rectangle{14.00, 28.00}

	 got := Perimeter(rectangle)
	 want := 84.00

	 if got != want {
		t.Errorf("got %v, but want %v", got, want)
	 }
}