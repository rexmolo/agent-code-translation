func IsSorted(lst []int) bool {
    // Count occurrences of each number
    countDigit := make(map[int]int)
    for _, num := range lst {
        countDigit[num]++
    }

    // Check if any number appears more than 2 times
    for _, count := range countDigit {
        if count > 2 {
            return false
        }
    }

    // Check if list is sorted in ascending order
    for i := 1; i < len(lst); i++ {
        if lst[i-1] > lst[i] {
            return false
        }
    }

    return true
}