package main

import "fmt"

// IsPrime returns true if a given number is prime, and false otherwise.
func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	// The original Python loop `for k in range(2, n - 1)` iterates from 2
	// up to, but not including, n-1. The Go loop `for k := 2; k < n-1; k++`
	// preserves this exact logic.
	for k := 2; k < n-1; k++ {
		if n%k == 0 {
			return false
		}
	}
	return true
}

// main function to demonstrate IsPrime with the examples from the Python docstring.
func main() {
	fmt.Printf("is_prime(6) -> %t\n", IsPrime(6))         // Expected: False
	fmt.Printf("is_prime(101) -> %t\n", IsPrime(101))     // Expected: True
	fmt.Printf("is_prime(11) -> %t\n", IsPrime(11))       // Expected: True
	fmt.Printf("is_prime(13441) -> %t\n", IsPrime(13441)) // Expected: True
	fmt.Printf("is_prime(61) -> %t\n", IsPrime(61))       // Expected: True
	fmt.Printf("is_prime(4) -> %t\n", IsPrime(4))         // Expected: False
	fmt.Printf("is_prime(1) -> %t\n", IsPrime(1))         // Expected: False
}
