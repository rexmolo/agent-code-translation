package main

import "fmt"

func Modp(n int, p int) int {
	ret := 1
	for i := 0; i < n; i++ {
		ret = (2 * ret) % p
	}
	return ret
}

func main() {
	fmt.Println(Modp(3, 5))
	fmt.Println(Modp(1101, 101))
	fmt.Println(Modp(0, 101))
	fmt.Println(Modp(3, 11))
	fmt.Println(Modp(100, 101))
}
