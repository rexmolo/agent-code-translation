package main

import (
	"fmt"
	"sort"
)

// Median returns the median of elements in the slice l.
func Median(l []int) float64 {
	// Create a copy to avoid modifying the original slice.
	// Python's sorted() also returns a new list.
	data := make([]int, len(l))
	copy(data, l)

	// Sort the copied slice
	sort.Ints(data)

	n := len(data)
	if n == 0 {
		// The median of an empty set is undefined. Returning 0.0 is a common practical choice.
		// The original Python code would raise an IndexError.
		return 0.0
	}

	if n%2 == 1 {
		// Odd number of elements: return the middle one.
		return float64(data[n/2])
	} else {
		// Even number of elements: return the average of the two middle ones.
		mid1 := data[n/2-1]
		mid2 := data[n/2]
		return (float64(mid1) + float64(mid2)) / 2.0
	}
}

// main function to demonstrate the Median function, similar to the Python doctests.
func main() {
	list1 := []int{3, 1, 2, 4, 5}
	fmt.Println(Median(list1))

	list2 := []int{-10, 4, 6, 1000, 10, 20}
	fmt.Println(Median(list2))
}
