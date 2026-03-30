package main

import "fmt"

// SumToN returns the sum of integers from 1 to n (inclusive).
// Example: SumToN(5) returns 15 (1+2+3+4+5)
func SumToN(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

func main() {
	// Test cases based on docstring examples
	fmt.Println(SumToN(30))  // 465
	fmt.Println(SumToN(100)) // 5050
	fmt.Println(SumToN(5))   // 15
	fmt.Println(SumToN(10))  // 55
	fmt.Println(SumToN(1))   // 1
}
