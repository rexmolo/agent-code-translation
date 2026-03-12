package main

import "slices"

func SortEven(l []int) []int {
	// Extract elements at even indices (0, 2, 4, ...)
	evens := []int{}
	for i := 0; i < len(l); i += 2 {
		evens = append(evens, l[i])
	}

	// Extract elements at odd indices (1, 3, 5, ...)
	odds := []int{}
	for i := 1; i < len(l); i += 2 {
		odds = append(odds, l[i])
	}

	// Sort the evens slice
	slices.Sort(evens)

	// Interleave evens and odds
	ans := []int{}
	for i := 0; i < len(odds); i++ {
		ans = append(ans, evens[i], odds[i])
	}

	// If there's one more even than odd, append the last even
	if len(evens) > len(odds) {
		ans = append(ans, evens[len(evens)-1])
	}

	return ans
}
