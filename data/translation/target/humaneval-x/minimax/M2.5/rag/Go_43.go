func PairsSumToZero(l []int) bool {
    for i, l1 := range l {
        for j := i + 1; j < len(l); j++ {
            if l1+l[j] == 0 {
                return true
            }
        }
    }
    return false
}