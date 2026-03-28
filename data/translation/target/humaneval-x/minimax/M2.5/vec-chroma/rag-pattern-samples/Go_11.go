package main

import "fmt"

func StringXor(a string, b string) string {
	// Create a byte slice to store the result
	result := make([]byte, len(a))
	
	// Iterate through each character of both strings
	for i := 0; i < len(a); i++ {
		// XOR: if characters are the same, result is '0', otherwise '1'
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