package main

import "fmt"

func AnyInt(x, y, z interface{}) bool {
	// Try to convert each interface{} to int using type assertion
	xInt, okX := x.(int)
	yInt, okY := y.(int)
zInt, okZ := z.(int)

	// If any conversion fails, not all are integers
	if !okX || !okY || !okZ {
		return false
	}

	// Check if one equals the sum of the other two
	if (xInt+yInt == zInt) || (xInt+zInt == yInt) || (yInt+zInt == xInt) {
		return true
	}

	return false
}

func main() {
	// Test cases
	fmt.Println(AnyInt(5, 2, 7))   // Expected: true (5+2=7)
	fmt.Println(AnyInt(3, 2, 2))   // Expected: false
	fmt.Println(AnyInt(3, -2, 1))  // Expected: true (3+(-2)=1)
	fmt.Println(AnyInt(3.6, -2.2, 2)) // Expected: false
}
