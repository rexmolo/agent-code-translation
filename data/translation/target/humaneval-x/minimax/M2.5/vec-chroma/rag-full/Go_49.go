package main

func Modp(n int, p int) int {
	ret := 1
	for i := 0; i < n; i++ {
		ret = (2 * ret) % p
	}
	return ret
}

func main() {
	// Example usage
	// fmt.Println(Modp(3, 5))    // Output: 3
	// fmt.Println(Modp(1101, 101)) // Output: 2
}