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

	for _, c := range strconv.Itoa(absNum) {
		digit := int(c - '0')
		if digit%2 == 0 {
			evenCount++
		} else {
			oddCount++
		}
	}

	return [2]int{evenCount, oddCount}
}
