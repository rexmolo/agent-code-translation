package main

import (
	"fmt"
	"math"
)

func main() {
	// Example usage
	fmt.Println(minSubArraySum([]int{2, 3, 4, 1, 2, 4})) // Output: 1
	fmt.Println(minSubArraySum([]int{-1, -2, -3}))       // Output: -6
}

func minSubArraySum(nums []int) int {
	// Handle empty slice
	if len(nums) == 0 {
		return 0
	}

	// Kadane's algorithm variant for finding minimum subarray sum
	// We track the minimum sum ending at each position
	minSum := nums[0]
	currentSum := nums[0]

	for i := 1; i < len(nums); i++ {
		// Either start a new subarray at current position,
		// or extend the previous subarray
		if currentSum+nums[i] < nums[i] {
			currentSum = nums[i]
		} else {
			currentSum = currentSum + nums[i]
		}

		// Update minimum if current sum is smaller
		if currentSum < minSum {
			minSum = currentSum
		}
	}

	// If array contains all positive numbers, return the smallest element
	allPositive := true
	for _, n := range nums {
		if n <= 0 {
			allPositive = false
			break
		}
	}
	if allPositive {
		return int(math.Min(float64(nums[0]), math.Inf(1)))
	}

	return minSum
}