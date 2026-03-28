func RightAngleTriangle(a, b, c int) bool {
	a2 := a * a
	b2 := b * b
	c2 := c * c
	return a2 == b2+c2 || b2 == a2+c2 || c2 == a2+b2
}