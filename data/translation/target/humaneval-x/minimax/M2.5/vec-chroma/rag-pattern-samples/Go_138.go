package main

import "fmt"

func IsEqualToSumEven(n int) bool {
	return n%2 == 0 && n >= 8
}

func main() {
	// Test cases based on docstring
	fmt.Println(IsEqualToSumEven(4)) // Expected: false
	fmt.Println(IsEqualToSumEven(6)) // Expected: false
	fmt.Println(IsEqualToSumEven(8)) // Expected: true
}
