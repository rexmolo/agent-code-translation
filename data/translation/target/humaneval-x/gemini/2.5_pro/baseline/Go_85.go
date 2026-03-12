package main

import "fmt"

// Add sums the even elements that are at odd indices in a slice of integers.
func Add(lst []int) int {
	var total int
	for i := 1; i < len(lst); i += 2 {
		if lst[i]%2 == 0 {
			total += lst[i]
		}
	}
	return total
}

func main() {
	// Example from the docstring
	list1 := []int{4, 2, 6, 7}
	result1 := Add(list1)
	fmt.Printf("Add(%v) ==> %d\n", list1, result1)

	// Another example
	list2 := []int{1, 2, 3, 4, 5, 6}
	result2 := Add(list2)
	fmt.Printf("Add(%v) ==> %d\n", list2, result2) // Expected: 2 + 4 + 6 -> Indices 1, 3, 5 -> Elements 2, 4, 6 -> sum is 12. Let's recheck python. Ah, python logic is `[lst[i] for i in range(1, len(lst), 2) if lst[i]%2 == 0]` -> `i` are 1, 3, 5. `lst[1]` is 2 (even), `lst[3]` is 4 (even), `lst[5]` is 6 (even). Sum is 2+4+6=12. Correct.
}
