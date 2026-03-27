package main

import (
	"math"
)

func Factorize(n int) []int {
	var fact []int
	i := 2
	for i <= int(math.Sqrt(float64(n)) + 1) {
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