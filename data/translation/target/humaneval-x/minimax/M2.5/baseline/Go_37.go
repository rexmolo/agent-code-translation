package main

import (
	"fmt"
	"sort"
)

func SortEven(l []int) []int {
	// Extract even-indexed elements (indices 0, 2, 4, ...)
	evens := []int{}
	for i := 0; i < len(l); i += 2 {
		evens = append(evens, l[i])
	}

	// Extract odd-indexed elements (indices 1, 3, 5, ...)
	odds := []int{}
	for i := 1; i < len(l); i += 2 {
		odds = append(odds, l[i])
	}

	// Sort the evens
	sort.Ints(evens)

	// Interleave evens and odds
	ans := []int{}
	minLen := len(evens)
	if len(odds) < len(evens) {
		minLen = len(odds)
	}

	for i := 0; i < minLen; i++ {
		ans = append(ans, evens[i], odds[i])
	}

	// If there's one more even than odd, append the last even
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