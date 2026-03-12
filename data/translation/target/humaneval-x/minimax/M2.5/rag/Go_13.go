package main

import "fmt"

func GreatestCommonDivisor(a int, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	// Test cases
	fmt.Println(GreatestCommonDivisor(3, 5))   // Expected: 1
	fmt.Println(GreatestCommonDivisor(25, 15)) // Expected: 5
}
