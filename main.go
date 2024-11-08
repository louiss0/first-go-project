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

func Sum(numbers [5]int) int {

	sum := 0

	for _, number := range numbers {

		sum += number

	}

	return sum

}

func main() {

	fmt.Println(Hello("World", ""))

}
