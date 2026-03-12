package main

import (
	"fmt"
)

func ProdSigns(arr []int) interface{} {
	if len(arr) == 0 {
		return nil
	}

	// Check if any element is 0 and count negatives
	hasZero := false
	negativeCount := 0

	for _, v := range arr {
		if v == 0 {
			hasZero = true
		} else if v < 0 {
			negativeCount++
		}
	}

	// Calculate product of signs: 0 if any zero, otherwise (-1)^negativeCount
	var prod int
	if hasZero {
		prod = 0
	} else if negativeCount%2 == 0 {
		prod = 1
	} else {
		prod = -1
	}

	// Calculate sum of absolute values
	sum := 0
	for _, v := range arr {
		if v < 0 {
			sum += -v
		} else {
			sum += v
		}
	}

	return prod * sum
}

func main() {
	// Test cases
	fmt.Println(ProdSigns([]int{}))           // nil
	fmt.Println(ProdSigns([]int{1, 2, 2, -4})) // -9
	fmt.Println(ProdSigns([]int{0, 1}))        // 0
}
