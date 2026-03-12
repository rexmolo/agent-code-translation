func HexKey(num string) int {
	primes := map[rune]bool{
		'2': true,
		'3': true,
		'5': true,
		'7': true,
		'B': true,
		'D': true,
	}

	total := 0
	for _, c := range num {
		if primes[c] {
			total++
		}
	}
	return total
}