package main

import "fmt"

func GenerateIntegers(a, b int) []int {
	lower := a
	if b < lower {
		lower = b
	}
	if lower < 2 {
		lower = 2
	}

	upper := a
	if b > upper {
		upper = b
	}
	if upper > 8 {
		upper = 8
	}

	var result []int
	for i := lower; i <= upper; i++ {
		if i%2 == 0 {
			result = append(result, i)
		}
	}

	return result
}

func main() {
	fmt.Println(GenerateIntegers(2, 8))   // [2 4 6 8]
	fmt.Println(GenerateIntegers(8, 2))   // [2 4 6 8]
	fmt.Println(GenerateIntegers(10, 14)) // []
}