package main

import "fmt"

// StrangeSortList returns a list in "strange" order where elements
// are picked alternating between minimum and maximum of remaining elements.
// Start with min, then max, then min, and so on.
func StrangeSortList(lst []int) []int {
	res := make([]int, 0, len(lst))
	remaining := make([]int, len(lst))
	copy(remaining, lst)
	isMin := true

	for len(remaining) > 0 {
		// Find min or max based on the current flag
		var extreme int
		if isMin {
			extreme = remaining[0]
			for _, v := range remaining {
				if v < extreme {
					extreme = v
				}
			}
		} else {
			extreme = remaining[0]
			for _, v := range remaining {
				if v > extreme {
					extreme = v
				}
			}
		}

		res = append(res, extreme)

		// Remove the extreme value from remaining
		for i, v := range remaining {
			if v == extreme {
				remaining = append(remaining[:i], remaining[i+1:]...)
				break
			}
		}

		isMin = !isMin
	}

	return res
}

func main() {
	// Test cases
	fmt.Println(fmt.Sprintf("%v", StrangeSortList([]int{1, 2, 3, 4}))) // [1 4 2 3]
	fmt.Println(fmt.Sprintf("%v", StrangeSortList([]int{5, 5, 5, 5}))) // [5 5 5 5]
	fmt.Println(fmt.Sprintf("%v", StrangeSortList([]int{})))          // []
}