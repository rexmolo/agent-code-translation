package main

import "fmt"

func IsMultiplyPrime(a int) bool {
	// Helper function to check if a number is prime
	isPrime := func(n int) bool {
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

	// Check all combinations of three primes from 2 to 100
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
	fmt.Println(IsMultiplyPrime(30)) // true: 30 = 2 * 3 * 5
}
