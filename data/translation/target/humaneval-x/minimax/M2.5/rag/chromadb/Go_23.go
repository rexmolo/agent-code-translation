package main

import "fmt"

// Strlen returns the length of the given string.
func Strlen(str string) int {
	return len(str)
}

func main() {
	// Test cases
	fmt.Println(Strlen(""))    // 0
	fmt.Println(Strlen("abc")) // 3
}
