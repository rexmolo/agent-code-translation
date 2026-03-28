package main

import (
	"fmt"
	"sort"
)

func SortedListSum(lst []string) []string {
	// Sort alphabetically first
	sort.Strings(lst)

	// Filter out strings with odd lengths
	var newLst []string
	for _, s := range lst {
		if len(s)%2 == 0 {
			newLst = append(newLst, s)
		}
	}

	// Sort by length first, then alphabetically for same length
	sort.Slice(newLst, func(i, j int) bool {
		if len(newLst[i]) == len(newLst[j]) {
			return newLst[i] < newLst[j]
		}
		return len(newLst[i]) < len(newLst[j])
	})

	return newLst
}

func main() {
	// Example usage
	result1 := SortedListSum([]string{"aa", "a", "aaa"})
	fmt.Println(result1) // Output: [aa]

	result2 := SortedListSum([]string{"ab", "a", "aaa", "cd"})
	fmt.Println(result2) // Output: [ab cd]
}