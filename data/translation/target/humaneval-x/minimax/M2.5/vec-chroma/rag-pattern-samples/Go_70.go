package main

import (
	"fmt"
	"sort"
)

func StrangeSortList(lst []int) []int {
	if len(lst) == 0 {
		return []int{}
	}

	// Sort a copy of the input
	arr := make([]int, len(lst))
	copy(arr, lst)
	sort.Ints(arr)

	result := []int{}
	left, right := 0, len(arr)-1
	useMin := true

	for left <= right {
		if useMin {
			result = append(result, arr[left])
			left++
		} else {
			result = append(result, arr[right])
			right--
		}
		useMin = !useMin
	}

	return result
}

func main() {
	// Test examples
	fmt.Println(StrangeSortList([]int{1, 2, 3, 4})) // [1 4 2 3]
	fmt.Println(StrangeSortList([]int{5, 5, 5, 5})) // [5 5 5 5]
	fmt.Println(StrangeSortList([]int{}))           // []
}