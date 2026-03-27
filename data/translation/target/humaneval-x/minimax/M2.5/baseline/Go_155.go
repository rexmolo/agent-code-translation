package main

import (
	"fmt"
	"math"
	"strconv"
)

func EvenOddCount(num int) [2]int {
	evenCount := 0
	oddCount := 0

	strNum := strconv.Itoa(int(math.Abs(float64(num))))

	for _, r := range strNum {
		digit := int(r - '0')
		if digit%2 == 0 {
			evenCount++
		} else {
			oddCount++
		}
	}

	return [2]int{evenCount, oddCount}
}

func main() {
	fmt.Println(EvenOddCount(-12)) // Output: [1 1]
	fmt.Println(EvenOddCount(123)) // Output: [1 2]
}