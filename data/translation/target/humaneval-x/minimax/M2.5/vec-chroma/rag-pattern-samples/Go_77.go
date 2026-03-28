package main

import (
	"fmt"
	"math"
)

func Iscube(a int) bool {
	if a < 0 {
		a = -a
	}
	x := math.Round(math.Pow(float64(a), 1.0/3.0))
	return int(x)*int(x)*int(x) == a
}

func main() {
	fmt.Println(Iscube(1))
	fmt.Println(Iscube(2))
	fmt.Println(Iscube(-1))
	fmt.Println(Iscube(64))
	fmt.Println(Iscube(0))
	fmt.Println(Iscube(180))
}
