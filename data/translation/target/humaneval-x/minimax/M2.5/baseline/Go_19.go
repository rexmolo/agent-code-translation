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
	parts := strings.Fields(numbers)

	// Sort based on the value map
	sort.Slice(parts, func(i, j int) bool {
		return valueMap[parts[i]] < valueMap[parts[j]]
	})

	return strings.Join(parts, " ")
}
