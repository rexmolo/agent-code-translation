package main

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
		maxVal := nums[0]
		for i := 1; i < len(nums); i++ {
			if nums[i] > maxVal {
				maxVal = nums[i]
			}
		}
		return -maxVal
	}

	return -maxSum
}
