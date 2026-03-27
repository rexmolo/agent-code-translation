package main

import (
	"strings"
)

func encodeCyclic(s string) string {
	// Split string into groups of 3
	var groups []string
	for i := 0; i < len(s); i += 3 {
		end := i + 3
		if end > len(s) {
			end = len(s)
		}
		groups = append(groups, s[i:end])
	}

	// Cycle elements in each group (only if group has exactly 3 elements)
	for i, group := range groups {
		if len(group) == 3 {
			groups[i] = group[1:] + group[0:1]
		}
	}

	// Join groups
	var builder strings.Builder
	for _, g := range groups {
		builder.WriteString(g)
	}
	return builder.String()
}

func DecodeCyclic(s string) string {
	return encodeCyclic(encodeCyclic(s))
}
