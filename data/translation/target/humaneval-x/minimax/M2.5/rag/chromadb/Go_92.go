package main

import "fmt"

func AnyInt(x, y, z interface{}) bool {
	// Check if all three values are integers using type assertions
	xi, ok1 := x.(int)
	yi, ok2 := y.(int)
	zi, ok3 := z.(int)

	// If any value is not an int, return false
	if !ok1 || !ok2 || !ok3 {
		return false
	}

	// Check if one number equals the sum of the other two
	if xi+yi == zi || xi+zi == yi || yi+zi == xi {
		return true
	}

	return false
}

func main() {
	// Test cases
	fmt.Println(AnyInt(5, 2, 7))      // true
	fmt.Println(AnyInt(3, 2, 2))      // false
	fmt.Println(AnyInt(3, -2, 1))     // true
	fmt.Println(AnyInt(3.6, -2.2, 2)) // false
}
