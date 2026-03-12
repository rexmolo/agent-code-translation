package main

import "fmt"

func main() {
	// Example usage
	fmt.Println(Tri(3))
}

func Tri(n int) []float64 {
	if n == 0 {
		return []float64{1}
	}
	myTri := []float64{1, 3}
	for i := 2; i <= n; i++ {
		if i%2 == 0 {
			myTri = append(myTri, float64(i)/2+1)
		} else {
			myTri = append(myTri, myTri[i-1]+myTri[i-2]+float64(i+3)/2)
		}
	}
	return myTri
}
