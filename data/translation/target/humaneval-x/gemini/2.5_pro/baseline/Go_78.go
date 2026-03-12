package main

import (
	"fmt"
)

// HexKey receives a hexadecimal number as a string and counts the number of
// hexadecimal digits that are primes (2, 3, 5, 7, B=11, D=13).
func HexKey(num string) int {
	total := 0
	// Iterate over each character (rune) in the input string.
	for _, char := range num {
		// Use a switch statement to check if the character is one of the
		// prime hexadecimal digits.
		switch char {
		case '2', '3', '5', '7', 'B', 'D':
			total++
		}
	}
	return total
}

// main function to demonstrate the usage of HexKey with the provided examples.
func main() {
	fmt.Println(HexKey("AB"))              // Expected: 1
	fmt.Println(HexKey("1077E"))           // Expected: 2
	fmt.Println(HexKey("ABED1A33"))        // Expected: 4
	fmt.Println(HexKey("123456789ABCDEF0"))// Expected: 6
	fmt.Println(HexKey("2020"))            // Expected: 2
}
