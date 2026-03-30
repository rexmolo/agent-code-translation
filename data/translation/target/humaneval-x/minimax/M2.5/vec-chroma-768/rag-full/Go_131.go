package main

import (
	"fmt"
	"strconv"
)

func Digits(n int) int {
	// Convert integer to string to iterate over digits
	s := strconv.Itoa(n)

	product := 1
	oddCount := 0

	// Iterate over each character in the string
	for _, c := range s {
		// Convert character to digit by subtracting ASCII value of '0'
		digit := int(c - '0')
		if digit%2 == 1 {
			product *= digit
			oddCount++
		}
	}

	// Return 0 if no odd digits were found
	if oddCount == 0 {
		return 0
	}
	return product
}

func main() {
	// Test cases to verify the translation
	fmt.Println(Digits(1))   // Expected: 1
	fmt.Println(Digits(4))   // Expected: 0
	fmt.Println(Digits(235)) // Expected: 15
}
