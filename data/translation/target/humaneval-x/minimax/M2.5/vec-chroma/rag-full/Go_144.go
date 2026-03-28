package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Simplify(x, n string) bool {
	// Split each fraction string by "/"
	partsX := strings.Split(x, "/")
	partsN := strings.Split(n, "/")

	// Convert string parts to integers using strconv.Atoi
	a, _ := strconv.Atoi(partsX[0])
	b, _ := strconv.Atoi(partsX[1])
	c, _ := strconv.Atoi(partsN[0])
	d, _ := strconv.Atoi(partsN[1])

	// Calculate the resulting fraction: (a/b) * (c/d) = (a*c)/(b*d)
	numerator := a * c
	denom := b * d

	// Check if the division results in a whole number
	// In Python, the check `numerator/denom == int(numerator/denom)` compares
	// float division with integer division. In Go, we check if numerator
	// is evenly divisible by denom using the modulo operator.
	return numerator%denom == 0
}

func main() {
	// Test cases
	fmt.Println(Simplify("1/5", "5/1"))   // true
	fmt.Println(Simplify("1/6", "2/1"))   // false
	fmt.Println(Simplify("7/10", "10/2")) // false
}