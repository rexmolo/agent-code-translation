package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func ClosestInteger(value string) int {
	// Remove trailing zeros after decimal point
	if strings.Contains(value, ".") {
		value = strings.TrimRight(value, "0")
		// Also remove trailing decimal point if present
		value = strings.TrimRight(value, ".")
	}

	num, _ := strconv.ParseFloat(value, 64)

	var res float64

	// Check if value ends with ".5"
	if strings.HasSuffix(value, ".5") {
		if num > 0 {
			res = math.Ceil(num)
		} else {
			res = math.Floor(num)
		}
	} else if len(value) > 0 {
		res = math.Round(num)
	} else {
		res = 0
	}

	return int(res)
}

func main() {
	// Test cases
	fmt.Println(ClosestInteger("10"))     // 10
	fmt.Println(ClosestInteger("15.3"))  // 15
	fmt.Println(ClosestInteger("14.5"))  // 15
	fmt.Println(ClosestInteger("-14.5")) // -15
}
