package main

import "testing"

func TestHello(t *testing.T) {
	// got := Hello("All")
	// want := "Hello, All"
	// if got != want {
	// 	t.Errorf("err: %q is got, %q is want.", got, want)
	// }

	assertCorrectMessage := func(t *testing.T, got, want string) {
		if got != want {
			t.Errorf("got is %q, want is %q", got, want)
		}
	}

	t.Run("Say hello name in default", func(t *testing.T) {
		got := Hello("TestUser", "English")
		want := "Hello, TestUser"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Say hello empty case, return world", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("world", "Spanish")
		want := "Hola, world"
		assertCorrectMessage(t, got, want)
	})
}
