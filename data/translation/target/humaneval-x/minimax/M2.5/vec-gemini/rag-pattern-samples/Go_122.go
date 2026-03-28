package main

import (
	"fmt"
	"strconv"
)

func AddElements(arr []int, k int) int {
	sum := 0
	for i := 0; i < k; i++ {
		elem := arr[i]
		// Check if element has at most 2 digits
		// Convert to string and check length
		if len(strconv.Itoa(elem)) <= 2 {
			sum += elem
		}
	}
	return sum
}

func main() {
	// Example usage
	arr := []int{111, 21, 3, 4000, 5, 6, 7, 8, 9}
	k := 4
	result := AddElements(arr, k)
	fmt.Println(result) // Output: 24
}