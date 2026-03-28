package main

import "fmt"

func FlipCase(str string) string {
	// Flip case: lowercase -> uppercase, uppercase -> lowercase
	result := make([]rune, len(str))
	for i, c := range str {
		if c >= 'A' && c <= 'Z' {
			result[i] = c + 32 // Convert to lowercase
		} else if c >= 'a' && c <= 'z' {
			result[i] = c - 32 // Convert to uppercase
		} else {
			result[i] = c // Keep non-letter characters unchanged
		}
	}
	return string(result)
}

func main() {
	// Test the function
	fmt.Println(FlipCase("Hello")) // Output: hELLO
	fmt.Println(FlipCase("hELLO")) // Output: Hello
	fmt.Println(FlipCase("123"))  // Output: 123
	fmt.Println(FlipCase("AbC"))  // Output: aBc
}
