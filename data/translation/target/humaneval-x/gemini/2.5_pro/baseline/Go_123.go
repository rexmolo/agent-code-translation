package main

import (
	"fmt"
	"sort"
)

// GetOddCollatz returns a sorted list of the odd numbers in the Collatz sequence for n.
// The Collatz conjecture concerns a sequence defined as follows: start with any
// positive integer n. Then each term is obtained from the previous term as follows:
// if the previous term is even, the next term is one half of the previous term. If
// the previous term is odd, the next term is 3 times the previous term plus 1.
// The conjecture is that no matter what value of n, the sequence will always reach 1.
func GetOddCollatz(n int) []int {
	if n <= 0 {
		// Collatz sequence is defined for positive integers.
		return []int{}
	}

	var oddCollatz []int

	for {
		// Check if the current number is odd and add it to the list.
		if n%2 != 0 {
			oddCollatz = append(oddCollatz, n)
		}

		// The sequence ends when it reaches 1.
		if n == 1 {
			break
		}

		// Calculate the next term in the sequence.
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
	}

	// Sort the collected odd numbers in increasing order.
	sort.Ints(oddCollatz)
	return oddCollatz
}

func main() {
	// Example from Python docstring: get_odd_collatz(5) returns [1, 5]
	result1 := GetOddCollatz(5)
	fmt.Printf("GetOddCollatz(5) -> %v\n", result1)

	// Example from Python docstring: Collatz(1) is [1]
	result2 := GetOddCollatz(1)
	fmt.Printf("GetOddCollatz(1) -> %v\n", result2)

	// Example with an even starting number
	result3 := GetOddCollatz(12)
	fmt.Printf("GetOddCollatz(12) -> %v\n", result3)
}
