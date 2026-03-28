package main

func TriplesSumToZero(l []int) bool {
	n := len(l)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if l[i]+l[j]+l[k] == 0 {
					return true
				}
			}
		}
	}
	return false
}