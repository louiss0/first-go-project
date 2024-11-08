package main

import "testing"

func TestHello(testUtils *testing.T) {

	got := Hello("Chris")

	want := "Hello, Chris"

	if got != want {

		testUtils.Errorf("got %q want %q", got, want)
	}

}
