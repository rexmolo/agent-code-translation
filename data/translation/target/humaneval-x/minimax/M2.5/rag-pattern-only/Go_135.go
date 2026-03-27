package main

import "fmt"

func CanArrange(arr []int) int {
	ind := -1
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			ind = i
		}
	}
	return ind
}

func main() {
	fmt.Println(CanArrange([]int{1, 2, 4, 3, 5})) // Output: 3
	fmt.Println(CanArrange([]int{1, 2, 3}))        // Output: -1
}