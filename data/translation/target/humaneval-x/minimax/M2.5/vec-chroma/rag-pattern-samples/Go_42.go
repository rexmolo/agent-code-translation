package main

import (
	"fmt"
)

func IncrList(l []int) []int {
	result := make([]int, 0, len(l))
	for _, e := range l {
		result = append(result, e+1)
	}
	return result
}

func main() {
	// Test cases from docstring
	fmt.Println(IncrList([]int{1, 2, 3}))
	fmt.Println(IncrList([]int{5, 3, 5, 2, 3, 3, 9, 0, 123}))
}