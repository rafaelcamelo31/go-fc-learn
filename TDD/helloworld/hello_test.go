package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Rafael", "")
		want := "Hello, Rafael"

		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})
	t.Run("in Portugues", func(t *testing.T) {
		got := Hello("Rafael", "Portugues")
		want := "Ola, Rafael"

		assertCorrectMessage(t, got, want)
	})
	t.Run("in Japanese", func(t *testing.T) {
		got := Hello("Rafael", "Japanese")
		want := "こんにちは, Rafael"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
