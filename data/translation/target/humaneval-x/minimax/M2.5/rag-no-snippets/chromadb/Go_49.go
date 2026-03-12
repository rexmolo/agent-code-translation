func Modp(n int, p int) int {
	ret := 1
	for i := 0; i < n; i++ {
		ret = (2 * ret) % p
	}
	return ret
}