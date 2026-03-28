package main

import (
	"strconv"
)

func CountNums(arr []int) int {
	// Helper function to compute signed digit sum
	digitsSum := func(n int) int {
		neg := 1
		if n < 0 {
			n = -n
			neg = -1
		}
		// Convert to string and extract digits
		str := strconv.Itoa(n)
		sum := 0
		for i, ch := range str {
			digit := int(ch - '0')
			if i == 0 {
				digit *= neg
			}
			sum += digit
		}
		return sum
	}

	// Count elements where digits sum > 0
	count := 0
	for _, n := range arr {
		if digitsSum(n) > 0 {
			count++
		}
	}
	return count
}
