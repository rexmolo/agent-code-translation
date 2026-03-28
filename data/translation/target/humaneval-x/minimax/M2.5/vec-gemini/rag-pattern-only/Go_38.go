package main

import (
	"strings"
)

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
	for i, group := range groups {
		if len(group) == 3 {
			groups[i] = group[1:] + group[0]
		}
	}

	return strings.Join(groups, "")
}

func DecodeCyclic(s string) string {
	return EncodeCyclic(EncodeCyclic(s))
}
