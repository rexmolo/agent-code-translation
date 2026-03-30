package main

import (
	"fmt"
)

func RollingMax(numbers []int) []int {
	if len(numbers) == 0 {
		return []int{}
	}

	runningMax := numbers[0]
	result := []int{runningMax}

	for i := 1; i < len(numbers); i++ {
		runningMax = max(runningMax, numbers[i])
		result = append(result, runningMax)
	}

	return result
}

func main() {
	// Test the function
	testCases := [][]int{
		{1, 2, 3, 2, 3, 4, 2},
		{1, 2, 3, 4, 5},
		{5, 4, 3, 2, 1},
		{},
	}

	for _, tc := range testCases {
		result := RollingMax(tc)
		fmt.Printf("Input: %v\n", tc)
		fmt.Printf("Output: %v\n\n", result)
	}
}
