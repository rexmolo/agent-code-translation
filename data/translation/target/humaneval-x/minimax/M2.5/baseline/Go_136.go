package main

import "fmt"

func LargestSmallestIntegers(lst []int) [2]interface{} {
	var maxNegative *int
	var minPositive *int

	for _, num := range lst {
		if num < 0 {
			if maxNegative == nil || num > *maxNegative {
				maxNegative = &num
			}
		}
		if num > 0 {
			if minPositive == nil || num < *minPositive {
				minPositive = &num
			}
		}
	}

	result := [2]interface{}{}

	if maxNegative != nil {
		result[0] = *maxNegative
	}

	if minPositive != nil {
		result[1] = *minPositive
	}

	return result
}

func main() {
	// Test cases
	fmt.Println(LargestSmallestIntegers([]int{2, 4, 1, 3, 5, 7}))  // [<nil> <nil>]
	fmt.Println(LargestSmallestIntegers([]int{}))                // [<nil> <nil>]
	fmt.Println(LargestSmallestIntegers([]int{0}))               // [<nil> <nil>]
	fmt.Println(LargestSmallestIntegers([]int{-5, -3, -1, 1, 3})) // [-1 1]
}