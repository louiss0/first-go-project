package main

import "testing"

func TestHello(testUtils *testing.T) {

	assertCorrectMessage := func(t testing.TB, got string, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	testUtils.Run("saying hello to people", func(t *testing.T) {

		got := Hello("Chris", "")

		want := "Hello, Chris"

		assertCorrectMessage(t, got, want)

	})

	testUtils.Run("says 'Hello World' when an empty string is supplied",
		func(testUtils *testing.T) {

			got := Hello("", "")
			want := "Hello, World"

			assertCorrectMessage(testUtils, got, want)

		},
	)

	testUtils.Run("in Spanish", func(testUtils *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(testUtils, got, want)
	})

}

func TestAdd(testUtils *testing.T) {

	sum := Add(2, 2)

	expected := 4

	if sum != expected {

		testUtils.Errorf("expected %d but got %d", expected, sum)

	}

}
