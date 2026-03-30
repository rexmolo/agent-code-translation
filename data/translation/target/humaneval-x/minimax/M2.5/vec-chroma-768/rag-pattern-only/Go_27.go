package main

import (
	"strings"
	"unicode"
)

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

func main() {
	// Example usage
	println(FlipCase("Hello")) // Output: hELLO
	println(FlipCase("hELLO")) // Output: Hello
	println(FlipCase("123ABCabc")) // Output: 123abcABC
}
