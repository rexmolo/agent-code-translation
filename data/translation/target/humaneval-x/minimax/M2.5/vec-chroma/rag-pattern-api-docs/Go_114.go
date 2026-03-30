package main

import "slices"

// minSubArraySum finds the minimum sum of any non-empty sub-array of nums.
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
		// max(-i for i in nums)
		negNums := make([]int, len(nums))
		for i, num := range nums {
			negNums[i] = -num
		}
		maxSum = slices.Max(negNums)
	}
	minSum := -maxSum
	return minSum
}

func main() {
	// Test cases
	println(minSubArraySum([]int{2, 3, 4, 1, 2, 4})) // == 1
	println(minSubArraySum([]int{-1, -2, -3}))       // == -6
}