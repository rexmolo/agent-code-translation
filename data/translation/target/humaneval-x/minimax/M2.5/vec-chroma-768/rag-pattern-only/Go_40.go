package main

import "fmt"

// TriplesSumToZero checks if any three distinct elements in the list sum to zero.
// Returns true if such a triplet exists, false otherwise.
func TriplesSumToZero(l []int) bool {
	n := len(l)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if l[i]+l[j]+l[k] == 0 {
					return true
				}
			}
		}
	}
	return false
}

func main() {
	// Example usage
	testCases := [][]int{
		{1, 3, 5, 0},
		{1, 3, -2, 1},
		{1, 2, 3, 7},
		{2, 4, -5, 3, 9, 7},
		{1},
	}
	for _, tc := range testCases {
		fmt.Printf("TriplesSumToZero(%v) = %v\n", tc, TriplesSumToZero(tc))
	}
}
