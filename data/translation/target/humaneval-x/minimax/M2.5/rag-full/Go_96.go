package main

import "fmt"

func CountUpTo(n int) []int {
	primes := []int{}
	for i := 2; i < n; i++ {
		isPrime := true
		// Only need to check up to sqrt(i) for primality
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primes = append(primes, i)
		}
	}
	return primes
}

func main() {
	// Test cases based on docstring examples
	fmt.Println(CountUpTo(5))   // => [2 3]
	fmt.Println(CountUpTo(11))  // => [2 3 5 7]
	fmt.Println(CountUpTo(0))   // => []
	fmt.Println(CountUpTo(20))  // => [2 3 5 7 11 13 17 19]
	fmt.Println(CountUpTo(1))   // => []
	fmt.Println(CountUpTo(18))  // => [2 3 5 7 11 13 17]
}
