package main

import "fmt"

func AnyInt(x, y, z interface{}) bool {
	// Type assert each parameter to check if they are integers
	xInt, xOk := x.(int)
	yInt, yOk := y.(int)
zInt, zOk := z.(int)

	// If any of them is not an int, return false
	if !xOk || !yOk || !zOk {
		return false
	}

	// Check if any two numbers sum to the third
	if xInt+yInt == zInt || xInt+zInt == yInt || yInt+zInt == xInt {
		return true
	}

	return false
}

func main() {
	fmt.Println(AnyInt(5, 2, 7))      // true
	fmt.Println(AnyInt(3, 2, 2))      // false
	fmt.Println(AnyInt(3, -2, 1))     // true
	fmt.Println(AnyInt(3.6, -2.2, 2)) // false
}
