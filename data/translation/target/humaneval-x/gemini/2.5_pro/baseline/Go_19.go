package main

import (
	"fmt"
	"sort"
	"strings"
)

// valueMap maps the string representation of numbers to their integer values.
var valueMap = map[string]int{
	"zero":  0,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

// SortNumbers takes a space-delimited string of number words ('zero' to 'nine')
// and returns a new string with the number words sorted numerically.
func SortNumbers(numbers string) string {
	words := strings.Split(numbers, " ")

	// Filter out empty strings that might result from multiple spaces,
	// mimicking Python's `[x for x in ... if x]` behavior.
	var numberWords []string
	for _, word := range words {
		if word != "" {
			numberWords = append(numberWords, word)
		}
	}

	// Sort the slice of number words using a custom comparison function.
	// This is the Go equivalent of Python's `sorted(..., key=...)`.
	// The comparison is based on the integer values from valueMap.
	sort.Slice(numberWords, func(i, j int) bool {
		return valueMap[numberWords[i]] < valueMap[numberWords[j]]
	})

	return strings.Join(numberWords, " ")
}

// main function to demonstrate SortNumbers, matching the Python docstring example.
func main() {
	input := "three one five"
	// The expected output is "one three five"
	fmt.Printf("'%s' -> '%s'\n", input, SortNumbers(input))

	input2 := "nine seven two eight zero"
	// The expected output is "zero two seven eight nine"
	fmt.Printf("'%s' -> '%s'\n", input2, SortNumbers(input2))

	input3 := "four  one   six"
	// The expected output is "one four six"
	fmt.Printf("'%s' -> '%s'\n", input3, SortNumbers(input3))
}