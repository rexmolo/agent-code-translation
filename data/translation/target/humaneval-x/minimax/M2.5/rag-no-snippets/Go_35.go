package main

func MaxElement(l []int) int {
	m := l[0]
	for _, e := range l {
		if e > m {
			m = e
		}
	}
	return m
}
