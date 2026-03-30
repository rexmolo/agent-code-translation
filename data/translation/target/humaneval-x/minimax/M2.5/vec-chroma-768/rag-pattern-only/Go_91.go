package main

import (
	"fmt"
	"regexp"
	"strings"
)

func IsBored(S string) int {
	re := regexp.MustCompile(`[.?!]\s*`)
	sentences := re.Split(S, -1)

	count := 0
	for _, sentence := range sentences {
		if strings.HasPrefix(sentence, "I ") {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(IsBored("Hello world"))
	fmt.Println(IsBored("The sky is blue. The sun is shining. I love this weather"))
}