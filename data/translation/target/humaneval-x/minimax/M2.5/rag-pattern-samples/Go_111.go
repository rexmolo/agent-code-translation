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
	
	// Split by space
	list1 := strings.Split(test, " ")
	
	// Count occurrences of each letter
	counts := make(map[rune]int)
	for _, item := range list1 {
		if item == "" {
			continue
		}
		// Convert first character to rune
		r := rune(item[0])
		counts[r]++
	}
	
	// Find maximum count
	t := 0
	for _, count := range counts {
		if count > t {
			t = count
		}
	}
	
	// Build result with letters having maximum count
	if t > 0 {
		for letter, count := range counts {
			if count == t {
				result[letter] = count
			}
		}
	}
	
	return result
}

func main() {
	// Test cases
	fmt.Println(Histogram("a b c"))      // map[a:1 b:1 c:1]
	fmt.Println(Histogram("a b b a"))    // map[a:2 b:2]
	fmt.Println(Histogram("a b c a b"))  // map[a:2 b:2]
	fmt.Println(Histogram("b b b b a"))  // map[b:4]
	fmt.Println(Histogram(""))           // map[]
}