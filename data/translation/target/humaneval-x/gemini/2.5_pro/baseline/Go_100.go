package main

import "fmt"

// MakeAPile creates a slice representing a pile of stones with n levels.
// The first level has n stones. The number of stones in each subsequent
// level increases by 2 from the previous level.
func MakeAPile(n int) []int {
	// Pre-allocate a slice of size n to store the number of stones for each level.
	pile := make([]int, n)

	// Loop n times, corresponding to the n levels.
	for i := 0; i < n; i++ {
		// The number of stones in level (i+1) is calculated as n + 2*i.
		pile[i] = n + 2*i
	}

	return pile
}

func main() {
	// Example usage from the docstring.
	result1 := MakeAPile(3)
	fmt.Println(result1)

	// Another example to test an even number.
	result2 := MakeAPile(4)
	fmt.Println(result2)
}
