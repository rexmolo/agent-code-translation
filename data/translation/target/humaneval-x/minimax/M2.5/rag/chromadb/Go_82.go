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
	fmt.Println(PrimeLength("Hello"))      // true (5 is prime)
	fmt.Println(PrimeLength("abcdcba"))    // true (7 is prime)
	fmt.Println(PrimeLength("kittens"))    // true (7 is prime)
	fmt.Println(PrimeLength("orange"))     // false (6 is not prime)
	// Additional test cases
	fmt.Println(PrimeLength(""))            // false (0 is not prime)
	fmt.Println(PrimeLength("a"))           // false (1 is not prime)
	fmt.Println(PrimeLength("ab"))          // false (2 is prime, so this should be true)
	fmt.Println(PrimeLength("abc"))          // true (3 is prime)
}
