func PairsSumToZero(l []int) bool {
	for i := 0; i < len(l); i++ {
		for j := i + 1; j < len(l); j++ {
			if l[i]+l[j] == 0 {
				return true
			}
		}
	}
	return false
}