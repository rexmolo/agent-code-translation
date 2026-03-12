package main

import (
	"strconv"
)

func CountNums(arr []int) int {
	// Helper function to compute digit sum with signed first digit for negative numbers
	digitSum := func(n int) int {
		neg := 1
		if n < 0 {
			n = -n
			neg = -1
		}

		// Convert number to string and extract digits
		sum := 0
		strN := strconv.Itoa(n)
		for i, c := range strN {
			digit := int(c - '0')
			if i == 0 {
				digit *= neg
			}
			sum += digit
		}
		return sum
	}

	// Count elements with digit sum > 0
	count := 0
	for _, n := range arr {
		if digitSum(n) > 0 {
			count++
		}
	}
	return count
}