package main

import (
	"fmt"
)

func Hello(word string) string {

	return fmt.Sprintf("Hello %s", word)
}

func main() {

	fmt.Println(Hello("World"))

}
