package main

import (
	"fmt"
)

// CountUpTo takes a non-negative integer and returns a slice of the
// prime numbers that are less than n.
// for example:
// CountUpTo(5) => [2,3]
// CountUpTo(11) => [2,3,5,7]
// CountUpTo(0) => []
// CountUpTo(20) => [2,3,5,7,11,13,17,19]
// CountUpTo(1) => []
// CountUpTo(18) => [2,3,5,7,11,13,17]
func CountUpTo(n int) []int {
	primes := []int{}
	for i := 2; i < n; i++ {
		isPrime := true
		for j := 2; j < i; j++ {
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

// main function is included to create a runnable program and demonstrate usage.
func main() {
	fmt.Printf("count_up_to(5) => %v\n", CountUpTo(5))
	fmt.Printf("count_up_to(11) => %v\n", CountUpTo(11))
	fmt.Printf("count_up_to(0) => %v\n", CountUpTo(0))
	fmt.Printf("count_up_to(20) => %v\n", CountUpTo(20))
	fmt.Printf("count_up_to(1) => %v\n", CountUpTo(1))
	fmt.Printf("count_up_to(18) => %v\n", CountUpTo(18))
}
