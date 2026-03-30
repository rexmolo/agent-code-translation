package main

import (
	"strings"
)

func WordsString(s string) []string {
	if s == "" {
		return nil
	}

	// Replace commas with spaces, then split by whitespace
	s = strings.Replace(s, ",", " ", -1)
	return strings.Fields(s)
}

func main() {
	// Example usage
	result := WordsString("Hi, my name is John")
	println(result) // For testing; in real use, format output as needed
}
