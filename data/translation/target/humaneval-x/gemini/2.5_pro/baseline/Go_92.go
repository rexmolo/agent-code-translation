package main

import "fmt"

// AnyInt checks if three numbers are all integers and if one of them
// is the sum of the other two.
func AnyInt(x, y, z interface{}) bool {
	// Use type assertion to check if each interface value holds an int.
	// The 'ok' variable will be true if the assertion succeeds, and false otherwise.
	xInt, okX := x.(int)
	yInt, okY := y.(int)
	zInt, okZ := z.(int)

	// The function should only return true if all three are integers.
	// The '&&' operator short-circuits, so the sum comparison is only
	// done if okX, okY, and okZ are all true.
	if okX && okY && okZ {
		// Check if the sum of any two numbers equals the third.
		if (xInt+yInt == zInt) || (xInt+zInt == yInt) || (yInt+zInt == xInt) {
			return true
		}
	}

	// Return false in all other cases.
	return false
}

func main() {
	// Examples from the original Python docstring
	fmt.Println("any_int(5, 2, 7) ->", AnyInt(5, 2, 7))
	fmt.Println("any_int(3, 2, 2) ->", AnyInt(3, 2, 2))
	fmt.Println("any_int(3, -2, 1) ->", AnyInt(3, -2, 1))
	fmt.Println("any_int(3.6, -2.2, 2) ->", AnyInt(3.6, -2.2, 2))

	// Additional test case with non-numeric type
	fmt.Println("any_int(5, 2, \"hello\") ->", AnyInt(5, 2, "hello"))
}
