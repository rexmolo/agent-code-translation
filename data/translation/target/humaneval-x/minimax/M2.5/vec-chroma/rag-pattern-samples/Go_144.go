package main

import (
	"fmt"
	"strings"
)

func Simplify(x, n string) bool {
	// Split both fractions into numerator and denominator
	xParts := strings.Split(x, "/")
	nParts := strings.Split(n, "/")

	// Parse the numerators and denominators
	a := xParts[0]
	b := xParts[1]
	c := nParts[0]
	d := nParts[1]

	// Convert to integers
	var aNum, bDen, cNum, dDen int
	fmt.Sscanf(a, "%d", &aNum)
	fmt.Sscanf(b, "%d", &bDen)
	fmt.Sscanf(c, "%d", &cNum)
	fmt.Sscanf(d, "%d", &dDen)

	// Multiply: (a/b) * (c/d) = (a*c) / (b*d)
	numerator := aNum * cNum
	denom := bDen * dDen

	// Check if the result is a whole number
	// A number is whole if numerator is divisible by denominator
	if numerator%denom == 0 {
		return true
	}
	return false
}
