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
	// Example usage for testing
	fmt.Println(BelowZero([]int{1, 2, 3}))      // false
	fmt.Println(BelowZero([]int{1, 2, -4, 5})) // true
}
