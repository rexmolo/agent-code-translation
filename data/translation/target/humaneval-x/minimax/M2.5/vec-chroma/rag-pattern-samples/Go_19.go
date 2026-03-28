package main

import (
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

	// Split by space and filter out empty strings
	parts := strings.Split(numbers, " ")
	var filtered []string
	for _, x := range parts {
		if x != "" {
			filtered = append(filtered, x)
		}
	}

	// Sort by numeric value using custom comparison
	sort.Slice(filtered, func(i, j int) bool {
		return valueMap[filtered[i]] < valueMap[filtered[j]]
	})

	return strings.Join(filtered, " ")
}
