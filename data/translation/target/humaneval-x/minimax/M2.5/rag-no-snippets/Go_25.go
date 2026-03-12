package main

import (
	"fmt"
	"math"
)

func Factorize(n int) []int {
	fact := []int{}
	i := 2
	for i <= int(math.Sqrt(float64(n)) + 1) {
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

func main() {
	// Test cases from docstring
	fmt.Println(Factorize(8))   // [2 2 2]
	fmt.Println(Factorize(25))  // [5 5]
	fmt.Println(Factorize(70)) // [2 5 7]
}
