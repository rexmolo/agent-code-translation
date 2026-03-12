package main

import (
	"regexp"
	"strings"
)

func IsBored(S string) int {
	// Split by '.', '?', or '!' followed by optional whitespace
	re := regexp.MustCompile(`[.?!]\s*`)
	sentences := re.Split(S, -1)

	count := 0
	for _, sentence := range sentences {
		// Check if sentence starts with "I " (I followed by space)
		if len(sentence) >= 2 && strings.HasPrefix(sentence, "I ") {
			count++
		}
	}
	return count
}