package main

import (
	"fmt"
	"strings"
)

func encodeCyclic(s string) string {
	// split string to groups. Each of length 3.
	numGroups := (len(s) + 2) / 3
	groups := make([]string, numGroups)
	for i := 0; i < numGroups; i++ {
		start := 3 * i
		end := start + 3
		if end > len(s) {
			end = len(s)
		}
		groups[i] = s[start:end]
	}
	// cycle elements in each group. Unless group has fewer elements than 3.
	for i, group := range groups {
		if len(group) == 3 {
			groups[i] = group[1:] + group[0:1]
		}
	}
	return strings.Join(groups, "")
}

func DecodeCyclic(s string) string {
	// takes as input string encoded with encode_cyclic function. Returns decoded string.
	return encodeCyclic(encodeCyclic(s))
}

func main() {
	// Test the functions
	testCases := []string{"abc", "abcdef", "a", "ab", "abcd", ""}
	for _, tc := range testCases {
		encoded := encodeCyclic(tc)
		decoded := DecodeCyclic(encoded)
		fmt.Printf("Original: %q, Encoded: %q, Decoded: %q\n", tc, encoded, decoded)
	}
}
