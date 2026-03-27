package main

import "fmt"

func IsMultiplyPrime(a int) bool {
	// Write a function that returns true if the given number is the multiplication of 3 prime numbers
	// and false otherwise.
	// Knowing that (a) is less then 100.
	// Example:
	// isMultiplyPrime(30) == true
	// 30 = 2 * 3 * 5

	for i := 2; i <= 100; i++ {
		if !isPrime(i) {
			continue
		}
		for j := 2; j <= 100; j++ {
			if !isPrime(j) {
				continue
			}
			for k := 2; k <= 100; k++ {
				if !isPrime(k) {
					continue
				}
				if i*j*k == a {
					return true
				}
			}
		}
	}
	return false
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for j := 2; j < n; j++ {
		if n%j == 0 {
			return false
		}
	}
	return true
}

func main() {
	// Test cases
	fmt.Println(IsMultiplyPrime(30))  // true: 2 * 3 * 5
	fmt.Println(IsMultiplyPrime(12))  // false
	fmt.Println(IsMultiplyPrime(100)) // false
}