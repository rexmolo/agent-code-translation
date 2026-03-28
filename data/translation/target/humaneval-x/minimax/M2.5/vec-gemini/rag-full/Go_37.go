package main

import (
	"fmt"
	"slices"
)

func SortEven(l []int) []int {
	evens := l[::2]
	odds := l[1::2]
	slices.Sort(evens)
	ans := make([]int, 0, len(l))
	for i := 0; i < len(odds); i++ {
		ans = append(ans, evens[i], odds[i])
	}
	if len(evens) > len(odds) {
		ans = append(ans, evens[len(evens)-1])
	}
	return ans
}

func main() {
	fmt.Println(SortEven([]int{1, 2, 3}))
	fmt.Println(SortEven([]int{5, 6, 3, 4}))
}
