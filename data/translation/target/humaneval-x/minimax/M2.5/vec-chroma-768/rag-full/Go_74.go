package main

// TotalMatch returns the list that has fewer total characters.
// If both lists have the same total character count, it returns the first list.
func TotalMatch(lst1 []string, lst2 []string) []string {
	totalChars := func(lst []string) int {
		total := 0
		for _, s := range lst {
			total += len(s)
		}
		return total
	}

	l1 := totalChars(lst1)
	l2 := totalChars(lst2)

	if l1 <= l2 {
		return lst1
	}
	return lst2
}