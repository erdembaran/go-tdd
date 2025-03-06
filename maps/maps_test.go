package main

import "testing"


func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unkown")
		want := "could not find the word you were looking for"

		if err == nil {
			// t.Fatal: stops the test if this fails
			t.Fatal("expected to get an error.")
		}
		assertStrings(t, err.Error(), want)

	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	
	if got != want {
		// t.Errorf: reports a test error
		t.Errorf("got %v want %v", got, want)
	}
}