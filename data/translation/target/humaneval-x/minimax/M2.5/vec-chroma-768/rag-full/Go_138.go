package main

// IsEqualToSumEven evaluates whether the given number n can be written as
// the sum of exactly 4 positive even numbers.
// The minimum sum of 4 positive even numbers is 2+2+2+2=8, so any even number
// >= 8 can be expressed this way.
func IsEqualToSumEven(n int) bool {
	return n%2 == 0 && n >= 8
}

func main() {
	// Example usage tests:
	// IsEqualToSumEven(4) == false
	// IsEqualToSumEven(6) == false
	// IsEqualToSumEven(8) == true
}
