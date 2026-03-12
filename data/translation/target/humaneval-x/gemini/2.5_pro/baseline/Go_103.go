package main

import (
	"fmt"
	"math"
)

// RoundedAvg computes the average of integers from n to m, rounds the result to the
// nearest integer, and returns its binary representation.
// It replicates the logic of the Python function, including the rounding behavior.
// If n is greater than m, it returns -1.
func RoundedAvg(n, m int) interface{} {
	// If n is greater than m, return -1.
	if m < n {
		return -1
	}

	// The average of an arithmetic series (like integers from n to m) is simply
	// the average of the first and last terms. This is more efficient than summing
	// all numbers in a loop.
	// We use float64 for the calculation to handle potential decimal results.
	avg := float64(n+m) / 2.0

	// Python's `round()` function uses "round half to even" logic.
	// For example, round(2.5) is 2, and round(3.5) is 4.
	// The example `rounded_avg(20, 33)` has an average of 26.5, which Python rounds
	// to 26, confirming this behavior.
	// Go's `math.RoundToEven` function, available since Go 1.10, does exactly this.
	rounded := int(math.RoundToEven(avg))

	// Convert the rounded integer to its binary string representation,
	// prefixed with "0b" to match Python's `bin()` output.
	return fmt.Sprintf("0b%b", rounded)
}

// main function to test the RoundedAvg function with examples.
func main() {
	// Example: rounded_avg(1, 5) => "0b11"
	fmt.Println(RoundedAvg(1, 5))

	// Example: rounded_avg(7, 5) => -1
	fmt.Println(RoundedAvg(7, 5))

	// Example: rounded_avg(10, 20) => "0b1111"
	fmt.Println(RoundedAvg(10, 20))

	// Example: rounded_avg(20, 33) => "0b11010"
	fmt.Println(RoundedAvg(20, 33))
}
