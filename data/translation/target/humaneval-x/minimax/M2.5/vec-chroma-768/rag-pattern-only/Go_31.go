package main

import (
	"fmt"
	"math"
	"os"
)

func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	// Check for even numbers and check up to sqrt(n)
	if n%2 == 0 {
		return false
	}
	limit := int(math.Sqrt(float64(n)))
	for k := 3; k <= limit; k += 2 {
		if n%k == 0 {
			return false
		}
	}
	return true
}

func main() {
	// Test cases from the docstring
	testCases := []struct {
		n      int
		expect bool
	}{
		{6, false},
		{101, true},
		{11, true},
		{13441, true},
		{61, true},
		{4, false},
		{1, false},
	}

	for _, tc := range testCases {
		result := IsPrime(tc.n)
		if result != tc.expect {
			fmt.Printf("IsPrime(%d) = %v, expected %v\n", tc.n, result, tc.expect)
			os.Exit(1)
		}
	}
	fmt.Println("All tests passed!")
}