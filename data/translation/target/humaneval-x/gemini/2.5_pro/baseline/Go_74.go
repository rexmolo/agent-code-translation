package main

// TotalMatch accepts two lists of strings and returns the list that has
// a total number of chars in all its strings less than the other list.
// If the two lists have the same number of chars, the first list is returned.
func TotalMatch(lst1 []string, lst2 []string) []string {
	l1 := 0
	for _, st := range lst1 {
		l1 += len(st)
	}

	l2 := 0
	for _, st := range lst2 {
		l2 += len(st)
	}

	if l1 <= l2 {
		return lst1
	} else {
		return lst2
	}
}
