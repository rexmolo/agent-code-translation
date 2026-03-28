package main

import (
	"fmt"
	"sort"
)

func Monotonic(l []int) bool {
	if len(l) <= 1 {
		return true
	}

	// Create sorted copy for ascending check
	asc := make([]int, len(l))
	copy(asc, l)
	sort.Ints(asc)

	// Check if ascending
	isAscending := true
	for i := range l {
		if l[i] != asc[i] {
			isAscending = false
			break
		}
	}
	if isAscending {
		return true
	}

	// Create sorted copy for descending check
	desc := make([]int, len(l))
	copy(desc, l)
	sort.Slice(desc, func(i, j int) bool {
		return desc[i] > desc[j]
	})

	// Check if descending
	isDescending := true
	for i := range l {
		if l[i] != desc[i] {
			isDescending = false
			break
		}
	}

	return isDescending
}

func main() {
	fmt.Println(Monotonic([]int{1, 2, 4, 20}))   // Expected: true
	fmt.Println(Monotonic([]int{1, 20, 4, 10}))  // Expected: false
	fmt.Println(Monotonic([]int{4, 1, 0, -10}))  // Expected: true
}
