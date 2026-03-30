package main

import (
	"fmt"
	"reflect"
)

func AnyInt(x, y, z interface{}) bool {
	// Check if all three values are integers using reflection
	xVal := reflect.ValueOf(x)
	yVal := reflect.ValueOf(y)
	zVal := reflect.ValueOf(z)

	// Ensure all are of type int (not float64, etc.)
	if xVal.Kind() != reflect.Int ||
		yVal.Kind() != reflect.Int ||
		zVal.Kind() != reflect.Int {
		return false
	}

	// Convert to int64 for calculations
	xi, yi, zi := xVal.Int(), yVal.Int(), zVal.Int()

	// Check if one equals the sum of the other two
	return xi+yi == zi || xi+zi == yi || yi+zi == xi
}

func main() {
	fmt.Println(AnyInt(5, 2, 7))       // true
	fmt.Println(AnyInt(3, 2, 2))      // false
	fmt.Println(AnyInt(3, -2, 1))     // true
	fmt.Println(AnyInt(3.6, -2.2, 2)) // false
}