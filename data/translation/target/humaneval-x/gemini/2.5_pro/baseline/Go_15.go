package main

import (
	"fmt"
	"strconv"
	"strings"
)

// StringSequence returns a string containing space-delimited numbers starting from 0 upto n inclusive.
func StringSequence(n int) string {
	if n < 0 {
		return ""
	}

	// Pre-allocate a slice of strings with the exact size needed (n+1 elements for 0 to n).
	numbers := make([]string, n+1)

	// Populate the slice by converting each number from 0 to n to a string.
	for i := 0; i <= n; i++ {
		numbers[i] = strconv.Itoa(i)
	}

	// Join the string elements with a space delimiter.
	return strings.Join(numbers, " ")
}

func main() {
	// Mimic the Python docstring examples
	fmt.Printf("'%s'\n", StringSequence(0))
	fmt.Printf("'%s'\n", StringSequence(5))
}