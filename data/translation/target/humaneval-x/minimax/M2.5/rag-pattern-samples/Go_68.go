package main

import "fmt"

func Pluck(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}

	minEven := -1
	minIndex := -1

	for i, v := range arr {
		if v%2 == 0 {
			if minEven == -1 || v < minEven {
				minEven = v
				minIndex = i
			}
		}
	}

	if minEven == -1 {
		return []int{}
	}

	return []int{minEven, minIndex}
}

func main() {
	// Test cases
	fmt.Println(Pluck([]int{}))                   // []
	fmt.Println(Pluck([]int{4, 2, 3}))            // [2, 1]
	fmt.Println(Pluck([]int{1, 2, 3}))            // [2, 1]
	fmt.Println(Pluck([]int{5, 0, 3, 0, 4, 2}))  // [0, 1]
}
