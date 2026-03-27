package main

import "fmt"

func IsEqualToSumEven(n int) bool {
	// A positive even number is at least 2
	// 4 positive even numbers sum to at least 8 (2+2+2+2)
	// Any even number >= 8 can be written as the sum of exactly 4 positive even numbers
	return n%2 == 0 && n >= 8
}

func main() {
	// Test cases from docstring
	fmt.Println(IsEqualToSumEven(4)) // False
	fmt.Println(IsEqualToSumEven(6)) // False
	fmt.Println(IsEqualToSumEven(8)) // True
}