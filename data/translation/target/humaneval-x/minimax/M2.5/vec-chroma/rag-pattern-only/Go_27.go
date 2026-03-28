package main

import "fmt"

func FlipCase(str string) string {
	// Convert string to rune slice for proper Unicode handling
	runes := []rune(str)
	for i, r := range runes {
		switch {
		case r >= 'a' && r <= 'z':
			// Convert lowercase to uppercase
			runes[i] = r - 'a' + 'A'
		case r >= 'A' && r <= 'Z':
			// Convert uppercase to lowercase
			runes[i] = r - 'A' + 'a'
		}
	}
	return string(runes)
}

func main() {
	// Test the function
	result := FlipCase("Hello")
	fmt.Println(result) // Output: hELLO
}
