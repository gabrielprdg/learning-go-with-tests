package main

import "testing"

// helper function
func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHello(t *testing.T) {
	got := Hello("Gabi", "")
	want := "Hello, Gabi"

	assertCorrectMessage(t, got, want)
}

func TestHelloCheering(t *testing.T) {
	got := Hello("Chris", "")
	want := "Hello, Chris"

	assertCorrectMessage(t, got, want)
}

// with subtests
func TestHelloWithAnEmptyString(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Gabriel", "")
		want := "Hello, Gabriel"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying 'Hello, World' when a empty string was supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Carla", "Spanish")
		want := "Hola, Carla"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Ilson", "French")
		want := "Bonjour, Ilson"

		assertCorrectMessage(t, got, want)
	})
}
