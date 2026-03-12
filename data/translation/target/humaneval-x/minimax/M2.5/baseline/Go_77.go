package main

import (
	"math"
)

func Iscube(a int) bool {
	a = absInt(a)
	cubeRoot := math.Round(math.Pow(float64(a), 1.0/3.0))
	return int(cubeRoot)*int(cubeRoot)*int(cubeRoot) == a
}

func absInt(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	// Test cases
	println(Iscube(1))   // true
	println(Iscube(2))   // false
	println(Iscube(-1))  // true
	println(Iscube(64))  // true
	println(Iscube(0))   // true
	println(Iscube(180)) // false
}