package main

import (
	"fmt"
	"sort"
)

func GetOddCollatz(n int) []int {
	var oddCollatz []int

	// If n is odd initially, add it to the list
	if n%2 != 0 {
		oddCollatz = append(oddCollatz, n)
	}

	// Generate Collatz sequence
	for n > 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = n*3 + 1
		}

		// If n is odd, add it to the list
		if n%2 != 0 {
			oddCollatz = append(oddCollatz, n)
		}
	}

	// Sort the result in increasing order
	sort.Ints(oddCollatz)

	return oddCollatz
}

func main() {
	// Example usage
	result := GetOddCollatz(5)
	fmt.Println(result)
}