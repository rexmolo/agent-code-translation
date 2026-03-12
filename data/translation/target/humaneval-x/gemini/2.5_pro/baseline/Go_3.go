package main

import (
	"fmt"
)

// BelowZero checks if the balance of an account ever falls below zero given a series of operations.
// The account starts with a zero balance.
// It returns true if the balance drops below zero at any point, otherwise it returns false.
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
	// Example from docstring: below_zero([1, 2, 3]) -> False
	fmt.Println(BelowZero([]int{1, 2, 3}))

	// Example from docstring: below_zero([1, 2, -4, 5]) -> True
	fmt.Println(BelowZero([]int{1, 2, -4, 5}))
}