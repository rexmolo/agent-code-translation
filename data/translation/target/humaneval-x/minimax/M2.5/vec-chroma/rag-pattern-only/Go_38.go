package main

import "fmt"

func EncodeCyclic(s string) string {
	// Calculate number of groups (ceil division)
	numGroups := (len(s) + 2) / 3

	var result []byte

	for i := 0; i < numGroups; i++ {
		start := 3 * i
		end := start + 3
		if end > len(s) {
			end = len(s)
		}

		group := s[start:end]

		if len(group) == 3 {
			// Rotate: group[1:] + group[0]
			result = append(result, group[1:]...)
			result = append(result, group[0])
		} else {
			result = append(result, group...)
		}
	}

	return string(result)
}

func DecodeCyclic(s string) string {
	return EncodeCyclic(EncodeCyclic(s))
}

func main() {
	// Test examples
	testCases := []string{"abc", "abcd", "abcdef", "abcde", ""}
	for _, tc := range testCases {
		encoded := EncodeCyclic(tc)
		decoded := DecodeCyclic(encoded)
		fmt.Printf("Original: %q, Encoded: %q, Decoded: %q\n", tc, encoded, decoded)
	}
}
