package main

import (
	"fmt"
)

func TotalMatch(lst1 []string, lst2 []string) []string {
	l1 := 0
	for _, st := range lst1 {
		l1 += len(st)
	}

	l2 := 0
	for _, st := range lst2 {
		l2 += len(st)
	}

	if l1 <= l2 {
		return lst1
	}
	return lst2
}

func main() {
	// Test cases
	testCases := []struct {
		name  string
		lst1  []string
		lst2  []string
	}{
		{"Empty lists", []string{}, []string{}},
		{"First example", []string{"hi", "admin"}, []string{"hI", "Hi"}},
		{"Second example", []string{"hi", "admin"}, []string{"hi", "hi", "admin", "project"}},
		{"Third example", []string{"hi", "admin"}, []string{"hI", "hi", "hi"}},
		{"Fourth example", []string{"4"}, []string{"1", "2", "3", "4", "5"}},
	}

	for _, tc := range testCases {
		result := TotalMatch(tc.lst1, tc.lst2)
		fmt.Printf("%s: %v\n", tc.name, result)
	}
}
