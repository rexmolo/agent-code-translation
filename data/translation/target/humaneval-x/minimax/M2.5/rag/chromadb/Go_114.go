package main

import (
	"fmt"
	"slices"
)

func minSubArraySum(nums []int) int {
	maxSum := 0
	s := 0
	for _, num := range nums {
		s += -num
		if s < 0 {
			s = 0
		}
		if s > maxSum {
			maxSum = s
		}
	}
	if maxSum == 0 {
		// All numbers were non-positive, find the maximum (least negative)
		maxVal := slices.Max(nums)
		maxSum = -maxVal
	}
	return -maxSum
}

func main() {
	// Test cases
	fmt.Println(minSubArraySum([]int{2, 3, 4, 1, 2, 4})) // Expected: 1
	fmt.Println(minSubArraySum([]int{-1, -2, -3}))      // Expected: -6
}