package main

import (
	"fmt"
	"math"
	"strconv"
)

func RoundedAvg(n, m int) interface{} {
	if n > m {
		return -1
	}

	// Calculate the sum of integers from n to m (inclusive)
	// Using the formula: sum = (n + m) * count / 2
	count := m - n + 1
	sum := (n + m) * count / 2

	// Calculate average and round to nearest integer
	avg := math.Round(float64(sum) / float64(count))

	// Convert to binary string with "0b" prefix
	return "0b" + strconv.FormatInt(int64(avg), 2)
}

func main() {
	fmt.Println(RoundedAvg(1, 5))
	fmt.Println(RoundedAvg(7, 5))
	fmt.Println(RoundedAvg(10, 20))
	fmt.Println(RoundedAvg(20, 33))
}