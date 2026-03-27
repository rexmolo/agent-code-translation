package main

import (
	"fmt"
	"sort"
	"strings"
)

func SortNumbers(numbers string) string {
	valueMap := map[string]int{
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	// Split by space and filter empty strings
	parts := strings.Split(numbers, " ")
	var words []string
	for _, x := range parts {
		if x != "" {
			words = append(words, x)
		}
	}

	// Sort by numeric value using sort.Slice with custom comparator
	sort.Slice(words, func(i, j int) int {
		return valueMap[words[i]] - valueMap[words[j]]
	})

	return strings.Join(words, " ")
}

func main() {
	// Example usage
	result := SortNumbers("three one five")
	fmt.Println(result)
}