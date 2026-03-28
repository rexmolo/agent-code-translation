package main

import (
	"fmt"
)

func OddCount(lst []string) []string {
	res := make([]string, 0, len(lst))
	for _, arr := range lst {
		n := 0
		for _, r := range arr {
			digit := int(r - '0')
			if digit%2 == 1 {
				n++
			}
		}
		res = append(res, fmt.Sprintf("the number of odd elements %dn the str%dnng %d of the %dnput.", n, n, n, n))
	}
	return res
}

func main() {
	// Example usage
	result := OddCount([]string{"1234567"})
	for _, s := range result {
		println(s)
	}
	result = OddCount([]string{"3", "11111111"})
	for _, s := range result {
		println(s)
	}
}
