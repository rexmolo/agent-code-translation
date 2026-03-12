package main

import "fmt"

// BelowThreshold returns true if all numbers in the slice l are below threshold t.
func BelowThreshold(l []int, t int) bool {
	for _, e := range l {
		if e >= t {
			return false
		}
	}
	return true
}

// main function to demonstrate the usage of BelowThreshold,
// mimicking the Python doctests.
func main() {
	fmt.Println(BelowThreshold([]int{1, 2, 4, 10}, 100)) // Expected: true
	fmt.Println(BelowThreshold([]int{1, 20, 4, 10}, 5))   // Expected: false
}
