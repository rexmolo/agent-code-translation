package main

import (
	"fmt"
)

func RoundedAvg(n, m int) interface{} {
	if m < n {
		return -1
	}

	// Calculate sum of integers from n to m (inclusive)
	// Using arithmetic series formula: (first + last) * count / 2
	count := m - n + 1
	summation := (n + m) * count / 2

	// Calculate average (using integer division for rounding)
	rounded := summation / count

	// Convert to binary string with "0b" prefix
	return fmt.Sprintf("0b%b", rounded)
}

func main() {
	// Test cases
	fmt.Println(RoundedAvg(1, 5))   // "0b11"
	fmt.Println(RoundedAvg(7, 5))  // -1
	fmt.Println(RoundedAvg(10, 20)) // "0b1111"
	fmt.Println(RoundedAvg(20, 33)) // "0b11010"
}
