package main

import "fmt"

func AnyInt(x, y, z interface{}) bool {
	// Type assertion to check if all three values are integers
	xInt, ok1 := x.(int)
	yInt, ok2 := y.(int)
zInt, ok3 := z.(int)

	// If any of them is not an int, return false
	if !ok1 || !ok2 || !ok3 {
		return false
	}

	// Check if one equals the sum of the other two
	if xInt+yInt == zInt || xInt+zInt == yInt || yInt+zInt == xInt {
		return true
	}

	return false
}

func main() {
	// Test examples
	fmt.Println(AnyInt(5, 2, 7))   // true
	fmt.Println(AnyInt(3, 2, 2))   // false
	fmt.Println(AnyInt(3, -2, 1))  // true
	fmt.Println(AnyInt(3.6, -2.2, 2)) // false
}