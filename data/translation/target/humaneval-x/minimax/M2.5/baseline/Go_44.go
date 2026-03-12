package main

import (
	"bufio"
	"fmt"
	"os"
)

func ChangeBase(x int, base int) string {
	if x == 0 {
		return "0"
	}

	result := ""
	for x > 0 {
		digit := x % base
		result = string(rune('0'+digit)) + result
		x = x / base
	}
	return result
}

func main() {
	// Example usage - reading from stdin if needed
	// The function can be called directly with: ChangeBase(8, 3)
	fmt.Println(ChangeBase(8, 3))
	fmt.Println(ChangeBase(8, 2))
	fmt.Println(ChangeBase(7, 2))
}
