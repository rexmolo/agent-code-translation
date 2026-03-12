package main

import "fmt"

// Add two numbers x and y.
// Example:
//
//	Add(2, 3) // returns 5
//	Add(5, 7) // returns 12
func Add(x int, y int) int {
	return x + y
}

// main function to demonstrate the Add function.
func main() {
	fmt.Println("Adding 2 and 3:", Add(2, 3))
	fmt.Println("Adding 5 and 7:", Add(5, 7))
}
