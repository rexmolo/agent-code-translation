package main

import (
	"fmt"
)

// GenerateIntegers returns the even single-digit numbers within the range defined by a and b.
// The logic mirrors the provided Python implementation.
func GenerateIntegers(a, b int) []int {
	// First, determine the ordered lower and upper bounds from the inputs.
	low, high := a, b
	if a > b {
		low, high = b, a
	}

	// The problem is about finding "even digits", so we are interested in the
	// intersection of the range [low, high] and the even digits [2, 4, 6, 8].

	// lower = max(2, min(a, b))
	lowerBound := 2
	if low > 2 {
		lowerBound = low
	}

	// upper = min(8, max(a, b))
	upperBound := 8
	if high < 8 {
		upperBound = high
	}

	// Pre-allocating slice capacity could be a small optimization, but for clarity,
	// we'll start with a zero-length slice and append.
	var result []int

	// To make the loop more efficient, we find the first even number
	// at or after lowerBound and then increment by 2.
	start := lowerBound
	if start%2 != 0 {
		// If lowerBound is odd, the first even number is the next one.
		start++
	}

	// Iterate from the adjusted start up to the upper bound, stepping by 2.
	for i := start; i <= upperBound; i += 2 {
		result = append(result, i)
	}

	// It's idiomatic in Go to return an empty, non-nil slice for empty collections.
	if result == nil {
		return []int{}
	}

	return result
}

func main() {
	// Example usage based on the Python docstring
	fmt.Println(GenerateIntegers(2, 8))
	fmt.Println(GenerateIntegers(8, 2))
	fmt.Println(GenerateIntegers(10, 14))
}
