package main

import (
	"fmt"
	"math"
	"strconv"
)

func RoundedAvg(n, m int) interface{} {
	if m < n {
		return -1
	}

	sum := 0
	for i := n; i <= m; i++ {
		sum += i
	}

	count := m - n + 1
	avg := float64(sum) / float64(count)
	rounded := int(math.Round(avg))

	return "0b" + strconv.FormatInt(int64(rounded), 2)
}

func main() {
	fmt.Println(RoundedAvg(1, 5))    // 0b11
	fmt.Println(RoundedAvg(7, 5))    // -1
	fmt.Println(RoundedAvg(10, 20))  // 0b1111
	fmt.Println(RoundedAvg(20, 33))  // 0b11011
}
