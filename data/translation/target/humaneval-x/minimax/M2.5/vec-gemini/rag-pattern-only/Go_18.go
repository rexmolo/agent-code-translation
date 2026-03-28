package main

import "fmt"

func HowManyTimes(str string, substring string) int {
	times := 0
	subLen := len(substring)
	strLen := len(str)

	for i := 0; i <= strLen-subLen; i++ {
		if str[i:i+subLen] == substring {
			times++
		}
	}

	return times
}

func main() {
	// Test cases
	fmt.Println(HowManyTimes("", "a"))      // 0
	fmt.Println(HowManyTimes("aaa", "a"))   // 3
	fmt.Println(HowManyTimes("aaaa", "aa")) // 3
}
