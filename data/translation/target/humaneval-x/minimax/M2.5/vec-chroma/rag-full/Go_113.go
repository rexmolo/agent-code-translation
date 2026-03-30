package main

import "fmt"

func OddCount(lst []string) []string {
	res := []string{}
	for _, arr := range lst {
		n := 0
		for _, d := range arr {
			digit := int(d - '0')
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
		fmt.Println(s)
	}
	
	result2 := OddCount([]string{"3", "11111111"})
	for _, s := range result2 {
		fmt.Println(s)
	}
}
