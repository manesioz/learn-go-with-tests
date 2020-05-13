package main

import (
	"reflect"
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := "could not find the word you were looking for"

		if err == nil {
			t.Fatal("expected to get an error.")
		}

		assertStrings(t, err.Error(), want)
	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("word contained in dict", func(t *testing.T) {
		_, err := dictionary.Add("test", "same key different val, should NOT replace")

		want := "Error: Word already exists"
		if err == nil {
			t.Fatal("expected to get an error.")
		}
		assertStrings(t, err.Error(), want)
	})

	t.Run("word not contained in dict", func(t *testing.T) {
		got, _ := dictionary.Add("random_key", "random value")
		want := dictionary
		assertDictionaries(t, got, want)

	})
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertDictionaries(t *testing.T, got, want Dictionary) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %#v want %#v", got, want)
	}
}
