package main

import "fmt"

func ProdSigns(arr []int) interface{} {
	if len(arr) == 0 {
		return nil
	}

	// Calculate sum of absolute values and check for zero/negative count
	hasZero := false
	negativeCount := 0
	sumAbs := 0

	for _, v := range arr {
		if v == 0 {
			hasZero = true
		}
		if v < 0 {
			negativeCount++
		}
		// Calculate absolute value (Go has no built-in abs for int)
		if v < 0 {
			sumAbs += -v
		} else {
			sumAbs += v
		}
	}

	// Calculate product of signs: 0 if any zero, otherwise (-1)^negativeCount
	prod := 0
	if !hasZero {
		if negativeCount%2 == 0 {
			prod = 1
		} else {
			prod = -1
		}
	}

	return prod * sumAbs
}

func main() {
	// Test cases
	fmt.Println(ProdSigns([]int(nil)))         // nil (empty)
	fmt.Println(ProdSigns([]int{}))            // nil (empty)
	fmt.Println(ProdSigns([]int{1, 2, 2, -4})) // -9
	fmt.Println(ProdSigns([]int{0, 1}))        // 0
}