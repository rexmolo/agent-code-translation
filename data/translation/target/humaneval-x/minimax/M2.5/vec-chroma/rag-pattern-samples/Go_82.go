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
	// Test cases from the examples
	fmt.Println(PrimeLength("Hello"))    // True
	fmt.Println(PrimeLength("abcdcba"))  // True
	fmt.Println(PrimeLength("kittens"))  // True
	fmt.Println(PrimeLength("orange"))   // False
}
