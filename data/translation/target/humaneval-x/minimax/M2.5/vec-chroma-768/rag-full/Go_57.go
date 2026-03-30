package main

import (
	"fmt"
	"sort"
)

func Monotonic(l []int) bool {
	if len(l) <= 1 {
		return true
	}

	// Check if already sorted in ascending order
	if sort.IntsAreSorted(l) {
		return true
	}

	// Check if sorted in descending order
	descSorted := make([]int, len(l))
	copy(descSorted, l)
	sort.Sort(sort.Reverse(sort.IntSlice(descSorted)))

	return sort.IntsAreSorted(descSorted)
}

func main() {
	// Test cases from docstring
	fmt.Println(Monotonic([]int{1, 2, 4, 20}))  // true
	fmt.Println(Monotonic([]int{1, 20, 4, 10})) // false
	fmt.Println(Monotonic([]int{4, 1, 0, -10})) // true
}
