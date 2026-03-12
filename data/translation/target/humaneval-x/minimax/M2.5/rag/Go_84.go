package main

import (
	"fmt"
	"strconv"
)

func Solve(N int) string {
	// Convert N to string
	str := strconv.Itoa(N)

	// Sum the digits
	sum := 0
	for _, c := range str {
		sum += int(c - '0')
	}

	// Convert sum to binary string
	return fmt.Sprintf("%b", sum)
}

func main() {
	// Read from stdin (single integer N)
	var N int
	fmt.Scan(&N)

	// Output the result
	fmt.Println(Solve(N))
}
