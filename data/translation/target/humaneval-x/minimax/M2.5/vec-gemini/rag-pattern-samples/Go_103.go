package main

import (
	"fmt"
	"math"
)

func RoundedAvg(n, m int) interface{} {
	if m < n {
		return -1
	}

	summation := 0
	for i := n; i <= m; i++ {
		summation += i
	}

	count := m - n + 1
	average := float64(summation) / float64(count)
	rounded := int(math.Round(average))

	return "0b" + fmt.Sprintf("%b", rounded)
}

func main() {
	// Test cases
	fmt.Println(RoundedAvg(1, 5))   // "0b11"
	fmt.Println(RoundedAvg(7, 5))   // -1
	fmt.Println(RoundedAvg(10, 20)) // "0b1111"
	fmt.Println(RoundedAvg(20, 33)) // "0b11010"
}
