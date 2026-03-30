package main

import (
	"fmt"
	"sort"
)

func Monotonic(l []int) bool {
	if len(l) <= 1 {
		return true
	}

	// Check if sorted in ascending order
	sortedAsc := make([]int, len(l))
	copy(sortedAsc, l)
	sort.Ints(sortedAsc)
	ascending := true
	for i, v := range l {
		if v != sortedAsc[i] {
			ascending = false
			break
		}
	}
	if ascending {
		return true
	}

	// Check if sorted in descending order
	sortedDesc := make([]int, len(l))
	copy(sortedDesc, l)
	sort.Sort(sort.Reverse(sort.IntSlice(sortedDesc)))
	descending := true
	for i, v := range l {
		if v != sortedDesc[i] {
			descending = false
			break
		}
	}

	return descending
}

func main() {
	fmt.Println(Monotonic([]int{1, 2, 4, 20}))   // true
	fmt.Println(Monotonic([]int{1, 20, 4, 10})) // false
	fmt.Println(Monotonic([]int{4, 1, 0, -10})) // true
}