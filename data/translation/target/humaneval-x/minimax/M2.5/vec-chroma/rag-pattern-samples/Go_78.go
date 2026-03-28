package main

import "fmt"

func HexKey(num string) int {
	// Prime hexadecimal digits: 2, 3, 5, 7, B (decimal 11), D (decimal 13)
	primes := map[rune]bool{
		'2': true,
		'3': true,
		'5': true,
		'7': true,
		'B': true,
		'D': true,
	}

	total := 0
	for _, r := range num {
		if primes[r] {
			total++
		}
	}
	return total
}

func main() {
	// Test cases
	fmt.Println(HexKey("AB"))           // 1
	fmt.Println(HexKey("1077E"))        // 2
	fmt.Println(HexKey("ABED1A33"))     // 4
	fmt.Println(HexKey("123456789ABCDEF0")) // 6
	fmt.Println(HexKey("2020"))        // 2
}
