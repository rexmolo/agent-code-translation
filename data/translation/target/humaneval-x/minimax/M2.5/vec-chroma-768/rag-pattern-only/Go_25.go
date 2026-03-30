package main

import (
	"math"
)

// Factorize returns a list of prime factors of the given integer in ascending order.
// Each factor appears the number of times corresponding to how many times it appears
// in the factorization. The product of all factors equals the input number.
//
// Examples:
//   Factorize(8)  -> [2, 2, 2]
//   Factorize(25) -> [5, 5]
//   Factorize(70) -> [2, 5, 7]
func Factorize(n int) []int {
	fact := []int{}
	i := 2
	for i <= int(math.Sqrt(float64(n))+1) {
		if n%i == 0 {
			fact = append(fact, i)
			n /= i
		} else {
			i++
		}
	}

	if n > 1 {
		fact = append(fact, n)
	}
	return fact
}
