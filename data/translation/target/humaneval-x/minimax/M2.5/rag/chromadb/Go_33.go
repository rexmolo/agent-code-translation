package main

import "sort"

func SortThird(l []int) []int {
	// Make a copy of the slice
	result := make([]int, len(l))
	copy(result, l)

	// Extract elements at indices divisible by 3 (0, 3, 6, ...)
	third := make([]int, 0)
	for i := 0; i < len(l); i += 3 {
		third = append(third, l[i])
	}

	// Sort the third elements
	sort.Ints(third)

	// Put them back into result at positions 0, 3, 6, ...
	for i := 0; i < len(third); i++ {
		result[i*3] = third[i]
	}

	return result
}

// Test with examples
func main() {
	// Example 1
	result1 := SortThird([]int{1, 2, 3})
	println(result1)

	// Example 2
	result2 := SortThird([]int{5, 6, 3, 4, 8, 9, 2})
	println(result2)
}
