package main

import (
	"sort"
)

func ByLength(arr []int) []string {
	dic := map[int]string{
		1: "One",
		2: "Two",
		3: "Three",
		4: "Four",
		5: "Five",
		6: "Six",
		7: "Seven",
		8: "Eight",
		9: "Nine",
	}

	// Sort in reverse (descending) order
	sortedArr := make([]int, len(arr))
	copy(sortedArr, arr)
	sort.Slice(sortedArr, func(i, j int) bool {
		return sortedArr[i] > sortedArr[j]
	})

	newArr := []string{}
	for _, v := range sortedArr {
		if val, ok := dic[v]; ok {
			newArr = append(newArr, val)
		}
	}
	return newArr
}