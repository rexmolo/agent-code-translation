package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// isPrime checks if a number is prime
func isPrime(k int) bool {
	if k < 2 {
		return false
	}
	for i := 2; i < k; i++ {
		if k%i == 0 {
			return false
		}
	}
	return true
}

// LargestPrimeFactor returns the largest prime factor of n.
// Assumes n > 1 and is not a prime.
func LargestPrimeFactor(n int) int {
	largest := 1
	for j := 2; j <= n; j++ {
		if n%j == 0 && isPrime(j) {
			if j > largest {
				largest = j
			}
		}
	}
	return largest
}

func main() {
	// Read from stdin
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	
	n, _ := strconv.Atoi(input)
	result := LargestPrimeFactor(n)
	
	// Write to stdout
	fmt.Println(result)
}
