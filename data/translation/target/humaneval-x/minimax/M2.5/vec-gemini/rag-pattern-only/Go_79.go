package main

import (
	"fmt"
)

func DecimalToBinary(decimal int) string {
	// Handle special case of 0
	if decimal == 0 {
		return "db0db"
	}

	// Convert decimal to binary string
	binary := ""
	for decimal > 0 {
		remainder := decimal % 2
		binary = fmt.Sprintf("%d%s", remainder, binary)
		decimal = decimal / 2
	}

	return "db" + binary + "db"
}

func main() {
	// Test cases
	fmt.Println(DecimalToBinary(15))  // Output: db1111db
	fmt.Println(DecimalToBinary(32))  // Output: db100000db
}