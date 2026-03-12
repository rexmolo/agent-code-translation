package main

func StartsOneEnds(n int) int {
    if n == 1 {
        return 1
    }
    result := 18
    for i := 0; i < n-2; i++ {
        result *= 10
    }
    return result
}

func main() {
    // Example usage
    // result := StartsOneEnds(1) // returns 1
    // result := StartsOneEnds(2) // returns 18
    // result := StartsOneEnds(3) // returns 180
}
