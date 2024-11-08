package main

import (
	"fmt"
)

const englishHello = "Hello"

func Hello(word string) string {

	if word == "" {
		word = "World"
	}

	return fmt.Sprintf("%s, %s", englishHello, word)
}

func main() {

	fmt.Println(Hello("World"))

}
