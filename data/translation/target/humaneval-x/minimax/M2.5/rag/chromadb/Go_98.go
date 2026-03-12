package main

import "fmt"

func CountUpper(s string) int {
	count := 0
	for i := 0; i < len(s); i += 2 {
		if s[i] == 'A' || s[i] == 'E' || s[i] == 'I' || s[i] == 'O' || s[i] == 'U' {
			count++
		}
	}
	return count
}

func main() {
	// Test cases
	fmt.Println(CountUpper("aBCdEf"))  // Expected: 1
	fmt.Println(CountUpper("abcdefg")) // Expected: 0
	fmt.Println(CountUpper("dBBE"))    // Expected: 0
}