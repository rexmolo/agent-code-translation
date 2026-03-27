package main

import "fmt"

func PrimeLength(s string) bool {
	l := len(s)
	if l == 0 || l == 1 {
		return false
	}
	for i := 2; i < l; i++ {
		if l%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	// Test cases
	fmt.Println(PrimeLength("Hello"))    // True (5 is prime)
	fmt.Println(PrimeLength("abcdcba")) // True (7 is prime)
	fmt.Println(PrimeLength("kittens")) // True (7 is prime)
	fmt.Println(PrimeLength("orange"))   // False (6 is not prime)
}
