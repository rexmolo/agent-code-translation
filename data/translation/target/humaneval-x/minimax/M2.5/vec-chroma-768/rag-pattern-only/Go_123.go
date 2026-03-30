package main

import (
	"fmt"
	"sort"
)

func GetOddCollatz(n int) []int {
	var oddCollatz []int

	// Initialize based on whether n is odd
	if n%2 == 1 {
		oddCollatz = append(oddCollatz, n)
	}

	// Generate Collatz sequence and collect odd numbers
	for n > 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = n*3 + 1
		}

		if n%2 == 1 {
			oddCollatz = append(oddCollatz, n)
		}
	}

	// Sort in increasing order
	sort.Ints(oddCollatz)


	return oddCollatz
}


func main() {
	fmt.Println(GetOddCollatz(5))  // Output: [1 5]
	fmt.Println(GetOddCollatz(6))  // Output: [1 3 5]
	fmt.Println(GetOddCollatz(1))  // Output: [1]
}
