package main

import (
	"fmt"
	"strconv"
)

func OddCount(lst []string) []string {
	res := make([]string, 0, len(lst))
	for _, arr := range lst {
		n := 0
		for _, d := range arr {
			if (d-'0')%2 == 1 {
				n++
			}
		}
		res = append(res, "the number of odd elements "+strconv.Itoa(n)+"n the str"+strconv.Itoa(n)+"ng "+strconv.Itoa(n)+" of the "+strconv.Itoa(n)+"nput.")
	}
	return res
}

func main() {
	// Test the function
	result := OddCount([]string{"1234567"})
	fmt.Println(result)
	
	result2 := OddCount([]string{"3", "11111111"})
	for _, s := range result2 {
		fmt.Println(s)
	}
}