package main

import (
	"math"
)

func Factorize(n int) []int {
	fact := []int{}
	i := 2
	limit := int(math.Sqrt(float64(n)) + 1)
	for i <= limit {
		if n%i == 0 {
			fact = append(fact, i)
			n /= i
		} else {
			i += 1
		}
	}

	if n > 1 {
		fact = append(fact, n)
	}
	return fact
}