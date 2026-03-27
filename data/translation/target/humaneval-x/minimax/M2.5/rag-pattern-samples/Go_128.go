package main

import "fmt"

func ProdSigns(arr []int) interface{} {
	// Return nil for empty array (equivalent to Python's None)
	if len(arr) == 0 {
		return nil
	}

	// Check if 0 is in the array
	hasZero := false
	for _, v := range arr {
		if v == 0 {
			hasZero = true
			break
		}
	}
	if hasZero {
		return 0
	}

	// Count negative elements
	negativeCount := 0
	for _, v := range arr {
		if v < 0 {
			negativeCount++
		}
	}

	// Product of signs: (-1)^(number of negative elements)
	prod := 1
	if negativeCount%2 == 1 {
		prod = -1
	}

	// Sum of absolute values
	sumAbs := 0
	for _, v := range arr {
		if v < 0 {
			sumAbs += -v
		} else {
			sumAbs += v
		}
	}

	return prod * sumAbs
}

func main() {
	// Test cases
	fmt.Println(ProdSigns([]int(nil)))    // nil (empty)
	fmt.Println(ProdSigns([]int{}))       // nil (empty)
	fmt.Println(ProdSigns([]int{1, 2, 2, -4})) // -9
	fmt.Println(ProdSigns([]int{0, 1}))   // 0
	fmt.Println(ProdSigns([]int{-1, -2, -3})) // -6 (6 * 1)
	fmt.Println(ProdSigns([]int{-1, 2, 3}))   // -6 (6 * -1)
}