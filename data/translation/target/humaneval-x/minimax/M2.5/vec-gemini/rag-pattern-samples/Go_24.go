package main

import "fmt"

func LargestDivisor(n int) int {
	for i := n - 1; i >= 1; i-- {
		if n%i == 0 {
			return i
		}
	}
	return 1
}

func main() {
	// Test cases
	fmt.Println(LargestDivisor(15)) // 5
	fmt.Println(LargestDivisor(12)) // 6
	fmt.Println(LargestDivisor(7))  // 1
	fmt.Println(LargestDivisor(4))  // 2
}
