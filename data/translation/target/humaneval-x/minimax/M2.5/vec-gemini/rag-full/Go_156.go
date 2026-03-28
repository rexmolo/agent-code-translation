package main

import (
	"fmt"
	"strings"
)

func IntToMiniRoman(number int) string {
	num := []int{1, 4, 5, 9, 10, 40, 50, 90,
		100, 400, 500, 900, 1000}
	sym := []string{"I", "IV", "V", "IX", "X", "XL",
		"L", "XC", "C", "CD", "D", "CM", "M"}
	i := 12
	res := ""
	for number > 0 {
		div := number / num[i]
		number %= num[i]
		for div > 0 {
			res += sym[i]
			div--
		}
		i--
	}
	return strings.ToLower(res)
}

func main() {
	// Test examples
	fmt.Println(IntToMiniRoman(19))
	fmt.Println(IntToMiniRoman(152))
	fmt.Println(IntToMiniRoman(426))
}