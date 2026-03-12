package main

import (
	"strings"
)

// EncodeShift returns an encoded string by shifting every character by 5 in the alphabet.
func EncodeShift(s string) string {
	var builder strings.Builder
	builder.Grow(len(s))
	for _, ch := range s {
		encoded := ((int(ch) - int('a') + 5) % 26) + int('a')
		builder.WriteByte(byte(encoded))
	}
	return builder.String()
}

// DecodeShift takes as input string encoded with encode_shift function. Returns decoded string.
func DecodeShift(s string) string {
	var builder strings.Builder
	builder.Grow(len(s))
	for _, ch := range s {
		decoded := ((int(ch) - int('a') - 5 + 26) % 26) + int('a')
		builder.WriteByte(byte(decoded))
	}
	return builder.String()
}
