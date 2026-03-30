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

	// Split by space and filter empty strings (Python's 'if x' in comprehension)
	words := strings.Split(numbers, " ")
	var nonEmpty []string
	for _, w := range words {
		if w != "" {
			nonEmpty = append(nonEmpty, w)
		}
	}

	// Sort by numeric value using sort.Slice with anonymous function (lambda equivalent)
	sort.Slice(nonEmpty, func(i, j int) bool {
		return valueMap[nonEmpty[i]] < valueMap[nonEmpty[j]]
	})

	return strings.Join(nonEmpty, " ")
}

func main() {
	fmt.Println(SortNumbers("three one five"))
}
