package main

import "fmt"

// Add adds two integers x and y and returns their sum.
func Add(x int, y int) int {
	return x + y
}

func main() {
	// Example usage as per docstring tests
	fmt.Println(Add(2, 3))   // 5
	fmt.Println(Add(5, 7))  // 12
}
