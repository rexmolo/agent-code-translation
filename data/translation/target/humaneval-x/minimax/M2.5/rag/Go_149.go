package main

import (
	"fmt"
	"sort"
)

func SortedListSum(lst []string) []string {
	// Sort alphabetically first
	sort.Strings(lst)

	// Filter to keep only even-length strings
	var new_lst []string
	for _, s := range lst {
		if len(s)%2 == 0 {
			new_lst = append(new_lst, s)
		}
	}

	// Sort by length (ascending), and for same length, alphabetically
	sort.Slice(new_lst, func(i, j int) bool {
		if len(new_lst[i]) == len(new_lst[j]) {
			return new_lst[i] < new_lst[j]
		}
		return len(new_lst[i]) < len(new_lst[j])
	})

	return new_lst
}

func main() {
	// Example usage:
	lst1 := []string{"aa", "a", "aaa"}
	fmt.Println(SortedListSum(lst1)) // Output: [aa]

	lst2 := []string{"ab", "a", "aaa", "cd"}
	fmt.Println(SortedListSum(lst2)) // Output: [ab cd]
}
