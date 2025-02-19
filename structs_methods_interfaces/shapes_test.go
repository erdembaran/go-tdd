package main

import (
	"math"
	"testing"
)

func TestRectangleArea(t* testing.T) {
	tests := []struct {
		name string
		rect Rectangle	
		want float64
	}{	
		{"Square (8x8)", Rectangle{Width: 8, Height: 8}, 64.00},
		{"Rectangle (4x6)", Rectangle{Width: 4, Height: 6}, 24.00},

	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.rect.Area()
			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}

		})
	}
}

func TestCircleArea(t *testing.T) {
	tests := []struct {
		name string
		circ Circle
		want float64
	}{
		{"Circle1", Circle{Radius: 4}, 50.26},
		{"Circle2", Circle{Radius: 8},	201.06},

	}

	const epsilon = 0.01

	for _, tc := range tests {
		t.Run(tc.name, func (t *testing.T) {
			got := tc.circ.Area()
			if math.Abs(got - tc.want) > epsilon {
                t.Errorf("got %v, want %v", got, tc.want)
            }
		})
	}


}