package main

import (
	"fmt"
	"strings"
)

func HexKey(num string) int {
	primes := "2357BD"
	total := 0
	for i := 0; i < len(num); i++ {
		if strings.Contains(primes, string(num[i])) {
			total++
		}
	}
	return total
}

func main() {
	// Test cases from the problem statement
	fmt.Println(HexKey("AB"))           // Expected: 1
	fmt.Println(HexKey("1077E"))       // Expected: 2
	fmt.Println(HexKey("ABED1A33"))    // Expected: 4
	fmt.Println(HexKey("123456789ABCDEF0")) // Expected: 6
	fmt.Println(HexKey("2020"))        // Expected: 2
}
