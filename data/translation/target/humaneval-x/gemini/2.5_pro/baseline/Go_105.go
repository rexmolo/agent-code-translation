package main

import (
	"sort"
)

// ByLength sorts integers from an array that are between 1 and 9 inclusive,
// reverses the result, and replaces each digit with its string name.
// It follows the logic of the original Python implementation:
// 1. Sort the entire input array in descending order.
// 2. Filter the sorted array, keeping only numbers between 1 and 9.
// 3. Convert the remaining numbers to their string representations.
func ByLength(arr []int) []string {
	// The Python `dic` is translated to a Go map.
	dic := map[int]string{
		1: "One",
		2: "Two",
		3: "Three",
		4: "Four",
		5: "Five",
		6: "Six",
		7: "Seven",
		8: "Eight",
		9: "Nine",
	}

	// Python's `sorted()` creates a new list. To replicate this and avoid
	// mutating the input slice, we create a copy.
	arrCopy := make([]int, len(arr))
	copy(arrCopy, arr)

	// Python's `sorted(..., reverse=True)` is equivalent to sorting in descending order.
	sort.Sort(sort.Reverse(sort.IntSlice(arrCopy)))

	var newArr []string
	// The try-except block is replaced by Go's idiomatic map lookup.
	for _, val := range arrCopy {
		if word, ok := dic[val]; ok {
			newArr = append(newArr, word)
		}
	}

	return newArr
}
