func RightAngleTriangle(a, b, c int) bool {
	return a*a == b*b+c*c || b*b == a*a+c*c || c*c == a*a+b*b
}