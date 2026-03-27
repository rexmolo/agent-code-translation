package main

import (
	"fmt"
	"strconv"
)

func Solve(N int) string {
	// Convert N to string to iterate over each digit
	str := strconv.Itoa(N)

	// Sum all digits
	sum := 0
	for _, c := range str {
		digit := int(c - '0')
		sum += digit
	}

	// Convert sum to binary string using %b format
	return fmt.Sprintf("%b", sum)
}
