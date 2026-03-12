package main

import (
	"fmt"
	"sort"
)

func SortArray(array []int) []int {
	if len(array) < 2 {
		newSlice := make([]int, len(array))
		copy(newSlice, array)
		return newSlice
	}

	result := make([]int, len(array))
	copy(result, array)

	sum := array[0] + array[len(array)-1]

	if sum%2 == 0 {
		sort.Sort(sort.Reverse(sort.IntSlice(result)))
	} else {
		sort.Ints(result)
	}

	return result
}

func main() {
	fmt.Println(SortArray([]int{}))
	fmt.Println(SortArray([]int{5}))
	fmt.Println(SortArray([]int{2, 4, 3, 0, 1, 5}))
	fmt.Println(SortArray([]int{2, 4, 3, 0, 1, 5, 6}))
}