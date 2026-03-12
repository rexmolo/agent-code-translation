package main

import "fmt"

// ProdSigns calculates the sum of magnitudes of integers multiplied by the product of all signs.
// The sign product is 1 for an even number of negative values, -1 for an odd number, and 0 if any value is zero.
// It returns nil for an empty slice, which is the Go equivalent of Python's None.
func ProdSigns(arr []int) interface{} {
	if len(arr) == 0 {
		return nil
	}

	signProduct := 1
	sumOfMagnitudes := 0
	hasZero := false

	for _, val := range arr {
		if val == 0 {
			hasZero = true
		} else if val < 0 {
			signProduct *= -1
		}

		if val < 0 {
			sumOfMagnitudes += -val
		} else {
			sumOfMagnitudes += val
		}
	}

	if hasZero {
		return 0
	}

	return signProduct * sumOfMagnitudes
}

func main() {
	// Example 1: prod_signs([1, 2, 2, -4]) == -9
	arr1 := []int{1, 2, 2, -4}
	fmt.Printf("Input: %v, Output: %v\n", arr1, ProdSigns(arr1))

	// Example 2: prod_signs([0, 1]) == 0
	arr2 := []int{0, 1}
	fmt.Printf("Input: %v, Output: %v\n", arr2, ProdSigns(arr2))

	// Example 3: prod_signs([]) == None
	arr3 := []int{}
	fmt.Printf("Input: %v, Output: %v\n", arr3, ProdSigns(arr3))
}
