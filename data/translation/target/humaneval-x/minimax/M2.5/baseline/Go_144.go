package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Simplify(x, n string) bool {
	// Parse x = a/b
	partsX := strings.Split(x, "/")
	a, _ := strconv.Atoi(partsX[0])
	b, _ := strconv.Atoi(partsX[1])

	// Parse n = c/d
	partsN := strings.Split(n, "/")
	c, _ := strconv.Atoi(partsN[0])
	d, _ := strconv.Atoi(partsN[1])

	// Multiply: (a/b) * (c/d) = (a*c)/(b*d)
	numerator := a * c
	denominator := b * d

	// Check if result is a whole number by checking divisibility
	return numerator%denominator == 0
}

func main() {
	fmt.Println(Simplify("1/5", "5/1"))   // true
	fmt.Println(Simplify("1/6", "2/1"))   // false
	fmt.Println(Simplify("7/10", "10/2")) // false
}