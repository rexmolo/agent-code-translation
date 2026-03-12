package main

import (
	"fmt"
	"strings"
)

// StringXor performs a binary XOR operation on two input strings a and b,
// which are assumed to consist only of '1's and '0's. It returns the
// result as a string. The operation is performed up to the length of the
// shorter string, mimicking Python's zip behavior.
func StringXor(a string, b string) string {
	n := len(a)
	if len(b) < n {
		n = len(b)
	}

	var result strings.Builder
	result.Grow(n) // Pre-allocate memory for efficiency

	for i := 0; i < n; i++ {
		if a[i] == b[i] {
			result.WriteByte('0')
		} else {
			result.WriteByte('1')
		}
	}

	return result.String()
}

func main() {
	// Example from the Python docstring
	fmt.Println(StringXor("010", "110"))

	// Additional test cases
	fmt.Println(StringXor("1", "1"))
	fmt.Println(StringXor("0", "1"))
	fmt.Println(StringXor("10101", "00110"))
}
