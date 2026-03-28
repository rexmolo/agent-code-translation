package main

import (
	"fmt"
)

func main() {
	fmt.Println(minSubArraySum([]int{2, 3, 4, 1, 2, 4})) // Output: 1
	fmt.Println(minSubArraySum([]int{-1, -2, -3}))       // Output: -6
}

// MinSubArraySum finds the minimum sum of any non-empty sub-array of nums.
// It uses Kadane's algorithm on the negated array to find the maximum
// subarray sum, then negates it to get the minimum.
func minSubArraySum(nums []int) int {
	maxSum := nums[0]
	s := 0

	for _, num := range nums {
		s += -num
		if s > maxSum {
			maxSum = s
		}
		if s < 0 {
			s = 0
		}
	}

	return -maxSum
}

// Using the requested function signature
func Minsubarraysum(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxSum := nums[0]
	s := 0

	for _, num := range nums {
		s += -num
		if s > maxSum {
			maxSum = s
		}
		if s < 0 {
			s = 0
		}
	}

	return -maxSum
}