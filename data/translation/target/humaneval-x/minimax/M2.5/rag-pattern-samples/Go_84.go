package main

import (
	"fmt"
	"strconv"
)

func Solve(N int) string {
	// Convert N to string to iterate over each digit
	strN := strconv.Itoa(N)

	// Sum all the digits
	sum := 0
	for _, c := range strN {
		sum += int(c - '0')
	}

	// Convert the sum to binary string
	// strconv.FormatInt with base 2 handles the conversion
	return strconv.FormatInt(int64(sum), 2)
}

func main() {
	var N int
	fmt.Scan(&N)
	fmt.Println(Solve(N))
}
