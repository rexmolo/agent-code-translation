package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Simplify(x, n string) bool {
	// Parse x fraction
	partsX := strings.Split(x, "/")
	a, _ := strconv.Atoi(partsX[0])
	b, _ := strconv.Atoi(partsX[1])

	// Parse n fraction
	partsN := strings.Split(n, "/")
	c, _ := strconv.Atoi(partsN[0])
	d, _ := strconv.Atoi(partsN[1])

	numerator := a * c
	denom := b * d

	// Check if numerator is divisible by denominator (whole number)
	if numerator%denom == 0 {
		return true
	}
	return false
}

func main() {
	fmt.Println(Simplify("1/5", "5/1"))   // true
	fmt.Println(Simplify("1/6", "2/1"))   // false
	fmt.Println(Simplify("7/10", "10/2")) // false
}