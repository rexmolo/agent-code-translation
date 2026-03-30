package main

import "fmt"

func HowManyTimes(str string, substring string) int {
	// Handle edge case: empty string or substring longer than string
	if len(str) == 0 || len(substring) > len(str) {
		return 0
	}
	
	// Handle edge case: empty substring
	if len(substring) == 0 {
		return 0
	}
	
	times := 0
	for i := 0; i <= len(str)-len(substring); i++ {
		if str[i:i+len(substring)] == substring {
			times++
		}
	}
	return times
}

func main() {
	// Example usage and tests
	fmt.Println(HowManyTimes("", "a"))     // 0
	fmt.Println(HowManyTimes("aaa", "a")) // 3
	fmt.Println(HowManyTimes("aaaa", "aa")) // 3
}
