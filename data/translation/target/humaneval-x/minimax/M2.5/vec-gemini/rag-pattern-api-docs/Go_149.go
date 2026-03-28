package main

import (
	"fmt"
	"sort"
)

func SortedListSum(lst []string) []string {
	// Make a copy to avoid modifying the input
	result := make([]string, len(lst))
	copy(result, lst)

	// Sort alphabetically (in place)
	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})

	// Filter strings with even lengths
	var filtered []string
	for _, s := range result {
		if len(s)%2 == 0 {
			filtered = append(filtered, s)
		}
	}

	// Sort by length, then alphabetically for ties
	sort.Slice(filtered, func(i, j int) bool {
		if len(filtered[i]) == len(filtered[j]) {
			return filtered[i] < filtered[j]
		}
		return len(filtered[i]) < len(filtered[j])
	})

	return filtered
}

func main() {
	// Test examples
	fmt.Println(SortedListSum([]string{"aa", "a", "aaa"}))      // Output: [aa]
	fmt.Println(SortedListSum([]string{"ab", "a", "aaa", "cd"})) // Output: [ab cd]
}