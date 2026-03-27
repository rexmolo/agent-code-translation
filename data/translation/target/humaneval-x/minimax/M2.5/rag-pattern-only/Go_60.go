package main

import "fmt"

func SumToN(n int) int {
	return n * (n + 1) / 2
}

func main() {
	// Test cases based on docstring examples
	fmt.Println(SumToN(30))  // 465
	fmt.Println(SumToN(100)) // 5050
	fmt.Println(SumToN(5))   // 15
	fmt.Println(SumToN(10))  // 55
	fmt.Println(SumToN(1))   // 1
}
