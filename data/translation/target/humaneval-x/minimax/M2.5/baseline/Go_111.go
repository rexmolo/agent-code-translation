package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Histogram(test string) map[rune]int {
	result := make(map[rune]int)
	
	if test == "" {
		return result
	}
	
	// Split by space to get individual characters
	parts := strings.Split(test, " ")
	
	// Count frequency of each character
	counts := make(map[rune]int)
	for _, part := range parts {
		if part == "" {
			continue
		}
		for _, ch := range part {
			counts[ch]++
		}
	}
	
	// Find maximum count
	maxCount := 0
	for _, count := range counts {
		if count > maxCount {
			maxCount = count
		}
	}
	
	if maxCount == 0 {
		return result
	}
	
	// Add all characters with max count to result
	for ch, count := range counts {
		if count == maxCount {
			result[ch] = count
		}
	}
	
	return result
}

func main() {
	// Read input from stdin
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a space separated string: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	
	result := Histogram(input)
	fmt.Println(result)
}
