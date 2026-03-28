package main

import (
	"fmt"
	"unicode/utf8"
)

func Strlen(str string) int {
	return utf8.RuneCountInString(str)
}

func main() {
	// Test the Strlen function
	fmt.Println(Strlen(""))   // 0
	fmt.Println(Strlen("abc")) // 3
}