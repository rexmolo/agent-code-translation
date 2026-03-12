package main

import "sort"

// StrangeSortList sorts a list of integers in a peculiar order.
// The sorting starts with the minimum value, then the maximum of the remaining integers,
// then the minimum of the remaining, and so on.
func StrangeSortList(lst []int) []int {
	if len(lst) == 0 {
		return []int{}
	}

	// Create a copy to sort. This is more efficient (O(n log n)) than the original
	// Python version's repeated min/max search and removal (O(n^2)).
	temp := make([]int, len(lst))
	copy(temp, lst)
	sort.Ints(temp)

	res := make([]int, 0, len(lst))
	left, right := 0, len(temp)-1

	// Build the result by alternating between the smallest (from the left)
	// and largest (from the right) elements of the sorted slice.
	for left <= right {
		// Append the minimum of the remaining elements
		res = append(res, temp[left])
		left++

		// If there are still elements, append the maximum of the remaining
		if left <= right {
			res = append(res, temp[right])
			right--
		}
	}

	return res
}
