package main

import "fmt"

func encode_cyclic(s string) string {
	// split string to groups. Each of length 3.
	var groups []string
	for i := 0; i < len(s); i += 3 {
		end := i + 3
		if end > len(s) {
			end = len(s)
		}
		groups = append(groups, s[i:end])
	}

	// cycle elements in each group. Unless group has fewer elements than 3.
	result := ""
	for _, group := range groups {
		if len(group) == 3 {
			result += group[1:] + group[0:1]
		} else {
			result += group
		}
	}
	return result
}

func DecodeCyclic(s string) string {
	return encode_cyclic(encode_cyclic(s))
}

func main() {
	// Test the function
	testCases := []string{"abc", "abcdef", "abcde", "a", "ab"}
	for _, tc := range testCases {
		encoded := encode_cyclic(tc)
		decoded := DecodeCyclic(encoded)
		fmt.Printf("Original: %s, Encoded: %s, Decoded: %s\n", tc, encoded, decoded)
	}
}
