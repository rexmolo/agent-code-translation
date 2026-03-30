package main

import (
	"fmt"
	"sort"
)

func Monotonic(l []int) bool {
	if len(l) <= 1 {
		return true
	}

	// Create a copy for ascending sort
	asc := make([]int, len(l))
	copy(asc, l)
	sort.Slice(asc, func(i, j int) bool {
		return asc[i] < asc[j]
	})

	// Create a copy for descending sort
	desc := make([]int, len(l))
	copy(desc, l)
	sort.Slice(desc, func(i, j int) bool {
		return desc[i] > desc[j]
	})

	// Check if original equals ascending or descending sorted
	ascEqual := true
	descEqual := true
	for i := range l {
		if l[i] != asc[i] {
			ascEqual = false
		}
		if l[i] != desc[i] {
			descEqual = false
		}
		if !ascEqual && !descEqual {
			return false
		}
	}

	return ascEqual || descEqual
}

func main() {
	// Test cases
	fmt.Println(Monotonic([]int{1, 2, 4, 20})) // true
	fmt.Println(Monotonic([]int{1, 20, 4, 10})) // false
	fmt.Println(Monotonic([]int{4, 1, 0, -10})) // true
}