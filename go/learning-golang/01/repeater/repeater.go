// Функция, която приема string s и разделител sep и връща функция, която построява повторенията определен брой пъти и ги връща.

// func Repeater(s, sep string) func (int) string
// Която може да се използва по следния начин

// Repeater("foo", ":")(3) // foo:foo:foo

package main

import (
	"fmt"
	"strings"
)

func Repeater(str string, separator string) func(number int) {
	return func(number int) {
		values := make([]string, number)

		for i := 0; i < number; i++ {
			values[i] = str
		}
		fmt.Print(strings.Join(values[:], separator))
	}
}

func main() {
	Repeater("foo", ":")(3)
}
