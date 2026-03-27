package main

import "fmt"

func HowManyTimes(str string, substring string) int {
	times := 0
	subLen := len(substring)
	strLen := len(str)

	// Edge case: empty substring
	if subLen == 0 {
		return 0
	}

	for i := 0; i <= strLen-subLen; i++ {
		if str[i:i+subLen] == substring {
			times++
		}
	}

	return times
}

func main() {
	// Test cases
	fmt.Println(HowManyTimes("", "a"))      // Expected: 0
	fmt.Println(HowManyTimes("aaa", "a"))   // Expected: 3
	fmt.Println(HowManyTimes("aaaa", "aa")) // Expected: 3
}