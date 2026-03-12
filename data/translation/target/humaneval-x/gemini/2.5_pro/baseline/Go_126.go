package main

import (
	"fmt"
)

// IsSorted checks if a list of numbers is sorted in ascending order.
// It returns false if any number appears more than twice.
// Assumes no negative numbers and only integers.
func IsSorted(lst []int) bool {
	// An empty or single-element list is considered sorted and has no invalid duplicates.
	if len(lst) <= 1 {
		return true
	}

	counts := make(map[int]int)
	// Account for the first element before starting the loop
	counts[lst[0]] = 1

	for i := 1; i < len(lst); i++ {
		// 1. Check if the slice is sorted in non-decreasing order
		if lst[i] < lst[i-1] {
			return false
		}

		// 2. Count element occurrences and check for more than 2 duplicates
		num := lst[i]
		counts[num]++
		if counts[num] > 2 {
			return false
		}
	}

	// If the loop completes, the list is sorted and has valid duplicate counts.
	return true
}

// main function to demonstrate the usage of IsSorted
func main() {
	fmt.Println("is_sorted([5]) ->", IsSorted([]int{5}))
	fmt.Println("is_sorted([1, 2, 3, 4, 5]) ->", IsSorted([]int{1, 2, 3, 4, 5}))
	fmt.Println("is_sorted([1, 3, 2, 4, 5]) ->", IsSorted([]int{1, 3, 2, 4, 5}))
	fmt.Println("is_sorted([1, 2, 3, 4, 5, 6]) ->", IsSorted([]int{1, 2, 3, 4, 5, 6}))
	fmt.Println("is_sorted([1, 2, 3, 4, 5, 6, 7]) ->", IsSorted([]int{1, 2, 3, 4, 5, 6, 7}))
	fmt.Println("is_sorted([1, 3, 2, 4, 5, 6, 7]) ->", IsSorted([]int{1, 3, 2, 4, 5, 6, 7}))
	fmt.Println("is_sorted([1, 2, 2, 3, 3, 4]) ->", IsSorted([]int{1, 2, 2, 3, 3, 4}))
	fmt.Println("is_sorted([1, 2, 2, 2, 3, 4]) ->", IsSorted([]int{1, 2, 2, 2, 3, 4}))
}