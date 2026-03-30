package main

import (
	"fmt"
	"sort"
)

func GetOddCollatz(n int) []int {
	var oddCollatz []int

	if n%2 != 0 {
		oddCollatz = []int{n}
	}

	for n > 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = n*3 + 1
		}

		if n%2 != 0 {
			oddCollatz = append(oddCollatz, n)
		}
	}

	sort.Slice(oddCollatz, func(i, j int) bool {
		return oddCollatz[i] < oddCollatz[j]
	})

	return oddCollatz
}

func main() {
	// Test examples
	fmt.Println(GetOddCollatz(5))  // Output: [1 5]
	fmt.Println(GetOddCollatz(8))  // Output: [1]
	fmt.Println(GetOddCollatz(1))  // Output: [1]
}