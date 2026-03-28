package main

import (
	"sort"
)

func GetOddCollatz(n int) []int {
	var oddCollatz []int

	// If n is odd, start with n in the list
	if n%2 != 0 {
		oddCollatz = []int{n}
	}

	// Generate Collatz sequence until we reach 1
	for n > 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = n*3 + 1
		}

		// If the new term is odd, add it to the list
		if n%2 != 0 {
			oddCollatz = append(oddCollatz, n)
		}
	}

	// Sort the odd numbers in increasing order
	sort.Ints(oddCollatz)

	return oddCollatz
}