package main

import (
	"fmt"
	"math"
)

func Factorize(n int) []int {
	fact := []int{}
	i := 2
	// Calculate the limit once before the loop, matching Python's behavior
	limit := int(math.Sqrt(float64(n))) + 1
	for i <= limit {
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

func main() {
	// Test cases from docstring
	fmt.Println(Factorize(8))
	fmt.Println(Factorize(25))
	fmt.Println(Factorize(70))
}