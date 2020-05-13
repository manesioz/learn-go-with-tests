package main

import (
	"math"
	"testing"
)

func TestGeometryFunctions(t *testing.T) {

	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("calling the area function for circle", func(t *testing.T) {

		circ := &Circle{10.0}
		checkArea(t, circ, 100.0*math.Pi)
	})

	t.Run("calling the area function for rectangle", func(t *testing.T) {

		rect := &Rectangle{10.0, 10.0}
		checkArea(t, rect, 100.0)
	})
}
