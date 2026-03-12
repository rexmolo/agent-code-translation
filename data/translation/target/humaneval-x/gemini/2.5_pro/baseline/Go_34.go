package main

import (
	"fmt"
	"sort"
)

// Unique returns a new sorted slice containing the unique elements of the input slice.
func Unique(l []int) []int {
	// Use a map to store the unique elements. The keys of the map will be the unique integers.
	// The value is an empty struct {} because it occupies zero memory, making it more efficient
	// than using a bool.
	uniqueElements := make(map[int]struct{})
	for _, item := range l {
		uniqueElements[item] = struct{}{}
	}

	// Create a new slice to hold the unique elements (the keys from the map).
	// Pre-allocating the slice capacity to the number of unique elements improves performance.
	keys := make([]int, 0, len(uniqueElements))
	for k := range uniqueElements {
		keys = append(keys, k)
	}

	// Sort the slice of unique elements in ascending order.
	sort.Ints(keys)

	return keys
}

// main function to demonstrate the Unique function, mimicking the Python docstring example.
func main() {
	input := []int{5, 3, 5, 2, 3, 3, 9, 0, 123}
	result := Unique(input)
	fmt.Println(result)
	// Expected output: [0 2 3 5 9 123]
}
