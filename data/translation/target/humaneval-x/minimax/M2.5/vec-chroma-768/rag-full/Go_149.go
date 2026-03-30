package main

import (
	"fmt"
	"slices"
	"sort"
)

func SortedListSum(lst []string) []string {
	// Sort the input list in place (alphabetically)
	slices.Sort(lst)

	// Filter out strings with odd lengths
	var new_lst []string
	for _, s := range lst {
		if len(s)%2 == 0 {
			new_lst = append(new_lst, s)
		}
	}

	// Sort by length first, then alphabetically for same length
	sort.Slice(new_lst, func(i, j int) bool {
		if len(new_lst[i]) == len(new_lst[j]) {
			return new_lst[i] < new_lst[j]
		}
		return len(new_lst[i]) < len(new_lst[j])
	})

	return new_lst
}

func main() {
	// Test cases
	test1 := []string{"aa", "a", "aaa"}
	fmt.Println(SortedListSum(test1)) // Expected: [aa]

	test2 := []string{"ab", "a", "aaa", "cd"}
	fmt.Println(SortedListSum(test2)) // Expected: [ab cd]
}
