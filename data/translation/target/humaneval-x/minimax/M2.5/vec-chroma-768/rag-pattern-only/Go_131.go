package main

import (
	"fmt"
	"strconv"
)

func Digits(n int) int {
	// Convert to string to iterate over each digit
	s := strconv.Itoa(n)
	product := 1
	oddCount := 0

	for _, char := range s {
		digit := int(char - '0')
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
	// Test cases from the docstring
	fmt.Println(Digits(1))   // 1
	fmt.Println(Digits(4))   // 0
	fmt.Println(Digits(235)) // 15
}