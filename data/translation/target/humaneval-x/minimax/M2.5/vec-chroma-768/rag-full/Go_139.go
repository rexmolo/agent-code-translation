package main

import "fmt"

func SpecialFactorial(n int) int {
	fact_i := 1
	special_fact := 1
	for i := 1; i <= n; i++ {
		fact_i *= i
		special_fact *= fact_i
	}
	return special_fact
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(SpecialFactorial(n))
}
