package main

import (
	"math"
	"testing"
)

func TestGeometry(t *testing.T) {
    checkArea := func(t testing.TB, shape Shape, want float64) {
        t.Helper()
        got := shape.Area()
        if math.Abs(got - want) > 0.01 {
            t.Errorf("got %.2f want %.2f", got, want)
        }
    }

    t.Run("calculate rectangle area", func(t *testing.T) {
        rectangle := Rectangle{Width: 10, Height: 5}
        checkArea(t, rectangle, 50.0)
    })

    t.Run("calculate circle area", func(t *testing.T) {
        circle := Circle{Radius: 5}
        checkArea(t, circle, 78.54)
    })
} 