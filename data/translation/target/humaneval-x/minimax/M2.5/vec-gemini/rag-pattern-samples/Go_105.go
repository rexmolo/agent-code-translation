package main

import (
	"fmt"
	"sort"
)

func ByLength(arr []int) []string {
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

	// Sort in descending order (equivalent to Python's sorted(arr, reverse=True))
	sortedArr := make([]int, len(arr))
	copy(sortedArr, arr)
	sort.Slice(sortedArr, func(i, j int) bool {
		return sortedArr[i] > sortedArr[j]
	})

	// Build result by filtering valid numbers (1-9) and mapping to strings
	// Equivalent to the try/except block in Python
	var result []string
	for _, v := range sortedArr {
		if name, ok := dic[v]; ok {
			result = append(result, name)
		}
	}

	return result
}

func main() {
	// Test examples
	examples := [][]int{
		{2, 1, 1, 4, 5, 8, 2, 3},
		{},
		{1, -1, 55},
	}

	for _, arr := range examples {
		result := ByLength(arr)
		fmt.Printf("Input: %v -> Output: %v\n", arr, result)
	}
}