package main

import "cmp"

func GenerateIntegers(a, b int) []int {
	lower := cmp.Max(2, cmp.Min(a, b))
	upper := cmp.Min(8, cmp.Max(a, b))

	var result []int
	for i := lower; i <= upper; i++ {
		if i%2 == 0 {
			result = append(result, i)
		}
	}

	return result
}