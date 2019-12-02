// Функция, която създава map reducer функция за int аргументи с подадени
// map функция, reduce функция и първоначална стойност initial за reduce функцията.

// func MapReducer(mapper func (int) int, reducer func (int, int) int, initial int) func (...int) int
// Която може да бъде използвана по следния начин

// powerSum := MapReducer(
//     func (v int) int { return v * v },
//     func (a, v int) int { return a + v },
//     0,
// )

// powerSum(1, 2, 3, 4) // 30
// Аргументите на reducer трябва да се подават от ляво на дясно. Иначе казано - напишете left-fold във вашата имплементация.

package main

import "fmt"

func MapReducer(mapper func(int) int, reducer func(int, int) int, initial int) func(...int) int {
	return func(args ...int) int {
		result := 0
		for _, v := range args {
			result = reducer(result, mapper(v))
		}
		fmt.Println(result)
		var mapped = mapper(initial)
		return mapped
	}
}

func main() {
	powerSum := MapReducer(
		func(v int) int { return v * v },
		func(a, v int) int { return a + v },
		0,
	)

	powerSum(1, 2, 3, 4) // 30
}
