package main

import (
	"strings"
)

func StringXor(a string, b string) string {
	xor := func(i, j byte) byte {
		if i == j {
			return '0'
		}
		return '1'
	}

	result := make([]byte, len(a))
	for i := 0; i < len(a); i++ {
		result[i] = xor(a[i], b[i])
	}

	return string(result)
}

// For testing - uncomment to run
/*
func main() {
	// Test the function
	result := StringXor("010", "110")
	println(result) // Expected: 100
}
*/
