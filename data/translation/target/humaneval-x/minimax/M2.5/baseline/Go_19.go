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

	parts := strings.Split(numbers, " ")

	// Filter out empty strings (equivalent to [x for x in numbers.split(' ') if x])
	var filtered []string
	for _, s := range parts {
		if s != "" {
			filtered = append(filtered, s)
		}
	}

	// Sort by numeric value using custom comparator
	sort.Slice(filtered, func(i, j int) bool {
		return valueMap[filtered[i]] < valueMap[filtered[j]]
	})

	return strings.Join(filtered, " ")
}
