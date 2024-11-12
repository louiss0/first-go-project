package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"reflect"
	"strings"
	"sync"
	"time"

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

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func Countdown(out io.Writer, sleeper Sleeper) {

	startCount := 1
	maxCount := 3
	finalWord := "Go!"

	lo.ForEach(
		lo.Reverse(lo.RangeFrom(startCount, maxCount)),
		func(item int, index int) {
			fmt.Fprintln(out, item)
			sleeper.Sleep()
		})

	fmt.Fprint(out, finalWord)
}

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {

	c.sleep(c.duration)

}

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	resultsChannel := make(chan result)

	for _, url := range urls {

		go func(value string) {

			resultsChannel <- result{value, wc(value)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultsChannel
		results[r.string] = r.bool
	}

	return results
}

var tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (winner string, error error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {

	ping := func(url string) chan struct{} {
		ch := make(chan struct{})
		go func() {
			http.Get(url)
			close(ch)
		}()
		return ch
	}

	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	walkValue := func(value reflect.Value) {

		walk(value.Interface(), fn)

	}

	switch val.Kind() {

	case reflect.String:
		fn(val.String())

	case reflect.Struct:
		lo.ForEach(
			lo.Range(val.NumField()),
			func(item int, _ int) {

				walkValue(val.Field(item))
			},
		)

	case reflect.Slice, reflect.Array:
		lo.ForEach(
			lo.Range(val.Len()),
			func(item int, index int) {

				walkValue(val.Index(item))

			},
		)

	case reflect.Map:
		lo.ForEach(
			val.MapKeys(),
			func(key reflect.Value, index int) {

				walkValue(val.MapIndex(key))

			},
		)

	case reflect.Chan:

		for {
			if v, ok := val.Recv(); ok {

				walkValue(v)
			} else {
				break
			}
		}

	case reflect.Func:
		lo.ForEach(
			val.Call(nil),
			func(item reflect.Value, index int) {

				walkValue(item)
			},
		)

	}

}

type Counter struct {
	mutex sync.Mutex
	value int
}

func (c *Counter) Inc() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())

		if err != nil {
			log.Print(err)
			return
		}

		fmt.Fprint(w, data)
	}

}

type Store interface {
	Fetch(ctx context.Context) (string, error)
}
type RomanNumeral struct {
	Value  int
	Symbol string
}

var allRomanNumerals = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(arabic int) string {

	var result strings.Builder

	lo.Reduce(
		allRomanNumerals,
		func(acc int, numeral RomanNumeral, _ int) int {

			itemCount := acc

			lo.ForEachWhile(
				lo.Range(itemCount),
				func(_ int, _ int) (goon bool) {

					if itemCount < numeral.Value {
						return false
					}

					result.WriteString(numeral.Symbol)

					itemCount -= numeral.Value

					return true

				})

			return itemCount

		}, arabic)

	return result.String()
}

// later..
func ConvertToArabic(roman string) int {

	type State struct {
		arabic int
		roman  string
	}

	state := lo.Reduce(
		allRomanNumerals,
		func(state State, numeral RomanNumeral, index int) State {

			arabic := state.arabic
			roman := state.roman

			for strings.HasPrefix(roman, numeral.Symbol) {
				arabic += numeral.Value
				roman = strings.TrimPrefix(roman, numeral.Symbol)
			}

			return State{arabic, roman}

		},
		State{
			arabic: 0,
			roman:  roman,
		},
	)

	return state.arabic

}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
