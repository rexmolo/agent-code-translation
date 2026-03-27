package main

import (
	"fmt"
	"slices"
)

func Common(l1 []int, l2 []int) []int {
	ret := make(map[int]struct{})
	for _, e1 := range l1 {
		for _, e2 := range l2 {
			if e1 == e2 {
				ret[e1] = struct{}{}
			}
		}
	}

	result := make([]int, 0, len(ret))
	for k := range ret {
		result = append(result, k)
	}
	slices.Sort(result)
	return result
}

func main() {
	fmt.Println(Common([]int{1, 4, 3, 34, 653, 2, 5}, []int{5, 7, 1, 5, 9, 653, 121}))
	fmt.Println(Common([]int{5, 3, 2, 8}, []int{3, 2}))
}
