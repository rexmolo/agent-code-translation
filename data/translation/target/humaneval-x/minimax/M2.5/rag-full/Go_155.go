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
	strNum := strconv.Itoa(absNum)

	// Iterate over each character (digit)
	for _, char := range strNum {
		digit := int(char - '0') // Convert char to numeric value
		if digit%2 == 0 {
			evenCount++
		} else {
			oddCount++
		}
	}

	return [2]int{evenCount, oddCount}
}
