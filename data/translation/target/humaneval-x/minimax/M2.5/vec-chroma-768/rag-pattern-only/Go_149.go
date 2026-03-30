package main

import (
	"fmt"
	"sort"
)

func SortedListSum(lst []string) []string {
	// Filter strings with even length
	var newLst []string
	for _, s := range lst {
		if len(s)%2 == 0 {
			newLst = append(newLst, s)
		}
	}

	// Sort by length first, then alphabetically for ties
	sort.Slice(newLst, func(i, j int) bool {
		if len(newLst[i]) == len(newLst[j]) {
			return newLst[i] < newLst[j]
		}
		return len(newLst[i]) < len(newLst[j])
	})

	return newLst
}

func main() {
	// Test cases
	fmt.Println(SortedListSum([]string{"aa", "a", "aaa"}))           // [aa]
	fmt.Println(SortedListSum([]string{"ab", "a", "aaa", "cd"}))   // [ab cd]
}