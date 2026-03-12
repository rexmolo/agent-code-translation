package main

func Minsubarraysum(nums []int) int {
	// The problem implies a non-empty array, but we handle the empty case for robustness.
	if len(nums) == 0 {
		return 0
	}

	// The Python code finds the minimum subarray sum by applying Kadane's algorithm
	// to the negated array to find its maximum subarray sum, and then negating the result.
	// We replicate that exact logic here.
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

	// The Python version of Kadane's algorithm used initializes max_sum to 0.
	// This leads to an incorrect result of 0 if all numbers in the negated array are negative
	// (i.e., all original numbers are positive). The Python code has a special check for this.
	if maxSum == 0 {
		// This block corresponds to `max_sum = max(-i for i in nums)`.
		// It finds the largest element in the negated array, which is the correct
		// answer when all elements are negative.
		maxOfNegated := -nums[0]
		for i := 1; i < len(nums); i++ {
			if -nums[i] > maxOfNegated {
				maxOfNegated = -nums[i]
			}
		}
		maxSum = maxOfNegated
	}

	// min_subarray_sum(nums) == -max_subarray_sum(-nums)
	minSum := -maxSum
	return minSum
}
