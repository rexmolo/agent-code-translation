package main

import (
	"fmt"
)

func Search(lst []int) int {
	// Find the maximum value in the list
	maxVal := lst[0]
	for _, v := range lst {
		if v > maxVal {
			maxVal = v
		}
	}

	// Create frequency array (index represents the integer value)
	frq := make([]int, maxVal+1)

	// Count frequencies for each integer
	for _, i := range lst {
		frq[i]++
	}

	// Find the greatest integer i where frequency >= i
	ans := -1
	for i := 1; i < len(frq); i++ {
		if frq[i] >= i {
			ans = i
		}
	}

	return ans
}

func main() {
	// Test cases
	fmt.Println(Search([]int{4, 1, 2, 2, 3, 1}))                   // Expected: 2
	fmt.Println(Search([]int{1, 2, 2, 3, 3, 3, 4, 4, 4}))          // Expected: 3
	fmt.Println(Search([]int{5, 5, 4, 4, 4}))                      // Expected: -1
}
