package main

import "fmt"

func Strlen(str string) int {
	return len(str)
}

func main() {
	// Example usage
	fmt.Println(Strlen(""))     // Output: 0
	fmt.Println(Strlen("abc"))  // Output: 3
}