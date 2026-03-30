package main

import "strings"

func EncodeCyclic(s string) string {
	// split string to groups. Each of length 3.
	numGroups := (len(s) + 2) / 3

	var groups []string
	for i := 0; i < numGroups; i++ {
		start := 3 * i
		end := start + 3
		if end > len(s) {
			end = len(s)
		}
		group := s[start:end]
		groups = append(groups, group)
	}

	// cycle elements in each group. Unless group has fewer elements than 3.
	var result []string
	for _, group := range groups {
		if len(group) == 3 {
			// Move first char to end: group[1], group[2], group[0]
			result = append(result, group[1:]+string(group[0]))
		} else {
			result = append(result, group)
		}
	}

	return strings.Join(result, "")
}

func DecodeCyclic(s string) string {
	// Decoding is just encoding twice (cyclic operation is its own inverse)
	return EncodeCyclic(EncodeCyclic(s))
}
