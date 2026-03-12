package main

import (
	"fmt"
	"math/big"
)

// Modp returns 2^n modulo p.
// This implementation uses math/big.Int to handle potential overflows
// when n and p are large, ensuring correctness for a wider range of inputs.
func Modp(n int, p int) int {
	if p <= 1 {
		// Modulo 1 is always 0, and modulo 0 or negative is undefined/panics.
		// Returning 0 for p <= 1 is a safe default.
		return 0
	}

	// Use big.Int for calculations to prevent overflow.
	bigP := big.NewInt(int64(p))
	ret := big.NewInt(1)
	two := big.NewInt(2)

	// The exponent is n, which is an int.
	// We can use the Exp method for efficient modular exponentiation.
	// It calculates (base^exponent) % modulus
	// Here, base=2, exponent=n, modulus=p
	bigN := big.NewInt(int64(n))
	ret.Exp(two, bigN, bigP)

	return int(ret.Int64())
}

func main() {
	fmt.Println(Modp(3, 5))
	fmt.Println(Modp(1101, 101))
	fmt.Println(Modp(0, 101))
	fmt.Println(Modp(3, 11))
	fmt.Println(Modp(100, 101))
}
