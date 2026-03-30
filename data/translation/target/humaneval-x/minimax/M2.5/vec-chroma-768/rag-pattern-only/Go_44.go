package main

import "fmt"

func ChangeBase(x int, base int) string {
	if x == 0 {
		return ""
	}

	var digits []byte
	for x > 0 {
		digit := x % base
		digits = append(digits, byte('0'+digit))
		x /= base
	}

	// Reverse the digits to get correct order
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}

	return string(digits)
}

func main() {
	fmt.Println(ChangeBase(8, 3))  // Expected: 22
	fmt.Println(ChangeBase(8, 2))  // Expected: 1000
	fmt.Println(ChangeBase(7, 2)) // Expected: 111
}
