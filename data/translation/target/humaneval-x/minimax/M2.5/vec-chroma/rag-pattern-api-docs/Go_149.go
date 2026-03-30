package main

import (
	"fmt"
	"sort"
)

func SortedListSum(lst []string) []string {
	// Sort the list alphabetically first
	sort.Strings(lst)

	// Filter out strings with odd lengths
	var result []string
	for _, s := range lst {
		if len(s)%2 == 0 {
			result = append(result, s)
		}
	}

	// Sort by length first, then alphabetically for ties
	sort.Slice(result, func(i, j int) bool {
		if len(result[i]) == len(result[j]) {
			return result[i] < result[j]
		}
		return len(result[i]) < len(result[j])
	})

	return result
}

func main() {
	// Test cases
	fmt.Println(SortedListSum([]string{"aa", "a", "aaa"}))      // [aa]
	fmt.Println(SortedListSum([]string{"ab", "a", "aaa", "cd"})) // [ab cd]
}
