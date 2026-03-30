package main

import "fmt"

func HowManyTimes(str string, substring string) int {
	times := 0

	s := []rune(str)
	sub := []rune(substring)
	subLen := len(sub)

	for i := 0; i <= len(s)-subLen; i++ {
		if string(s[i:i+subLen]) == string(sub) {
			times++
		}
	}

	return times
}

func main() {
	// Test cases
	fmt.Println(HowManyTimes("", "a"))    // Expected: 0
	fmt.Println(HowManyTimes("aaa", "a")) // Expected: 3
	fmt.Println(HowManyTimes("aaaa", "aa")) // Expected: 3
}
