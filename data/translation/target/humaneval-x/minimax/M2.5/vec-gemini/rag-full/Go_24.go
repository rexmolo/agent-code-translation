package main

import "fmt"

func LargestDivisor(n int) int {
	// Find the largest divisor smaller than n
	for i := n - 1; i >= 1; i-- {
		if n%i == 0 {
			return i
		}
	}
	return 0 // No divisor found (e.g., for n <= 1)
}

func main() {
	fmt.Println(LargestDivisor(15))
}