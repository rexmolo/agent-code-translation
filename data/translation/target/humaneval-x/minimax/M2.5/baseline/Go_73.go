package main

import "fmt"

func SmallestChange(arr []int) int {
	ans := 0
	for i := 0; i < len(arr)/2; i++ {
		if arr[i] != arr[len(arr)-i-1] {
			ans++
		}
	}
	return ans
}

func main() {
	// Test cases
	fmt.Println(SmallestChange([]int{1, 2, 3, 5, 4, 7, 9, 6})) // Expected: 4
	fmt.Println(SmallestChange([]int{1, 2, 3, 4, 3, 2, 2}))   // Expected: 1
	fmt.Println(SmallestChange([]int{1, 2, 3, 2, 1}))          // Expected: 0
}