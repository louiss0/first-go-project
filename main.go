package main

import (
	"fmt"
	"strings"

	"github.com/samber/lo"
)

const englishHello = "Hello"
const spanishHello = "Hola"
const frenchHello = "Bonjour"

func Hello(word string, language string) string {

	word = lo.If(word == "", "World").Else(word)

	return lo.Switch[string, string](language).
		Case("Spanish", fmt.Sprintf("%s, %s", spanishHello, word)).
		Case("French", fmt.Sprintf("%s, %s", frenchHello, word)).
		Default(fmt.Sprintf("%s, %s", englishHello, word))

}

// Takes two integers and returns the sum of them.
func Add(x, y int) int {

	return x + y
}

func Repeat(word string, times int) string {

	return strings.Join(
		lo.RepeatBy(times,
			func(index int) string {
				return word
			},
		),
		"",
	)

}

func Sum(numbers []int) int {

	sum := 0

	for _, number := range numbers {

		sum += number

	}

	return sum

}

func SumAll(numbersToSum ...[]int) []int {
	return lo.Map(numbersToSum, func(numbers []int, _ int) int {

		return Sum(numbers)

	})
}

func SumAllTails(numbersToSum ...[]int) []int {
	return lo.Map(numbersToSum, func(numbers []int, _ int) int {

		return lo.IfF(
			len(numbers) == 0,
			func() int { return 0 },
		).
			ElseF(
				func() int { return Sum(numbers[1:]) },
			)
	})

}

func Perimeter(width float64, height float64) float64 {
	return 2 * (width + height)
}

func main() {

	fmt.Println(Hello("World", ""))

}
