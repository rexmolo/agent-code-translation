package main

import (
	"fmt"
	"strings"
)

// WordsString splits a string by commas or spaces into a slice of words.
func WordsString(s string) []string {
	if s == "" {
		return []string{}
	}

	// Replace commas with spaces to treat them as a uniform delimiter.
	noCommas := strings.ReplaceAll(s, ",", " ")

	// strings.Fields splits the string around one or more consecutive white space
	// characters, which is equivalent to Python's s.split().
	return strings.Fields(noCommas)
}

func main() {
	fmt.Printf("%q\n", WordsString("Hi, my name is John"))
	fmt.Printf("%q\n", WordsString("One, two, three, four, five, six"))
	fmt.Printf("%q\n", WordsString(""))
	fmt.Printf("%q\n", WordsString("  leading,trailing spaces  "))
}
