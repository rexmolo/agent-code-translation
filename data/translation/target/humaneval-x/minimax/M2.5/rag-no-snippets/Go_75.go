package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func IsMultiplyPrime(a int) bool {
	// Helper function to check if a number is prime
	isPrime := func(n int) bool {
		if n < 2 {
			return false
		}
		for j := 2; j < n; j++ {
			if n%j == 0 {
				return false
			}
		}
		return true
	}

	// Triple nested loop to find if a can be expressed as product of 3 primes
	for i := 2; i <= 100; i++ {
		if !isPrime(i) {
			continue
		}
		for j := 2; j <= 100; j++ {
			if !isPrime(j) {
				continue
			}
			for k := 2; k <= 100; k++ {
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
	// Read input from stdin
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	
	a, _ := strconv.Atoi(line)
	
	result := IsMultiplyPrime(a)
	fmt.Println(result)
}