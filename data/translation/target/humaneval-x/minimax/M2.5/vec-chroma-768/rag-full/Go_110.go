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
