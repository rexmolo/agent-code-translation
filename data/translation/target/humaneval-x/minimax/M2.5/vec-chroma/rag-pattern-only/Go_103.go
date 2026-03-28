package main

import (
	"math"
	"strconv"
)

func RoundedAvg(n, m int) interface{} {
	if m < n {
		return -1
	}

	var sum int64 = 0
	for i := n; i <= m; i++ {
		sum += int64(i)
	}

	count := float64(m - n + 1)
	avg := float64(sum) / count
	rounded := int64(math.Round(avg))

	return "0b" + strconv.FormatInt(rounded, 2)
}

func main() {
	// Example usage
	// RoundedAvg(1, 5)    // "0b11"
	// RoundedAvg(7, 5)     // -1
	// RoundedAvg(10, 20)   // "0b1111"
	// RoundedAvg(20, 33)   // "0b11010"
}
