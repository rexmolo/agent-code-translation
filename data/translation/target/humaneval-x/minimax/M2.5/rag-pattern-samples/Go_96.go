package main

import "fmt"

func CountUpTo(n int) []int {
	primes := []int{}
	for i := 2; i < n; i++ {
		isPrime := true
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
	// Test cases
	testCases := []int{5, 11, 0, 20, 1, 18}
	for _, n := range testCases {
		fmt.Printf("count_up_to(%d) => %v\n", n, CountUpTo(n))
	}
}