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

	// Sort integers in descending order (reverse=True in Python)
	sortedArr := make([]int, len(arr))
	copy(sortedArr, arr)
	sort.Slice(sortedArr, func(i, j int) bool {
		return sortedArr[i] > sortedArr[j]
	})

	// Filter numbers between 1-9 and map to string names
	newArr := make([]string, 0)
	for _, val := range sortedArr {
		if name, exists := dic[val]; exists {
			newArr = append(newArr, name)
		}
	}

	return newArr
}