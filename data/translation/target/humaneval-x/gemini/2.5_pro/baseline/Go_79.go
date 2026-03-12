package main

import (
	"fmt"
)

// DecimalToBinary converts a decimal integer to its binary string representation,
// wrapped with "db" at the beginning and end.
//
// Examples:
// DecimalToBinary(15)   // returns "db1111db"
// DecimalToBinary(32)   // returns "db100000db"
func DecimalToBinary(decimal int) string {
	// In Go, fmt.Sprintf with the %b verb converts an integer to its binary string representation.
	// This is equivalent to Python's bin(decimal)[2:].
	binaryString := fmt.Sprintf("%b", decimal)

	// Concatenate the prefix, binary string, and suffix.
	return "db" + binaryString + "db"
}

// main function to demonstrate the DecimalToBinary function.
func main() {
	fmt.Println(DecimalToBinary(15))
	fmt.Println(DecimalToBinary(32))
}
