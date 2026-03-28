package main

import (
	"fmt"
	"sort"
)

func SortedListSum(lst []string) []string {
	// Sort alphabetically first
	sort.Strings(lst)

	// Filter to keep only strings with even length
	var new_lst []string
	for _, s := range lst {
		if len(s)%2 == 0 {
			new_lst = append(new_lst, s)
		}
	}

	// Sort by length (primary) and alphabetically (secondary for stable sort)
	sort.SliceStable(new_lst, func(i, j int) bool {
		if len(new_lst[i]) == len(new_lst[j]) {
			return new_lst[i] < new_lst[j]
		}
		return len(new_lst[i]) < len(new_lst[j])
	})

	return new_lst
}

func main() {
	// Test cases
	fmt.Println(SortedListSum([]string{"aa", "a", "aaa"}))           // [aa]
	fmt.Println(SortedListSum([]string{"ab", "a", "aaa", "cd"})) // [ab cd]
}
