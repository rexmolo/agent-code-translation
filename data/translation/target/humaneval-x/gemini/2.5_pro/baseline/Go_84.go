package main

import (
	"fmt"
	"strconv"
)

// Solve calculates the sum of the digits of N and returns that sum represented as a binary string.
// This function implements the logic from the original Python code.
func Solve(N int) string {
	// The Python equivalent `str(0)` -> `sum([0])` -> `bin(0)[2:]` results in "0".
	// The loop below would produce sum 0, so this explicit check is for clarity
	// but not strictly necessary if N >= 0.
	if N == 0 {
		return "0"
	}

	digitSum := 0
	tempN := N
	// To sum the digits, we can repeatedly take the number modulo 10 to get the last digit
	// and then divide by 10 to remove the last digit, until the number is 0.
	// This is a more idiomatic Go approach for numerical operations than converting to a string.
	for tempN > 0 {
		digitSum += tempN % 10
		tempN /= 10
	}

	// strconv.FormatInt converts an integer to a string in a given base.
	// The base 2 argument specifies a binary representation.
	// This is the Go equivalent of Python's `bin(number)[2:]`.
	return strconv.FormatInt(int64(digitSum), 2)
}

// main function to make the code a runnable program.
// It reads an integer from standard input, calls Solve, and prints the output.
func main() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		// In case of non-integer input, you might want to handle the error.
		// For this problem, we assume valid input as per constraints.
		return
	}
	fmt.Println(Solve(n))
}
