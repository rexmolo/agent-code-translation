package main

import "fmt"

func CountUpper(s string) int {
	vowels := map[rune]bool{
		'A': true,
		'E': true,
		'I': true,
		'O': true,
		'U': true,
	}
	
	count := 0
	for i := 0; i < len(s); i += 2 {
		if vowels[rune(s[i])] {
			count++
		}
	}
	return count
}

func main() {
	// Test cases
	fmt.Println(CountUpper("aBCdEf")) // Expected: 1
	fmt.Println(CountUpper("abcdefg")) // Expected: 0
	fmt.Println(CountUpper("dBBE")) // Expected: 0
}
