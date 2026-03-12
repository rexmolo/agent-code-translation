package main

import "fmt"

func Intersperse(numbers []int, delimeter int) []int {
	if len(numbers) == 0 {
		return []int{}
	}

	result := []int{}

	for i := 0; i < len(numbers)-1; i++ {
		result = append(result, numbers[i])
		result = append(result, delimeter)
	}

	result = append(result, numbers[len(numbers)-1])

	return result
}

func main() {
	// Test cases from docstring
	fmt.Println(Intersperse([]int{}, 4))
	fmt.Println(Intersperse([]int{1, 2, 3}, 4))
}
