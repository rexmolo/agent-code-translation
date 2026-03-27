package main

import (
	"fmt"
	"sort"
)

func SortedListSum(lst []string) []string {
	// Sort alphabetically first
	sort.Strings(lst)

	// Filter to keep only strings with even length
	var newLst []string
	for _, s := range lst {
		if len(s)%2 == 0 {
			newLst = append(newLst, s)
		}
	}

	// Sort by length (ascending), then alphabetically for same length
	sort.Slice(newLst, func(i, j int) bool {
		if len(newLst[i]) == len(newLst[j]) {
			return newLst[i] < newLst[j]
		}
		return len(newLst[i]) < len(newLst[j])
	})

	return newLst
}

func main() {
	// Test examples
	result1 := SortedListSum([]string{"aa", "a", "aaa"})
	fmt.Printf("%v\n", result1) // [aa]

	result2 := SortedListSum([]string{"ab", "a", "aaa", "cd"})
	fmt.Printf("%v\n", result2) // [ab cd]
}
