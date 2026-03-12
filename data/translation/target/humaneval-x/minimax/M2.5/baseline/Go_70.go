package main

import "fmt"

func StrangeSortList(lst []int) []int {
	res := make([]int, 0, len(lst))
	switchMin := true

	for len(lst) > 0 {
		var val int
		if switchMin {
			val = min(lst)
		} else {
			val = max(lst)
		}
		res = append(res, val)
		lst = remove(lst, val)
		switchMin = !switchMin
	}
	return res
}

func min(lst []int) int {
	minVal := lst[0]
	for _, v := range lst {
		if v < minVal {
			minVal = v
		}
	}
	return minVal
}

func max(lst []int) int {
	maxVal := lst[0]
	for _, v := range lst {
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal
}

func remove(lst []int, val int) []int {
	for i, v := range lst {
		if v == val {
			return append(lst[:i], lst[i+1:]...)
		}
	}
	return lst
}

func main() {
	// Test cases
	fmt.Println(StrangeSortList([]int{1, 2, 3, 4})) // [1 4 2 3]
	fmt.Println(StrangeSortList([]int{5, 5, 5, 5})) // [5 5 5 5]
	fmt.Println(StrangeSortList([]int{}))           // []
}
