package main

import (
	"fmt"
	"strconv"
)

func Solve(N int) string {
	// Convert N to string
	strN := fmt.Sprintf("%d", N)

	// Sum the digits
	sum := 0
	for _, c := range strN {
		sum += int(c - '0')
	}

	// Convert sum to binary string
	return strconv.FormatInt(int64(sum), 2)
}

// For testing purposes (not part of the solution but allows running)
func main() {
	// Example usage:
	fmt.Println(Solve(1000)) // Output: 1
	fmt.Println(Solve(150))  // Output: 110
	fmt.Println(Solve(147))  // Output: 1100
}
