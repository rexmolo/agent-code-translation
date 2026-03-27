package main

import "sort"

func GetOddCollatz(n int) []int {
	var oddCollatz []int

	// Start with n if it's odd
	if n%2 == 1 {
		oddCollatz = append(oddCollatz, n)
	}

	// Generate Collatz sequence until we reach 1
	for n > 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = n*3 + 1
		}

		// Collect odd numbers in the sequence
		if n%2 == 1 {
			oddCollatz = append(oddCollatz, n)
		}
	}

	// Sort the result in increasing order
	sort.Slice(oddCollatz, func(i, j int) bool {
		return oddCollatz[i] < oddCollatz[j]
	})

	return oddCollatz
}
