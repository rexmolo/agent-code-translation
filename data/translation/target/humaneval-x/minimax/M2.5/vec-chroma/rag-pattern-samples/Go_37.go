package main

import (
	"fmt"
	"sort"
)

// SortEven takes a list l and returns a list l' such that
// l' is identical to l in the odd indices, while its values at the even indices are equal
// to the values of the even indices of l, but sorted.
func SortEven(l []int) []int {
	// Extract elements at even indices (0, 2, 4, ...)
	evens := make([]int, 0, (len(l)+1)/2)
	for i := 0; i < len(l); i += 2 {
		evens = append(evens, l[i])
	}

	// Extract elements at odd indices (1, 3, 5, ...)
	odds := make([]int, 0, len(l)/2)
	for i := 1; i < len(l); i += 2 {
		odds = append(odds, l[i])
	}

	// Sort the even-indexed elements
	sort.Ints(evens)

	// Interleave sorted evens with odds
	ans := make([]int, 0, len(l))
	for i := 0; i < len(odds); i++ {
		ans = append(ans, evens[i], odds[i])
	}

	// If there's one more even element than odd, append the last even
	if len(evens) > len(odds) {
		ans = append(ans, evens[len(evens)-1])
	}

	return ans
}

func main() {
	// Test cases
	fmt.Println(SortEven([]int{1, 2, 3}))       // [1 2 3]
	fmt.Println(SortEven([]int{5, 6, 3, 4}))   // [3 6 5 4]
}
