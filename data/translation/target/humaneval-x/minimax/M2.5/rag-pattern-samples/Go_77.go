package main

import (
	"fmt"
	"math"
)

func Iscube(a int) bool {
	absA := math.Abs(float64(a))
	root := math.Round(math.Cbrt(absA))
	return root*root*root == absA
}

func main() {
	// Test cases
	fmt.Println(Iscube(1))   // True
	fmt.Println(Iscube(2))   // False
	fmt.Println(Iscube(-1))  // True
	fmt.Println(Iscube(64))  // True
	fmt.Println(Iscube(0))   // True
	fmt.Println(Iscube(180)) // False
}