package main

import (
	"fmt"
	"slices"
)

func StrangeSortList(lst []int) []int {
	result := make([]int, 0, len(lst))
	isMin := true

	for len(lst) > 0 {
		var val int
		if isMin {
			val = slices.Min(lst)
		} else {
			val = slices.Max(lst)
		}
		result = append(result, val)

		// Find and remove the element from the slice
		for i, v := range lst {
			if v == val {
				lst = slices.Delete(lst, i, i+1)
				break
			}
		}

		isMin = !isMin
	}

	return result
}

func main() {
	// Test cases
	testCases := [][]int{
		{1, 2, 3, 4},
		{5, 5, 5, 5},
		{},
		{1},
		{1, 2},
	}

	for _, tc := range testCases {
		result := StrangeSortList(tc)
		fmt.Println(result)
	}
}
