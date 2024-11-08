package main

import (
	"fmt"
	"math"
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

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}
func main() {

	fmt.Println(Hello("World", ""))

}
