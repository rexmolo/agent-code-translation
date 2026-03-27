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

	// Create a copy to sort (to avoid modifying original slice)
	sortedArr := make([]int, len(arr))
	copy(sortedArr, arr)

	// Sort in reverse order
	sort.Slice(sortedArr, func(i, j int) bool {
		return sortedArr[i] > sortedArr[j]
	})

	// Filter valid numbers (1-9) and map to string names
	var result []string
	for _, v := range sortedArr {
		if val, ok := dic[v]; ok {
			result = append(result, val)
		}
	}

	return result
}

func main() {
	// Test cases
	fmt.Println(ByLength([]int{2, 1, 1, 4, 5, 8, 2, 3}))
	fmt.Println(ByLength([]int{}))
	fmt.Println(ByLength([]int{1, -1, 55}))
}
