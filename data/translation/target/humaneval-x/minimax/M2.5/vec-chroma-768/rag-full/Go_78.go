package main

import "fmt"

func HexKey(num string) int {
	primes := map[byte]bool{
		'2': true,
		'3': true,
		'5': true,
		'7': true,
		'B': true,
		'D': true,
	}
	total := 0
	for i := 0; i < len(num); i++ {
		if primes[num[i]] {
			total++
		}
	}
	return total
}

func main() {
	// Test cases
	fmt.Println(HexKey("AB"))       // Expected: 1
	fmt.Println(HexKey("1077E"))   // Expected: 2
	fmt.Println(HexKey("ABED1A33")) // Expected: 4
	fmt.Println(HexKey("123456789ABCDEF0")) // Expected: 6
	fmt.Println(HexKey("2020"))    // Expected: 2
}
