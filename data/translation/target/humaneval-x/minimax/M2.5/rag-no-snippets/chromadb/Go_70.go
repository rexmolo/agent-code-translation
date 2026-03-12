package main

import (
	"slices"
)

func StrangeSortList(lst []int) []int {
	res := make([]int, 0, len(lst))
	switchVal := true

	for len(lst) > 0 {
		var val int
		if switchVal {
			val = slices.Min(lst)
		} else {
			val = slices.Max(lst)
		}
		res = append(res, val)

		// Find and remove the first occurrence of the value
		for i, v := range lst {
			if v == val {
				lst = slices.Delete(lst, i, i+1)
				break
			}
		}

		switchVal = !switchVal
	}

	return res
}