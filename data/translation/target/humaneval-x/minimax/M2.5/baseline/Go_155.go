package main

import (
	"fmt"
	"math"
)

func EvenOddCount(num int) [2]int {
	evenCount := 0
	oddCount := 0

	// Get absolute value to handle negative numbers
	n := int(math.Abs(float64(num)))

	// Convert number to string to iterate over digits
	s := fmt.Sprintf("%d", n)

	// Count even and odd digits
	for _, c := range s {
		digit := int(c - '0')
		if digit%2 == 0 {
			evenCount++
		} else {
			oddCount++
		}
	}

	return [2]int{evenCount, oddCount}
}