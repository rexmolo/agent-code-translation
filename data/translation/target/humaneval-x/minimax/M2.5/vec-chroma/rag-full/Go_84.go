package main

import (
	"fmt"
	"strconv"
)

func Solve(N int) string {
	// Convert N to string
	s := strconv.Itoa(N)
	
	// Sum all digits
	sum := 0
	for _, c := range s {
		digit := int(c - '0')
		sum += digit
	}
	
	// Convert to binary string
	return fmt.Sprintf("%b", sum)
}

func main() {
	var N int
	fmt.Scan(&N)
	fmt.Println(Solve(N))
}