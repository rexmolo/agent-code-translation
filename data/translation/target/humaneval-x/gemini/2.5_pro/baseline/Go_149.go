package main

import "sort"

// SortedListSum accepts a list of strings, filters out those with odd lengths,
// and returns the remaining strings sorted first by length (ascending) and then
// alphabetically for strings of the same length.
func SortedListSum(lst []string) []string {
	// Filter the list to keep only strings with even length.
	var evenLengthStrings []string
	for _, s := range lst {
		if len(s)%2 == 0 {
			evenLengthStrings = append(evenLengthStrings, s)
		}
	}

	// Sort the filtered list. The primary sort key is the length of the string (ascending),
	// and the secondary sort key is the string itself (alphabetical).
	// This is achieved using sort.Slice with a custom comparison function.
	sort.Slice(evenLengthStrings, func(i, j int) bool {
		s1 := evenLengthStrings[i]
		s2 := evenLengthStrings[j]

		// If lengths are different, sort by length.
		if len(s1) != len(s2) {
			return len(s1) < len(s2)
		}

		// If lengths are the same, sort alphabetically.
		return s1 < s2
	})

	return evenLengthStrings
}
