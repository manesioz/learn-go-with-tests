package main

import (
	"testing"
)

//TestHelloWorld is a unit test func
func TestHelloWorld(t *testing.T) {

	t.Run("saying hello to an unspecified person", func(t *testing.T) {
		got := HelloWorld("")
		want := "Hello, World"

		if got != want {
			t.Errorf("Error: got %q, but want %q", got, want)
		}
	})

	t.Run("saying hello to a specified person", func(t *testing.T) {
		got := HelloWorld("Zack")
		want := "Hello, Zack"

		if got != want {
			t.Errorf("Error: got %q, but want %q", got, want)
		}
	})
}
