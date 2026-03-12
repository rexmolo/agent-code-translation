package main

import "fmt"

// ProdSigns computes the sum of magnitudes multiplied by the product of all signs
func ProdSigns(arr []int) interface{} {
	if len(arr) == 0 {
		return nil
	}

	hasZero := false
	negCount := 0
	sumAbs := 0

	for _, x := range arr {
		if x == 0 {
			hasZero = true
		} else if x < 0 {
			negCount++
		}
		// Calculate absolute value
		if x < 0 {
			sumAbs += -x
		} else {
			sumAbs += x
		}
	}

	var prod int
	if hasZero {
		prod = 0
	} else if negCount%2 == 0 {
		prod = 1
	} else {
		prod = -1
	}

	return prod * sumAbs
}

func main() {
	// Test cases
	fmt.Println(ProdSigns([]int{}))              // nil
	fmt.Println(ProdSigns([]int{1, 2, 2, -4}))   // -9
	fmt.Println(ProdSigns([]int{0, 1}))          // 0
}