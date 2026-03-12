package main

import (
	"fmt"
	"math"
)

func Iscube(a int) bool {
	if a < 0 {
		a = -a
	}
	root := int(math.Round(math.Pow(float64(a), 1.0/3.0)))
	return root*root*root == a
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
