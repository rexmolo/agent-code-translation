package main

import "slices"

func Longest(strings []string) interface{} {
	if len(strings) == 0 {
		return nil
	}

	// Use slices.MaxFunc to find the longest string
	// Return first one in case of ties (when compare returns 0)
	longest := slices.MaxFunc(strings, func(a, b string) int {
		return len(a) - len(b)
	})

	return longest
}