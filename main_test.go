package main

import (
	"slices"
	"testing"
)

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

	testUtils.Run("in French", func(testUtils *testing.T) {
		got := Hello("Étienne", "Spanish")
		want := "Hola, Étienne"
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

func TestRepeat(testUtils *testing.T) {

	repeated := Repeat("a", 5)

	expected := "aaaaa"

	if repeated != expected {

		testUtils.Errorf("expected %s but got %s", expected, repeated)

	}

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func TestSum(testUtils *testing.T) {

	testUtils.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	testUtils.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

}

func TestSumAll(t *testing.T) {

	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(testUtils *testing.T) {

	checkSums := func(testAndBenchmarkUtils testing.TB, got []int, want []int) {
		testAndBenchmarkUtils.Helper()
		if !slices.Equal(got, want) {
			testAndBenchmarkUtils.Errorf("got %v want %v", got, want)
		}
	}

	testUtils.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	testUtils.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		checkSums(t, got, want)
	})

}

func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 10.0)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
