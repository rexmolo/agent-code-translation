package main

import (
	"fmt"
	"slices"
)

func Pluck(arr []int) []int {
	if len(arr) == 0 {
		return nil
	}
	
	// Filter even numbers (manually, as there's no direct filter equivalent)
	var evens []int
	for _, v := range arr {
		if v%2 == 0 {
			evens = append(evens, v)
		}
	}
	
	if len(evens) == 0 {
		return nil
	}
	
	// Find minimum value using slices.Min
	minVal := slices.Min(evens)
	
	// Find index of the minimum value in original array using slices.Index
	index := slices.Index(arr, minVal)
	
	return []int{minVal, index}
}

func main() {
	// Test the function
	fmt.Println(Pluck([]int{4, 2, 3}))        // [2, 1]
	fmt.Println(Pluck([]int{1, 2, 3}))        // [2, 1]
	fmt.Println(Pluck([]int{}))               // []
	fmt.Println(Pluck([]int{5, 0, 3, 0, 4, 2})) // [0, 1]
}