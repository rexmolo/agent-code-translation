package main

import "fmt"

func Exchange(lst1, lst2 []int) string {
	odd := 0

even := 0

	for _, v := range lst1 {
		if v%2 == 1 {
			odd++
		}
	}

	for _, v := range lst2 {
		if v%2 == 0 {
			even++
		}
	}

	if even >= odd {
		return "YES"
	}
	return "NO"
}

func main() {
	// Test cases from the problem
	fmt.Println(Exchange([]int{1, 2, 3, 4}, []int{1, 2, 3, 4})) // => YES
	fmt.Println(Exchange([]int{1, 2, 3, 4}, []int{1, 5, 3, 4})) // => NO
}