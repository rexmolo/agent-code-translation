package main

import (
	"fmt"
	"strings"
)

func IntToMiniRoman(number int) string {
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	var builder strings.Builder

	for i := 0; i < len(values); i++ {
		for number >= values[i] {
			builder.WriteString(symbols[i])
			number -= values[i]
		}
	}

	return strings.ToLower(builder.String())
}

func main() {
	fmt.Println(IntToMiniRoman(19))
	fmt.Println(IntToMiniRoman(152))
	fmt.Println(IntToMiniRoman(426))
}
