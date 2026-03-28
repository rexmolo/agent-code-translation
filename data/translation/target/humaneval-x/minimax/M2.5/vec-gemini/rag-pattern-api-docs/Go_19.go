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

	words := strings.Split(numbers, " ")

	// Filter out empty strings (handles multiple consecutive spaces like Python's [x for x in ... if x])
	var filtered []string
	for _, w := range words {
		if w != "" {
			filtered = append(filtered, w)
		}
	}

	// Sort using custom comparator based on numeric value
	sort.Slice(filtered, func(i, j int) bool {
		return valueMap[filtered[i]] < valueMap[filtered[j]]
	})

	return strings.Join(filtered, " ")
}