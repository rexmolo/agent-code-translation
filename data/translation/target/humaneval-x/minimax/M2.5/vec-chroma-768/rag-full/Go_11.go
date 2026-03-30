package main

import "fmt"

func StringXor(a string, b string) string {
	// Determine the length (minimum of both strings)
	length := len(a)
	if len(b) < length {
		length = len(b)
	}

	// Build result as a slice of bytes
	result := make([]byte, length)

	for i := 0; i < length; i++ {
		// XOR: if bits are same, result is '0', otherwise '1'
		if a[i] == b[i] {
			result[i] = '0'
		} else {
			result[i] = '1'
		}
	}

	return string(result)
}

func main() {
	// Test cases
	fmt.Println(StringXor("010", "110")) // Output: 100
}