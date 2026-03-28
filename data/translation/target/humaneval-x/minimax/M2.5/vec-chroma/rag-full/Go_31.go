package main

import "fmt"

func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	for k := 2; k < n-1; k++ {
		if n%k == 0 {
			return false
		}
	}
	return true
}

func main() {
	// Test cases from docstring
	fmt.Println(IsPrime(6))    // false
	fmt.Println(IsPrime(101))   // true
	fmt.Println(IsPrime(11))    // true
	fmt.Println(IsPrime(13441)) // true
	fmt.Println(IsPrime(61))    // true
	fmt.Println(IsPrime(4))     // false
	fmt.Println(IsPrime(1))     // false
}
