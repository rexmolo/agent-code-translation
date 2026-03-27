package main

import (
	"fmt"
	"slices"
)

func StrangeSortList(lst []int) []int {
	res := make([]int, 0, len(lst))
	switchMin := true

	for len(lst) > 0 {
		var val int
		if switchMin {
			val = slices.Min(lst)
		} else {
			val = slices.Max(lst)
		}
		res = append(res, val)

		// Remove the first occurrence of val from lst
		for i := 0; i < len(lst); i++ {
			if lst[i] == val {
				lst = append(lst[:i], lst[i+1:]...)
				break
			}
		}

		switchMin = !switchMin
	}
	return res
}

func main() {
	// Test cases
	fmt.Println(StrangeSortList([]int{1, 2, 3, 4})) // [1 4 2 3]
	fmt.Println(StrangeSortList([]int{5, 5, 5, 5})) // [5 5 5 5]
	fmt.Println(StrangeSortList([]int{}))           // []
}
