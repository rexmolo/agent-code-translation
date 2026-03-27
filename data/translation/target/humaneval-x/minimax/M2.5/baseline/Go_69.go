func Search(lst []int) int {
    if len(lst) == 0 {
        return -1
    }

    // Find maximum value in the list
    maxVal := lst[0]
    for _, v := range lst {
        if v > maxVal {
            maxVal = v
        }
    }

    // Create frequency array where index represents the integer
    frq := make([]int, maxVal+1)
    for _, v := range lst {
        frq[v]++
    }

    // Find the greatest integer where frequency >= integer value
    ans := -1
    for i := 1; i < len(frq); i++ {
        if frq[i] >= i {
            ans = i
        }
    }

    return ans
}