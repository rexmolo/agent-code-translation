package main

import "fmt"

func StringXor(a string, b string) string {
	// Determine the length to iterate over (using shorter length to match Python's zip behavior)
	n := len(a)
	if len(b) < n {
		n = len(b)
	}

	// Build result byte slice
	result := make([]byte, n)
	for i := 0; i < n; i++ {
		if a[i] == b[i] {
			result[i] = '0'
		} else {
			result[i] = '1'
		}
	}

	return string(result)
}

func main() {
	fmt.Println(StringXor("010", "110")) // Output: 100
}
