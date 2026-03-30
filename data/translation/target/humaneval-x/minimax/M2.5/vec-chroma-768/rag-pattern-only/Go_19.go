package main

import (
	"fmt"
	"slices"
	"strings"
)

func SortNumbers(numbers string) string {
	valueMap := map[string]int{
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		""five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	words := strings.Fields(numbers)

	slices.SortFunc(words, func(a, b string) int {
		return valueMap[a] - valueMap[b]
	})

	return strings.Join(words, " ")
}

func main() {
	examples := []string{
		"three one five",
		"zero four two eight",
		"nine five eight one",
	}
	for _, ex := range examples {
		fmt.Printf("Input: %q -> Output: %q\n", ex, SortNumbers(ex))
	}
}
