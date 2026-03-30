package main

import "fmt"

func SumToN(n int) int {
	return n * (n + 1) / 2
}

func main() {
	// Test cases from Python docstring
	testCases := []struct {
		input    int
		expected int
	}{
		{30, 465},
		{100, 5050},
		{5, 15},
		{10, 55},
		{1, 1},
	}

	for _, tc := range testCases {
		result := SumToN(tc.input)
		fmt.Printf("SumToN(%d) = %d (expected %d)\n", tc.input, result, tc.expected)
	}
}