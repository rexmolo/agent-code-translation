package main

import (
	"fmt"
	"strconv"
)

func Digits(n int) int {
	product := 1
	oddCount := 0

	// Convert n to string to iterate over each digit
	s := strconv.Itoa(n)
	for _, c := range s {
		digit := int(c - '0') // Convert rune to integer digit
		if digit%2 == 1 {
			product *= digit
			oddCount++
		}
	}

	if oddCount == 0 {
		return 0
	}
	return product
}

func main() {
	// Test cases from the Python docstring
	fmt.Println(Digits(1))  // Expected: 1
	fmt.Println(Digits(4))  // Expected: 0
	fmt.Println(Digits(235)) // Expected: 15
}