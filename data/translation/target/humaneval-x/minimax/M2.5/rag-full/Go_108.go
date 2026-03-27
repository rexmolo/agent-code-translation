package main

import "fmt"

func CountNums(arr []int) int {
	digitsSum := func(n int) int {
		neg := 1
		if n < 0 {
			n = -n
			neg = -1
		}
		// Convert number to string and extract each digit
		strN := fmt.Sprintf("%d", n)
		sum := 0
		for i, c := range strN {
			digit := int(c - '0')
			if i == 0 {
				digit *= neg
			}
			sum += digit
		}
		return sum
	}

	count := 0
	for _, n := range arr {
		if digitsSum(n) > 0 {
			count++
		}
	}
	return count
}

func main() {
	// Test cases
	fmt.Println(CountNums([]int{}))           // Expected: 0
	fmt.Println(CountNums([]int{-1, 11, -11})) // Expected: 1
	fmt.Println(CountNums([]int{1, 1, 2}))     // Expected: 3
}