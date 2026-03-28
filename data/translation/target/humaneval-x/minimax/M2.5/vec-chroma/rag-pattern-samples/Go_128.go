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

	// Count negative numbers
	negCount := 0
	for _, v := range arr {
		if v < 0 {
			negCount++
		}
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

	// Calculate product of signs
	var prod int
	if hasZero {
		prod = 0
	} else {
		if negCount%2 == 0 {
			prod = 1
		} else {
			prod = -1
		}
	}

	return prod * sum
}

func main() {
	// Test cases
	fmt.Println(ProdSigns([]int{1, 2, 2, -4})) // -9
	fmt.Println(ProdSigns([]int{0, 1}))        // 0
	fmt.Println(ProdSigns([]int{}))            // nil
}