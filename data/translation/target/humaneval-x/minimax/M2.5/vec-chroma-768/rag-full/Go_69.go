package main

import "slices"

func Search(lst []int) int {
	// Handle empty slice case
	if len(lst) == 0 {
		return -1
	}

	// Find max value in the list
	maxVal := slices.Max(lst)

	// Create frequency array where index represents the integer value
	frq := make([]int, maxVal+1)
	for _, v := range lst {
		frq[v]++
	}

	// Find the largest integer i where frequency >= i
	ans := -1
	for i := 1; i <= maxVal; i++ {
		if frq[i] >= i {
			ans = i
		}
	}

	return ans
}
