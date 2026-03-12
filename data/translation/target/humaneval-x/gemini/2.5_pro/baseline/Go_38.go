package main

import (
	"strings"
)

// encodeCyclic is a helper function that encodes a string by performing a left
// cyclic shift on groups of three characters. Groups with fewer than 3
// characters are left unchanged.
// This function is not exported (private to the package) as it's a helper
// for the main DecodeCyclic function.
func encodeCyclic(s string) string {
	var builder strings.Builder
	// Pre-allocating the builder's capacity can improve performance by
	// reducing the number of allocations needed as the string is built.
	builder.Grow(len(s))

	for i := 0; i < len(s); i += 3 {
		end := i + 3
		if end > len(s) {
			end = len(s)
		}
		group := s[i:end]

		if len(group) == 3 {
			// In Python: group[1:] + group[0]
			// This performs the left cyclic shift, e.g., "abc" -> "bca".
			builder.WriteString(group[1:])
			builder.WriteByte(group[0])
		} else {
			// Groups with 1 or 2 characters are appended without modification.
			builder.WriteString(group)
		}
	}
	return builder.String()
}

// DecodeCyclic takes a string encoded with the cyclic shift logic and returns
// the original, decoded string.
func DecodeCyclic(s string) string {
	// The Python implementation decodes by applying the encoding function twice.
	// This is because for a 3-element cycle, two left shifts are equivalent to
	// one right shift, which is the inverse operation.
	// Let L be the encode operation (left shift).
	// s_encoded = L(s_original)
	// To decode, we need L_inverse(s_encoded).
	// L(L(L(s_original))) = s_original, so L_inverse = L(L).
	// Therefore, decode(s_encoded) = L(L(s_encoded)) which is L(L(L(s_original))) = s_original.
	return encodeCyclic(encodeCyclic(s))
}
