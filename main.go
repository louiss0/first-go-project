package main

import (
	"errors"
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

func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

type Bitcoin int

func (bitcoin Bitcoin) String() string {

	return fmt.Sprintf("%d BTC", bitcoin)

}

type Wallet struct {
	balance Bitcoin
}

var InsufficientFundsError = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {

	return lo.If(
		amount > w.balance,
		InsufficientFundsError,
	).ElseF(func() error {

		w.balance -= amount

		return nil
	})

}

func (w *Wallet) Deposit(amount Bitcoin) {

	w.balance += amount

}

func (w Wallet) Balance() Bitcoin {

	return w.balance

}

type Dictionary map[string]string

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it already exists")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) AddWord(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) {

	delete(d, word)
}

func main() {

	fmt.Println(Hello("World", ""))

}
