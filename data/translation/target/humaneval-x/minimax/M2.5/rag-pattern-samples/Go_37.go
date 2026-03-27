package main

import (
	"fmt"
	"sort"
)

func SortEven(l []int) []int {
	var evens []int
	var odds []int

	for i := 0; i < len(l); i++ {
		if i%2 == 0 {
			evens = append(evens, l[i])
		} else {
			odds = append(odds, l[i])
		}
	}

	sort.Ints(evens)

	var ans []int
	for i := 0; i < len(evens); i++ {
		ans = append(ans, evens[i])
		if i < len(odds) {
			ans = append(ans, odds[i])
		}
	}

	return ans
}

func main() {
	// Example usage
	fmt.Println(SortEven([]int{1, 2, 3}))
	fmt.Println(SortEven([]int{5, 6, 3, 4}))
}