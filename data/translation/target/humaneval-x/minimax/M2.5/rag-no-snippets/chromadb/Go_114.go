package main

import (
	"math"
	"slices"
)

func minSubArraySum(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

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
		// All numbers are positive, find max(-i for i in nums)
		// which is -min(nums)
		minVal := slices.Min(nums)
		return -minVal
	}

	return -maxSum
}
