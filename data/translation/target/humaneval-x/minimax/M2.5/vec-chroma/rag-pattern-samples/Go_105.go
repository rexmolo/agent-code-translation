package main

import (
	"fmt"
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

	var newArr []string
	for _, v := range sortedArr {
		if name, ok := dic[v]; ok {
			newArr = append(newArr, name)
		}
	}

	return newArr
}

func main() {
	// Test cases
	arr1 := []int{2, 1, 1, 4, 5, 8, 2, 3}
	result1 := ByLength(arr1)
	fmt.Println(result1) // [Eight Five Four Three Two Two One One]

	arr2 := []int{}
	result2 := ByLength(arr2)
	fmt.Println(result2) // []

	arr3 := []int{1, -1, 55}
	result3 := ByLength(arr3)
	fmt.Println(result3) // [One]
}
