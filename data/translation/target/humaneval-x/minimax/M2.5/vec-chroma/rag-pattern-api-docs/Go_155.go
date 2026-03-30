package main

import (
	"math"
	"strconv"
)

func EvenOddCount(num int) [2]int {
	evenCount := 0
	oddCount := 0

	// Get absolute value and convert to string
	absNum := int(math.Abs(float64(num)))
	numStr := strconv.Itoa(absNum)

	// Iterate through each digit character
	for _, c := range numStr {
		digit := int(c - '0') // Convert rune to integer digit
		if digit%2 == 0 {
			evenCount++
		} else {
			oddCount++
		}
	}

	return [2]int{evenCount, oddCount}
}
