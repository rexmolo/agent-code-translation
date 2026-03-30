package main

import "fmt"

// Tri returns the first n+1 numbers of the Tribonacci sequence.
// Tribonacci is defined as:
//   tri(1) = 3
//   tri(n) = 1 + n/2, if n is even
//   tri(n) = tri(n-1) + tri(n-2) + tri(n+1), if n is odd
func Tri(n int) []float64 {
    if n == 0 {
        return []float64{1}
    }

    myTri := make([]float64, 0, n+1)
    myTri = append(myTri, 1, 3)

    for i := 2; i <= n; i++ {
        if i%2 == 0 {
            myTri = append(myTri, float64(i)/2+1)
        } else {
            val := myTri[i-1] + myTri[i-2] + float64(i+3)/2
            myTri = append(myTri, val)
        }
    }

    return myTri
}

func main() {
    fmt.Println(Tri(3)) // [1 3 2 8]
    fmt.Println(Tri(4)) // [1 3 2 8 3]
}
