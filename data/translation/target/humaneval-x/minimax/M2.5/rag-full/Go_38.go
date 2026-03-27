package main

import (
	"strings"
)

func encodeCyclic(s string) string {
	// Calculate number of groups (ceiling division)
	numGroups := (len(s) + 2) / 3

	var groups []string
	for i := 0; i < numGroups; i++ {
		start := 3 * i
		end := start + 3
		if end > len(s) {
			end = len(s)
		}
		group := s[start:end]

		// Cycle elements if group has exactly 3 elements
		if len(group) == 3 {
			group = group[1:] + string(group[0])
		}
		groups = append(groups, group)
	}
	return strings.Join(groups, "")
}

func DecodeCyclic(s string) string {
	// Decode by applying encode twice (rotation is its own inverse for 3-char groups)
	return encodeCyclic(encodeCyclic(s))
}
