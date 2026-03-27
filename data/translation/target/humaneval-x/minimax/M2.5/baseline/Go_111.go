package main

import (
	"fmt"
)

func Histogram(test string) map[rune]int {
	result := make(map[rune]int)

	if test == "" {
		return result
	}

	// Count occurrences of each letter
	counts := make(map[rune]int)
	maxCount := 0

	for _, r := range test {
		if r == ' ' {
			continue
		}
		counts[r]++
		if counts[r] > maxCount {
			maxCount = counts[r]
		}
	}

	if maxCount == 0 {
		return result
	}

	// Add letters with max count to result
	for letter, count := range counts {
		if count == maxCount {
			result[letter] = count
		}
	}

	return result
}

func main() {
	// Test cases
	fmt.Println(Histogram("a b c"))      // map[a:1 b:1 c:1]
	fmt.Println(Histogram("a b b a"))   // map[a:2 b:2]
	fmt.Println(Histogram("a b c a b")) // map[a:2 b:2]
	fmt.Println(Histogram("b b b b a")) // map[b:4]
	fmt.Println(Histogram(""))          // map[]
}