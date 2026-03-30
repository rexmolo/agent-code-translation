package main

import (
	"fmt"
	"strconv"
)

func AddElements(arr []int, k int) int {
	sum := 0
	for _, elem := range arr[:k] {
		if len(strconv.Itoa(elem)) <= 2 {
			sum += elem
		}
	}
	return sum
}

func main() {
	// Test with example from docstring
	arr := []int{111, 21, 3, 4000, 5, 6, 7, 8, 9}
	k := 4
	result := AddElements(arr, k)
	fmt.Println(result) // Output: 24
}