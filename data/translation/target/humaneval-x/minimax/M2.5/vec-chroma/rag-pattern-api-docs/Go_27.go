package main

import (
	"fmt"
	"strings"
	"unicode"
)

func FlipCase(str string) string {
	result := make([]rune, len(str))
	for i, r := range str {
		if unicode.IsUpper(r) {
			result[i] = unicode.ToLower(r)
		} else if unicode.IsLower(r) {
			result[i] = unicode.ToUpper(r)
		} else {
			result[i] = r
		}
	}
	return string(result)
}

func main() {
	fmt.Println(FlipCase("Hello"))
}