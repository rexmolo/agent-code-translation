package main

import (
	"strings"
	"unicode"
)

// FlipCase for a given string, flips lowercase characters to uppercase and uppercase to lowercase.
func FlipCase(str string) string {
	var builder strings.Builder
	builder.Grow(len(str))

	for _, r := range str {
		if unicode.IsUpper(r) {
			builder.WriteRune(unicode.ToLower(r))
		} else if unicode.IsLower(r) {
			builder.WriteRune(unicode.ToUpper(r))
		} else {
			builder.WriteRune(r)
		}
	}

	return builder.String()
}
