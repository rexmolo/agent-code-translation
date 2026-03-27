package main

import (
	"math"
	"strconv"
)

func RoundedAvg(n, m int) interface{} {
	if n > m {
		return -1
	}

	// Calculate sum of integers from n to m (inclusive)
	// Using arithmetic series formula: sum = (first + last) * count / 2
	count := m - n + 1
	sum := (n + m) * count / 2

	// Calculate average and round to nearest integer
	avg := float64(sum) / float64(count)
	rounded := int(math.Floor(avg + 0.5))

	// Convert to binary string with "0b" prefix
	return "0b" + strconv.FormatInt(int64(rounded), 2)
}

func main() {
	// Test cases
	result1 := RoundedAvg(1, 5)
	println(result1)

	result2 := RoundedAvg(7, 5)
	println(result2)

	result3 := RoundedAvg(10, 20)
	println(result3)

	result4 := RoundedAvg(20, 33)
	println(result4)
}
