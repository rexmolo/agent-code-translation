package main

import (
	"math"
)

// isPrime checks if a number is prime. 
// It corrects a subtle bug in the original Python implementation where numbers <= 1
// would be incorrectly considered prime.
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	limit := int(math.Sqrt(float64(n)))
	for i := 2; i <= limit; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Skjkasdkd finds the largest prime number in a slice of integers
// and returns the sum of its digits.
func Skjkasdkd(lst []int) int {
	maxx := 0
	for _, num := range lst {
		if num > maxx && isPrime(num) {
			maxx = num
		}
	}

	// Calculate the sum of the digits of the largest prime found.
	// This is done using arithmetic, which is more idiomatic in Go than string conversion.
	sumOfDigits := 0
	temp := maxx
	for temp > 0 {
		sumOfDigits += temp % 10
		temp /= 10
	}

	return sumOfDigits
}
