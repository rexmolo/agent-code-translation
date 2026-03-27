package main

import (
	"fmt"
)

func ProdSigns(arr []int) interface{} {
	if len(arr) == 0 {
		return nil
	}

	// Check for zero and count negatives
	hasZero := false
	negCount := 0
	sumAbs := 0

	for _, v := range arr {
		if v == 0 {
			hasZero = true
		}
		if v < 0 {
			negCount++
		}
		sumAbs += abs(v)
	}

	// Calculate product of signs
	prod := 1
	if hasZero {
		prod = 0
	} else if negCount%2 == 1 {
		prod = -1
	}

	return prod * sumAbs
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	// Test cases
	fmt.Println(ProdSigns([]int(nil)))      // nil
	fmt.Println(ProdSigns([]int{}))          // nil
	fmt.Println(ProdSigns([]int{1, 2, 2, -4})) // -9
	fmt.Println(ProdSigns([]int{0, 1}))      // 0
}
