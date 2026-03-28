package main

// TotalMatch compares two lists of strings by their total character count
// and returns the list with fewer characters (or the first list if equal)
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
	}
	return lst2
}
