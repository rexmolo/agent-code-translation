package main

import "slices"

func SortEven(l []int) []int {
	// Extract even-indexed elements (indices 0, 2, 4, ...)
	evens := make([]int, 0)
	for i := 0; i < len(l); i += 2 {
		evens = append(evens, l[i])
	}

	// Extract odd-indexed elements (indices 1, 3, 5, ...)
	odds := make([]int, 0)
	for i := 1; i < len(l); i += 2 {
		odds = append(odds, l[i])
	}

	// Sort the even-indexed elements
	slices.Sort(evens)

	// Interleave evens and odds back together
	ans := make([]int, 0, len(l))
	for i := 0; i < len(odds); i++ {
		ans = append(ans, evens[i], odds[i])
	}

	// If there's one more even element than odd, append it at the end
	if len(evens) > len(odds) {
		ans = append(ans, evens[len(evens)-1])
	}

	return ans
}