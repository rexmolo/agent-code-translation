package main

import "fmt"

func FlipCase(str string) string {
	// For a given string, flip lowercase characters to uppercase and uppercase to lowercase.
	// Example: FlipCase("Hello") returns "hELLO"
	result := make([]byte, len(str))
	for i, c := range []byte(str) {
		// XOR with 32 to flip between uppercase and lowercase ASCII letters
		if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') {
			result[i] = c ^ 32
		} else {
			result[i] = c
		}
	}
	return string(result)
}

func main() {
	fmt.Println(FlipCase("Hello"))
}