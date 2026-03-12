package main

import "fmt"

// IsEqualToSumEven evaluates whether the given number n can be written as the sum of exactly 4 positive even numbers.
// Example:
// IsEqualToSumEven(4) == false
// IsEqualToSumEven(6) == false
// IsEqualToSumEven(8) == true
func IsEqualToSumEven(n int) bool {
	// The smallest sum of 4 positive even numbers is 2 + 2 + 2 + 2 = 8.
	// The sum of four even numbers is always even.
	// Therefore, any even number n >= 8 can be expressed as the sum of four positive even numbers.
	// For example, n = 2 + 2 + 2 + (n - 6). Since n is an even number >= 8, (n - 6) is a positive even number.
	return n%2 == 0 && n >= 8
}

func main() {
	fmt.Println(IsEqualToSumEven(4))  // Expected: false
	fmt.Println(IsEqualToSumEven(6))  // Expected: false
	fmt.Println(IsEqualToSumEven(8))  // Expected: true
	fmt.Println(IsEqualToSumEven(10)) // Expected: true
	fmt.Println(IsEqualToSumEven(11)) // Expected: false
}