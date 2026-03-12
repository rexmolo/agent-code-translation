package main

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
	// Test cases
	lst1 := []int{1, 2, 3, 4}
	lst2 := []int{1, 2, 3, 4}
	println(Exchange(lst1, lst2)) // Should print: YES

	lst1 = []int{1, 2, 3, 4}
	lst2 = []int{1, 5, 3, 4}
	println(Exchange(lst1, lst2)) // Should print: NO
}
