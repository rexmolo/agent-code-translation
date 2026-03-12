package main

import (
	"fmt"
	"strconv"
)

// CircularShift performs a circular right shift on the digits of an integer.
// If shift is greater than the number of digits, the function returns the digits reversed.
// The result is returned as a string.
func CircularShift(x int, shift int) string {
	s := strconv.Itoa(x)
	n := len(s)

	// Python's slicing behavior for negative shifts with the original formula
	// effectively results in a no-op. To prevent panics from Go's stricter
	// slicing and to replicate the behavior, we handle negative shifts explicitly.
	if shift < 0 {
		return s
	}

	if shift > n {
		// Reverse the string as per the requirement
		runes := []rune(s)
		for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		return string(runes)
	}

	// Perform right circular shift for 0 <= shift <= n
	return s[n-shift:] + s[:n-shift]
}

func main() {
	fmt.Printf("\"%s\"\n", CircularShift(12, 1))
	fmt.Printf("\"%s\"\n", CircularShift(12, 2))

	// Additional test cases to match python docstring format
	// >>> circular_shift(123, 4) -> "321"
	fmt.Printf("\"%s\"\n", CircularShift(123, 4))

	// >>> circular_shift(-12, 1) -> "2-1"
	fmt.Printf("\"%s\"\n", CircularShift(-12, 1))
}
