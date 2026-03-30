package main

import "fmt"

func HowManyTimes(str string, substring string) int {
	// Count how many times substring appears in str, including overlapping cases
	times := 0

	strLen := len(str)
	subLen := len(substring)

	// Edge case: if substring is longer than string, return 0
	if subLen > strLen {
		return 0
	}

	// Check each possible starting position
	for i := 0; i <= strLen-subLen; i++ {
		if str[i:i+subLen] == substring {
			times++
		}
	}

	return times
}

func main() {
	// Test cases
	fmt.Println(HowManyTimes("", "a"))   // Expected: 0
	fmt.Println(HowManyTimes("aaa", "a")) // Expected: 3
	fmt.Println(HowManyTimes("aaaa", "aa")) // Expected: 3
}
