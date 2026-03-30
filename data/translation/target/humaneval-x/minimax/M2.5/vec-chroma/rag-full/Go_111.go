package main

import (
	"fmt"
	"strings"
)

func Histogram(test string) map[rune]int {
	result := make(map[rune]int)
	
	if test == "" {
		return result
	}
	
	parts := strings.Split(test, " ")
	
	// First pass: count each character and track maximum
	maxCount := 0
	for _, part := range parts {
		if part == "" {
			continue
		}
		// Handle space-separated single characters
		r := rune(part[0])
		result[r]++
		if result[r] > maxCount {
			maxCount = result[r]
		}
	}
	
	// Second pass: collect only characters with max count
	finalResult := make(map[rune]int)
	for r, count := range result {
		if count == maxCount && maxCount > 0 {
			finalResult[r] = count
		}
	}
	
	return finalResult
}

func main() {
	var input string
	fmt.Scanln(&input)
	
	for k, v := range Histogram(input) {
		fmt.Printf("%c: %d\n", k, v)
	}
}