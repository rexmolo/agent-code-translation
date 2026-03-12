package main

import (
	"fmt"
	"unicode/utf8"
)

func Strlen(str string) int {
	return utf8.RuneCountInString(str)
}

func main() {
	// Demonstration mirroring the Python docstring examples
	fmt.Println(Strlen(""))
	fmt.Println(Strlen("abc"))
}
