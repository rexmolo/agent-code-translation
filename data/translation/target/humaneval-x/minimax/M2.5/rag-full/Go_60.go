package main

import "fmt"

// SumToN sums numbers from 1 to n.
// Example: SumToN(30) = 465
// Example: SumToN(100) = 5050
// Example: SumToN(5) = 15
// Example: SumToN(10) = 55
// Example: SumToN(1) = 1
func SumToN(n int) int {
	return n * (n + 1) / 2
}

func main() {
	// Test cases
	fmt.Println(SumToN(30))  // 465
	fmt.Println(SumToN(100)) // 5050
	fmt.Println(SumToN(5))   // 15
	fmt.Println(SumToN(10))  // 55
	fmt.Println(SumToN(1))   // 1
}