package main

import "fmt"

func BelowZero(operations []int) bool {
	balance := 0

	for _, op := range operations {
		balance += op
		if balance < 0 {
			return true
		}
	}

	return false
}

func main() {
	// Test cases
	fmt.Println(BelowZero([]int{1, 2, 3}))       // Expected: false
	fmt.Println(BelowZero([]int{1, 2, -4, 5}))   // Expected: true
}