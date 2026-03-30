package main

import "slices"

func Minsubarraysum(nums []int) int {
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
		maxSum = -slices.Min(nums)
	}
	minSum := -maxSum
	return minSum
}
