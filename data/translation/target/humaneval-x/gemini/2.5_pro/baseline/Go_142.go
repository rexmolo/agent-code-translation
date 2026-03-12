package main

import (
	"fmt"
)

// SumSquares will take a slice of integers. For all entries in the slice, the function shall square the integer entry if its index is a 
// multiple of 3 and will cube the integer entry if its index is a multiple of 4 and not a multiple of 3. The function will not 
// change the entries in the slice whose indexes are not a multiple of 3 or 4. The function shall then return the sum of all entries. 
func SumSquares(lst []int) int {
	totalSum := 0
	for i, val := range lst {
		if i%3 == 0 {
			totalSum += val * val
		} else if i%4 == 0 {
			totalSum += val * val * val
		} else {
			totalSum += val
		}
	}
	return totalSum
}

// main is the entry point of the program. 
// It demonstrates the SumSquares function with examples from the original Python docstring.
func main() {
	lst1 := []int{1, 2, 3}
	fmt.Println(SumSquares(lst1)) // Expected: 6

	lst2 := []int{}
	fmt.Println(SumSquares(lst2)) // Expected: 0

	lst3 := []int{-1, -5, 2, -1, -5}
	fmt.Println(SumSquares(lst3)) // Expected: -126
}