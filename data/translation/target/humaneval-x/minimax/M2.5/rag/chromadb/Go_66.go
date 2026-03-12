func Digitsum(x string) int {
	sum := 0
	for _, r := range x {
		if r >= 'A' && r <= 'Z' {
			sum += int(r)
		}
	}
	return sum
}