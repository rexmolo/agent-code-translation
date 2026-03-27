package main

import (
	"fmt"
)

func IncrList(l []int) []int {
	result := make([]int, len(l))
	for i, e := range l {
		result[i] = e + 1
	}
	return result
}

func main() {
	// Test cases
	fmt.Println(IncrList([]int{1, 2, 3}))
	fmt.Println(IncrList([]int{5, 3, 5, 2, 3, 3, 9, 0, 123}))
}