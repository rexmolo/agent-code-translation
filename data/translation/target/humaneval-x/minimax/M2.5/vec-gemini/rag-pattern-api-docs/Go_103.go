package main

import (
	"fmt"
	"math"
)

func RoundedAvg(n, m int) interface{} {
	if m < n {
		return -1
	}

	// Sum all integers from n to m inclusive
	var summation int
	for i := n; i <= m; i++ {
		summation += i
	}

	count := float64(m - n + 1)
	avg := float64(summation) / count

	// Python's round() uses banker's rounding (round half to even)
	// Go's math.Round uses standard rounding, so we implement banker's rounding
	var rounded int
	if avg > 0 {
		rounded = int(math.Floor(avg + 0.5))
		// Adjust for banker's rounding: if we rounded up a .5, check if it should round down
		frac := avg - math.Floor(avg)
		if frac == 0.5 {
			if rounded%2 == 0 {
				// Round down (even)
				rounded = int(math.Floor(avg))
			}
		}
	} else {
		rounded = int(math.Ceil(avg - 0.5))
	}

	return "0b" + fmt.Sprintf("%b", rounded)
}

func main() {
	fmt.Println(RoundedAvg(1, 5))   // "0b11"
	fmt.Println(RoundedAvg(7, 5))    // -1
	fmt.Println(RoundedAvg(10, 20)) // "0b1111"
	fmt.Println(RoundedAvg(20, 33)) // "0b11010"
}