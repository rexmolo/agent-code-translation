package main

import (
	"fmt"
)

// PrimeLength takes a string and returns true if the string
// length is a prime number or false otherwise.
func PrimeLength(s string) bool {
	l := len(s)
	if l < 2 { // 0 and 1 are not prime numbers
		return false
	}
	// A number is prime if it's not divisible by any number from 2 to n-1.
	for i := 2; i < l; i++ {
		if l%i == 0 {
			return false
		}
	}
	return true
}

// main function to demonstrate the usage of PrimeLength
func main() {
	fmt.Println(PrimeLength("Hello"))
	fmt.Println(PrimeLength("abcdcba"))
	fmt.Println(PrimeLength("kittens"))
	fmt.Println(PrimeLength("orange"))
}
