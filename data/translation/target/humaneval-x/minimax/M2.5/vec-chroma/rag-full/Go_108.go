package main

import "strconv"

// CountNums returns the number of elements in arr whose digit sum is greater than 0.
// For negative numbers, the first digit is treated as negative.
// For example, -123 has signed digits -1, 2, and 3 with sum 4.
func CountNums(arr []int) int {
	count := 0
	for _, n := range arr {
		if digitSum(n) > 0 {
			count++
		}
	}
	return count
}

// digitSum calculates the sum of digits of n, with the first digit
// multiplied by neg (-1 for negative numbers, 1 for non-negative).
func digitSum(n int) int {
	neg := 1
	if n < 0 {
		n = -n
		neg = -1
	}
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum * neg
}

// For testing - could be replaced with actual tests
func main() {
	// Example usage:
	// fmt.Println(CountNums([]int{}))          // 0
	// fmt.Println(CountNums([]int{-1, 11, -11})) // 1
	// fmt.Println(CountNums([]int{1, 1, 2}))    // 3
}
