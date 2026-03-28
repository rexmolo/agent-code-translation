package main

func Intersection(interval1 [2]int, interval2 [2]int) string {
    // Helper function to check if a number is prime
    isPrime := func(num int) bool {
        if num == 1 || num == 0 {
            return false
        }
        if num == 2 {
            return true
        }
        for i := 2; i < num; i++ {
            if num%i == 0 {
                return false
            }
        }
        return true
    }

    // Calculate intersection boundaries
    l := interval1[0]
    if interval2[0] > l {
        l = interval2[0]
    }
    r := interval1[1]
    if interval2[1] < r {
        r = interval2[1]
    }

    length := r - l

    if length > 0 && isPrime(length) {
        return "YES"
    }
    return "NO"
}