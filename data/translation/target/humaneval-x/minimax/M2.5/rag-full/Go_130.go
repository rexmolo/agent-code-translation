package main

import "fmt"

func Tri(n int) []float64 {
	if n == 0 {
		return []float64{1}
	}
	myTri := make([]float64, n+1)
	myTri[0] = 1
	myTri[1] = 3
	for i := 2; i <= n; i++ {
		if i%2 == 0 {
			myTri[i] = float64(i)/2 + 1
		} else {
			myTri[i] = myTri[i-1] + myTri[i-2] + float64(i+3)/2
		}
	}
	return myTri
}

func main() {
	fmt.Println(Tri(3))
}
