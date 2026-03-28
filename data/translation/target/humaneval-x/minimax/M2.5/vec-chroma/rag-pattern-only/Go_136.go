package main

import "fmt"

func LargestSmallestIntegers(lst []int) [2]interface{} {
	var largestNegative int
	var smallestPositive int
	hasNegative := false
	hasPositive := false

	for _, x := range lst {
		if x < 0 {
			if !hasNegative || x > largestNegative {
				largestNegative = x
				hasNegative = true
			}
		}
		if x > 0 {
			if !hasPositive || x < smallestPositive {
				smallestPositive = x
				hasPositive = true
			}
		}
	}

	var result [2]interface{}
	if hasNegative {
		result[0] = largestNegative
	}
	if hasPositive {
		result[1] = smallestPositive
	}

	return result
}

func main() {
	// Test cases
	fmt.Println(LargestSmallestIntegers([]int{2, 4, 1, 3, 5, 7})) // [nil 1]
	fmt.Println(LargestSmallestIntegers([]int{}))                // [nil nil]
	fmt.Println(LargestSmallestIntegers([]int{0}))               // [nil nil]
}