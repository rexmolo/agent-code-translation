package main

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