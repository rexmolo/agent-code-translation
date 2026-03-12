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
	ascending := sort.IntsAreSorted(l)

	// Check if sorted in descending order using sort.Reverse
	descending := sort.IsSorted(sort.Reverse(sort.IntSlice(l)))

	return ascending || descending
}

func main() {
	fmt.Println(Monotonic([]int{1, 2, 4, 20}))  // True
	fmt.Println(Monotonic([]int{1, 20, 4, 10})) // False
	fmt.Println(Monotonic([]int{4, 1, 0, -10})) // True
}
