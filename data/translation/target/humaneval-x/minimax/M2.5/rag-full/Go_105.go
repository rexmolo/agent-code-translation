package main

import (
	"fmt"
	"sort"
)

func ByLength(arr []int) []string {
	dict := map[int]string{
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

	// Sort in descending order (reverse=True)
	sortedArr := make([]int, len(arr))
	copy(sortedArr, arr)
	sort.Sort(sort.Reverse(sort.IntSlice(sortedArr)))

	var newArr []string
	for _, v := range sortedArr {
		if name, ok := dict[v]; ok {
			newArr = append(newArr, name)
		}
	}

	return newArr
}

func main() {
	// Test examples
	fmt.Println(ByLength([]int{2, 1, 1, 4, 5, 8, 2, 3}))
	fmt.Println(ByLength([]int{}))
	fmt.Println(ByLength([]int{1, -1, 55}))
}
