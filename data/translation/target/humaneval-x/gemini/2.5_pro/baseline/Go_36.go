package main

import (
	"fmt"
	"strconv"
	"strings"
)

// FizzBuzz returns the number of times the digit 7 appears in integers less than n
// which are divisible by 11 or 13.
func FizzBuzz(n int) int {
	var builder strings.Builder
	for i := 0; i < n; i++ {
		if i%11 == 0 || i%13 == 0 {
			builder.WriteString(strconv.Itoa(i))
		}
	}
	s := builder.String()
	return strings.Count(s, "7")
}

func main() {
	// Replicating the doctests from the Python example
	fmt.Println(FizzBuzz(50))  // Expected: 0
	fmt.Println(FizzBuzz(78))  // Expected: 2
	fmt.Println(FizzBuzz(79))  // Expected: 3
}
