package main

import (
	"fmt"
)

// F implements the function that takes n as a parameter,
// and returns a slice of size n, such that the value of the element at index i is the factorial of i if i is even
// or the sum of numbers from 1 to i otherwise.
// i starts from 1.
func F(n int) []int {
	if n <= 0 {
		return []int{}
	}

	ret := make([]int, n)

	// The problem statement uses 1-based indexing for 'i'.
	for i := 1; i <= n; i++ {
		// Go slices are 0-indexed, so we store the result for 'i' at index 'i-1'.
		index := i - 1

		if i%2 == 0 {
			// i is even: calculate factorial of i.
			factorial := 1
			for j := 1; j <= i; j++ {
				factorial *= j
			}
			ret[index] = factorial
		} else {
			// i is odd: calculate sum from 1 to i.
			// The formula for the sum of an arithmetic series (n * (n + 1) / 2) is more efficient.
			sum := i * (i + 1) / 2
			ret[index] = sum
		}
	}
	return ret
}

func main() {
	// Example from the Python docstring
	result := F(5)
	fmt.Println(result)

	// Another test case
	result2 := F(7)
	fmt.Println(result2)
}
