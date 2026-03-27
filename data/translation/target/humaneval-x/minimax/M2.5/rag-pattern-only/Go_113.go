package main

import (
	"fmt"
	"strconv"
)

func OddCount(lst []string) []string {
	res := make([]string, 0, len(lst))
	for _, arr := range lst {
		n := 0
		for _, r := range arr {
			digit, err := strconv.Atoi(string(r))
			if err != nil {
				continue
			}
			if digit%2 == 1 {
				n++
			}
		}
		res = append(res, fmt.Sprintf("the number of odd elements %dn the str%dnng %d of the %dnput.", n, n, n, n))
	}
	return res
}

func main() {
	// Test cases
	result1 := OddCount([]string{"1234567"})
	for _, s := range result1 {
		fmt.Println(s)
	}

	result2 := OddCount([]string{"3", "11111111"})
	for _, s := range result2 {
		fmt.Println(s)
	}
}
