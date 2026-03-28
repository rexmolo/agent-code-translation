package main

func RightAngleTriangle(a, b, c int) bool {
	return a*a == b*b+c*c || b*b == a*a+c*c || c*c == a*a+b*b
}

func main() {
	// Example usage:
	println(RightAngleTriangle(3, 4, 5)) // true
	println(RightAngleTriangle(1, 2, 3)) // false
}