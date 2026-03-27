package main

import (
	"sort"
	"strconv"
)

func OrderByPoints(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}

	// Create a copy to avoid modifying the original slice
	result := make([]int, len(nums))
	copy(result, nums)

	// Use SliceStable to preserve original order for equal digit sums
	// (similar to Python's stable sorted)
	sort.SliceStable(result, func(i, j int) bool {
		return digitSum(result[i]) < digitSum(result[j])
	})

	return result
}

// digitSum calculates the sum of digits with special handling for negative numbers
// For negative numbers: convert to positive, calculate digit sum, but multiply first digit by -1
func digitSum(n int) int {
	neg := 1
	if n < 0 {
		n = -n
		neg = -1
	}

	s := strconv.Itoa(n)
	digits := []rune(s)

	// First digit gets the negative sign applied
	firstDigit := int(digits[0]-'0') * neg
	sum := firstDigit

	// Rest of the digits are added normally
	for i := 1; i < len(digits); i++ {
		sum += int(digits[i] - '0')
	}

	return sum
}