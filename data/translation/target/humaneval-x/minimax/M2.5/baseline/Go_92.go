package main

import "fmt"

func AnyInt(x, y, z interface{}) bool {
	// Try type assertion for each parameter to check if they are integers
	xVal, xOk := x.(int)
	yVal, yOk := y.(int)
zVal, zOk := z.(int)

	// Check if all three values are integers
	if xOk && yOk && zOk {
		// Check if one equals the sum of the other two
		if xVal+yVal == zVal || xVal+zVal == yVal || yVal+zVal == xVal {
			return true
		}
		return false
	}
	return false
}

func main() {
	fmt.Println(AnyInt(5, 2, 7))       // true
	fmt.Println(AnyInt(3, 2, 2))       // false
	fmt.Println(AnyInt(3, -2, 1))      // true
	fmt.Println(AnyInt(3.6, -2.2, 2)) // false
}
