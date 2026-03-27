package main

import (
	"fmt"
	"strconv"
	"strings"
)

func CompareOne(a, b interface{}) interface{} {
	// Convert to string representation first
	tempA := fmt.Sprintf("%v", a)
	tempB := fmt.Sprintf("%v", b)
	
	// Replace comma with dot for string values (decimal separator)
	tempA = strings.Replace(tempA, ",", ".", -1)
	tempB = strings.Replace(tempB, ",", ".", -1)
	
	// Convert to float64 for numeric comparison
	floatA, errA := strconv.ParseFloat(tempA, 64)
	floatB, errB := strconv.ParseFloat(tempB, 64)
	
	// If conversion fails, return nil (shouldn't happen with valid input)
	if errA != nil || errB != nil {
		return nil
	}
	
	// Compare and return the appropriate value
	if floatA == floatB {
		return nil
	}
	if floatA > floatB {
		return a
	}
	return b
}

func main() {
	// Test cases
	fmt.Println(CompareOne(1, 2.5))       // 2.5
	fmt.Println(CompareOne(1, "2,3"))      // 2,3
	fmt.Println(CompareOne("5,1", "6"))   // 6
	fmt.Println(CompareOne("1", 1))        // <nil>
}