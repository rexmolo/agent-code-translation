package main

import "fmt"

func HowManyTimes(str string, substring string) int {
	// Edge case: if substring is longer than str, return 0
	if len(substring) > len(str) {
		return 0
	}
	
	times := 0
	
	// Iterate through each possible starting position
	// In Python: range(len(string) - len(substring) + 1)
	for i := 0; i <= len(str)-len(substring); i++ {
		// Python string[i:i+len(substring)] translates to Go slice notation
		if str[i:i+len(substring)] == substring {
			times++
		}
	}
	
	return times
}

func main() {
	// Test cases
	fmt.Println(HowManyTimes("", "a"))       // Expected: 0
	fmt.Println(HowManyTimes("aaa", "a"))    // Expected: 3
	fmt.Println(HowManyTimes("aaaa", "aa")) // Expected: 3
}
