package main

import "fmt"

func CanArrange(arr []int) int {
	ind := -1
	i := 1
	for i < len(arr) {
		if arr[i] < arr[i-1] {
			ind = i
		}
		i++
	}
	return ind
}

func main() {
	// Test examples
	fmt.Println(CanArrange([]int{1, 2, 4, 3, 5})) // Output: 3
	fmt.Println(CanArrange([]int{1, 2, 3}))      // Output: -1
}
