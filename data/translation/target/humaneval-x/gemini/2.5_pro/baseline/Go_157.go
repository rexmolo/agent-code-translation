package main

import "fmt"

// RightAngleTriangle checks if the three sides form a right-angled triangle.
// Given the lengths of the three sides of a triangle, it returns true if they
// form a right-angled triangle, false otherwise.
// A right-angled triangle is a triangle in which one angle is a right angle (90 degrees).
// Example:
// RightAngleTriangle(3, 4, 5) == true
// RightAngleTriangle(1, 2, 3) == false
func RightAngleTriangle(a, b, c int) bool {
	// The Pythagorean theorem states that in a right-angled triangle,
	// the square of the hypotenuse (the side opposite the right angle)
	// is equal to the sum of the squares of the other two sides.
	// We check all three possibilities for which side could be the hypotenuse.
	return a*a == b*b+c*c || b*b == a*a+c*c || c*c == a*a+b*b
}

func main() {
	fmt.Println("Is 3, 4, 5 a right-angled triangle?", RightAngleTriangle(3, 4, 5))
	fmt.Println("Is 1, 2, 3 a right-angled triangle?", RightAngleTriangle(1, 2, 3))
	fmt.Println("Is 5, 12, 13 a right-angled triangle?", RightAngleTriangle(5, 12, 13))
}
