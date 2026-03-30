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

	// Filter out empty strings
	var filtered []string
	for _, s := range list1 {
		if s != "" {
			filtered = append(filtered, s)
		}
	}

	if len(filtered) == 0 {
		return result
	}

	// Find maximum count
	t := 0
	for _, s := range filtered {
		count := 0
		for _, item := range filtered {
			if item == s {
				count++
			}
		}
		if count > t {
			t = count
		}
	}

	// Build dictionary with letters that have max count
	if t > 0 {
		for _, s := range filtered {
			count := 0
			for _, item := range filtered {
				if item == s {
					count++
				}
			}
			if count == t {
				r := rune(s[0])
				result[r] = count
			}
		}
	}

	return result
}

func main() {
	// Test cases
	fmt.Println(Histogram("a b c"))
	fmt.Println(Histogram("a b b a"))
	fmt.Println(Histogram("a b c a b"))
	fmt.Println(Histogram("b b b b a"))
	fmt.Println(Histogram(""))
}