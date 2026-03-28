package main

import (
	"fmt"
	"sort"
)

func GetOddCollatz(n int) []int {
	var oddCollatz []int

	// Add n if it's odd (initial number in the sequence)
	if n%2 != 0 {
		oddCollatz = append(oddCollatz, n)
	}

	// Generate the Collatz sequence
	for n > 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = n*3 + 1
		}

		// If the new term is odd, add it to the result
		if n%2 != 0 {
			oddCollatz = append(oddCollatz, n)
		}
	}

	// Sort the result in increasing order
	sort.Ints(oddCollatz)

	return oddCollatz
}

func main() {
	// Test the function
	fmt.Println(GetOddCollatz(5)) // Output: [1 5]
}
