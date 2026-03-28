package main

import (
	"fmt"
	"sort"
)

func SortEven(l []int) []int {
	evens := make([]int, 0)
	odds := make([]int, 0)

	for i := 0; i < len(l); i++ {
		if i%2 == 0 {
			evens = append(evens, l[i])
		} else {
			odds = append(odds, l[i])
		}
	}

	sort.Ints(evens)

	ans := make([]int, 0)
	minLen := len(evens)
	if len(odds) < minLen {
		minLen = len(odds)
	}

	for i := 0; i < minLen; i++ {
		ans = append(ans, evens[i], odds[i])
	}

	if len(evens) > len(odds) {
		ans = append(ans, evens[len(evens)-1])
	}

	return ans
}

func main() {
	// Test cases
	fmt.Println(SortEven([]int{1, 2, 3}))          // [1 2 3]
	fmt.Println(SortEven([]int{5, 6, 3, 4}))       // [3 6 5 4]
}
