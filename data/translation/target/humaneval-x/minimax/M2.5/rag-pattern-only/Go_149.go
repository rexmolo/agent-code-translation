package main

import (
	"sort"
)

func SortedListSum(lst []string) []string {
	// Sort alphabetically first - this determines order for equal-length strings
	sort.Strings(lst)
	
	// Filter: keep only strings with even length
	var newLst []string
	for _, s := range lst {
		if len(s)%2 == 0 {
			newLst = append(newLst, s)
		}
	}
	
	// Sort by length (ascending); for equal lengths, alphabetical order is preserved
	// because sort.Slice is stable in Go
	sort.Slice(newLst, func(i, j int) bool {
		if len(newLst[i]) == len(newLst[j]) {
			return newLst[i] < newLst[j]
		}
		return len(newLst[i]) < len(newLst[j])
	})
	
	return newLst
}