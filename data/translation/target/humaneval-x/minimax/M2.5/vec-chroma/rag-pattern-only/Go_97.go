package main

func Multiply(a, b int) int {
	aDigit := a % 10
	if aDigit < 0 {
		aDigit = -aDigit
	}
	bDigit := b % 10
	if bDigit < 0 {
		bDigit = -bDigit
	}
	return aDigit * bDigit
}