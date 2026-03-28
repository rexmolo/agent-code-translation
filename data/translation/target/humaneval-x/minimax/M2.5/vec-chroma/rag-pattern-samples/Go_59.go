package main

import "fmt"

func main() {
	fmt.Println(LargestPrimeFactor(13195))
	fmt.Println(LargestPrimeFactor(2048))
}

func LargestPrimeFactor(n int) int {
	largest := 1

	// Handle factor of 2 separately
	for n%2 == 0 {
		largest = 2
		n /= 2
	}

	// Check odd factors from 3
	for i := 3; i*i <= n; i += 2 {
		for n%i == 0 {
			largest = i
			n /= i
		}
	}

	// If n > 1, then it's a prime factor larger than sqrt(original n)
	if n > 1 {
		largest = n
	}

	return largest
}
