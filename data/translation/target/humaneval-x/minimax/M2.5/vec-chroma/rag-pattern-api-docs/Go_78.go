package main

import (
	"strings"
)

func HexKey(num string) int {
	primes := "2357BD"

	total := 0
	for _, ch := range num {
		if strings.ContainsRune(primes, ch) {
			total++
		}
	}
	return total
}

func main() {
	// Test cases
	println(HexKey("AB"))          // 1
	println(HexKey("1077E"))      // 2
	println(HexKey("ABED1A33"))   // 4
	println(HexKey("123456789ABCDEF0")) // 6
	println(HexKey("2020"))       // 2
}
