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
	// Test the function
	fmt.Println(TotalMatch([]string{}, []string{}))
	fmt.Println(TotalMatch([]string{"hi", "admin"}, []string{"hI", "Hi"}))
	fmt.Println(TotalMatch([]string{"hi", "admin"}, []string{"hi", "hi", "admin", "project"}))
	fmt.Println(TotalMatch([]string{"hi", "admin"}, []string{"hI", "hi", "hi"}))
	fmt.Println(TotalMatch([]string{"4"}, []string{"1", "2", "3", "4", "5"}))
}