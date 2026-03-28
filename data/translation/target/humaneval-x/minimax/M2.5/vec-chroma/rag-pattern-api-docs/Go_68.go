package main

import "fmt"

// Pluck returns the smallest even value and its index from the array.
// Returns an empty slice if there are no even values or the array is empty.
func Pluck(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}

	// Find the minimum even value
	minEven := -1
	for _, v := range arr {
		if v%2 == 0 {
			if minEven == -1 || v < minEven {
				minEven = v
			}
		}
	}

	// If no even values found, return empty slice
	if minEven == -1 {
		return []int{}
	}

	// Find the first occurrence index of minEven in original array
	index := -1
	for i, v := range arr {
		if v == minEven {
			index = i
			break
		}
	}

	return []int{minEven, index}
}

func main() {
	// Test cases
	fmt.Println(Pluck([]int{4, 2, 3}))   // Output: [2 1]
	fmt.Println(Pluck([]int{1, 2, 3}))   // Output: [2 1]
	fmt.Println(Pluck([]int{}))          // Output: []
	fmt.Println(Pluck([]int{5, 0, 3, 0, 4, 2})) // Output: [0 1]
}
