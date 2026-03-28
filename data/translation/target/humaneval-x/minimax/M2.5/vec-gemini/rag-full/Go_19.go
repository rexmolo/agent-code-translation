package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func SortNumbers(numbers string) string {
	valueMap := map[string]int{
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

	// Split by space and filter out empty strings
	parts := strings.Split(numbers, " ")
	var words []string
	for _, x := range parts {
		if x != "" {
			words = append(words, x)
		}
	}

	// Sort by numeric value using sort.Slice with custom comparator
	sort.Slice(words, func(i, j int) int {
		return valueMap[words[i]] - valueMap[words[j]]
	})

	return strings.Join(words, " ")
}

func main() {
	// Read from stdin
	scanner := bufio.NewScanner(os.Stdin)
	var input string
	for scanner.Scan() {
		input += scanner.Text() + " "
	}
	input = strings.TrimSpace(input)

	result := SortNumbers(input)
	fmt.Println(result)
}