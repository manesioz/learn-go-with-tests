package main

import "testing"

func TestRepeat(t *testing.T) {

	t.Run("calling the repeated function", func(t *testing.T) {
		repeated := Repeat("a", 3)
		expected := "aaa"

		if repeated != expected {
			t.Errorf("Error: got %q, but want %q", repeated, expected)
		}
	})
}
