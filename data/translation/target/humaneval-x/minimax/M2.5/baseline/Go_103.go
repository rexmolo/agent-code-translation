package main

import (
	"fmt"
	"math"
)

func RoundedAvg(n, m int) interface{} {
	if m < n {
		return -1
	}

	// Calculate average of arithmetic series: (n + m) / 2
	avg := float64(n+m) / 2.0

	// Round to nearest integer (banker's rounding - same as Python's round())
	rounded := int(math.Round(avg))

	// Convert to binary string with "0b" prefix
	return fmt.Sprintf("0b%d", rounded)
}

func main() {
	// Test cases
	fmt.Println(RoundedAvg(1, 5))   // "0b11"
	fmt.Println(RoundedAvg(7, 5))   // -1
	fmt.Println(RoundedAvg(10, 20)) // "0b1111"
	fmt.Println(RoundedAvg(20, 33)) // "0b11010"
}
