package main

import "slices"

func SortEven(l []int) []int {
	// Extract even indices (0, 2, 4, ...)
	evens := make([]int, 0, (len(l)+1)/2)
	for i := 0; i < len(l); i += 2 {
		evens = append(evens, l[i])
	}

	// Extract odd indices (1, 3, 5, ...)
	odds := make([]int, 0, len(l)/2)
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

	// If there's an extra even element (more evens than odds), append it
	if len(evens) > len(odds) {
		ans = append(ans, evens[len(evens)-1])
	}

	return ans
}
