package main

import "fmt"

// Strlen returns the length of the given string.
func Strlen(str string) int {
	return len(str)
}

func main() {
	// Example usage demonstrating the function
	fmt.Println(Strlen(""))   // Output: 0
	fmt.Println(Strlen("abc")) // Output: 3
}