package main

import (
	"fmt"
	"strconv"
)

func Digits(n int) int {
	product := 1
	oddCount := 0

	// Convert n to string and iterate through each digit
	for _, d := range strconv.Itoa(n) {
		digit := int(d - '0') // Convert rune to integer digit
		if digit%2 == 1 {
			product *= digit
			oddCount++
		}
	}

	if oddCount == 0 {
		return 0
	}
	return product
}

func main() {
	// Test cases
	fmt.Println(Digits(1))   // == 1
	fmt.Println(Digits(4))   // == 0
	fmt.Println(Digits(235)) // == 15
}