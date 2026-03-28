package main

import (
	"cmp"
	"slices"
)

func FindMax(words []string) string {
	if len(words) == 0 {
		return ""
	}

	slices.SortFunc(words, func(a, b string) int {
		uniqueA := countUniqueChars(a)
		uniqueB := countUniqueChars(b)

		// Primary: descending order of unique character count
		// (larger count comes first)
		if uniqueA != uniqueB {
			return cmp.Compare(uniqueB, uniqueA)
		}
		// Secondary: ascending lexicographical order
		return cmp.Compare(a, b)
	})

	return words[0]
}

func countUniqueChars(s string) int {
	unique := make(map[rune]bool)
	for _, ch := range s {
		unique[ch] = true
	}
	return len(unique)
}
