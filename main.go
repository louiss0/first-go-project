package main

import (
	"fmt"

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

func main() {

	fmt.Println(Hello("World", ""))

}
