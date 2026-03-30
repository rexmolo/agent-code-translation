package main

import "fmt"

func Exchange(lst1, lst2 []int) string {
	odd := 0
even := 0

for _, i := range lst1 {
	if i%2 == 1 {
		odd++
	}
}

for _, i := range lst2 {
	if i%2 == 0 {
		even++
	}
}

if even >= odd {
	return "YES"
}
return "NO"
}

func main() {
	// Example usage
	lst1 := []int{1, 2, 3, 4}
	lst2 := []int{1, 2, 3, 4}
	fmt.Println(Exchange(lst1, lst2))

	lst1 = []int{1, 2, 3, 4}
	lst2 = []int{1, 5, 3, 4}
	fmt.Println(Exchange(lst1, lst2))
}
