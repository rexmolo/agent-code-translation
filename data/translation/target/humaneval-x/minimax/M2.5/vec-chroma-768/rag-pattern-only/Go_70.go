package main

import "fmt"

func StrangeSortList(lst []int) []int {
	res := make([]int, 0)
	isMin := true

	// Make a working copy of the input slice so we can modify it
	remaining := make([]int, len(lst))
	copy(remaining, lst)

	for len(remaining) > 0 {
		var value int
		if isMin {
			value = min(remaining)
		} else {
			value = max(remaining)
		}
		res = append(res, value)
		remaining = removeElement(remaining, value)
		isMin = !isMin
	}

	return res
}

func min(lst []int) int {
	m := lst[0]
	for _, v := range lst[1:] {
		if v < m {
			m = v
		}
	}
	return m
}

func max(lst []int) int {
	m := lst[0]
	for _, v := range lst[1:] {
		if v > m {
			m = v
		}
	}
	return m
}

func removeElement(lst []int, val int) []int {
	for i, v := range lst {
		if v == val {
			return append(lst[:i], lst[i+1:]...)
		}
	}
	return lst
}

func main() {
	// Test cases
	testCases := [][]int{
		{1, 2, 3, 4},
		{5, 5, 5, 5},
		{},
		{1},
		{3, 1, 2},
	}

	for _, tc := range testCases {
		result := StrangeSortList(tc)
		fmt.Printf("Input: %v, Output: %v\n", tc, result)
	}
}