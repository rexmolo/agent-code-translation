package main

import (
	"strings"
)

func HowManyTimes(str string, substring string) int {
	// Edge cases: empty string or substring longer than string
	if len(str) == 0 || len(substring) == 0 || len(substring) > len(str) {
		return 0
	}

	times := 0
	subLen := len(substring)

	for i := 0; i <= len(str)-subLen; i++ {
		if str[i:i+subLen] == substring {
			times++
		}
	}

	return times
}