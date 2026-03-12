package main

import (
	"fmt"
)

func IsMultiplyPrime(a int) bool {
	isPrime := func(n int) bool {
		for j := 2; j < n; j++ {
			if n%j == 0 {
				return false
			}
		}
		return true
	}

	for i := 2; i < 101; i++ {
		if !isPrime(i) {
			continue
		}
		for j := 2; j < 101; j++ {
			if !isPrime(j) {
				continue
			}
			for k := 2; k < 101; k++ {
				if !isPrime(k) {
					continue
				}
				if i*j*k == a {
					return true
				}
			}
		}
	}
	return false
}

func main() {
	fmt.Println(IsMultiplyPrime(30)) // true (2 * 3 * 5)
}
