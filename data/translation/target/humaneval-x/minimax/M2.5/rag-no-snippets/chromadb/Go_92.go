package main

import (
	"fmt"
	"reflect"
)

func AnyInt(x, y, z interface{}) bool {
	// Check if all values are integers using reflection
	xv := reflect.ValueOf(x)
	yv := reflect.ValueOf(y)
	zv := reflect.ValueOf(z)

	// Helper to check if a reflect.Kind is an integer type
	isIntKind := func(kind reflect.Kind) bool {
		switch kind {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return true
		}
		return false
	}

	// Check if all three values are integers
	if !isIntKind(xv.Kind()) || !isIntKind(yv.Kind()) || !isIntKind(zv.Kind()) {
		return false
	}

	// Convert to int64 for arithmetic comparison
	xn := xv.Int()
	yn := yv.Int()
	zn := zv.Int()

	// Check if one equals the sum of the other two
	return (xn+yn == zn) || (xn+zn == yn) || (yn+zn == xn)
}

func main() {
	// Test cases
	fmt.Println(AnyInt(5, 2, 7))    // true
	fmt.Println(AnyInt(3, 2, 2))     // false
	fmt.Println(AnyInt(3, -2, 1))   // true
	fmt.Println(AnyInt(3.6, -2.2, 2)) // false
}
