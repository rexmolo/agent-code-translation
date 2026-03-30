package main

import "fmt"

func BelowThreshold(l []int, t int) bool {
	for _, e := range l {
		if e >= t {
			return false
		}
	}
	return true
}

func main() {
	// Test cases
	fmt.Println(BelowThreshold([]int{1, 2, 4, 10}, 100)) // Expected: true
	fmt.Println(BelowThreshold([]int{1, 20, 4, 10}, 5))  // Expected: false
}