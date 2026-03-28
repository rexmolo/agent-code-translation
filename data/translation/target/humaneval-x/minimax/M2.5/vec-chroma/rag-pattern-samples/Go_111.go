package main

import (
	"fmt"
	"strings"
)

func Histogram(test string) map[rune]int {
	// Handle empty string
	if test == "" {
		return make(map[rune]int)
	}

	// Split by space
	list1 := strings.Split(test, " ")

	// Find the maximum count
	t := 0
	for _, i := range list1 {
		if i == "" {
			continue
		}
		count := 0
		for _, j := range list1 {
			if j == i {
				count++
			}
		}
		if count > t {
			t = count
		}
	}

	// Build the result map with letters that have the maximum count
	result := make(map[rune]int)
	if t > 0 {
		seen := make(map[string]bool)
		for _, i := range list1 {
			if i == "" || seen[i] {
				continue
			}
			count := 0
			for _, j := range list1 {
				if j == i {
					count++
				}
			}
			if count == t {
				result[[]rune(i)[0]] = count
				seen[i] = true
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