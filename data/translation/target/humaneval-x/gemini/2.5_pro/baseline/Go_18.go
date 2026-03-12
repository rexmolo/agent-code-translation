package main

import "fmt"

// HowManyTimes finds how many times a given substring can be found in the original string.
// It counts overlapping cases.
func HowManyTimes(str string, substring string) int {
	times := 0
	lenSubstring := len(substring)
	lenStr := len(str)

	// The loop's upper bound is `lenStr - lenSubstring`.
	// This ensures that the slice `str[i:i+lenSubstring]` will not go out of bounds.
	// It correctly handles all edge cases, including empty strings or a substring longer than the string.
	for i := 0; i <= lenStr-lenSubstring; i++ {
		if str[i:i+lenSubstring] == substring {
			times++
		}
	}

	return times
}

// main function to demonstrate the HowManyTimes function, mimicking the Python doctests.
func main() {
	fmt.Println(HowManyTimes("", "a"))
	fmt.Println(HowManyTimes("aaa", "a"))
	fmt.Println(HowManyTimes("aaaa", "aa"))
}
