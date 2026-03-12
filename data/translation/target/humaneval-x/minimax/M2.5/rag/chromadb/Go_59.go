package main

import "fmt"

func LargestPrimeFactor(n int) int {
	// Return the largest prime factor of n. Assume n > 1 and is not a prime.
	// >>> largest_prime_factor(13195)
	// 29
	// >>> largest_prime_factor(2048)
	// 2

	isPrime := func(k int) bool {
		if k < 2 {
			return false
		}
		for i := 2; i < k-1; i++ {
			if k%i == 0 {
				return false
			}
		}
		return true
	}

	largest := 1
	for j := 2; j <= n; j++ {
		if n%j == 0 && isPrime(j) {
			if j > largest {
				largest = j
			}
		}
	}
	return largest
}

func main() {
	fmt.Println(LargestPrimeFactor(13195)) // 29
	fmt.Println(LargestPrimeFactor(2048))   // 2
}
