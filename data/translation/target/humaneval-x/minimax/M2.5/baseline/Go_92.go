package main

import "fmt"

func AnyInt(x, y, z interface{}) bool {
	// Check if all three arguments are integers using type assertion
	xi, xOk := x.(int)
	yi, yOk := y.(int)
	zi, zOk := z.(int)

	// All three must be integers
	if xOk && yOk && zOk {
		// Check if one number equals the sum of the other two
		if xi+yi == zi || xi+zi == yi || yi+zi == xi {
			return true
		}
		return false
	}
	return false
}

func main() {
	// Test cases matching the Python examples
	fmt.Println(AnyInt(5, 2, 7))    // true
	fmt.Println(AnyInt(3, 2, 2))    // false
	fmt.Println(AnyInt(3, -2, 1))   // true
	fmt.Println(AnyInt(3.6, -2.2, 2)) // false
}