package main

import (
	"fmt"
	"math"
)

func RoundedAvg(n, m int) interface{} {
	if m < n {
		return -1
	}

	sum := 0
	for i := n; i <= m; i++ {
		sum += i
	}

	avg := float64(sum) / float64(m-n+1)
	rounded := int(math.Round(avg))

	return fmt.Sprintf("%b", rounded)
}
