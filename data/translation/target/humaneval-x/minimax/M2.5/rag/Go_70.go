package main

import (
	"fmt"
	"slices"
)

func StrangeSortList(lst []int) []int {
	res := make([]int, 0, len(lst))
	switchFlag := true

	for len(lst) > 0 {
		var val int
		if switchFlag {
			val = slices.Min(lst)
		} else {
			val = slices.Max(lst)
		}
		res = append(res, val)

		// Find index and remove the element from lst
		idx := -1
		for i, v := range lst {
			if v == val {
				idx = i
				break
			}
		}
		if idx != -1 {
			lst = append(lst[:idx], lst[idx+1:]...)
		}

		switchFlag = !switchFlag
	}

	return res
}

func main() {
	// Test cases
	testCases := [][]int{
		{1, 2, 3, 4},
		{5, 5, 5, 5},
		{},
	}

	for _, tc := range testCases {
		result := StrangeSortList(tc)
		fmt.Println(result)
	}
}