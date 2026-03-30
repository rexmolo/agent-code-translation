package main

import "fmt"

// isPrime checks if a number is prime
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

// IsMultiplyPrime returns true if the given number is the multiplication
// of 3 prime numbers and false otherwise.
// Example: IsMultiplyPrime(30) == true (30 = 2 * 3 * 5)
func IsMultiplyPrime(a int) bool {
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

func main() {
	// Test the function
	fmt.Println(IsMultiplyPrime(30)) // true
	fmt.Println(IsMultiplyPrime(27)) // false
}