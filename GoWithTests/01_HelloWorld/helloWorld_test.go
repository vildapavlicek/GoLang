package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("Got '%s' want '%s'", got, want)
		}
	}

	t.Run("test chris input", func(t *testing.T) {
		got := hello("Chris")
		want := "Hello, Chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("test empty input", func(t *testing.T) {
		got := hello("")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})
}