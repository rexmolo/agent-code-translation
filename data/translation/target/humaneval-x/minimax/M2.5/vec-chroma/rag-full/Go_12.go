package main

import "fmt"

func Longest(strings []string) interface{} {
	if len(strings) == 0 {
		return nil
	}

	maxLen := 0
	for _, s := range strings {
		if len(s) > maxLen {
			maxLen = len(s)
		}
	}

	for _, s := range strings {
		if len(s) == maxLen {
			return s
		}
	}

	return nil
}

func main() {
	// Test cases
	fmt.Println(Longest([]string{}))
	fmt.Println(Longest([]string{"a", "b", "c"}))
	fmt.Println(Longest([]string{"a", "bb", "ccc"}))
}
