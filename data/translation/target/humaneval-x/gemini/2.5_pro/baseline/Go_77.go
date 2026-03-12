package main

import (
	"fmt"
	"math"
)

// Iscube checks if an integer is a perfect cube.
// It's a translation of the provided Python function.
func Iscube(a int) bool {
	// The Python version takes the absolute value of `a` first.
	// `a = abs(a)`
	absA := a
	if a < 0 {
		absA = -a
	}

	// The Python code `int(round(a ** (1. / 3)))` is translated below.
	// We convert to float64, calculate the cube root, round it, and cast back to int.
	// math.Cbrt is preferred over math.Pow for cube roots due to precision.
	root := int(math.Round(math.Cbrt(float64(absA))))

	// The final step in Python is `... ** 3 == a`.
	// We cube the integer root and compare it with the absolute value.
	return root*root*root == absA
}

// main is a driver function to test Iscube with the provided examples.
func main() {
	fmt.Printf("iscube(1) ==> %t\n", Iscube(1))
	fmt.Printf("iscube(2) ==> %t\n", Iscube(2))
	fmt.Printf("iscube(-1) ==> %t\n", Iscube(-1))
	fmt.Printf("iscube(64) ==> %t\n", Iscube(64))
	fmt.Printf("iscube(0) ==> %t\n", Iscube(0))
	fmt.Printf("iscube(180) ==> %t\n", Iscube(180))
}
