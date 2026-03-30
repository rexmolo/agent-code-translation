package main

import "fmt"

func LargestDivisor(n int) int {
	// Find the largest divisor of n that is smaller than n
	for i := n - 1; i >= 1; i-- {
		if n%i == 0 {
			return i
		}
	}
	return 0
}

func main() {
	fmt.Println(LargestDivisor(15))
}
