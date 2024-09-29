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
		_, got := dictionary.Search("unknown")
		if got == nil {
			t.Fatal("expected to get an error")
		}
		assertErrors(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is a real dvalue"
		err := dictionary.Add(word, definition)
		assertErrors(t, err, nil)
		assertDefinitions(t, dictionary, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is a real value"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new test")

		assertErrors(t, err, ErrWordAlreadyExists)
		assertDefinitions(t, dictionary, word, definition)
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q but want %q ", got, want)
	}
}

func assertErrors(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}

func assertDefinitions(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Errorf("got %q, but want %q", dictionary[word], definition)
	}

	assertStrings(t, got, definition)
}
