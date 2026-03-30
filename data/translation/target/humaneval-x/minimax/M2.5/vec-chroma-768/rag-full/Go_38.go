package main

import (
	"slices"
	"strings"
)

func encodeCyclic(s string) string {
	// split string to groups. Each of length 3.
	numGroups := (len(s) + 2) / 3
	groups := make([]string, 0, numGroups)
	for i := 0; i < numGroups; i++ {
		start := 3 * i
		end := slices.Min([]int{3*i + 3, len(s)})
		groups = append(groups, s[start:end])
	}

	// cycle elements in each group. Unless group has fewer elements than 3.
	for i, group := range groups {
		if len(group) == 3 {
			groups[i] = group[1:] + string(group[0])
		}
	}
	return strings.Join(groups, "")
}

func DecodeCyclic(s string) string {
	return encodeCyclic(encodeCyclic(s))
}
