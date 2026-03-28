package main

import (
	"strconv"
)

func CountNums(arr []int) int {
	// Helper function to compute digit sum with signed first digit
	digitsSum := func(n int) int {
		neg := 1
		if n < 0 {
			n = -n
			neg = -1
		}

		// Convert to string to iterate over digits
		str := strconv.Itoa(n)

		// Process digits: first digit gets the sign, rest are positive
		sum := 0
		for i, c := range str {
			digit := int(c - '0')
			if i == 0 {
				digit = digit * neg
			}
			sum += digit
		}
		return sum
	}

	// Count elements where digit sum > 0
	count := 0
	for _, n := range arr {
		if digitsSum(n) > 0 {
			count++
		}
	}

	return count
}

func main() {
	// Test cases
	_ = CountNums([]int{}) == 0
	_ = CountNums([]int{-1, 11, -11}) == 1
	_ = CountNums([]int{1, 1, 2}) == 3
}
