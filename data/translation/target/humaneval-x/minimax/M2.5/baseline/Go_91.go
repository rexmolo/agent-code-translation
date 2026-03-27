package main

import (
	"regexp"
)

func IsBored(S string) int {
	// Split by '.', '?', or '!' followed by optional whitespace
	re := regexp.MustCompile(`[.?!]\s*`)
	sentences := re.Split(S, -1)

	count := 0
	for _, sentence := range sentences {
		if len(sentence) >= 2 && sentence[0:2] == "I " {
			count++
		}
	}
	return count
}
