package main

import "fmt"

func StrangeSortList(lst []int) []int {
	result := []int{}
	if len(lst) == 0 {
		return result
	}

	// Create a copy to work with (avoid modifying original)
	working := make([]int, len(lst))
	copy(working, lst)

	isMin := true

	for len(working) > 0 {
		var targetIdx int
		if isMin {
			// Find minimum
			targetIdx = 0
			for i := 1; i < len(working); i++ {
				if working[i] < working[targetIdx] {
					targetIdx = i
				}
			}
		} else {
			// Find maximum
			targetIdx = 0
			for i := 1; i < len(working); i++ {
				if working[i] > working[targetIdx] {
					targetIdx = i
				}
			}
		}

		result = append(result, working[targetIdx])

		// Remove element at targetIdx using slice manipulation
		working = append(working[:targetIdx], working[targetIdx+1:]...)

		// Toggle between min and max
		isMin = !isMin
	}

	return result
}

func main() {
	// Test examples
	fmt.Println(StrangeSortList([]int{}))
	fmt.Println(StrangeSortList([]int{1, 2, 3, 4}))
	fmt.Println(StrangeSortList([]int{5, 5, 5, 5}))
}
