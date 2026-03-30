package main

import (
	"fmt"
	"regexp"
)

func IsBored(S string) int {
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

func main() {
	// Test cases
	fmt.Println(IsBored("Hello world"))
	fmt.Println(IsBored("The sky is blue. The sun is shining. I love this weather"))
}
