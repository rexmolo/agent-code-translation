package main

import (
	"fmt"
	"unicode"
)

func FlipCase(str string) string {
	runes := []rune(str)
	for i, r := range runes {
		if unicode.IsUpper(r) {
			runes[i] = unicode.ToLower(r)
		} else if unicode.IsLower(r) {
			runes[i] = unicode.ToUpper(r)
		}
	}
	return string(runes)
}

func main() {
	fmt.Println(FlipCase("Hello"))
}
