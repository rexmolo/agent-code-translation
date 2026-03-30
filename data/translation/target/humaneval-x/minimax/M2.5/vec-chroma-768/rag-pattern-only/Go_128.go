package main

import "fmt"

func ProdSigns(arr []int) interface{} {
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

	// Calculate product of signs
	var prod int
	if hasZero {
		prod = 0
	} else {
		// Count negative numbers
		negCount := 0
		for _, v := range arr {
			if v < 0 {
				negCount++
			}
		}
		// (-1)^negCount
		if negCount%2 == 0 {
			prod = 1
		} else {
			prod = -1
		}
	}

	// Calculate sum of absolute values
	sumAbs := 0
	for _, v := range arr {
		sumAbs += abs(v)
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
	fmt.Println(ProdSigns([]int{}))           // nil
	fmt.Println(ProdSigns([]int{0, 1}))       // 0
	fmt.Println(ProdSigns([]int{1, 2, 2, -4})) // -9
}
