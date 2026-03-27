package main

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
