package main

import "fmt"

func GetMaxTriples(n int) int {
	// a[i] = i*i - i + 1 = i*(i-1) + 1
	// For any i, a[i] ≡ 1 (mod 3)
	// So a[i] + a[j] + a[k] ≡ 3 ≡ 0 (mod 3) for any i, j, k
	// Therefore, all C(n,3) triples are valid
	return n * (n - 1) * (n - 2) / 6
}

func main() {
	n := 5
	result := GetMaxTriples(n)
	fmt.Println(result)
}