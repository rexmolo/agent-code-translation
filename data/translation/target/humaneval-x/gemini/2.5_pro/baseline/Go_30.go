package main

import "fmt"

// GetPositive returns only the positive numbers from a slice of integers.
func GetPositive(l []int) []int {
	var positiveNums []int
	for _, num := range l {
		if num > 0 {
			positiveNums = append(positiveNums, num)
		}
	}
	return positiveNums
}

// main function to demonstrate GetPositive with examples from the docstring.
func main() {
	list1 := []int{-1, 2, -4, 5, 6}
	fmt.Println(GetPositive(list1))

	list2 := []int{5, 3, -5, 2, -3, 3, 9, 0, 123, 1, -10}
	fmt.Println(GetPositive(list2))
}
