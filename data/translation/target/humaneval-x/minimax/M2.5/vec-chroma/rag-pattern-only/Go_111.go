package main

import (
	"fmt"
	"strings"
)

func Histogram(test string) map[rune]int {
	// Split by space
	list1 := strings.Split(test, " ")
	
	// Count occurrences of each letter
	counts := make(map[string]int)
	for _, item := range list1 {
		if item != "" {
			counts[item]++
		}
	}
	
	// Find maximum count
	t := 0
	for _, count := range counts {
		if count > t {
			t = count
		}
	}
	
	// Build result with letters that have max count
	result := make(map[rune]int)
	if t > 0 {
		for item, count := range counts {
			if count == t {
				// Convert string key to rune
				for _, r := range item {
					result[r] = count
				}
			}
		}
	}
	
	return result
}

func main() {
	// Test cases
	fmt.Println(Histogram("a b c"))       // map[a:1 b:1 c:1]
	fmt.Println(Histogram("a b b a"))     // map[a:2 b:2]
	fmt.Println(Histogram("a b c a b"))   // map[a:2 b:2]
	fmt.Println(Histogram("b b b b a"))   // map[b:4]
	fmt.Println(Histogram(""))            // map[]
}