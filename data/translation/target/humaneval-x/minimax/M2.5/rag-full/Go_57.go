package main

import (
	"fmt"
	"sort"
)

func Monotonic(l []int) bool {
	// Create copies to avoid modifying the original slice
	asc := make([]int, len(l))
	desc := make([]int, len(l))
	copy(asc, l)
	copy(desc, l)

	// Sort in ascending order
	sort.Ints(asc)

	// Sort in descending order
	sort.Sort(sort.Reverse(sort.IntSlice(desc)))

	// Check if original equals either sorted version
	for i := range l {
		if l[i] != asc[i] && l[i] != desc[i] {
			return false
		}
	}
	return true
}

func main() {
	// Test cases from docstring
	fmt.Println(Monotonic([]int{1, 2, 4, 20}))   // true
	fmt.Println(Monotonic([]int{1, 20, 4, 10})) // false
	fmt.Println(Monotonic([]int{4, 1, 0, -10})) // true
}
