package main

import "fmt"

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
		maxSum = nums[0]
		for _, num := range nums {
			if num > maxSum {
				maxSum = num
			}
		}
	}
	return -maxSum
}

func main() {
	// Test cases
	fmt.Println(Minsubarraysum([]int{2, 3, 4, 1, 2, 4})) // Expected: 1
	fmt.Println(Minsubarraysum([]int{-1, -2, -3}))      // Expected: -6
}