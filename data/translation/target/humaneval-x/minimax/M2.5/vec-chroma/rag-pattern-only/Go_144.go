package main

import (
	"strconv"
	"strings"
)

func Simplify(x, n string) bool {
	// Split x into numerator and denominator
	partsX := strings.Split(x, "/")
	a := partsX[0]
	b := partsX[1]
	
	// Split n into numerator and denominator
	partsN := strings.Split(n, "/")
	c := partsN[0]
	d := partsN[1]
	
	// Convert to integers
	numA, _ := strconv.Atoi(a)
	denomB, _ := strconv.Atoi(b)
	numC, _ := strconv.Atoi(c)
	denomD, _ := strconv.Atoi(d)
	
	// Multiply numerators and denominators
	numerator := numA * numC
	denom := denomB * denomD
	
	// Check if it's a whole number (denominator divides evenly into numerator)
	return numerator%denom == 0
}

func main() {
	// Test cases
	println(Simplify("1/5", "5/1"))   // true
	println(Simplify("1/6", "2/1"))   // false
	println(Simplify("7/10", "10/2")) // false
}