package main

import (
	"sort"
	"strconv"
)

// digitsSum calculates the sum of a number's digits based on the specific rules.
// For negative numbers, the most significant digit is treated as negative.
func digitsSum(n int) int {
	if n == 0 {
		return 0
	}

	isNegative := false
	workNum := n
	if workNum < 0 {
		isNegative = true
		workNum = -workNum
	}

	s := strconv.Itoa(workNum)
	sum := 0
	for _, r := range s {
		sum += int(r - '0')
	}

	if isNegative {
		// The original logic negates the first digit. For example, for -123,
		// the sum is -1 + 2 + 3 = 4. Our current sum is 1 + 2 + 3 = 6.
		// To get to 4, we subtract the first digit twice from the total sum.
		firstDigit := int(s[0] - '0')
		sum -= 2 * firstDigit
	}

	return sum
}

// OrderByPoints sorts the given slice of integers in ascending order
// according to the sum of their digits. If sums are equal, the original
// relative order is preserved.
func OrderByPoints(nums []int) []int {
	// Create a copy to avoid modifying the original slice, mimicking Python's sorted().
	result := make([]int, len(nums))
	copy(result, nums)

	// Use SliceStable for a stable sort, which preserves the original order
	// of elements with equal sort keys.
	sort.SliceStable(result, func(i, j int) bool {
		return digitsSum(result[i]) < digitsSum(result[j])
	})

	return result
}
