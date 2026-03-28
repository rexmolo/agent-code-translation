package main

import "slices"

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
		maxSum = -slices.Max(func(nums []int) []int {
			result := make([]int, len(nums))
			for i, n := range nums {
				result[i] = -n
			}
			return result
		}(nums))
	}
	minSum := -maxSum
	return minSum
}

func Minsubarraysum(nums []int) int {
	return minSubArraySum(nums)
}

func main() {
	// Test cases can be added here
}
