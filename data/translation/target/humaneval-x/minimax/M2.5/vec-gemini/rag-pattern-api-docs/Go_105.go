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

	// Sort in reverse (descending) order using sort.Reverse
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))

	var newArr []string
	for _, v := range arr {
		// Check if the value exists in the map (i.e., is between 1 and 9)
		if name, ok := dic[v]; ok {
			newArr = append(newArr, name)
		}
		// Values not in the map are silently ignored (equivalent to Python's try/except pass)
	}

	return newArr
}

func main() {
	// Test case 1: normal case
	arr1 := []int{2, 1, 1, 4, 5, 8, 2, 3}
	result1 := ByLength(arr1)
	fmt.Println(result1)
	// Expected: [Eight Five Four Three Two Two One One]

	// Test case 2: empty array
	arr2 := []int{}
	result2 := ByLength(arr2)
	fmt.Println(result2)
	// Expected: []

	// Test case 3: array with strange numbers
	arr3 := []int{1, -1, 55}
	result3 := ByLength(arr3)
	fmt.Println(result3)
	// Expected: [One]
}
