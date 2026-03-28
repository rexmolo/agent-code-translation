package main

// isPrime checks if a number is prime
func isPrime(n int) bool {
	for j := 2; j < n; j++ {
		if n%j == 0 {
			return false
		}
	}
	return true
}

// IsMultiplyPrime returns true if the given number is the multiplication
// of 3 prime numbers and false otherwise.
// a is less than 100.
// Example:
// IsMultiplyPrime(30) == true
// 30 = 2 * 3 * 5
func IsMultiplyPrime(a int) bool {
	for i := 2; i < 101; i++ {
		if !isPrime(i) {
			continue
		}
		for j := 2; j < 101; j++ {
			if !isPrime(j) {
				continue
			}
			for k := 2; k < 101; k++ {
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