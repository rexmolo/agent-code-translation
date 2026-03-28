package main

import "fmt"

func Pluck(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}

	minVal := -1
	minIdx := -1

	for i, v := range arr {
		if v%2 == 0 {
			if minVal == -1 || v < minVal {
				minVal = v
				minIdx = i
			}
		}
	}

	if minVal == -1 {
		return []int{}
	}

	return []int{minVal, minIdx}
}

func main() {
	// Test cases
	fmt.Println(Pluck([]int{}))
	fmt.Println(Pluck([]int{4, 2, 3}))
	fmt.Println(Pluck([]int{1, 2, 3}))
	fmt.Println(Pluck([]int{5, 0, 3, 0, 4, 2}))
}