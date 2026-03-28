package main

import "fmt"

func EncodeCyclic(s string) string {
	// Split into groups of 3 characters
	var result []byte
	for i := 0; i < len(s); i += 3 {
		end := i + 3
		if end > len(s) {
			end = len(s)
		}
		group := s[i:end]

		if len(group) == 3 {
			// Cycle: group[1:] + group[0]
			result = append(result, group[1], group[2], group[0])
		} else {
			// Keep group as is if fewer than 3 elements
			result = append(result, group...)
		}
	}
	return string(result)
}

func DecodeCyclic(s string) string {
	// Decoding is done by applying encode twice
	return EncodeCyclic(EncodeCyclic(s))
}

func main() {
	// Test the functions
	testCases := []string{"abc", "abcdef", "abcdefg", "ab", "a", ""}
	for _, tc := range testCases {
		encoded := EncodeCyclic(tc)
		decoded := DecodeCyclic(encoded)
		fmt.Printf("Original: %q, Encoded: %q, Decoded: %q\n", tc, encoded, decoded)
	}
}
