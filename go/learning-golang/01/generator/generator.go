// Функция, създава "генератор" функция за int числа.

// func Generator(gen func (int) int, initial int) func() int
// На Generator се подава gen функция и първоначалната стойност в поредицата - initial.
// gen взима като аргумент предишно изчислената стойност и връща следващата.

// Примерна употреба

// counter := Generator(
//     func (v int) int { return v + 1 },
//     0,
// )
// power := Generator(
//     func (v int) int { return v * v },
//     2,
// )

// counter() // 0
// counter() // 1
// power() // 2
// power() // 4
// counter() // 2
// power() // 16
// counter() // 3
// power() // 256

package main

import "fmt"

func Generator(gen func(int) int, initial int) func() int {
	return func() int {
		fmt.Println(initial)
		initial = gen(initial)
		return initial
	}
}

func main() {
	counter := Generator(
		func(v int) int { return v + 1 },
		0,
	)
	power := Generator(
		func(v int) int { return v * v },
		2,
	)

	counter() // 0
	counter() // 1
	power()   // 2
	power()   // 4
	counter() // 2
	power()   // 16
	counter() // 3
	power()   // 256
}
