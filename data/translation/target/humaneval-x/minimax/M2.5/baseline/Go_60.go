package main

import "fmt"

// SumToN sums numbers from 0 to n (inclusive).
// This is equivalent to the Python sum(range(n + 1))
func SumToN(n int) int {
	sum := 0
	for i := 0; i <= n; i++ {
		sum += i
	}
	return sum
}

func main() {
	// Test cases from Python docstring
	fmt.Println(SumToN(30)) // 465
	fmt.Println(SumToN(100)) // 5050
	fmt.Println(SumToN(5))  // 15
	fmt.Println(SumToN(10)) // 55
	fmt.Println(SumToN(1))  // 1
}
