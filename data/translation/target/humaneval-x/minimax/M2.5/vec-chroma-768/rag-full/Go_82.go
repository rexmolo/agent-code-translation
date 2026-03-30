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
	fmt.Println(PrimeLength("Hello"))      // true
	fmt.Println(PrimeLength("abcdcba"))     // true
	fmt.Println(PrimeLength("kittens"))     // true
	fmt.Println(PrimeLength("orange"))       // false
}