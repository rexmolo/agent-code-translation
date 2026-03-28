package main

import "fmt"

func MaxElement(l []int) int {
	if len(l) == 0 {
		return 0
	}
	m := l[0]
	for _, e := range l {
		if e > m {
			m = e
		}
	}
	return m
}

func main() {
	fmt.Println(MaxElement([]int{1, 2, 3}))
	fmt.Println(MaxElement([]int{5, 3, -5, 2, -3, 3, 9, 0, 123, 1, -10}))
}