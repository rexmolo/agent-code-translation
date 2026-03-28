package main

import (
	"fmt"
	"strconv"
)

func Solve(N int) string {
	// Convert integer to string
	s := strconv.Itoa(N)
	
	// Sum all digits
	sum := 0
	for _, c := range s {
		sum += int(c - '0')
	}
	
	// Convert sum to binary string
	result := fmt.Sprintf("%b", sum)
	
	return result
}

// For local testing:
// func main() {
// 	fmt.Println(Solve(1000)) // Output: 1
// 	fmt.Println(Solve(150))  // Output: 110
// 	fmt.Println(Solve(147))  // Output: 1100
// }