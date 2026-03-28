package main

import "fmt"

func Pluck(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}

	// Find all even values
	var evens []int
	for _, v := range arr {
		if v%2 == 0 {
			evens = append(evens, v)
		}
	}

	if len(evens) == 0 {
		return []int{}
	}

	// Find minimum even value
	minEven := evens[0]
	for _, v := range evens {
		if v < minEven {
			minEven = v
		}
	}

	// Find first index of minEven in original array
	for i, v := range arr {
		if v == minEven {
			return []int{minEven, i}
		}
	}

	return []int{}
}

func main() {
	// Test cases
	fmt.Println(Pluck([]int{4, 2, 3}))      // [2 1]
	fmt.Println(Pluck([]int{1, 2, 3}))      // [2 1]
	fmt.Println(Pluck([]int{}))             // []
	fmt.Println(Pluck([]int{5, 0, 3, 0, 4, 2})) // [0 1]
}
