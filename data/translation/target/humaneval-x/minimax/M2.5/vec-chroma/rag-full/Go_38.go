package main

import (
	"strings"
)

// EncodeCyclic returns encoded string by cycling groups of three characters.
func EncodeCyclic(s string) string {
	// split string to groups. Each of length 3.
	groupCount := (len(s) + 2) / 3
	groups := make([]string, 0, groupCount)

	for i := 0; i < groupCount; i++ {
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
			// group[1:] + group[0] - move first char to end
			groups[i] = group[1:] + group[0:1]
		}
	}

	return strings.Join(groups, "")
}

// DecodeCyclic takes as input string encoded with encode_cyclic function.
// Returns decoded string.
func DecodeCyclic(s string) string {
	return EncodeCyclic(EncodeCyclic(s))
}

func main() {
	// Example usage - can be tested with stdin/stdout
	s := "Hello World"
	encoded := EncodeCyclic(s)
	decoded := DecodeCyclic(encoded)
	_ = encoded
	_ = decoded
}
