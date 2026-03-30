package main

import "fmt"

func SumToN(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

func main() {
	// Test cases matching the Python docstring examples
	fmt.Println(SumToN(30))  // 465
	fmt.Println(SumToN(100)) // 5050
	fmt.Println(SumToN(5))   // 15
	fmt.Println(SumToN(10))  // 55
	fmt.Println(SumToN(1))   // 1
}
