package main

import "fmt"

func GetPositive(l []int) []int {
	var result []int
	for _, v := range l {
		if v > 0 {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println(GetPositive([]int{-1, 2, -4, 5, 6}))
	// Test case 2
	fmt.Println(GetPositive([]int{5, 3, -5, 2, -3, 3, 9, 0, 123, 1, -10}))
}